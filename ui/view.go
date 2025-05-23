package ui

import (
	"fmt"
	"go-gpu/monitor"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderView(
	gpu monitor.GPUStats,
	showInfo, showTemp, showMemory, showEnergy, showHelp bool,
	frame int,
) string {
	// ---------- Estilos dinâmicos ----------
	tempColor := pickColor(gpu.Temperature, 50, 70, 80)
	clockColor := pickColorPct(float64(gpu.ClockCore)/float64(gpu.ClockMax), 0.6, 0.8)
	memColor := pickColorPct(gpu.MemUsed/gpu.MemTotal, 0.7, 0.9)
	utilColor := pickColorPct(float64(gpu.Utilization)/100, 0.7, 0.9)
	powerColor := pickColorPct(gpu.PowerDraw/gpu.PowerLimit, 0.7, 0.9)
	fanColor := pickColor(gpu.FanSpeed, 50, 70, 85)

	var infoCard, tempCard, memCard, powerCard string
	var footer string

	// ---------- Estilos fixos ----------
	block := func(title string, content string) string {
		return lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("8")).
			Padding(0, 1).
			MarginBottom(1).
			Width(cardWidth).
			Render(fmt.Sprintf("%s\n%s", title, content))
	}

	colored := func(text string, color string) string {
		return lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Bold(true).Render(text)
	}

	// ---------- Seções ----------
	if showInfo {
		infoCard = block("🔧 Info Geral",
			fmt.Sprintf("%s %s",
				LabelStyle.Render("Nome:"), colored(gpu.Name, "15"),
			))
	} else {
		infoCard = collapsedCard("🔧 Info Geral")
	}

	if showTemp {
		tempCard = block("🔥 Temperatura & Clock",
			fmt.Sprintf(
				"%s %s\n%s\n%s %s\n%s",
				LabelStyle.Render("Temp:"),
				colored(fmt.Sprintf("%d°C", gpu.Temperature), tempColor),
				drawBar(float64(gpu.Temperature), 100, 20), // Temperatura entre 0–100%
				LabelStyle.Render("Clock GPU:"),
				colored(fmt.Sprintf("%d / %d MHz", gpu.ClockCore, gpu.ClockMax), clockColor),
				drawBar(float64(gpu.ClockCore), float64(gpu.ClockMax), 20),
			),
		)
	} else {
		tempCard = collapsedCard("🔥 Temperatura & Clock")
	}

	if showMemory {
		memCard = block("💾 Memória & Uso",
			fmt.Sprintf(
				"%s %s\n%s %s",
				LabelStyle.Render("Memória:"),
				colored(fmt.Sprintf("%.0f / %.0f MB", gpu.MemUsed, gpu.MemTotal), memColor),
				LabelStyle.Render("Utilização:"),
				colored(fmt.Sprintf("%d%%", gpu.Utilization), utilColor),
			),
		)
	} else {
		memCard = collapsedCard("💾 Memória & Uso")
	}

	if showEnergy {
		powerCard = block("⚡ Fan & Energia",
			fmt.Sprintf(
				"%s %s %s\n%s %s",
				LabelStyle.Render("Fan Speed:"),
				colored(fmt.Sprintf("%d%%", gpu.FanSpeed), fanColor),
				getFanSpinner(frame), // ← aqui o giro!
				LabelStyle.Render("Power:"),
				colored(fmt.Sprintf("%.1f / %.1f W", gpu.PowerDraw, gpu.PowerLimit), powerColor),
			),
		)
	} else {
		powerCard = collapsedCard("⚡ Fan & Energia")
	}

	if showHelp {
		footer = FooterStyle.Render("I = Info | T = Temp e Clock | M = Memória | E = Energia | Q = Sair | H = Ajuda")
	} else {
		footer = FooterStyle.Render("H = AJUDA | Q = SAIR")
	}

	return lipgloss.NewStyle().Padding(1, 2).Render(
		fmt.Sprintf("%s\n\n%s\n%s\n%s\n%s",
			infoCard,
			tempCard,
			memCard,
			powerCard,
			footer,
		),
	)
}

func getFanSpinner(frame int) string {
	frames := []string{"|", "/", "-", "\\"}
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("7")).
		Render(frames[frame%len(frames)])
}

func collapsedCard(title string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("8")).
		Italic(true).
		MarginBottom(1).
		Width(cardWidth).
		Render("▼ " + title + " (oculto)")
}

func pickColor(value int, greenMax, yellowMax, redMax int) string {
	switch {
	case value < greenMax:
		return "10" // verde
	case value < yellowMax:
		return "3" // amarelo
	case value < redMax:
		return "1" // vermelho
	default:
		return "9" // magenta/erro
	}
}

func pickColorPct(pct float64, warn, danger float64) string {
	switch {
	case pct < warn:
		return "10" // verde
	case pct < danger:
		return "3" // amarelo
	default:
		return "1" // vermelho
	}
}

func drawBar(used float64, total float64, width int) string {
	if total <= 0 {
		total = 1
	}
	percent := used / total
	filled := int(percent * float64(width))
	if filled < 0 {
		filled = 0
	}
	if filled > width {
		filled = width
	}

	fg := lipgloss.Color("10")
	if percent > 0.7 {
		fg = lipgloss.Color("3")
	}
	if percent > 0.85 {
		fg = lipgloss.Color("1")
	}

	bar := strings.Repeat("█ ", filled) + strings.Repeat("░ ", width-filled)
	return lipgloss.NewStyle().Foreground(fg).Bold(true).Render(bar)
}

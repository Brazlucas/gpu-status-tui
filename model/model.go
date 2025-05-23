package model

import (
	"time"

	"go-gpu/monitor"
	"go-gpu/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type gpuStats = monitor.GPUStats

type Model struct {
	GPU        gpuStats
	Err        error
	ShowInfo   bool
	ShowTemp   bool
	ShowMemory bool
	ShowEnergy bool
	ShowHelp   bool
	Frame      int
}

type tickMsg time.Time
type dataMsg time.Time

func InitialModel() Model {
	return Model{
		ShowInfo:   true,
		ShowTemp:   true,
		ShowMemory: true,
		ShowEnergy: true,
		ShowHelp:   false,
	}
}

func tickVisual() tea.Cmd {
	return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func tickData() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return dataMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tickData(), tickVisual())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tickMsg:
		m.Frame++
		return m, tickVisual()

	case dataMsg:
		info, err := monitor.GetGPUInfo()
		if err != nil {
			m.Err = err
		} else {
			m.GPU = info
		}
		return m, tickData()

	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "i" {
			m.ShowInfo = !m.ShowInfo
		}
		if msg.String() == "t" {
			m.ShowTemp = !m.ShowTemp
		}
		if msg.String() == "m" {
			m.ShowMemory = !m.ShowMemory
		}
		if msg.String() == "e" {
			m.ShowEnergy = !m.ShowEnergy
		}
		if msg.String() == "h" {
			m.ShowHelp = !m.ShowHelp
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.Err != nil {
		return "Erro: " + m.Err.Error()
	}
	return ui.RenderView(m.GPU, m.ShowInfo, m.ShowTemp, m.ShowMemory, m.ShowEnergy, m.ShowHelp, m.Frame)
}

package ui

import "github.com/charmbracelet/lipgloss"

const cardWidth = 50

var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("12")).
			Bold(true).
			Border(lipgloss.NormalBorder()).
			Padding(0, 2).
			Align(lipgloss.Center).
			Width(36)

	LabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("8")).
			Width(14)

	ValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Bold(true)

	SectionStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			MarginTop(1)

	FooterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("7")).
			MarginTop(2).
			Italic(true)
)

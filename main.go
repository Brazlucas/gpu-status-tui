package main

import (
	"fmt"

	"go-gpu/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model.InitialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Erro ao iniciar:", err)
	}
}

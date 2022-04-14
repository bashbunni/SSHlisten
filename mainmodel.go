package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	comment viewport.Model
}

func NewModel(comment string) *Model {
	commentViewport := viewport.New(10, 10)
	commentViewport.SetContent(comment)
	return &Model{comment: commentViewport}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.comment.View()
}

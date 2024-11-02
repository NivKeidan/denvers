package main

import (
	"fmt"
	"os"

	"denvers/critter"
	"denvers/log"
	"denvers/ui"
	"denvers/world"

	tea "github.com/charmbracelet/bubbletea"
)

// This hold the current state of the application
type model struct {
	denver    *critter.Denver
	showStats bool
	position  *world.Position
	world     *world.Map
}

func (m *model) toggleShowStats() {
	m.showStats = !m.showStats
}
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c": // Quit on 'q' or Ctrl+C
			return m, tea.Quit
		case "w", "up":
			m.position.Up()
		case "s", "down":
			m.position.Down()
		case "a", "left":
			m.position.Left()
		case "d", "right":
			m.position.Right()
		case "tab":
			m.toggleShowStats()
		}
	}
	return m, nil
}

// View renders the TUI on each update
func (m model) View() string {
	header := ui.CreateElement(ui.Header, m.denver.Name)
	characterInfo := ui.CreateElement(ui.Basic, m.position)

	if m.showStats {
		statsView := ui.CreateElement(ui.StatsView, m.denver.Health, m.denver.MaxHealth, m.denver.Strength, m.denver.Speed)
		characterInfo = characterInfo.AppendRight(statsView)
	}

	worldView := ""
	for i, row := range m.world.Tiles {
		for j, tile := range row {
			if m.position.X() == i && m.position.Y() == j {
				worldView += "D"
				continue
			}
			worldView += string(tile)
		}
		worldView += "\n"
	}

	return fmt.Sprintf("%s\n%s\n\n%s", ui.PrependWithSpaces(2, header), ui.PrependWithSpaces(2, characterInfo), ui.PrependWithSpaces(2, ui.Element(worldView)))
}

func main() {
	defer log.Close()
	log.Log("hi")

	world, startingPosition := world.DefaultMap()
	m := model{
		denver:   critter.NewDenver(),
		position: startingPosition,
		world:    world,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
		os.Exit(1)
	}
}

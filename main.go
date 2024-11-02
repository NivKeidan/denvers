package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Position struct {
	x int
	y int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Denver struct {
	Name      string
	Strength  int
	Speed     int
	Health    int
	MaxHealth int
	Position  Position
}

func (c *Denver) moveUp() {
	c.Position.y++
}

func (c *Denver) moveDown() {
	c.Position.y--
}

func (c *Denver) moveLeft() {
	c.Position.x--
}

func (c *Denver) moveRight() {
	c.Position.x++
}

func NewDenver() *Denver {
	// TODO make random
	return &Denver{
		Name:      "MyDenver",
		Strength:  10,
		Speed:     10,
		Health:    10,
		MaxHealth: 10,
	}
}

// This hold the current state of the application
type model struct {
	denver    *Denver
	showStats bool
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
			m.denver.moveUp()
		case "s", "down":
			m.denver.moveDown()
		case "a", "left":
			m.denver.moveLeft()
		case "d", "right":
			m.denver.moveRight()
		case "tab":
			m.toggleShowStats()
		}
	}
	return m, nil
}

// View renders the TUI on each update
func (m model) View() string {
	if m.showStats {
		return fmt.Sprintf("=== Critter Stats ===\n\nHealth: %d/%d\nStrength: %d\nSpeed: %d\n\nPress 'TAB' to return.", m.denver.Health, m.denver.MaxHealth, m.denver.Strength, m.denver.Speed)
	}

	return fmt.Sprintf("Denver: %s\nPosition %s\n\nUse a,w,s,d or arrows to move around\nPress TAB to show Denver stats\nPress q to quit.", m.denver.Name, m.denver.Position)
}

func main() {
	p := tea.NewProgram(model{denver: NewDenver()})
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
		os.Exit(1)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func getMaxRight(s string) int {
	max := 0
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		l := len(scanner.Text())
		if l > max {
			max = l
		}
	}
	return max
}

func padWithSpaces(index int, s string) string {
	if len(s) >= index {
		return s
	}
	padding := strings.Repeat(" ", index-len(s))
	return fmt.Sprintf("%s%s", s, padding)
}

func addToRight(base, addon string) string {
	maxRight := getMaxRight(base)
	space := 4
	minStartindIndex := maxRight + space

	baseScanner := bufio.NewScanner(strings.NewReader(base))
	addonScanner := bufio.NewScanner(strings.NewReader(addon))
	finalString := ""

	for baseScanner.Scan() {
		currentLine := baseScanner.Text()
		if addonScanner.Scan() {
			currentLine = padWithSpaces(minStartindIndex, currentLine)
			currentLine += addonScanner.Text()
		}

		finalString += fmt.Sprintf("%s\n", currentLine)
	}

	for addonScanner.Scan() {
		currentLine := padWithSpaces(minStartindIndex, "")
		currentLine += addonScanner.Text()
		finalString += fmt.Sprintf("%s\n", currentLine)
	}

	return finalString
}

// View renders the TUI on each update
func (m model) View() string {
	header := fmt.Sprintf(`
  ========================================
              Welcome, %s!
  ========================================`, m.denver.Name)

	mainView := fmt.Sprintf(`
  Position:    %s
  ----------------------------------------
  Use WASD or arrow keys to move around.
  Press TAB to view stats.
  Press Q to quit.
`, m.denver.Position)

	statsView := fmt.Sprintf(`
  ┌──────────────────────┐
  │    Denver Stats      │
  ├──────────────────────┤
  │ Health:    %d/%d     │
  │ Strength:  %d        │
  │ Speed:     %d        │
  └──────────────────────┘
`, m.denver.Health, m.denver.MaxHealth, m.denver.Strength, m.denver.Speed)

	finalView := mainView
	if m.showStats {
		finalView = addToRight(mainView, statsView)
	}

	// Return the main view with the stats box on the right
	return fmt.Sprintf("%s\n%s", header, finalView)
}

func main() {
	p := tea.NewProgram(model{denver: NewDenver()}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
		os.Exit(1)
	}
}

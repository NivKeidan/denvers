package ui

type Template string

func (t Template) String() string {
	return string(t)
}

func NewTemplate(s string) Template {
	return Template(s)
}

func NewTemplateWithNewLine(s string) Template {
	return Template(s[1:])
}

const Header = `
========================================
            Welcome, %s!
========================================`

const Basic = `
Position:    %s
----------------------------------------
Use WASD or arrow keys to move around.
Press TAB to view stats.
Press Q to quit.
`

const StatsView = `
┌──────────────────────┐
│    Denver Stats      │
├──────────────────────┤
│ Health:    %d/%d     │
│ Strength:  %d        │
│ Speed:     %d        │
└──────────────────────┘`

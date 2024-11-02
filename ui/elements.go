package ui

import (
	"bufio"
	"fmt"
	"strings"
)

type Element string

func CreateElement(template Template, args ...any) Element {
	return Element(fmt.Sprintf(template.String(), args...))
}

func (e Element) String() string {
	return string(e)
}

func (e Element) AppendRight(addon Element) Element {
	maxRight := getMaxRight(e)
	space := 4
	minStartindIndex := maxRight + space

	baseScanner := bufio.NewScanner(strings.NewReader(e.String()))
	addonScanner := bufio.NewScanner(strings.NewReader(addon.String()))
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

	return Element(finalString)
}

func getMaxRight(e Element) int {
	max := 0
	scanner := bufio.NewScanner(strings.NewReader(e.String()))
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

func PrependWithSpaces(count int, e Element) string {
	finalString := ""
	scanner := bufio.NewScanner(strings.NewReader((e.String())))
	for scanner.Scan() {
		finalString += fmt.Sprintf("%s%s\n", strings.Repeat(" ", count), scanner.Text())
	}
	return finalString
}

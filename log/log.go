package log

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var f *os.File

func Start() (err error) {
	f, err = tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	return err
}

func Close() {
	f.Close()
}

func Log(s string) (err error) {
	if _, err = f.Write([]byte(s)); err != nil {
		return err
	}
	if _, err = f.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}

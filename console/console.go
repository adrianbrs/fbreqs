package console

import (
	"github.com/spf13/viper"
)

// Console settings type
type consoleType struct {
	Separator  string
	Width      int
	LeftOffset int
}

var consoleConfig *consoleType

// Init console configuration
func Init() {
	consoleConfig = &consoleType{
		Separator:  "-",
		Width:      viper.GetInt("app.shell.width"),
		LeftOffset: viper.GetInt("app.shell.offset"),
	}
}

// SetWidth sets console width
func SetWidth(width int) {
	consoleConfig.Width = width
}

// GetWidth returns console width
func GetWidth() int {
	return consoleConfig.Width
}

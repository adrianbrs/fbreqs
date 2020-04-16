package console

import (
	"strings"

	"github.com/fatih/color"
)

// Info ...
func Info(msg string, args ...interface{}) {
	color.New(color.BgHiCyan).Print(" ")
	color.New(color.FgCyan).Printf(LeftW("[INFO] ", consoleConfig.LeftOffset)+msg+"\n", args...)
}

// Warn ...
func Warn(msg string, args ...interface{}) {
	color.New(color.FgBlack, color.BgHiYellow).Print(LeftW(" [WARN]", consoleConfig.LeftOffset))
	color.New(color.FgHiYellow).Printf(" "+msg+"\n", args...)
}

// Error ...
func Error(msg string, args ...interface{}) {
	color.New(color.BgHiRed, color.FgBlack).Print(LeftW(" [ERROR]", consoleConfig.LeftOffset))
	color.New(color.FgRed).Printf(" "+msg+"\n", args...)
}

// Success ...
func Success(msg string, args ...interface{}) {
	color.New(color.BgHiGreen, color.FgBlack).Print(" ")
	color.New(color.FgGreen).Printf(LeftW("[SUCCESS] ", consoleConfig.LeftOffset)+msg+"\n", args...)
}

// Log ...
func Log(msg string, args ...interface{}) {
	color.New(color.BgHiWhite).Print(" ")
	color.New(color.FgWhite).Printf(LeftW("[LOG] ", consoleConfig.LeftOffset)+msg+"\n", args...)
}

// Separator ...
func Separator() {
	width := consoleConfig.Width + consoleConfig.LeftOffset
	repeat := (width / len(consoleConfig.Separator)) + 1
	color.New(color.FgHiBlack).Println(strings.Repeat(consoleConfig.Separator, repeat)[:width])
}

// Primary ...
func Primary(msg string, args ...interface{}) {
	width := consoleConfig.Width + consoleConfig.LeftOffset
	color.New(color.BgBlack, color.FgWhite).Printf(CenterW(msg, width)+"\n", args...)
}

// Logger any printf function
type Logger func(msg string, args ...interface{})

// Logger variables
var (
	DefaultLogger = Log
	InfoLogger    = Info
	WarnLogger    = Warn
	ErrorLogger   = Error
	SuccessLogger = Success
	PrimaryLogger = Primary
)

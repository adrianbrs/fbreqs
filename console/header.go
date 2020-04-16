package console

import (
	"fmt"
	"strings"
)

type headerType struct {
	TopBar    string
	MiddleBar string
	BottomBar string
	Values    []string
	Align     string
	Logger    Logger
}

func Header() *headerType {
	return &headerType{
		TopBar:    "=",
		MiddleBar: "=",
		BottomBar: "=",
		Align:     "center",
		Logger:    DefaultLogger,
	}
}

func (h *headerType) Left(str string) *headerType {
	h.Align = "left"
	return h
}

func (h *headerType) Right(str string) *headerType {
	h.Align = "right"
	return h
}

func (h *headerType) Center(str string) *headerType {
	h.Align = "center"
	return h
}

func (h *headerType) SetTopBar(str string) *headerType {
	h.TopBar = str
	return h
}

func (h *headerType) SetMiddleBar(str string) *headerType {
	h.MiddleBar = str
	return h
}

func (h *headerType) SetBottomBar(str string) *headerType {
	h.BottomBar = str
	return h
}

func (h *headerType) SetBar(str string) *headerType {
	h.TopBar = str
	h.MiddleBar = str
	h.BottomBar = str
	return h
}

func (h *headerType) SetValues(str ...string) *headerType {
	h.Values = str
	return h
}

func (h *headerType) AppendValues(str ...string) *headerType {
	h.Values = append(h.Values, str...)
	return h
}

func (h *headerType) SetLogger(logger Logger) *headerType {
	h.Logger = logger
	return h
}

func (h *headerType) Render() {
	// Repeating bar count
	repeatTop := (consoleConfig.Width / len(h.TopBar)) + 1
	repeatBot := (consoleConfig.Width / len(h.BottomBar)) + 1

	// Add top bar
	h.Logger(strings.Repeat(h.TopBar, repeatTop)[:consoleConfig.Width])

	// Add values
	for _, value := range h.Values {
		width := consoleConfig.Width - len(h.MiddleBar)*2
		centeredValue := CenterW(value, width)
		h.Logger(fmt.Sprintf("%s%s%s", h.MiddleBar, centeredValue, h.MiddleBar))
	}

	// Add bottom bar
	h.Logger(strings.Repeat(h.BottomBar, repeatBot)[:consoleConfig.Width])
}

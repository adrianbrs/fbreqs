package console

import (
	"fmt"
)

// Right align text right
func Right(str string, args ...interface{}) string {
	str = fmt.Sprintf(str, args...)
	return RightW(str, consoleConfig.Width)
}

// RightW align text right
func RightW(str string, width int) string {
	return fmt.Sprintf(fmt.Sprintf("%%%ds", width), str)
}

// Left align text left
func Left(str string, args ...interface{}) string {
	str = fmt.Sprintf(str, args...)
	return LeftW(str, consoleConfig.Width)
}

// LeftW align text left
func LeftW(str string, width int) string {
	return fmt.Sprintf(fmt.Sprintf("%%-%ds", width), str)
}

// Center align text center
func Center(str string, args ...interface{}) string {
	str = fmt.Sprintf(str, args...)
	return CenterW(str, consoleConfig.Width)
}

// CenterW align text center
func CenterW(str string, width int) string {
	return fmt.Sprintf("%[1]*s", -width, fmt.Sprintf("%[1]*s", (width+len(str))/2, str))
}

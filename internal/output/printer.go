package output

import "fmt"

func Success(message string) {
	fmt.Println(successStyle.Render("✓ " + message))
}

func Error(message string) {
	fmt.Println(errorStyle.Render("✗ " + message))
}

func Info(message string) {
	fmt.Println(infoStyle.Render(message))
}

func Warning(message string) {
	fmt.Println(warningStyle.Render("! " + message))
}

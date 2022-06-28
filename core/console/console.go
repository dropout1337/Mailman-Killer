package console

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/admin100/util/console"
)

var (
	PrimaryColor   string = "\x1b[38;5;147m" // Color customization
	SecondaryColor string = "\x1b[0m"
)

func Clear() {
	c := exec.Command("cmd", "/c", "cls") // Only on windows
	c.Stdout = os.Stdout
	c.Run()
}

func SetTitle(title string) {
	console.SetConsoleTitle(title) // Set the console title
}

func Log(message string, newline bool) {
	content := fmt.Sprintf("[%v%s%v] %v", PrimaryColor, time.Now().Format("15:04:05"), SecondaryColor, message) // Just a nice clean logging
	if newline {
		fmt.Println(content) // Print the content with a new line
	} else {
		fmt.Print(content) // Print without a new line
	}
}

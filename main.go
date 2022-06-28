package main

import (
	"fmt"
	"strings"

	"emails.go/core/console"
	"emails.go/core/sender"
)

var (
	target  string
	threads int

	refinedInput string
	useRefined   bool = false

	loopInput string
	loop      bool = false
)

func main() {
	console.Clear()
	console.SetTitle("[emails.go] - fastest email spammer on earth!") // Set the console title

	// Use refined nodes?
	console.Log(fmt.Sprintf("Refined%v:%v ", console.PrimaryColor, console.SecondaryColor), false)
	fmt.Scanln(&refinedInput)

	if strings.Contains(refinedInput, "y") {
		useRefined = true
	}

	// Constantly loop send emails?
	console.Log(fmt.Sprintf("Loop%v:%v ", console.PrimaryColor, console.SecondaryColor), false)
	fmt.Scanln(&loopInput)

	if strings.Contains(loopInput, "y") {
		loop = true
	}

	// Ask for the target email
	console.Log(fmt.Sprintf("Email%v:%v ", console.PrimaryColor, console.SecondaryColor), false)
	fmt.Scanln(&target)

	// Ask for the amount of threads
	console.Log(fmt.Sprintf("Threads%v:%v ", console.PrimaryColor, console.SecondaryColor), false)
	fmt.Scanln(&threads)

	// Send email
	fmt.Println() // Print a blank new line
	if !loop {
		sender.SendEmails(target, threads, useRefined)
	} else {
		for {
			sender.SendEmails(target, threads, useRefined)
		}
	}
}

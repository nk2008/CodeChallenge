/* <!DOCTYPE html>
<html lang="en">
<head>
    <title>Display Random Words</title>
</head>
<body>
  	<form id="form1" runat="server" name="RandomWords.go">
        <label>Enter Words (as comma separated): <textarea rows="3" cols="30" name="listOfWords" id="listOfWords"></textarea>
        </label>
        <input id="btnStart" type="submit" value="Start" OnClick="Start()">
        <input id="btnClear" type="reset" value="Clear">
    </form>
</body>
</html> */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	input := []string{"One", "Two", "Three", "Four", "Five"}
	if len(input) == 0 {
		fmt.Println("User has not entered any words to continue, process halted.")
		return
	}
	commandRunning := false
	commandTimeRemaining := time.Duration(0)

	totalCycleTime := time.Duration(60)                             //Total Cycle Time
	displayTimer := time.Duration(int(totalCycleTime) / len(input)) // how long to display each word
	cutoffTimer := time.Duration(0)                                 // when to stop timer and reset/halt the process

	// Trigger the method immediately
	for t := range time.Tick(1 * time.Second) {
		go Start(t)
		// Command Running is logged here
		commandRunning = true
		fmt.Println("Command Running: ", commandRunning)

		cutoffTimer += displayTimer
		// Command Time Remaining is logged here
		commandTimeRemaining = (totalCycleTime - cutoffTimer) * time.Second
		fmt.Println("Command Time Remaining: ", commandTimeRemaining)

		// Wait for x interval based on words slice length and total interval of totalCycleTime seconds
		time.Sleep(displayTimer * time.Second)

		// Stop and reset the process when cutofftimer has met total interval of totalCycleTime seconds or more
		if cutoffTimer >= totalCycleTime {
			// Command Running is logged here
			commandRunning = false
			fmt.Println("Command Running: ", commandRunning)
			break
		}
	}
}

func Start(t time.Time) {
	words := []string{"One", "Two", "Three", "Four", "Five"}
	// Using Random method from Math Package to get random word to display
	fmt.Println("Current Word Selected: ", words[rand.Intn(len(words))])
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"

	hook "github.com/robotn/gohook"
)

var isRunning bool = false
var keysToPressMap = map[byte]bool{
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
}
var delayBetweenKeys = time.Millisecond * 50
var delayBetweenCycles = time.Millisecond * 50

func keysToPressString() string {
	keys := make([]byte, 0, len(keysToPressMap))
	for key := range keysToPressMap {
		keys = append(keys, key)
	}
	return string(keys)
}

func main() {
	fmt.Println("Starting macro engine...")
	go addEvents()

	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) != 2 {
			fmt.Println("Current keys are:", keysToPressString())
			continue
		}

		var operation = input[0]
		var keyboardKey = input[1]

		switch operation {
		case '+':
			keysToPressMap[keyboardKey] = true
			fmt.Printf("Added key %s to macro\n", string(keyboardKey))
		case '-':
			delete(keysToPressMap, keyboardKey)
			fmt.Printf("Removed key %s from macro\n", string(keyboardKey))
		default:
			fmt.Println("Current keys are:", keysToPressString())
		}
	}
}

func pressKeysAndCycle() {
	for {
		if !isRunning {
			return
		}

		for key := range keysToPressMap {
			robotgo.KeyTap(string(key))
			time.Sleep(delayBetweenKeys)
		}
		time.Sleep(delayBetweenCycles)
	}
}

func startMacro() {
	go pressKeysAndCycle()
	// The above will press keys fairly robotically/rhythmically
	// For specialized behavior, add new recursive functions below
}

func addEvents() {
	fmt.Println("--- Please press alt + q to enable macro        ---")
	fmt.Println("---              alt + w to pause macro         ---")
	fmt.Println("--- Add keys with +<key> and remove with -<key> ---")
	fmt.Println("--- List current keys by pressing enter         ---")
	fmt.Println("Current keys are:", keysToPressString())
	hook.Register(hook.KeyDown, []string{"q", "alt"}, func(e hook.Event) {
		if !isRunning {
			isRunning = true
			go startMacro()
		}
	})

	hook.Register(hook.KeyDown, []string{"w", "alt"}, func(e hook.Event) {
		isRunning = false
	})

	s := hook.Start()
	<-hook.Process(s)
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"

	hook "github.com/robotn/gohook"
)

var isRunning bool = false

func main() {
	fmt.Println("Starting Path of Exile macro engine...")
	go addEvents()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs // Wait for an interrupt signal to gracefully shut down

	fmt.Println("Exiting application...")
}

func press1() {
	if !isRunning {
		return
	}
	robotgo.KeyTap("1")
	time.Sleep(time.Second * 7)
	go press1()
}

func pressr() {
	if !isRunning {
		return
	}
	robotgo.KeyTap("r")
	time.Sleep(time.Millisecond * 3190)
	go pressr()
}

func presse() {
	if !isRunning {
		return
	}
	robotgo.KeyTap("e")
	time.Sleep(time.Millisecond * 3190)
	go presse()
}

func startMacro() {
	go press1()
	time.Sleep(time.Millisecond * 50) // Delay to avoid using two actions at once
	go pressr()
	time.Sleep(time.Millisecond * 1595) // Delay to offset r and e skills, so they are alternating
	go presse()
}

func addEvents() {
	fmt.Println("--- Please press alt + c to enable macro ---")
	fmt.Println("---              alt + v to pause macro  ---")
	hook.Register(hook.KeyDown, []string{"c", "alt"}, func(e hook.Event) {
		if !isRunning {
			isRunning = true
			go startMacro()
		}
	})

	hook.Register(hook.KeyDown, []string{"v", "alt"}, func(e hook.Event) {
		isRunning = false
	})

	s := hook.Start()
	<-hook.Process(s)
}

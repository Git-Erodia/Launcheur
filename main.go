package main

import (
	_ "embed"
	"erodialuncher/src/Zwuiix"
	"github.com/kbinani/win"
	"os"
	"os/signal"
	"syscall"
)

var c = make(chan os.Signal)

func main() {
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	handler()
}

func handler() {
	win.SetConsoleTitle("")
	Zwuiix.Luncher{}.Start()
}

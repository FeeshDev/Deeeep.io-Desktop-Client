package main

import (
	"log"
	"github.com/jchv/go-webview2"
)

func main() {
	debug := true
	w := webview2.New(debug)
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview2.HintFixed)
	w.Navigate("https://deeeep.io")
	w.Run()
}
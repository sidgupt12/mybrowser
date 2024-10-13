package main

import (
	webview "github.com/webview/webview_go"
)

func main() {
	//First we need to create a webview window
	//debug mode is enabled
	debug := true

	//create the window
	w := webview.New(debug)

	//defer the destruction of the window
	defer w.Destroy()

	//set the title of the window
	w.SetTitle("Skibbidi Browser")

	//set the size of the window
	//1000 is the width, 800 is the height
	//webview.HintNone is the hint for the size
	//hint means that the window will be resizable
	w.SetSize(1000, 800, webview.HintNone)

	//load the url
	w.Navigate("https://www.duckduckgo.com")

	//run the window
	w.Run()
}

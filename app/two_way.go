package main

import (
	"gocore/core"
	"gocore/ui"
)

func TwoWay() {
	name := core.NewObservable("Your Name")

	label := ui.NewTextField().
		Padding(8).Background("lightblue").
		Border("2px solid grey")
	label.BindTo(name)

	sep := ui.NewDiv().Height(10)

	input := ui.NewTextField().
		Padding(8).Background("lightblue").
		Border("2px solid grey")
	input.BindTo(name)

	win := ui.NewWindow()
	win.Add(input, sep, label)
	win.Run()
}

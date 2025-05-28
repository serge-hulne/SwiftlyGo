package main

import (
	"fmt"
	"gocore/core"
	"gocore/ui"
)

func OneWay() {
	source := core.NewObservable("")

	in := ui.NewTextField()
	in.BindTo(source)

	out := ui.NewTextField()

	btn := ui.NewButton("Copy")
	btn.OnClick(func() {
		fmt.Println("Clicked:", source.Get())
		out.SetText(source.Get())
	})

	w := ui.NewWindow()
	w.Add(in, out, btn)
	w.Run()
}

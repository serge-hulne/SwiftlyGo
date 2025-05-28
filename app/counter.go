package main

import (
	"fmt"
	"gocore/core"
	"gocore/ui"
)

func counter() {
	count := core.NewObservable(0)

	// 1st formulation

	label := ui.NewLabel("").
		Padding(8).
		Border("1px solid #ccc").
		Center()

	count.Subscribe(func(v int) {
		label.SetText(fmt.Sprintf("Count: %d", v))
	})

	// 2nd formulation

	// Label bound to count

	// label := ui.NewLabel("").
	//		Padding(8).
	//		Border("1px solid #ccc").
	//		Center()
	//	label.BindTo(core.Derive(func() string {
	//		return fmt.Sprintf("Count: %d", count.Get())
	//	}))

	// Increment button
	inc := ui.NewButton("➕").
		Padding(8).
		Background("lightgreen")
	inc.OnClick(func() {
		count.Set(count.Get() + 1)
	})

	// Decrement button
	dec := ui.NewButton("➖").
		Padding(8).
		Background("salmon")
	dec.OnClick(func() {
		count.Set(count.Get() - 1)
	})

	row := ui.NewHBox(dec, label, inc).Padding(12)

	win := ui.NewWindow()
	win.Add(row)
	win.Run()
}

package main

import (
	"fmt"
	"gocore/core"
	"gocore/ui"
	"strings"
)

type Task struct {
	Label   string
	Checked *core.Observable[bool]
}

func todo() {
	win := ui.NewWindow()

	taskText := core.NewObservable("")
	tasks := core.NewObservable([]*Task{})

	// Input
	input := ui.NewTextField()
	input.BindTo(taskText)
	input.Padding(8)

	// Add Button
	addBtn := ui.NewButton("Add").Padding(8)

	// Task list container
	list := ui.NewVBox().Padding(8)

	// Header
	header := ui.NewHBox()
	header.Add(input, addBtn)

	// Add button logic (no derived, no subscribe)
	addBtn.OnClick(func() {
		text := strings.TrimSpace(taskText.Get())
		if text == "" {
			return
		}

		task := &Task{
			Label:   text,
			Checked: core.NewObservable(false),
		}
		taskText.Set("")

		current := tasks.Get()
		tasks.Set(append(current, task))

		checkbox := ui.NewCheckBox()
		checkbox.SetChecked(false)
		checkbox.OnChange(func(checked bool) {
			task.Checked.Set(checked)
			tasks.Set(tasks.Get())
		})

		label := ui.NewLabel(task.Label)

		row := ui.NewHBox()
		row.Add(checkbox, label)
		list.Add(row)
	})

	remaining := core.Derive(func() string {
		count := 0
		for _, task := range tasks.Get() {
			if !task.Checked.Get() {
				count++
			}
		}
		return fmt.Sprintf("🕒 Remaining tasks: %d", count)
	})

	remainingLabel := ui.NewLabel("")
	remainingLabel.BindTo(remaining)

	// Derived count of total tasks
	taskCount := core.Derive(func() string {
		return fmt.Sprintf("📦 Total tasks: %d", len(tasks.Get()))
	})

	countLabel := ui.NewLabel("")
	countLabel.BindTo(taskCount)

	// Main layout
	layout := ui.NewVBox(header, list, countLabel, remainingLabel).Padding(16)

	win.Add(layout)
	win.Run()
}

package main

import (
	"gocore/core"
	"gocore/shared"
	"gocore/ui"
	"strconv"
)

// Cell represents a reactive cell in the Game of Life.

type Cell struct {
	Alive *core.Observable[bool]
	View  shared.Widget
}

func newCell() *Cell {
	alive := core.NewObservable(false)

	view := ui.NewDiv().Width(20).Height(20).Border("1px solid #ccc")
	view.BindStyle(core.Derive(func() map[string]string {
		if alive.Get() {
			return map[string]string{"background": "black"}
		}
		return map[string]string{"background": "white"}
	}))

	view.OnClick(func() {
		alive.Set(!alive.Get())
	})

	return &Cell{Alive: alive, View: view}
}

func ConwayReactive() {
	const rows, cols = 20, 20
	grid := make([][]*Cell, rows)

	for i := range grid {
		grid[i] = make([]*Cell, cols)
		for j := range grid[i] {
			grid[i][j] = newCell()
		}
	}

	container := ui.NewVBox()

	for _, row := range grid {
		h := ui.NewHBox()
		for _, cell := range row {
			h.Add(cell.View)
		}
		container.Add(h)
	}

	step := func() {
		next := make([][]bool, rows)
		for i := range next {
			next[i] = make([]bool, cols)
			for j := range next[i] {
				liveNeighbors := countLiveNeighbors(grid, i, j)
				currAlive := grid[i][j].Alive.Get()
				next[i][j] = liveRule(currAlive, liveNeighbors)
			}
		}
		for i := range grid {
			for j := range grid[i] {
				grid[i][j].Alive.Set(next[i][j])
			}
		}
	}

	run := ui.NewButton("▶️ Step").Padding(8)
	run.OnClick(step)

	count := core.Derive(func() string {
		total := 0
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j].Alive.Get() {
					total++
				}
			}
		}
		return "Alive: " + strconv.Itoa(total)
	})

	label := ui.NewLabel("").Padding(8).Border("1px solid grey")
	label.BindTo(count)

	root := ui.NewVBox(container, run, label).Padding(12)

	win := ui.NewWindow()
	win.Add(root)
	win.Run()
}

func countLiveNeighbors(grid [][]*Cell, x, y int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx, ny := x+dx, y+dy
			if dx == 0 && dy == 0 {
				continue
			}
			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
				if grid[nx][ny].Alive.Get() {
					count++
				}
			}
		}
	}
	return count
}

func liveRule(alive bool, neighbors int) bool {
	switch {
	case alive && (neighbors == 2 || neighbors == 3):
		return true
	case !alive && neighbors == 3:
		return true
	default:
		return false
	}
}

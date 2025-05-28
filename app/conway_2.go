
package main

/*
import (
	"gocore/core"
	"gocore/shared"
	"gocore/ui"
	"math/rand"
	"strconv"
)

type Cell struct {
	Alive *core.Observable[bool]
	Next  *core.Derived[bool]
	View  shared.Widget
}

func newCell() *Cell {
	alive := core.NewObservable(rand.Float32() < 0.2)

	view := ui.NewDiv().
		Width(20).Height(20).Border("1px solid #ccc")

	view.BindStyle(core.Derive(func() map[string]string {
		if alive.Get() {
			return map[string]string{"background": "black"}
		}
		return map[string]string{"background": "white"}
	}))

	return &Cell{Alive: alive, View: view}
}

func ConwaySafeReactive() {
	const rows, cols = 20, 20
	grid := make([][]*Cell, rows)

	for i := range grid {
		grid[i] = make([]*Cell, cols)
		for j := range grid[i] {
			grid[i][j] = newCell()
		}
	}

	// Derive Next state safely
	for i := range grid {
		for j := range grid[i] {
			cell := grid[i][j]
			neighbors := getNeighbors(grid, i, j)

			cell.Next = core.Derive(func() bool {
				count := 0
				for _, n := range neighbors {
					if n.Alive.Get() {
						count++
					}
				}
				return liveRule(cell.Alive.Get(), count)
			})
		}
	}

	// UI Setup
	container := ui.NewVBox()
	for _, row := range grid {
		h := ui.NewHBox()
		for _, cell := range row {
			h.Add(cell.View)
		}
		container.Add(h)
	}

	stepBtn := ui.NewButton("▶️ Step").Padding(8)

	stepBtn.OnClick(func() {
		// Precompute next states (forces recompute)
		for i := range grid {
			for j := range grid[i] {
				_ = grid[i][j].Next.Get()
			}
		}
		// Now apply changes
		for i := range grid {
			for j := range grid[i] {
				next := grid[i][j].Next.Get()
				if grid[i][j].Alive.Get() != next {
					grid[i][j].Alive.Set(next)
				}
			}
		}
	})

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

	win := ui.NewWindow()
	win.Add(container, stepBtn, label)
	win.Run()
}

func getNeighbors(grid [][]*Cell, x, y int) []*Cell {
	var result []*Cell
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx, ny := x+dx, y+dy
			if dx == 0 && dy == 0 {
				continue
			}
			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
				result = append(result, grid[nx][ny])
			}
		}
	}
	return result
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

*/

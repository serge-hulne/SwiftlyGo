package main

import (
	"gocore/ui"
	"math/rand"
	"strconv"
	"sync"
	"time"

	dom "honnef.co/go/js/dom/v2"
)

const (
	rows = 20
	cols = 20
)

type Cell struct {
	i, j    int
	alive   bool
	lock    sync.Mutex
	view    dom.HTMLElement
	grid    *[][]*Cell
	changed chan struct{}
}

func newCell(i, j int, grid *[][]*Cell) *Cell {
	doc := dom.GetWindow().Document()
	div := doc.CreateElement("div").(dom.HTMLElement)
	div.Style().SetProperty("width", "20px", "")
	div.Style().SetProperty("height", "20px", "")
	div.Style().SetProperty("border", "1px solid #ccc", "")
	div.Style().SetProperty("box-sizing", "border-box", "")
	div.Style().SetProperty("background", "white", "")

	c := &Cell{
		i:       i,
		j:       j,
		grid:    grid,
		view:    div,
		changed: make(chan struct{}, 1),
	}

	// Random start
	if rand.Float32() < 0.2 {
		c.alive = true
		c.updateVisual()
	}

	// Allow clicking to toggle manually
	div.AddEventListener("click", false, func(dom.Event) {
		c.lock.Lock()
		c.alive = !c.alive
		c.updateVisual()
		c.lock.Unlock()
		c.broadcast()
	})

	return c
}

func (c *Cell) updateVisual() {
	if c.alive {
		c.view.Style().SetProperty("background", "black", "")
	} else {
		c.view.Style().SetProperty("background", "white", "")
	}
}

func (c *Cell) broadcast() {
	select {
	case c.changed <- struct{}{}:
	default:
	}
}

func (c *Cell) run() {
	for {
		time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)

		c.lock.Lock()
		neighbors := c.getNeighbors()
		aliveCount := 0
		for _, n := range neighbors {
			n.lock.Lock()
			if n.alive {
				aliveCount++
			}
			n.lock.Unlock()
		}

		next := c.alive
		switch {
		case c.alive && (aliveCount < 2 || aliveCount > 3):
			next = false
		case !c.alive && aliveCount == 3:
			next = true
		}

		if next != c.alive {
			c.alive = next
			c.updateVisual()
			c.broadcast()
		}
		c.lock.Unlock()
	}
}

func (c *Cell) getNeighbors() []*Cell {
	var result []*Cell
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			ni, nj := c.i+dx, c.j+dy
			if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
				result = append(result, (*c.grid)[ni][nj])
			}
		}
	}
	return result
}

func ConwayReactiveGrid() {
	grid := make([][]*Cell, rows)
	for i := range grid {
		grid[i] = make([]*Cell, cols)
	}

	// Initialize grid
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = newCell(i, j, &grid)
		}
	}

	// Start goroutines
	for i := range grid {
		for j := range grid[i] {
			go grid[i][j].run()
		}
	}

	// Display grid
	container := ui.NewDiv()
	container.Element().Style().SetProperty("display", "grid", "")
	container.Element().Style().SetProperty("grid-template-columns", "repeat("+strconv.Itoa(cols)+", 20px)", "")
	container.Element().Style().SetProperty("gap", "0px", "")

	for i := range grid {
		for j := range grid[i] {
			container.Element().AppendChild(grid[i][j].view)
		}
	}

	win := ui.NewWindow()
	win.Add(container)
	win.Run()
}

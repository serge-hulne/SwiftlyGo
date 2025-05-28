package ui

import dom "honnef.co/go/js/dom/v2"

type Window struct {
	children []Widget
}

func NewWindow() *Window {
	return &Window{}
}

func (w *Window) Add(widgets ...Widget) {
	w.children = append(w.children, widgets...)
	for _, widget := range widgets {
		AppendToBody(widget.Element())
	}
}

func MessageSnackbar(msg string) {
	dom.GetWindow().Alert(msg)
}

func (w *Window) Run() {
	// Prevents exit, but lets the event loop run
	c := make(chan struct{})
	<-c
}

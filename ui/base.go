package ui

import (
	"fmt"
	"gocore/core"
	"gocore/shared"

	dom "honnef.co/go/js/dom/v2"
)

// -------- Widgets (Reusable) --------

// Label

type Label struct{ *shared.BaseWidget }

func NewLabel(text string) *Label {
	doc := dom.GetWindow().Document()
	el := doc.CreateElement("div").(dom.HTMLElement)
	el.SetTextContent(text)
	return &Label{&shared.BaseWidget{Inner: el, El: el}}
}

// func (l *Label) BindTo(obs *Observable[string]) {
// 	l.BindText(obs)
// }

func (l *Label) Padding(px int) *Label      { l.BaseWidget = l.BaseWidget.Padding(px); return l }
func (l *Label) Background(c string) *Label { l.BaseWidget = l.BaseWidget.Background(c); return l }
func (l *Label) Border(s string) *Label     { l.BaseWidget = l.BaseWidget.Border(s); return l }
func (l *Label) Center() *Label             { l.BaseWidget = l.BaseWidget.Center(); return l }

// Button

type Button struct{ *shared.BaseWidget }

func NewButton(label string) *Button {
	doc := dom.GetWindow().Document()
	btn := doc.CreateElement("button").(dom.HTMLElement)
	btn.SetTextContent(label)
	return &Button{&shared.BaseWidget{Inner: btn, El: btn}}
}

func (b *Button) OnClick(handler func()) {
	b.Inner.AddEventListener("click", false, func(dom.Event) {
		handler()
	})
}

func (b *Button) Padding(px int) *Button      { b.BaseWidget = b.BaseWidget.Padding(px); return b }
func (b *Button) Background(c string) *Button { b.BaseWidget = b.BaseWidget.Background(c); return b }
func (b *Button) Border(s string) *Button     { b.BaseWidget = b.BaseWidget.Border(s); return b }
func (b *Button) Center() *Button             { b.BaseWidget = b.BaseWidget.Center(); return b }

// TextField

type TextField struct {
	*shared.BaseWidget
	input *dom.HTMLInputElement
}

func NewTextField() *TextField {
	doc := dom.GetWindow().Document()
	in := doc.CreateElement("input").(*dom.HTMLInputElement)
	return &TextField{
		BaseWidget: &shared.BaseWidget{Inner: in, El: in},
		input:      in,
	}
}

func (t *TextField) Text() string {
	return t.input.Value()
}

func (t *TextField) BindTo(obs *core.Observable[string]) {
	obs.Subscribe(func(val string) {
		t.input.SetValue(val)
	})
	t.input.AddEventListener("input", false, func(dom.Event) {
		obs.Set(t.Text())
	})
}

func (d *TextField) Padding(px int) *TextField { d.BaseWidget = d.BaseWidget.Padding(px); return d }
func (d *TextField) Background(c string) *TextField {
	d.BaseWidget = d.BaseWidget.Background(c)
	return d
}
func (d *TextField) Border(s string) *TextField { d.BaseWidget = d.BaseWidget.Border(s); return d }
func (d *TextField) Center() *TextField         { d.BaseWidget = d.BaseWidget.Center(); return d }

func (t *TextField) Input() *dom.HTMLInputElement {
	return t.input
}

func (t *TextField) SetText(val string) {
	t.input.SetValue(val)
}

// Div

type Div struct {
	*shared.BaseWidget
}

func NewDiv() *Div {
	doc := dom.GetWindow().Document()
	d := doc.CreateElement("div").(dom.HTMLElement)
	return &Div{&shared.BaseWidget{Inner: d, El: d}}
}

func (d *Div) Add(children ...Widget) {
	for _, c := range children {
		d.El.AppendChild(c.Element())
	}
}

func (d *Div) Padding(px int) *Div      { d.BaseWidget = d.BaseWidget.Padding(px); return d }
func (d *Div) Background(c string) *Div { d.BaseWidget = d.BaseWidget.Background(c); return d }
func (d *Div) Border(s string) *Div     { d.BaseWidget = d.BaseWidget.Border(s); return d }
func (d *Div) Center() *Div             { d.BaseWidget = d.BaseWidget.Center(); return d }

// ------------

// --- VBox ---
type VBox struct {
	*shared.BaseWidget
}

func NewVBox(children ...Widget) *VBox {
	doc := dom.GetWindow().Document()
	container := doc.CreateElement("div").(dom.HTMLElement)

	container.Style().SetProperty("display", "flex", "")
	container.Style().SetProperty("flex-direction", "column", "")
	container.Style().SetProperty("gap", "0.5rem", "")

	vbox := &VBox{&shared.BaseWidget{Inner: container, El: container}}

	if len(children) > 0 {
		vbox.Add(children...)
	}
	return vbox
}

func (v *VBox) Add(children ...Widget) {
	for _, c := range children {
		v.El.AppendChild(c.Element())
	}
}

func (v *VBox) Padding(px int) *VBox      { v.BaseWidget = v.BaseWidget.Padding(px); return v }
func (v *VBox) Background(c string) *VBox { v.BaseWidget = v.BaseWidget.Background(c); return v }
func (v *VBox) Border(s string) *VBox     { v.BaseWidget = v.BaseWidget.Border(s); return v }
func (v *VBox) Center() *VBox             { v.BaseWidget = v.BaseWidget.Center(); return v }

// --- HBox ---
type HBox struct {
	*shared.BaseWidget
}

func NewHBox(children ...Widget) *HBox {
	doc := dom.GetWindow().Document()
	container := doc.CreateElement("div").(dom.HTMLElement)

	container.Style().SetProperty("display", "flex", "")
	container.Style().SetProperty("flex-direction", "row", "")
	container.Style().SetProperty("gap", "0.5rem", "")

	hbox := &HBox{&shared.BaseWidget{Inner: container, El: container}}

	if len(children) > 0 {
		hbox.Add(children...)
	}

	return hbox
}

func (h *HBox) Add(children ...Widget) {
	for _, c := range children {
		h.El.AppendChild(c.Element())
	}
}

func (h *HBox) Padding(px int) *HBox      { h.BaseWidget = h.BaseWidget.Padding(px); return h }
func (h *HBox) Background(c string) *HBox { h.BaseWidget = h.BaseWidget.Background(c); return h }
func (h *HBox) Border(s string) *HBox     { h.BaseWidget = h.BaseWidget.Border(s); return h }
func (h *HBox) Center() *HBox             { h.BaseWidget = h.BaseWidget.Center(); return h }

// --- TextArea ---
type TextArea struct {
	*shared.BaseWidget
	area *dom.HTMLTextAreaElement
}

func NewTextArea() *TextArea {
	doc := dom.GetWindow().Document()
	ta := doc.CreateElement("textarea").(*dom.HTMLTextAreaElement)

	return &TextArea{
		BaseWidget: &shared.BaseWidget{Inner: ta, El: ta},
		area:       ta,
	}
}

func (t *TextArea) Text() string {
	return t.area.Value()
}

func (t *TextArea) SetText(val string) {
	t.area.SetValue(val)
}

func (t *TextArea) BindTo(obs *core.Observable[string]) {
	obs.Subscribe(func(val string) {
		t.area.SetValue(val)
	})
	t.area.AddEventListener("input", false, func(dom.Event) {
		obs.Set(t.Text())
	})
}

func (t *TextArea) Padding(px int) *TextArea { t.BaseWidget = t.BaseWidget.Padding(px); return t }

func (t *TextArea) Background(c string) *TextArea {
	t.BaseWidget = t.BaseWidget.Background(c)
	return t
}
func (t *TextArea) Border(s string) *TextArea { t.BaseWidget = t.BaseWidget.Border(s); return t }

func (t *TextArea) Center() *TextArea { t.BaseWidget = t.BaseWidget.Center(); return t }

func (b *TextArea) wrapWithStyle(styles map[string]string) *TextArea {
	if b.IsWrapped {
		for prop, val := range styles {
			b.El.Style().SetProperty(prop, val, "")
		}
		return b
	}

	doc := dom.GetWindow().Document()
	wrapper := doc.CreateElement("div").(dom.HTMLElement)

	for prop, val := range styles {
		wrapper.Style().SetProperty(prop, val, "")
	}

	// ✅ STEP 1: Detach from DOM if attached
	if parent := b.El.ParentNode(); parent != nil {
		parent.RemoveChild(b.El)
	}

	// ✅ STEP 2: Append to wrapper safely
	wrapper.AppendChild(b.El)

	// ✅ STEP 3: Track new outermost
	b.El = wrapper
	b.IsWrapped = true

	return b
}

func (b *TextArea) Width(percent string) *TextArea {
	return b.wrapWithStyle(map[string]string{
		"width": fmt.Sprintf("%s", percent),
	})
}

// Updated BindText
func BindText(el dom.HTMLElement, obs core.ReadonlyObservable[string]) {
	obs.Subscribe(func(val string) {
		el.SetTextContent(val)
	})
}

// Updated Label.BindTo
func (l *Label) BindTo(obs core.ReadonlyObservable[string]) {
	l.BindText(obs)
}

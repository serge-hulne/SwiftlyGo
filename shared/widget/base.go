package widget

import (
	"fmt"

	"gocore/reactive"

	dom "honnef.co/go/js/dom/v2"
)

type BaseWidget struct {
	Inner     dom.HTMLElement
	El        dom.HTMLElement
	IsWrapped bool
}

func (b *BaseWidget) Element() dom.HTMLElement {
	return b.El
}

func (b *BaseWidget) SetText(text string) {
	b.Inner.SetTextContent(text)
}

func (b *BaseWidget) SetStyle(prop, value string) {
	b.Inner.Style().SetProperty(prop, value, "")
}

func (b *BaseWidget) wrapWithStyle(styles map[string]string) *BaseWidget {
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

	if parent := b.El.ParentNode(); parent != nil {
		parent.RemoveChild(b.El)
	}

	wrapper.AppendChild(b.El)
	b.El = wrapper
	b.IsWrapped = true
	return b
}

func (b *BaseWidget) Padding(px int) *BaseWidget {
	return b.wrapWithStyle(map[string]string{
		"padding": fmt.Sprintf("%dpx", px),
	})
}

func (b *BaseWidget) Background(color string) *BaseWidget {
	return b.wrapWithStyle(map[string]string{
		"background": color,
	})
}

func (b *BaseWidget) Border(style string) *BaseWidget {
	return b.wrapWithStyle(map[string]string{
		"border": style,
	})
}

func (b *BaseWidget) BindText(obs reactive.ReadonlyObservable[string]) {
	obs.Subscribe(func(val string) {
		b.SetText(val)
	})
}

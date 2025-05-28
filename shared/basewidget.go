package shared

import (
	"fmt"
	"strings"

	dom "honnef.co/go/js/dom/v2"
)

type ReadonlyObservable[T any] interface {
	Get() T
	Subscribe(func(T))
}

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

func (b *BaseWidget) SetHTML(html string) {
	b.Inner.SetInnerHTML(html)
}

func (b *BaseWidget) SetClass(class string) {
	b.Inner.SetAttribute("class", class)
}

func (b *BaseWidget) AddClass(class string) {
	curr := b.Inner.GetAttribute("class")
	if curr != "" {
		class = curr + " " + class
	}
	b.Inner.SetAttribute("class", class)
}

func (b *BaseWidget) RemoveClass(class string) {
	curr := strings.Fields(b.Inner.GetAttribute("class"))
	var result []string
	for _, c := range curr {
		if c != class {
			result = append(result, c)
		}
	}
	b.Inner.SetAttribute("class", strings.Join(result, " "))
}

func (b *BaseWidget) SetStyle(prop, value string) {
	b.Inner.Style().SetProperty(prop, value, "")
}

func (b *BaseWidget) Show() {
	b.Inner.Style().SetProperty("display", "", "")
}

func (b *BaseWidget) Hide() {
	b.Inner.Style().SetProperty("display", "none", "")
}

func (b *BaseWidget) SetAttr(key, val string) {
	b.Inner.SetAttribute(key, val)
}

func (b *BaseWidget) GetAttr(key string) string {
	return b.Inner.GetAttribute(key)
}

//func (b *BaseWidget) OnClick(handler func()) {
//	b.Inner.AddEventListener("click", false, func(dom.Event) {
//		handler()
//	})
//}

func (b *BaseWidget) OnClick(handler func()) {
	b.El.AddEventListener("click", false, func(dom.Event) {
		handler()
	})
}

func (b *BaseWidget) On(event string, handler func()) {
	b.Inner.AddEventListener(event, false, func(dom.Event) {
		handler()
	})
}

func (b *BaseWidget) AppendTo(parent dom.Element) {
	parent.AppendChild(b.El)
}

func (b *BaseWidget) Remove() {
	if parent := b.El.ParentNode(); parent != nil {
		parent.RemoveChild(b.El)
	}
}

func (b *BaseWidget) BindText(obs ReadonlyObservable[string]) {
	obs.Subscribe(func(val string) {
		b.SetText(val)
	})
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

func (b *BaseWidget) Center() *BaseWidget {
	b.El.Style().SetProperty("justify-content", "center", "")
	b.El.Style().SetProperty("align-items", "center", "")
	if b.El.Style().GetPropertyValue("display") == "" {
		b.El.Style().SetProperty("display", "flex", "")
	}
	return b
}

func (b *BaseWidget) Padding(px int) *BaseWidget {
	return b.wrapWithStyle(map[string]string{
		"padding": fmt.Sprintf("%dpx", px),
	})
}

func (b *BaseWidget) Margin(px int) *BaseWidget {
	return b.wrapWithStyle(map[string]string{
		"margin": fmt.Sprintf("%dpx", px),
	})
}

func (b *BaseWidget) Width(px int) *BaseWidget {
	return b.wrapWithStyle(map[string]string{
		"width": fmt.Sprintf("%dpx", px),
	})
}

func (b *BaseWidget) Height(px int) *BaseWidget {
	return b.wrapWithStyle(map[string]string{
		"height": fmt.Sprintf("%dpx", px),
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

type Widget interface {
	Element() dom.HTMLElement
}

//func (b *BaseWidget) BindStyle(obs ReadonlyObservable[map[string]string]) {
//	obs.Subscribe(func(styles map[string]string) {
//		for prop, val := range styles {
//			b.Inner.Style().SetProperty(prop, val, "")
//		}
//	})
//}

func (b *BaseWidget) BindStyle(obs ReadonlyObservable[map[string]string]) {
	obs.Subscribe(func(styles map[string]string) {
		for prop, val := range styles {
			b.El.Style().SetProperty(prop, val, "")
		}
	})
}

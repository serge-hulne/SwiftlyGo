package ui

import (
	"gocore/shared"

	dom "honnef.co/go/js/dom/v2"
)

type CheckBox struct {
	*shared.BaseWidget
	input *dom.HTMLInputElement
}

func NewCheckBox() *CheckBox {
	doc := dom.GetWindow().Document()
	in := doc.CreateElement("input").(*dom.HTMLInputElement)
	in.SetType("checkbox")

	return &CheckBox{
		BaseWidget: &shared.BaseWidget{Inner: in, El: in},
		input:      in,
	}
}

func (c *CheckBox) Checked() bool {
	return c.input.Checked()
}

func (c *CheckBox) SetChecked(v bool) {
	c.input.SetChecked(v)
}

func (c *CheckBox) OnChange(handler func(bool)) {
	c.input.AddEventListener("change", false, func(dom.Event) {
		handler(c.input.Checked())
	})
}

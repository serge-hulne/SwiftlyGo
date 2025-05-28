package ui

import dom "honnef.co/go/js/dom/v2"

type Widget interface {
	Element() dom.HTMLElement
}

func AppendToBody(el dom.Element) {
	doc := dom.GetWindow().Document()
	body := doc.GetElementsByTagName("body")[0].(*dom.HTMLBodyElement)
	body.AppendChild(el)
}

package gforms

import (
	"bytes"
)

type textAreaWidget struct {
	Type  string
	Attrs map[string]string
	Widget
}

func (wg *textAreaWidget) html(f FieldInterface) string {
	var buffer bytes.Buffer
	err := Template.ExecuteTemplate(&buffer, "TextAreaWidget", widgetContext{
		Field: f,
		Attrs: wg.Attrs,
		Value: f.GetV().RawStr,
	})
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Generate text area field: <textarea ...>
func TextAreaWidget(attrs map[string]string) Widget {
	w := new(textAreaWidget)
	w.Type = "textarea"
	if attrs == nil {
		attrs = map[string]string{}
	}
	w.Attrs = attrs
	return w
}

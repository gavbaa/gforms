package gforms

import (
	"reflect"
)

// It maps value to FormInstance.CleanedData as type `string`.
type FileField struct {
	BaseField
}

// Create a new TextField instance.
func (f *FileField) New() FieldInterface {
	fi := new(FileFieldInstance)
	fi.Model = f
	fi.V = nilV("")
	return fi
}

// Instance for TextField
type FileFieldInstance struct {
	FieldInstance
}

// Create a new TextField with validators and widgets.
func NewFileField(name string, vs Validators, ws ...Widget) *FileField {
	f := new(FileField)
	f.Name = name
	f.Validators = vs
	if len(ws) > 0 {
		f.Widget = ws[0]
	}
	return f
}

// Get a value from request data, and clean it as type `string`.
func (f *FileFieldInstance) Clean(data Data) error {
	m, hasField := data[f.Model.GetName()]
	if hasField {
		f.V = m
		v := m.rawValueAsString()
		m.Kind = reflect.String
		if v != nil {
			m.Value = *v
			m.IsNil = false
		}
	}
	return nil
}

func (f *FileFieldInstance) html() string {
	return renderTemplate("FileTypeField", newTemplateContext(f))
}

// Get as HTML format.
func (f *FileFieldInstance) Html() string {
	return fieldToHtml(f)
}

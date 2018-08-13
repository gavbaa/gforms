package gforms

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
)

// It maps value to FormInstance.CleanedData as type `bool`.
type BooleanField struct {
	BaseField
}

// Create a new BooleanField instance.
func (f *BooleanField) New() FieldInterface {
	fi := new(BooleanFieldInstance)
	fi.Model = f
	fi.V = nilV("")
	return fi
}

// Instance for BooleanField.
type BooleanFieldInstance struct {
	FieldInstance
}

type booleanContext struct {
	Field   FieldInterface
	Checked bool
}

// Create a new BooleanField with validators and widgets.
func NewBooleanField(name string, vs Validators, ws ...Widget) *BooleanField {
	f := new(BooleanField)
	f.Name = name
	f.Validators = vs
	if len(ws) > 0 {
		f.Widget = ws[0]
	}
	return f
}

// Get a value from request data, and clean it as type `bool`.
func (f *BooleanFieldInstance) Clean(data Data) error {
	m, hasField := data[f.GetName()]
	if hasField {
		f.V = m
		v := false
		if m.Kind == reflect.String {
			vs := m.rawValueAsString()
			if vs != nil {
				if *vs == "" {
					// It's very confusing to allow the empty value to be true, but we need it for checkboxes.
					v = true
				} else {
					// Per the HTML standard, "on" is the default value for a checkbox with no value.
					if strings.ToLower(*vs) == "on" {
						v = true
					} else {
						b, err := strconv.ParseBool(*vs)
						if err != nil {
							m.IsNil = true
						} else {
							v = b
						}
					}
				}
			}
		} else if m.Kind == reflect.Bool {
			v = m.rawValueAsBool()
		}
		m.Value = v
		m.Kind = reflect.Bool
		m.IsNil = false
		return nil
	}
	nv := newV("", false, reflect.Bool)
	nv.Value = false
	nv.IsNil = false
	f.V = nv
	return nil
}

func (f *BooleanFieldInstance) SetInitial(v string) {
	f.V.RawStr = v
	f.V.RawValue = []string{v}
	f.V.IsNil = false
	b, err := strconv.ParseBool(v)
	if err != nil {
		f.V.IsNil = true
	} else {
		f.V.Value = b
	}
}

func (f *BooleanFieldInstance) html() string {
	var buffer bytes.Buffer
	cx := new(booleanContext)
	cx.Field = f
	checked, _ := f.V.Value.(bool)
	cx.Checked = checked
	err := Template.ExecuteTemplate(&buffer, "BooleanTypeField", cx)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Get as HTML format.
func (f *BooleanFieldInstance) Html() string {
	return fieldToHtml(f)
}


// Code generated by go-genaccessor; DO NOT EDIT.

package example

import (
	"encoding"
)

func (m Person) ID() string {
	return m.id
}

func (m Person) Name() string {
	return m.name
}

func (m *Person) Rename(s string) {
	m.name = s
}

func (m Person) Tags() []string {
	return m.tags
}

func (m *Person) SetTags(s []string) {
	m.tags = s
}

func (m Person) Text() encoding.TextMarshaler {
	return m.text
}

func (m *Person) SetText(s encoding.TextMarshaler) {
	m.text = s
}

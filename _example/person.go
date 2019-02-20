package example

import (
	"encoding"
)

//go:generate go-genaccessor

type Person struct {
	id   string                 `getter:""`
	name string                 `getter:"" setter:"Rename"`
	tags []string               `getter:"" setter:""`
	text encoding.TextMarshaler `getter:"" setter:""`
}

package dagpb

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	"github.com/ipld/go-ipld-prime/schema"
)

const (
	midvalue  = schema.Maybe(4)
	allowNull = schema.Maybe(5)
)

type maState uint8

const (
	maState_initial maState = iota
	maState_midKey
	maState_expectValue
	maState_midValue
	maState_finished
)

type laState uint8

const (
	laState_initial laState = iota
	laState_midValue
	laState_finished
)

package dagpb

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _RawNode struct{ x []byte }
type RawNode = *_RawNode

func (n RawNode) Bytes() []byte {
	return n.x
}
func (_RawNode__Prototype) FromBytes(v []byte) (RawNode, error) {
	n := _RawNode{v}
	return &n, nil
}

type _RawNode__Maybe struct {
	m schema.Maybe
	v RawNode
}
type MaybeRawNode = *_RawNode__Maybe

func (m MaybeRawNode) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeRawNode) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeRawNode) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeRawNode) AsNode() ipld.Node {
	switch m.m {
	case schema.Maybe_Absent:
		return ipld.Absent
	case schema.Maybe_Null:
		return ipld.Null
	case schema.Maybe_Value:
		return m.v
	default:
		panic("unreachable")
	}
}
func (m MaybeRawNode) Must() RawNode {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}

var _ ipld.Node = (RawNode)(&_RawNode{})
var _ schema.TypedNode = (RawNode)(&_RawNode{})

func (RawNode) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Bytes
}
func (RawNode) LookupByString(string) (ipld.Node, error) {
	return mixins.Bytes{"dagpb.RawNode"}.LookupByString("")
}
func (RawNode) LookupByNode(ipld.Node) (ipld.Node, error) {
	return mixins.Bytes{"dagpb.RawNode"}.LookupByNode(nil)
}
func (RawNode) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Bytes{"dagpb.RawNode"}.LookupByIndex(0)
}
func (RawNode) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return mixins.Bytes{"dagpb.RawNode"}.LookupBySegment(seg)
}
func (RawNode) MapIterator() ipld.MapIterator {
	return nil
}
func (RawNode) ListIterator() ipld.ListIterator {
	return nil
}
func (RawNode) Length() int {
	return -1
}
func (RawNode) IsAbsent() bool {
	return false
}
func (RawNode) IsNull() bool {
	return false
}
func (RawNode) AsBool() (bool, error) {
	return mixins.Bytes{"dagpb.RawNode"}.AsBool()
}
func (RawNode) AsInt() (int, error) {
	return mixins.Bytes{"dagpb.RawNode"}.AsInt()
}
func (RawNode) AsFloat() (float64, error) {
	return mixins.Bytes{"dagpb.RawNode"}.AsFloat()
}
func (RawNode) AsString() (string, error) {
	return mixins.Bytes{"dagpb.RawNode"}.AsString()
}
func (n RawNode) AsBytes() ([]byte, error) {
	return n.x, nil
}
func (RawNode) AsLink() (ipld.Link, error) {
	return mixins.Bytes{"dagpb.RawNode"}.AsLink()
}
func (RawNode) Prototype() ipld.NodePrototype {
	return _RawNode__Prototype{}
}

type _RawNode__Prototype struct{}

func (_RawNode__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _RawNode__Builder
	nb.Reset()
	return &nb
}

type _RawNode__Builder struct {
	_RawNode__Assembler
}

func (nb *_RawNode__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_RawNode__Builder) Reset() {
	var w _RawNode
	var m schema.Maybe
	*nb = _RawNode__Builder{_RawNode__Assembler{w: &w, m: &m}}
}

type _RawNode__Assembler struct {
	w *_RawNode
	m *schema.Maybe
}

func (na *_RawNode__Assembler) reset() {}
func (_RawNode__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.BytesAssembler{"dagpb.RawNode"}.BeginMap(0)
}
func (_RawNode__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.BytesAssembler{"dagpb.RawNode"}.BeginList(0)
}
func (na *_RawNode__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.BytesAssembler{"dagpb.RawNode"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	}
	panic("unreachable")
}
func (_RawNode__Assembler) AssignBool(bool) error {
	return mixins.BytesAssembler{"dagpb.RawNode"}.AssignBool(false)
}
func (_RawNode__Assembler) AssignInt(int) error {
	return mixins.BytesAssembler{"dagpb.RawNode"}.AssignInt(0)
}
func (_RawNode__Assembler) AssignFloat(float64) error {
	return mixins.BytesAssembler{"dagpb.RawNode"}.AssignFloat(0)
}
func (_RawNode__Assembler) AssignString(string) error {
	return mixins.BytesAssembler{"dagpb.RawNode"}.AssignString("")
}
func (na *_RawNode__Assembler) AssignBytes(v []byte) error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	}
	if na.w == nil {
		na.w = &_RawNode{}
	}
	na.w.x = v
	*na.m = schema.Maybe_Value
	return nil
}
func (_RawNode__Assembler) AssignLink(ipld.Link) error {
	return mixins.BytesAssembler{"dagpb.RawNode"}.AssignLink(nil)
}
func (na *_RawNode__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_RawNode); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v2, err := v.AsBytes(); err != nil {
		return err
	} else {
		return na.AssignBytes(v2)
	}
}
func (_RawNode__Assembler) Prototype() ipld.NodePrototype {
	return _RawNode__Prototype{}
}
func (RawNode) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n RawNode) Representation() ipld.Node {
	return (*_RawNode__Repr)(n)
}

type _RawNode__Repr = _RawNode

var _ ipld.Node = &_RawNode__Repr{}

type _RawNode__ReprPrototype = _RawNode__Prototype
type _RawNode__ReprAssembler = _RawNode__Assembler

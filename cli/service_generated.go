// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package cli

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
)

type ODictMethod int8

const (
	ODictMethodLookup  ODictMethod = 1
	ODictMethodSplit   ODictMethod = 2
	ODictMethodIndex   ODictMethod = 3
	ODictMethodSearch  ODictMethod = 4
	ODictMethodCompile ODictMethod = 5
	ODictMethodWrite   ODictMethod = 6
	ODictMethodLexicon ODictMethod = 7
)

var EnumNamesODictMethod = map[ODictMethod]string{
	ODictMethodLookup:  "Lookup",
	ODictMethodSplit:   "Split",
	ODictMethodIndex:   "Index",
	ODictMethodSearch:  "Search",
	ODictMethodCompile: "Compile",
	ODictMethodWrite:   "Write",
	ODictMethodLexicon: "Lexicon",
}

var EnumValuesODictMethod = map[string]ODictMethod{
	"Lookup":  ODictMethodLookup,
	"Split":   ODictMethodSplit,
	"Index":   ODictMethodIndex,
	"Search":  ODictMethodSearch,
	"Compile": ODictMethodCompile,
	"Write":   ODictMethodWrite,
	"Lexicon": ODictMethodLexicon,
}

func (v ODictMethod) String() string {
	if s, ok := EnumNamesODictMethod[v]; ok {
		return s
	}
	return "ODictMethod(" + strconv.FormatInt(int64(v), 10) + ")"
}

type LookupPayload struct {
	_tab flatbuffers.Table
}

func GetRootAsLookupPayload(buf []byte, offset flatbuffers.UOffsetT) *LookupPayload {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LookupPayload{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsLookupPayload(buf []byte, offset flatbuffers.UOffsetT) *LookupPayload {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &LookupPayload{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *LookupPayload) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LookupPayload) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LookupPayload) Follow() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *LookupPayload) MutateFollow(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *LookupPayload) Split() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LookupPayload) MutateSplit(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *LookupPayload) Queries(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *LookupPayload) QueriesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func LookupPayloadStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func LookupPayloadAddFollow(builder *flatbuffers.Builder, follow bool) {
	builder.PrependBoolSlot(0, follow, false)
}
func LookupPayloadAddSplit(builder *flatbuffers.Builder, split int32) {
	builder.PrependInt32Slot(1, split, 0)
}
func LookupPayloadAddQueries(builder *flatbuffers.Builder, queries flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(queries), 0)
}
func LookupPayloadStartQueriesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func LookupPayloadEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type WritePayload struct {
	_tab flatbuffers.Table
}

func GetRootAsWritePayload(buf []byte, offset flatbuffers.UOffsetT) *WritePayload {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &WritePayload{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsWritePayload(buf []byte, offset flatbuffers.UOffsetT) *WritePayload {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &WritePayload{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *WritePayload) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *WritePayload) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *WritePayload) Xml() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *WritePayload) Out() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func WritePayloadStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func WritePayloadAddXml(builder *flatbuffers.Builder, xml flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(xml), 0)
}
func WritePayloadAddOut(builder *flatbuffers.Builder, out flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(out), 0)
}
func WritePayloadEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type SplitPayload struct {
	_tab flatbuffers.Table
}

func GetRootAsSplitPayload(buf []byte, offset flatbuffers.UOffsetT) *SplitPayload {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SplitPayload{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSplitPayload(buf []byte, offset flatbuffers.UOffsetT) *SplitPayload {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SplitPayload{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SplitPayload) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SplitPayload) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SplitPayload) Threshold() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SplitPayload) MutateThreshold(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *SplitPayload) Query() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func SplitPayloadStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SplitPayloadAddThreshold(builder *flatbuffers.Builder, threshold int32) {
	builder.PrependInt32Slot(0, threshold, 0)
}
func SplitPayloadAddQuery(builder *flatbuffers.Builder, query flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(query), 0)
}
func SplitPayloadEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type SearchPayload struct {
	_tab flatbuffers.Table
}

func GetRootAsSearchPayload(buf []byte, offset flatbuffers.UOffsetT) *SearchPayload {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SearchPayload{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSearchPayload(buf []byte, offset flatbuffers.UOffsetT) *SearchPayload {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SearchPayload{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SearchPayload) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SearchPayload) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SearchPayload) Force() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *SearchPayload) MutateForce(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *SearchPayload) Exact() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *SearchPayload) MutateExact(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *SearchPayload) Query() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func SearchPayloadStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func SearchPayloadAddForce(builder *flatbuffers.Builder, force bool) {
	builder.PrependBoolSlot(0, force, false)
}
func SearchPayloadAddExact(builder *flatbuffers.Builder, exact bool) {
	builder.PrependBoolSlot(1, exact, false)
}
func SearchPayloadAddQuery(builder *flatbuffers.Builder, query flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(query), 0)
}
func SearchPayloadEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type CompilePayload struct {
	_tab flatbuffers.Table
}

func GetRootAsCompilePayload(buf []byte, offset flatbuffers.UOffsetT) *CompilePayload {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CompilePayload{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCompilePayload(buf []byte, offset flatbuffers.UOffsetT) *CompilePayload {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CompilePayload{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CompilePayload) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CompilePayload) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CompilePayload) Path() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CompilePayload) Out() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func CompilePayloadStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func CompilePayloadAddPath(builder *flatbuffers.Builder, path flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(path), 0)
}
func CompilePayloadAddOut(builder *flatbuffers.Builder, out flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(out), 0)
}
func CompilePayloadEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

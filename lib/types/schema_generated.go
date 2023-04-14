// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package types

import (
	"bytes"
	"strconv"
	flatbuffers "github.com/google/flatbuffers/go"
)

type POS int8

const (
	POSun        POS = 0
	POSadj       POS = 1
	POSadv       POS = 2
	POSart       POS = 3
	POSconj      POS = 4
	POSintj      POS = 5
	POSn         POS = 6
	POSpart      POS = 7
	POSpref      POS = 8
	POSprep      POS = 9
	POSpostp     POS = 10
	POSpron      POS = 11
	POSsuff      POS = 12
	POSv         POS = 13
	POSabv       POS = 14
	POSadf       POS = 15
	POSaff       POS = 16
	POSaux_adj   POS = 17
	POSaux_v     POS = 18
	POSaux       POS = 19
	POSchr       POS = 20
	POSconj_c    POS = 21
	POSconj_s    POS = 22
	POScop       POS = 23
	POScf        POS = 24
	POSctr       POS = 25
	POSdet       POS = 26
	POSexpr      POS = 27
	POSinf       POS = 28
	POSintf      POS = 29
	POSname      POS = 30
	POSnum       POS = 31
	POSphr_adv   POS = 32
	POSphr_adj   POS = 33
	POSphr_prep  POS = 34
	POSphr       POS = 35
	POSpropn     POS = 36
	POSprov      POS = 37
	POSpunc      POS = 38
	POSsym       POS = 39
	POSvi        POS = 40
	POSvt        POS = 41
	POSadj_f     POS = 42
	POSadj_ix    POS = 43
	POSadj_kari  POS = 44
	POSadj_ku    POS = 45
	POSadj_na    POS = 46
	POSadj_nari  POS = 47
	POSadj_no    POS = 48
	POSadj_pn    POS = 49
	POSadj_shiku POS = 50
	POSadj_t     POS = 51
	POSadv_to    POS = 52
	POSn_adv     POS = 53
	POSn_pref    POS = 54
	POSn_suf     POS = 55
	POSn_t       POS = 56
	POSv_unspec  POS = 57
	POSv1_s      POS = 58
	POSv1        POS = 59
	POSv2a_s     POS = 60
	POSv2b_k     POS = 61
	POSv2b_s     POS = 62
	POSv2d_k     POS = 63
	POSv2d_s     POS = 64
	POSv2g_k     POS = 65
	POSv2g_s     POS = 66
	POSv2h_k     POS = 67
	POSv2h_s     POS = 68
	POSv2k_k     POS = 69
	POSv2k_s     POS = 70
	POSv2m_k     POS = 71
	POSv2m_s     POS = 72
	POSv2n_s     POS = 73
	POSv2r_k     POS = 74
	POSv2r_s     POS = 75
	POSv2s_s     POS = 76
	POSv2t_k     POS = 77
	POSv2t_s     POS = 78
	POSv2w_s     POS = 79
	POSv2y_k     POS = 80
	POSv2y_s     POS = 81
	POSv2z_s     POS = 82
	POSv4b       POS = 83
	POSv4g       POS = 84
	POSv4h       POS = 85
	POSv4k       POS = 86
	POSv4m       POS = 87
	POSv4n       POS = 88
	POSv4r       POS = 89
	POSv4s       POS = 90
	POSv4t       POS = 91
	POSv5aru     POS = 92
	POSv5b       POS = 93
	POSv5g       POS = 94
	POSv5k_s     POS = 95
	POSv5k       POS = 96
	POSv5m       POS = 97
	POSv5n       POS = 98
	POSv5r_i     POS = 99
	POSv5r       POS = 100
	POSv5s       POS = 101
	POSv5t       POS = 102
	POSv5u_s     POS = 103
	POSv5u       POS = 104
	POSv5uru     POS = 105
	POSvk        POS = 106
	POSvn        POS = 107
	POSvr        POS = 108
	POSvs_c      POS = 109
	POSvs_i      POS = 110
	POSvs_s      POS = 111
	POSvs        POS = 112
	POSvz        POS = 113
)

var EnumNamesPOS = map[POS]string{
	POSun:        "un",
	POSadj:       "adj",
	POSadv:       "adv",
	POSart:       "art",
	POSconj:      "conj",
	POSintj:      "intj",
	POSn:         "n",
	POSpart:      "part",
	POSpref:      "pref",
	POSprep:      "prep",
	POSpostp:     "postp",
	POSpron:      "pron",
	POSsuff:      "suff",
	POSv:         "v",
	POSabv:       "abv",
	POSadf:       "adf",
	POSaff:       "aff",
	POSaux_adj:   "aux_adj",
	POSaux_v:     "aux_v",
	POSaux:       "aux",
	POSchr:       "chr",
	POSconj_c:    "conj_c",
	POSconj_s:    "conj_s",
	POScop:       "cop",
	POScf:        "cf",
	POSctr:       "ctr",
	POSdet:       "det",
	POSexpr:      "expr",
	POSinf:       "inf",
	POSintf:      "intf",
	POSname:      "name",
	POSnum:       "num",
	POSphr_adv:   "phr_adv",
	POSphr_adj:   "phr_adj",
	POSphr_prep:  "phr_prep",
	POSphr:       "phr",
	POSpropn:     "propn",
	POSprov:      "prov",
	POSpunc:      "punc",
	POSsym:       "sym",
	POSvi:        "vi",
	POSvt:        "vt",
	POSadj_f:     "adj_f",
	POSadj_ix:    "adj_ix",
	POSadj_kari:  "adj_kari",
	POSadj_ku:    "adj_ku",
	POSadj_na:    "adj_na",
	POSadj_nari:  "adj_nari",
	POSadj_no:    "adj_no",
	POSadj_pn:    "adj_pn",
	POSadj_shiku: "adj_shiku",
	POSadj_t:     "adj_t",
	POSadv_to:    "adv_to",
	POSn_adv:     "n_adv",
	POSn_pref:    "n_pref",
	POSn_suf:     "n_suf",
	POSn_t:       "n_t",
	POSv_unspec:  "v_unspec",
	POSv1_s:      "v1_s",
	POSv1:        "v1",
	POSv2a_s:     "v2a_s",
	POSv2b_k:     "v2b_k",
	POSv2b_s:     "v2b_s",
	POSv2d_k:     "v2d_k",
	POSv2d_s:     "v2d_s",
	POSv2g_k:     "v2g_k",
	POSv2g_s:     "v2g_s",
	POSv2h_k:     "v2h_k",
	POSv2h_s:     "v2h_s",
	POSv2k_k:     "v2k_k",
	POSv2k_s:     "v2k_s",
	POSv2m_k:     "v2m_k",
	POSv2m_s:     "v2m_s",
	POSv2n_s:     "v2n_s",
	POSv2r_k:     "v2r_k",
	POSv2r_s:     "v2r_s",
	POSv2s_s:     "v2s_s",
	POSv2t_k:     "v2t_k",
	POSv2t_s:     "v2t_s",
	POSv2w_s:     "v2w_s",
	POSv2y_k:     "v2y_k",
	POSv2y_s:     "v2y_s",
	POSv2z_s:     "v2z_s",
	POSv4b:       "v4b",
	POSv4g:       "v4g",
	POSv4h:       "v4h",
	POSv4k:       "v4k",
	POSv4m:       "v4m",
	POSv4n:       "v4n",
	POSv4r:       "v4r",
	POSv4s:       "v4s",
	POSv4t:       "v4t",
	POSv5aru:     "v5aru",
	POSv5b:       "v5b",
	POSv5g:       "v5g",
	POSv5k_s:     "v5k_s",
	POSv5k:       "v5k",
	POSv5m:       "v5m",
	POSv5n:       "v5n",
	POSv5r_i:     "v5r_i",
	POSv5r:       "v5r",
	POSv5s:       "v5s",
	POSv5t:       "v5t",
	POSv5u_s:     "v5u_s",
	POSv5u:       "v5u",
	POSv5uru:     "v5uru",
	POSvk:        "vk",
	POSvn:        "vn",
	POSvr:        "vr",
	POSvs_c:      "vs_c",
	POSvs_i:      "vs_i",
	POSvs_s:      "vs_s",
	POSvs:        "vs",
	POSvz:        "vz",
}

var EnumValuesPOS = map[string]POS{
	"un":        POSun,
	"adj":       POSadj,
	"adv":       POSadv,
	"art":       POSart,
	"conj":      POSconj,
	"intj":      POSintj,
	"n":         POSn,
	"part":      POSpart,
	"pref":      POSpref,
	"prep":      POSprep,
	"postp":     POSpostp,
	"pron":      POSpron,
	"suff":      POSsuff,
	"v":         POSv,
	"abv":       POSabv,
	"adf":       POSadf,
	"aff":       POSaff,
	"aux_adj":   POSaux_adj,
	"aux_v":     POSaux_v,
	"aux":       POSaux,
	"chr":       POSchr,
	"conj_c":    POSconj_c,
	"conj_s":    POSconj_s,
	"cop":       POScop,
	"cf":        POScf,
	"ctr":       POSctr,
	"det":       POSdet,
	"expr":      POSexpr,
	"inf":       POSinf,
	"intf":      POSintf,
	"name":      POSname,
	"num":       POSnum,
	"phr_adv":   POSphr_adv,
	"phr_adj":   POSphr_adj,
	"phr_prep":  POSphr_prep,
	"phr":       POSphr,
	"propn":     POSpropn,
	"prov":      POSprov,
	"punc":      POSpunc,
	"sym":       POSsym,
	"vi":        POSvi,
	"vt":        POSvt,
	"adj_f":     POSadj_f,
	"adj_ix":    POSadj_ix,
	"adj_kari":  POSadj_kari,
	"adj_ku":    POSadj_ku,
	"adj_na":    POSadj_na,
	"adj_nari":  POSadj_nari,
	"adj_no":    POSadj_no,
	"adj_pn":    POSadj_pn,
	"adj_shiku": POSadj_shiku,
	"adj_t":     POSadj_t,
	"adv_to":    POSadv_to,
	"n_adv":     POSn_adv,
	"n_pref":    POSn_pref,
	"n_suf":     POSn_suf,
	"n_t":       POSn_t,
	"v_unspec":  POSv_unspec,
	"v1_s":      POSv1_s,
	"v1":        POSv1,
	"v2a_s":     POSv2a_s,
	"v2b_k":     POSv2b_k,
	"v2b_s":     POSv2b_s,
	"v2d_k":     POSv2d_k,
	"v2d_s":     POSv2d_s,
	"v2g_k":     POSv2g_k,
	"v2g_s":     POSv2g_s,
	"v2h_k":     POSv2h_k,
	"v2h_s":     POSv2h_s,
	"v2k_k":     POSv2k_k,
	"v2k_s":     POSv2k_s,
	"v2m_k":     POSv2m_k,
	"v2m_s":     POSv2m_s,
	"v2n_s":     POSv2n_s,
	"v2r_k":     POSv2r_k,
	"v2r_s":     POSv2r_s,
	"v2s_s":     POSv2s_s,
	"v2t_k":     POSv2t_k,
	"v2t_s":     POSv2t_s,
	"v2w_s":     POSv2w_s,
	"v2y_k":     POSv2y_k,
	"v2y_s":     POSv2y_s,
	"v2z_s":     POSv2z_s,
	"v4b":       POSv4b,
	"v4g":       POSv4g,
	"v4h":       POSv4h,
	"v4k":       POSv4k,
	"v4m":       POSv4m,
	"v4n":       POSv4n,
	"v4r":       POSv4r,
	"v4s":       POSv4s,
	"v4t":       POSv4t,
	"v5aru":     POSv5aru,
	"v5b":       POSv5b,
	"v5g":       POSv5g,
	"v5k_s":     POSv5k_s,
	"v5k":       POSv5k,
	"v5m":       POSv5m,
	"v5n":       POSv5n,
	"v5r_i":     POSv5r_i,
	"v5r":       POSv5r,
	"v5s":       POSv5s,
	"v5t":       POSv5t,
	"v5u_s":     POSv5u_s,
	"v5u":       POSv5u,
	"v5uru":     POSv5uru,
	"vk":        POSvk,
	"vn":        POSvn,
	"vr":        POSvr,
	"vs_c":      POSvs_c,
	"vs_i":      POSvs_i,
	"vs_s":      POSvs_s,
	"vs":        POSvs,
	"vz":        POSvz,
}

func (v POS) String() string {
	if s, ok := EnumNamesPOS[v]; ok {
		return s
	}
	return "POS(" + strconv.FormatInt(int64(v), 10) + ")"
}

type Definition struct {
	_tab flatbuffers.Table
}

func GetRootAsDefinition(buf []byte, offset flatbuffers.UOffsetT) *Definition {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Definition{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDefinition(buf []byte, offset flatbuffers.UOffsetT) *Definition {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Definition{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Definition) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Definition) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Definition) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Definition) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Definition) Examples(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Definition) ExamplesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DefinitionStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func DefinitionAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func DefinitionAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func DefinitionAddExamples(builder *flatbuffers.Builder, examples flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(examples), 0)
}
func DefinitionStartExamplesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DefinitionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Etymology struct {
	_tab flatbuffers.Table
}

func GetRootAsEtymology(buf []byte, offset flatbuffers.UOffsetT) *Etymology {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Etymology{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEtymology(buf []byte, offset flatbuffers.UOffsetT) *Etymology {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Etymology{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Etymology) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Etymology) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Etymology) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Etymology) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Etymology) Usages(obj *Usage, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Etymology) UsagesByKey(obj *Usage, key POS) bool{
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *Etymology) UsagesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func EtymologyStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func EtymologyAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func EtymologyAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(description), 0)
}
func EtymologyAddUsages(builder *flatbuffers.Builder, usages flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(usages), 0)
}
func EtymologyStartUsagesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EtymologyEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Group struct {
	_tab flatbuffers.Table
}

func GetRootAsGroup(buf []byte, offset flatbuffers.UOffsetT) *Group {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Group{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsGroup(buf []byte, offset flatbuffers.UOffsetT) *Group {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Group{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Group) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Group) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Group) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Group) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Group) Definitions(obj *Definition, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Group) DefinitionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func GroupStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func GroupAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func GroupAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(description), 0)
}
func GroupAddDefinitions(builder *flatbuffers.Builder, definitions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(definitions), 0)
}
func GroupStartDefinitionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func GroupEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Usage struct {
	_tab flatbuffers.Table
}

func GetRootAsUsage(buf []byte, offset flatbuffers.UOffsetT) *Usage {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Usage{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUsage(buf []byte, offset flatbuffers.UOffsetT) *Usage {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Usage{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Usage) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Usage) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Usage) Pos() POS {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return POS(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *Usage) MutatePos(n POS) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func UsageKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &Usage{}
	obj2 := &Usage{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return obj1.Pos() < obj2.Pos()
}

func (rcv *Usage) LookupByKey(key POS, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &Usage{}
		obj.Init(buf, tableOffset)
		val := obj.Pos()
		comp := 0
		if val > key {
			comp = 1
		} else if val < key {
			comp = -1
		}
		if comp > 0 {
			span = middle
		} else if comp < 0 {
			middle += 1
			start += middle
			span -= middle
		} else {
			rcv.Init(buf, tableOffset)
			return true
		}
	}
	return false
}

func (rcv *Usage) Definitions(obj *Definition, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Usage) DefinitionsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Usage) Groups(obj *Group, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Usage) GroupsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func UsageStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func UsageAddPos(builder *flatbuffers.Builder, pos POS) {
	builder.PrependInt8Slot(0, int8(pos), 0)
}
func UsageAddDefinitions(builder *flatbuffers.Builder, definitions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(definitions), 0)
}
func UsageStartDefinitionsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func UsageAddGroups(builder *flatbuffers.Builder, groups flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(groups), 0)
}
func UsageStartGroupsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func UsageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Entry struct {
	_tab flatbuffers.Table
}

func GetRootAsEntry(buf []byte, offset flatbuffers.UOffsetT) *Entry {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Entry{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEntry(buf []byte, offset flatbuffers.UOffsetT) *Entry {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Entry{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Entry) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Entry) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Entry) Key() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func EntryKeyCompare(o1, o2 flatbuffers.UOffsetT, buf []byte) bool {
	obj1 := &Entry{}
	obj2 := &Entry{}
	obj1.Init(buf, flatbuffers.UOffsetT(len(buf)) - o1)
	obj2.Init(buf, flatbuffers.UOffsetT(len(buf)) - o2)
	return string(obj1.Key()) < string(obj2.Key())
}

func (rcv *Entry) LookupByKey(key string, vectorLocation flatbuffers.UOffsetT, buf []byte) bool {
	span := flatbuffers.GetUOffsetT(buf[vectorLocation - 4:])
	start := flatbuffers.UOffsetT(0)
	bKey := []byte(key)
	for span != 0 {
		middle := span / 2
		tableOffset := flatbuffers.GetIndirectOffset(buf, vectorLocation+ 4 * (start + middle))
		obj := &Entry{}
		obj.Init(buf, tableOffset)
		comp := bytes.Compare(obj.Key(), bKey)
		if comp > 0 {
			span = middle
		} else if comp < 0 {
			middle += 1
			start += middle
			span -= middle
		} else {
			rcv.Init(buf, tableOffset)
			return true
		}
	}
	return false
}

func (rcv *Entry) Term() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Entry) Pronunciation() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Entry) See() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Entry) Etymologies(obj *Etymology, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Entry) EtymologiesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func EntryStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func EntryAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func EntryAddTerm(builder *flatbuffers.Builder, term flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(term), 0)
}
func EntryAddPronunciation(builder *flatbuffers.Builder, pronunciation flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(pronunciation), 0)
}
func EntryAddSee(builder *flatbuffers.Builder, see flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(see), 0)
}
func EntryAddEtymologies(builder *flatbuffers.Builder, etymologies flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(etymologies), 0)
}
func EntryStartEtymologiesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func EntryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Dictionary struct {
	_tab flatbuffers.Table
}

func GetRootAsDictionary(buf []byte, offset flatbuffers.UOffsetT) *Dictionary {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Dictionary{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDictionary(buf []byte, offset flatbuffers.UOffsetT) *Dictionary {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Dictionary{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Dictionary) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Dictionary) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Dictionary) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Dictionary) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Dictionary) Entries(obj *Entry, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Dictionary) EntriesByKey(obj *Entry, key string) bool{
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		return obj.LookupByKey(key, x, rcv._tab.Bytes)
	}
	return false
}

func (rcv *Dictionary) EntriesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func DictionaryStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func DictionaryAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func DictionaryAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func DictionaryAddEntries(builder *flatbuffers.Builder, entries flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(entries), 0)
}
func DictionaryStartEntriesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DictionaryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
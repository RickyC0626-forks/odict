// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package types

import (
	"strconv"
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

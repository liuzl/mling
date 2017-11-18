package mling

import (
	"github.com/liuzl/go-porterstemmer"
	"github.com/liuzl/segment"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

type MlingTokenizer struct {
	rmPunct bool
	doFold  bool
	doLower bool
	doStem  bool
	err     error
}

func NewMlingTokenizer(rmPunct, doFold, doLower, doStem bool) *MlingTokenizer {
	return &MlingTokenizer{rmPunct: rmPunct, doFold: doFold, doLower: doLower, doStem: doStem}
}

func DefaultMlingTokenizer() *MlingTokenizer {
	return NewMlingTokenizer(true, true, true, false)
}

func (t *MlingTokenizer) Tokenize(text string) []string {
	var ret []string
	var trans transform.Transformer
	if t.doFold {
		isMn := func(r rune) bool {
			return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
		}
		trans = transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	}
	seg := segment.NewSegmenterDirect([]byte(text))
	for seg.Segment() {
		term := strings.TrimSpace(seg.Text())
		if term == "" {
			continue
		}
		if t.rmPunct {
			r := []rune(term)
			if len(r) == 1 && unicode.IsPunct(r[0]) {
				continue
			}
		}
		if t.doFold {
			res, _, err := transform.String(trans, term)
			if err != nil {
				t.err = err
			} else {
				term = res
			}
		}
		if t.doLower {
			term = strings.ToLower(term)
		}
		if t.doStem {
			term = porterstemmer.StemString(term)
		}
		if term == "" {
			continue
		}
		ret = append(ret, term)
	}
	return ret
}

func SegmentBytes(bytes []byte) []string {
	// utf-8 bytes
	var ret []string
	seg := segment.NewSegmenterDirect(bytes)
	for seg.Segment() {
		term := strings.TrimSpace(seg.Text())
		if term == "" {
			continue
		}
		ret = append(ret, term)
	}
	return ret
}

func Segment(text string) []string {
	return SegmentBytes([]byte(text))
}

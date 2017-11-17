package mling

import (
	"github.com/liuzl/segment"
	"strings"
)

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

package wubi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCode_WuBi86(t *testing.T) {
	testTable := []struct {
		testName string
		char     rune
		codes    []string
	}{
		{
			testName: "在",
			char:     '在',
			codes:    []string{"d", "dhf", "dhfd"},
		},
		{
			testName: "的",
			char:     '的',
			codes:    []string{"r", "rqy", "rqyy"},
		},
		{
			testName: "餰",
			char:     '餰',
			codes:    []string{"wpth", "wyvh"},
		},
	}

	c := New86()
	for _, v := range testTable {
		codes := c.GetCode(v.char)
		require.Equal(t, v.codes, codes)
	}
}

func TestGetCode_WuBi86_Code2Chars(t *testing.T) {
	testTable := []struct {
		testName string
		code     string
		chars    []string
	}{
		{
			testName: "d",
			chars:    []string{"在", "石"},
			code:     "d",
		},
		{
			testName: "r",
			chars:    []string{"白", "的"},
			code:     "r",
		},
		{
			testName: "wpth",
			chars:    []string{"飾", "餰"},
			code:     "wpth",
		},
	}

	c := New86()
	for _, v := range testTable {
		chars := c.GetChar(v.code)
		require.Equal(t, v.chars, chars)
	}
}

func TestGetCode_WuBi86_Codes2Chars(t *testing.T) {
	testTable := []struct {
		testName string
		code     []string
		chars    [][]string
	}{
		{
			testName: "d r",
			chars:    [][]string{{"在", "石"}, {"白", "的"}},
			code:     []string{"d", "r"},
		},
		{
			testName: "wpth p",
			chars:    [][]string{{"飾", "餰"}, {"之", "这"}},
			code:     []string{"wpth", "p"},
		},
	}

	c := New86()
	for _, v := range testTable {
		chars := c.GetChars(v.code)
		require.Equal(t, v.chars, chars)
	}
}

func TestGetCodes_WuBi86(t *testing.T) {
	testTable := []struct {
		testName string
		chars    string
		codes    [][]string
	}{
		{
			testName: "干一行",
			chars:    "干一行，爱一行，一行行，行行行，一行不行，行行不行",
			codes: [][]string{{"fggh"}, {"g", "ggl", "ggll"}, {"tf", "tfhh"}, {"，"}, {"ep", "epd", "epdc"},
				{"g", "ggl", "ggll"}, {"tf", "tfhh"}, {"，"}, {"g", "ggl", "ggll"}, {"tf", "tfhh"}, {"tf", "tfhh"},
				{"，"}, {"tf", "tfhh"}, {"tf", "tfhh"}, {"tf", "tfhh"}, {"，"}, {"g", "ggl", "ggll"}, {"tf", "tfhh"},
				{"i", "gi", "gii"}, {"tf", "tfhh"}, {"，"}, {"tf", "tfhh"}, {"tf", "tfhh"}, {"i", "gi", "gii"},
				{"tf", "tfhh"}},
		},
	}

	c := New86()
	for _, v := range testTable {
		codes := c.GetCodes(v.chars)
		require.Equal(t, v.codes, codes)
	}
}

func TestGetCode_WuBi98(t *testing.T) {
	testTable := []struct {
		testName string
		char     rune
		codes    []string
	}{
		{
			testName: "在",
			char:     '在',
			codes:    []string{"d", "dhf", "dhfd"},
		},
		{
			testName: "的",
			char:     '的',
			codes:    []string{"r", "rqy", "rqyy"},
		},
		{
			testName: "餰",
			char:     '餰',
			codes:    []string{"wvts"},
		},
	}

	c := New98()
	for _, v := range testTable {
		codes := c.GetCode(v.char)
		require.Equal(t, v.codes, codes)
	}
}

func BenchmarkWubi86Char2Codes(b *testing.B) {
	c := New86()
	for i := 0; i < b.N; i++ {
		_ = c.GetCode('一')
	}
}

func BenchmarkWubi86Chars2Codes(b *testing.B) {
	c := New86()
	for i := 0; i < b.N; i++ {
		_ = c.GetCodes("干一行")
	}
}

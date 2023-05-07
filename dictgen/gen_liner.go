package dictgen

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/dlclark/regexp2"
)

var (
	patternGroupNameChar = "char"
	patternGroupNameCode = "code"
)

var (
	ErrLinerInvalidPattern = errors.New("invalid pattern")
)

// NewLiner 生成一个按行解析器
// pattern: 解析器使用得正则表达式，此表达式至少有2个命名分组。
//
//	分组命名格式：
//		char: 表示汉字
//		code: 表示五笔码
func NewLiner(pattern string, debug bool) (Liner, error) {
	rgp, err := regexp2.Compile(pattern, regexp2.None)
	if err != nil {
		return Liner{}, err
	}

	codeIdx := -1
	charIdx := -1
	for _, n := range rgp.GetGroupNames() {
		if n == patternGroupNameCode {
			codeIdx = rgp.GroupNumberFromName(n)
		} else if n == patternGroupNameChar {
			charIdx = rgp.GroupNumberFromName(n)
		}
	}

	if codeIdx < 0 || charIdx < 0 {
		return Liner{}, ErrLinerInvalidPattern
	}

	return Liner{
		regexp:  rgp,
		charIdx: charIdx,
		codeIdx: codeIdx,
		debug:   debug,
	}, nil
}

// Liner 逐行扫描生成器
type Liner struct {
	regexp *regexp2.Regexp

	debug bool

	charIdx int
	codeIdx int
}

func (l Liner) findAllMatch(s string) []*regexp2.Match {
	var matches []*regexp2.Match
	m, _ := l.regexp.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m)
		m, _ = l.regexp.FindNextMatch(m)
	}
	return matches
}

// ParseSingle 解析单行数据
func (l Liner) ParseSingle(raw []byte) (codes []string, chars []string, err error) {
	matches := l.findAllMatch(string(raw))
	for _, match := range matches {
		group := match.GroupByNumber(l.charIdx)
		for _, capture := range group.Captures {
			chars = append(chars, capture.String())
		}
		group = match.GroupByNumber(l.codeIdx)
		for _, capture := range group.Captures {
			codes = append(codes, capture.String())
		}
	}

	if len(codes) > 1 && len(chars) > 1 {
		return nil, nil, ErrMultiCharsAndMultiCodes
	}

	if len(codes) == 0 || len(chars) == 0 {
		return nil, nil, ErrMissCharOrCode
	}

	return
}

// ParseAll 解析全量数据
func (l Liner) ParseAll(raw io.Reader) (Dict, error) {
	scanner := bufio.NewScanner(raw)

	dict := NewDict()
	for scanner.Scan() {
		bts := scanner.Bytes()
		str := string(bts)
		codes, chars, err := l.ParseSingle(bts)
		if err != nil {
			l.log("warning: %s %s\n", err.Error(), str)
			continue
		}

		if len(codes) == 1 {
			dict.AddCodeWithChars(codes[0], chars)
		} else {
			dict.AddCharWithCodes(chars[0], codes)
		}
	}

	return dict, nil
}

func (l Liner) log(format string, args ...any) {
	if !l.debug {
		return
	}
	fmt.Printf(format, args...)
}

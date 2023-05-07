package wubi

import (
	"errors"
	"sort"
	"strings"

	"github.com/gopherlib/wubi/dict"
)

// Dictionary 五笔码表
type Dictionary map[string]string

// Version 五笔版本
type Version string

// 默认支持的版本
const (
	Version86 Version = "86"
	Version98 Version = "98"
	Version06 Version = "06"
)

var (
	dictList = map[Version]Dictionary{
		Version86: dict.Dict86,
		Version98: dict.Dict98,
		Version06: dict.Dict06,
	}
)

// RegisterVersion 注册自定义版本，可以覆盖默认的码表数据
func RegisterVersion(ver Version, dict Dictionary) {
	dictList[ver] = dict
}

// NewConvertOf86 86版
func NewConvertOf86() *Convert {
	return &Convert{
		dictChar2Code: dict.Dict86,
		dictCode2Char: reverseDict(dict.Dict86),
	}
}

// NewConvertOf98 98版
func NewConvertOf98() *Convert {
	return &Convert{
		dictChar2Code: dict.Dict98,
		dictCode2Char: reverseDict(dict.Dict98),
	}
}

// NewConvertOf06 新世纪版
func NewConvertOf06() *Convert {
	return &Convert{
		dictChar2Code: dict.Dict06,
		dictCode2Char: reverseDict(dict.Dict06),
	}
}

// NewConvert 创建一个转换器实例
func NewConvert(ver Version) (*Convert, error) {
	if d, ok := dictList[ver]; ok {
		return &Convert{
			dictChar2Code: d,
			dictCode2Char: reverseDict(d),
		}, nil
	}

	return nil, errors.New("invalid version")
}

// NewConvertWithDict 使用自定义码表创建转换器实例
func NewConvertWithDict(d Dictionary) *Convert {
	return &Convert{
		dictChar2Code: d,
		dictCode2Char: reverseDict(d),
	}
}

// Convert 五笔码转换器
type Convert struct {
	dictChar2Code Dictionary
	dictCode2Char Dictionary
}

// GetCode 获取单字的五笔码
func (c Convert) GetCode(char rune) []string {
	if code, ok := c.dictChar2Code[string(char)]; ok {
		codes := strings.Split(code, ",")
		c.sort(codes)
		return codes
	}

	return nil
}

// GetCodes 获取字符串的五笔码列表
func (c Convert) GetCodes(chars string) [][]string {
	codes := make([][]string, 0, len(chars))
	for _, r := range []rune(chars) {
		code := c.GetCode(r)
		if len(code) == 0 {
			// 如果字符没有对应的五笔码，则原样返回
			codes = append(codes, []string{string(r)})
		} else {
			// 对五笔码进行排序，简码在前，全码在后
			if len(code) > 0 {
				c.sort(code)
			}
			codes = append(codes, code)
		}
	}
	return codes
}

// GetChar 获取单个五笔码对应的汉字
func (c Convert) GetChar(code string) []string {
	if char, ok := c.dictCode2Char[code]; ok {
		chars := strings.Split(char, ",")
		sort.Sort(sort.StringSlice(chars))
		return chars
	}

	return nil
}

// GetChars 获取五笔码列表对应的汉字
func (c Convert) GetChars(codes []string) [][]string {
	chars := make([][]string, 0, len(codes))
	for _, code := range codes {
		cs := c.GetChar(code)
		if len(cs) == 0 {
			// 如果五笔码没有对应的汉字，则原样返回
			chars = append(chars, []string{code})
		} else {
			if len(cs) > 0 {
				sort.Sort(sort.StringSlice(cs))
			}
			chars = append(chars, cs)
		}
	}
	return chars
}

func (c Convert) sort(codes []string) {
	sort.Sort(codeSlice(codes))
}

// reverseDict 翻转字典
func reverseDict(dict Dictionary) Dictionary {
	newArrDict := make(map[string][]string, len(dict))
	newDict := make(Dictionary, len(dict))

	// key,val 原始字典中的 key => val
	for key, val := range dict {
		// 原始字典的val是以 , 分隔的数组字符串, 这里还原为数组格式
		vals := strings.Split(val, ",")
		for _, v := range vals {
			// 对每个val进行反转
			if _, ok := newArrDict[v]; ok {
				newArrDict[v] = append(newArrDict[v], key)
			} else {
				newArrDict[v] = []string{key}
			}
		}
	}

	for key, vals := range newArrDict {
		newDict[key] = strings.Join(vals, ",")
	}
	return newDict
}

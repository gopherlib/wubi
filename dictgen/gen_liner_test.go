package dictgen

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMultiData(t *testing.T) {
	liner, err := NewLiner(`^\d+\t(?<code>[a-zA-Z]+)\t(?<char>[一-龥〇])\t\d+`, false)
	require.NoError(t, err)

	buf := bytes.NewReader([]byte(`type	keys	words	sets
1	a	工	0
1	b	了	0
1	c	以	0`))
	dict, err := liner.ParseAll(buf)
	require.NoError(t, err)
	fmt.Println(dict)
	_ = dict
}

func TestParseSingle(t *testing.T) {
	liner, err := NewLiner(`^(?<code>[a-zA-Z]+)|\s+~*(?<char>[一-龥〇])(?![一-龥〇])`, false)
	require.NoError(t, err)

	testTables := []struct {
		testName string
		raw      string
		codes    []string
		chars    []string
		err      error
	}{
		{
			testName: "单码两字",
			raw:      `a 工 戈`,
			codes:    []string{"a"},
			chars:    []string{"工", "戈"},
			err:      nil,
		},
		{
			testName: "多字母单码单字",
			raw:      `aah 芽`,
			codes:    []string{"aah"},
			chars:    []string{"芽"},
			err:      nil,
		},
		{
			testName: "多字母单码单词组",
			raw:      `aagk 工整`, // 只支持单字转换，忽略词组，所以只有词组时会返回ErrMissCharOrCode
			codes:    nil,
			chars:    nil,
			err:      ErrMissCharOrCode,
		},
		{
			testName: "多字母单码单字与词组",
			raw:      `aaaa 工 恭恭敬敬`, // 只支持单字转换，忽略词组，所以只返回单字
			codes:    []string{"aaaa"},
			chars:    []string{"工"},
			err:      nil,
		},
		{
			testName: "〇",
			raw:      `llll 〇`,
			codes:    []string{"llll"},
			chars:    []string{"〇"},
			err:      nil,
		},
		{
			testName: "~",
			raw:      `lllf 田园 ~壘 ~壨 ~畾`,
			codes:    []string{"lllf"},
			chars:    []string{"壘", "壨", "畾"},
			err:      nil,
		},
	}

	for _, item := range testTables {
		t.Run(item.testName, func(t *testing.T) {
			codes, chars, err := liner.ParseSingle([]byte(item.raw))
			require.Equal(t, item.err, err)
			require.Equal(t, item.codes, codes)
			require.Equal(t, item.chars, chars)
		})
	}
}

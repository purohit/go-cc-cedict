package cedict

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(nil)
	tests := []struct {
		in  string
		exp Entry
	}{
		{
			in: "人怕出名豬怕肥 人怕出名猪怕肥 [ren2 pa4 chu1 ming2 zhu1 pa4 fei2] /lit. people fear getting famous like pigs fear fattening up (for the slaughter)/fig. fame has its price/",
			exp: Entry{
				Traditional: "人怕出名豬怕肥",
				Simplified:  "人怕出名猪怕肥",
				Pinyin:      "ren2 pa4 chu1 ming2 zhu1 pa4 fei2",
				Definitions: []Definition{
					{Text: "lit. people fear getting famous like pigs fear fattening up (for the slaughter)"},
					{Text: "fig. fame has its price"},
				},
			},
		},
	}
	for _, t := range tests {
		got, err := ParseEntry(t.in)
		assert.Nil(err)
		assert.Equal(t.exp, *got)
	}
}

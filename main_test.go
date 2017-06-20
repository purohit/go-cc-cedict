package cedict

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)
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
				Abbreviations: []Abbreviation{},
				Variants:      []Variant{},
				References:    []Reference{},
			},
		},
	}
	for _, t := range tests {
		got, err := ParseEntry(t.in)
		assert.Nil(err)
		assert.EqualValues(t.exp, *got)
	}
}

func TestParseDictionary(t *testing.T) {
	assert := assert.New(t)
	// Picked 10 random definitions and a few comment lines.
	r := strings.NewReader(
		`# CC-CEDICT
# Community maintained free Chinese-English dictionary.
栗斑腹鵐 栗斑腹鹀 [li4 ban1 fu4 wu2] /(bird species of China) Jankowski's bunting (Emberiza jankowskii)/
粉轉黑 粉转黑 [fen3 zhuan3 hei1] /(Internet slang) to go from being an admirer to being a detractor/
震源機制 震源机制 [zhen4 yuan2 ji1 zhi4] /focal mechanism of earthquake/
道孚縣 道孚县 [Dao4 fu2 xian4] /Dawu county (Tibetan: rta 'u rdzong) in Garze Tibetan autonomous prefecture 甘孜藏族自治州[Gan1 zi1 Zang4 zu2 zi4 zhi4 zhou1], Sichuan (formerly in Kham province of Tibet)/
心醉 心醉 [xin1 zui4] /enchanted/fascinated/charmed/
膠印 胶印 [jiao1 yin4] /offset printing/
夯砣 夯砣 [hang1 tuo2] /rammer/tamper/
明理 明理 [ming2 li3] /sensible/reasonable/an obvious reason, truth or fact/to understand the reason or reasoning/
少安毋躁 少安毋躁 [shao3 an1 wu2 zao4] /keep calm, don't get excited/don't be impatient/
興山縣 兴山县 [Xing1 shan1 xian4] /Xingshan county in Yichang 宜昌[Yi2 chang1], Hubei/
`)
	dict, err := ParseDictionary(r)
	assert.Nil(err)
	assert.Equal(len(dict.Entries), 10)
}

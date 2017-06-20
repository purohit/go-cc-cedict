package cedict

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Dictionary represents a parsed CC-CEDICT dictionary.
type Dictionary struct {
	Entries []*Entry
}

// Reference is used when a dictionary item
// says "see: xxx|xxx"
type Reference struct {
	Simplified  string
	Traditional string
}

// Variant is used when a dictionary item
// says "[old/erhua] variant of xxx|xxx"
type Variant struct {
	Erhua bool
	Old   bool
}

// Definition is the English definition of
// this word, along with some markers on register and
// regional dialects.
type Definition struct {
	Text           string
	Idiom          bool
	FigureofSpeech bool
	Slang          bool
	Colloquial     bool
	// Definition only applies regionally:
	Cantonese bool
	Taiwanese bool
}

type Abbreviation struct {
}

// Entry represents a single line in the dictionary
type Entry struct {
	Simplified    string
	Traditional   string
	Pinyin        string
	Definitions   []Definition
	Abbreviations []Abbreviation
	Variants      []Variant
	References    []Reference
}

func (e *Entry) String() string {
	if len(e.Definitions) > 0 {
		return fmt.Sprintf("%s\t%s\t%s",
			e.Simplified, e.Pinyin, e.Definitions[0].Text)
	}
	// TODO: Better default printing for words without definitions
	return fmt.Sprintf("%s\t%s", e.Simplified, e.Pinyin)
}

// ParseDictionary reads a CC-CEDICT buffer, and parses
// all entries into Entries, skipping comments and breaking
// on errors.
func ParseDictionary(r io.Reader) (*Dictionary, error) {
	scanner := bufio.NewScanner(r)
	dict := &Dictionary{Entries: make([]*Entry, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") { // Ignore comments
			continue
		}
		entry, err := ParseEntry(line)
		if err != nil {
			return nil, fmt.Errorf("parsing line [%s]: %s", line, err)
		}
		dict.Entries = append(dict.Entries, entry)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading input:", err)
	}
	return dict, nil
}

func newEntry() *Entry {
	return &Entry{
		Definitions:   make([]Definition, 0),
		Abbreviations: make([]Abbreviation, 0),
		Variants:      make([]Variant, 0),
		References:    make([]Reference, 0),
	}
}

var entryRegexp = regexp.MustCompile(`^(.*) (.*) \[(.*)\] /(?:(.*)/)+`)

// ParseEntry parses a single dictionary line
func ParseEntry(line string) (*Entry, error) {
	matches := entryRegexp.FindStringSubmatch(line)
	if matches == nil {
		return nil, fmt.Errorf("Entry doesn't match regular expression")
	}
	e := newEntry()
	e.Traditional = matches[1]
	e.Simplified = matches[2]
	e.Pinyin = strings.ToLower(matches[3])
	defs := strings.Split(matches[4], "/")
	if len(defs) == 0 {
		return nil, fmt.Errorf("No definitions found")
	}
	for _, def := range defs {
		e.Definitions = append(e.Definitions, Definition{
			Text: def,
		})
	}
	return e, nil
}

package cedict

import (
	"bufio"
	"fmt"
	"os"
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
	Text       string
	Idiom      bool
	Slang      bool
	Colloquial bool
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

// ParseDictionary reads the CC-CEDICT .txt file, and parses
// all entries into Entries, skipping comments and breaking
// on errors.
func ParseDictionary(filename string) (*Dictionary, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
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

func ParseEntry(line string) (*Entry, error) {
	parts := strings.SplitN(line, " ", 4)
	if len(parts) < 4 {
		return nil, fmt.Errorf("Not enough parts in dict")
	}
	e := newEntry()
	e.Traditional = parts[0]
	e.Simplified = parts[1]
	e.Pinyin = strings.ToLower(strings.Trim(parts[2], "[]"))
	parts = strings.Split(parts[3], "/")
	if len(parts) < 3 {
		return nil, fmt.Errorf("No definitions found")
	}
	// Throw away first and last parts since definition start has /
	parts = parts[1:]
	parts = parts[:len(parts)-1]
	for _, part := range parts[1:] {
		e.Definitions = append(e.Definitions, Definition{
			Text: part,
		})
	}
	return e, nil
}

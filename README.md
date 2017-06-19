# go-cc-cedict
[![Build
Status](https://travis-ci.org/purohit/go-cc-cedict.svg?branch=master)](https://travis-ci.org/purohit/go-cc-cedict)
[![GoDoc](https://godoc.org/github.com/purohit/go-cc-cedict?status.svg)](https://godoc.org/github.com/purohit/go-cc-cedict)

A parser for the wonderful Chinese <-> English [CC-CEDICT](https://cc-cedict.org/wiki/) project that returns structured entries including definitions, Pinyin, Simplified and Traditional variants, references and abbreviations.

The dictionary as written was not meant for machine consumption, but I'm doing my damndest
to parse it for applications.

## Usage

You can either use a convenience function to load up the whole dictionary given the [raw data file](https://www.mdbg.net/chinese/dictionary?page=cedict) or parse individual lines, as you
please.

```go
import (
    "fmt"

    cedict "github.com/purohit/go-cc-cedict"
)

dict, _ := cedict.ParseDictionary("/home/data/cedict_1_0_ts_utf-8_mdbg.txt")
for i, entry := range dict.Entries {
    fmt.Println(e)
    if i > 10 {
        break
    }
}

// Alternatively, parse the lines yourself:

entry, _ := cedict.ParseEntry("企投 企投 [qi4 tou2] /to have fun (Taiwanese, POJ pr. [chhit-thô])/")
fmt.Println(e.Simplified, e.Pinyin, e.EnglishDefinitions)
```

## Contributing

I'd appreciate contributions that help make `Entry`s even more specific
-- there are a lot of exceptions in the definitions since they're somewhat
free-form. For example, adding a specific "Abbreviations" field to the struct
where all abbreviations are parsed out instead of remaining in
`Definitions`.

Please add your own tests and make sure the current ones pass! I'm okay with reasonable
backwards-incompatibility (I expect this library to be low-usage) via modifications to `Entry`.

## License

The license for this parsing software itself is MIT, however the actual dictionary data is licensed under a CC license (see the [main project
page](https://cc-cedict.org/wiki/) for more info).

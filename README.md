### go-cc-cedict

A parser for the wonderful [CC-CEDICT](https://cc-cedict.org/wiki/) project that parses the dictionary and
returns structured `Entry` structs with definitions, Pinyin, Simplified and
Traditional variants, as well as links and abbreviations.

The dictionary was not meant for machine consumption, but I'm doing my damndest
to try to parse it for applications.

### Usage

You can either use a convenience function to load up the whole dictionary given the [raw text data file](https://www.mdbg.net/chinese/dictionary?page=cedict) or parse individual lines, as you
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

### Contributing

I'd appreciate contributions that help make `Entry` structs even more specific
-- there are a lot of exceptions in the definitions since they're somewhat
free-form. For example, adding a specific "Abbreviations" field to the struct
where all abbreviations are parsed out instead of remaining in
"EnglishDefinitions".

Please add your own tests and make sure the current ones pass! I'm okay with
backwards-incompatibility (I expect this library to be low-usage) through modifying the structure of `Entry` as long as it seems reasonable.

### License

The license for this parsing software itself is MIT, however the actual
raw dictionary data is licensed under a CC license (see the [main project
page](https://cc-cedict.org/wiki/) for more info).

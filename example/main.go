package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/taimoorgit/moredefaults"
)

type Gender string

type Sample struct {
	Name    string `default:"John Smith"`
	Age     int    `default:"27"`
	Gender  Gender `default:"m"`
	Working bool   `default:"true"`

	SliceInt    []int    `default:"[1, 2, 3]"`
	SlicePtr    []*int   `default:"[1, 2, 3]"`
	SliceString []string `default:"[\"a\", \"b\"]"`

	MapNull            map[string]int          `default:"{}"`
	Map                map[string]int          `default:"{\"key1\": 123}"`
	MapOfStruct        map[string]OtherStruct  `default:"{\"Key2\": {\"Foo\":123}}"`
	MapOfPtrStruct     map[string]*OtherStruct `default:"{\"Key3\": {\"Foo\":123}}"`
	MapOfStructWithTag map[string]OtherStruct  `default:"{\"Key4\": {\"Foo\":123}}"`

	Struct    OtherStruct  `default:"{\"Foo\": 123}"`
	StructPtr *OtherStruct `default:"{\"Foo\": 123}"`

	NoTag    OtherStruct // Recurses into a nested struct by default
	NoOption OtherStruct `default:"-"` // no option
}

type OtherStruct struct {
	Hello  string `default:"world"` // Tags in a nested struct also work
	Foo    int    `default:"-"`
	Random int    `default:"-"`
}

// SetDefaults implements moredefaults.Setter interface
func (s *OtherStruct) SetDefaults() {
	if moredefaults.CanUpdate(s.Random) { // Check if it's a zero value (recommended)
		s.Random = rand.Int() // Set a dynamic value
	}
}

func main() {
	obj := &Sample{}
	if err := moredefaults.Set(obj); err != nil {
		panic(err)
	}

	out, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// Output:
	// {
	// 	"Name": "John Smith",
	// 	"Age": 27,
	// 	"Gender": "m",
	// 	"Working": true,
	// 	"SliceInt": [
	// 		1,
	// 		2,
	// 		3
	// 	],
	// 	"SlicePtr": [
	// 		1,
	// 		2,
	// 		3
	// 	],
	// 	"SliceString": [
	// 		"a",
	// 		"b"
	// 	],
	// 	"MapNull": {},
	// 	"Map": {
	// 		"key1": 123
	// 	},
	// 	"MapOfStruct": {
	// 		"Key2": {
	// 			"Hello": "world",
	// 			"Foo": 123,
	// 			"Random": 5577006791947779410
	// 		}
	// 	},
	// 	"MapOfPtrStruct": {
	// 		"Key3": {
	// 			"Hello": "world",
	// 			"Foo": 123,
	// 			"Random": 8674665223082153551
	// 		}
	// 	},
	// 	"MapOfStructWithTag": {
	// 		"Key4": {
	// 			"Hello": "world",
	// 			"Foo": 123,
	// 			"Random": 6129484611666145821
	// 		}
	// 	},
	// 	"Struct": {
	// 		"Hello": "world",
	// 		"Foo": 123,
	// 		"Random": 4037200794235010051
	// 	},
	// 	"StructPtr": {
	// 		"Hello": "world",
	// 		"Foo": 123,
	// 		"Random": 3916589616287113937
	// 	},
	// 	"NoTag": {
	// 		"Hello": "world",
	// 		"Foo": 0,
	// 		"Random": 6334824724549167320
	// 	},
	// 	"NoOption": {
	// 		"Hello": "",
	// 		"Foo": 0,
	// 		"Random": 0
	// 	}
	// }
}

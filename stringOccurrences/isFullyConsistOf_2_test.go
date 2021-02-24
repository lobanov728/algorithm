package stringOccurrences

import (
	"fmt"
	"testing"
)

var elems1 = []string{
	"h",
	"he",
	"li",
	"be",
	"b",
	"c",
	"n",
	"o",
	"f",
	"ne",
	"na",
	"mg",
	"al",
	"si",
	"p",
	"s",
	"cl",
	"ar",
	"cl",
	"ar",
	"ti",
}

var inputs1 = []string {
	"clicarti",
	"clicartit",
	//"he",
	//"cbe",
	//"cberg",
	//"cb",
	//"clx",
	//"xcl",
	//"cbx",
	//"mgb",
	//"mgxmgmg",
	//"mgbxmgbb",
	//"nac",
	//"hoh",
	//"hxoh",
	//"",
	//"x",
	//"xx",
	//"xxx",
	//"c",
	//"cc",
	//"ccc",
	//"cxb",
	//"xcb",
	//"mgcl",
}

func TestIsFullyConsistOf2(t *testing.T) {
	for _, input := range inputs1 {
		fmt.Println("word", input, isFullyConsistOf_2(input, elems1))
	}
}
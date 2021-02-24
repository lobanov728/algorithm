package stringOccurrences
//
//import (
//	"fmt"
//	"testing"
//)
//
//var elems = []string{
//	"h",
//	"he",
//	"li",
//	"be",
//	"b",
//	"c",
//	"n",
//	"o",
//	"f",
//	"ne",
//	"na",
//	"mg",
//	"al",
//	"si",
//	"p",
//	"s",
//	"cl",
//	"ar",
//	"cli",
//	"art",
//}
//
//var inputs = []string {
//	"clizcart",
//	//"he",
//	//"cbe",
//	//"cberg",
//	//"cb",
//	//"clx",
//	//"xcl",
//	//"cbx",
//	//"mgb",
//	//"mgxmgmg",
//	//"mgbxmgbb",
//	//"nac",
//	//"hoh",
//	//"hxoh",
//	//"",
//	//"x",
//	//"xx",
//	//"xxx",
//	//"c",
//	//"cc",
//	//"ccc",
//	//"cxb",
//	//"xcb",
//	//"mgcl",
//}
//
//func TestIsFullyConsistOf(t *testing.T) {
//	for _, input := range inputs {
//		fmt.Println("word", input, IsFullyConsistOf(input, elems, 3))
//	}
//}
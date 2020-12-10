package algorithm

import (
	"fmt"
	"reflect"
	"strconv"
)

type boxable interface {
	IsBox() bool
}

type Item struct {
	name string
	isBox bool
	isKey bool
	Items []*Item
}

func (i *Item) IsBox() bool {
	return i.isBox
}

func (i Item) IsKey() bool {
	return i.isKey
}

func nameless(input boxable)  {
	switch input.(type) {

	}
}

func CreateBoxes() Item {
	rootBox := Item{isBox: true, name:"rootBox"}
	fmt.Println(reflect.TypeOf(rootBox))

	nameless(&rootBox)
	box11 := Item{isBox: true, name:"box11"}
	box12 := Item{isBox: true, name:"box12"}
	box13 := Item{isBox: true, name:"box13"}
	box121 := Item{isBox: true, name:"box121"}
	box122 := Item{isBox: true, name:"box122"}
	box131 := Item{isBox: true, name:"box131"}
	box132 := Item{isBox: true, name:"box132"}
	box133 := Item{isBox: true, name:"box133"}
	key := Item{isKey: true, name:"key"}

	rootBox.Items = append(rootBox.Items, &box11, &box12, &box13)
	box12.Items = append(box12.Items, &box121, &box122)
	box13.Items = append(box13.Items, &box131, &box132, &box133)
	box131.Items = append(box131.Items, &key)

	return rootBox
}

var founded bool
func LookForKey(item Item) string  {
	for _, inItem := range item.Items {
		if founded {
			break
		}
		fmt.Println("look in", inItem.name)
		if inItem.isKey {
			founded = true
			return "founded"
		} else {
			LookForKey(*inItem)
		}
	}
	if founded {
		return "founded"
	}
	return strconv.FormatBool(founded) + item.name + " not found"
}
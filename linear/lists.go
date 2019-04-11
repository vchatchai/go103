package linear

import (
	"container/list"
	"fmt"
)

func List() {
	var listInt list.List
	listInt.PushBack(11)
	listInt.PushBack(22)
	listInt.PushBack(33)
	for element := listInt.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}

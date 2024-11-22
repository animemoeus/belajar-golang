package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("arter")
	data.PushBack("tendean")
	data.PushBack(1337)
	data.PushBack("angie")

	//var head *list.Element = data.Front()
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

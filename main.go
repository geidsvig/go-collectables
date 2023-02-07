package main

import (
	"log"
	"reflect"

	collectables "github.com/geidsvig/go-collectables/lib"
)

type Collectable struct {
	name  string
	value int
}

func main() {
	log.Println("lets go")
	T := reflect.TypeOf(Collectable{})

	list := collectables.Collection{}
	item1 := Collectable{name: "A", value: 1}
	item2 := Collectable{name: "K", value: 13}
	err := list.Append(T, item1)
	if err != nil {
		log.Println("Failed to append ", item1, err)
	}
	err = list.Append(T, item2)
	if err != nil {
		log.Println("Failed to append ", item2, err)
	}
	index := list.IndexOf(item1)
	log.Println("index of A ", index)
	index2 := list.IndexOf(item2)
	log.Println("index of K ", index2)
	log.Println("length of list ", list.Length())
	log.Println(list)

	item3 := Collectable{name: "Q", value: 12}
	err = list.Prepend(T, item3)
	if err != nil {
		log.Println("Failed to prepend ", item3, err)
	}
	head, err := list.GetFirst(T)
	if err != nil {
		log.Println("Failed to get first ", err)
	}
	log.Println(head)
	list.RemoveLast()
	log.Println("Removed last ", list)
	list.RemoveAt(1)
	log.Println("Removed index 1 ", list)
	log.Println(list)
	log.Println("is empty ", list.IsEmpty())

	err = list.Append(T, 1)
	if err != nil {
		log.Println("Failed to append ", 1, err)
	}
	log.Println(list) // i added in safety with reflection

	log.Println("index of wrong type ", list.IndexOf(2))

}

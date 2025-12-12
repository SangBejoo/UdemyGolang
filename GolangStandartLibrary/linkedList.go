package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack(1)
	data.PushBack(2)
	data.PushBack(3)
	data.PushBack("INi gimana?")

	//fmt.Println(data.Front())
	//fmt.Println(data.Back())
	//fmt.Println(data)

	head := data.Front()
	fmt.Println(head.Value) // 1

	next := head.Next()
	fmt.Println(next.Value) // 2

	next = next.Next()
	fmt.Println(next.Value) // 3

	//for e := data.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}
	antrian := list.New()
	dataAntrian := []string{"teh", "kopi"}

	antrian.PushBack(dataAntrian[0])
	antrian.PushBack(dataAntrian[1])

	antrianPertama := antrian.Front()
	fmt.Println(antrianPertama.Value)

	antrianKedua := antrianPertama.Next()
	fmt.Println(antrianKedua.Value)

	for e := antrian.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

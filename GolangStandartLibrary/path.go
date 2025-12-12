package main

import (
	"fmt"
	"path"
)

func main() {
	// path package digunakan untuk memanipulasi path file dan direktori
	fmt.Println(path.Dir("Hello/world.go"))
	fmt.Println(path.Base("Hello/world.go"))
	fmt.Println(path.Ext("Hello/world.go"))
	fmt.Println(path.Join("Hello", "world.go", "main.go"))
	fmt.Println(path.IsAbs("Hello/world.go"))

	//filepath package digunakan untuk memanipulasi path file dan direktori sesuai dengan sistem operasi

}

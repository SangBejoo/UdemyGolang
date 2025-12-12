package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("Hello/world.go"))
	fmt.Println(filepath.Base("Hello/world.go"))
	fmt.Println(filepath.Ext("Hello/world.go"))
	fmt.Println(filepath.Join("Hello", "world.go", "main.go"))
	fmt.Println(filepath.IsAbs("Hello/world.go"))
	fmt.Println(filepath.IsLocal("Hello/world.go"))
}

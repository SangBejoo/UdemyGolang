package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		User{"Jane", 30},
		User{"John", 32},
		User{"Paul", 31},
		User{"Paul", 34},
		User{"Jack", 33},
		User{"Jack", 35},
	}
	fmt.Println(users)

	sort.Sort(UserSlice(users))
}

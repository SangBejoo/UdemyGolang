package main

func ReverseString(s []byte) {
	size := len(s)

	for i := 0; i < size/2; i++ {
		s[i], s[size-1-i] = s[size-1-i], s[i]
	}
	return
}

func main() {
	data := []byte("hello")
	ReverseString(data)
	println(string(data)) // Output: "olleh"
}

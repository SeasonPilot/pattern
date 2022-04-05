package main

import "fmt"

func main() {
	//b := &IntBuilder{}
	b := &StrBuilder{}
	d := NewDirector(b)
	d.Make()
	fmt.Println(b.GetResult())
}

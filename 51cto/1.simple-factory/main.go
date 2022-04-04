package main

import "fmt"

type Api interface {
	Say(string) string
}

// NewApi 简单工厂模式，核心代码
func NewApi(s string) Api {
	if s == "ch" {
		return Chinese{}
	} else if s == "en" {
		return English{}
	} else {
		return nil
	}
}

type Chinese struct{}

func (Chinese) Say(s string) string {
	return "你好," + s
}

type English struct{}

func (English) Say(s string) string {
	return "Hello," + s
}

func main() {
	api := NewApi("ch")
	fmt.Println(api.Say("yincheng"))

	//api := NewApi("en")
	//fmt.Println(api.Say("yincheng"))
}

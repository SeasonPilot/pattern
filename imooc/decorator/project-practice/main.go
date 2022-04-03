package main

import (
	"fmt"
	"log"
	"net/http"
)

// 具体装饰器逻辑
func logging(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("URL: %s ,METHOD: %s", r.URL, r.Method)
		// 调用被装饰类的方法，传进来的是具体的类，这里使用抽象类进行调用方法
		handler.ServeHTTP(w, r)
	}
}

// 实例化对象。具体组件，抽象组件是 Handler ，HandlerFunc 类型 实现了 Handler 接口，hello 是 HandlerFunc 类型
func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "hello")
}

func main() {
	http.HandleFunc("/", logging(hello))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

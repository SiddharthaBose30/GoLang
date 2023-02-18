package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//  defer_()
	// defer2()
	// panic_call()
	// panic_defer_both()
	// Web_Application()
	recover_func()
	println("hi")
}

func defer_() {
	defer println("1")
	defer println("2")
	defer println("3")
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	println(res)
	robots, err := ioutil.ReadAll(res.Body)
	println("", res.StatusCode, "\n", res.Request.Method, "\n", res.Status, "\n", res.Request, "\n", res.Request.URL, "\n", res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}

func defer2() {
	a := 2
	defer println(a)
	a = 3
}
func panic_call() {
	println(1)
	panic("3")
	println(2)
}
func panic_defer_both() {
	println(1)
	defer println(2)
	println(3)
	panic("wrong")

}

func Web_Application() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

}
func recover_func() {
	println(2)
	defer func() {
		if err := recover(); err != nil {
			println("carry on")
			panic("Again")
		}
	}()
	panic("Wrong")

}

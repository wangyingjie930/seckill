package main

import (
	"fmt"
	"net/http"
	"sync"
)

var sum int64 = 0

//互斥锁
var mutex sync.Mutex

//预存商品数量
var productNum int64 = 1000000

func GetOneProduct() bool {
	mutex.Lock()
	defer mutex.Unlock()

	if sum < productNum {
		sum += 1
		fmt.Printf("%d\n", sum)
		return true
	}
	return false
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if GetOneProduct() {
		w.Write([]byte("true"))
		return
	}
	w.Write([]byte("false"))
	return
}

func main() {
	http.HandleFunc("/getOne", GetProduct)
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"github.com/kataras/iris"
	"log"
	"net/http"
	"time"
)

func main() {
	n := 1000
	c := make(chan int64)
	url := "http://www.baidu.com"
	for i := 0; i < n; i++ {
		go doHttpGet(i+1, url, c)
	}

	var max, min, ave, total, index int64
	for t := range c {
		if t > max {
			max = t
		}
		if t < min || min == 0 {
			min = t
		}
		total += t
		index++
		fmt.Printf("[%4d] %4d %3d %4d\n", index, max, min, t)
	}

	ave = total / int64(n)
	qps := 1000 * n / int(ave)
	fmt.Printf("max: %dms, ave: %dms, min: %dms, qps: %d\n", max, ave, min, qps)
	//fmt.Println(msg)
	//doHttpGet(1, "http://localhost:8080/api")
}

func doHttpGet(i int, url string, in chan int64) {
	t1 := time.Now().UnixNano()
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != iris.StatusOK {
		log.Println(resp.StatusCode)
	}
	t2 := time.Now().UnixNano()
	t := (t2 - t1) / 1e6
	//fmt.Printf("[%4d] %s %dms", i, url, t)
	in <- t
}

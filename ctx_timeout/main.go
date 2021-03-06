package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

func process() {
	ctx, cancle := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancle()

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)

	//req, err := http.NewRequest("GET", "http://www.google.com", nil)
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println("http request failed ,err:", err)
		return
	}

	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		c <- pack
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		res := <-c
		fmt.Printf("timeout  2s ,err :%s\n", res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response:%s\n", out)
	}

}

func main() {
	process()
}

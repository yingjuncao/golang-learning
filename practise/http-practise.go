package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	Url = "http://www.baidu.com"
)

func main() {
	client := &http.Client{}
	// 发送一个请求
	req, err := http.NewRequest("POST", Url, strings.NewReader("key=value"))
	// 必须加 close
	defer req.Body.Close()
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	// 增加header 或者 cookes
	req.Header.Add("Content-Type", "application/json")

    // 增加 cookes
    cookie1 := &http.Cookie{Name: "name", Value: "caoyingjun"}
	req.AddCookie(cookie1)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Print(string(data))
}

package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time int64  `json:"time"`
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m) // b == []byte(`{"name":"Alice","body":"Hello","time":1294706395881547000}`)
	fmt.Println(string(b))

	var m1 Message
	if err := json.Unmarshal(b, &m1); err == nil {
		fmt.Println(m1)
	}
}

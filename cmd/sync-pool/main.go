package main

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

type SomeEl struct {
	Count int    `json:"count"`
	Text  string `json:"text"`
}

func NoPool(w io.Writer, data []byte) {
	var el SomeEl
	_ = json.Unmarshal(data, &el)
	_, _ = w.Write([]byte(fmt.Sprintf("%v", el)))
}

var pool = sync.Pool{
	New: func() any {
		return &SomeEl{}
	},
}

func WithPool(w io.Writer, data []byte) {
	el := pool.Get().(*SomeEl)
	_ = json.Unmarshal(data, &el)
	_, _ = w.Write([]byte(fmt.Sprintf("%v", el)))
	pool.Put(el)
}

package main

import (
	"encoding/json"
	"fmt"
)

type Word struct {
	Word string `json:"word"`
	Zhuyin string `json:"zhuyin"`
}
func main() {
	save("幹", "ㄍㄢˋ")
}

func save(word string, zhuyin string) {
	zhi := Word {
		Word: word,
		Zhuyin: zhuyin,
	}
	j, _ := json.Marshal(zhi)
	json := string(j)
	
}
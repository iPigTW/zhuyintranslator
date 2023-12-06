package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Worde struct {
	Word string `json:"word"`
	Zhuyin string `json:"zhuyin"`
}
func main() {
	save("靠", "ㄎㄠˋ")
	fmt.Println(checkWord("嗨"))
}
func checkWord(word string) bool {
	var tmp []Worde
	readfile, _ := os.ReadFile("data.json")
	json.Unmarshal(readfile, &tmp)
	for _, w := range tmp {
		if w.Word == word {
			return true
		}
	}
	return false
}

func save(word string, zhuyin string) {
	var tmp []Worde
	zhi := Worde {
		Word: word,
		Zhuyin: zhuyin,
	}
	readfile, _ := os.ReadFile("data.json")
	json.Unmarshal(readfile, &tmp)
	
	tmp = append(tmp, zhi)
	e,_ := json.Marshal(tmp)
	os.WriteFile("data.json", e, 0644)
}
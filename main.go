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
	var word string
	fmt.Print("Enter a word to check...")
	fmt.Scan(&word)
	check, zhu := check(word)
	if !check {
		var zhuyin string
		fmt.Print("Word not found! Please enter the zhuyin...")
		fmt.Scan(&zhuyin)
		save(word, zhuyin)
	}else {
		fmt.Print("Word found! The zhuyin is "+zhu)
	}
}

func check(word string) (bool, string) {
	var tmp []Worde
	readfile, _ := os.ReadFile("data.json")
	json.Unmarshal(readfile, &tmp)
	for _, w := range tmp {
		if w.Word == word {
			return true, w.Zhuyin
		}
	}
	return false, ""
}

func save(word string, zhuyin string) {
	var tmp []Worde
	zhi := Worde {
		Word: word,
		Zhuyin: zhuyin,
	}
	readfile, _ := os.ReadFile("data.json")
	json.Unmarshal(readfile, &tmp)
	for _, w := range tmp {
		if w == zhi {
			return
		}
	}
	tmp = append(tmp, zhi)
	e,_ := json.Marshal(tmp)
	os.WriteFile("data.json", e, 0644)
}
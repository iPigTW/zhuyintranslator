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
	choice()
}
func choice() {
	var choice int
	fmt.Println(`Choose your input method:
	1.File
	2.Enter it manually
	`)
	fmt.Scanln(&choice)
	switch choice {
	case 2:
		oneWord()
	case 1:
		fileInput()
	}
}
func oneWord() {
	var word string
	fmt.Print("Enter words to check...")
	fmt.Scan(&word)
	for _, w := range word {
		check, zhu := check(string(w))
		if !check && len(string(w)) != 0 {
			var zhuyin string
			fmt.Print("\nWord "+string(w)+" not found! Please enter the zhuyin...")
			fmt.Scan(&zhuyin)
			if len(zhuyin) != 0 {
				save(string(w), zhuyin)
			}
			
		}else {
			fmt.Print(zhu)
		}
	}
	
}
func fileInput() {
	var path string
	fmt.Print("Enter file path:")
	fmt.Scan(&path)
	file, _ := os.ReadFile(path)
	for _, w := range string(file) {
		check, zhuyin := check(string(w))
		if !check && len(string(w)) != 0 {
			var zhu string
			fmt.Print("\nWord "+string(w)+" not found! Please enter the zhuyin...")
			fmt.Scan(&zhu)
			if len(zhu) != 0 {
				save(string(w), zhu)
			}
		}else {
			fmt.Print(zhuyin)
		}
		
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
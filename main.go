package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("layout.html"))

type Number struct {
	Nome  string      `json:"nome"`
	Email string      `json:"email"`
	Cns   interface{} `json:"cns"`
	Cbo   string      `json:"cbo"`
}

func main() {

	dat, erre := os.ReadFile("exp.json")
	fmt.Print(erre)
	data := Number{}
	json.Unmarshal([]byte(dat), &data)

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		fmt.Println("Template erro!", err)
		return
	}
	err = os.WriteFile("index.html", buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("Erro!")
		return
	}
}

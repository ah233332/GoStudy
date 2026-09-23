package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name string `json:"Name"`
	Animal
}

type Animal struct {
	Age   int    `json:"-"`               //忽略字段
	Color string `json:"color,omitempty"` //忽略空值
}

func (c *Cat) Setname(name string) {
	c.Name = name
}

func main() {
	animal := Animal{3, "black"}
	cat := Cat{"大哈", animal}
	cat.Setname("花花")
	byteData, _ := json.Marshal(cat)
	fmt.Println(string(byteData))
	fmt.Println(cat.Name)
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Monster struct {
	Name  string `json:"姓名"`
	Age   int    `json:"年龄"`
	Skill string `json:"技能"`
}

func (p *Monster) Store() {
	data, err := json.Marshal(&p)
	if err != nil {
		return
	}
	jsonPath := "D:/code/go/study/02-stdlib/testing/testcase/js.json"
	file, err := os.OpenFile(jsonPath, os.O_CREATE|os.O_WRONLY, 0666)

	writer := bufio.NewWriter(file)
	writer.WriteString(string(data))
	writer.Flush()
	file.Close()
}

func (p *Monster) ReStore(jsPath string) {
	var per Monster

	file, err := os.OpenFile(jsPath, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	data, err := io.ReadAll(reader)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &per)

	fmt.Println()
	fmt.Printf("json文件反序列化为:%v", per)

}

func main() {
	p := Monster{
		Name:  "李超",
		Age:   88,
		Skill: "打瓦",
	}

	p.Store()
	str := "D:/code/go/study/02-stdlib/testing/testcase/js.json"
	p.ReStore(str)

}

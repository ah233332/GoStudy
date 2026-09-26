package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string `json:"姓名"`
	Age   int    `json:"年龄"`
	Skill string `json:"技能"`
}

// 序列化，保存到文件内。
func (p *Monster) Store(jsonPath string) bool {
	data, err := json.Marshal(&p)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}

	file, err := os.OpenFile(jsonPath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("openfile err =", err)
		return false
	}
	//我用的是缓存写入的方法
	//还可以一次性写入方法 os.WriterFile()
	writer := bufio.NewWriter(file)
	writer.WriteString(string(data))
	writer.Flush()
	file.Close()
	return true
}

// 将文件内容反序列化
func (p *Monster) ReStore(jsPath string) bool {
	//我最开始用的os.openfile + io.readall,适合自定义打开文件(读+写+追加,权限)
	// 如果只读没必要引入两个包包，直接os.readfile()
	//可以用一次性的读取方法 os.ReadFile()
	data, err := os.ReadFile(jsPath) //正好返回data为[]byte类型，与unmarshal适配
	if err != nil {
		fmt.Println("readfile err :", err)
		return false
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		fmt.Println("unmarshal err:", err)
		return false
	}

	fmt.Printf("json文件反序列化为:%v", *p)

	return true
}

func main() {
	p := Monster{
		Name:  "李超",
		Age:   88,
		Skill: "打瓦",
	}

	str := "D:/code/go/study/02-stdlib/testing/testcase/js.json"
	p.Store(str)
	p.ReStore(str)

}

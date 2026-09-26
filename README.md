# Go 学习笔记

2026 年 9 月开始，从零学 Go 的练习仓库。

内容大都是跟着B站课程手搓的，排序分类也很随意。

参考价值不大，就是学习日记，记录一下。

代码注释保留中文原文。


## 运行方式

```bash
cd D:\code\go\study   # 先进入仓库根目录
go run ./01-basics/map   # 再运行某一个练习
```


## 目录结构

```text
study/
├── go.mod                        module study
├── 01-basics/                    基础语法
│   ├── input-output/             输入输出、基本数据类型、fmt 格式化动词
│   ├── if-else/                  if 语句、逻辑运算符
│   ├── switch-case/              switch 语句、fallthrough 穿透
│   ├── for-loop/                 for 循环、range 遍历
│   ├── array-slice/              数组、切片、append 与 make、sort.Slice
│   ├── map/                      map 增删改查、comma-ok 判断
│   ├── func-menu/                函数是一等公民：map[int]func() 做菜单分发
│   ├── func-variadic/            变参函数、range 求和、九九乘法表
│   ├── struct-json/              结构体、匿名嵌入、方法、struct tag 与 json
│   └── init-defer/               init 函数、defer 执行顺序
├── 02-stdlib/                    标准库
│   ├── flag/
│   │   ├── args/                 os.Args 命令行参数
│   │   └── flagdemo/             flag 解析命名命令行参数
│   ├── factoryModel/             工厂模式：私有结构体、构造函数与访问方法
│   │   ├── main/main.go          调用 NewStudent 并输出学生信息
│   │   └── model/model.go        student 封装、NewStudent、GetScore
│   ├── file/
│   │   ├── file-study/
│   │   │   ├── create-read/      文件创建、写入、清空、追加、读取
│   │   │   ├── read-two-ways/    bufio 逐行读 与 os.ReadFile 一次性读
│   │   │   └── copy-file/        读一个文件写到另一个文件、os.Stat
│   │   └── copy-dir/
│   │       ├── srcdata/          拷贝练习使用的源文件
│   │       └── main.go           io.Copy 拷贝任意二进制文件
│   ├── json/
│   │   ├── serial/               结构体、map、切片、基本类型序列化
│   │   └── unmarshal/            结构体、map、切片反序列化
│   └── testing/
│       ├── calTest/              cal.go 与 cal_test.go，基础单元测试
│       └── testcase/             JSON 读写函数及其测试用例
└── 03-practice/                  综合练习
    └── char-count/               统计文件里的英文、数字、空格、其他字符
```



## 常用命令

```bash
go run ./01-basics/map   # 运行某一个练习
go build ./...           # 编译全部，检查有没有报错
go vet ./...             # 静态检查
go test ./...            # 跑测试
gofmt -l .               # 列出没有格式化的文件
gofmt -w .               # 格式化并写回
```


## 待办

接着学习GO语言进阶内容，再学习Gin框架，后续完善打通项目前后端

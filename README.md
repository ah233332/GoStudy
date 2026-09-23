# Go 学习笔记

2026 年 9 月开始，从零学 Go 的练习仓库。

内容大都是跟着B站课程手搓的，排序分类也很随意。

参考价值不大，就是学习日记，记录一下。

代码注释保留中文原文。

## 为什么上传到GitHub

因为我最近学习了Git,试着在vscode里面用git连接到了github

## 运行方式

```bash
cd D:\code\go\study   # 先进入仓库根目录
go run ./01-basics/map   # 再运行某一个练习
```

必须从仓库根目录运行：程序里的文件路径是相对**当前工作目录**的，不是相对 .go 文件所在目录。
想验证的话，在程序里 `fmt.Println(os.Getwd())` 打印一下当前工作目录就知道了。

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
│   ├── args/                     os.Args 命令行参数
│   ├── file-study/
│   │   ├── create-read/          文件创建、写入、清空、追加、读取
│   │   ├── read-two-ways/        bufio 逐行读 与 os.ReadFile 一次性读
│   │   └── copy-file/            读一个文件写到另一个文件、os.Stat
│   └── copy-dir/                 io.Copy 拷贝任意二进制文件
└── 03-practice/                  综合练习
    └── char-count/               统计文件里的英文、数字、空格、其他字符
```

每个练习目录里只有一个 `main.go`。原因是：同一个包里不能有两个 `func main`，
所以每个可执行程序必须独占一个目录。

## 常用命令

```bash
go run ./01-basics/map   # 运行某一个练习
go build ./...           # 编译全部，检查有没有报错
go vet ./...             # 静态检查
go test ./...            # 跑测试（目前还没有测试）
gofmt -l .               # 列出没有格式化的文件
gofmt -w .               # 格式化并写回
```

## 为什么目录名是英文

Go 的导入路径不允许出现非 ASCII 字符。中文目录名会让 `go build ./...`、`go vet ./...`、
`go test ./...` 全部失败，报错形如：

```text
malformed import path "study/01-basics/for循环": invalid char '循'
```

所以目录名使用英文，文件名和注释保持中文。

## 待办

接着学习GO语言进阶内容，再学习Gin框架，后续完善打通项目前后端

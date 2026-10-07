# Go 学习笔记

2026 年 9 月开始，从零学 Go 的练习仓库。

内容大都是跟着B站课程手搓的，排序分类也很随意。

参考价值不大，就是学习日记，记录一下。


## 运行方式

```bash
cd D:\code\go\study   # 先进入仓库根目录
go run ./01-basics/map   # 再运行某一个练习
```

## 目录结构

```text
study/
├── go.mod                        module study；Go 1.27.0
├── README.md
├── 01-basics/                    基础语法
│   ├── input-output/             输入输出、基本数据类型、fmt 格式化动词
│   ├── const/                    const 常量、iota 自增、常量编译期限制
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
│   │   ├── char-count/           统计文件里的英文、数字、空格、其他字符
│   │   ├── copy-dir/
│   │   │   ├── srcdata/          拷贝练习使用的源文件
│   │   │   └── main.go           io.Copy 拷贝任意二进制文件
│   │   └── file-study/
│   │       ├── create-read/      文件创建、写入、清空、追加、读取
│   │       ├── read-two-ways/    bufio 逐行读 与 os.ReadFile 一次性读
│   │       ├── copy-file/        读一个文件写到另一个文件、os.Stat
│   │       ├── 111.txt
│   │       ├── newtest.txt
│   │       └── test.txt
│   ├── json/
│   │   ├── serial/               结构体、map、切片、基本类型序列化
│   │   └── unmarshal/            结构体、map、切片反序列化
│   └── testing/
│       ├── calTest/              cal.go 与 cal_test.go，基础单元测试
│       └── testcase/             JSON 读写函数及其测试用例
└── 03-advance/                   进阶学习
    ├── goroutine-channel/        协程与管道  
    │   ├── 协程管道基本.txt       进程/线程、并发/并行、goroutine 特点、channel 队列与只读/只写、select/recover
    │   ├── 1-channel/             channel 基础：缓冲队列、cap/len、interface{} 与类型断言
    │   ├── 2-channel/             channel 关闭与 range 遍历，避免用变化的 len 遍历导致漏读
    │   ├── channel-case1/         Person 随机数据：生成 10 个结构体写入 channel 后遍历输出
    │   ├── channel-case2/         writeData/readData 两个协程协作，exitChan 通知主协程等待
    │   ├── channel-case3/         8 个 worker 计算前缀和，WaitGroup 等待后关闭 resChan 并顺序输出
    │   ├── channel-case4/         待完成：goroutine + channel + 文件排序，已有需求图 case4.png 和空的 mian.go
    │   ├── channel-case5/
    │   │   ├── way1/              1way-main.go：12 个 worker 判断 1-200000 中的素数，WaitGroup 等待后关闭 primeChan 并排序输出
    │   │   └── way2/              2way-main.go：4 个 worker 判断 0-8000 中的素数，exitChan 收集退出信号后关闭 primeChan
    │   ├── channel-case6/         待完成：1 个协程写 1-2000 到 numChan，8 个协程取出 n 计算 1+...+n 写入 resChan，已有需求图 case6.png 和空的 main.go
    │   ├── select/                select 多路复用：多个 channel 同时等待，default 分支避免 deadlock
    │   ├── recover/               defer + recover 捕获子协程 panic，防止整个程序崩溃
    │   ├── goLock/                20 个协程计算阶乘写入共享 map，sync.Mutex 保护并发写
    │   ├── goTest/                main 与子协程每秒交替打印，观察并发执行和主协程退出
    │   └── runtime/               runtime.NumCPU 与 GOMAXPROCS 查看/限制 CPU 核数
    └── reflect/                   反射
        ├── 1-reflect/             reflect.TypeOf / ValueOf、Kind、Interface() 与类型断言，基本类型和结构体
        ├── 2-reflect/             通过反射修改变量的值：reflect.ValueOf(&x).Elem().SetInt/SetString，必须传指针
        ├── reflect-case1/         小练习：float64 取 Type/Kind/值，Value→interface{}→float64，并用反射改 string 变量
        ├── reflect-case2/         Cal 结构体：反射遍历 Num1/Num2 字段，输入名字后反射调用 GetSub 输出减法
        ├── reflect-caseBest/      反射操作结构体：遍历字段、读 json tag、按名字调用方法、FieldByName 修改字段值
        ├── 反射基本.txt           反射的优缺点、interface{} / reflect.Value / 原类型 三种转换
        └── 反射转换示意图.png
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

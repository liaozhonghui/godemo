# runoob 教程学习

## 语言基本属性

main.go
```go
package main
import (
    "fmt"
)
func main() {
    fmt.Println("Hello, World!");
}
```

go run 命令执行
go 语言的主要特性：
1. 自动垃圾回收
2. 更丰富的内置类型
3. 函数多返回值
4. 错误处理
5. 匿名函数和闭包
6. 类型和接口
7. 并发编程
8. 反射
9. 语言交互性

## go 语言结构

go程序的基本组成部分
- 包声明package
- 引入包import
- 函数func 
- 语句&表达式
- 注释 // /***/

go 程序基础语法
关键字， 标识符， 常量， 变量， 字符串， 符号
一个标识符由字母，数字，下划线表示，并且首字符为字母或者下划线而不能是数字.

go语言数据类型
- 布尔型
- 数字类型
- 字符串类型
- 派生类型
指针类型，数组类型，结构体类型，channel类型，函数类型，切片类型，接口类型，map类型
- 其他数字类型
byte: 类似uint8, rune: 类似int32, int, uintptr

go语言变量
变量定义`var identifier type`
声明多个变量：`var identifier1, identifier2 type`
如果没有初始化， 则默认为类型的零值, 数字类型为0, 字符串类型为"", 布尔类型为false, 其他类型为nil.
_用于被抛弃值
变量交换小技巧 `a, b = b, a`
使用`:=` 符号用来声明一个变量并进行赋值.

小知识：对于值类型的数据，可以通过哟&i来进行取地址，值类型变量的值存储在栈中，内存地址会根据机器的不同而有所不同，因为每台机器可能有不容的存储器布局，并且位置分配可以能不同.


variable.go
```go
package runoob
import "fmt"
func main(){
    _, numb, strs := numbers()
    fmt.Println(numb, strs)
}
func numbers() (int, int, string) {
    a, b, c := 1, 2, "str"
    reutrn a, b, c
}
```

go语言常量
常量定义`const identifier [type] = value`
常量可以使用len(), cap(), unsafe.Sizeof()函数计算表达式的值。
iota特殊常量，
iota在const关键词出现的时候被重置为0(const 内部的第一行代码之前)，const每新增一行代码声明可以使iota计数增加一次 

iota.go
```go
package runoob
import "fmt"
func main() {
    const (
        a = iota // 0
        b // 1
        c // 2
        d = "ha" // 独立值，iota += 1
        e // "ha" iota += 1
        f = 100 // 独立值 iota += 1
        g // 100 iota += 1
        h = iota // 7
        i // 8
    )
    fmt.Println(a, b, c, d, e, f, g, h, i)
}
```

go 语言运算符号
&& || !
位运算符： & | ^(异或) << >>
指针符号： &返回变量存储地址， *指针变量
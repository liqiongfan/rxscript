# RXScript 脚本语言语法规范指南

欢迎使用 `RXScript`。

## Hello World 入门

```RXScript
main() {
  println("Hello world")
}
```

## 安装 RXScript 

- 已安装 `Go` 1.18以及以上版本的可以运行如下命令直接安装:
```shell
go install github.com/liqiongfan/rxscript@latest
```

- 没有安装 `Go` 1.18以及以上版本的直接下载二进制包或者编译安装

```shell
go install .
```

- [rxscript-linux-amd64](rxscript-linux-amd64) 是linux平台的二进制可执行文件
- [rxscript-mac-arm64](rxscript-mac-arm64) 是macOS的arm64平台二进制可执行文件
- [rxscript-mac-amd64](rxscript-mac-amd64) 是 macOS的amd64平台可执行文件
- [rxscript.exe](rxscript.exe) 是Windows平台的二进制可执行文件

- 验证是否安装成功

```shell
rxscript -v
```

- 执行 `RXScript` 脚本

```shell
rxscript xxx.rxt
```

如果文件以 `.rxt` 结尾，那么可以省略后缀：

```shell
rxscript xxx
```


## 语法指南

### 包导入

- 系统包导入: `import 包名` 

```RXScript
import std
```

系统包不需要 `双引号`, 添加了双引号的是自定义包。

- 自定义包导入: `import "包名"`

```RXScript
import "std"
```

目前自带的 `std` 标准库包含了字符串 `String`, `Atomic`, `WaitGroup`, `Http`, `Timer`, `Ticker`，后续类正在迭代中.

### 类型系统

|    数据类型    |                 概念说明                  | 代码样例                             |
|:----------:|:-------------------------------------:|----------------------------------|
|  Integer   | 整数数值, 支持二进制、八进制、十六进制、十进制，支持自增/自减及算术运算 | 	b = 10  b = 044 b=0x1f b=0b1010 |
|   Float    |      (浮点型)	带有小数点的数值。支持写成 .05 形式       | 999.9, .05, 2.58                 |
|   String   |            (字符串)	由双引号包裹的文本            | "hello"                          |
|  Boolean   |        (布尔型)	条件判断的基本类型，只有两个值。         | true, false                      |
|    Null    |        (空值)	表示变量没有分配任何有效值或对象为空        | null                             |
|   Array    |           方括号包裹的用逗号分割的数据集合            | [1,2,3,true,"",[]]               |
| Dictionary |            大括号包裹的用逗号分割的键值对            | {a:1, b:true}                    |

`RXScript` 是弱类型的语言，定义变量的时候直接定义即可，不需要声明类型：

```RXScript
a = 1 // 十进制
a = 0xabcdef // 十六进制
a = 0122 // 八进制
a = 0b1010 // 二进制
a = 1e9 // 科学计数法
b = true
c = {}
```

### 优先级与运算符支持


|     优先级     | 运算符                                     | 描述           | 结合性   |
|:-----------:|:----------------------------------------|:-------------|:------|
| **1 (最高)**  | `()`                                    | 圆括号（更改常规优先级） | 从内到外  |
|             | `[]`                                    | 数组/切片下标访问    | 左结合   |
|             | `++`, `--`                              | 自增、自减后缀      | 后缀单目  |
|    **2**    | `-`                                     | 单目负号（一元负数）   | 右结合   |
|    **3**    | `*`, `/`, `%`                           | 乘法、除法、取模（取余） | 左结合   |
|    **4**    | `+`, `-`                                | 加法、减法        | 左结合   |
|    **5**    | `<<`, `>>`                              | 左移、右移        | 左结合   |
|    **6**    | `&`, `\|`, `^`                          | 按位与、按位或、按位异或 | 左结合   |
|    **7**    | `>`, `>=`, `<`, `<=`, `!=`, `==`, `===` | 关系与相等性比较     | 左结合   |
|    **8**    | `&&`                                    | 逻辑与          | 右结合   |
|    **9**    | `\|\|`                                  | 逻辑或          | 右结合   |
|   **10**    | `? :`                                   | 三元条件运算符      | 右结合   |
| **11 (最低)** | `+=`, `-=`                              | 赋值运算符        | 视作原子值 |


`RXScript` 中数字类型的字符串数据，会被采用高精度运算，如果需要设置精度，可以使用

```shell
rxscript -p 32 xxx.rxt
```

`-p`: 表示字符串数字的运算精度，默认为小数点后32位


### 类与函数

- 类与面向对象 (OOP)

i. 类的定义与实例化

在 `RXScript` 中，使用 `class` 关键字来声明一个类。类是对象的蓝图，包含属性（成员变量）和方法（成员函数）。
- 声明类：可以使用 `public` 或 `private` 关键字来修饰属性和方法。
- 实例化：直接像调用函数一样使用 类名(参数) 即可创建对象，无需 `new` 关键字。
  代码样例：
```RXScript
class B {
  public name = "1111"  // 类属性
  public getName() {    // 类方法
    return this.name
  }
}

// 实例化对象
b = B()
```

ii. 构造函数

构造函数用于在创建对象时进行初始化操作。RXScript 的构造函数有以下特点：
- 命名规则：构造函数的名称必须与类名完全相同。
- 如果需要调用父类的构造函数，直接使用
```RXScript 
this.父类名称()
```
- 参数传递：可以接受外部传入的参数，用于初始化内部属性。

代码样例：
```RXScript
class Hello : Base, Animal {
  public Hello(id, name) {
    this.Base() // 调用Base的构造函数
    this.id = id  // 将参数赋值给对象的属性
    this.name = name
    this.address = "my address."
  }
}

// 实例化时传入参数，触发构造函数
a = Hello(99, "Rxe")
```

iii. 访问控制与封装

封装将数据（属性）和行为（方法）绑定在类内部，并通过访问修饰符控制外部的访问权限：
- `public`（共有）：可以在类的外部被直接访问和修改（例如：a.doMoney(999.9)）。
- `private`（私有）：只能在类的内部（方法中）被访问，外部无法直接调用，从而保护内部数据的安全性。
  代码样例：
```RXScript
class Base {
  private address = "do.1.2" // 私有属性，外部无法通过 base.address 访问
  public Base() { ... }
}

class Hello : Base, Animal {
  private printlnGetAddress(p) { // 私有方法，只能在 Hello 类内部的方法中被调用
    x = p(" 100")
    return x
  }
}
```

iiii. `this` 关键字

在类的方法内部，this 关键字代表当前调用该方法的对象实例。
- 用于在方法内访问或修改当前对象的属性：`this.id = id`。
- 用于在方法内调用当前对象的其他方法：`this.getMoney()`。
- `RXScript` 支持动态属性扩展，即使类中没有预先声明，也可以通过 `this.name2 = 9` 在运行时为当前实例动态添加新属性。

iiiii. 多重继承

这是 `RXScript` 中非常核心且高级的 OOP 特性。一个子类可以通过冒号 : 同时继承多个父类，从而获得所有父类的属性和方法。
代码样例：
```RXScript
// Hello 类同时继承了 Base 类和 Animal 类
class Hello : Base, Animal
{
  // Hello 实例将同时拥有 Base 和 Animal 的公有属性 address
}
```

- 函数

在 `RXScript` 中，全局函数是独立于类之外、直接定义在脚本根级别的执行单元。

i. 函数的语法格式

`RXScript` 的函数定义非常简洁，采用弱类型设计，不需要声明参数类型和返回值类型。

定义语法

```RXScript
函数名(参数1, 参数2, ...) {
// 执行代码块
return 返回值
}
```

- 参数列表：参数仅需写出变量名，多个参数用逗号 , 分隔。如果没有参数，保持括号 () 为空即可。
- 返回值：使用 return 关键字返回结果。支持弱类型返回，一个函数可以根据逻辑返回任意类型的数据。如果不需要返回值，可以省略 return。

ii. 调用语法

```RXScript
变量 = 函数名(实参1, 实参2, ...)
```

iii. 函数的用法特性

基础调用与传参

函数通过传入指定的参数来执行内部逻辑。例如：
```RXScript
// 定义：接收一个参数 ver
getAddress(ver) {
  return "hello - getAddress " + ver
}

// 调用：传入字符串 "v0.1.0"，接收返回的字符串结果
result = getAddress("v0.1.0")
println(result)
```


iiii. 函数定义时的 ...：可变参数

当 ... 紧跟在函数或方法的最后一个形参名称后面时，它表示该函数可以接收任意数量的实参。

语法格式

```RXScript
class 类名 {
  public 方法名(形参名 ...) {
    // 此时 形参名 在方法内部会自动转为一个 Array (数组)
  }
}
```

    注意点：... 必须放在形参列表的最后一位。

用法特性与代码实例

```RXScript
class A {
  // 接收任意多个参数，并将它们打包进名为 a 的数组中
  public say(a ...) {
    println(a) // 打印出来的是一个完整的数组
  }
}

```

函数调用时的 ...：打散/展开数组

当 ... 紧跟在一个数组变量的后面并作为实参传入函数时，它表示将这个数组拆解、打散为一个个独立的个体，分别传给函数。

语法格式

```RXScript
数组变量 = [值1, 值2, 值3]
函数名(数组变量...) // 等价于：函数名(值1, 值2, 值3)
```

用法特性与代码实例

在 main() 函数中，您展示了将可变参数与数组展开配合使用的绝佳样例：

```RXScript
main() {
  v = [1, 2, 3, 4] // v 是一个数组
  
  a = A()
  
  // 注意这里的 v...
  a.say("RXScript", "do", v..., {ab: 20})
}
```

⚙️ 运行时深度解析：

如果没有 ...，传入的第三个参数将是整个数组 [1, 2, 3, 4]。
但是因为使用了 v...，该数组在传入前被平铺展开了。因此，a.say 实际上接收到了 6 个独立的参数：
1. "RXScript"
2. "do"
3. 1（来自 v[0]）
4. 2（来自 v[1]）
5. 3（来自 v[2]）
6. 4（来自 v[3]）
7. {ab: 20}
   最终，这 6 个参数由于 say(a ...) 的可变参数定义，重新在函数内部被打包打印。

总结：核心区别

| 出现位置     | 语法学术语          | 核心作用                | 视觉记忆法      |
|----------|----------------|---------------------|------------|
| 函数/方法定义时 |                |                     |            |
| (a ...)  | 可变参数 (Rest)    | 把零散的多个参数打包成一个数组。    | 📥 收拢 / 聚合 |
| 函数/方法调用时 |                |                     |            |
| (v...)   | 展开运算符 (Spread) | 把一个现成的数组拆开成零散的多个参数。 | 📤 铺开 / 释放 |


### 控制流与条件分支

- 条件分支

`RXScript` 支持标准的 `if`、`else if` 和 `else` 结构，用于根据布尔表达式的真假来决定执行哪个代码块。

语法格式

```RXScript
if 条件表达式1 {
  // 条件1成立时执行
} else if 条件表达式2 {
  // 条件2成立时执行
} else {
  // 所有条件都不成立时执行
}
```

用法特性与代码实例

• 多分支判断：
```RXScript
if this.id < 10 {
  println(this.name + " script is very excellent")
} else if this.name2 {
  // name2 为真时执行（RXScript 支持隐式类型转换，非零值/非空通常视为真）
} else {
  println("go")
}
```

• 三元运算符（简写条件分支）：
`RXScript` 支持 (条件) ? 表达式1 : 表达式2 语法，用于简化双分支的赋值或输出。

```RXScript
println((this.name2 && this.id > 100) ? this.z : this.name)
```

- 循环控制 (Loops)

代码中展示了两种循环用法：无限循环（通过 break 退出）和 遍历循环（配合 range 和 series）。

无限循环与中断 (for + break)

```RXScript
for {
  if num >= 10 {
    break // 当满足条件时，使用 break 关键字强制跳出循环
  }
  num++
}
```

遍历循环 (for...range)

用于遍历数组、字典或序列。语法结构为 for 键, 值 = range 集合。
• 遍历数组：

```RXScript
for name, value = range this.xz {
  println("name: ", name, " value: ", value) // name 为索引，value 为元素值
}
```

• 遍历等差数列序列（配合 `series` 函数）：
```RXScript
// 从 0 开始，到 10 结束，步长为 2 (生成的序列为 0, 2, 4, 6, 8, 10)
for idx, value = range series(0, 10, 2) {
  println("idx=", idx, " value=", value)
}
```

- 特殊控制流

`RXScript` 引入了类似于 `Go` 语言的现代高级控制流关键字。

延迟执行 (defer)

- 语法与用法：defer { 代码块 }
- 特性：被 defer 包裹的代码块不会立即执行，而是会推迟到当前函数或当前作用域即将退出（返回）之前才执行。如果有多个 defer，它们会按照“后进先出”（LIFO）的顺序逆序执行。
- 代码实例：
```RXScript
go {
  defer { println("first defer") } // 离开当前 go 块前执行
  defer { println("second defer, but run first") }    // 先于上面的 defer 执行
  println("goroutine1")
}
```

### 高并发与线程安全

高并发核心机制：协程与通道


并发关键字 `go`

- 语法：`go { /* 代码块 */ }`
- 用法与原理：go 关键字会启动一个轻量级协程独立运行，它不会阻塞当前主线程。
- 代码实例：
```RXScript
go {
  // 协程 1：后台执行通道接收、原子读取等
  d = <-c
  println("goroutine1, b=", b)
}

go {
  // 协程 2：后台并发执行原子写入
  storeAtomicInt(atomic, readAtomicInt(atomic) + 10)
}
```

并发通信 `channel` (通道)

在多协程并发中，`RXScript` 遵循“不要通过共享内存来通信，而要通过通信来共享内存”的原则。
- 语法与用法：
  - `c = channel(capacity)`：创建一个具备指定容量的并发安全管道。
  - `c <- value`：向通道发送数据（若通道满则阻塞）。
  - `v = <- c`：从通道接收数据（若通道空则阻塞）。
  - 代码实例：
```RXScript
c = channel(2) // 创建一个容量为 2 的安全缓冲区
c <- 10        // 主协程发送数据

go {
  d = <-c      // 子协程安全地接收数据 10
}
```

线程安全与并发同步控制

当多个协程并发访问共享资源时，必须保证线程安全。代码中展示了两种典型的同步与安全控制手段：

计数器同步组 `wait` 函数 或者 `std.WaitGroup` 类

为了防止主线程（main 函数）在子协程还没执行完就提前退出，必须使用 wait 家族函数进行阻塞同步。
核心函数：
- `wg = wait()`：初始化一个同步组。
- `waitAdd(wg, delta)`：增加需要等待的协程数量。
- `waitDone(wg)`：协程执行完毕时，减少等待计数（通常结合 defer 确保必定触发）。
- `waitOk(wg)`：主线程阻塞点，直到计数归零才放行。

代码控制流模型：
```RXScript
wg = wait()

waitAdd(wg, 1) // 计数 +1
go {
  defer { waitDone(wg) } // 无论协程内部如何，退出时计数 -1
  // 协程业务...
}

waitOk(wg) // 阻塞，直到子协程全部执行完毕
```

或者使用类版本
```RXScript

import "std"

wg = std.WaitGroup()
wg.add(1)
go {
  defer { wg.done() }
  // TODO...
  ...
}
wg.wait()

```

无锁线程安全 `atomicInt` 或者 `std.Atomic` 类 (原子操作)

当多个协程需要同时读写同一个整数变量（如并发计数器）时，传统的 `b++` 不是线程安全的。`RXScript` 提供了原子操作来保证硬件级别的线程安全。
核心函数：
- atomic = atomicInt(initialValue)：创建并初始化一个原子整型变量。
- readAtomicInt(atomic)：线程安全地读取当前值。
- storeAtomicInt(atomic, newValue)：线程安全地写入新值。
  代码实例：
```RXScript
atomic = atomicInt(10) // 初始化为 10

go {
  // 协程 1 安全读取并修改
  storeAtomicInt(atomic, 110)
}

go {
  // 协程 2 安全累加：先原子读，再加10，最后原子写
  storeAtomicInt(atomic, readAtomicInt(atomic) + 10)
}
```

或者类版本:

```RXScript
import "std"

atomic = std.Atomic(10) // 初始化为 10

go {
  // 协程 1 安全读取并修改
  atomic.store(10)
}

go {
  // 协程 2 安全累加：先原子读，再加10，最后原子写
  atomic.store(atomic.load() + 10)
}
```

并发与线程安全对比摘要

| 特性 / 工具   | 解决的问题                 | 语法/函数               | 安全级别             |
|-----------|-----------------------|---------------------|------------------|
| go        | 开启异步并行任务	`go { ... }` | 依赖内部逻辑              |                  |
| channel   | 协程间安全传递数据/传递对象        | channel(n)、<-       | 绝对安全（内建锁保护）      |
| wait 系列   | 协调多协程的生命周期，防止主线程过早退出  | wait Ok/Add/Done    | 流程同步安全           |
| atomicInt | 多协程并发读写同一数值时的数据竞争问题   | read/storeAtomicInt | 绝对安全（CPU硬件级原子保护） |


### 系统内置函数与 API 参考手册

本语言（RxEngine）提供了一系列原生的全局内置函数，涵盖了 I/O 打印、文件系统操作、时间时区处理、数值序列以及动态类型检查。

基础与 I/O 函数


println(arg1, arg2, ...)

• 说明：向标准输出（控制台）打印一行或多行数据。
• 参数：接受任意数量、任意类型的参数，参数之间用逗号 , 分隔。

文件系统操作 API

listDir(path)

• 说明：列出指定目录下的所有文件和子目录名称。
• 参数：path (String) - 目标目录路径。
• 返回值：一个包含文件名字符串的数组 (Array)。

readFile(filePath)

• 说明：读取并返回指定文件的全部文本内容。
• 返回值：文件内容的字符串 (String)。

putFile(filePath, content, permMode, isAppend)

• 说明：向指定文件写入数据。如果文件不存在会自动创建。
• 参数：permMode 通常为八进制格式权限位字符串（如 "0644"）；isAppend 为 true 时追加写入，false 时覆盖原文件。

delFile(filePath) / rename(oldPath, newPath)

• 说明：分别用于文件物理删除以及重命名/移动路径。

时间与时区处理 API

getTimeResource([timeStr])

• 说明：获取高精度时间对象。不传参默认返回当前时间，或解析包含 "YYYY-MM-DD HH:mm:ss" 格式的字符串。
• 返回值：时间对象实例，支持链式调用 .In(loc) 和 .Format(layout)（采用经典的 Go 语言占位符 "2006-01-02 15:04:05"）。

getLocation(timezoneStr)

• 说明：加载并返回指定时区的位置资源对象（如 "Asia/Shanghai"）。

动态类型检查函数（Type Assertion）

系统提供了一组以 is 开头的全局断言函数，用于在运行时动态校验变量或对象的具体类型。所有断言函数均返回布尔值（true 或 false）。

函数名	说明	示例
+ isInt(val)	检查变量是否为整型 (Int)	isInt(10) (返回 true)
+ isFloat(val)	检查变量是否为浮点型 (Double/Float)	isFloat(9.9) (返回 true)
+ isNull(val)	检查变量是否为空值 (Null)	isNull(null) (返回 true)
+ isBool(val)	检查变量是否为布尔型 (Bool)	isBool(true) (返回 true)
+ isString(val)	检查变量是否为字符串 (String)	isString("rx") (返回 true)
+ isArray(val)	检查变量是否为动态数组 (Array)	isArray(arr) (返回 true)
+ isDictionary(val)	检查变量是否为字典/映射 (Dictionary)	isDictionary(kv) (返回 true)
+ isFunc(val)	检查变量是否为函数或类方法 (Function)	isFunc(a.sayHello) (返回 true)

join —— 字符串拼接

join 函数用于将一个数组（Array）中的所有元素，通过指定的分隔符（Delimiter）连接成一个单一的字符串（String）。

语法格式
```RXScript
新字符串 = join(目标数组, 分隔符字符串)
```
- 参数 1：目标数组。可以包含数字、字符串等，RXScript 会在拼接时进行隐式类型转换。
- 参数 2：分隔符字符串。用于插入到每个数组元素之间的文本（如 ,, -, | 等）。

explode —— 字符串分割

explode 函数执行相反的操作。它根据指定的分隔符（Delimiter），将一个完整的字符串（String）切分成多个碎片，并返回一个包含这些碎片的数组（Array）。

语法格式
```RXScript
结果数组 = explode(原始字符串, 分隔符字符串)
```
- 参数 1：原始字符串。需要被裁剪拆分的文本。
- 参数 2：分隔符字符串。作为裁剪基准的切片标识符。

延迟执行函数 sleep

sleep 函数用于挂起当前协程/线程，函数会根据传入参数的 数据类型 自动解析对应的时间单位：

参数类型	默认单位	说明	示例
string	显式指定	必须携带单位后缀，支持: ns, us, ms, s, m, h	sleep("1s") (等待 1 秒)
int	秒 (s)	传入整数时，默认按秒计算	sleep(3) (等待 3 秒)
double	毫秒 (ms)	传入浮点数时，默认按毫秒计算	sleep(1.5) (等待 1.5 毫秒)

### 程序入口

脚本执行时，默认寻找全局的 `main()` 函数作为程序的启动入口点。
```RXScript
main() {
  // TODO...
  return 0
}
```







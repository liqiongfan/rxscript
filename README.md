# RXScript 脚本语言语法规范指南

欢迎使用 `RXScript`。

## Hello World 入门

```RXScript
main() {
  println("Hello world")
}
```

## 为什么取名 RXScript ?

取自女儿名字拼音的首字母 **R(若) X(曦)** 


## Go 项目接入

`RXScript` 是一门建立在 `Go` 应用生态上的动态脚本执行语言，弥补了 `Go` 语言编译后无法动态更新与加载的问题。

在 `Go` 中接入 `RXScript` 非常便捷, 并且提供了参数控制脚本中可用的最大协程数量，如下是`10`个协程:
```Go
ast.RunScript(scriptPath, 0, 10/*限制RXScript可用的协程数量*/)
```

### 目录

1. 模块导入
2. 数据类型与字面量
3. 类与面向对象 (OOP)
4. 函数与方法
5. 控制流与条件分支
6. 运算符与优先级
7. 高并发与线程安全
8. 内置标准库与全局函数 
9. 程序入口

#### 

1.模块导入

在 `RXScript` 中，`import` 关键字用于引入外部模块或标准库，以便在当前脚本中使用其提供的功能、类或函数。

1.语法格式

  `RXScript` 支持两种主要的导入方式：默认导入 和 别名导入。
  `RXScript` 自带 `std` 模块。
  目前自带的 `std` 标准库包含了字符串 `String`, `Atomic`, `WaitGroup`, `Http` 四个类，后续类正在迭代中.

  [ **注意** ] 系统模块导入 不需要 双引号 如：

```RXScript
import std
```

如下代码导入的是用户自定义 `std` 模块:

```RXScript
import "std"
```
  
  1.1 默认导入

  在 `RXScript` 中，模块路径是相对于 `main` 函数的位置来确定
  直接通过模块路径进行导入。导入后，需要使用模块的原始名称作为命名空间来调用其内部的成员。
  ```RXScript
  import "模块名称"
  ```
  
  1.2 别名导入（使用 as 关键字）
  
  当模块路径较长，或者为了避免命名冲突时，可以使用 as 关键字为导入的模块指定一个更简短或不同的命名空间别名。
  ```RXScript
  import "模块名称" as 别名
  ```

2.数据类型与字面量

本语言为动态类型，变量无需声明类型，直接赋值即可使用。支持以下基础和复合类型：
  
  2.1 基础数据类型 (Primitive Types)

| 数据类型    | 概念说明                       | 代码样例                      |
|---------|----------------------------|---------------------------|
| Integer | 整数数值，支持自增/自减及算术运算          | 	b = 10, this.id = 0, b++ |
| Float   | (浮点型)	带有小数点的数值。支持写成 .05 形式 | 999.9, .05, 2.58          |
| String  | (字符串)	由双引号包裹的文本            | "hello"                   |
| Boolean | (布尔型)	条件判断的基本类型，只有两个值。     | true, false               |
| Null    | (空值)	表示变量没有分配任何有效值或对象为空    | null                      |
  
  2.2 复合/高级数据类型 (Composite & Advanced Types)
  
  
  🟩 Array (数组 / 列表)
  
+ 说明：使用方括号 [] 定义的有序集合。RXScript 的数组是动态且弱类型的，同一个数组内可以同时存放整型、浮点型、Null、字符串、甚至嵌套另一个数组。
- 代码样例：
```RXScript
  xz = [1, 200, null, this.name, "rx", ["go", "rxengine"]]
```
  
  🟩 Dictionary (字典 / 键值对)
  
- 说明：使用大括号 {} 定义的键值对集合。可以通过点语法 .属性名 动态添加或修改键值。
- 代码样例：
```RXScript
  kv = {}       // 初始化空字典, 也允许初始化有值字典
  kv.ab = 99.99 // 赋值
```
  
  🟩 Class & Instance (类与实例对象)
  
- 说明：通过 class 关键字定义的结构，支持多重继承（如 Hello : Base, Animal）。通过 类名() 可以实例化为对象。
- 代码样例：
```RXScript
  a = Hello(99, "Rxe") // 'a' 是 Hello 类的实例对象
```
  
  2.3 并发与同步数据类型 (Concurrency Types)

  🟩 Channel (通道 / 管道)
  
+ 说明：用于在不同的 Goroutine (并发协程) 之间传递数据的安全管道。使用 channel(容量) 初始化，通过 <- 运算符进行发送和接收。
+ [ **注意** ] 对于容量为0通道的,没有写入者之前直接读取会死锁，或者没有读取者之前写入也会死锁。
+ 代码样例：
```RXScript
  c = channel(2) // 创建容量为 2 的通道
  c <- 10        // 发送数据
  v = <- c       // 接收数据
```
  
  🟩 WaitGroup (等待组)
  
- 说明：用于阻塞等待一组并发协程完成的同步原语。在内建函数中通过 wait() 创建，使用 waitAdd、waitDone、waitOk 进行控制。
- 代码样例：
```RXScript
  wg = wait()       // 创建等待组
  waitAdd(wg, 1)    // 计数加 1
  waitDone(wg)      // 计数减 1 (通常在 defer 中)
  waitOk(wg)        // 阻塞等待直到计数清零
```
  
  🟩 AtomicInt (原子整型)
  
- 说明：支持线程安全（协程安全）原子操作的整数类型。不能直接修改，必须通过 readAtomicInt 和 storeAtomicInt 函数进行读写。
- 代码样例：
```RXScript
  atomic = atomicInt(10)          // 创建并初始化为 10
  storeAtomicInt(atomic, 110)     // 原子写入
  readAtomicInt(atomic)           // 原子读取
```
  
  2.3 内建特种类型 (Special Built-in Types)
  
- TimeResource (时间资源)：通过 getTimeResource() 获取的时间对象，拥有 .Format()、.In() 等内建方法。
- Location (时区对象)：通过 getLocation("America/Los_Angeles") 获取的时区标识。
- Series (序列/迭代器)：通过 series(start, end, step) 生成的等差数列集合，专门用于 for...range 循环。
- Function (函数类型)：函数在 RXScript 中是一等公民，可以作为参数传递。可以通过 isFunc(funcX) 来校验。


3.类与面向对象 (OOP)

3.1 类的定义与实例化

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

3.2 构造函数

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

3.3 访问控制与封装

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

3.4 `this` 关键字

在类的方法内部，this 关键字代表当前调用该方法的对象实例。
- 用于在方法内访问或修改当前对象的属性：`this.id = id`。
- 用于在方法内调用当前对象的其他方法：`this.getMoney()`。
- `RXScript` 支持动态属性扩展，即使类中没有预先声明，也可以通过 `this.name2 = 9` 在运行时为当前实例动态添加新属性。

3.5 多重继承

这是 `RXScript` 中非常核心且高级的 OOP 特性。一个子类可以通过冒号 : 同时继承多个父类，从而获得所有父类的属性和方法。
代码样例：
```RXScript
// Hello 类同时继承了 Base 类和 Animal 类
class Hello : Base, Animal
{
  // Hello 实例将同时拥有 Base 和 Animal 的公有属性 address
}
```

4.函数

在 `RXScript` 中，全局函数是独立于类之外、直接定义在脚本根级别的执行单元。

4.1 函数的语法格式

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

4.2 调用语法

```RXScript
变量 = 函数名(实参1, 实参2, ...)
```

4.2 函数的用法特性

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

  作为“一等公民”传递（函数指针/闭包）

在 `RXScript` 中，函数是一等公民（First-Class Function）。这意味着函数名本身可以作为一个变量，赋值给其他变量，或者作为参数传递给另一个函数。
• 将函数作为参数传递：
```RXScript
// 类中的成员方法，接收一个参数 p（期望它是一个函数）
private printlnGetAddress(p) {
  x = p(" 100") // 在内部像调用普通函数一样调用参数 p
  return x
}

// 调用时，直接把全局函数名 getAddress 作为参数传进去（注意：不要加括号）
println(this.printlnGetAddress(getAddress))
```

4.3 函数定义时的 ...：可变参数

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


5.控制流与条件分支

5.1 条件分支

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

5.2 循环控制 (Loops)

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

  特殊控制流

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

6.运算符与优先级

6.1 运算符分类与语法

  算术运算符

用于执行基础数学运算。
- 基础运算符：`+`（加/字符串拼接）、-（减）、*（乘）、/（除）。
- 特殊语法：代码中展示了字符串减法（隐式类型转换后计算数值）`"50" - "30"`。
- 自增/自减：`b++` (自增1)、b-- (自减1)。注意：在 `RXScript` 中，它们通常作为独立语句使用。

  赋值运算符

用于给变量赋值，支持复合赋值。
- 基础赋值：`=`。例如：`b = 10`。
- 复合算术赋值：`+=、-=`。例如：`this.name2 = this.id-=5/5`。

  比较运算符

用于比较两个值，结果返回布尔值（true 或 false）。
- 符号：`<（小于）、>（大于）、<=（小于等于）、>=（大于等于）、==（等于）`。

  逻辑运算符

用于连接多个布尔条件。
- 逻辑与 (`&&`)：两边条件同时为真时结果为真。
- 逻辑或 (`||`)（样例未直接出现但支持）：只要有一边为真结果就为真。

  通道操作符

专门用于并发协程间 Channel 通信的独有符号 <-。
- 发送数据：`c <- 10` （将 10 发送到通道 c 中）。
- 接收数据：`v = <- c` （从通道 c 中读取数据并赋值给 v）。

  三元条件运算符

- 语法：`条件 ? 表达式1 : 表达式2`
- 用法：如果条件为真，执行并返回表达式1的值；否则执行并返回表达式2的值。

二、 运算符优先级

当一个表达式中包含多个运算符时，RXScript 会按照优先级从高到低的顺序进行计算。以下根据代码样例的运行逻辑，整理出优先级从高到低的对应关系：

| 优先级 | 运算符类型   | 运算符符号                       | 结合性   | 代码样例与解析                                                                                      |
|-----|---------|-----------------------------|-------|----------------------------------------------------------------------------------------------|
| 1   | 成员访问与调用 | . (点语法)、() (函数调用)、[] (数组索引) | 从左到右  | this.xz[5][0] 先访问 xz 属性，再取索引 5，最后取二级索引 0。                                                    |
| 2   | 单目与通道接收 | ++、--、<- (接收)               | 从右到左  | v = <- c  先从通道 c 中取出值，再执行赋值。                                                                 |
| 3   | 乘除运算    | *、/                         | 从左到右  | 5/5   在复合运算中优先被计算。                                                                           |
| 4   | 加减运算    | +、-                         | 从左到右  | this.getMoney()+2.58 先调用方法，再做加法。                                                             |
| 5   | 比较运算    | <, >, <=, >=, ==            | 从左到右  | this.id > 100    优先于逻辑与 && 执行。                                                               |
| 6   | 逻辑与	    | &&                          | 从左到右  | this.name2 && this.id > 100   先计算右侧比较，再进行逻辑与。                                                |
| 7   | 三元条件    | ? :                         | 从右到左  | (条件) ? this.z : this.name    在逻辑运算符计算完毕后进行分支选择。                                              |
| 8   | 赋值运算	   | =, -=, +=                   | 从右到左	 | this.name2 = this.id-=5/5   右结合典范：先算 5/5 得到 1；然后执行 this.id -= 1；最后将整个减法表达式的结果赋值给 this.name2。 |

**注意：**
当字符串内容为数字时, 内部自动采用大整数计算执行 +, -, *, / 等算术：
```RXScript
println("50" - "30") // 输出 20
```

7.高并发与线程安全

  高并发核心机制：协程与通道


7.1 并发关键字 `go`

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

7.2 并发通信 `channel` (通道)

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


8.系统内置函数与 API 参考手册

本语言（RxEngine）提供了一系列原生的全局内置函数，涵盖了 I/O 打印、文件系统操作、时间时区处理、数值序列以及动态类型检查。

8.1 基础与 I/O 函数


println(arg1, arg2, ...)

• 说明：向标准输出（控制台）打印一行或多行数据。
• 参数：接受任意数量、任意类型的参数，参数之间用逗号 , 分隔。

8.2 文件系统操作 API

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

8.3 时间与时区处理 API

getTimeResource([timeStr])

• 说明：获取高精度时间对象。不传参默认返回当前时间，或解析包含 "YYYY-MM-DD HH:mm:ss" 格式的字符串。
• 返回值：时间对象实例，支持链式调用 .In(loc) 和 .Format(layout)（采用经典的 Go 语言占位符 "2006-01-02 15:04:05"）。

getLocation(timezoneStr)

• 说明：加载并返回指定时区的位置资源对象（如 "Asia/Shanghai"）。

8.4 动态类型检查函数（Type Assertion）

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

8.5 join —— 字符串拼接

join 函数用于将一个数组（Array）中的所有元素，通过指定的分隔符（Delimiter）连接成一个单一的字符串（String）。

  语法格式
```RXScript
新字符串 = join(目标数组, 分隔符字符串)
```
- 参数 1：目标数组。可以包含数字、字符串等，RXScript 会在拼接时进行隐式类型转换。
- 参数 2：分隔符字符串。用于插入到每个数组元素之间的文本（如 ,, -, | 等）。

8.6 explode —— 字符串分割

explode 函数执行相反的操作。它根据指定的分隔符（Delimiter），将一个完整的字符串（String）切分成多个碎片，并返回一个包含这些碎片的数组（Array）。

  语法格式
```RXScript
结果数组 = explode(原始字符串, 分隔符字符串)
```
- 参数 1：原始字符串。需要被裁剪拆分的文本。 
- 参数 2：分隔符字符串。作为裁剪基准的切片标识符。

8.5 延迟执行函数 sleep

sleep 函数用于挂起当前协程/线程，函数会根据传入参数的 数据类型 自动解析对应的时间单位：

参数类型	默认单位	说明	示例
string	显式指定	必须携带单位后缀，支持: ns, us, ms, s, m, h	sleep("1s") (等待 1 秒)
int	秒 (s)	传入整数时，默认按秒计算	sleep(3) (等待 3 秒)
double	毫秒 (ms)	传入浮点数时，默认按毫秒计算	sleep(1.5) (等待 1.5 毫秒)

9.程序入口

脚本执行时，默认寻找全局的 `main()` 函数作为程序的启动入口点。
```RXScript
main() {
  // TODO...
  return 0
}
```

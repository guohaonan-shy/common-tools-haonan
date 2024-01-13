golang常见问题:  

## 1.new与make二者均是为变量分配内存的, 但是这两个关键字的区别是什么？  
##### 1). new可以给所有的类型(基本类型 + 结构体(包括map、channel、slice以及自定义结构体))进行内存分配,但是make只能给slice、channel以及map进行内存分配。  
##### 2). 返回类型不一样，new返回指向变量的指针，make返回变量本身  
##### 3). new分配的空间被清零。make 分配空间后，会进行初始化，即slice、channel以及map分别包含了初始的底层数据结构  
##### 多说一句，内存清零指的是将分配给某个变量或数据结构的内存中的所有位都设置为零，int就是0，string就是"",指针就是nil  
##### 4). 这块明确一点，这两个关键字均是在堆中进行内存分配的  
##### 1️⃣ make因为创建的是动态的数据结构，随着runtime会不断扩容或者缩容，因此会从堆进行内存分配  
##### 2️⃣ new返回的是一个指定类型的指针，这个指针可能随着程序的执行，在创建声明函数意外被引用，比如将该指针作为函数返回值返回到上层函数，不适合在栈中分配
##### case1(基本类型):  
`s1 := new(int)` 声明一个整型变量  
`assert.Equal(t, 0, *s1)` 变量s1是一个整型指针(即指向一块存有int类型的内存)  
##### case2(指针类型):
`s2 := new(*int)` 声明一个整型指针变量，*int也是类型  
`assert.Equal(t, (*int)(nil), *s2)` 
s2的类型是**int，s2是指向一块存有*int类型的内存单元，即s2是一个指向整型指针的指针变量
##### case3(slice, map以及channel): 
`s4 := new([]int)`  
`assert.Equal(t, ([]int)(nil), *s4)`s4则是*[]int，即一个整型切片指针  
`s5 := new(chan int)`  
`assert.Equal(t, (chan int)(nil), *s5)`s5则是一个*chan int，即一个无缓冲整型通道的指针  
`s6 := new(map[int]int)`  
`assert.Equal(t, (map[int]int)(nil), *s6)`s6则是一个*map[int]int，即一个key和value均为int的map指针  
后续深入研究着三个go自带的数据结构时候，在runtime层面的源码，这三个结构本质上是结构体，
底层包含一个pointer指向具体的数据存储单元，这里的nil指的是对应结构内的pointer，
像[]int、chan int、map[int]int也可理解为类型，
即s4指向结构为[]int的内存单元，s4是一个指针，channel和map同理  
##### case4(结构体):
```go
type TestObject struct { 
   val int 
}

type TestStructForMemory struct {
    integer   int
    dict      map[string]string
    ptr       *TestObject
    structure TestObject
}

func main() {
    s7 := new(TestStructForMemory)
    assert.Equal(t, TestStructForMemory{
        integer: 0,
        dict:    (map[string]string)(nil),
        ptr:     nil,
		structure: TestObject{val: 0,},
        }, *s7)
}
```
结构体则是初始化一个结构体，结构体内的field根据类型，默认初始化成对应类型的零值  
说回到本问题上，make只用于slice, channel以及map，并且分配好指向底层数据的内存，而非像new一样的nil

## 2. defer函数的执行逻辑：  
##### 1). 作用：defer延迟函数，主要用于释放资源，收尾工作；如释放锁，关闭文件，关闭链接；捕获panic;  
##### 2). 编程原则：defer函数最好紧跟在资源打开后面，否则defer可能得不到执行，导致内存泄露。如果资源的创建和对应的defer中间有一段程序，并且return了，那么之后的defer不会执行
##### 3). 执行原理：待到return，开始处理defer，按照LIFO的顺序处理多个defer逻辑

关键知识点：无名返回值 vs 有名返回值，无参函数 vs 有参函数  
首先，先阐明一点，go中的参数传递均为值传递，无论是基本类型、结构体都是将值复制到对应的参数变量内，当然如果复制的是一个指针变量，则类同引用传递 eg. 切片  
回到defer这个话题上，看几个经典的case  
##### case1(有参函数、无参函数)
```go
package main

import "fmt"

func test(a int) {//无返回值函数
	defer fmt.Println("1、a =", a) //方法
	defer func(v int) { fmt.Println("2、a =", v)} (a) //有参函数
	defer func() { fmt.Println("3、a =", a)} () //无参函数
	a++
}
func main() {
	test(1)
}
```
`3、a = 2`  
`2、a = 1`  
`1、a = 1`  
顺序如之前所言，按照先入后出的顺序执行，但是2和1的输出打印值，没有收到a++这行代码的影响呢？  
答：对于有参函数(有参匿名函数和方法)，运行时运行到defer代码会捕获a的当前值，return之后执行defer函数的具体逻辑，所以前两行没有收到1和2的影响  
但是对于第三行的无参函数函数来说，内部引用的变量a被闭包捕获，保存的是外部函数中变量a的引用，那么具体的值需要运行到defer函数逻辑时，才能确认值，所以返回了执行了a++的结果  
##### 总之，引用go官方文档中的话：A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns
##### defer只是将一个函数执行触发的时机延迟到了return的时刻；而函数的传参，按照代码顺序是在程序运行到defer语句(区别执行defer的函数逻辑和执行defer语句)就确认了的

##### case2(有名返回值 vs 无名返回值)
```go
package main

import "fmt"

func a1() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i //或者直接写成return
}

func a2() int {
    var i int
    defer func() {
        i++
        fmt.Println("defer2:", i)
    }()
    defer func() {
        i++
        fmt.Println("defer1:", i)
    }()
    return i
}
```  
a1和a2两个函数唯一的区别在于函数签名中的返回值，一个是有名的，一个是无名的：  
本质上，return并非一个原子的操作，先将变量值复制给“返回值”(无名)，然后执行defer逻辑，最后返回给上层调用“返回值”  
```
    // 承接之前的两个case，有名 vs 无名，假定有名返回就是i
    if 返回值无名:
        default_value(类型编译时确认) = i
    
    if has_defer:
        LIFO -> defer3(i), defer2(i), defer1(i)
    
    if 返回值无名:
        return default_value 
    else:
        return i    
```
1). 对于无名返回值，return会初始化一个对应类型的返回值，就取名为“返回值”吧，初始值为return后的变量值，然后执行defer逻辑，最后返回  
2). 有名返回值，return并不会执行初始化过程，初始化早在函数开始时就初始化好了，然后执行defer的逻辑，最后返回  
最后，上述两个case的关键区别也出来了，本质上就是defer操作的整型变量和返回值是否是同一变量；当然，如果返回值是指针的话，那defer操作是会对无名返回值有影响的了，原理不再赘述  
### 在这个问题思考结束之后，感谢这篇文档：https://blog.csdn.net/Cassie_zkq/article/details/108567205 以及chatgpt与我的对话
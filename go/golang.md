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

## 3. rune和byte类型：  
这个问题很简单本质上，但是单独拎出来说是为了概念的明确，不要带着模模糊糊的概念往下走，这条路是长期的，不是面试一次就完了的  
```
var str = "abcdefg中国  
for i, s := range str {
    print(str[i])
    print(s)
} 
```
这两个输出的内容一样吗？- 可能一样，也可能不一样，怎么解释？  
#### 可能一样 - golang的字符串底层使用字节数组实现的，对于基本的ASCII码来说，一个字节即可表示，所以二者输出是相同的  
#### 首先多说一句，在计算机科学的领域中unicode是一个文字符号的唯一标识，是一个共识协议; 经过不同的编码方式，如utf-8，将unicode转化为字节数组从而在计算机系统存储和传递
#### 可能不一样 - 回到这个问题，因为底层存储是以utf-8编码实现的字节数组，而for range遍历的是字符串的unicode(字符的唯一标识), 如果字符unicode -> utf-8的字节数组不止一个字节，那边遍历到下一个字符字节的游标需要往后移相应的数(参考go_test.go 170)
#### 即str[i]返回的是utf8编码之后的字节，如果是非一个字节表示的符号，此时返回的是utf-8编码的第一个字节，而s返回的是该符号的unicode，如需要，需要自己编码

## 4. slice：
切片的扩容不再赘述，随处可见的文章都有讲解
```
type slice struct {
    data unsafe.pointer
    len int
    cap int
}
```
首先，golang函数传参是值传递，因此切片结构体以副本形式传递给调用函数，但是切片存储数据的底层是一个指针，指向动态数组的首元素  
在容量充足的情况，即不考虑数组扩容，多个切片变量会共享同一块内存，修改是同步的  
更进一步，append作为切片操作最常用的函数，根据当前切片的长度进行元素尾追加，看下面的case
```
s := make([]int, 0, 4)
s = append(s, 1)
s1 := append(s,2)
s = append(s, 3)
```
很有意思的问题，现在s的长度是多少，s1[1]还是2吗？  
答案是s的长度是2，s1[1]值变为3  
首先，因为s的容量为4，因此append操作的时候是没有扩容，重新分配内存存储的；
执行完第三行s1的赋值代码后，s1这个切片的data指针是存有两个元素的内存块，并与s共享；
但是s的len是1，就代表切片s无法索引到data[1]，之后s的append操作，按照s的长度为1，容量为4进行操作，将3放入data[1]的位置，覆盖了s1[1]

## 5. map: 
map底层数据结构
```
type hmap struct {     //hashmap
    count     int                  // 元素个数     
    flags     uint8     
    B         uint8                // 扩容常量相关字段B是buckets数组的长度的对数 2^B     
    noverflow uint16               // 溢出的bucket个数     
    hash0     uint32               // hash seed     
    buckets    unsafe.Pointer      // buckets 数组指针     
    oldbuckets unsafe.Pointer      // 结构扩容的时候用于赋值的buckets数组     
    nevacuate  uintptr             // 搬迁进度     
    extra *mapextra                // 用于扩容的结构，里面存有新旧桶里面溢出的bmap指针
}

type bmap struct { // bucket的数据结构
    topbits  [8]uint8 // hash(key)的高8位
    keys     [8]keytype // 8个key
    values   [8]valuetype // 8个value
    pad      uintptr
    overflow uintptr 如果相同该bmap无空位，会新建bmap与该bmap相连
}
```
数据结构类似于C的hashmap，不同于C的一个bmap存储一个kv键值对，go的bmap一个结构存储8个kv键值  
先不考虑扩容，map是如何查找和写入的？  
#### 查找逻辑：
1. 根据hashmap内的hash种子计算出key的hash 
2. has值与bucket数目取模确认bucket位置
3. bucket内的bmap结点进行遍历，先比较高位的hash是否相同，若相同在比较key是否相同，知道找到目标key或者遍历完仍未发现目标key

#### 写入逻辑：
1. 根据key值算出哈希值
2. 哈希值与bucket数目取模确定bucket位置
3. 查找该key是否已经存在，如果存在则直接更新值
4. 如果没找到将key，将key插入

当map经过多次的增删改之后，会出现两种极端情况，一种是bucket内极度紧致，负载因子过大，导致hashmap的查找退化成了链表；另一种则是bucket内极度稀疏，导致空间使用率很低；因此map需要进行增量扩容和等量扩容来不断调整map的内容

#### 增量扩容(负载因子过大，需要rehash):
1. 每当添加元素的过程中，hashmap会根据`负载因子 = 元素个数 / bucket数目>6.5`判断是否需要增量扩容
2. 当扩容开始，首先`B+1`，则新的map的bucket数目为`2^(B+1)`, 申请新的bucket内存，并将原来的指针设置为oldbuckets
3. 新的key根据写入逻辑，写入新的buckets存储内，同时更新flag和nevacuate这两个变量，表示map正在进行迁移；
4. 扩容期间，每次map的读写都会先访问oldBucket，如果oldbucket的tophash[0]是正常的hash填充值，则说明该buckets没有迁移，将该bucket内所有的kv rehash至新桶(保证一个bucket要不然包含全部的key，要不然全是空)
5. 当迁移完成后，oldBuckets为空,buckets是新的bucket空间

#### 等量扩容 (overflow的个数很大，但负载因子较低):
1. 当hmap中的nonoverflow数目过大时，触发扩容，只不过新申请的buckets个数同原先一样，flag添加一个新的标签sameSizeGrow
2. 其他步骤同上述扩容过程

## 6. channel
channel底层数据结构：
```
type hchan struct {
	qcount   uint           // current number of elements in this channel
	dataqsiz uint           // the total length of this channel
	buf      unsafe.Pointer // mem unit that points to an array of dataqsiz elements
	elemsize uint16         // element size(eg.int, string, struct{})
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index if this channel has buffer 
	recvx    uint   // receive index if this channel has buffer 
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters
	
	lock mutex // 防止其他的receive或者并发的send事件
}
```
channel整体由一个环型数组和两个队列(接收事件+发送事件)组成，如果channel是一个无缓冲通道，那环形数组不会用到；同时用锁避免channel出现race condition   
#### channel的创建：
当创建channel时，make函数会根据size判断是否需要给buf分配内存空间；如果size不为0，那么程序会在固定的channelSize基础上，加上`元素类型大小 * 元素个数`的内存空间  

假设我们初始化一个channel
`c := make(chan int, 5)` 即c是一个长度为5，元素类型为int的channel
#### send:
1. 当缓冲队列还没满时，程序将入队的元素变量`复制`到队列下一个可插入位置(即channel不是共享内存，而是通信)，更新可插入位置和元素数目  
2. 当缓冲队列满了，程序会将sudog(同步元素以及其关联的goroutine的实体)放入sendq末尾，然后调用gopark使goroutine陷入等待，同时p调度一个新的goroutine执行
#### receive:
1. 当缓冲队列没满时，程序将recvx的元素，复制到目标对象，迭代recx和元素数目;  
2. 当缓冲队列满了，程序将recvx的元素，复制到目标对象；然后将接收等待队列中队尾对象复制到buf内，然后迭代recx和元素数目
3. 当缓冲队列为空，将接收事件添加等待队列，并调用gopark切换执行goroutine  
##### 接收事件有个细节: 接收事件`<-chan`的左边可能有写入变量也可能没有；
##### 如果有目标变量(eg.`for task := range chan`), 则channel会将recev的对象复制到目标对象，并清空对应位置的内存，迭代recvx)
##### 如果无目标变量(eg.`for _ := range chan`)，则省略目标变量复制的步骤

当channel长度为0，即无buf缓冲，则按照send队列满和receive队列为空来处理

#### close:
设置为关闭状态，将两个事件的等待队列内所有sudog更新为runnable的状态(但不清空缓冲队列的元素)，等待对应p的调度  
因此，当关闭之后，如果仍往channel继续添加元素，`会直接panic；`  
接收元素会消费完队列内所有的元素，之后每次`<-c`都只会得到空值，但是`for _ = range c`不再执行

#### channel为nil的情况：
channel为nil，send和recv事件会直接将对应的goroutine挂起，同时fatal
close(channel)也无法解决，因为nil的channel关闭会直接panic
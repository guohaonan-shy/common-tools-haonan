golang常见问题:  

## 1.new与make二者均是为变量分配内存的, 但是这两个关键字的区别是什么？  
##### 1). new可以给所有的类型(基本类型 + 结构体(包括map、channel、slice以及自定义结构体))进行内存分配,但是make只能给slice、channel以及map进行内存分配。  
##### 2). 返回类型不一样，new返回指向变量的指针，make返回变量本身  
##### 3). new分配的空间被清零。make 分配空间后，会进行初始化，即slice、channel以及map分别包含了初始的底层数据结构  
##### 多说一句，内存清零指的是将分配给某个变量或数据结构的内存中的所有位都设置为零，int就是0，string就是"",指针就是nil  
##### 4). 这块明确一点，这两个关键字均是在堆中进行内存分配的  
##### 1️⃣ make因为创建的是动态的数据结构，随着runtime会不断扩容或者缩容，因此会从堆进行内存分配  
##### 2️⃣ new返回的是一个指定类型的指针，这个指针可能随着程序的执行，在创建声明函数以外被引用，比如将该指针作为函数返回值返回到上层函数，不适合在栈中分配
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
var str = "abcdefg中国"  
for i, s := range str {
    print(str[i]) // 返回的是unicode
    print(s) // 返回的是utf-8 的计算机编码
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

## Go的调度模型(GMP)
关于GMP分别是什么: goroutine(调度单位)、processor(处理器)以及machine(工作线程)  
1. OS level的工作线程m不再关注线程之间切换调度的细节，而是通过处理器来获取下一个goroutine  
2. user level的协程goroutine, 只关注当前自己关联的是哪个m以及自己的祖先g，不关心其他g的状态与关系
3. p维护多个本地处理器以及一个全局处理器，通过本地对象和全局对象的交互，提升go程序的工作效率  

首先，在讨论go的调度模型之前，我们必须要了解甚至明白runtime调度的基本单位goroutine以及其生命周期
### 创建
都知道goroutine怎么用  
`go func() {...}`  
runtime先为该goroutine申请一个栈空间，初始化goroutine的状态和上下文，然后放入队列等待调度
### 销毁
整体上，goroutine的状态更新为dead，然后进行下一轮调度，执行下一个runnable的goroutine
### 调度
首先从本地队列了获取runnable的goroutine进行执行，如果本地队列g的数目不足，p会去其他p的本地队列或者全局队列获取一半的goroutine执行

看到着，有几个问题对程序如何运行至关重要
1. 当goroutine很多的时候，goroutine的内存申请和释放都要切换到内核态，由操作系统对内存进行处理，效率会有损耗，怎么解决？
2. GMP中哪些状态和上下文，需要engineers了解呢?  
3. 放入队列的goroutine 怎么配合本地和全局队列从而保证调度的公平性呢？

关于第一个问题：  
其实每个processor以及全局scheduler都维护了一个协程池(cache),当goroutine销毁时，内存空间并不会被os回收，而是放入本地或者全局的协程池；
本地只会维护size为64的池子，一旦超过会将一半的的协程单元放入全局协程池，供其他proc复用，进一步提升资源利用；这样创建和销毁goroutine可以直接复用之前已经申请过的空间，减少内核切换的频率；
这里还有一个细节，如果栈的大小不是startingStackSize，会将空间清零；本地队列统一一个队列维护，但是全局队列需要维护在NstackQ(无栈队列)；后续看到gc关注这块

关于第二个问题：  
新创建的goroutine优先放入对应m的p上的本地队列，并且当caller goroutine的时间片结束后，会优先调度执行该g；
p上有个指针runnext(should be run next instead of what's in runq if there's time remaining in the running G's time slice),
即会优先执行runnext的协程，而非runq队列内的，如果当前g的事件片还有剩余  

更好的局部性: 
1. 因为这两个goroutine是父子关系，访问统一处理器可以提升处理器缓存的命中率，提升响应速度；
2. 减少同其他goroutine的竞争，更快的启动和响应

关于第三个问题：  
本地队列的存在极大减少了全局队列的锁冲突，同时调度过程，定期会从全局队列获取goroutine，避免全局队列内的g饥饿；

## Go的内存管理：
page: Go程序与OS之间进行内存分配和释放的基本管理单元  
mspan: Go程序内部进行内存的管理、分配以及释放的基本单元，通常都是由整数个page组成
sizeClass: 每个mspan和每个object都有对应的sizeClass，每个mspan按照它自身的属性Size Class的大小分割成若干个object，每个object可存储一个对象

当创建一个大小为n bytes的对象时，根据object的list可以得到sizeClass,然后找到对应mspan的大小  
问题就变成在指定sizeClass的mspan内寻找一个未使用的object空间进行分配

再讲具体的分配策略之前，需要阐明一些高效的内存组织方式
### 池化技术：
程序动态申请内存空间，是要使用系统调用的，比如 Linux 系统上是调用 mmap 方法实现的。但对于大型系统服务来说，直接调用 mmap 申请内存，会有一定的代价。比如：
1. 系统调用会导致程序进入内核态，内核分配完内存后（也就是上篇所讲的，对虚拟地址和物理地址进行映射等操作），再返回到用户态。 
2. 频繁申请很小的内存空间，容易出现大量内存碎片，增大操作系统整理碎片的压力。 
3. 为了保证内存访问具有良好的局部性，开发者需要投入大量的精力去做优化，这是一个很重的负担。

如何解决上面的问题呢？有经验的人，可能很快就想到解决方案，那就是我们常说的对象池（也可以说是缓存）  
假设系统需要频繁动态申请内存来存放一个数据结构，比如 [10]int 。那么我们完全可以在程序启动之初，一次性申请几百甚至上千个 [10]int 。这样完美的解决了上面遇到的问题：
1. 不需要频繁申请内存了，而是从对象池里拿，程序不会频繁进入内核态
2. 因为一次性申请一个连续的大空间，对象池会被重复利用，不会出现碎片。
3. 程序频繁访问的就是对象池背后的同一块内存空间，局部性良好。

这样做会造成一定的内存浪费，我们可以定时检测对象池的大小，保证可用对象的数量在一个合理的范围，少了就提前申请，多了就自动释放。
如果某种资源的申请和回收是昂贵的，我们都可以通过建立资源池的方式来解决，其他比如连接池，内存池等等，都是一个思路

### Golang的内存池：
首先golang的内存池是在程序启动时，申请了一大块虚拟内存存入一个mheap的结构，mheap一方面跟os交互进行内存的申请和释放；一方面以mspan为单位对go程序内的内存进行管理

为了提升性能和减少锁冲突，mheap分配了一块空间叫mcentral，同时每个processor内又一个mcache结构，给goroutine分配heap空间，大幅减少了访问中心化的mcentral和mheap的概率

当一个对象申请内存空间，首先根据其大小判断是否为空对象，如[0]int, struct{},则会返回一个固定指针
```
if size == 0 {
    return unsafe.Pointer(&zerobase) // 减少空对象进行内存申请释放的动作
}
```
如果该对象是一个非指针对象且大小属于tiny对象(小于16 bytes),则先对内存进行对齐，比如如果申请对象的大小为7 byte，则对齐至8 byte；如果tinyCache没有值，则从sizeClass = 2进行对象的分配

对于指针对象或者普通对象(小于32KB(大对象))，则根据class_size获取对应的mspan，然后先从span缓存的对象池看是否有free object，如果没有再从span后面分配(层层缓存啊)，
根据objectSize和freeIdx判断是否还有可分配的object；如果没有向mcentral对应的class申请获取更多mspan

mcentral每个classSize分有partial和free两个span set，优先从partial分配指定classSize的page大小；若无再从空的span裁剪出合适的pages给到mcache

如果mcentral也没有合适的空间，那就只能从mheap获取，或者从os申请更大的内存

总之，Go在程序启动时，会向操作系统申请一大块内存，之后自行管理。
Go内存管理的基本单元是mspan，它由若干个页组成，每种mspan可以分配特定大小的object。
mcache, mcentral, mheap是Go内存管理的三大组件，层层递进。mcache管理线程在本地缓存的mspan；mcentral管理全局的mspan供所有线程使用；mheap管理Go的所有动态分配内存。
极小对象会分配在一个object中，以节省资源，使用tiny分配器分配内存；一般小对象通过mspan分配内存；大对象则直接由mheap分配内存。

## Go的垃圾回收
go的runtime会定时或者当内存使用达到指定阈值时，触发垃圾回收过程，这个过程分为标记、染色以及清除
首先，当cpu获取到了gc线程的执行权，先进入标记染色的状态，从根对象出发，将引用对象标记为灰色，并放入灰色队列，遍历完成的对象标记为黑色；知道灰色队列为空，STW，清除所有的白色对象

根对象包括每个活动的goroutine的栈上的变量，局部变量和指向堆上对象的指针，全局变量和静态变量。

然而，传统的垃圾收集算法会在垃圾收集的执行期间暂停应用程序，一旦触发垃圾收集，垃圾收集器会抢占 CPU 的使用权占据大量的计算资源以完成标记和清除工作，然而很多追求实时的应用程序无法接受长时间的 STW

因此，golang提出了并发增量的处理模式，即在应用程序的运行过程中，gc collector会获取到cpu执行权，在应用程序修改对象之间关系的同时，标记清除；  

但是，这样会衍生出了新的问题，即已被标记为黑色的对象被应用修改直接引用某个白色的对象，即清除了本不应该清除的对象，因此为了解决这个问题，采用了写入屏障防止该情况出现

大致的准则，两条满足其一即可保证没问题：
1. 强三色不变性：黑色对象不会指向白色对象，只会指向灰色对象或者黑色对象
2. 弱三色不变性：黑色对象指向的白色对象必须包含一条从灰色对象经由多个白色对象的可达路径；这样保证后续标记过程中，该对象可达

通常堆上的根对象会使用写入屏障进行保护，会栈因为周期跟随goroutine，所以不会用写入屏障进行保护
插入屏障（Dijkstra）- 灰色赋值器：
当应用程序每次插入新对象，将原无引用的A对象，增加了新的引用下游；或修改引用关系，将A对象引用从B修改为C，则新对象会被标记为灰色；满足第一条强三色一致性

删除屏障 （Yuasa）- 黑色赋值器
当A->B的引用关系被删除，如果B为白色，则将B标记为灰色；防止B及其引用对象被错误清理
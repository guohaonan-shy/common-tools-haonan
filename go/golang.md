golang常见问题:  

1.new与make二者均是为变量分配内存的, 但是这两个关键字的区别是什么？  
1). new可以给所有的类型(基本类型 + 结构体(包括map、channel、slice以及自定义结构体))进行内存分配,但是make只能给slice、channel以及map进行内存分配。  
2). 返回类型不一样，new返回指向变量的指针，make返回变量本身  
3). new分配的空间被清零。make 分配空间后，会进行初始化，即slice、channel以及map分别包含了初始的底层数据结构  
多说一句，内存清零指的是将分配给某个变量或数据结构的内存中的所有位都设置为零，int就是0，string就是"",指针就是nil
  
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
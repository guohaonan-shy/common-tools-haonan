package _go

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
)

type Object struct {
	number int
}

type Toy struct {
	Object
}

func Test_Toy(t *testing.T) {
	toy1 := new(Toy)
	toy2 := &Toy{}
	print(toy1)
	print(toy2)
}

func Test_Defer(t *testing.T) {
	func(a int) {
		defer fmt.Printf("\n1. defer a=%d\n", a) // 一样， 将入参作为值传递给对应方法
		defer func(a int) {                      // 编译之后已经将a值传递到defer的函数结构内了
			fmt.Printf("\n2. defer a=%d\n", a)
			print(&a)
		}(a)
		defer func() {
			fmt.Printf("\n3. defer a=%d\n", a)
			print(&a)
		}()
		a++
		print(&a)
	}(1)
}

type TestInterface interface {
	Test(a, b int)

	Insert(a int) bool
}

type TestStruct struct {
}

func (t TestStruct) Test(a, b int) {

}

func (t TestStruct) Insert(a int) bool {
	return false
}

func Test_Interface(t *testing.T) {
	var s TestInterface = TestStruct{}
	ty := reflect.TypeOf(s)

	print(ty)

	print(s)
}

func DiGui(mapping map[int]struct{}) {
	key := rand.Intn(5)
	if _, ok := mapping[key]; ok {
		return
	}
	mapping[key] = struct{}{}
	DiGui(mapping)
	return
}

func Test_Map(t *testing.T) {
	dict := make(map[int]struct{}, 0)
	DiGui(dict)
}

func sliceParam(s []int) {
	if len(s) == 5 {
		return
	}
	sliceParam(append(s, 1))
	s = append(s, 2)
}

func Test_Slice(t *testing.T) {
	s1 := make([]int, 0, 100)
	sliceParam(append(s1, 1))
}

type TestObject struct {
	val int
}

type TestStructForMemory struct {
	integer   int
	dict      map[string]string
	ptr       *TestObject
	structure TestObject
}

func Test_NewMake(t *testing.T) {
	s1 := new(int)
	println(&s1)
	assert.Equal(t, 0, *s1)
	//assert.Equal(t, (*int)(nil), *s1)

	s2 := new(*int)
	assert.Equal(t, (*int)(nil), *s2)

	s3 := new(string)
	assert.Equal(t, "", *s3)

	s4 := new([]int)
	//assert.Equal(t, []int{}, *s4)
	assert.Equal(t, ([]int)(nil), *s4)

	s5 := new(chan int)
	//assert.Equal(t, make(chan int, 0), *s5)
	assert.Equal(t, (chan int)(nil), *s5)

	s6 := new(map[int]int)
	//assert.Equal(t, map[int]int{}, *s6)
	assert.Equal(t, (map[int]int)(nil), *s6)

	s7 := new(TestStructForMemory)
	assert.Equal(t, TestStructForMemory{
		integer: 0,
		dict:    (map[string]string)(nil),
		ptr:     nil,
		structure: TestObject{
			val: 0,
		},
	}, *s7)
}

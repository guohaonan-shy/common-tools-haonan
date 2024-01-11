package _go

import (
	"fmt"
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

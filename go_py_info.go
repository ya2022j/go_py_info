package main

// '' 表示rune类型，ASCII编码 ，字节    " "表示string ，字符串 `` 可以输出原本的大字符串，也是string
import "fmt"

// 具名函数对比python中的def函数，
// 1 都必须带类型，但是返回值如果只有一个可以省略返回参数名
// 2 如果 返回参数名写清楚了，那么在｛｝中可以只写一个return
// 3 函数作为参数送入其他函数中，对比python中的lambda表达式集合内置函数filter等使用
// var a interface{} 空接口 对比 python中的 a=None
// 哈希与不可哈希
// golang中值类型int系列、float系列、bool、string、数组和结构体（只是拷贝参数）；--->对比python中list和dict的地位即可
// 指针类型(就改变了内存地址),指针、slice切片、管道channel、接口interface、map、函数等
//new(int) --->new返回指针
// 1匿名函数－－可以自动执行  最后的()是输入参数的地方
// 2 defer 可以逆序操作，但是也是排在return之后

//接口类型你个判定，3种方法  1  if v,ok :=varI.(T);ok{}  2 switch str:=value.(type){}  str,ok :=value.(string) if ok{}
//变参函数对比python中的*args **kwars
// ...int 变参用    i := []int{1,2,3}  ---> test(i...)  i...解包用
//函数作为一个参数，起到了函数内部的控制的作用
type student struct {
	name string
}

func filter(stu []student, f func(s student) bool) []student {
	var r []student
	for _, s := range stu {
		if f(s) == true {
			r = append(r, s)
		}
	}
	return r
}
func test(a ...int) (b int) {
	//匿名函数 or  匿名函数也可以封装   f = (data int) {fmt.Println("hell") } --->f()
	func(data int) {
		fmt.Println("hell")
	}(100)

	// defer
	fmt.Println("start--->")
	defer fmt.Println("this is defer --->3") // this is defer --->1 2 3
	defer fmt.Println("this is defer --->2")
	defer fmt.Println("this is defer --->1")
	b = a + 3
	return
}

//接口(接口中是方法)+类型(类型中是具体的类型，int也可以使用指针)+方法(有接收器，相当于类型的实例化) 后面使用类型或结构都是可以的 类比python中的类，类中函数，还有函数的参数

type I interface {
	f()
}
type A struct {
	Name string
}

//方法中可以传值也可以传指针
//传指针时    //调用方法前，还得用类型区确认一下 A("asdf")
type Recevie{
	Name string 
}
// 在写方法时，接收器里的结构体常写成指针类型 func (re *Recevie)
//所谓面向对象，golang中就是对接口中的类型(多用结构体struct)，在类型作为接收器时，使用方法实现继承，重写（多态）

func (a A) f() { // 传值时 a = "asd" ret :=
	fmt.Println(a)
}
func main() {
	arr := [...]int{1, 2, 3} //arr ---->[1 2 3]  golang数组对比python中的列表list

	slice := arr[1:2] // slice --->[]int{1, 2, 3} [1,2,3] or slice := arr[1:2] -->[2]   golang的切片有两种方式，对比python中的索引

	m := make(map[string]int) // or m = map[string]int{"a":1}   golang中map对比python中的字典dict，特别有一个make可以用来分配内存空间
	m["a"] = 1

	fmt.Println(slice, arr, m)
	s := "aaa"
	ret := A{"adsf"}
	ret.f() //ret.f()--->  "aaa"

}

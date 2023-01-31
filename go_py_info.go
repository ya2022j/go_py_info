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
//interface可以被任意的对象实现   一个对象可以实现任意多个interface   任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface
//interface值-->interface里面到底能存什么值呢？如果我们定义了一个interface的变量，那么这个变量里面可以存实现这个interface的任意类型的对象。
//interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现
//嵌入interface -->如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method
// interface结构的嵌套，可以类似python的类 继承  ，广度优先还是深度优先
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


//反射就是能检查程序在运行时的状态 --->运用reflect包--->使用reflect一般分成三步，
//下面简要的讲解一下：要去反射是一个类型的值(这些值都实现了空interface)，
//首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)

使用reflect一般分成三步，下面简要的讲解一下：要去反射是一个类型的值(这些值都实现了空interface)，首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)。这两种获取方式如下：

t := reflect.TypeOf(i)    //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
v := reflect.ValueOf(i)   //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
转化为reflect对象之后我们就可以进行一些操作了，也就是将reflect对象转化成相应的值，例如

tag := t.Elem().Field(0).Tag  //获取定义在struct里面的标签
name := v.Elem().Field(0).String()  //获取存储在第一个字段里面的值
获取反射值能返回相应的类型和数值

var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
最后，反射的话，那么反射的字段必须是可修改的，我们前面学习过传值和传引用，这个里面也是一样的道理。反射的字段必须是可读写的意思是，如果下面这样写，那么会发生错误

var x float64 = 3.4
v := reflect.ValueOf(x)
v.SetFloat(7.1)
如果要修改相应的值，必须这样写

var x float64 = 3.4
p := reflect.ValueOf(&x)
v := p.Elem()
v.SetFloat(7.1)


// goroutine是Go并行设计的核心-->goroutine是通过Go的runtime管理的一个线程管理器

//多个goroutine运行在同一个进程里面，共享内存数据
//想要发挥多核处理器的并行，需要在我们的程序中显式调用 runtime.GOMAXPROCS(n) 告诉调度器同时使用多个线程。GOMAXPROCS 设置了同时运行逻辑代码的系统线程的最大数量，并返回之前的设置。如果n < 1，不会改变当前设置

//goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。
//那么goroutine之间如何进行数据的通信呢--->channel类型。定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel：

//channel通过操作符<-来接收和发送数据



// memory <----共享内存-->  channel(args chan) <--send/receive--> (app)/int/string/struct
// ch <- v    // 发送v到channel ch.
// v := <-ch  // 从ch中接收数据，并赋值给v
//默认情况下，channel接收和发送数据都是阻塞的，
//除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具。


//Buffered Channels
//ch := make(chan type, value)
//当 value = 0 时，channel 是无缓冲阻塞读写的，
//当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。

// Range和Close

// 使用range遍历往通道里面添加数据，然后goroutine再与内存交互
// 使用close关闭通道

//可以显式的关闭channel，生产者通过内置函数close关闭channel。关闭channel之后就无法再发送任何数据了，在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭。

// 记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic

// 另外记住一点的就是channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的


// select --->多个channel
//上面介绍的都是只有一个channel的情况，那么如果存在多个channel的时候，我们该如何操作呢
//Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。

// select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是随机的选择一个执行的。

//在select里面还有default语法，select其实就是类似switch的功能，default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）。
// 超时
// 有时候会出现goroutine阻塞的情况，那么我们如何避免整个程序进入阻塞的情况呢？我们可以利用select来设置超时，通过如下的方式实现：

func main() {
    c := make(chan int)
    o := make(chan bool)
    go func() {
        for {
            select {
                case v := <- c:
                    println(v)
                case <- time.After(5 * time.Second):
                    println("timeout")
                    o <- true
                    break
            }
        }
    }()
    <- o
}
// /runtime goroutine
// runtime包中有几个处理goroutine的函数：

// Goexit

// 退出当前执行的goroutine，但是defer函数还会继续调用

// Gosched

// 让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。

// NumCPU

// 返回 CPU 核数量

// NumGoroutine

// 返回正在执行和排队的任务总数

// GOMAXPROCS

// 用来设置可以并行计算的CPU核数的最大值，并返回之前的值。

// break    default      func    interface    select
// case     defer        go      map          struct
// chan     else         goto    package      switch
// const    fallthrough  if      range        type
// continue for          import  return       var
// var和const ，变量和常量申明，可以批量声明
// package和import 
// func 用于定义函数和方法
// return 用于从函数返回
// defer 用于类似析构函数  ---> 多个defer可以逆向操作
// go 用于并发
// select 用于选择不同类型的通讯
// interface 用于定义接口，参考2.6小节
// struct 用于定义抽象数据类型，参考2.5小节
// chan用于channel通讯（在并发时，作为基础类型或应用与内存之前交互共享数据的媒介）
// type用于声明自定义类型（type  interface｛｝  type struct{}）
// map用于声明map类型数据 ---> make(map[string]int)
// range用于读取slice、map、channel数据

	// check type  
	// 1. %T fmt.Printf
	// 2. reflect.TypeOf(s)
	// 2. reflect.TypeOf(s).Kind()
	
	fmt.Printf("%T",s)

	
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(s).Kind())

	//     1.数值类型和string类型之间的相互转换可能造成值部分丢失；其他的转换仅是类型的转换，不会造成值的改变;
        // 2.string和数字之间转换可使用标准库strconv;
	// strconv.Itoa(8)---> "8"
	// strconv.Atoi("102")--->102
        // 3.或者想要转换byte数组（[]byte或 []rune）为string字符串类型，这种情况下可以用string(): []byte -> string 

 

	
	//Go虽然保留了指针，但与其他编程语言不同的是，在Go当中不支持指针运算以及"->"运算符，而直接采用"."选择符来操作指针目标对象的成员。
// 操作符"&"取变量地址，使用“*”通过指针间接访问目标对象
// 默认值是nil而非NULL

// 一个指针变量可以指向任何一个值的内存地址。指针变量类似于变量和常量，在使用指针前需要声明。指针声明格式如下：
// var var_name *var-type

// 如何使用指针
// 1. 定义指针变量
// 2. 为指针变量赋值
// 3. 访问指针变量中指向地址的值
// 4. 在指针类型前面加上*号(前缀)来获取指针所指向的内容。

func main() {
	var b int
	a := 2
	b = a  //传值引用 ---> 只是拷贝。无法改变原有变量

	b = 3    //传值引用---> 只是拷贝。无法改变原有变量
	
	fmt.Println(b)
	fmt.Println(a)
	var ptrInt *int // 1. 定义指针变量  equal : ptrInt := new(int)

	ptrInt = &a         // 2. 为指针变量赋值   //地址引用 改变引用内内容，原有变量也发生变化
	fmt.Println(ptrInt) // 3. 访问指针变量中指向地址的值

	*ptrInt = 66  // 地址引用 改变引用内内容，原有变量也发生变化

	fmt.Println(strings.Repeat("-", 12))

	fmt.Println(*ptrInt) // 4. 在指针类型前面加上*号(前缀)来获取指针所指向的内容。
	fmt.Println(a)

}


//接口interface

// 接口是一个多个方法签名的集合
// 只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示声明实现了哪个接口，这称为Structural Typing
// 接口只有方法声明，没有实现，没有数据字段
// 接口可以匿名嵌入其他接口，或嵌入结构中
// 将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针
// 只有当接口存储的类型和对象都为nil时，接口才等于nil
//接口调用不会做receiver的自动转换
// 接口同样支持匿名字段方法
// 接口也可以实现类似OOP中的多台
// 空接口可以作为任何类型数据的容器(var a interface{})

// 类型断言
// 通过类型断言的ok pattern可以判断接口的数据类型
// 使用type swtich则可针对空接口进行比较全面的类型判断
// 接口转换
// 可以将拥有超集的接口转换为子集的接口


//反射reflection

// 反射可大大提高程序的灵活性，使得interface{}有更大的发挥余地
// 反射使用typeOf和valueOf函数从接口中获取目标对象信息

func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct { // 1. 类型判断
		fmt.Println("XXX")
		return
	}

	v := reflect.ValueOf(o)

	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ { // 2.字段信息
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s : %v = %v \n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ { // 3. 方法信息
		m := t.Method(i)
		fmt.Printf("%6s : %v \n", m.Name, m.Type)
	}
}

//反射会将匿名字段作为独立字段(匿名字段本质)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User  //匿名字段
	title string
}

func main() {
	m := Manager{User: User{1, "s", 12}, title: "ddd"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.Field(1))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
}

// 想要利用反射修改对象状态，前提是interface.data是settabele,即pointer-interface
// 基本数据类型：

func main() {
	x := 12
	v := reflect.ValueOf(&x)
	fmt.Println(v)
	v.Elem().SetInt(99)
	fmt.Println(x)
}


//复杂数据类型：

package main
 
import (
	"fmt"
	"reflect"
)
 
type User struct {
	Id   int
	Name string
	Age  int
}
 
func main() {
	u := User{1, "LJ", 19}
	Set(&u)
	fmt.Println(u)
}
 
func Set(o interface{}) {
	v := reflect.ValueOf(o)
 
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() { //指针 是否可设置判断
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}
 
	f := v.FieldByName("Name") //通过字段名获取字段
	if !f.IsValid() {          //字段是否存在判断
		fmt.Println("BAD")
		return
	}
 
	if f.Kind() == reflect.String { //类型判断
		f.SetString("BYEBYE")
	}
}

//通过反射可以“动态”调用方法

package main
 
import (
	"fmt"
	"reflect"
)
 
type User struct {
	Id   int
	Name string
	Age  int
}
 
func (u User) Hello(name string) {
	fmt.Println("Hello", name, "my name is", u.Name)
}
 
func main() {
	u := User{1, "LJ", 19}
	v := reflect.ValueOf(u)
 
	mv := v.MethodByName("Hello")  //根据名称获取方法
 
	args := []reflect.Value{reflect.ValueOf("joe")}  //参数
	mv.Call(args)  //调用
}


// Channel
// Channel是goroutine沟通的桥梁，大都是阻塞同步的
//通过make创建，close关闭

// func main(){
// 	c := make(chan bool)
// 	go func(){
// 		fmt.Println("gogooooo")
// 		c <- true
// 	}()
// 	<-c
// }

//Channel是引用类型
//可以使用for range来迭代不断操作channel

// func main() {
// 	c := make(chan bool)

// 	go func() {
// 		fmt.Println("gooo")
// 		c <- true
// 		close(c)
// 	}()

// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }

// Select
// 可处理一个或多个channel的发送与接收
// 同时有多个可用的channel时按随机处理
// 可用空的select来阻塞main函数

// 可设置超时

func main() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}



//有buffer的channel --> 先“放”后“取”
//无buffer的channel--> 先“取”放“”






var a string 
var c = make(chan int,10) //有buffer的channel --> 先“放”后“取”
func f(){
	a = "saf"
	c <- 0

}

func main(){ //有buffer的channel --> 先“放”后“取”
	go f()
	<-c 
	print(a)
}


func f(){
	a = "asdf"
	<- c 
}
//无buffer的channel--> 先“取”放“”
func main(){
	go f()
	c <- 0 
	print(a)
}

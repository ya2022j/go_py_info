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

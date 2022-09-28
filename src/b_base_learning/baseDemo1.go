// baseDemo1.go
package main

import (
	"fmt"     //提供格式化输出
	"strconv" //提供字符串与基本类型的转换
)

// 变量的集体定义方式，如果开头字母大写就是全部变量
var (
	packageName   string = "main"
	packageImport string = "fmt"
	isMain        bool   = true
)

// 常量只能是布尔，数字和字符串
const is_main_package = true
const count_of_import = 1
const name_of_package = "main"

// 常量也可以定义在一起
const (
	chinese_name = "中国"
	english_name = "China"
	age          = 60
)

// 定义结构体，javaBean
type Person struct {
	id   int
	name string
}

// package main下的main函数必不可少，两者齐备才能生成可执行文件
func main() {
	fmt.Println("Hello World!")
	//基本类型 var+变量名+变量类型+等号+赋初始值
	var intA int = 1
	var boolB bool = false
	var float32C float32 = 112.123123
	var float64D float64 = -234234234.2342355342
	//变量定义了必须用，不用的话会编译不过
	fmt.Println(intA)
	fmt.Println(boolB)
	fmt.Println(float32C)
	fmt.Println(float64D)
	//数字类型 go的数字类型好复杂，挺多的
	var uintA uint = 1
	var uint8A uint8 = 255
	var uint16A uint16 = 65535
	var uint32A uint32 = 4294967295
	var uint64A uint64 = 18446744073709551615
	var int8A int8 = -128
	var int16A int16 = -32768
	var int32A int32 = -2147483648
	var int64A int64 = -9223372036854775808
	var runeA rune = -2147483648
	fmt.Println(uintA)
	fmt.Println(uint8A)
	fmt.Println(uint16A)
	fmt.Println(uint32A)
	fmt.Println(uint64A)
	fmt.Println(int8A)
	fmt.Println(int16A)
	fmt.Println(int32A)
	fmt.Println(int64A)
	fmt.Println(runeA)

	//var uintptrA uintptr = *int64A
	//fmt.Println(uintptrA)
	//默认值 只定义变量，不赋值的话，会被赋默认值
	var dva *int             //int型指针
	var dvb []int            //int型的切片
	var dvc map[string]int   //map<string, int> map
	var dve chan int         //定义了一个channel，里边会流动int型的数据
	var dvf func(string) int //定义一个函数，输入string，输出int
	var dvg error            //定义一个error的接口
	fmt.Println("开始默认初始值测试")
	fmt.Println(dva)
	fmt.Println(dvb)
	fmt.Println(dvc)
	fmt.Println(dve)
	fmt.Println(dvf)
	fmt.Println(dvg)

	//变量声明和赋值 两个变量可以一起声明赋值
	var p1, p2 string = "abc", "efg"
	fmt.Println(p1, p2)
	//不给初始值的时候就会有默认值 string的默认值不是nil而是空
	var p3 string
	fmt.Println("p3=" + p3)
	//不声明变量类型但是赋值了的话，编译器会自行推断变量类型
	var p4 = 234
	var p5 = true
	fmt.Println(p4, p5)
	//声明赋值一步到位 var可以声明变量，：=也可以声明变量，但两者不能同时使用，同时使用就相当于把一个变量声明了两遍，这就涉及到谁覆盖谁的问题了
	//为了保证不出现二义性，go语言编译阶段就会报错。为了避免这种二义性，go语言也不支持隐形类型转换
	//省略var的声明方式如下
	p6 := true
	fmt.Println(p6)

	fmt.Println(packageName)
	//下划线，空变量占位符，如果某个变量不得不取出来，但又不用的话，就可以用占位符取出来。这样后面的程序是不能引用它的值的
	p7, _, p9 := true, "abc", 123

	fmt.Println(p7, p9)

	var rA, rB = func1()
	fmt.Println(rA, rB)

	fmt.Println(is_main_package, name_of_package, count_of_import)

	//is_main_package = false 编译错误

	fmt.Println(chinese_name, english_name, age)

	//fmt.Println(iota) iota只能在常量定义时用，就是个常量计数器

	const (
		const_a = "sss"
		const_b = iota
		const_c = 123
		const_d = true
		const_e
		const_f
		const_g = iota
	)
	//sss 1 123 true true true 6 充分显示iota就是个常量个数计数器
	fmt.Println(const_a, const_b, const_c, const_d, const_e, const_f, const_g)
	var aaa = 11
	if aaa < 5 {
		fmt.Println("go here aaa<5")
	} else if aaa == 5 {
		fmt.Println("go here aaa=5")
	} else {
		fmt.Println("go here aaa>5")
	}
	//case里面自带break,匹配到case之后就不会执行下面的了，如果想让它继续把下面的都跑完，就家fallthrough
	switch aaa {
	case 1:
		fmt.Println("go here aaa=1")
	case 2:
		fmt.Println("go here aaa=2")
	case 3:
		fmt.Println("go here aaa=3")
	case 4:
		fmt.Println("go here aaa=4")
	case 5:
		fmt.Println("go here aaa=5")
	case 6:
		fmt.Println("go here aaa=6")
	case 7:
		fmt.Println("go here aaa=7")
	case 8:
		fmt.Println("go here aaa=8")
	case 9:
		fmt.Println("go here aaa=9")
	case 10:
		fmt.Println("go here aaa=10")
	case 11:
		fmt.Println("go here aaa=11")
		fallthrough
	case 12:
		fmt.Println("go here aaa=12")
		fallthrough
	default:
		fmt.Println("go here aaa=default")
	}
	//select就是一个while(true),直到某个生产者或者消费者匹配到执行，否则就会阻塞。  谨慎使用，因为会阻塞
	// var bbb = 4
	// select {
	// case 1:
	// 	fmt.Println("select case 1")
	// case 2:
	// 	fmt.Println("select case 1")
	// case 3:
	// 	fmt.Println("select case 1")
	// case 4:
	// 	fmt.Println("select case 1")
	// case 5:
	// 	fmt.Println("select case 1")
	// default:
	// 	fmt.Println("select case default")

	// }

	//go语言里只有for一种循环方式，没有while和do while
	for i := 0; i < 10; i++ {
		fmt.Println("for i=" + strconv.Itoa(i))
	}

	j := 10
	for j > 7 {
		j--
		fmt.Println("for j=" + strconv.Itoa(j))
	}

	for {
		fmt.Println("for true")
		break
	}
	//函数，可以返回同时返回多个值
	var retA, retB = func2(13, "func2++++")
	fmt.Println("func2 return=" + retA + strconv.Itoa(retB))
	//函数还可以作为参数传入函数中， 函数就是个接收输入值，给出输出值的工具，不读取外部变量
	funcParam1 := func(intA int) string { return "int to string=" + strconv.Itoa(intA) }
	func2(11, funcParam1(3333))

	//传入函数实现回调
	func3(44444, func(a int) string {
		return "convert int to " + strconv.Itoa(a)
	})

	//闭包，就是函数里面内嵌函数
	fmt.Println(closePackageFunc(99))

	//数组
	var array1 = [3]int{2, 1, 9}
	fmt.Println(array1)
	fmt.Println(array1[1])

	//指针 定义
	var intAaPoint *int
	var stringBbPoint *string

	var aaaa int = 2555
	var bbbb string = "bbbb2555"
	fmt.Println(aaaa)
	fmt.Println(bbbb)
	//指针 定义
	intAaPoint = &aaaa
	stringBbPoint = &bbbb
	//指针 内存地址值
	fmt.Println(intAaPoint)
	fmt.Println(stringBbPoint)
	//指针 指针对应的真实值
	fmt.Println(*intAaPoint)
	fmt.Println(*stringBbPoint)

	//结构体，就是javabean，脱胎于C语言
	var aPerson Person = Person{1, "zhangsan"}
	fmt.Println(aPerson)
	fmt.Println(aPerson.name)

	//语言切片 就是创建list
	//新建list
	var list1 = make([]int, 0)
	fmt.Println(list1)
	//获取list的size
	fmt.Println("list1 length=" + strconv.Itoa(len(list1)))
	//向list中添加元素
	list1 = append(list1, 1, 0)
	list1 = append(list1, 9, 6, 3)
	fmt.Println(list1)
	//读取list中的元素
	fmt.Println(list1[1])
	fmt.Println(list1[1:3])
	fmt.Println(list1[:3])
	fmt.Println(list1[2:])
	//删除list中的元素?这个go里边根本没这功能，只能是把删除元素之外的元素拼接起来塞到原来的list里来实现

	//range相当于java中的迭代器
	list2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(list2)
	for indexOfList2, valueOfList2 := range list2 {
		fmt.Printf("index==>value::%d => %d\n", indexOfList2, valueOfList2)
		//fmt.Println("index=%d, value=%d", indexOfList2, valueOfList2)
	}

	//map
	//map的创建
	var map1 map[int]string = make(map[int]string)
	var map2 map[string]int = make(map[string]int)
	//map的添加元素
	map1[0] = "zero"
	map1[1] = "one"
	map1[2] = "two"
	map2["three"] = 3
	map2["four"] = 4
	map2["five"] = 5
	//map的删除元素
	delete(map1, 1)
	delete(map2, "four")
	//map的遍历元素
	fmt.Println(map1)
	fmt.Println(map2)
	//对map来讲，range遍历的是key
	for val := range map1 {
		fmt.Println(val, map1[val])
	}
	//map的是否存在元素
	var anyValue, isExists = map2["five"]
	fmt.Println(anyValue, isExists)

	//递归，就是方法里自己调用自己

	//类型转换
	var aaaaa int = 1
	//var bbbbb string = "123"
	fmt.Println(string(aaaaa))
	//fmt.Println(int(bbbbb) + 2)

	//接口
	var aPhone = new(HuaWeiPhone)
	fmt.Println(aPhone.speak())
	//aPhone.listen("huaweiphone is listenning") 报编译错误，编译器不是找的接口下的方法，而是找的接口实现类下的方法
	var bPhone = new(XiaoMiPhone)
	//bPhonePoint := &bPhone
	fmt.Println(bPhone.speak())

	//Error接口 一个特殊的接口，只有Error() string这一个方法

}

// 接口的例子 #############
type Phone interface {
	speak() string
	listen(string)
}
type HuaWeiPhone struct {
}

type XiaoMiPhone struct {
}

func (a HuaWeiPhone) speak() string {
	return "this is huawei phone speaking"
}

func (a *XiaoMiPhone) speak() string {
	return "this is xiaomi phone speaking"
}

//接口的例子 #############

// 闭包，函数里边套函数
func closePackageFunc(a int) string {
	var mm = a
	var intMM int = 454
	cp1 := func(i int, j int) string {
		return strconv.Itoa(i) + "___" + strconv.Itoa(j) + "---" + strconv.Itoa(mm)
	}
	return cp1(6, intMM)
}

// go中函数也是一个对象，所以，它可以被定义成一个类型，可以作为参数传递，可以在函数内被引用
// 这就是把函数定义成了一个类型，这样它就可以被写进函数的参数里了
type funcIntToString func(int) string

// 一个函数，没有返回值的
func func3(a int, b funcIntToString) {
	fmt.Println("a convert to string =" + b(a))
}

// 一个函数，有多个返回值
func func2(a int, b string) (string, int) {
	fmt.Println(a + 1)
	fmt.Println(b)
	return b, a
}

// 一个函数，没有入参，只有多个返回值
func func1() (int, string) {
	return 1, "2324rvsdf"
}

package main
/*
*@Description: Go语言基本语法的学习，基本涵盖go语言所有语法，入门级
*@Author: lvzhonghou
*@Date: 2014.11.12
*/
import "fmt"
import "math"
import "errors"
import "io"
import "os"

//错误处理-recover
func g(i int) {
	if i > 1 {
		fmt.Println("Panic!")
		panic(fmt.Sprintf("%v", i))
	}
	
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("Calling g with ", i)
		g(i)
		fmt.Println("Returned normally from g.")
	}
	
}

//错误处理-Panic
func init1() {
	user := os.Getenv("USER")

	if user == "" {
		panic("no value for $USER")
	}
	
}

//错误处理-Defer
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return 
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return 
}

/*引用defer的错误处理*/
func CopyFile_v1(dstName, srcName string) (written int64, err error){
	src, err := os.Open(srcName)
	if err != nil {
		return 
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return 
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

//错误处理
/*自定义的出错结构*/
type myError struct {
	arg int
	errMsg string
}

/*实现Error接口*/
func (e *myError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.errMsg)
}

/*两种出错*/
func error_test(arg int) (int, error){
	if arg < 0 {
		return -1, errors.New("Bad Arguments - negtive!")
	} else if arg > 256 {
		return -1, &myError{arg, "Bad Arguments - too large!"}
	}
	return arg * arg, nil
}

//接口和多态
type shape interface {
	area() float64
	perimeter() float64  
}

/*长方形*/
type rect_v1 struct {
	width, height float64
}

func (r *rect_v1) area() float64 {
	return r.width * r.height
}

func (r *rect_v1) perimeter() float64{
	return 2 * (r.width + r.height)	
}

/*圆形*/
type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius	
}

/*-------接口的使用-------*/
func interface_test() {
	r := rect_v1 {width: 2.9, height: 4.8}
	c := circle {radius: 4.3}

	s := []shape{&r, &c}

	for _, sh := range s {
		fmt.Println(sh)
		fmt.Println(sh.area())
		fmt.Println(sh.perimeter())
		
	}
}

//结构体方法
type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int{  //求周长
	return 2 * (r.width + r.height)
	
}

//结构体
type Person struct {
	name string
	age int
	email string
}

//递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

//返回匿名函数
func nextNum() func() int {
	i, j := 1, 1
	return func() int {
		var tmp = i + j
		i, j = j, tmp
		return tmp
	}
}

//函数
func max(a int, b int) int {
	if a > b {
		return b
	}
	return b
}

//返回多个返回值
func multi_ret(key string)(int, bool) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	var err bool
	var val int

	val, err = m[key]

	return val, err	
}

//函数不定参数
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	//错误处理-Panic
	// init1()

	fmt.Println("hello world")

	fmt.Printf("%t\n" , 1==2)
	fmt.Printf("二进制：%b\n", 255)
	fmt.Printf("浮点数：%f\n", math.Pi)

	//变量
	var x = 100
	var str string = "hello world"
	var i, j, k = 1, 2, 3
	fmt.Println(x)
	fmt.Println(str)
	fmt.Println(i)
	fmt.Println(j) 
	fmt.Println(k)

	//常量
	const s string = "hello world"
	const pi float32 = 3.1415926
	fmt.Println(s)
	fmt.Println(pi)

	//数组
	var a [5]int
	fmt.Println(a)

	a[1] = 10
	a[3] = 30
	fmt.Println("assign:", a)

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("init:", b)

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("2d: ", c)

	//数组的切片操作
	a1 := [5]int{1, 2, 3, 4, 5}

	b1 := a1[2:4]
	b2 := b1
	fmt.Println(b2)

	b1 = a1[:4]
	fmt.Println(b1)

	b1 = a1[2:]
	fmt.Println(b1)

	//分支循环语句
	if  x % 2 == 0 {
		fmt.Println("x is even")
	}

	//switch语句
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5, 6:
		fmt.Println("four, five, six")
	default:
		fmt.Println("invalid value!")
	}

	//for语句
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i1 := 1
	for i1 < 10{
		fmt.Println(i1)
		i1++
	}

	i2 := 1
	for  {
		if i2 > 10 {
			break
		}
		i2++
	}

	//map
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	fmt.Println(m)
	fmt.Println(len(m))

	v := m["two"]
	fmt.Println(v)

	delete(m, "two")
	fmt.Println(m)

	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(m1)

	for key, val := range m1 {
		fmt.Printf("%s => %d \n", key, val)
	}

	//指针
	var i3 int = 1
	var pInt *int = &i3

	fmt.Printf("i3=%d\tpInt=%p\t*pInt=%d\n", i3, pInt, *pInt)

	//内存分配
	var p *[]int = new([]int)
	fmt.Println(p)

	var v1 []int = make([]int, 10)
	fmt.Println(v1)

	fmt.Println(max(4, 5))
	fmt.Println(multi_ret("four"))

	//变长参数
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	nextNumFunc := (nextNum)
	for i := 0; i < 10; i++ {
			fmt.Println(nextNumFunc())
	}

	fmt.Println(fact(7))

	//结构体
	person := Person{"Tom", 30, "tom@qq.com"}
	person = Person{name:"Tom", age:30, email:"tom@gmail.com"}

	fmt.Println(person)

	pPerson := &person

	fmt.Println(pPerson)

	pPerson.age = 40
	person.name = "Jerry"
	fmt.Println(person)

	//接口测试
	r := rect{width: 10, height: 15}

	fmt.Println("面积： ", r.area())
	fmt.Println("周长： ", r.perimeter())

	rp := &r
	fmt.Println("面积：", rp.area())
	fmt.Println("周长：", rp.perimeter())	

	interface_test();

	//error测试
	for _, i := range []int{-1, 4, 1000} {
		if r, e := error_test(i); e != nil {
			fmt.Println("failed:", e)
		} else {
			fmt.Println("success:", r)
		}
		
	}

	//recover测试
	f()
	fmt.Println("Returned normally from f.")

}
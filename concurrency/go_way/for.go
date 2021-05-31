package main

import "fmt"

// 要求循环 15 次然后使用 fmt 包来打印计数器的值
// var count int = 0
// var i int = 1
// for i:=0; i<15; i++{
//     count += 1
//     fmt.Printf("count = %d\n", count)
// }

// 使用 goto 语句重写循环，要求不能使用 for 关键字
// LOOP:
// fmt.Printf("The counter is at %d\n", i)
// i ++
// if i <= 15{
//     goto LOOP
// }

// 使用 2 层嵌套 for 循环
// for i:=1;i<=25;i++{
//     for j:=1;j<=i;j++{
//         fmt.Print("G")
//     }
//     fmt.Println()
// }

// 使用一层 for 循环以及字符串截断(一)
// var G_str string = "GGGGGGGGGGGGGGGGGGGGGGGGG"
// for i:=1; i<25; i++{
//     fmt.Println(G_str[0:i])
// }

// 使用一层 for 循环以及字符串截断(二)
// str := "G"
// for i:=0; i<25;i++{
//     fmt.Println(str)
//     str += "G"
// }

// 写一个从 1 打印到 100 的程序，但是每当遇到 3 的倍数时，不打印相应的数字，但打印一次 "Fizz"。
// 遇到 5 的倍数时，打印 Buzz 而不是相应的数字。对于同时为 3 和 5 的倍数的数，打印 FizzBuzz
// for i:=1; i<=100; i++{
//     switch true{
//         case i%3==0:
//         fmt.Println("Fizz")
//         case i%5==0:
//         fmt.Println("Buzz")
//         case i%15==0:
//         fmt.Println("FizzBuzz")
//         default:
//         fmt.Println(i)
//     }
// }

//使用 * 符号打印宽为 20，高为 10 的矩形
// for i:=1; i<=10; i++{
//     for j:=1; j<=20; j++{
//         fmt.Print("*")
//     }
//     fmt.Println()
// }

// for 结构的第二种形式是没有头部的条件判断迭代（类似其它语言中的 while 循环）
// var i = 5
// for i > 0 {
//     i = i -1
//     fmt.Println(i)
// }

// 这是 Go 特有的一种的迭代结构，它可以迭代任何一个集合,要注意的是，循环始终为集
// 合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值，除非是指针类型
// str := "Go is a beautiful language!"
// str2 := "Chinese: 日本語にほんご"
// fmt.Printf("The length srt is %d.", len(str2))
// for pos, char := range str2 {
//     fmt.Printf("The srt2 position %d character is %c\n", pos, char)
// }

// 程序的输出结果是什么？
// for i := 0; i < 5; i++ {
// var v int
// fmt.Printf("%d ", v)
// }

// 打印1-7的奇数
//     for i := 0; i<7 ; i++ {
//     if i%2 == 0 { continue }
//     fmt.Println("Odd:", i)
// }
// }

// 下面的函数将不会被编译，为什么呢？大家可以试着纠正过来。
// type *Stack struct{}
// func (st *Stack) Pop() int {
//     v := 0
//     for ix := len(st) - 1; ix >= 0; ix-- {
//         if v = st[ix]; v != 0 {
//             st[ix] = 0
//             return v
//         }
//     }

// 编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值
// func NoReturnName(i, j int) (int, int, int) {
// 	return i + j, i - j, i * j
// }

// func ReturnName(i, j int) (sum int, diff int, mult int) {
// 	sum, diff, mult = i+j, i-j, i*j
// 	return
// }

// func main() {
// 	sum, diff, mult := NoReturnName(3, 4)
// 	fmt.Println("Sum:", sum, "| Diff=", diff, "| Mult=", mult)
// 	sum, diff, mult = ReturnName(3, 4)
// 	fmt.Println("Sum:", sum, "| Diff=", diff, "| Mult=", mult)
// }

// 编写一个名字为 MySqrt 的函数，计算一个 float64 类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。
// 编写两个版本，一个是非命名返回值，一个是命名返回值
// func MySqrt(f float64) (res float64, err error) {
// 	if f < 0 {
// 		res = float64(math.NaN())
// 		err = errors.New("I won't be able to do a sqrt of negative number")
// 	} else {
// 		res = math.Sqrt(f)
// 	}
// 	return
// }

// func MySqrtN(f float64) (float64, error) {
// 	if f < 0 {
// 		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
// 	}
// 	return math.Sqrt(f), nil
// }

// func main() {
// 	fmt.Print("First example with -1: ")
// 	res, err := MySqrt(-1)
// 	if err != nil {
// 		fmt.Println("Error! Return values are: ", res, err)
// 	} else {
// 		fmt.Println("It's ok! Return values are: ", res, err)
// 	}

// 	fmt.Println(MySqrtN(5))
// }

//空白符用来匹配一些不需要的值，然后丢弃掉
// func main() {
// 	i1, _, f1 := ThreeValues()
// 	fmt.Printf("The int: %d, the float: %f \n", i1, f1)
// }

// func ThreeValues() (int, int, float64) {
// 	return 1, 2, 3.5
// }

//传递指针给函数不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不再需要使用 return 返回
// func Multiply(a, b int, reply *int) {
// 	*reply = a * b
// }

// func main() {
// 	n := 0
// 	reply := &n
// 	Multiply(3, 4, reply)
// 	fmt.Print(*reply)
// }

//参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数调用变参函数
// func min(s ...int) int {
// 	if len(s) == 0 {
// 		return 0
// 	}
// 	min := s[0]
// 	for _, v := range s {
// 		if v < min {
// 			min = v
// 		}
// 	}
// 	return min
// }

// func main() {
// 	x := min(1, 2, 4, 6, 8, 0, 4, 2)
// 	fmt.Printf("The minimum is: %d\n", x)

// 	slice := []int{7, 9, 3, 5, 1}
// 	y := min(slice...)
// 	fmt.Printf("The minimum is: %d\n", y)
// }

//如果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}，这样就可以接受任何类型的参数。
//该方案不仅可以用于长度未知的参数，还可以用于任何不确定类型的参数
// func typecheck(value ...interface{}) {
// 	for _, val := range(value){
// 		switch v := val.(type){
// 		case int: fmt.Print("")
// 		case string:...
// 		case float64:...
// 		...
// 		}
// 	}
// }

//判断变量的类型
// func main() {
// 	var i = 123
// 	fmt.Println(reflect.TypeOf(i))
// }

//递归函数
// func fibonacci(n int) (res int) {
// 	if n <= 1 {
// 		res = 1
// 	} else {
// 		res = fibonacci(n-2) + fibonacci(n-1)
// 	}
// 	return
// }

// func main() {
// 	result := 0
// 	for i := 0; i <= 10; i++ {
// 		result = fibonacci(i)
// 		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
// 	}
// }

//使用递归函数从 10 打印到 1
// func main() {
// 	printrec(1)
// }

// func printrec(i int) {
// 	if i > 10 {
// 		return
// 	}
// 	printrec(i + 1)
// 	fmt.Printf("%d ", i)
// }

//实现一个输出前 30 个整数的阶乘的程序
//特别注意的是，使用 int 类型最多只能计算到 12 的阶乘，因为一般情况下 int 类型的大小为 32 位，继续计算会导致溢出错误
func main(){

}

func jiecheng(n int) (res int){
	if n
}
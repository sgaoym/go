package main

import (
	"fmt"
)

func test() {
	a := 21
	b := 10
	var c int

	c = a + b
	fmt.Println(c)

	c = a - b
	fmt.Println(c)

	c = a/b;
	fmt.Println(c)
	c = a*b;
	fmt.Println(c)
	c = a%b;
	fmt.Println(c)
	a++
	fmt.Println(a)
	a = 21
	a--
	fmt.Println(a)
}

func test2(){
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	ptr  = &a
	fmt.Println(a)
	fmt.Println(*ptr)
}

func test3(){
	var a int = 10
	if a < 20 {
		fmt.Println("a < 20")
	} else {
		fmt.Println("a >= 20")
	}

	a = -1;
	switch a {
	case -1:
		fmt.Println("case -1")
	default:
		fmt.Println("switch default.")
	}

	fmt.Println("the value of a is:",a);

	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println("type is ",i)
	case int:
		fmt.Println("type is ",i)
	case float64:
		fmt.Println("type is ",i)
	case func(int) float64:
		fmt.Println("type is ",i)
	case bool,string:
		fmt.Println("type is ",i)
	default:
		fmt.Println("unknown")
	}

}

func test4(){
	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		fmt.Println("received", i1, "from c1\n")
	case c2 <- i2:
		fmt.Println("sent", i2, "to c2\n")
	case i3,ok := (<-c3):
		if ok{
			fmt.Println("received", i3, "from c3\n")
		} else {
			fmt.Println("c3 is closed\n")
		}
	default:
		fmt.Println("no communication\n")
	}
}

func test5(){
	a := 10
	for a < 20 {
		fmt.Println("a:", a)
		a++
		if a > 15 {
			break;
		}
	}

	var i, j int
	for i=2; i < 100; i++ {
		for j = 2; j <= (i/j); j++ {
			if(i %j == 0) {
				break;
			}
		}

		if(j > i/j){
			fmt.Println(i, "是素数")
		}
	}

	a = 10
	
	LOOP: for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}

		fmt.Println("a == ", a)
		a++
	}
}


func max(n1, n2 int ) int {
	var r int

	if n1 > n2 {
		r = n1
	} else {
		r = n2
	}

	return r;
}

func swap(x, y string) (string, string) {
	return y,x
}

func f_array() {
	var balance = [5]float32{1000,2,3,4,50.5}
	var balance2 = [...]float32{1000,2,3,4,7,50.5}
	balance2[1] = 0
	balance[4] = 50.5

	var n [10]int
	var i,j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

func f_ptr(){
	var a int  = 20
	var p *int
	p = &a

	fmt.Printf("a'address:%x\n", &a)
	fmt.Printf("p store address:%x\n", p)
	fmt.Printf("p store value:%d\n", *p)

	var ptr *int
	fmt.Printf("ptr is %x\n", ptr)

	if ptr == nil {
		fmt.Printf("null pointer\n")
	} else {
		fmt.Printf("not null pointer\n")
	}
}

const MAX int = 3

func f_ptrArray (){
	a := []int {10,100,200}
	var i int
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d ", i, a[i])
	}

	var ptr[MAX]*int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

type Books struct {
	title string
	author string
	subject string
	book_id int
}
func f_struct(){
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程",6495497})
	fmt.Println(Books{title:"Go 语言", author:"www.runoob.com",subject:"Go 语言教程",book_id:649540})
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}

func f_slice(){
	var ns = make([]int,3,5)
	numbers := []int {0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	fmt.Println("numbers == ", numbers);

	fmt.Println("numbers[1:4] == ", numbers[:4])

	printSlice(ns);

	var nu []int
	printSlice(nu)

	nu = append(nu,0)
	printSlice(nu)

	nu = append(nu,1)
	printSlice(nu)

	nu = append(nu, 2,3,4)
	printSlice(nu)

	nu_new := make([]int, len(nu), cap(nu)*2)
	copy(nu_new, nu)
	printSlice(nu)
}

func printSlice(x []int){
	fmt.Printf("len=%d  cap=%d slice=%v\n", len(x), cap(x),x)
}

func f_range(){
	nums := []int {2,3,4}
	sum := 0
	for _,num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	kvs := map[string]string{"a":"apple","b", "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k,v)
	}

}

func main() {
	f_range();
}
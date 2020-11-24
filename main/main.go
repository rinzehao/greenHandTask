package main

import (
	"fmt"
)

const Pi float32 = 3.1415926
const MaxThread = 10

var isActive bool

type person struct {
	name string
	age  int
}

func main() {

	type cat struct {
		name      string
		age       int
		sex       string
		character string
	}

	var P person
	P.name = "Hayashi"
	P.age = 22
	fmt.Printf("the person's name is %s \n", P.name)
	P2 := person{"Daniel", 24}
	P3 := person{age: 27, name: "Jack"}
	fmt.Printf("%s %s \n", P2.name, P3.name)

	var name string = "粑粑"
	fmt.Println("请输入你的名字:")
	//fmt.Scanln(&name)
	fmt.Println("你好", name)

	var vname1, vname2, vname3 = 100, 100.20, true
	vname4, vname5, vname6 := 21, 45.6565, true
	fmt.Println(vname1, vname2, vname3, vname4, vname5, vname6)

	isActive = true

	var ar = [7]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	var a, b, c []byte
	a = ar[2:5]
	b = ar[3:7]
	c = b[1:3]
	fmt.Printf("%d %d %d", a, b, c)

	var numbers map[string]int
	numbers = make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Println()
	fmt.Println("第三个数字是", numbers["three"])
	numbers["three"] = 33
	fmt.Println("修改后的第三个数是", numbers["three"])
	fmt.Println(len(numbers))

	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 5.1, "C#": 2}
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating assciated with C# in the map ")
	}

	delete(rating, "C")

	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "sqwadika"
	fmt.Println("m[hello] now is...", m["Hello"])

	var s1 []int
	if s1 == nil {
		fmt.Printf("s1 is nil --> %#v \n", s1)
	}
	s2 := make([]int, 3)
	if s2 == nil {
		fmt.Printf("s2 is nil--> %#v \n", s2)
	} else {
		fmt.Printf("s2 is not nil--> %#v \n", s2)
	}

	if x := numbers["three"]; x > 10 {
		fmt.Printf("x is bigger than 10")
	} else {
		fmt.Printf("x is less than 10")
	}

	i := 0
Here:
	fmt.Printf("i now =%d \n", i)
	i++
	if i <= 10 {
		goto Here
	}
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	for sum < 1000 {
		sum += sum
	}

	for index := 10; index > 0; index-- {
		if index == 5 {
			break
		}
		fmt.Println(index)
	}

	for k, v := range numbers {
		fmt.Println("map's key :", k)
		fmt.Println("map 's val", v)
	}
	a_array := [7]int{1, 2, 3, 4, 5, 6, 7}
	for index, value := range a_array {
		fmt.Println(index, value)
	}
	var PriceOfBooks map[string]float32
	PriceOfBooks = make(map[string]float32)
	PriceOfBooks["红楼梦"] = 12.5
	PriceOfBooks["青楼梦"] = 122.5
	PriceOfBooks["做美梦"] = 28.77
	for k, v := range PriceOfBooks {
		if price := v; price > 100 {
			fmt.Println("别做梦")
		} else {
			fmt.Println(k, v)
		}
	}
	var flag = 10
	switch flag {
	case 1:
		fmt.Println("mmmiya")
	case 2:
		fmt.Println("mmmiya")
	case 10:
		fmt.Println("shabixg")
		fallthrough
	default:
		fmt.Println("发生甚么事了？")
	}
	_max := max(numbers["three"], numbers["one"])
	fmt.Println("大意了啊，最大值是", _max)

	var s []int = a_array[:]
	findMax, findMin := findMaxAndMin(s)
	fmt.Print("马老师，最小值为", findMin)
	fmt.Println("  最大值为", findMax)
	y, z := 4, 10
	fmt.Printf("max(%d,%d)=%d\n", y, z, max(y, z))

	p_older, p_gap := Older(P2, P3)

	fmt.Printf("the elder man is %s ,and the gap of age is %d", p_older.name, p_gap)

	Stu := Student{person{name: "年枣糕", age: 21}, "20170070117"}
	fmt.Println("His name is ", Stu.name)
	fmt.Println("His age is ", Stu.age)
	fmt.Println("His Sid is ", Stu.sid)
	Stu.age = 69
	fmt.Println("年枣糕变成69岁的年枣糕", Stu.age)

	type Employee struct {
		person
		age   int32
		phone int64
	}

	Bob := Employee{person{name: "马宝国", age: 69}, 5, 13246766666}
	fmt.Println("Bob‘s work age is ", Bob.age)
	fmt.Println("Bob's age is", Bob.person.age)

}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findMaxAndMin(a []int) (max int, min int) {
	fmt.Println("找最值")
	max = a[0]
	min = a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

type Student struct {
	person
	sid string
}

//package main
//
//import "fmt"
//
//func main() {
//var symbol int
//symbol = 's'
//fmt.Println(symbol)
//
//fmt.Println(string(symbol))
//
//var (
//	a = 2
//	b = "string"
//	c = true
//)
//fmt.Println(a, b, c)

//fmt.Println(10 % 3)

//var name string
//var age int
//
//fmt.Println("name")
//fmt.Scan(&name)
//
//fmt.Println("age")
//fmt.Scan(&age)
//
//fmt.Println(name, age)

//var a int
//fmt.Scan(&a)
//fmt.Println((a * 2) + 100)

//var a int
//var b int
//
//fmt.Scan(&a) // считаем переменную 'a' с консоли
//fmt.Scan(&b) // считаем переменную 'b' с консоли
//
//fmt.Println((a * a) + (b * b))

//var a = 2053
//fmt.Println((a % 100) / 10)

//var d int
//fmt.Scan(&d) //93
//
//var h, m int
//m = (d * 2) % 60 //6 min
//h = (d * 2) / 60 //3 h
//fmt.Println("It is", h, "hours", m, "minutes.")
//const (
//	a = 5
//	b
//	c
//	d = iota
//	e
//)
//
//fmt.Println(a, b, c, d, e)
//
//const (
//	a = iota + 1
//	_
//	b
//	c
//	d = c + 2
//	t
//	i
//	i2 = iota + 2
//)

//i := 3
//
//switch i {
//case 1:
//	fmt.Println("1")
//case 2:
//	fmt.Println("2")
//case 3:
//	fmt.Println("3")
//	fallthrough
//case 4:
//	fmt.Println("4")
//case 5:
//	fmt.Println("5")
//}

//var a int
//fmt.Scan(&a)
//
//switch {
//case a >= 1 && a <= 10:
//	fmt.Println("1 - 10")
//case a >= 11 && a <= 20:
//	fmt.Println("11 - 20")
//default:
//	fmt.Println("default")
//}

//var a int
//fmt.Scan(&a)
//
//switch {
//case a < 0:
//	fmt.Println("Число отрицательное")
//case a > 0:
//	fmt.Println("Число положительное")
//case a == 0:
//	fmt.Println("Ноль")
//}

//var number int
//fmt.Scan(&number)
//
//firstDigit := number / 100
//secondDigit := (number / 10) % 10
//thirdDigit := number % 10
//
//switch {
//case firstDigit != secondDigit && secondDigit != thirdDigit && firstDigit != thirdDigit:
//	fmt.Println("YES")
//default:
//	fmt.Println("NO")
//}

//}

//package main
//
//import "fmt"
//
//func main() {
//	var number int
//	fmt.Scan(&number)
//
//	switch {
//	case number == 10000:
//		fmt.Println(number / 10000)
//	case number < 10000 && number >= 1000:
//		fmt.Println(number / 1000)
//	case number < 1000 && number >= 100:
//		fmt.Println(number / 100)
//	case number < 100 && number >= 10:
//		fmt.Println(number / 10)
//	case number < 10:
//		fmt.Println(number)
//
//	}
//}

//package main
//
//import "fmt"
//
//func main() {
//	var n int
//	fmt.Scan(&n)
//	switch {
//	case n < 10:
//		fmt.Println(n)
//	case n < 100:
//		fmt.Println(n / 10)
//	case n < 1000:
//		fmt.Println(n / 100)
//	case n < 10000:
//		fmt.Println(n / 1000)
//	case n < 100000:
//		fmt.Println(n / 10000)
//	}
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	var n int
//	fmt.Scan(&n)
//
//	var (
//		n1, n2, n3, n4, n5, n6 int
//	)
//
//	n1 = n / 100000
//	n2 = (n / 10000) % 10
//	n3 = (n / 1000) % 10
//	n4 = (n / 100) % 10
//	n5 = (n / 10) % 10
//	n6 = n % 10
//
//	var sumLeft = n1 + n2 + n3
//	var sumRight = n4 + n5 + n6
//
//	switch {
//	case sumLeft == sumRight:
//		fmt.Println("YES")
//	default:
//		fmt.Println("NO")
//	}
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	input := 444
//
//	n1 := input%4 == 0
//	fmt.Println(n1)
//
//	n2 := input%100 == 0
//	fmt.Println(n2)
//
//	n3 := input%400 == 0
//	fmt.Println(n3)
//
//	if n1 {
//		if n2 {
//			if n3 {
//				fmt.Println("true 400")
//				return
//			}
//			fmt.Println("false 100")
//			return
//		}
//		fmt.Println("true 4")
//	}
//
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	var a, b, sum int
//	fmt.Scan(&a, &b)
//
//	for i := a; i <= b; i++ {
//		sum += i
//	}
//	fmt.Println(sum)
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	var count int
//	fmt.Scan(&count)
//
//	var output int
//
//	for i := 1; i <= count; i++ {
//		var number int
//		fmt.Scan(&number)
//		if number%8 == 0 {
//			if number > 9 && number < 100 {
//				output += number
//			}
//		}
//	}
//	fmt.Println(output)
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	var number, count, before int
//	count = 0
//	before = 1
//	//var slice []int
//
//	for {
//		fmt.Scan(&number)
//
//		if before == number {
//			count++
//		}
//
//		if before < number {
//			count = 1
//		}
//
//		before = number
//
//		if number == 0 {
//			break
//		}
//	}
//	fmt.Println(count)
//}
//
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var num, max, count int
//
//	//fmt.Println("Enter a sequence of natural numbers (end with 0):")
//
//	for {
//		fmt.Scan(&num)
//		if num == 0 {
//			break
//		}
//		if num > max {
//			max = num
//			count = 1
//		} else if num == max {
//			count++
//		}
//	}
//
//	//fmt.Printf("The number of elements equal to the maximum (%d) is: %d\n", max, count)
//	fmt.Println(count)
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	var n, c, d int
//	fmt.Scan(&n)
//	fmt.Scan(&c)
//	fmt.Scan(&d)
//
//	for i := 1; i <= n; i++ {
//		if i%c == 0 && i%d != 0 {
//			fmt.Println(i)
//			break
//		}
//	}
//
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	var n int
//
//	for {
//		fmt.Scan(&n)
//
//		if n > 100 {
//			break
//		}
//
//		if n < 10 {
//			continue
//		}
//		fmt.Println(n)
//	}
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	var x, p, y, years int
//
//	fmt.Scan(&x) //сумма
//	fmt.Scan(&p) //процент
//	fmt.Scan(&y) //желаемая сумма
//
//	//result = 100 + (100 * 10 / 100)
//
//	for x < y {
//		x = x + (x * p / 100)
//		years++
//	}
//	fmt.Println(years)
//
//}

package main

import (
	"fmt"
)

func main() {

	var number1, number2 int

	fmt.Scan(&number1, &number2)

	var digits1, digits2, result []int

	for number1 > 0 {
		digit := number1 % 10
		digits1 = append([]int{digit}, digits1...)
		number1 /= 10
	}
	for number2 > 0 {
		digit := number2 % 10
		digits2 = append([]int{digit}, digits2...)
		number2 /= 10
	}

	for _, value1 := range digits1 {
		for _, value2 := range digits2 {
			if value1 == value2 {
				result = append(result, value1)
			}
		}
	}

	for i, value := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(value)
	}
	fmt.Println()
}

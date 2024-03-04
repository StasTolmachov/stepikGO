package main

func main() {
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

	const (
		a = iota + 1
		_
		b
		c
		d = c + 2
		t
		i
		i2 = iota + 2
	)

}

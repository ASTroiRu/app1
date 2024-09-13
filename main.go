package main

import "fmt"

func main() {
	l := [11]string{"Ноль", "Один", "Два", "Три", "Четыре", "Пять",
		"Шесть", "Семь", "Восемь", "Девять", "Десять"}
	var n1, n2, n3 int
	fmt.Scan(&n1, &n2, &n3)
	fmt.Println(l[n1])
	fmt.Println(l[n2])
	fmt.Println(l[n3])
}

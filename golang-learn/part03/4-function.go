package main

import "fmt"

//递归,斐波拉契

func factorial(i int )int{
	if i <= 1{
		return 1
	}
	return i *factorial(i-1)
}

func factorial_list(i int) int {

	if i == 0{
		return 0
	}
	if i == 1{
		return 1
	}
	return factorial_list(i-1)+factorial_list(i-2)
}


func main()  {
	var i int =7
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))

	var j int
	var s []int
	for j = 0; j<10; j++{
		s = append(s, factorial_list(j))
	}

	fmt.Printf("Factorial list of %d is %d\n", j, s)
}

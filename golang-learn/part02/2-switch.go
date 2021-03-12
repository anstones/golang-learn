package main

import "fmt"

func main() {
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50,60,70: grade ="C"
	default:
		grade = "D"

	}

	switch grade{
	case "A":
		fmt.Println("优秀!\n")
	case "B", "C":
		fmt.Println("良好\n" )
	case "D":
		fmt.Println("及格\n" )
	default:
		fmt.Println("不及格\n" )

	}
	fmt.Printf("你的等级是 %s\n", grade)
}
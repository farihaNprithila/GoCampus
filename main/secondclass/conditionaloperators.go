package secondclass

import (
	"fmt"
	"strconv"
)

func ConditionalOperators() {
	//traditional loop
	//var i=0 wont work here
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//conditional loop(like while)
	//var j int = 0
	//both will work here
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	//infinite
	var k int = 0
	for {
		fmt.Println("Running...")
		k++
		if k == 5 {
			break
		} // You need a break or other logic to exit
	}

	//loop over collection
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	if x < 5 {
		fmt.Println("x is less than 5")
	} else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is greater than 5")
	}

	if num, err := strconv.Atoi("123758576967605068098096789076"); err == nil {
		fmt.Println("Converted number:", num)
	} else {
		fmt.Println("Conversion failed:", err)
	}

	x = 7

	//Basic switch Statement
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	//Switch with condition
	x = 75
	switch {
	case x < 50:
		fmt.Println("Less than 50")
	case x >= 50 && x <= 100:
		fmt.Println("Between 50 and 100")
	default:
		fmt.Println("More than 100")
	}

	//Switch with short statement
	switch n := 3; n {
	case 1:
		fmt.Println("One")
	case 2, 3, 4:
		fmt.Println("Two to Four")
	default:
		fmt.Println("Other")
	}

}

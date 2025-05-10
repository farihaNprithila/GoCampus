package secondclass

import "fmt"

var y = "hello scope from package"

// exported
func PrintHelloScope() {
	//x := "hello scope"
	//{
	//	x := "new hello scope"
	//	fmt.Println(x)
	//}
	//fmt.Println(x) //will print hello scope here because x := "new hello scope" is outside the scope

	var x = "hello scope"
	{
		x = "new hello scope"
		fmt.Println(x)
	}
	fmt.Println(x) //this will replace the hello scope with hello new scope

	fmt.Println(y)

}

// private
func printByeScope() {}

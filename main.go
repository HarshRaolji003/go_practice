package main

import "fmt"

func main(){
	// var intNum int = 32767
	// intNum = intNum +1
	// fmt.Println(intNum);

	// var floatNum float64 = 12345678.9
	// fmt.Println(floatNum);

	// var floatNum2 float32 = 10.1
	// var intNum2 int32 = 2
	// var result float32 = floatNum2 +float32(intNum2);
	// fmt.Println("Result:", result);

	// var myString string = "Hello World"
	// var myString string = 
	// `Hello World`
	// fmt.Println(myString)

	// fmt.Println(len("$"))

	// var myRune rune = 'a'
	// fmt.Println(myRune) >>> 97

	// var myBoolean bool= false
	// fmt.Println(myBoolean)

	// myVar := "text"
	// fmt.Println(myVar)

	// cannot mutate the value
	//const pi float32 = 3.1515 
	// fmt.Println(pi)

	// Functions and Control Structure

	printValue:= "Function Called"
	printMe(printValue);

	var num int= 11
	var den int = 2
	var result int = intDivision(num,den)
	fmt.Println(result)

}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) int {
	var result int = numerator%denominator
	return result
}

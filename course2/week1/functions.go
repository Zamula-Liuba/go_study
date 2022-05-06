package main
import "fmt"

func PrintHello() {
	fmt.Printf("Hello, world.\n")
}
func multiply(x int, y int) {
	fmt.Print(x * y)
	fmt.Printf("\n")
}
func withReturn(x int) int {
	return x + 1
}
func withTwoReturns(x int) (int, int) {
	return x, x + 1
}
//passed arguments  are copied to parameters
//modifying parametrs has no effect outside the function
//programmer can pass a pointer as an argument
func foo(y *int) {
	*y = *y + 1
}
func arrayFoo(x [3]int) int {
	return x[0]
}
//possible to pass array pointers
func arrayPointers(x *[3]int)int {
	return (*x)[0] + 100
}
//but YOU DONT USE POINTERS IN FUNC IN GO
//USE SLICES INSTESD
//slices contain a pionter to the array
func sliceFoo(sli[] int) int {
	sli[0] = slo[0] + 1
}

func main() {
	PrintHello()

	multiply(3, 5)

	y := withReturn(3) //4
	a, b := withTwoReturns(4)
	foo(&y)
	fmt.Print(a, b, y, "\n")

	c := [3]int{1, 2, 3}
	fmt.Println(arrayFoo(c))
	arrayPointers(&c)
	fmt.Println(arrayPointers(c))

	d := []int{1,2,3}
	sliceFoo(d)
	fmt.Print(d)
},
package main
import "fmt"

func main(){
	// Variable declaration:
	var x int // x is assigned a zero value (0, 0.0, false, nil)
	fmt.Println(x)

	var y float32 = 3.9
	fmt.Println(y)

	// Short declaration operator:
	a := 42 // compiler assigns a type
	fmt.Println(a)

	b, c, d, _, f := 1, "hi", 33, 41, "welcome" // _ is a blank identifier, can be assigned but not used
	fmt.Println(b, c, d, f)

	// This would not work:(declared and not used: e)
	/* 
		b, c, d, e := 1, "hi", 33, 41
		fmt.Println(b, c, d)
	*/

	fmt.Printf("42 in binary is %b\n", a)
	m := 3.56
	fmt.Printf("%v is of type %T\n", m, m)
	n := float32(m)
	fmt.Printf("%v is of type %T", n, n)
}

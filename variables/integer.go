package variables

import (
	"fmt"
)

func Integers() {
	var c int

	var a int = 10
	var b int = 20

	c = a + b

	d := int32(10)
	e := int64(20)

	fmt.Println("int", c)
	fmt.Println("int32", d)
	fmt.Println("int64", e)
}


package main 

import ( 
	"fmt" 
) 
// START OMIT
func b() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	fmt.Printf("x: ptr=%p len=%d cap=%d %x\n", &x[0], len(x), cap(x), x)
	x = append(x, 2) 
	y := append(x, 3)
	z := append(x, 4)
	fmt.Printf("x: ptr=%p len=%d cap=%d %x\n", &x[0], len(x), cap(x), x)
	fmt.Printf("y: ptr=%p len=%d cap=%d %x\n", &y[0], len(y), cap(y), y)
	fmt.Printf("z: ptr=%p len=%d cap=%d %x\n", &z[0], len(z), cap(z), z)
}
//END OMIT
func main() {
	b()
}

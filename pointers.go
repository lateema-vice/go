package main

import "fmt"

//passed in by value
func zeroval(ival int) {
    ival = 0
}

//passed in by reference
//* dereferences this pointer
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    
    i := 1
    fmt.Println("initial: ", i)

    zeroval(i)
    fmt.Println("zeroval: ", i)

    //& gives the memory address
    zeroptr(&i)
    fmt.Println("zeroptr: ", i)    

    fmt.Println("pointer: ", &i)
}

package main

import "fmt"

func main(){
	nums := []int{ 2, 3, 4}
	sum := 0
	
    //Here we dont need the index
    //so let's ignore it
    for _ , num := range nums {
		sum += num
	}
	
	fmt.Println( "sum: ", sum)
    
    /*
    range on arrays and slices
    provides both the index and value
    for each entry.
    */
    for i, num := range nums{
        if num == 3 {
            fmt.Println("index: ", i)
        }
    }

    kvs := map[string]string{"a" : "apple", "b" : "banana"}

    for k, v:= range kvs{
        fmt.Printf("%s -> %s\n", k, v)
    }

}

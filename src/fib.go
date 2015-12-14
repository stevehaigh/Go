package main

import "fmt"
import "os"
import "strconv"

func fib(x int) int {
    
    if x <= 1 {
        return 1
    }   
   
    m, n := 0, 1
    for i := 1; i <= x; i++ {
        m, n = n, m + n   
    }
        
    return m    
}

func main() {
    
    x, _ := strconv.Atoi(os.Args[1])    
    fmt.Println(x, "Fibonacci number is:", fib(x))
}
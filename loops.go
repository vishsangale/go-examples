package main

import "fmt"


func main() {

    i := 1

    for i <= 10 {
        fmt.Println(i)
        i++
    }

    for j := 0; j < 5; j++ {
        fmt.Println(j)
    }


    val := 23


    if val < 25 {
        fmt.Println("Less tha 25")
    } else if val < 20 {
        fmt.Println("Less than 20")
    }


}

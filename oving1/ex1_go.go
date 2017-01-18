// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i int = 0

func add{
    for a := 0; a < 1000000; a++ {
        i += 1
    }
}

func sub{
    for a := 0; a < 10; a++ {
        i += 1
    }
}



func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
                                            // Try doing the exercise both with and without it!
    go add()                      // This spawns someGoroutine() as a goroutine
    go sub()

    for i := 0; i < 50; i++ {
        Println(i)
    }
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(200*time.Millisecond)
    Println("Done: ", i)

}




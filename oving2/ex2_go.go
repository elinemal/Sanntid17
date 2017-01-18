// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt"
    . "runtime"
    //"time"
)

var i int = 0
var lock = make(chan int, 1)
var masterlock = make(chan int, 2)

func add() {
    for a := 0; a < 1000000; a++ {
        <-lock
        i += 1
        lock <- 1
    }
    masterlock <- 1
}

func sub() {
    for a := 0; a < 100000; a++ {
        <- lock
        i -= 1
        lock <- 1
    }
    masterlock <- 1
}



func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
                                            // Try doing the exercise both with and without it!
    lock <- 1
    go add()                      // This spawns someGoroutine() as a goroutine
    go sub()

    <-masterlock
    <-masterlock
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    //time.Sleep(200*time.Millisecond)
    Println("Done: ", i)

}

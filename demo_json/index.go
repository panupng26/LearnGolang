// sync.Mutx is a lib that locks to prevent conflicts from running. In the context of these two types of Go, that is, prevent multiple goroutines from competing for resources.

// when you use >>> want to restrict access to multiple goroutines

// case in pratice >>>

package main

import (
	"fmt"
	"sync"
	"time"
)

var funcsync sync.Mutex

var n = 10

func sumInN(state1 int) {
    funcsync.Lock()
    var x int = n + state1
    fmt.Printf("\n%d\nfinish in sumInN",x)
    time.Sleep(1 * time.Second) 
    funcsync.Unlock()
}

func main() {
    val1 := 20
    val2 := 30
    go sumInN(val1)
    sumInN(val2)
    time.Sleep(3 * time.Second)
    fmt.Println("")
    fmt.Println("Finish")
}

// if you seen this code you'll seen in function sumInN has "funcsync" used sync.Mutex such as call "Lock()" This prevents other goroutines from running the code from now on until the method Unlock is called. So if we call it in main that separates another goroutine, these two goroutines will block each other, it will run (go SumInN(val1)) until it's called. (funcsync.Unlock())The function of SumInN(val2) will continue from the call line. (funcsync.Lock())
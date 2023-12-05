package main

import (
	"fmt"
	"sync"
)

func main() {
    var waitGroup sync.WaitGroup
    waitGroup.Add(2)

    fmt.Println("Defining coroutines")
    go serveApi()
    go serveHttp()

    fmt.Println("Waiting coroutines")
    waitGroup.Wait()
    fmt.Println("Coroutines ended.")
}

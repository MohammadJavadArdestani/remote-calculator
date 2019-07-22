package main

import (
	"simple_calculator_by_go/client"
	. "simple_calculator_by_go/server"
	"sync"
)

var wg sync.WaitGroup
func main()  {
	go RunServer()
	wg.Add(1)
	go client.RunClient()
	wg.Wait()

//server.Calculate("4+6")

}

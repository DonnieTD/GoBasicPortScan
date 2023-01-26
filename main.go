package main

import (
	"fmt"
	"net"
	"sync"
)

const (
	PORT_RANGE = 100
	// PORT_RANGE = 65536
	ADDRESS = "scanme.nmap.org"
	PORT    = "80"
)

func Worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		_, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ADDRESS, PORT))
		if err == nil {
			fmt.Printf("Port %v is open \n", p)
		}
		// WE NEED TO EN THE WAIT GROUP AFTER EACH ITTERATION BECAUSE GOLANG WILL KIND OF SUSPEND
		// THIS LOOP AND KICK OFF WHEN THE CHANNEL FILLS freeing the channel of that poort ????
		wg.Done()
	}
}

func main() {
	// Create a buffered channel of 771 so no more than 771 connections happen at a time
	ports := make(chan int, 100)

	var wg sync.WaitGroup

	// create a go routine for each worker
	for i := 0; i < cap(ports); i++ {
		// each worker will process a channel of 100 ports?
		go Worker(ports, &wg)
	}

	// Run the jobs or block when the worker pool is full . Do this till all the
	for port := 1; port < PORT_RANGE; port++ {
		wg.Add(1)
		ports <- port
	}
	wg.Wait()

	close(ports)
}

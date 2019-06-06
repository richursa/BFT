package main

import "time"

func main() {
	var chans [3]chan bool
	for i := 0; i < 3; i++ {
		chans[i] = make(chan bool)
	}
	for i := 0; i < 3; i++ {
		go lieutenant(i, true, chans)
	}
	time.Sleep(200 * time.Millisecond)
}

package main

import "fmt"

func lieutenant(index int, decision bool, messenger [3]chan bool) {
	trueCount := 0
	falseCount := 0
	d := true
	if index != 1 {
		go func() {
			for i := 0; i < 3; i++ {
				messenger[index] <- decision
			}
		}()
	} else {
		go func() {
			for i := 0; i < 3; i++ {
				messenger[index] <- !decision
			}
		}()
	}

	for i := 0; i < 3; i++ {
		d = <-messenger[i]
		if d == true {
			trueCount++
		} else {
			falseCount++
		}
	}
	if trueCount > falseCount {
		fmt.Println("Attack ! index ", index)
	} else {
		fmt.Println("Retreat ! index ", index)
	}
}

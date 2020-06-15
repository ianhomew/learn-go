package main

import (
	"fmt"
	"time"
)

//const (
//	Nanosecond  Duration = 1
//	Microsecond          = 1000 * Nanosecond
//	Millisecond          = 1000 * Microsecond
//	Second               = 1000 * Millisecond
//	Minute               = 60 * Second
//	Hour                 = 60 * Minute
//)

func main() {

	i := 0
	for {
		if i > 10 {
			break
		}
		//time.Sleep(1 * time.Second) // 1 second
		time.Sleep(100 * time.Millisecond) // 0.1 second
		fmt.Println("i=", i)
		i++
	}

}

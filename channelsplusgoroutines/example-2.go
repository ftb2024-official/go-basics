package channelsplusgoroutines

import (
	"fmt"
	"time"
)

func say(txt string) {
	for i := 0; i < 5; i++ {
		fmt.Println(txt)
		time.Sleep(time.Microsecond * 500)
	}
}

func Example2() {
	go say("Hello")
	say("world")
}

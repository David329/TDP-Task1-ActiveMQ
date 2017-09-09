package main

import "github.com/jjeffery/stomp"
import "fmt"


func main() {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		fmt.Println(err)
	}

	sub, err := conn.Subscribe("SampleQueue", stomp.AckAuto)
	if err != nil {
		fmt.Println(err)
	}
	for {
		msg := <-sub.C
		fmt.Println(string(msg.Body))
	}

	err = sub.Unsubscribe()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Disconnect()
}
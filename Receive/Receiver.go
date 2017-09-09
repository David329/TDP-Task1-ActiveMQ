package main

import "github.com/jjeffery/stomp"
import "fmt"


func main() {
	//Qeue
	conn, err := stomp.Dial("tcp", "localhost:61613")

	//Manejo de Errores
	if err != nil {
		fmt.Println(err)
	}

	//Nombre de la cola DonL
	sub, err := conn.Subscribe("DonL", stomp.AckAuto)
	if err != nil {
		fmt.Println(err)
	}

	//Imprimir mensajes
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
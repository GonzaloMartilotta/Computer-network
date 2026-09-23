package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:65432")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	mensaje := []byte("Hola prueba por TCP")
	_, err = conn.Write(mensaje)
	if err != nil {
		panic(err)
	}

	fmt.Println("Mensaje enviado")
}

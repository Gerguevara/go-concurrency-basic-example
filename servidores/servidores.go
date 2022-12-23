package servidores

import (
	"fmt"
	"net/http"
	"time"
)

// data := <- c // leer de un canal c
// c <- data // enviar data a  canal c
func GetServidores() {

	inicio := time.Now()
	c := make(chan string)

	// fmt.Println(reflect.TypeOf(inicio))
	servidores := []string{
		"http://facebook.com",
		"http://instagram.com",
		"http://twitter.com",
		"http://google.com",
		"http://amazon.com",
		"http://microsoft.com",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, c)
		fmt.Println(<-c)
	}

	fmt.Println("finished in: ", time.Since(inicio))
}

// para verficar el tipo de un paquete de go solo se pone el paquete de Go y se busca el struct que debe empezar con mayuscula
// para saber el tipo de una variable se puede utilizar t := reflect.TypeOf(x)
// func revisarServidor(servidor string, t time.Time, c chan string) {..}

func revisarServidor(servidor string, c chan<- string) {
	timePerRequest := time.Now()
	_, err := http.Get(servidor)

	if err != nil {
		msj := fmt.Sprintf("%s no esta disponible ", servidor)
		c <- msj
	} else {
		msj := fmt.Sprintf("servidor %s funciona correctamente y respondio en %s ", servidor, time.Since(timePerRequest))
		c <- msj
	}

}

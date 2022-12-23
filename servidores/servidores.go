package servidores

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

func GetServidores() {

	inicio := time.Now()

	fmt.Println(reflect.TypeOf(inicio))
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor, inicio)
	}

	fmt.Println("total: ", time.Since(inicio))
}

// para verficar el tipo de un paquete de go solo se pone el paquete de Go y se busca el struct que debe empezar con mayuscula
// para saber el tipo de una variable se puede utilizar t := reflect.TypeOf(x)
func revisarServidor(servidor string, t time.Time) {
	_, err := http.Get(servidor)

	if err != nil {
		fmt.Println(servidor, " no esta disponible ")
	} else {
		fmt.Println(servidor, " funciona correctamente ", time.Since(t))
	}

}

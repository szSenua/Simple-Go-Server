package main

import (
	"fmt"
	"log"
	"net/http"
)

// request es lo que el usuario envía al servidor y la respuesta es lo que el servidor devuelve
// al usuario.
func helloHandler(w http.ResponseWriter, r *http.Request) {

	//Verificamos si la url de la petición es /hello
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	//Verificamos que la solicitud es un GET
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	//Analiza los datos del formulario con la función ParseForm
	//Si diese un fallo, se imprimirá un mensaje de error con el valor de error que tire el formulario.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	//servidor de archivos, le decimos que queremos que revise automáticamente el directorio static
	//y que mire el fichero index.html
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	//La forma en que mostrará el formulario y la función hello
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	//ListenAndServer es una función dentro del paquete net/http que se encarga de levantar el servidor
	//le indicamos que si ocurre algún error a la hora de levantarlo, se registrará un mensaje
	//con log.Fatal
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

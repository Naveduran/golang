// Pasos para crear una función en go

// 0. Instale go
// "wget https://go.dev/dl/go1.19.linux-amd64.tar.gz"
// "sudo tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz"
// "export PATH=$PATH:/usr/local/go/bin"
// "go version"

// Te metes en la carpeta en que vas a crear todo(yo usé "hello")
// 1. Ejecutas "go mod init packagename" para manejar dependencias
// yo puse "go mod init hello"
// Esto crea el archivo go.mod en el directorio actual.
// Este archivo maneja las dependencias en go 
// 2. Creas un archivo.go 
// 3. Declaras el nombre del paquete siempre al inicio
package main

// 4. Importas paquetes de la librería standar.
// Este paquete contiene la funcion para imprimir
import "fmt"

// 5. Importas paquetes externos de https://pkg.go.dev/ (esto es el pip de go)
import "rsc.io/quote"
// Después de declararlos corre el comando: "go mod tidy"
// Esto busca todos los paquetes nuevos en tu código y los descarga


// 6. Escribe tu funcion

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}

// 7. Corre la función con "go run ."

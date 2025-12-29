package main
//módulos necesarios a importar que usara nuestra app
import (
	"fmt"
	"log"
	"net/http"
	"time" 
)

func main() {
	// 1. Decirle al servidor dónde están los archivos estáticos (la imagen)
	// Esto sirve la carpeta ./static en la url /static/
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. La ruta principal (Home)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Escribimos HTML simple directamente en la respuesta
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<div style="text-align: center; margin-top: 50px;">
				<h1>👻 Hola PhantomSec 👻</h1>
				<p>Despliegue Seguro - Actividad 4</p>
				<img src="/static/logo.png" alt="Logo" style="width: 300px; border: 2px solid #333;">
				<p>Estado: <strong>Producción Segura</strong></p>
			</div>
		`)
	})

	// 3. Iniciar el servidor
	port := "8080"
	log.Printf("Servidor arrancando en http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}

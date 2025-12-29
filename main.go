package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

	
	port := "8080"
	log.Printf("Servidor arrancando en http://localhost:%s", port)


	server := &http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

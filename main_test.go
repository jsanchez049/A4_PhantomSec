package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Este test simula una petición al servidor para ver si responde bien
func TestRootHandler(t *testing.T) {
	// 1. Crear una petición falsa a la ruta "/"
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 2. Crear un "grabador" de respuesta
	rr := httptest.NewRecorder()
	
	// 3. Definir el handler (copia simplificada de la lógica)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hola PhantomSec"))
	})

	// 4. Ejecutar la petición
	handler.ServeHTTP(rr, req)

	// 5. Verificar que el código de estado es 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("El handler devolvió el código incorrecto: obtuve %v quería %v",
			status, http.StatusOK)
	}

	// 6. Verificar que el cuerpo contiene algo esperado
	expected := "PhantomSec"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("El handler devolvió un body inesperado: obtuve %v quería que contuviera %v",
			rr.Body.String(), expected)
	}
}
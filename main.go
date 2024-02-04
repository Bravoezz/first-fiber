package main

import (
	"log"
	"os"

	"github.com/Bravoezz/first-fiber/db"
	"github.com/Bravoezz/first-fiber/server"
)

func main() {
	// db connection
	db.InitDbConnection()
	// db.Migrate()

	dir := "storage"

    // Verificar si el directorio ya existe
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        // El directorio no existe, así que lo creamos
        if err := os.Mkdir(dir, os.ModePerm); err != nil {
            log.Fatal(err)
        }
        log.Println("Directorio creado:", dir)
    } else if err != nil {
        // Error al verificar la existencia del directorio
        log.Fatal(err)
    } else {
        // El directorio ya existe, no es necesario crearlo nuevamente
        log.Println("El directorio", dir, "ya existe.")
    }

	server.Start("4000")
}

package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//gorm para manejar la BD
var db *gorm.DB

func main() {
	dsn := "ren:app_password@tcp(localhost:3306)/lab6db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//Se declaran las solicitudes cors autorizadas, para poder realizar los procesos solicitados en la tarea.
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS","PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
	//Se declaran todos los métodos de manejo programados, así como sus respectivos endpoints.
	r.Post("/api/series",HandleRegistration)
	r.Get("/api/series",HandleAllSeries)
	r.Get("/api/series/{id}",HandleSerieID)
	r.Delete("/api/series/{id}",HandleDeleteID)
	r.Put("/api/series/{id}",HandleEditSerieID)
	r.Patch("/api/series/{id}/status", HandleStatusPatch)
	r.Patch("/api/series/{id}/episode",HandleEpisodePatch)
	r.Patch("/api/series/{id}/upvote", HandleUpvotePatch)
	r.Patch("/api/series/{id}/downvote", HandleDownvotePatch)
	//El backend se comunica/le habla al puerto 8080, requerido por la tarea. 
	http.ListenAndServe(":8080", r)
}

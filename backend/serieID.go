package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
	"strconv"
)

//Función que maneja el general a una serie en específico. Recibe el id del url y manda a traer la Serie que en la BD tiene exactamente ese id. Devuelve esa Serie. 
func HandleSerieID(w http.ResponseWriter, r *http.Request) {
    serieID := chi.URLParam(r, "id")
    if serieID == "" {
        respondWithError(w, "ID de la serie requerido", http.StatusBadRequest)
        return
    }

    var result Serie

    if err := db.Raw("SELECT * FROM lab6db.series WHERE id = ?", serieID).Scan(&result).Error; err != nil {
        respondWithError(w, "Database Error", http.StatusInternalServerError)
        return
    }

idSerieMem, _ := strconv.Atoi(serieID)
    if result.ID != idSerieMem {
        respondWithError(w, "Serie no encontrada", http.StatusNotFound)
        return
    }

    respondWithJSON(w, result)
}


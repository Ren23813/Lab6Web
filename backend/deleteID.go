package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
)

//Función para manejar la eliminación de una serie. Recibe del url el id específico de la serie a borrar; accede a la BD y borra la serie
func HandleDeleteID(w http.ResponseWriter, r *http.Request) {
    serieID := chi.URLParam(r, "id")
    if serieID == "" {
        respondWithError(w, "ID de la serie requerido", http.StatusBadRequest)
        return
    }

    var result Serie

    if err := db.Raw("DELETE FROM lab6db.series WHERE id = ?", serieID).Scan(&result).Error; err != nil {
        respondWithError(w, "Database Error", http.StatusInternalServerError)
        return
    }


    respondWithJSON(w, result)
}


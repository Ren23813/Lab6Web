package main

import (
    "net/http"
)
//Función para manejar el recibimiento de todas las series almacenadas en la BD; se mandan como una lista JSON del modelo Serie
func HandleAllSeries(w http.ResponseWriter, r *http.Request) {
    var result []Serie
    if err := db.Raw("SELECT * FROM lab6db.series").Scan(&result).Error; err != nil {
        respondWithError(w, "Database Error", http.StatusInternalServerError)
        return
    }

//Si no existen resultados para series, o sea, no hay series ingresadas, se devuelve una lista vacía.
    if len(result) == 0 {
        respondWithJSON(w, []Serie{})
    }

   respondWithJSON(w, result)
}

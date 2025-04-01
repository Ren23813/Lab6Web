package main


import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "strconv"
)

//Función que maneja el click en la flecha hacia arriba de ranking en el front. Recibe el id de la url, busca en la BD esa serie y le aumenta en 1 su ranking previo; se actualiza la BD.
func HandleUpvotePatch(w http.ResponseWriter, r *http.Request) {
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

serieIDMem, _ := strconv.Atoi(serieID)
    if result.ID !=serieIDMem {
        respondWithError(w, "Serie no encontrada", http.StatusNotFound)
        return
    }



*result.Ranking += 1

 if err := db.Save(&result).Error; err != nil {
        respondWithError(w, "Error al actualizar la serie", http.StatusInternalServerError)
        return
    }


    respondWithJSON(w, result)
}

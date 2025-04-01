package main


import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "strconv"
    "encoding/json"
)

//Función que maneja el click del desplegable para cambio de estatus. Recibe el id de la serie, así como su contenido en forma de Serie. Se sobreescribe por encima de la existente con su mismo id y se guarda
func HandleStatusPatch(w http.ResponseWriter, r *http.Request) {
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

 var updatedSerie Serie
    if err := json.NewDecoder(r.Body).Decode(&updatedSerie); err != nil {
        respondWithError(w, "Error al leer los datos de la serie", http.StatusBadRequest)
        return
    }


if updatedSerie.Status != "" {result.Status = updatedSerie.Status}


 if err := db.Save(&result).Error; err != nil {
        respondWithError(w, "Error al actualizar la serie", http.StatusInternalServerError)
        return
    }


    respondWithJSON(w, result)
}


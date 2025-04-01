package main


import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "strconv"
    "encoding/json"
)

//Función para editar una serie. Lee el id de la url, luego accede a esta serie en la BD, entonces, se crea un objeto Serie extraído del front, el cual se descodifica y se sobrepone encima del de la BD. Se actualiza la BD. 
func HandleEditSerieID(w http.ResponseWriter, r *http.Request) {
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


if updatedSerie.Title != "" {result.Title = updatedSerie.Title}

if updatedSerie.Status != "" {result.Status = updatedSerie.Status}

if updatedSerie.TotalEpisodes != 0 {result.TotalEpisodes = updatedSerie.TotalEpisodes}

if updatedSerie.LastEpisodeWatched != nil {result.LastEpisodeWatched = updatedSerie.LastEpisodeWatched}

if updatedSerie.Ranking != nil {result.Ranking = updatedSerie.Ranking}

 if err := db.Save(&result).Error; err != nil {
        respondWithError(w, "Error al actualizar la serie", http.StatusInternalServerError)
        return
    }


    respondWithJSON(w, result)
}


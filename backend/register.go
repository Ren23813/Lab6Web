package main
import (
	"encoding/json"
	"net/http"
)

//Función para manejar el ingreso de una serie. Básicamente, se validan que no hayan datos vacíos, y se ingresan a la BD, según el payload mandado por el frontend
func HandleRegistration(w http.ResponseWriter, r *http.Request){
var req Serie
if err := json.NewDecoder(r.Body).Decode(&req); err!=nil{
	respondWithError(w, "Invalid Request", http.StatusBadRequest)
	return }
if req.Title == "" || req.Status == "" || req.LastEpisodeWatched == nil || req.Ranking==nil{
		respondWithError(w, "Data is invalid", http.StatusBadRequest)
	return }
result := db.Exec("INSERT INTO lab6db.series(nombre,estado,total_capitulos,capitulos_vistos,puntuacion) VALUES (?,?,?,?,?)",
req.Title, req.Status, req.TotalEpisodes, req.LastEpisodeWatched ,req.Ranking)

if result.Error != nil {
		respondWithError(w, "Error creating the series.", http.StatusInternalServerError)
		return
	}

respondWithJSON(w, ApiResponse{
		Success: true,
		Message: "Registration successful",
	})
}




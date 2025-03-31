package main

import (
    "net/http"
//"fmt"
)

func HandleAllSeries(w http.ResponseWriter, r *http.Request) {
    var result []Serie
    if err := db.Raw("SELECT * FROM lab6db.series").Scan(&result).Error; err != nil {
        respondWithError(w, "Database Error", http.StatusInternalServerError)
        return
    }

   
    if len(result) == 0 {
        respondWithJSON(w, []Serie{})
    }

   respondWithJSON(w, result)
}

package main

//Modelo para estructurar objetos de tipo Serie, principales para este lab. Todas las series tendrán esta forma. 
//En el gorm, se colocan los nombres de las columnas en la BD, así como sus restricciones. 
//En el json, se coloca el nombre exacto que espera el frontend para recibir y mandar sus datos al back.
type Serie struct {
    ID		      int    `gorm:"column:id;primaryKey" json:"id"`
    Title             string `gorm:"column:nombre;not null" json:"title"`
    Status            string `gorm:"column:estado;not null" json:"status"`
    TotalEpisodes     int    `gorm:"column:total_capitulos" json:"totalEpisodes"`
    LastEpisodeWatched *int   `gorm:"column:capitulos_vistos;not null" json:"lastEpisodeWatched"`
    Ranking           *int   `gorm:"column:puntuacion;not null" json:"ranking"`
}



//Estructura genérica de respuesta al front. 
type ApiResponse struct {
    Success bool     `json:"success"`
    Message string   `json:"message"`
    Series  []Serie  `json:"series,omitempty"`
}

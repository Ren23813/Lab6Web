package main

type Serie struct {
    ID		      int    `gorm:"column:id;primaryKey" json:"id"`
    Title             string `gorm:"column:nombre;not null" json:"title"`
    Status            string `gorm:"column:estado;not null" json:"status"`
    TotalEpisodes     int    `gorm:"column:total_capitulos" json:"totalEpisodes"`
    LastEpisodeWatched *int   `gorm:"column:capitulos_vistos;not null" json:"lastEpisodeWatched"`
    Ranking           *int   `gorm:"column:puntuacion;not null" json:"ranking"`
}




type ApiResponse struct {
    Success bool     `json:"success"`
    Message string   `json:"message"`
    Series  []Serie  `json:"series,omitempty"`
}

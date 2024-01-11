package main

import (
	"fmt"
	"net/http"

	"github.com/MauriPinoRicci/fortnite_go/utils"
	"github.com/go-chi/chi/v5"
)

type Skin struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Color  string `json:"color"`
	Season string `json:"season"`
}

type Gun struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Season  string `json:"season"`
	Cuality string `json:"cuality"`
}

func main() {
	router := chi.NewRouter()
	router.Get("/v1/skins", func(w http.ResponseWriter, r *http.Request) {
		utils.RenderResponse(w, r, 200, Skin{Id: "1", Name: "Caballero Negro", Color: "Negro y Rojo", Season: "Temporada 1"})
	})
	
	router.Get("/v2/skins/guns", func(w http.ResponseWriter, r *http.Request) {
		utils.RenderResponse(w, r, 200, Gun{Id: "1", Name: "Escopeta Corredera", Season: "Todas las temporadas", Cuality: "Legendaria"})
	})
	port := "3000"

	fmt.Println(fmt.Sprintf("Server running in port %s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

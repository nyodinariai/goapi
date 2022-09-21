package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

//Retonar um router com as rota configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
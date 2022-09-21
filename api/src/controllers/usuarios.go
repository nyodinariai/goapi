package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Criando Usuario"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscar Usuarios"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscar Usuario"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizar Usuario"))
}

func Deletarusuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletar Usuario"))
}
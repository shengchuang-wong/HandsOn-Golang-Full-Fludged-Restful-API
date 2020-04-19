package controllers

import (
	"net/http"

	"github.com/shengchuang-wong/HandsOn-Golang-Full-Fludged-Restful-API/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}

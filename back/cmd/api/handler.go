package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"

	"github.com/wesleycremonini/back/internal/response"
	"go.uber.org/zap"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "OK",
	}

	err := response.Success(w, http.StatusOK, "OK", data)
	if err != nil {
		serverError(w, err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	var data []string
	response.Error(w, http.StatusNotFound, "Not found", data)
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	var data []string
	response.Error(w, http.StatusNotFound, message, data)
}

func serverError(w http.ResponseWriter, err error) {
	zap.L().Error(err.Error())
	message := "The server encountered a problem and could not process your request"
	var data []string
	response.Error(w, http.StatusInternalServerError, message, data)
}

func (app *application) handleHello(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("hello endpoint hit")

	hellos := []string{
		"Olá",
		"Hello",
		"Hola",
		"Bonjour",
		"こんにちは",
		"你好",
		"안녕하세요",
		"Здравствуйте",
		"مرحبا",
	}

	randomInt, _ := rand.Int(rand.Reader, big.NewInt(int64(len(hellos))))

	err := response.Success(w, http.StatusOK, "OK", map[string]string{"hello": hellos[randomInt.Int64()]})
	if err != nil {
		serverError(w, err)
	}
}

package controllers

import (
	"net/http"

	"github.com/BarunKGP/nlquery/core"
	"github.com/BarunKGP/nlquery/ports"
	"github.com/julienschmidt/httprouter"
)

func HandleCreateThread(
	e ports.EnvReader,
	w http.ResponseWriter,
	r *http.Request,
	p httprouter.Params) error {
	chathub := core.NewChatHub()
	webClient := core.NewClient(1, &chathub)

	return nil

}

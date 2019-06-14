package controllers

import (
	"net/http"

	"hackathon/controllers/index"
)

var Routers = map[string]http.HandlerFunc{
	"/":   index.Index,
	"/ws": index.Ws,
}

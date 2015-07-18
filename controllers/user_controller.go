package controllers

import (
	"github.com/carrot/gcore/models"
	"github.com/carrot/gcore/response"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type UserController struct{}

func (c *UserController) ReadSingle(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Prepping response
	resp := response.New()
	defer resp.Output(writer)

	// Parsing out params
	id, err := strconv.ParseUint(params.ByName("id"), 0, 64)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.ErrorCode = response.ErrorInvalidParameters
		return
	}

	// Loading User
	user := new(models.User)
	user.Load(id)

	// Prepping success response
	resp.Content = user
}

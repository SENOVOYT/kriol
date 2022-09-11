// Filename: cmd/api/helpers.go
package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)
func (app *application) readIDParam(r *http.Request) (int64, error){
	//use the "paramsfromcontext()" function to get the request context a slice
	params := httprouter.ParamsFromContext(r.Context())
	//GET the value of the "id" parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		
		return 0, errors.New("Invalid Id Parameter")
	}
	return id, nil
}
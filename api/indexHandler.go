package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/context"
	"github.com/maleck13/testapp/config"
	
	
	
)

//Example route handler
func IndexHandler(rw http.ResponseWriter, req *http.Request) HttpError {
	encoder := json.NewEncoder(rw)
	resData := make(map[string]string)
	resData["example"] = config.Conf.GetExample()

	val,has := context.GetOk(req,"test")
	if has{
		resData["context"] = val.(string)
	}

	if err := encoder.Encode(resData); err != nil {
		return NewHttpError(err, http.StatusInternalServerError)
	}
	return nil
}




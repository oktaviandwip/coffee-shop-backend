package pkg

import (
	"encoding/json"
	"net/http"
	"reflect"
)

func Response(w http.ResponseWriter, statusCode int, result interface{}) {
	desc := ""

	switch statusCode {
	case http.StatusOK:
		desc = "OK"
	case http.StatusCreated:
		desc = "Created"
	case http.StatusNoContent:
		desc = "No Content"
	case http.StatusBadRequest:
		desc = "Bad Request"
	case http.StatusUnauthorized:
		desc = "Unauthorized"
	case http.StatusNotFound:
		desc = "Page Not Found"
	case http.StatusInternalServerError:
		desc = "Internal Server Error"
	case http.StatusBadGateway:
		desc = "Bad Gateway"
	case http.StatusNotModified:
		desc = "Not Modified"
	default:
		desc = ""
	}

	response := map[string]interface{}{
		"status": statusCode,
		"desc":   desc,
	}

	if statusCode >= http.StatusInternalServerError {
		if err, ok := result.(string); ok {
			response["error"] = err
		}
	} else if statusCode >= http.StatusOK {
		if message, ok := result.(string); ok {
			response["message"] = message
		} else if result != nil && reflect.TypeOf(result).Kind() == reflect.Slice {
			response["data"] = result
		}
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

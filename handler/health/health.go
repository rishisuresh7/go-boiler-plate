package health

import (
	"net/http"

	"go-boiler-plate/response"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	response.Success{Success: "I'm alive"}.Send(w)
}

package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/dukesp69/bookings/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
	app.InfoLog.Println("Client error with the status of ", status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

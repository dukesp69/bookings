package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/dukesp69/bookings/internal/config"
	"github.com/dukesp69/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	testApp.InProduction = false
	//what am I going to put in the session
	gob.Register(models.Reservation{})

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	//true when https enabled
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

// all methods that satisface the interface

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw *myWriter) Write(b []byte) (int, error) {
	lenght := len(b)
	return lenght, nil
}

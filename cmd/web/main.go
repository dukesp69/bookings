package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/dukesp69/bookings/internal/config"
	"github.com/dukesp69/bookings/internal/driver"
	"github.com/dukesp69/bookings/internal/handlers"
	"github.com/dukesp69/bookings/internal/helpers"
	"github.com/dukesp69/bookings/internal/models"
	"github.com/dukesp69/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	// wont close the database connection until run funcion is finish
	defer db.SQL.Close()

	defer close(app.MailChan)
	listenForMail()
	fmt.Println("Starting mail server fake")
	// msg := models.MailData{
	// 	To:      "test@test.com",
	// 	From:    "te@te.com",
	// 	Subject: "Hello",
	// 	Content: "<strong>asd</strong>",
	// }

	// app.MailChan <- msg

	fmt.Println("Starting application")

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	app.InProduction = false
	//what am I going to put in the session
	gob.Register(models.User{})
	gob.Register(models.Reservation{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(models.RoomRestriction{})
	gob.Register(map[string]int{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	//true when https enabled
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//connect to database
	log.Println("Connecting to database")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=")
	if err != nil {
		log.Fatal("Cannot connect to database...Die!")
	}
	log.Println("Connected to database")
	tc, err := render.CreateTemplateCache()
	log.Println(err)
	if err != nil {
		log.Fatal("Cannot create template cache")

		return nil, err
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}

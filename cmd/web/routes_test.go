package main

import (
	"fmt"
	"testing"

	"github.com/dukesp69/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing; test passwed
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is: %T", v))
	}
}

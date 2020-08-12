package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App struct
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize App
func (a *App) Initialize(user, password, dbname string) {}

// Run App
func (a *App) Run(addr string) {}

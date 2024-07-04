package api

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// SetupRouter configures the routes for the application
func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	return router
}

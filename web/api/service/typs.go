package service

import (
	"mymap/db"
	"github.com/jinzhu/gorm"
)

// Server .
type Server struct {
	*gorm.DB
}

// GetServer .
func GetServer() *Server {
	return &Server{
		db.GetDB(),
	}
}

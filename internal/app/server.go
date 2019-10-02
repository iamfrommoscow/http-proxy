package server

import "proxy/internal/pkg/db"

func StartApp() {
	db.Connect()
}

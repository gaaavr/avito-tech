package main

import (
	_ "avito/docs"
	"avito/internal/app"
)

// @title Avito-tech
// @version 1.0
// @description API Server for working with user balance
// @contact.name   Sergey Gavrilin
// @contact.email  ssg0808@yandex.ru
// @host localhost:8080
// @BasePath /
// @Schemes http
func main() {
	app.InitApi()
}

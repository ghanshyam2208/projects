package api

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Server struct {
	Router    *echo.Echo
	sqlClient *sqlx.DB
}

func NewServer() *Server {
	server := &Server{
		Router: echo.New(),
	}

	// add routes
	server.Router.POST("/accounts", CreateAccount)

	// Connect to the database
	server.connectToDB()

	return server
}

func (s *Server) connectToDB() {
	// connect to db
	sqlDb, err := sqlx.Connect("postgres", "user=postgres password=password dbname=simple_bank sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	s.sqlClient = sqlDb
	log.Println("Successfully connected to the database")
}

// handlers
func CreateAccount(c echo.Context) error {
	// Implement your account creation logic here
	return c.String(http.StatusOK, "Account created")
}

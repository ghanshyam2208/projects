package api

import (
	"log"
	"net/http"
	"simple_banking_app/data"
	"time"

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
	server.Router.POST("/accounts", server.CreateAccount)

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
func (s *Server) CreateAccount(c echo.Context) error {
	// Define a struct to receive the request data
	type CreateAccountRequest struct {
		Owner    string `json:"owner" validate:"required"`
		Balance  int64  `json:"balance" validate:"required"`
		Currency string `json:"currency" validate:"oneof=ECDSASecp256k1VerificationKey Ed25519VerificationKey RSAVerificationKey"`
	}

	var req CreateAccountRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	// Create an account instance
	account := data.Accounts{
		Owner:     req.Owner,
		Balance:   req.Balance,
		Currency:  req.Currency,
		CreatedAt: time.Now(),
	}

	// Insert the account into the database
	query := `INSERT INTO accounts (owner, balance, currency, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := s.sqlClient.QueryRowx(query, account.Owner, account.Balance, account.Currency, account.CreatedAt).Scan(&account.ID)
	if err != nil {
		log.Println("Failed to create account:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create account"})
	}

	return c.JSON(http.StatusOK, account)
}

package api

import (
	"fmt"
	"log"
	"net/http"
	"simple_banking_app/data"
	"time"

	"github.com/go-playground/validator"
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
		Currency string `json:"currency" validate:"required,oneof=USD INR EUR"`
	}

	var createAccountRequest CreateAccountRequest
	err := c.Bind(&createAccountRequest)

	if err != nil {
		log.Println("bad request error")
		return c.String(http.StatusBadRequest, "bad request")
	}

	// Create a new validator instance
	v := validator.New()

	// Validate the request struct
	err = v.Struct(createAccountRequest)

	if err != nil {
		// Return a 422 Unprocessable Entity response with error details
		log.Println("validation error")
		// Get the validation errors
		validationErrors := err.(validator.ValidationErrors)

		// Create a map to store the error messages
		errorMessages := make(map[string]string)

		// Iterate over the validation errors and create a user-friendly error message
		for _, validationError := range validationErrors {
			field := validationError.Field()
			tag := validationError.Tag()

			switch tag {
			case "required":
				errorMessages[field] = fmt.Sprintf("The %s field is required", field)
			default:
				errorMessages[field] = fmt.Sprintf("Invalid %s", field)
			}
		}

		// Return a 422 Unprocessable Entity response with error details
		return c.JSON(http.StatusUnprocessableEntity, errorMessages)
	}

	// Create an account instance
	account := data.Accounts{
		Owner:     createAccountRequest.Owner,
		Balance:   createAccountRequest.Balance,
		Currency:  createAccountRequest.Currency,
		CreatedAt: time.Now(),
	}

	// Insert the account into the database
	query := `INSERT INTO accounts (owner, balance, currency, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err = s.sqlClient.QueryRowx(query, account.Owner, account.Balance, account.Currency, account.CreatedAt).Scan(&account.ID)
	if err != nil {
		log.Println("Failed to create account:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create account"})
	}

	return c.JSON(http.StatusOK, account)
}

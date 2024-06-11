package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"simple_banking_app/data"
	"simple_banking_app/utils"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Server struct {
	Router          *echo.Echo
	sqlClient       *sqlx.DB
	serverValidator *validator.Validate
	config          *utils.Config
}

func removeTrailingSlash(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := c.Request().URL.Path
		if strings.HasSuffix(p, "/") && p != "/" {
			c.Request().URL.Path = p[:len(p)-1]
		}
		return next(c)
	}
}

func NewServer() *Server {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	server := &Server{
		Router: echo.New(),
		config: &config,
	}

	// Create a new validator instance
	server.serverValidator = validator.New()

	server.config = &config

	log.Println("config in server:", server.config)
	log.Println("config in server:", server.config.PostgresConnStr)

	// middleware to remove the trailing slash
	server.Router.Use(removeTrailingSlash)
	// add routes
	server.Router.POST("/accounts/", server.CreateAccount)
	server.Router.POST("/accounts", server.CreateAccount)
	server.Router.GET("/accounts/:id", server.GetAccountById)

	// Connect to the database
	server.connectToDB()
	return server
}

func (s *Server) connectToDB() {

	// connect to db
	sqlDb, err := sqlx.Connect("postgres", s.config.PostgresConnStr)
	if err != nil {
		log.Fatalln(err)
	}

	s.sqlClient = sqlDb
	log.Println("Successfully connected to the database")
}

func (s *Server) GetAccountById(c echo.Context) error {
	idParam := c.Param("id")
	id := struct {
		ID int `validate:"required,gt=0"`
	}{}

	id.ID, _ = strconv.Atoi(idParam)

	err := s.serverValidator.Struct(id)
	if err != nil {
		return validationError(c, err)
	}

	// Now you can use the validated id
	log.Println("id:", id.ID)

	// Query the database for the account by ID
	getACustomerSql := "SELECT id, owner, balance, currency, created_at from accounts WHERE id = $1"

	var account data.Accounts
	err = s.sqlClient.Get(&account, getACustomerSql, id.ID)
	if err != nil {
		// Handle database errors
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "Account not found")
		}
		log.Println("Database error:", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Return the account as a JSON response
	return c.JSON(http.StatusOK, account)
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

	// Validate the request struct
	err = s.serverValidator.Struct(createAccountRequest)

	if err != nil {
		return validationError(c, err)
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

func validationError(c echo.Context, err error) error {
	log.Println("validation error")
	validationErrors := err.(validator.ValidationErrors)

	errorMessages := make(map[string]string)

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

	return c.JSON(http.StatusUnprocessableEntity, errorMessages)
}

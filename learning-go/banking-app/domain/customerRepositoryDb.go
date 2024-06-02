package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/ghanshyam2208/banking/errs"
	"github.com/ghanshyam2208/banking/logger"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllSql string
	var err error
	var rows *sql.Rows

	if status == "" {
		findAllSql = "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql = "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers WHERE status = ?"
		rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		// log.Println("error while querying the database " + err.Error())
		logger.Error("error while querying the database " + err.Error())
		return nil, errs.NewInternalServerError()
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("no customer found") // errors.errors.New("customer not found")
			} else {
				// log.Println("error while querying the database " + err.Error())
				logger.Error("error while querying the database " + err.Error())
				return nil, errs.NewInternalServerError("unexpected database error") // errors.New("unexpected database error")
			}
		}
		customers = append(customers, c)
	}
	return customers, nil

}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	getACustomerSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers WHERE customer_id = ?"

	row := d.client.QueryRow(getACustomerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found") // errors.errors.New("customer not found")
		} else {
			log.Println("error while querying the database " + err.Error())
			logger.Error("error while querying the database " + err.Error())
			return nil, errs.NewInternalServerError("unexpected database error") // errors.New("unexpected database error")
		}

	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}

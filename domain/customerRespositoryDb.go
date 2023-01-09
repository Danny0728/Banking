package domain

import (
	"database/sql"
	"fmt"

	"github.com/Danny0728/BankAPI/errs"
	"github.com/Danny0728/BankAPI/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRespositoryDb struct {
	client *sqlx.DB
}

var statusMap = map[string]int{
	"active":   1,
	"inactive": 0,
}

func (d CustomerRespositoryDb) FindAll(status *string) ([]Customer, *errs.AppError) {
	customers := make([]Customer, 0)

	findAllsql := "select customer_id, name,city,zipcode,date_of_birth, status from customers"
	if status != nil {
		statusValue := statusMap[*status]
		findAllsql += fmt.Sprintf(" where status = %d", statusValue)
	}
	//this Select function does the work on .Query as well as the rows.Next() or the StructScan function in sqlx
	err := d.client.Select(&customers, findAllsql)
	// rows, err := d.client.Query(findAllsql)
	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected Error from serverside")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customer table" + err.Error())
	// 	return nil, errs.NewUnExpectedError("Unexpected Error from serverside")
	// }
	return customers, nil
}

// func (d CustomerRespositoryDb) AddCustomer(c Customer) (Customer, error) {

// }

func (d CustomerRespositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var c Customer
	customerSql := "select customer_id, name,city,zipcode,date_of_birth, status from customers where customer_id = ?"
	err := d.client.Get(&c, customerSql, id)

	switch {
	case err == sql.ErrNoRows:
		logger.Error("No Customer with provided id " + err.Error())
		return nil, errs.NewNotFoundError("Customer code does not exist")
	case err != nil:
		logger.Error("Error while scanning customer table" + err.Error())
		return nil, errs.NewUnExpectedError("Unexpected database error")
	}
	return &c, nil
}
func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRespositoryDb {
	return CustomerRespositoryDb{client: dbClient}
}

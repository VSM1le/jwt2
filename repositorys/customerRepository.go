package repositorys

import (
	"time"

	"github.com/VSM1le/jwt2/models"
	"github.com/gofiber/fiber/v2"
)

func (r *PostgreSQLRepository) SelectAllCustomer(ctx *fiber.Ctx) ([]models.Customer, error) {
	var customers []models.Customer
	query := `SELECT * FROM customers`
	err := r.db.SelectContext(ctx.Context(), &customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *PostgreSQLRepository) CreateCustomer(ctx *fiber.Ctx, customer *models.Customer) error {
	query := `INSERT INTO 
		customers 
			(cust_code,
			cust_name,
			cust_address_1,
			cust_address_2,
			cust_zipcode,
			cust_branch,
			cust_type,
			created_by)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := r.db.ExecContext(ctx.Context(), query,
		customer.CustCode,
		customer.CustName,
		customer.CustAddress1,
		customer.CustAddress2,
		customer.CustZipcode,
		customer.CustBranch,
		customer.CustType,
		customer.CreatedBy,
	)
	if err != nil {
		return err
	}
	return nil
}
func (r *PostgreSQLRepository) UpdateCustomer(ctx *fiber.Ctx, customer *models.Customer, id int64) error {
	query := `UPDATE customers 
		SET
			cust_code = $1,
			cust_name = $2,
			cust_address_1 = $3,
			cust_address_2 = $4,
	    	cust_zipcode = $5,
			cust_branch = $6,
			cust_type = $7,
			updated_at = $8,
			updated_by = $9
		where id = $10`
	_, err := r.db.ExecContext(ctx.Context(), query,
		customer.CustCode,
		customer.CustName,
		customer.CustAddress1,
		customer.CustAddress2,
		customer.CustZipcode,
		customer.CustBranch,
		customer.CustType,
		time.Now(),
		customer.UpdatedBy,
		id,
	)
	if err != nil {
		return err
	}

	query = `SELECT * FROM customers WHERE id = $1`

	err = r.db.GetContext(ctx.Context(), customer, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgreSQLRepository) GetCustomer(ctx *fiber.Ctx, id int64) (*models.Customer, error) {
	var customer models.Customer
	query := `SELECT * FROM customers WHERE id=$1`
	err := r.db.GetContext(ctx.Context(), &customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

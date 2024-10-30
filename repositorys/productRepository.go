package repositorys

import (
	"database/sql"
	"fmt"

	"github.com/VSM1le/jwt2/models"
	"github.com/gofiber/fiber/v2"
)

func (r *PostgreSQLRepository) SelectAllProduct(ctx *fiber.Ctx) ([]models.Product, error) {
	var products []models.Product
	query := `SELECT * FROM product_services`
	err := r.db.SelectContext(ctx.Context(), &products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (r *PostgreSQLRepository) CreateProduct(ctx *fiber.Ctx, product *models.Product) error {
	query := `INSERT INTO 
				product_services (ps_code,ps_name_th,ps_name_en,ps_vat,ps_whtax,ps_gov_whtax,created_by) 
				VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := r.db.ExecContext(ctx.Context(), query, product.PsCode, product.PsNameTh, product.PsNameTh, product.PsVat, product.PsWhtax, product.PsGovWhtax, product.CreatedBy)
	if err != nil {
		return err
	}
	return nil
}
func (r *PostgreSQLRepository) GetProduct(ctx *fiber.Ctx, id string) (*models.Product, error) {
	var product models.Product
	query := `SELECT * FROM product_services WHERE id = $1`

	// Use ctx.Context() for context-aware database queries
	err := r.db.GetContext(ctx.Context(), &product, query, id)

	// Check if no rows were found
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}

	return &product, nil
}

func (r *PostgreSQLRepository) UpdateProduct(ctx *fiber.Ctx, product *models.Product) error {
	query := `UPDATE product_services
			  SET ps_code,ps_name_th,ps_name_en,ps_vat,ps_whtax,ps_gov_whtax,updated_by,updated_at
			  VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
			  WHERER id=$9`
	_, err := r.db.ExecContext(ctx.Context(), query, product.PsCode, product.PsNameTh, product.PsVat, product.PsGovWhtax, product.PsGovWhtax, product.UpdatedBy, product.UpdatedAt, product.ID)
	if err != nil {
		return err
	}

	query = `SELECT * FROM product_services where id = $1`
	err = r.db.GetContext(ctx.Context(), product, query, product.ID)
	if err != nil {
		return err
	}
	return nil

}

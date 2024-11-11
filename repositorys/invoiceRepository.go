package repositorys

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/VSM1le/jwt2/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func (r *PostgreSQLRepository) SelectInvoice(c *fiber.Ctx) ([]models.InvoiceHeader, error) {
	var invoices []models.InvoiceHeader
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	query := `SELECT * FROM invoice_headers`
	err = tx.SelectContext(c.Context(), &invoices, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invoice not found")
		}
		return nil, err
	}
	for i := range invoices {
		var details []models.InvoiceDetail
		query = `SELECT * FROM invoice_details where invoice_header_id=$1`
		err = tx.SelectContext(c.Context(), &details, query, invoices[i].Id)
		if err != nil {
			return nil, err
		}
		invoices[i].InvoiceDetail = details
	}
	tx.Commit()
	return invoices, nil
}

func (r *PostgreSQLRepository) CreateInvoice(c *fiber.Ctx, invoice *models.InvoiceHeader) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var id int64
	query := `INSERT INTO 
		invoice_headers (
			inv_no,
			inv_date,
			customer_id,
			inv_cust_name,
			inv_cust_address_1,
			inv_cust_address_2,
			inv_cust_zipcode,
			inv_cust_branch,
			created_by)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id `
	err = tx.QueryRowContext(c.Context(), query,
		invoice.InvNo,
		invoice.InvDate,
		invoice.CustomerId,
		invoice.InvCustName,
		invoice.InvCustAddress1,
		invoice.InvCustAddress2,
		invoice.InvCustZipcode,
		invoice.InvCustbranch,
		invoice.CreatedBy,
	).Scan(&id)
	if err != nil {
		return err
	}

	for i := range invoice.InvoiceDetail {
		var detail models.InvoiceDetail
		var DeId int64
		query = `INSERT INTO 
			invoice_details(
				invoice_header_id,
				invd_ps_code,
				invd_ps_name_th,
				invd_ps_name_en,
				invd_vat,
				invd_whtax,
				invd_amt,
				invd_vat_amt,
				invd_whtax_amt,
				invd_net_amt,
				created_by,
				product_id
				)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
			RETURNING id`
		err := tx.QueryRowxContext(c.Context(), query,
			id,
			invoice.InvoiceDetail[i].InvdPsCode,
			invoice.InvoiceDetail[i].InvdPsNameTh,
			invoice.InvoiceDetail[i].InvdPsNameEn,
			invoice.InvoiceDetail[i].InvdVat,
			invoice.InvoiceDetail[i].InvdWhtax,
			invoice.InvoiceDetail[i].InvdAmt,
			invoice.InvoiceDetail[i].InvdVatAmt,
			invoice.InvoiceDetail[i].InvdWhtaxAmt,
			invoice.InvoiceDetail[i].InvdNetAmt,
			invoice.CreatedBy,
			invoice.InvoiceDetail[i].ProductId,
		).Scan(&DeId)
		if err != nil {
			fmt.Println("here")
			return err
		}
		query = `SELECT * FROM invoice_details where id=$1`
		err = tx.GetContext(c.Context(), &detail, query, DeId)
		if err != nil {
			fmt.Println("here 2")
			return err
		}
		invoice.InvoiceDetail[i] = detail
	}
	tx.Commit()
	return nil

}

func (r *PostgreSQLRepository) CheckInvoice(c *fiber.Ctx, id int) error {
	query := `SELECT * FROM invoice_headers WHERE id = $1`
	var exists int
	err := r.db.GetContext(c.Context(), &exists, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("error there is no invoice id that you are looking for %s", err)
		}
		return err
	}
	query = `SELECT * FROM invoice_details WHERE invoice_header_id = $id and invd_receipt_flag = yes`
	err = r.db.GetContext(c.Context(), &exists, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return fmt.Errorf("there is an invoice detail that already paid")
}

func (r *PostgreSQLRepository) CancelInvoice(c *fiber.Ctx, userId int64, id int) error {
	query := `UPDATE invoice_headers
			SET inv_status = 'inactive' , updated_at = $1, updated_by = $2 WHERE id = $3`
	_, err := r.db.ExecContext(c.Context(), query, userId, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgreSQLRepository) UpdateInvoice(c *fiber.Ctx, invoice *models.InvoiceHeader, id int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `UPDATE invoice_headers
		SET 
			inv_no=$1,
			inv_inv_date=$2,
			customer_id=$3,
			inv_cust_name=$4,
			inv_cust_address_1=$5,
			inv_cust_address_2=$6,
			inv_cust_zipcode=$7,
			inv_cust_branch=$8,
			updated_at=$9,
			updated_by=$10
		WHERE id=$11`
	_, err = tx.ExecContext(c.Context(), query,
		invoice.InvNo,
		invoice.InvDate,
		invoice.CustomerId,
		invoice.InvCustName,
		invoice.InvCustAddress1,
		invoice.InvCustAddress2,
		invoice.InvCustZipcode,
		invoice.InvCustbranch,
		invoice.UpdatedAT,
		invoice.UpdatedBy,
		id,
	)
	if err != nil {
		return err
	}

	var ids []int64
	for _, detail := range invoice.InvoiceDetail {
		if detail.Id != nil {
			ids = append(ids, *detail.Id)
		}
	}

	if len(ids) > 0 {
		query, args, err := sqlx.In(`DELETE FROM invoice_details WHERE invoice_header_id = ? AND id NOT IN (?)`, id, ids)
		if err != nil {
			return err
		}
		query = tx.Rebind(query)
		_, err = tx.ExecContext(c.Context(), query, args...)
		if err != nil {
			return err
		}
	}
	for i := range invoice.InvoiceDetail {
		var detail models.InvoiceDetail
		if invoice.InvoiceDetail[i].Id != nil {
			query = `UPDATE invoice_details 
			SET 
				invd_ps_code=$1
				invd_ps_name_th=$2
				invd_ps_name_en=$3
				invd_vat=$4
				invd_whtax=$5
				invd_amt=$6
				invd_whtax_amt=$7
				invd_net_amt=$8
				updated_at=$9
				updated_by=$10
			WHERE id = $11 and invoice_header_id = $12 
			`
			_, err := tx.ExecContext(c.Context(), query, invoice.InvoiceDetail[i].InvdPsCode,
				invoice.InvoiceDetail[i].InvdPsNameTh,
				invoice.InvoiceDetail[i].InvdPsNameEn,
				invoice.InvoiceDetail[i].InvdVat,
				invoice.InvoiceDetail[i].InvdWhtax,
				invoice.InvoiceDetail[i].InvdAmt,
				invoice.InvoiceDetail[i].InvdWhtaxAmt,
				invoice.InvoiceDetail[i].InvdNetAmt,
				time.Now(),
				invoice.InvoiceDetail[i].UpdatedBy,
				invoice.InvoiceDetail[i].Id,
				id,
			)
			if err != nil {
				return err
			}
			query = `SELECT * FROM invoice_details WHERE = $1`
			row, err := tx.QueryxContext(c.Context(), query, invoice.InvoiceDetail[i].Id)
			if err != nil {
				return err
			}
			err = row.StructScan(&detail)
			if err != nil {
				return err
			}
			invoice.InvoiceDetail[i] = detail
		} else {
			var newID int64
			query := `INSERT INTO 
				invoice_detail (
					invoice_headers_id,
					invd_ps_code,	
					invd_ps_name_th,
					invd_ps_name_en,
					invd_vat,
					invd_whtax,
					invd_amt,
					invd_vat_amt,
					invd_whtax_amt,
					invd_net_amt,
					created_by
				)
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`

			err := tx.QueryRowContext(c.Context(), query,
				invoice.Id,
				detail.InvdPsCode,
				detail.InvdPsNameTh,
				detail.InvdPsNameEn,
				detail.InvdVat,
				detail.InvdWhtax,
				detail.InvdAmt,
				detail.InvdVatAmt,
				detail.InvdWhtaxAmt,
				detail.InvdNetAmt,
				detail.CreatedBy,
				time.Now(),
			).Scan(&newID)
			if err != nil {
				return err
			}
			query = `SELECT * FROM invoice_details WHERE = $1`
			row, err := tx.QueryxContext(c.Context(), query, newID)
			if err != nil {
				return err
			}
			err = row.StructScan(&detail)
			if err != nil {
				return err
			}
			invoice.InvoiceDetail[i] = detail
		}
	}
	return nil
}

func (r *PostgreSQLRepository) CheckInvoiceDetail(c *fiber.Ctx, ids []int64) (bool, error) {
	query := `SELECT 
			COUNT(id) AS existing_count,
			COUNT(CASE WHEN invd_receipt_flag = 'yes' THEN 1 END) AS has_status_yes_count
			FROM invoice_details
			WHERE id = ANY($1)`
	var existingCount, hasStatusYes int
	// var hasStatusYes bool
	err := r.db.QueryRowxContext(c.Context(), query, ids).Scan(&existingCount, &hasStatusYes)
	if err != nil {
		return true, err
	}
	if existingCount == len(ids) && hasStatusYes == 0 {
		return false, nil
	}
	return true, nil
}

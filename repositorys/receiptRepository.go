package repositorys

import (
	"database/sql"
	"fmt"

	"github.com/VSM1le/jwt2/models"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func (r *PostgreSQLRepository) SelectReceipt(c *fiber.Ctx) ([]models.ReceiptHeader, error) {

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var receipts []models.ReceiptHeader
	query := `SELECT * FROM receipt_headers`
	err = tx.SelectContext(c.Context(), receipts, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("receipt not found")
		}
		return nil, err
	}
	if len(receipts) == 0 {
		return nil, fmt.Errorf("no receipts found")
	}

	var receiptDetails []models.ReceiptDetail
	query = `SELECT * FROM receipt_details WHERE receipt_header_id = ANY($1)`
	receiptHeaderIDs := make([]int64, len(receipts))
	for i, r := range receipts {
		receiptHeaderIDs[i] = r.Id
	}

	err = tx.SelectContext(c.Context(), &receiptDetails, query, pq.Array(receiptHeaderIDs))
	if err != nil {
		return nil, err
	}

	receiptDetailsMap := make(map[int64][]models.ReceiptDetail)
	for _, detail := range receiptDetails {
		receiptDetailsMap[detail.ReceiptHeaderId] = append(receiptDetailsMap[detail.ReceiptHeaderId], detail)
	}

	for i := range receipts {
		receipts[i].ReceiptDetails = receiptDetailsMap[receipts[i].Id]
	}
	tx.Commit()
	return receipts, nil
}

func (r *PostgreSQLRepository) CreateReceipt(c *fiber.Ctx, receipt *models.ReceiptHeader) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `INSERT INTO 
			invoice_headers (
				rec_no,
				rec_date,
				rec_payment_amt,
				customer_id,
				rec_cust_name,
				rec_cust_address_1,
				rec_cust_address_2,
				rec_cust_zipcode,
				rec_cust_branch,
				created_by,
				)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`
	var id int64
	err = tx.QueryRowxContext(c.Context(), query,
		receipt.RecNo,
		receipt.RecDate,
		receipt.RecPaymentAmt,
		receipt.CustomerId,
		receipt.RecCustName,
		receipt.RecCustAddress1,
		receipt.RecCustAddress2,
		receipt.RecCustZipcode,
		receipt.RecCustBranch,
		receipt.CreatedBy,
	).Scan(&id)
	if err != nil {
		return err
	}
	query = `INSERT INTO 
			invoice_details (
				receipt_header_id,
				invoice_detail_id,
				recd_inv_no,
				recd_ps_code,
				recd_ps_name_th,
				recd_ps_name_en,
				recd_vat,
				recd_whtax,
				recd_amt,
				recd_vat_amt,
				recd_whtax_amt,
				recd_net_amt,
				recd_wh_pay,
				created_by,
			)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`
	for i := range receipt.ReceiptDetails {
		_, err = tx.ExecContext(c.Context(), query,
			id,
			receipt.ReceiptDetails[i].InvoiceDetailId,
			receipt.ReceiptDetails[i].RecdInvNo,
			receipt.ReceiptDetails[i].RecdPsCode,
			receipt.ReceiptDetails[i].RecdPsNameTh,
			receipt.ReceiptDetails[i].RecdPsNameEn,
			receipt.ReceiptDetails[i].RecdVat,
			receipt.ReceiptDetails[i].RecdWhtax,
			receipt.ReceiptDetails[i].RecdAmt,
			receipt.ReceiptDetails[i].RecdVatAmt,
			receipt.ReceiptDetails[i].RecdWhtaxAmt,
			receipt.ReceiptDetails[i].RecdNetAmt,
			receipt.ReceiptDetails[i].RecdWhPay,
			receipt.CreatedBy,
		)
		if err != nil {
			return err
		}
		query = `UPDATE invoice_detail 
			SET invd_receipt_flag = yes 
			WHERE id = $2`
		_, err = tx.ExecContext(c.Context(), query, receipt.ReceiptDetails[i].InvoiceDetailId)
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}

func (r *PostgreSQLRepository) CancelReceipt(c *fiber.Ctx, id int) error {

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	query := `UPDATE receipt_headers
			SET rec_status = 'inactive'
			WHERE id = $1`
	_, err = tx.ExecContext(c.Context(), query, id)
	if err != nil {
		return err
	}
	query = `UPDATE invoice_details AS ids
			SET invd_status = 'no'
			FROM receipt_details AS rd
			WHERE ids.id = rd.invoice_detail_id AND rd.receipt_header_id = $1 AND ids.invd_status != 'no'
			`
	_, err = tx.ExecContext(c.Context(), query, id)
	if err != nil {
		return err
	}
	return nil
}

package models

import (
	"time"
)

type InvoiceHeader struct {
	Id              *int64          `json:"id" db:"id"`
	InvStatus       string          `json:"inv_status" db:"inv_status"`
	InvNo           string          `json:"inv_no" db:"inv_no"`
	InvDate         string          `json:"inv_date" db:"inv_date" validate:"datetime=2006-01-02"`
	CustomerId      int64           `json:"customer_id" db:"customer_id"`
	InvCustName     string          `json:"inv_cust_name" db:"inv_cust_name"`
	InvCustAddress1 string          `json:"inv_cust_address_1" db:"inv_cust_address_1"`
	InvCustAddress2 string          `json:"inv_cust_address_2" db:"inv_cust_address_2"`
	InvCustZipcode  string          `json:"inv_cust_zipcode" db:"inv_cust_zipcode"`
	InvCustbranch   string          `json:"inv_cust_branch" db:"inv_cust_branch"`
	InvRemark       *string         `json:"inv_remark" db:"inv_remark"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
	UpdatedAT       *time.Time      `json:"updated_at" db:"updated_at"`
	CreatedBy       int64           `json:"created_by" db:"created_by"`
	UpdatedBy       *int64          `json:"updated_by" db:"updated_by"`
	InvoiceDetail   []InvoiceDetail `json:"invoice_detail"`
}

type InvoiceDetail struct {
	Id              *int64     `json:"id" db:"id"`
	InvoiceHeaderId int64      `json:"invoice_header_id" db:"invoice_header_id"`
	ProductId       int64      `json:"product_id" db:"product_id"`
	InvdPsCode      string     `json:"invd_ps_code" db:"invd_ps_code"`
	InvdPsNameTh    string     `json:"invd_ps_name_th" db:"invd_ps_name_th"`
	InvdPsNameEn    string     `json:"invd_ps_name_en" db:"invd_ps_name_en"`
	InvdVat         float32    `json:"invd_vat" db:"invd_vat"`
	InvdWhtax       float32    `json:"invd_whtax" db:"invd_whtax"`
	InvdAmt         float64    `json:"invd_amt" db:"invd_amt"`
	InvdVatAmt      float64    `json:"invd_vat_amt" db:"invd_vat_amt"`
	InvdWhtaxAmt    float64    `json:"invd_whtax_amt" db:"invd_whtax_amt"`
	InvdNetAmt      float64    `json:"invd_net_amt" db:"invd_net_amt"`
	InvdReceiptFlag string     `json:"invd_receipt_flag" db:"invd_receipt_flag"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAT       *time.Time `json:"updated_at" db:"updated_at"`
	CreatedBy       int64      `json:"created_by" db:"created_by"`
	UpdatedBy       *int64     `json:"updated_by" db:"updated_by"`
}

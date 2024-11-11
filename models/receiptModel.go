package models

import "time"

type ReceiptHeader struct {
	Id              int64      `json:"id" db:"id"`
	RecStatus       string     `json:"rec_status" db:"rec_status"`
	RecNo           string     `json:"rec_no" db:"rec_no"`
	RecDate         string     `json:"rec_date" db:"rec_date"`
	RecPaymentAmt   float64    `json:"rec_payment_amt" db:"rec_payment_amt"`
	CustomerId      int64      `json:"customer_id" db:"customer_id"`
	RecCustName     string     `json:"rec_cust_name" db:"rec_cust_name"`
	RecCustAddress1 string     `json:"rec_cust_address_1" db:"rec_cust_address_1"`
	RecCustAddress2 string     `json:"rec_cust_address_2" db:"rec_cust_address_2"`
	RecCustZipcode  string     `json:"rec_cust_zipcode" db:"rec_cust_zipcode"`
	RecCustBranch   string     `json:"rec_cust_branch" db:"rec_cust_branch"`
	RecRemark       string     `json:"rec_remark" db:"rec_remark"`
	CreatedBy       int64      `json:"created_by" db:"created_by"`
	UpdatedBy       int64      `json:"updated_by" db:"updated_by"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at"`
	ReceiptDetails  []ReceiptDetail
}

type ReceiptDetail struct {
	Id              int64      `json:"id" db:"id"`
	ReceiptHeaderId int64      `json:"receipt_header_id" db:"receipt_header_id"`
	InvoiceDetailId int64      `json:"invoice_detail_id" db:"invoice_detail_id"`
	RecdInvNo       string     `json:"recd_inv_no" db:"recd_inv_no"`
	RecdPsCode      string     `json:"recd_ps_code" db:"recd_ps_code"`
	RecdPsNameTh    string     `json:"recd_ps_name_th" db:"recd_ps_name_th"`
	RecdPsNameEn    string     `json:"recd_ps_name_en" db:"recd_ps_name_en"`
	RecdVat         float32    `json:"recd_vat" db:"recd_vat"`
	RecdWhtax       float32    `json:"recd_whtax" db:"recd_whtax"`
	RecdAmt         float64    `json:"recd_amt" db:"recd_amt"`
	RecdVatAmt      float64    `json:"recd_vat_amt" db:"recd_vat_amt"`
	RecdWhtaxAmt    float64    `json:"recd_whtax_amt" db:"recd_whtax_amt"`
	RecdNetAmt      float64    `json:"recd_net_amt" db:"recd_net_amt"`
	RecdWhPay       float64    `json:"recd_wh_pay" db:"recd_wh_pay"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at"`
	CreatedBy       int64      `json:"created_by" db:"created_by"`
	UpdatedBy       *int64     `json:"updated_by" db:"updated_by"`
}

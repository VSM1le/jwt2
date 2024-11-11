package models

import (
	"time"
)

type Customer struct {
	ID           int64      `json:"id" db:"id"`
	CustCode     string     `json:"cust_code" db:"cust_code" validate:"required,min=5,max=5"`
	CustName     string     `json:"cust_name" db:"cust_name" validate:"required,min=1,max=255"`
	CustAddress1 string     `json:"cust_address_1" db:"cust_address_1" validate:"required,min=1,max=255"`
	CustAddress2 string     `json:"cust_address_2" db:"cust_address_2" validate:"required,min=1,max=255"`
	CustZipcode  string     `json:"cust_zipcode" db:"cust_zipcode" validate:"required,min=1,max=10"`
	CustBranch   string     `json:"cust_branch" db:"cust_branch" validate:"required,min=1,max=255"`
	CustType     string     `json:"cust_type" db:"cust_type" validate:"required,min=1,max=255,oneof=person company go"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
	CreatedBy    int64      `json:"created_by" db:"created_by"`
	UpdatedBy    *int64     `json:"updated_by" db:"updated_by"`
}

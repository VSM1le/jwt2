package models

import "time"

type Product struct {
	ID         int64      `json:"id" db:"id"`
	PsCode     string     `json:"ps_code" db:"ps_code" validate:"required,min=4,max=4"`
	PsNameTh   string     `json:"ps_name_th" db:"ps_name_th" vaidate:"required,min=1,max=255"`
	PsNameEn   string     `json:"ps_name_en" db:"ps_name_en" validate:"required,min=1,max=255"`
	PsVat      float32    `json:"ps_vat" db:"ps_vat" validate:"required,number"`
	PsWhtax    float32    `json:"ps_whtax" db:"ps_whtax" validate:"required,number"`
	PsGovWhtax float32    `json:"ps_gov_whtax" db:"ps_gov_whtax" validate:"required,number"`
	CreatedBy  int64      `json:"created_by" db:"created_by"`
	UpdatedBy  *int64     `json:"updated_by" db:"updated_by"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" db:"updated_at"`
}

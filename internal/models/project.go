package models

import "time"

type Project struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	TenantID     int       `json:"tenant_id"`
	LanguageCode string    `json:"language_code"`
	CreateDate   time.Time `json:"create_date"`
	UpdateDate   time.Time `json:"update_date"`
}

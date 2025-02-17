package models

import "time"

type Question struct {
	ID           int       `json:"id"`
	ProjectID    int       `json:"project_id"`
	Question     string    `json:"question"`
	Type         string    `json:"type"`
	CreateDate   time.Time `json:"create_date"`
	UpdateDate   time.Time `json:"update_date"`
	LanguageCode string    `json:"language_code"`
	TenantID     int       `json:"tenant_id"`
}

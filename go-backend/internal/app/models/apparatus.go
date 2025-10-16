package models

import "time"

type Apparatus struct {
	ID           int64     `json:"id"`
	UniversityID int64     `json:"university_id"`
	ExperimentID *int64    `json:"experiment_id,omitempty"`
	Name         string    `json:"name"`
	Type         string    `json:"type,omitempty"`
	Icon         string    `json:"icon,omitempty"`
	ModelPath    string    `json:"model_path,omitempty"`
	Meta         JSONMap   `json:"meta,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

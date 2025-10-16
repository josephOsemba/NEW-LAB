package models

import "time"

type Lab struct {
	ID           int64     `json:"id"`
	UniversityID int64     `json:"university_id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	Description  string    `json:"description,omitempty"`
	Thumbnail    string    `json:"thumbnail,omitempty"`
	Config       JSONMap   `json:"config,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type LabWithExperiments struct {
	Lab
	Experiments []Experiment `json:"experiments,omitempty"`
}

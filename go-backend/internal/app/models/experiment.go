package models

import "time"

type Difficulty string

const (
	DifficultyBeginner     Difficulty = "beginner"
	DifficultyIntermediate Difficulty = "intermediate"
	DifficultyAdvanced     Difficulty = "advanced"
)

type Experiment struct {
	ID           int64      `json:"id"`
	LabID        int64      `json:"lab_id"`
	UniversityID int64      `json:"university_id"`
	Title        string     `json:"title"`
	Slug         string     `json:"slug"`
	Summary      string     `json:"summary,omitempty"`
	Difficulty   Difficulty `json:"difficulty"`
	Config       JSONMap    `json:"config,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type ExperimentWithApparatus struct {
	Experiment
	Apparatus []Apparatus `json:"apparatus,omitempty"`
}

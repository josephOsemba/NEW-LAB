package store

import (
	"database/sql"

	"github.com/josephOsemba/go-backend/internal/app/models"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore(dataSourceName string) (*MySQLStore, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLStore{db: db}, nil
}

func (s *MySQLStore) Close() error {
	return s.db.Close()
}

// University methods
func (s *MySQLStore) GetUniversities() ([]models.University, error) {
	query := `SELECT id, slug, name, metadata, created_at, updated_at FROM universities`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var universities []models.University
	for rows.Next() {
		var u models.University
		err := rows.Scan(&u.ID, &u.Slug, &u.Name, &u.Metadata, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		universities = append(universities, u)
	}

	return universities, nil
}

func (s *MySQLStore) GetUniversityBySlug(slug string) (*models.University, error) {
	query := `SELECT id, slug, name, metadata, created_at, updated_at FROM universities WHERE slug = ?`
	row := s.db.QueryRow(query, slug)

	var u models.University
	err := row.Scan(&u.ID, &u.Slug, &u.Name, &u.Metadata, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Lab methods
func (s *MySQLStore) GetLabsByUniversity(universityID int64) ([]models.Lab, error) {
	query := `SELECT id, university_id, name, slug, description, thumbnail, config, created_at, updated_at 
	          FROM labs WHERE university_id = ?`
	rows, err := s.db.Query(query, universityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var labs []models.Lab
	for rows.Next() {
		var l models.Lab
		err := rows.Scan(&l.ID, &l.UniversityID, &l.Name, &l.Slug, &l.Description, &l.Thumbnail, &l.Config, &l.CreatedAt, &l.UpdatedAt)
		if err != nil {
			return nil, err
		}
		labs = append(labs, l)
	}

	return labs, nil
}

func (s *MySQLStore) GetLabBySlug(universityID int64, slug string) (*models.LabWithExperiments, error) {
	query := `SELECT id, university_id, name, slug, description, thumbnail, config, created_at, updated_at 
	          FROM labs WHERE university_id = ? AND slug = ?`
	row := s.db.QueryRow(query, universityID, slug)

	var lab models.Lab
	err := row.Scan(&lab.ID, &lab.UniversityID, &lab.Name, &lab.Slug, &lab.Description, &lab.Thumbnail, &lab.Config, &lab.CreatedAt, &lab.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Get experiments for this lab
	experiments, err := s.GetExperimentsByLab(lab.ID)
	if err != nil {
		return nil, err
	}

	return &models.LabWithExperiments{
		Lab:         lab,
		Experiments: experiments,
	}, nil
}

// Experiment methods
func (s *MySQLStore) GetExperimentsByLab(labID int64) ([]models.Experiment, error) {
	query := `SELECT id, lab_id, university_id, title, slug, summary, difficulty, config, created_at, updated_at 
	          FROM experiments WHERE lab_id = ?`
	rows, err := s.db.Query(query, labID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiments []models.Experiment
	for rows.Next() {
		var e models.Experiment
		err := rows.Scan(&e.ID, &e.LabID, &e.UniversityID, &e.Title, &e.Slug, &e.Summary, &e.Difficulty, &e.Config, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}
		experiments = append(experiments, e)
	}

	return experiments, nil
}

func (s *MySQLStore) GetExperimentBySlug(labID int64, slug string) (*models.ExperimentWithApparatus, error) {
	query := `SELECT id, lab_id, university_id, title, slug, summary, difficulty, config, created_at, updated_at 
	          FROM experiments WHERE lab_id = ? AND slug = ?`
	row := s.db.QueryRow(query, labID, slug)

	var exp models.Experiment
	err := row.Scan(&exp.ID, &exp.LabID, &exp.UniversityID, &exp.Title, &exp.Slug, &exp.Summary, &exp.Difficulty, &exp.Config, &exp.CreatedAt, &exp.UpdatedAt)
	if err != nil {
		return nil, err
	}

	apparatus, err := s.GetApparatusByExperiment(exp.ID)
	if err != nil {
		return nil, err
	}

	return &models.ExperimentWithApparatus{
		Experiment: exp,
		Apparatus:  apparatus,
	}, nil
}

// Apparatus methods
func (s *MySQLStore) GetApparatusByExperiment(experimentID int64) ([]models.Apparatus, error) {
	query := `SELECT id, university_id, experiment_id, name, type, icon, model_path, meta, created_at, updated_at 
	          FROM apparatus WHERE experiment_id = ?`
	rows, err := s.db.Query(query, experimentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apparatus []models.Apparatus
	for rows.Next() {
		var a models.Apparatus
		err := rows.Scan(&a.ID, &a.UniversityID, &a.ExperimentID, &a.Name, &a.Type, &a.Icon, &a.ModelPath, &a.Meta, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}
		apparatus = append(apparatus, a)
	}

	return apparatus, nil
}

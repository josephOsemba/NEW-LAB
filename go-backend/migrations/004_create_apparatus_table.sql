CREATE TABLE apparatus (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    university_id BIGINT NOT NULL,
    experiment_id BIGINT,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(100),
    icon VARCHAR(512),
    model_path VARCHAR(1024),
    meta JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (university_id) REFERENCES universities(id) ON DELETE CASCADE,
    FOREIGN KEY (experiment_id) REFERENCES experiments(id) ON DELETE
    SET NULL,
        INDEX idx_university_id (university_id),
        INDEX idx_experiment_id (experiment_id)
);
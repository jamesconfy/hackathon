package repo

import (
	"database/sql"
	"project-name/internal/models"
)

type DepositRepo interface {
	Add(deposit *models.Deposit) (*models.Deposit, error)
	Get(depositId string) (*models.Deposit, error)
	GetAll() ([]*models.Deposit, error)
}

type depositRepo struct {
	conn *sql.DB
}

func (d *depositRepo) Add(deposit *models.Deposit) (*models.Deposit, error) {
	var dep models.Deposit

	query := `INSERT INTO deposits (back_image, front_image) VALUES($1, $2) RETURNING id, back_image, front_image, user_id, date_created, date_updated`

	err := d.conn.QueryRow(query, deposit.BackImage, deposit.FrontImage).Scan(&dep.Id, &dep.BackImage, &dep.FrontImage, &dep.UserId, &dep.DateCreated, &dep.DateUpdated)
	if err != nil {
		return nil, err
	}

	return &dep, nil
}

func (d *depositRepo) Get(depositId string) (*models.Deposit, error) {
	var dep models.Deposit

	query := `SELECT id, back_image, front_image, user_id, date_created, date_updated FROM deposits WHERE id = $1`

	err := d.conn.QueryRow(query, depositId).Scan(&dep.Id, &dep.BackImage, &dep.FrontImage, &dep.UserId, &dep.DateCreated, &dep.DateUpdated)
	if err != nil {
		return nil, err
	}

	return &dep, nil
}

func (d *depositRepo) GetAll() ([]*models.Deposit, error) {
	var deposits []*models.Deposit

	query := `SELECT id, back_image, front_image, user_id, date_created, date_updated FROM deposits`

	rows, err := d.conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dep models.Deposit

		err := rows.Scan(&dep.Id, &dep.BackImage, &dep.FrontImage, &dep.UserId, &dep.DateCreated, &dep.DateUpdated)
		if err != nil {
			return nil, err
		}

		deposits = append(deposits, &dep)
	}

	return deposits, nil
}

func NewDepositRepo(conn *sql.DB) DepositRepo {
	return &depositRepo{conn: conn}
}

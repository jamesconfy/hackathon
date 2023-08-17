package repo

import (
	"database/sql"
	"project-name/config"
	"project-name/internal/models"
)

type DepositRepo interface {
	Add(deposit *models.Deposit) (*models.Deposit, error)
	Get(depositId string) (*models.Deposit, error)
	GetAll() ([]*models.Deposit, error)
	Update(deposit *models.Deposit) (*models.Deposit, error)
}

type depositRepo struct {
	conn *sql.DB
}

func (d *depositRepo) Add(deposit *models.Deposit) (*models.Deposit, error) {
	var dep models.Deposit

	query := `INSERT INTO deposits (user_id, account_id, back_image, front_image) VALUES($1, $2, $3, $4) RETURNING id, back_image, front_image, status, date_created, date_updated`

	userId := config.AppConfig.DEFAULT_USER_ID
	accountId := config.AppConfig.DEFAULT_ACCOUNT_ID

	err := d.conn.QueryRow(query, userId, accountId, deposit.BackImage, deposit.FrontImage).Scan(&dep.Id, &dep.BackImage, &dep.FrontImage, &dep.Status, &dep.DateCreated, &dep.DateUpdated)
	if err != nil {
		return nil, err
	}

	// fmt.Println("Deposit: ", &dep)
	return d.Get(dep.Id)
}

func (d *depositRepo) Get(depositId string) (*models.Deposit, error) {
	var user models.User
	var dep models.Deposit

	query := `SELECT u.id, u.first_name, u.last_name, u.date_created, u.date_updated, d.id, d.back_image, d.front_image, d.status, d.date_created, d.date_updated, a.number FROM deposits d INNER JOIN users u ON u.id = d.user_id INNER JOIN accounts a ON a.id = d.account_id WHERE d.id = $1`

	err := d.conn.QueryRow(query, depositId).Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.DateUpdated, &dep.Id, &dep.BackImage, &dep.FrontImage, &dep.Status, &dep.DateCreated, &dep.DateUpdated, &dep.AccountNumber)
	if err != nil {
		return nil, err
	}

	user.GetFullName()
	dep.User = &user
	return &dep, nil
}

func (d *depositRepo) GetAll() ([]*models.Deposit, error) {
	var deposits []*models.Deposit

	query := `SELECT u.id, u.first_name, u.last_name, u.date_created, u.date_updated, d.id, d.back_image, d.front_image, d.status, d.date_created, d.date_updated, a.number FROM deposits d INNER JOIN users u ON u.id = d.user_id INNER JOIN accounts a ON a.id = d.account_id ORDER BY d.date_created desc`

	rows, err := d.conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var dep models.Deposit
		var user models.User

		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateCreated, &user.DateUpdated, &dep.Id, &dep.BackImage, &dep.FrontImage, &dep.Status, &dep.DateCreated, &dep.DateUpdated, &dep.AccountNumber)
		if err != nil {
			return nil, err
		}

		user.GetFullName()
		dep.User = &user
		deposits = append(deposits, &dep)
	}

	return deposits, nil
}

// Update implements DepositRepo.
func (d *depositRepo) Update(deposit *models.Deposit) (*models.Deposit, error) {
	var id string

	query := `UPDATE deposits SET status = $1, date_updated = CURRENT_TIMESTAMP WHERE id = $2 RETURNING id`

	err := d.conn.QueryRow(query, deposit.Status, deposit.Id).Scan(&id)
	if err != nil {
		return nil, err
	}

	return d.Get(id)
}

func NewDepositRepo(conn *sql.DB) DepositRepo {
	return &depositRepo{conn: conn}
}

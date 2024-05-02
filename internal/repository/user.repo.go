package repository

import (
	"coffeeshop/config"
	"coffeeshop/internal/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type RepoUserIF interface {
	CreateUser(data *models.User) (*config.Result, error)
	FetchUser(id string) (*config.Result, error)
	UpdateUser(id string, data *models.User) (*config.Result, error)
	RemoveUser(id string) (*config.Result, error)
}
type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

// Create User
func (r *RepoUser) CreateUser(data *models.User) (*config.Result, error) {
	q := `
	INSERT INTO users(
		photo_profile,
		email,
		password,
		role,
		address,
		phone_number,
		display_name,
		first_name,
		last_name,
		birthday,
		gender	
	)
	VALUES(
		:photo_profile,
		:email,
		:password,
		:role,
		:address,
		:phone_number,
		:display_name,
		:first_name,
		:last_name,
		:birthday,
		:gender
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return nil, err
	}

	return &config.Result{Message: "1 data user created"}, nil
}

// Get User
func (r *RepoUser) FetchUser(id string) (*config.Result, error) {
	q := "SELECT * FROM users WHERE user_id = ?"
	data := models.User{}

	if err := r.Get(&data, r.Rebind(q), id); err != nil {
		return nil, err
	}

	return &config.Result{Data: data}, nil
}

// Update User
func (r *RepoUser) UpdateUser(id string, data *models.User) (*config.Result, error) {
	q := `
	UPDATE users
	SET 
    photo_profile = COALESCE(NULLIF(:photo_profile, ''), photo_profile),
    email = COALESCE(NULLIF(:email, ''), email),
    password = COALESCE(NULLIF(:password, ''), password),
		role = COALESCE(NULLIF(:role, ''), role),
		address = COALESCE(NULLIF(:address, ''), address),
    phone_number = COALESCE(NULLIF(:phone_number, ''), phone_number),
    display_name = COALESCE(NULLIF(:display_name, ''), display_name),
    first_name = COALESCE(NULLIF(:first_name, ''), first_name),
    last_name = COALESCE(NULLIF(:last_name, ''), last_name),
    birthday = COALESCE(CAST(NULLIF(:birthday, '') AS DATE), birthday),
    gender = COALESCE(NULLIF(:gender, ''), gender),
		updated_at = NOW()
	WHERE
    user_id = :user_id
`
	data.User_id = id
	result, err := r.NamedExec(q, data)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no user was updated")
	}

	return &config.Result{Message: "1 data user updated"}, nil
}

// Delete User
func (r *RepoUser) RemoveUser(id string) (*config.Result, error) {
	q := `
	DELETE FROM users
	WHERE user_id = $1
	`

	result, err := r.Exec(q, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no user was deleted")
	}

	return &config.Result{Message: "1 data user deleted"}, nil
}

// Authentication
func (r *RepoUser) GetAuthData(email string) (*models.User, error) {
	result := models.User{}
	q := "SELECT user_id, password, role FROM users WHERE email = ?"

	if err := r.Get(&result, r.Rebind(q), email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("email not found")
		}

		return nil, err
	}

	return &result, nil
}

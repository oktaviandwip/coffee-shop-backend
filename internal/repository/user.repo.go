package repository

import (
	"coffee/internal/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

// Create User
func (r *RepoUser) CreateUser(data *models.User) (string, error) {
	q := `INSERT INTO public.users(
		photo_profile,
		email,
		password,
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
		return "", err
	}

	return "1 data user created", nil
}

// Get User
func (r *RepoUser) ReadUser(offset int) ([]*models.User, error) {
	q := `SELECT * FROM public.users LIMIT $1 OFFSET $2 `

	limit := 10
	rows, err := r.Queryx(q, limit, offset)
	if err != nil {
		return nil, err
	}

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Update User
func (r *RepoUser) UpdateUser(id string, data *models.User) (string, error) {
	q := `
	UPDATE public.users
	SET 
    photo_profile = COALESCE(NULLIF(:photo_profile, ''), photo_profile),
    email = COALESCE(NULLIF(:email, ''), email),
    password = COALESCE(NULLIF(:password, ''), password),
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
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", errors.New("no user was updated")
	}

	return "1 data user updated", nil
}

// Delete User
func (r *RepoUser) RemoveUser(id string, data *models.User) (string, error) {
	q := `
	DELETE FROM public.users
	WHERE
		user_id = :user_id
`
	data.User_id = id
	result, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", errors.New("no user was deleted")
	}

	return "1 data user deleted", nil
}

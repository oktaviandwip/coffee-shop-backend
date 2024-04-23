package repository

import (
	"coffee/internal/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type RepoFavorite struct {
	*sqlx.DB
}

func NewFavorite(db *sqlx.DB) *RepoFavorite {
	return &RepoFavorite{db}
}

// Create Favorite
func (r *RepoFavorite) CreateFavorite(data *models.Favorite) (string, error) {
	q := `INSERT INTO public.favorites(
				user_id,
				product_id
			)
			VALUES(
				:user_id,
				:product_id
			)`

	_, err := r.NamedExec(
		q, data)
	if err != nil {
		return "", err
	}

	return "1 data favorite created", nil
}

// Get Favorite
func (r *RepoFavorite) ReadFavorite(offset int) ([]*models.Favorite, error) {
	q := `SELECT * FROM public.favorites LIMIT $1 OFFSET $2`

	limit := 10
	rows, err := r.Queryx(q, limit, offset)
	if err != nil {
		return nil, err
	}

	var favorites []*models.Favorite
	for rows.Next() {
		var favorite models.Favorite
		if err := rows.StructScan(&favorite); err != nil {
			return nil, err
		}

		favorites = append(favorites, &favorite)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return favorites, nil
}

// Update Favorite
func (r *RepoFavorite) UpdateFavorite(user_id, product_id string, data *models.Favorite) (string, error) {
	q := `
		UPDATE public.favorites
		SET 
			user_id = COALESCE(CAST(NULLIF($1, '') AS UUID), user_id),
			product_id = COALESCE(CAST(NULLIF($2, '') AS UUID), product_id),
			updated_at = NOW()
		WHERE
			user_id = $3 AND product_id = $4
	`

	result, err := r.Exec(
		q,
		data.User_id,
		data.Product_id,
		user_id,
		product_id,
	)

	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", errors.New("no user was updated ")
	}

	return "1 data favorite updated", nil
}

// Delete Favorite
func (r *RepoFavorite) RemoveFavorite(userId, productId string, data *models.Favorite) (string, error) {
	q := `
	DELETE FROM public.favorites
	WHERE
		user_id = :user_id AND product_id = :product_id
`
	data.User_id = userId
	data.Product_id = productId

	result, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", errors.New("no user was deleted ")
	}

	return "1 data favorite deleted", nil
}

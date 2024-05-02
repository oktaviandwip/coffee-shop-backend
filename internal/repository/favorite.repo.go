package repository

import (
	"coffeeshop/config"
	"coffeeshop/internal/models"
	"errors"
	"math"

	"github.com/jmoiron/sqlx"
)

type RepoFavoriteIF interface {
	CreateFavorite(data *models.Favorite) (*config.Result, error)
	FetchFavorite(user_id string, page, offset int) (*config.Result, error)
	UpdateFavorite(user_id, product_id string, data *models.Favorite) (*config.Result, error)
	RemoveFavorite(userId, productId string) (*config.Result, error)
}
type RepoFavorite struct {
	*sqlx.DB
}

func NewFavorite(db *sqlx.DB) *RepoFavorite {
	return &RepoFavorite{db}
}

// Create Favorite
func (r *RepoFavorite) CreateFavorite(data *models.Favorite) (*config.Result, error) {
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
		return nil, err
	}

	return &config.Result{Message: "1 data favorite created"}, nil
}

// Get Favorite
func (r *RepoFavorite) FetchFavorite(user_id string, page, offset int) (*config.Result, error) {
	q := "SELECT * FROM favorites WHERE user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?"
	data := models.Favorites{}
	limit := 10

	if err := r.Select(&data, r.Rebind(q), user_id, limit, offset); err != nil {
		return nil, err
	}

	// Meta Data
	var count int
	var metas config.Metas

	m := "SELECT COUNT(user_id) as count FROM favorites WHERE user_id = ?"
	err := r.Get(&count, r.Rebind(m), user_id)
	if err != nil {
		return nil, err
	}

	check := math.Ceil(float64(count) / float64(10))
	metas.Total = count
	if count > 0 && page != int(check) {
		metas.Next = page + 1
	}

	if page != 1 {
		metas.Prev = page - 1
	}

	return &config.Result{Data: data, Meta: metas}, nil
}

// Update Favorite
func (r *RepoFavorite) UpdateFavorite(user_id, product_id string, data *models.Favorite) (*config.Result, error) {
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
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no user was updated ")
	}

	return &config.Result{Message: "1 data favorite updated"}, nil
}

// Delete Favorite
func (r *RepoFavorite) RemoveFavorite(userId, productId string) (*config.Result, error) {
	q := `
	DELETE FROM public.favorites
	WHERE user_id = $1 AND product_id = $2
	`
	result, err := r.Exec(q, userId, productId)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no user was deleted ")
	}

	return &config.Result{Message: "1 data favorite deleted"}, nil
}

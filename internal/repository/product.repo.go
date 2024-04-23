package repository

import (
	"coffee/internal/models"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

// Create Product
func (r *RepoProduct) CreateProduct(data *models.Product) (string, error) {
	q := `INSERT INTO public.products(
							photo_product,
							product_name,
							price,
							description,
							size,
							delivery_method,
							start_hour,
							end_hour,
							stock,
							product_type
			)
			VALUES(
							$1,
							$2,
							$3,
							$4,
							$5,
							$6,
							$7,
							$8,
							$9,
							$10
			)`

	size := pq.Array(data.Size)
	deliveryMethod := pq.Array(data.Delivery_method)

	_, err := r.Exec(
		q,
		data.Photo_product,
		data.Product_name,
		data.Price,
		data.Description,
		size,
		deliveryMethod,
		data.Start_hour,
		data.End_hour,
		data.Stock,
		data.Product_type,
	)
	if err != nil {
		return "", err
	}

	return "1 data product created", nil
}

// Get Product
func (r *RepoProduct) ReadProduct(offset int) ([]*models.Product, error) {
	q := `SELECT * FROM public.products LIMIT $1 OFFSET $2`

	limit := 10
	rows, err := r.Queryx(q, limit, offset)
	if err != nil {
		return nil, err
	}

	var products []*models.Product
	for rows.Next() {
		var product models.Product
		var size string
		var deliveryMethod string

		if err := rows.Scan(
			&product.Product_id,
			&product.Photo_product,
			&product.Product_name,
			&product.Price,
			&product.Description,
			&size,
			&deliveryMethod,
			&product.Start_hour,
			&product.End_hour,
			&product.Stock,
			&product.Product_type,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}

		size = strings.ReplaceAll(size, "{", "")
		size = strings.ReplaceAll(size, "}", "")
		size = strings.ReplaceAll(size, "\"", "")

		deliveryMethod = strings.ReplaceAll(deliveryMethod, "{", "")
		deliveryMethod = strings.ReplaceAll(deliveryMethod, "}", "")
		deliveryMethod = strings.ReplaceAll(deliveryMethod, "\"", "")

		product.Size = strings.Split(size, ",")
		product.Delivery_method = strings.Split(deliveryMethod, ",")

		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// Search Product
func (r *RepoProduct) SearchProduct(searchStr string, offset int) ([]*models.Product, error) {
	q := `SELECT * FROM public.products WHERE product_name ILIKE $1 LIMIT $2 OFFSET $3`

	search := "%" + searchStr + "%"
	limit := 10

	rows, err := r.Queryx(q, search, limit, offset)
	if err != nil {
		return nil, err
	}

	var products []*models.Product
	for rows.Next() {
		var product models.Product
		var size string
		var deliveryMethod string

		if err := rows.Scan(
			&product.Product_id,
			&product.Photo_product,
			&product.Product_name,
			&product.Price,
			&product.Description,
			&size,
			&deliveryMethod,
			&product.Start_hour,
			&product.End_hour,
			&product.Stock,
			&product.Product_type,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}

		size = strings.ReplaceAll(size, "{", "")
		size = strings.ReplaceAll(size, "}", "")
		size = strings.ReplaceAll(size, "\"", "")

		deliveryMethod = strings.ReplaceAll(deliveryMethod, "{", "")
		deliveryMethod = strings.ReplaceAll(deliveryMethod, "}", "")
		deliveryMethod = strings.ReplaceAll(deliveryMethod, "\"", "")

		product.Size = strings.Split(size, ",")
		product.Delivery_method = strings.Split(deliveryMethod, ",")

		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// Sort Product
func (r *RepoProduct) SortProduct(sortStr string, offset int) ([]*models.Product, error) {
	q := `SELECT * FROM public.products WHERE product_type = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`

	limit := 10
	rows, err := r.Queryx(q, sortStr, limit, offset)
	if err != nil {
		return nil, err
	}

	var products []*models.Product
	for rows.Next() {
		var product models.Product
		var size string
		var deliveryMethod string

		if err := rows.Scan(
			&product.Product_id,
			&product.Photo_product,
			&product.Product_name,
			&product.Price,
			&product.Description,
			&size,
			&deliveryMethod,
			&product.Start_hour,
			&product.End_hour,
			&product.Stock,
			&product.Product_type,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			return nil, err
		}

		size = strings.ReplaceAll(size, "{", "")
		size = strings.ReplaceAll(size, "}", "")
		size = strings.ReplaceAll(size, "\"", "")

		deliveryMethod = strings.ReplaceAll(deliveryMethod, "{", "")
		deliveryMethod = strings.ReplaceAll(deliveryMethod, "}", "")
		deliveryMethod = strings.ReplaceAll(deliveryMethod, "\"", "")

		product.Size = strings.Split(size, ",")
		product.Delivery_method = strings.Split(deliveryMethod, ",")

		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// Update Product
func (r *RepoProduct) UpdateProduct(id string, data *models.Product) (string, error) {
	q := `
		UPDATE public.products
		SET 
			photo_product = COALESCE(NULLIF($1, ''), photo_product),
			product_name = COALESCE(NULLIF($2, ''), product_name),
			price = COALESCE(NULLIF($3, 0), price),
			description = COALESCE(NULLIF($4, ''), description),
			size = COALESCE(NULLIF($5, '{}'::varchar[]), size),
			start_hour = COALESCE(CAST(NULLIF($6, '') AS TIME), start_hour),
			end_hour = COALESCE(CAST(NULLIF($7, '') AS TIME), end_hour),
			stock = COALESCE(NULLIF($8, 0), stock),
			product_type = COALESCE(NULLIF($9, ''), product_type),
			updated_at = NOW()
		WHERE
			product_id = $10
	`

	result, err := r.Exec(q,
		data.Photo_product,
		data.Product_name,
		data.Price,
		data.Description,
		pq.Array(data.Size),
		data.Start_hour,
		data.End_hour,
		data.Stock,
		data.Product_type,
		id,
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

	return "1 data product updated", nil
}

// Delete Product
func (r *RepoProduct) RemoveProduct(id string, data *models.Product) (string, error) {
	q := `
	DELETE FROM public.products
	WHERE
		product_id = :product_id
`
	data.Product_id = id
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

	return "1 data product deleted", nil
}

package repository

import (
	"coffeeshop/config"
	"coffeeshop/internal/models"
	"errors"
	"math"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type RepoProductIF interface {
	CreateProduct(data *models.Product) (*config.Result, error)
	FetchProduct(page, offset int) (*config.Result, error)
	SearchProduct(searchStr string, page, offset int) (*config.Result, error)
	SortProduct(sortStr string, page, offset int) (*config.Result, error)
	UpdateProduct(id string, data *models.Product) (*config.Result, error)
	RemoveProduct(id string) (*config.Result, error)
}
type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

// Create Product
func (r *RepoProduct) CreateProduct(data *models.Product) (*config.Result, error) {
	q := `
			INSERT INTO products(
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

	_, err := r.Exec(
		q,
		data.Photo_product,
		data.Product_name,
		data.Price,
		data.Description,
		pq.Array(data.Size),
		pq.Array(data.Delivery_method),
		data.Start_hour,
		data.End_hour,
		data.Stock,
		data.Product_type,
	)
	if err != nil {
		return nil, err
	}

	return &config.Result{Message: "1 data product created"}, nil
}

// Get Product
func (r *RepoProduct) FetchProduct(page, offset int) (*config.Result, error) {

	// All Products
	q := "SELECT * FROM products ORDER BY product_id LIMIT $1 OFFSET $2"
	limit := 10

	rows, err := r.Queryx(q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := models.Products{}

	for rows.Next() {
		product := models.Product{}
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

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Meta Data
	var count int
	var metas config.Metas

	m := "SELECT COUNT(product_id) as count FROM products"
	err = r.Get(&count, r.Rebind(m))
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

	return &config.Result{Data: products, Meta: metas}, nil
}

// Search Product
func (r *RepoProduct) SearchProduct(searchStr string, page, offset int) (*config.Result, error) {
	q := "SELECT * FROM products WHERE product_name ILIKE $1 ORDER BY product_name ASC LIMIT $2 OFFSET $3"
	search := "%" + searchStr + "%"
	limit := 10

	rows, err := r.Queryx(q, search, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := models.Products{}

	for rows.Next() {
		product := models.Product{}
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

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Meta Data
	var count int
	var metas config.Metas

	m := "SELECT COUNT(product_id) as count FROM products WHERE product_name ILIKE ?"
	err = r.Get(&count, r.Rebind(m), search)
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

	return &config.Result{Data: products, Meta: metas}, nil
}

// Sort Product
func (r *RepoProduct) SortProduct(sortStr string, page, offset int) (*config.Result, error) {
	q := "SELECT * FROM products WHERE product_type = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3"
	limit := 10

	rows, err := r.Queryx(q, sortStr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := models.Products{}

	for rows.Next() {
		product := models.Product{}
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

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Meta Data
	var count int
	var metas config.Metas

	m := "SELECT COUNT(product_id) as count FROM products WHERE product_type = ?"
	err = r.Get(&count, r.Rebind(m), sortStr)
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

	return &config.Result{Data: products, Meta: metas}, nil
}

// Update Product
func (r *RepoProduct) UpdateProduct(id string, data *models.Product) (*config.Result, error) {
	q := `
		UPDATE products
		SET 
			photo_product = COALESCE(NULLIF($1, ''), photo_product),
			product_name = COALESCE(NULLIF($2, ''), product_name),
			price = COALESCE(NULLIF($3, 0), price),
			description = COALESCE(NULLIF($4, ''), description),
			size = COALESCE(NULLIF($5, null)::varchar[], size),
			delivery_method = COALESCE(NULLIF($6, null)::varchar[], delivery_method),
			start_hour = COALESCE(CAST(NULLIF($7, '') AS TIME), start_hour),
			end_hour = COALESCE(CAST(NULLIF($8, '') AS TIME), end_hour),
			stock = COALESCE(NULLIF($9, 0), stock),
			product_type = COALESCE(NULLIF($10, ''), product_type),
			updated_at = NOW()
		WHERE
			product_id = $11
	`

	result, err := r.Exec(q,
		data.Photo_product,
		data.Product_name,
		data.Price,
		data.Description,
		pq.Array(data.Size),
		pq.Array(data.Delivery_method),
		data.Start_hour,
		data.End_hour,
		data.Stock,
		data.Product_type,
		id,
	)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no product was updated ")
	}

	return &config.Result{Message: "1 data product updated"}, nil
}

// Delete Product
func (r *RepoProduct) RemoveProduct(id string) (*config.Result, error) {
	q := `
	DELETE FROM products
	WHERE product_id = $1
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
		return nil, errors.New("no product was deleted")
	}

	return &config.Result{Message: "1 product deleted"}, nil
}

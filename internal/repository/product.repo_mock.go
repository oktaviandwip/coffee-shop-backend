package repository

import (
	"coffeeshop/config"
	"coffeeshop/internal/models"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) CreateProduct(data *models.Product) (*config.Result, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) FetchProduct(page, offset int) (*config.Result, error) {
	args := r.Mock.Called(page, offset)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) SearchProduct(searchStr string, page, offset int) (*config.Result, error) {
	args := r.Mock.Called(searchStr, page, offset)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) SortProduct(sortStr string, page, offset int) (*config.Result, error) {
	args := r.Mock.Called(sortStr, page, offset)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) UpdateProduct(id string, data *models.Product) (*config.Result, error) {
	args := r.Mock.Called(id, data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) RemoveProduct(id string) (*config.Result, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(*config.Result), args.Error(1)
}

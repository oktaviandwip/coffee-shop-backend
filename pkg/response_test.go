package pkg

import (
	"coffeeshop/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRes(t *testing.T) {
	t.Run("HTTP 200 OK", func(t *testing.T) {
		data := &config.Result{
			Data: "some data",
		}
		response := NewRes(200, data)
		assert.NotNil(t, response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "OK", response.Status)
		assert.Equal(t, "some data", response.Data)
		assert.Nil(t, response.Meta)
		assert.Nil(t, response.Description)
	})

	t.Run("HTTP 400 Bad Request", func(t *testing.T) {
		data := &config.Result{
			Data: "error message",
		}
		response := NewRes(400, data)
		assert.NotNil(t, response)
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "Bad Request", response.Status)
		assert.Equal(t, "error message", response.Description)
		assert.Nil(t, response.Meta)
	})

	t.Run("HTTP 404 Not Found", func(t *testing.T) {
		data := &config.Result{
			Data: "data not found",
		}
		response := NewRes(404, data)
		assert.NotNil(t, response)
		assert.Equal(t, 404, response.Code)
		assert.Equal(t, "Not Found", response.Status)
		assert.Equal(t, "data not found", response.Description)
		assert.Nil(t, response.Meta)
	})

	t.Run("HTTP 201 Created", func(t *testing.T) {
		data := &config.Result{
			Data: "new data created",
		}
		response := NewRes(201, data)
		assert.NotNil(t, response)
		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "Created", response.Status)
		assert.Equal(t, "new data created", response.Data)
		assert.Nil(t, response.Meta)
	})
}

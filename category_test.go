package order_service

import (
	"app/api/models"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

var c int64

func TestFilm(t *testing.T) {
	c = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createCategory(t)
			fmt.Println(id)
			// DeleteCategory(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", c)
}

func createCategory(t *testing.T) int {
	response := &models.Category{}

	request := &models.CreateCategory{
		CategoryName: faker.Word(),
	}

	resp, err := PerformRequest(http.MethodPost, "/category", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.CategoryId
}

func updateCategory(t *testing.T, id string) int {
	response := &models.Category{}
	request := &models.UpdateCategory{
		CategoryId: 1,
		CategoryName: faker.Word(),
	}

	resp, err := PerformRequest(http.MethodPut, "/category/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.CategoryId
}

func DeleteCategory(t *testing.T, id int) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/category/%d", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}

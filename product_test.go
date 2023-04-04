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

var p int64

func TestProduct(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createoProduct(t)
			fmt.Println(id)
			Deleteproduct(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createoProduct(t *testing.T) int {
	response := &models.Product{}

	request := &models.CreateProduct{
		ProductName: faker.Word(),
		BrandId: 5,
		CategoryId: 5,
		ModelYear: 2023,
		ListPrice: 100,
	}

	resp, err := PerformRequest(http.MethodPost, "/product", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.ProductId
}

func Updateproduct(t *testing.T, id string) int {
	response := &models.Product{}
	request := &models.UpdateProduct{
		ProductId: 100,
		ProductName: faker.Word(),
		BrandId: 5,
		CategoryId: 5,
		ModelYear: 2023,
		ListPrice: 100,
	}

	resp, err := PerformRequest(http.MethodPut, "/product/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.ProductId
}

func Deleteproduct(t *testing.T, id int) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/product/%d", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}

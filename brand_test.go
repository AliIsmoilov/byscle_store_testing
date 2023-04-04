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

var s int64

func TestBrand(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createBrand(t)
			fmt.Println(id)
			DeleteBrand(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createBrand(t *testing.T) int {
	response := &models.Brand{}

	request := &models.CreateBrand{
		BrandName: faker.Word(),
	}

	resp, err := PerformRequest(http.MethodPost, "/brand", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.BrandId
}

func UpdateBrand(t *testing.T, id string) int {
	response := &models.Brand{}
	request := &models.UpdateBrand{
		BrandId: 1,
		BrandName: faker.Word(),
	}

	resp, err := PerformRequest(http.MethodPut, "/brand/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.BrandId
}

func DeleteBrand(t *testing.T, id int) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/brand/%d", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}

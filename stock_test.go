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

var sto int64

func TestStock(t *testing.T) {
	st = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createStock(t)
			fmt.Println(id)
			DeleteStock(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", st)
}

func createStock(t *testing.T) int {
	response := &models.Stock{}

	request := &models.CreateStock{
		StoreId: 3,
		ProductId: 35,
		Quantity: int(faker.Latitude()),
	}

	resp, err := PerformRequest(http.MethodPost, "/stock", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.StoreId
}

func UpdateStock(t *testing.T, id string) int {
	response := &models.Stock{}
	request := &models.UpdateStock{
		StoreId: 3,
		ProductId: 35,
		Quantity: int(faker.Latitude()),
	}

	resp, err := PerformRequest(http.MethodPut, "/stock/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.StoreId
}

func DeleteStock(t *testing.T, id int) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/stock/%d", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}

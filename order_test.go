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

var o int64

func TestOrder(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createorder(t)
			fmt.Println(id)
			DeleteOrder(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createorder(t *testing.T) int {
	response := &models.Order{}

	request := &models.CreateOrder{
		CustomerId: 1,
		OrderStatus: 1,
		OrderDate: faker.Date(),
		RequiredDate: faker.Date(),
		ShippedDate: faker.Date(),
		StoreId: 1,
		StaffId: 1,
	}

	resp, err := PerformRequest(http.MethodPost, "/order", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.OrderId
}

func Updateorder(t *testing.T, id string) int {
	response := &models.Order{}
	request := &models.UpdateOrder{
		OrderId: 122,
		CustomerId: 1,
		OrderStatus: 1,
		OrderDate: faker.Date(),
		RequiredDate: faker.Date(),
		ShippedDate: faker.Date(),
		StoreId: 1,
		StaffId: 1,
	}

	resp, err := PerformRequest(http.MethodPut, "/order/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.OrderId
}

func DeleteOrder(t *testing.T, id int) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/order/%d", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}

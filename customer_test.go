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

var cu int64

func TestCustomer(t *testing.T) {
	c = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createCustomer(t)
			fmt.Println(id)
			DeleteCustomer(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", c)
}

func createCustomer(t *testing.T) int {
	response := &models.Customer{}

	request := &models.CreateCustomer{
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Phone: faker.Phonenumber(),
		Email: faker.Email(),
		Street: faker.Word(),
		City: faker.Word(),
		State: faker.Word(),
		ZipCode: 200100,
	}

	resp, err := PerformRequest(http.MethodPost, "/customer", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.CustomerId
}

func updateCustomer(t *testing.T, id string) int {
	response := &models.Customer{}
	request := &models.UpdateCustomer{
		CustomerId: 1,
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Phone: faker.Phonenumber(),
		Email: faker.Email(),
		Street: faker.Word(),
		City: faker.Word(),
		State: faker.Word(),
		ZipCode: 200100,
	}

	resp, err := PerformRequest(http.MethodPut, "/category/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.CustomerId
}

func DeleteCustomer(t *testing.T, id int) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/customer/%d", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}

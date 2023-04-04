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

var str int64

func TestStore(t *testing.T) {
	str = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createStore(t)
			fmt.Println(id)
			// DeleteStore(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", st)
}

func createStore(t *testing.T) int {
	response := &models.Store{}

	request := &models.CreateStore{
		StoreId: 3,
		StoreName: faker.Word(),
		Phone: faker.Phonenumber(),
		Email: faker.Email(),
		Street: faker.Word(),
		City: faker.Word(),
		State: faker.Word(),
		ZipCode: faker.LONGITUDE,
	}

	resp, err := PerformRequest(http.MethodPost, "/store", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.StoreId
}

func UpdateStore(t *testing.T, id string) int {
	response := &models.Store{}
	request := &models.UpdateStore{
		StoreId: 3,
		StoreName: faker.Word(),
		Phone: faker.Phonenumber(),
		Email: faker.Email(),
		Street: faker.Word(),
		City: faker.Word(),
		State: faker.Word(),
		ZipCode: faker.LONGITUDE,
	}

	resp, err := PerformRequest(http.MethodPut, "/store/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.StoreId
}

func DeleteStore(t *testing.T, id int) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/store/%d", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}

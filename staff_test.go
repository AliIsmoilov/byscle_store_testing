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

var st int64

func TestStaff(t *testing.T) {
	st = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createStaff(t)
			fmt.Println(id)
			DeleteStaff(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", st)
}

func createStaff(t *testing.T) int {
	response := &models.Staff{}

	request := &models.CreateStaff{
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Email: faker.Email(),
		Phone: faker.Phonenumber(),
		Active: 1,
		StoreId: 2,
		ManagerId: 2,
	}

	resp, err := PerformRequest(http.MethodPost, "/staff", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.StaffId
}

func UpdateStaff(t *testing.T, id string) int {
	response := &models.Staff{}
	request := &models.UpdateStaff{
		StaffId: 7,
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Email: faker.Email(),
		Phone: faker.Phonenumber(),
		Active: 1,
		StoreId: 2,
		ManagerId: 2,
	}

	resp, err := PerformRequest(http.MethodPut, "/staff/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.StaffId
}

func DeleteStaff(t *testing.T, id int) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/staff/%d", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}

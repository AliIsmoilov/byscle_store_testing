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

var pr int64

func TestPromocode(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createPromoccode(t)
			fmt.Println(id)
			DeletePromocode(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", pr)
}

func createPromoccode(t *testing.T) int {
	response := &models.PromoCode{}

	request := &models.CreatePromoCode{
		Name: faker.WORD,
		Discount: 100,
		Discount_Type: "fixed",
		Ordred_limit_price: 500,
	}

	resp, err := PerformRequest(http.MethodPost, "/promocode", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.PromoCodeId
}

func UpdatePromocode(t *testing.T, id string) int {
	response := &models.PromoCode{}
	request := &models.UpdatePromoCode{
		PromoCodeId: 5,
		Name: faker.WORD,
		Discount: 100,
		Discount_Type: "fixed",
		Ordred_limit_price: 500,
	}

	resp, err := PerformRequest(http.MethodPut, "/promocode/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.PromoCodeId
}

func DeletePromocode(t *testing.T, id int) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/promocode/%d", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}

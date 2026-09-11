package orders

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeProductClient struct {
	products map[string]Product
}

func newFakeProductClient() *fakeProductClient {
	return &fakeProductClient{
		products: map[string]Product{
			"1": {ID: "1", Name: "Notebook", Price: 1200, Stock: 10},
			"2": {ID: "2", Name: "Mouse", Price: 25, Stock: 30},
			"3": {ID: "3", Name: "Keyboard", Price: 60, Stock: 20},
		},
	}
}

func (c *fakeProductClient) GetProduct(_ context.Context, id string) (Product, error) {
	product, ok := c.products[id]
	if !ok {
		return Product{}, ErrProductNotFound
	}

	return product, nil
}

func TestCreateOrderValid(t *testing.T) {
	service := NewService(NewRepository(), newFakeProductClient())
	handler := NewHandler(service)

	body := bytes.NewBufferString(`{"productId":"1","quantity":2}`)
	req := httptest.NewRequest(http.MethodPost, "/orders", body)
	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}
}

func TestCreateOrderInvalidQuantity(t *testing.T) {
	service := NewService(NewRepository(), newFakeProductClient())
	handler := NewHandler(service)

	body := bytes.NewBufferString(`{"productId":"1","quantity":0}`)
	req := httptest.NewRequest(http.MethodPost, "/orders", body)
	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

func TestCreateOrderInsufficientStock(t *testing.T) {
	service := NewService(NewRepository(), newFakeProductClient())
	handler := NewHandler(service)

	body := bytes.NewBufferString(`{"productId":"1","quantity":99}`)
	req := httptest.NewRequest(http.MethodPost, "/orders", body)
	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

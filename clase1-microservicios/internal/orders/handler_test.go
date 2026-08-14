package orders

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"clase1/starter/internal/products"
)

func TestCreateOrderValid(t *testing.T) {
	service := NewService(NewRepository(), products.NewRepository())
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
	service := NewService(NewRepository(), products.NewRepository())
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
	service := NewService(NewRepository(), products.NewRepository())
	handler := NewHandler(service)

	body := bytes.NewBufferString(`{"productId":"1","quantity":99}`)
	req := httptest.NewRequest(http.MethodPost, "/orders", body)
	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

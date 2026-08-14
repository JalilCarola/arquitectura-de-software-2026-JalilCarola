package products

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProductByID(t *testing.T) {
	service := NewService(NewRepository())
	handler := NewHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	handler.GetByID(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}

func TestGetProductByIDNotFound(t *testing.T) {
	service := NewService(NewRepository())
	handler := NewHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/products/999", nil)
	req.SetPathValue("id", "999")
	rec := httptest.NewRecorder()

	handler.GetByID(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rec.Code)
	}
}

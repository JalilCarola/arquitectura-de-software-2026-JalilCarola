package orders

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPProductClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewHTTPProductClient(baseURL string) *HTTPProductClient {
	return &HTTPProductClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *HTTPProductClient) GetProduct(ctx context.Context, id string) (Product, error) {
	url := fmt.Sprintf("%s/products/%s", c.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Product{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Product{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return Product{}, ErrProductNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return Product{}, fmt.Errorf("product service returned status %d", resp.StatusCode)
	}

	var product Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return Product{}, err
	}

	return product, nil
}

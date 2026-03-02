package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiOverviewHandler(t *testing.T) {
	t.Run("returns 200 and expected route keys on GET", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api", nil)
		w := httptest.NewRecorder()

		apiOverviewHandler(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", w.Code)
		}

		if ct := w.Header().Get("Content-Type"); ct != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", ct)
		}

		var routes map[string]string
		if err := json.NewDecoder(w.Body).Decode(&routes); err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}

		expectedKeys := []string{
			"create_user_url",
			"current_user_url",
			"user_token_url",
			"recipes_url",
			"ingredients_url",
			"tags_url",
		}
		for _, key := range expectedKeys {
			if _, ok := routes[key]; !ok {
				t.Errorf("expected route key %q missing from response", key)
			}
		}
	})

	t.Run("returns 405 on non-GET", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api", nil)
		w := httptest.NewRecorder()

		apiOverviewHandler(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status 405, got %d", w.Code)
		}
	})
}

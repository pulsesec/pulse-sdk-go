package pulse

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

const (
	testSiteKey   = "site_key"
	testSecretKey = "secret_key"
	testToken     = "token"
)

func ptr[T any](v T) *T {
	return &v
}

func TestClassify(t *testing.T) {
	client := New(testSiteKey, testSecretKey)

	t.Run("classify", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, urlClassify,
			func(req *http.Request) (*http.Response, error) {
				var payload classifyPayload
				if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
					return nil, err
				}

				if payload.Token != testToken {
					t.Errorf("Expected token %q, got %q", testToken, payload.Token)
				}

				if payload.SiteKey != testSiteKey {
					t.Errorf("Expected site key %q, got %q", testSiteKey, payload.SiteKey)
				}

				if payload.SecretKey != testSecretKey {
					t.Errorf("Expected secret key %q, got %q", testSecretKey, payload.SecretKey)
				}

				return httpmock.NewJsonResponse(http.StatusOK, classifyResponse{
					IsBot: ptr(true),
				})
			},
		)

		isBot, err := client.Classify("token")
		if err != nil {
			t.Fatalf("Failed to classify token: %v", err)
		}

		if !isBot {
			t.Error("Expected token to be classified as bot")
		}
	})
}

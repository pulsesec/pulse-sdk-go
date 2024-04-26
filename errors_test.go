package pulse

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestErrors(t *testing.T) {
	client := New(testSiteKey, testSecretKey)

	t.Run("error token used", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, urlClassify,
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewJsonResponse(http.StatusOK, classifyResponse{
					errorResponse: errorResponse{
						Errors: []Error{
							{
								Code:    "TOKEN_USED",
								Message: "Token has already been used",
							},
						},
					},
				})
			},
		)

		_, err := client.Classify("token")
		if !errors.Is(err, ErrTokenUsed) {
			t.Errorf("Expected error %q, got %v", ErrTokenUsed, err)
		}
	})

	t.Run("error unknown token", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, urlClassify,
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewJsonResponse(http.StatusOK, classifyResponse{
					errorResponse: errorResponse{
						Errors: []Error{
							{
								Code:    "TOKEN_NOT_FOUND",
								Message: "Token not found",
							},
						},
					},
				})
			},
		)

		_, err := client.Classify("token")
		if !errors.Is(err, ErrTokenNotFound) {
			t.Errorf("Expected error %q, got %v", ErrTokenUsed, err)
		}
	})
}

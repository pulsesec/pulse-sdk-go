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

		testError := Error{
			Code:    "TOKEN_USED",
			Message: "Token has already been used",
		}

		httpmock.RegisterResponder(http.MethodPost, urlClassify,
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewJsonResponse(http.StatusOK, classifyResponse{
					errorResponse: errorResponse{
						Errors: []Error{testError},
					},
				})
			},
		)

		_, err := client.Classify("token")
		if !errors.Is(err, ErrTokenUsed) {
			t.Errorf("Expected error %q, got %v", ErrTokenUsed, err)
		}

		var apiErr Error
		if !errors.As(err, &apiErr) {
			t.Errorf("Expected error to be of type %T", apiErr)
		}

		if apiErr != testError {
			t.Errorf("Expected error %v, got %v", testError, apiErr)
		}
	})

	t.Run("error unknown token", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		testError := Error{
			Code:    "TOKEN_NOT_FOUND",
			Message: "Token not found",
		}

		httpmock.RegisterResponder(http.MethodPost, urlClassify,
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewJsonResponse(http.StatusOK, classifyResponse{
					errorResponse: errorResponse{
						Errors: []Error{testError},
					},
				})
			},
		)

		_, err := client.Classify("token")
		if !errors.Is(err, ErrTokenNotFound) {
			t.Errorf("Expected error %q, got %v", ErrTokenUsed, err)
		}

		var apiErr Error
		if !errors.As(err, &apiErr) {
			t.Errorf("Expected error to be of type %T", apiErr)
		}

		if apiErr != testError {
			t.Errorf("Expected error %v, got %v", testError, apiErr)
		}
	})
}

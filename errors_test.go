package pulse

import (
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestErrors(t *testing.T) {
	client := New(testSiteKey, testSecretKey)

	t.Run("errors must be mapped", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		tests := []struct {
			Err   error
			Value Error
		}{
			{
				ErrTokenNotFound,
				Error{
					Code:    "TOKEN_NOT_FOUND",
					Message: "Token not found",
				},
			},
			{
				ErrTokenExpired,
				Error{
					Code:    "TOKEN_EXPIRED",
					Message: "Token has expired",
				},
			},
			{
				ErrTokenUsed,
				Error{
					Code:    "TOKEN_USED",
					Message: "Token has already been used",
				},
			},
		}

		for _, e := range tests {
			httpmock.RegisterResponder(http.MethodPost, urlClassify,
				func(req *http.Request) (*http.Response, error) {
					return httpmock.NewJsonResponse(http.StatusOK, classifyResponse{
						errorResponse: errorResponse{
							Errors: []Error{e.Value},
						},
					})
				},
			)

			_, err := client.Classify(testToken)
			if !errors.Is(err, e.Err) {
				t.Errorf("Expected error %q, got %v", e.Err, err)
			}

			var apiErr Error
			if !errors.As(err, &apiErr) {
				t.Errorf("Expected error to be of type %T", apiErr)
			}

			if apiErr != e.Value {
				t.Errorf("Expected error %v, got %v", e.Value, apiErr)
			}

			httpmock.Reset()
		}
	})
}

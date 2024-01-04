package handlers

import (
	"github.com/LobovVit/url-shortener/internal/app/config"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testRequest(t *testing.T, ts *httptest.Server, method,
	path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestUpdateHandler(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
	}
	tests := []struct {
		name  string
		metod string
		path  string
		body  string
		want  want
	}{
		{
			name:  "test1",
			metod: http.MethodPost,
			path:  "/",
			body:  "www.ya.ru",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  201,
			},
		},
	}

	config.GetConfig()
	mux := chi.NewRouter()
	mux.Post("/", SetShortHandler)
	ts := httptest.NewServer(mux)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, _ := testRequest(t, ts, tt.metod, tt.path, strings.NewReader(tt.body))
			assert.Equal(t, tt.want.statusCode, resp.StatusCode)
			assert.Equal(t, tt.want.contentType, resp.Header.Get("Content-Type"))
			resp.Body.Close()
		})
	}
}

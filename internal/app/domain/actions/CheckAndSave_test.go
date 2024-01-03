package actions

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCheckAndSave(t *testing.T) {
	tests := []struct {
		name string
		url  string
		tp   string
		code string
		val  string
		want int
	}{
		{
			name: "+ test #1",
			url:  "www.123.ee",
			tp:   "counter",
			code: "someMetric",
			val:  "527",
			want: http.StatusOK,
		}, {
			name: "+ test #2",
			url:  "www.123.ee",
			tp:   "counter",
			code: "someMetric",
			val:  "527",
			want: http.StatusOK,
		}, {
			name: "- test #3",
			url:  "www.123.ee",
			tp:   "counter",
			code: "someMetric",
			val:  "527",
			want: http.StatusOK,
		}, {
			name: "- test #4",
			url:  "www.123.ee",
			tp:   "counter",
			code: "someMetric",
			val:  "527",
			want: http.StatusOK,
		}, {
			name: "- test #5",
			url:  "www.123.ee",
			tp:   "counter",
			code: "someMetric",
			val:  "527",
			want: http.StatusOK,
		}, {
			name: "+ test #6",
			url:  "www.123.ee/update/gauge/someMetric/527",
			tp:   "gauge",
			code: "someMetric",
			val:  "527",
			want: http.StatusOK,
		}, {
			name: "+ test #7",
			url:  "www.123.ee/update/gauge/someMetric/527/",
			tp:   "gauge",
			code: "someMetric",
			val:  "527",
			want: http.StatusOK,
		}, {
			name: "- test #8",
			url:  "www.123.ee/update/gauge/some",
			tp:   "gauge",
			code: "someMetric",
			val:  "527",
			want: http.StatusOK,
		}, {
			name: "- test #9",
			url:  "www.123.ee/update/gauge/somete/gauge/some",
			tp:   "gauge",
			code: "someMetric",
			val:  "444",
			want: http.StatusOK,
		}, {
			name: "- test #10",
			url:  "www.123.ee/update/gauge/somemeMetric/hello",
			tp:   "gauge",
			code: "someMetric",
			val:  "444",
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CheckAndSave(tt.url)
			assert.NoError(t, err, tt.want)
		})
	}
}

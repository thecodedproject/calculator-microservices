package calculator_api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewServerAndRouter(t *testing.T) *httprouter.Router {
	r := httprouter.New()
	var err error
	_, err = New(r)
	require.NoError(t, err)
	return r
}

func NewJsonRequest(t *testing.T, method string, path string, i interface{}) *http.Request {
	buf, err := json.Marshal(i)
	require.NoError(t, err)
	req, err := http.NewRequest(method, path, bytes.NewBuffer(buf))
	require.NoError(t, err)
	return req
}

func TestAdd(t *testing.T) {

	type request struct {
		Inputs []float64 `json:"inputs"`
	}

	type resultResp struct {
		Value float64 `json:"value"`
	}

	type expectedResp struct {
		Ok bool `json:"ok"`
		Err string `json:"err"`
		Result resultResp `json:"result"`
	}

	tests := []struct {
		name string
		req interface{}
		status int
		expectedRes expectedResp
	}{
		{
			name: "Single input",
			req: request{
				Inputs: []float64{2.0},
			},
			status: http.StatusOK,
			expectedRes: expectedResp{
				Ok: true,
				Err: "",
				Result: resultResp{
					Value: 2.0,
				},
			},
		},
		{
			name: "Multiple inputs returns sum of all",
			req: request{
				Inputs: []float64{2.0, 3.2, 1.4, 5.0},
			},
			status: http.StatusOK,
			expectedRes: expectedResp{
				Ok: true,
				Err: "",
				Result: resultResp{
					Value: 11.6,
				},
			},
		},
		{
			name: "No inputs returns zero",
			req: request{
				Inputs: []float64{},
			},
			status: http.StatusOK,
			expectedRes: expectedResp{
				Ok: true,
				Err: "",
				Result: resultResp{
					Value: 0.0,
				},
			},
		},
		{
			name: "Missing inputs field returns 400",
			req: request{},
			status: http.StatusBadRequest,
			expectedRes: expectedResp{
				Ok: false,
				Err: "Missing inputs field",
			},
		},
		{
			name: "Unknown field returns error",
			req: struct {A int}{},
			status: http.StatusBadRequest,
			expectedRes: expectedResp{
				Ok: false,
				Err: "Request body contains unknown field \"A\"",
			},
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			req := NewJsonRequest(t, "POST", "/add", test.req)

			r := NewServerAndRouter(t)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, test.status, w.Code)

			var actual expectedResp
			err := json.Unmarshal(w.Body.Bytes(), &actual)
			require.NoError(t, err)

			assert.Equal(t, test.expectedRes, actual)
		})
	}
}

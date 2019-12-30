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

func TestAddWithSingleInput(t *testing.T) {

	in := struct {
		Inputs []float64 `json:"inputs"`
	}{
		Inputs: []float64{2.0},
	}
	buf, err := json.Marshal(in)
	require.NoError(t, err)

	r := httprouter.New()
	_, err = New(r)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(buf))
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type resultResp struct {
		Value float64 `json:"value"`
	}

	type expectedResp struct {
		Ok bool `json:"ok"`
		Err string `json:"err"`
		Result resultResp `json:"result"`
	}

	expected := expectedResp{
		Ok: true,
		Err: "",
		Result: resultResp{
			Value: 2.0,
		},
	}

	var actual expectedResp

	err = json.Unmarshal(w.Body.Bytes(), &actual)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}

package calculator_api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/thecodedproject/calculator_microservices/add"
	add_client"github.com/thecodedproject/calculator_microservices/add/client"
)

type Server struct {
	router *httprouter.Router

	addClient add.Client
}

func New(r *httprouter.Router) (*Server, error) {
	addClient, err := add_client.New()
	if err != nil {
		return nil, err
	}

	s := &Server{
		router: r,
		addClient: addClient,
	}
	s.routes()
	return s, nil
}

func (s *Server) handleAddPost() httprouter.Handle {
	type request struct {
		Inputs []float64 `json:"inputs"`
	}
	type result struct {
		Value float64 `json:"value"`
	}
	type response struct {
		Ok bool `json:"ok"`
		Err string `json:"err"`
		Result result `json:"result"`
	}
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		req := request{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			panic(err)
		}

		res, err := s.addClient.Calc(r.Context(), req.Inputs)

		json.NewEncoder(w).Encode(&response{
			Ok: true,
			Result: result{
				Value: res,
			},
		})

	}
}

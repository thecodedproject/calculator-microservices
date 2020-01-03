package calculator_api

import (
	"log"
	"encoding/json"
	"errors"
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

		var req request
		err := decodeJSONBody(w, r, &req)
		if err != nil {
			var mr *malformedRequest
			if errors.As(err, &mr) {

				msg, err := json.Marshal(&response{
					Ok:false,
					Err: mr.msg,
				})
				if err != nil {
					log.Println(err.Error())
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
				http.Error(w, string(msg), mr.status)
			} else {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		if req.Inputs == nil {
			msg, err := json.Marshal(&response{
				Ok:false,
				Err: "Missing inputs field",
			})
			if err != nil {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			http.Error(w, string(msg), http.StatusBadRequest)
			return
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

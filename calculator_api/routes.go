package calculator_api

func (s *Server) routes() {

	s.router.POST("/add", s.handleAddPost())

}

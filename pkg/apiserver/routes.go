package apiserver

func (s *server) configureRouter() {
	s.router.Handle("/api/v1/stats/usage", s.handleActivity()).Methods("GET", "OPTIONS")
	s.router.Handle("/api/v1/stats/usage/{from}/{to}", s.handleActivityPeriod()).Methods("GET", "OPTIONS")
	s.router.Handle("/api/v1/stats/usage/{from}", s.handleActivityPeriod()).Methods("GET", "OPTIONS")
}

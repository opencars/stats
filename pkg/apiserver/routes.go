package apiserver

func (s *server) configureRouter() {
	s.router.Handle("/api/v1/stats/usage", s.handleActivity())
	s.router.Handle("/api/v1/stats/usage/{from}/{to}", s.handleActivityPeriod())
}

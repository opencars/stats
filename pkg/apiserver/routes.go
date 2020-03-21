package apiserver

func (s *server) configureRouter() {
	s.router.Handle("/api/v1/stats/admin/authorizations/{token}", s.handleActivityByToken())
	s.router.Handle("/api/v1/stats/admin/authorizations/{from:[0-9]+}/{to:[0-9]+}", s.handleStatsForTimeInterval())
	s.router.Handle("/api/v1/stats/admin/authorizations/{duration}", s.handleStatsForDuration())
}

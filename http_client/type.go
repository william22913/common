package http_client

type APIConnector interface {
	HitAPI(
		method string,
		host string,
		path string,
		header map[string]string,
		body interface{},
		result interface{},

	) (
		int,
		error,
	)
}

package mql_request_filter

type EntityModel interface {
	GetFilterWhitelist() (map[string]bool, error)
}

package common
package common

// EmptyReq is a request with no fields
type EmptyReq struct{}

// HealthResp is the response for health check
type HealthResp struct {
	Status string `json:"status"`
}

// IndexResp is the response for index
type IndexResp struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

// PingResp is the response for ping
type PingResp struct {
	Message string `json:"message"`
}

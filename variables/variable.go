package variables

type Response struct {
	Error            error
	APIResponseArray []*APIResponse `json:"APIResponseArray,omitempty"`
	APIResponse      APIResponse    `json:"APIResponse,omitempty"`
	ChannelLogId     string
	HttpStatus       int
}
type APIResponse struct {
	Code    string `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type AddExoPlanetRequest struct {
	Name              string  `json:"name"binding:"required"`
	Description       string  `json:"description"binding:"required"`
	DistanceFromEarth int     `json:"DistanceFromEarth"binding:"required"`
	Mass              float64 `json:"mass,omitempty"`
	Radius            float64 `json:"radius"binding:"required"`
	ExoPlanetType     string  `json:"ExoPlanetType"binding:"required"`
}
type Request struct {
	Name string `json:"name"binding:"required"`
	Crew string `json:"crew,omitempty"`
}
type UpdateExoPlanetRequest struct {
	Name              string  `json:"name"binding:"required"`
	NewName           string  `json:"newName,omitempty"`
	Description       string  `json:"description,omitempty"`
	DistanceFromEarth int     `json:"DistanceFromEarth,omitempty"`
	Mass              float64 `json:"mass,omitempty"`
	Radius            float64 `json:"radius,omitempty"`
	ExoPlanetType     string  `json:"ExoPlanetType,omitempty"`
}
type FuelEstimation struct {
	DistanceFromEarth int     `json:"DistanceFromEarth"binding:"required"`
	Mass              float64 `json:"mass,omitempty"`
	Radius            float64 `json:"radius"binding:"required"`
	ExoPlanetType     string  `json:"ExoPlanetType"binding:"required"`
}

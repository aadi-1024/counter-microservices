package models

type JsonResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	Value   int    `json:"value,omitempty"`
}

package types

type Response struct {
	Messsage string `json:"message"`
	Key      string `json:"key,omitempty"`
}

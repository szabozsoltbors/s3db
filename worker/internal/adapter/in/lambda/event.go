package lambda

type Event struct {
	Key  string `json:"key"`
	Data []byte `json:"data"`
}

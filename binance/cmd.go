package v3

//	Необходим для удобного создания подписок

type Cmd struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int      `json:"id"`
}

package response

// codec is a Codec implementation
type resp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
	Status string `json:"status"`
}

// func Register() {
// 	tkit.Response = resp{}
// }

// func (r resp) Result() any {
// 	return r
// }

// func (r resp) Name() string {
// 	return "default"
// }

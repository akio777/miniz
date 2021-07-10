package response

type Response struct {
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"Data" xml:"Data"`
}

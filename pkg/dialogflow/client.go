package dialogflow

const BaseURL = "https://api.dialogflow.com/v1"

//Client represent a DialogFlow client
type Client struct {
	Url string
	Token string
}

//NewClient create and return a simple Dialogflow client
func NewClient(token string) *Client {
	c := Client{
		Url: BaseURL,
		Token: token,
	}

	return &c
}
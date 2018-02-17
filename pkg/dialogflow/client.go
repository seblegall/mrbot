package dialogflow

const BaseURL = "https://api.dialogflow.com/v1"

type Client struct {
	Url string
	Token string
}

func NewClient(token string) *Client {
	c := Client{
		Url: BaseURL,
		Token: token,
	}

	return &c
}
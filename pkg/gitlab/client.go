package gitlab

import (
	gitlab "github.com/xanzy/go-gitlab"
)

//Client is a gitlab client
type Client struct {
	*gitlab.Client
}

//NewClient creates a new gitlab client that statisfied the same interface
func NewClient(baseURL, token string) *Client {

	git := gitlab.NewClient(nil, token)
	git.SetBaseURL(baseURL)

	return &Client{git}
}

//Stream start a new stream that listen message containing
//a reference (using "@") to the mentionname specified
func (c *Client) Stream(groups []string) (stream *MergeRequestStream) {
	return c.newMergeRequestStream(groups)
}

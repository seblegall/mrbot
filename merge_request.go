package main

import (
	"fmt"
	"time"
)

//ListenMergeRequest publish a message for each new gitlab merge request created
func (b *Bot) ListenMergeRequest(groups []string) {

	go func(b *Bot) {
		stream := b.gitlab.Stream(groups)

		for mergeRequests := range stream.C {
			for _, mr := range mergeRequests {
				b.room.Send(fmt.Sprintf("New merge request !\nOpened by : %s\nOn project : %s\nAt : %s\nTitle:%s\nSee : %s", mr.CreatedBy, mr.Project, mr.Title, mr.CreatedAt.Format(time.RFC3339), mr.URL))
			}
		}
	}(b)

}

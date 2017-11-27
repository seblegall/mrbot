package gitlab

import (
	"log"
	"sync"
	"time"

	gitlab "github.com/xanzy/go-gitlab"
)

const waitTime = 60

//MergeRequest represents a gitlab Merge Request.
type MergeRequest struct {
	Project   string
	Title     string
	CreatedAt *time.Time
	CreatedBy string
	URL       string
}

//MergeRequests represents a slice of MergeRequest
type MergeRequests []MergeRequest

//MergeRequestStream represents a stream of MergeRequest.
//new open MergeRequest are sent to the channel C.
type MergeRequestStream struct {
	client *Client
	groups []string
	from   time.Time
	C      chan MergeRequests
	run    bool
}

func (c *Client) newMergeRequestStream(groups []string) *MergeRequestStream {
	stream := MergeRequestStream{
		client: c,
		groups: groups,
		from:   time.Now(),
		C:      make(chan MergeRequests),
	}

	stream.start()
	return &stream

}

func (s *MergeRequestStream) start() {
	s.run = true
	go s.loop()
}

func (s *MergeRequestStream) stop() {
	s.run = false
}

func (s *MergeRequestStream) loop() {
	defer close(s.C)
	for s.run {
		mr := s.GetOpenMergeRequests()
		if mr != nil {
			s.C <- mr
		}
		time.Sleep(waitTime * time.Second)
		s.from = s.from.Add(waitTime * time.Second)
	}
}

//GetOpenMergeRequests returns a list a open merge request from groups passed as parameter.
//This function return only merge request created after a "from" parameters.
func (s *MergeRequestStream) GetOpenMergeRequests() MergeRequests {

	var results MergeRequests
	groups := s.GetGroups()
	projectsList := s.GetProjects(groups)

	pageOptions := gitlab.ListOptions{
		PerPage: 100,
	}

	mrOptions := &gitlab.ListMergeRequestsOptions{
		ListOptions: pageOptions,
		State:       gitlab.String("opened"),
		OrderBy:     gitlab.String("created_at"),
		Scope:       gitlab.String("all"),
	}

	mergeRequests, _, err := s.client.Client.MergeRequests.ListMergeRequests(mrOptions)

	if err != nil {
		log.Fatal(err)
	}

	for _, mr := range mergeRequests {
		if p, ok := projectsList[mr.ProjectID]; ok {
			if mr.CreatedAt.After(s.from) {
				results = append(results, MergeRequest{
					Project:   p.Name,
					Title:     mr.Title,
					CreatedAt: mr.CreatedAt,
					CreatedBy: mr.Author.Name,
					URL:       mr.WebURL,
				})
			}
		}
	}

	return results
}

//GetGroups search groups sepcified as parameter and return found groups as a slice.
func (s *MergeRequestStream) GetGroups() []*gitlab.Group {

	var groups []*gitlab.Group
	var wgg sync.WaitGroup
	mu := &sync.Mutex{}

	for _, gtg := range s.groups {
		wgg.Add(1)
		go func(gtg string, mu *sync.Mutex) {
			gs, _, err := s.client.Client.Groups.ListGroups(&gitlab.ListGroupsOptions{
				Search: gitlab.String(gtg),
			})
			if err != nil {
				log.Fatal(err)
			}

			for _, g := range gs {
				mu.Lock()
				groups = append(groups, g)
				mu.Unlock()
			}
			wgg.Done()
		}(gtg, mu)
	}

	wgg.Wait()

	return groups
}

//GetProjects find all projects from a given slice of groups passed as parameter.
//This function returns a map of project where map keys are the corresponding project's ID
func (s *MergeRequestStream) GetProjects(groups []*gitlab.Group) map[int]*gitlab.Project {
	var wgp sync.WaitGroup
	projectsList := make(map[int]*gitlab.Project)
	mu := &sync.Mutex{}

	for _, g := range groups {
		wgp.Add(1)

		go func(g *gitlab.Group, mu *sync.Mutex) {
			//Find all projects
			projects, _, err := s.client.Client.Groups.ListGroupProjects(g.ID, &gitlab.ListGroupProjectsOptions{
				gitlab.ListOptions{
					PerPage: 100,
				},
			})
			if err != nil {
				log.Fatal(err)
			}

			for _, p := range projects {
				mu.Lock()
				projectsList[p.ID] = p
				mu.Unlock()
			}
			wgp.Done()
		}(g, mu)
	}
	wgp.Wait()

	return projectsList
}

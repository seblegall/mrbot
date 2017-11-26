package gitlab

import (
	"log"
	"sync"
	"time"

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

//GetOpenMergeRequests returns a list a open merge request from groups passed as parameter.
//This function return only merge request created after a "from" parameters.
func (c *Client) GetOpenMergeRequests(groups []string, from time.Time) map[string][]map[string]string {

	results := make(map[string][]map[string]string, 500)
	groupsCh := c.GetGroups(groups)
	projectsList := c.GetProjects(groupsCh)

	pageOptions := gitlab.ListOptions{
		PerPage: 100,
	}

	mrOptions := &gitlab.ListMergeRequestsOptions{
		ListOptions: pageOptions,
		State:       gitlab.String("opened"),
		OrderBy:     gitlab.String("created_at"),
		Scope:       gitlab.String("all"),
	}

	mergeRequests, _, err := c.Client.MergeRequests.ListMergeRequests(mrOptions)

	if err != nil {
		log.Fatal(err)
	}

	for _, mr := range mergeRequests {
		if p, ok := projectsList[mr.ProjectID]; ok {
			if mr.CreatedAt.After(from) {
				mr := map[string]string{
					"projectName": p.Name,
					"title":       mr.Title,
					"createdAt":   mr.CreatedAt.Format(time.RFC3339),
					"createdBy":   mr.Author.Name,
				}

				results[mr["projectName"]] = append(results[mr["projectName"]], mr)
			}
		}
	}

	return results
}

//GetGroups search groups sepcified as parameter and return found groups as a slice.
func (c *Client) GetGroups(groupsToSearch []string) []*gitlab.Group {

	groups := make([]*gitlab.Group, len(groupsToSearch))
	var wgg sync.WaitGroup
	mu := &sync.Mutex{}

	for _, gtg := range groupsToSearch {
		wgg.Add(1)
		go func(gtg string) {
			gs, _, err := c.Client.Groups.ListGroups(&gitlab.ListGroupsOptions{
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
		}(gtg)
	}

	wgg.Wait()

	return groups
}

//GetProjects find all projects from a given slice of groups passed as parameter.
//This function returns a map of project where map keys are the corresponding project's ID
func (c *Client) GetProjects(groups []*gitlab.Group) map[int]*gitlab.Project {
	var wgp sync.WaitGroup
	projects := map[int]*gitlab.Project{}
	mu := &sync.Mutex{}

	for _, g := range groups {
		wgp.Add(1)
		go func(g *gitlab.Group) {
			//Find all projects
			projects, _, err := c.Client.Groups.ListGroupProjects(g.ID, &gitlab.ListGroupProjectsOptions{
				gitlab.ListOptions{
					PerPage: 100,
				},
			})
			if err != nil {
				log.Fatal(err)
			}

			for _, p := range projects {
				mu.Lock()
				projects[p.ID] = p
				mu.Unlock()
			}
			wgp.Done()
		}(g)
	}
	wgp.Wait()

	return projects
}

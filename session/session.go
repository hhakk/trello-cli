package session

import (
	"fmt"
	"github.com/adlio/trello"

	"gitlab.com/hhakk/trello-cli/config"
	"gitlab.com/hhakk/trello-cli/levenshtein"
)

type Session struct {
	Client  *trello.Client
	Members map[string]*trello.Member
	Lists   map[string]*trello.List
	User    *trello.Member
}

func NearestCard(q string, cards []*trello.Card) *trello.Card {
	var nearest *trello.Card
	min := -1
	for _, card := range cards {
		dist := levenshtein.Distance(q, (*card).Name)
		if min == -1 || dist < min {
			min = dist
			nearest = card
		}
	}
	return nearest
}

func (s *Session) NearestList(q string) *trello.List {
	var nearest *trello.List
	min := -1
	for _, list := range (*s).Lists {
		dist := levenshtein.Distance(q, (*list).Name)
		if min == -1 || dist < min {
			min = dist
			nearest = list
		}
	}
	return nearest
}

func (s *Session) NearestMember(q string) *trello.Member {
	var nearest *trello.Member
	min := -1
	for _, m := range (*s).Members {
		dist := levenshtein.Distance(q, (*m).Username)
		if min == -1 || dist < min {
			min = dist
			nearest = m
		}
	}
	return nearest
}

func BoardFromURL(url string, c *trello.Client) (*trello.Board, error) {
	brds, err := c.GetMyBoards(trello.Defaults())
	if err != nil {
		return nil, err
	}
	for _, b := range brds {
		if (*b).URL == url {
			return b, nil
		}
	}
	return nil, fmt.Errorf("No board found with specified URL.")
}

func Init(cfg *config.Config) (*Session, error) {
	s := new(Session)
	client := trello.NewClient((*cfg).Key, (*cfg).Token)
	s.Client = client
	b, err := BoardFromURL((*cfg).DefaultBoard, s.Client)
	if err != nil {
		return nil, err
	}
	lists, err := b.GetLists(trello.Defaults())
	if err != nil {
		return nil, err
	}
	listmap := make(map[string]*trello.List)
	for _, l := range lists {
		listmap[(*l).ID] = l
	}
	s.Lists = listmap
	members, err := b.GetMembers()
	if err != nil {
		return nil, err
	}
	mmap := make(map[string]*trello.Member)
	for _, m := range members {
		mmap[(*m).ID] = m
	}
	s.Members = mmap
	user, err := client.GetMyMember(trello.Defaults())
	if err != nil {
		return nil, err
	}
	s.User = user
	return s, nil
}

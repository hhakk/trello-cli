package cmd

import (
	"fmt"
	"time"

	"github.com/adlio/trello"

	"gitlab.com/hhakk/trello-cli/colors"
	"gitlab.com/hhakk/trello-cli/session"
)

func Add(s *session.Session, list, name, desc, due, mem string) error {
	lid := ""
	mid := ""
	lName := ""
	mName := ""
	dt, err := time.Parse("2006-01-02", due)
	if err != nil {
		return err
	}

	if name == "" {
		return fmt.Errorf("Card name needs to be non-empty.")
	}

	if list == "" {
		return fmt.Errorf("No list chosen for the card.")
	} else {
		l := s.NearestList(list)
		lid = (*l).ID
		lName = (*l).Name
	}

	if mem == "" {
		// default member id = yours
		mid = (*s.User).ID
		mName = (*s.User).Username
	} else {
		m := s.NearestMember(mem)
		mid = (*m).ID
		mName = (*m).Username
	}

	card := trello.Card{
		Name:      name,
		Desc:      desc,
		Due:       &dt,
		Pos:       65536.0, // first in list, next = 2*65536 + 1
		IDList:    lid,
		IDLabels:  []string{},
		IDMembers: []string{mid},
	}
	cardInfo := fmt.Sprintf(
		colors.Color("List: ", colors.TEAL) +
			lName +
			colors.Color("\nName: ", colors.TEAL) +
			name +
			colors.Color("\nDescription: ", colors.TEAL) +
			desc +
			colors.Color("\nDue: ", colors.TEAL) +
			due +
			colors.Color("\nAssigned to: ", colors.TEAL) +
			mName,
	)
	prompt := fmt.Sprintf(
		colors.Color("Add a NEW card:\n", colors.GREEN) +
			cardInfo +
			colors.Color("\n(y/n): ", colors.ORANGE),
	)
	if confirm(prompt) {
		err = s.Client.CreateCard(&card, trello.Defaults())
		if err != nil {
			return err
		}

		fmt.Println("Added card successfully.")
	}
	return nil
}

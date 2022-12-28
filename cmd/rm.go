package cmd

import (
	"fmt"

	"github.com/adlio/trello"

	"gitlab.com/hhakk/trello-cli/colors"
	"gitlab.com/hhakk/trello-cli/session"
)

func Rm(s *session.Session, list, name string) error {
	var lName string
	var l *trello.List
	if name == "" {
		return fmt.Errorf("Card name needs to be non-empty.")
	}

	if list == "" {
		return fmt.Errorf("No list chosen for the card.")
	} else {
		l = s.NearestList(list)
		lName = (*l).Name
	}
	cards, err := l.GetCards(trello.Defaults())
	if err != nil {
		return err
	}
	card := session.NearestCard(name, cards)
	cardInfo := fmt.Sprintf(
		colors.Color("List: ", colors.TEAL) +
			lName +
			colors.Color("\nName: ", colors.TEAL) +
			(*card).Name,
	)
	prompt := fmt.Sprintf(
		colors.Color("ARCHIVE a card:\n", colors.ORANGE) +
			cardInfo +
			colors.Color("\n(y/n): ", colors.ORANGE),
	)
	if confirm(prompt) {
		err := card.Archive()
		if err != nil {
			return err
		}
		fmt.Println("Archived card successfully.")
	}
	return nil
}

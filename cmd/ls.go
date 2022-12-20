package cmd

import (
	"fmt"

	"github.com/adlio/trello"
	"gitlab.com/hhakk/trello-cli/colors"
	"gitlab.com/hhakk/trello-cli/session"
)

func printMembers(s *session.Session, my bool) {
	if my {
		fmt.Printf("%s\t%s\t%s\n", (*s.User).ID, (*s.User).Username, (*s.User).FullName)
	} else {
		for k, v := range s.Members {
			fmt.Printf("%s\t%s\t%s\n", k, (*v).Username, (*v).FullName)
		}
	}
}

func formatCard(s *session.Session, card *trello.Card) string {
	due := ""
	ms := "" // members
	// format due date in parenthesis
	if (*card).Due != nil {
		due = "(" + (*card).Due.Format("2006-01-02") + ")"
		due = colors.Color(due, colors.ORANGE)
	}
	// format members in parenthesis
	if len((*card).IDMembers) > 0 {
		num := len((*card).IDMembers)
		ms = "("
		for j, m := range (*card).IDMembers {
			ms += (*s.Members[m]).Username
			if j != num-1 {
				ms += ", "
			}
		}
		ms += ")"
		ms = colors.Color(ms, colors.YELLOW)
	}
	res := fmt.Sprintf(
		"%s %s %s",
		(*card).Name,
		due,
		ms,
	)
	return res
}

func hasCardUser(card *trello.Card, id string) bool {
	for _, m := range (*card).IDMembers {
		if m == id {
			return true
		}
	}
	return false
}

func printCards(s *session.Session, my bool) error {
	for _, cur := range s.Lists {
		cards, err := cur.GetCards(trello.Defaults())
		if err != nil {
			return err
		}
		hasCards := len(cards) > 1
		hasMyCards := false
		if !hasCards {
			continue
		}
		for _, card := range cards {
			if hasCardUser(card, (*s.User).ID) {
				hasMyCards = true
				break
			}
		}
		if my && !hasMyCards {
			continue
		}
		// list title
		fmt.Printf(
			colors.Color((*cur).Name, colors.GREEN) +
				"\n",
		)
		// card
		for i, card := range cards {
			myCard := hasCardUser(card, (*s.User).ID)
			if my && !myCard {
				continue
			}
			fmt.Printf(
				"\t[%s] %s\n",
				colors.Color(fmt.Sprintf("%d", i), colors.TEAL),
				formatCard(s, card),
			)
		}
	}
	return nil
}

func Ls(s *session.Session, t string, my bool) {
	switch t {
	case "members":
		printMembers(s, my)
	case "cards":
		printCards(s, my)
	}
}

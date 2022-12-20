package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gitlab.com/hhakk/trello-cli/cmd"
	"gitlab.com/hhakk/trello-cli/colors"
	"gitlab.com/hhakk/trello-cli/config"
	"gitlab.com/hhakk/trello-cli/session"
)

func NextFriday() string {
	dt := int(time.Now().Weekday())
	toFriday := dt - 5
	if toFriday <= 0 {
		toFriday = 7 + dt
	}
	nf := time.Now().AddDate(0, 0, toFriday).Format("2006-01-02")
	return nf
}

func TrelloUsage() {
	w := flag.CommandLine.Output()
	prog := os.Args[0]
	fmt.Fprintf(
		w,
		"Usage: "+
			colors.Color(prog, colors.GREEN)+
			colors.Color(" <command> ", colors.YELLOW)+
			colors.Color("[OPTIONS]\n", colors.ORANGE),
	)
	flag.PrintDefaults()
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(
		w,
		colors.Color(prog, colors.GREEN)+
			colors.Color(" ls ", colors.YELLOW)+
			"{cards, members} "+
			colors.Color("[--user]\n", colors.ORANGE)+
			"  List cards or members of your board.\n"+
			"  --user string\n\tPrint items related to user.\n\n",
	)
	fmt.Fprintf(
		w,
		colors.Color(prog, colors.GREEN)+
			colors.Color(" add ", colors.YELLOW)+
			colors.Color("-n NAME -l LIST [-d DESCRIPTION] [-t DUE] [-m MEMBER]\n", colors.ORANGE)+
			"  Add a new card to specified list.\n"+
			"  Supplied arguments are fuzzy matched.\n"+
			"  To target user 'bob123' you may simply write '-m bob'\n"+
			"  -n string\n\tCard name\n"+
			"  -l string\n\tList name\n"+
			"  -t string\n\tDue date (YYYY-MM-DD), defaults to next Friday\n"+
			"  -m string\n\tMember username, defaults to yours\n\n",
	)
	fmt.Fprintf(
		w,
		colors.Color(prog, colors.GREEN)+
			colors.Color(" rm ", colors.YELLOW)+
			colors.Color("-n NAME -l LIST\n", colors.ORANGE)+
			"  Archive a card in the specified list.\n"+
			"  Supplied arguments are fuzzy matched.\n"+
			"  -n string\n\tCard name\n"+
			"  -l string\n\tList name\n",
	)
}

func main() {
	flag.Usage = TrelloUsage
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	// config path
	cfgdir := filepath.Join(home, ".config/trello-cli")
	err = os.MkdirAll(cfgdir, 0750)
	if err != nil {
		panic(err)
	}
	cfgp := filepath.Join(cfgdir, "config.json")

	// ls: list cards, lists, memebers
	lsType := "cards" // ls type, either cards or members
	userOnly := false // whether to retrieve user specific info
	lsCmd := flag.NewFlagSet("ls", flag.ExitOnError)
	lsCmd.StringVar(&lsType, "t", lsType, "type of info retrieved")
	lsCmd.BoolVar(&userOnly, "user", userOnly, "whether to retrieve only user info")

	// add: add card
	name := ""          // name of new card
	desc := ""          // description of card
	due := NextFriday() // due date
	list := ""          // list name
	mem := ""           // member mem
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCmd.StringVar(&list, "l", list, "list name for new card")
	addCmd.StringVar(&mem, "m", mem, "assign card to this member (defaults to you)")
	addCmd.StringVar(&name, "n", name, "name of new card")
	addCmd.StringVar(&desc, "d", desc, "description of new card")
	addCmd.StringVar(&due, "t", due, "due date for new card (defaults to next friday)")

	// rm: archive a card
	rmCmd := flag.NewFlagSet("rm", flag.ExitOnError)
	rmCmd.StringVar(&name, "n", name, "name of new card")
	rmCmd.StringVar(&list, "l", list, "list name for new card")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(2)
	}
	switch args[0] {
	case "ls":
		lsCmd.Parse(args[1:])
	case "add":
		addCmd.Parse(args[1:])
	case "rm":
		rmCmd.Parse(args[1:])
	default:
		flag.Usage()
		os.Exit(2)
	}

	var cfg config.Config

	err = cfg.Load(cfgp)
	if err != nil {
		cfg.Input()
		err = cfg.Save(cfgp)
		if err != nil {
			panic(err)
		}
	}

	s, err := session.Init(&cfg)
	if err != nil {
		panic(err)
	}
	if lsCmd.Parsed() {
		cmd.Ls(s, lsType, userOnly)
	} else if addCmd.Parsed() {
		err := cmd.Add(s, list, name, desc, due, mem)
		if err != nil {
			panic(err)
		}
	} else if rmCmd.Parsed() {
		err := cmd.Rm(s, list, name)
		if err != nil {
			panic(err)
		}
	}
}

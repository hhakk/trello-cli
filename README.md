# trello-cli

A CLI application using [github.com/adlio/trello](https://github.com/adlio/trello) to list, add and remove Trello cards.

You need a Trello API key and API Token for this CLI to work.

You can get them by following instructions from the [official documentation](https://developer.atlassian.com/cloud/trello/guides/rest-api/authorization/).

On the first time running this application, you will be prompted your API key, API Token and default board URL.

## installation

You may simply run
```
sudo make install
```

## usage

### 1. List cards (optionally  only ones assigned to you)

```
trello ls {cards, members} [--user]
```

### 2. Add a new card

```
trello add -n NAME -l LIST [-d DESCRIPTION] [-t DEADLINE] [-m MEMBER]
NAME = name for card
LIST = list name for card
DESCRIPTION = description for card
DEADLINE = due date (defaults to next Friday, format: YYYY-MM-DD)
MEMBER = member name (defaults to you)

MEMBER and LIST are fuzzy matched, so writing "sales" matches "Sales"
```

### 3. Archive a card

```
trello rm -n NAME -l LIST
NAME = name for card
LIST = list name for card
```

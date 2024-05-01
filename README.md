# Deck application

# How to run locally

```bash
$ make run
```

# How to run tests

```bash
$ make test
```

# Endpoints

## Create Deck

```
$ curl -X POST http://localhost:8080/deck
{"deck_id":"....","shuffled":false,"remaining":52}
```

Possible url params:
* `cards=AS,2S`: define cards in the deck
* `shuffle=true/false`: shuffles the decks

## Open Deck

```
$ curl http://localhost:8080/deck/:id
{
    "deck_id":"...",
    "shuffled":true,
    "remaining":1,
    "cards":[{"value":"9","suit":"DIAMONDS","code":"9D"}]
}
```
Return details of the deck


## Draw a card

```
$ curl http://localhost:8080/deck/:id/draw?count=1
{"cards":[{"value":"4","suit":"SPADES","code":"4S"}]}
```

Params:
* `count`: amount of cards to be drawn

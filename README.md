# jokenpo

This project is a simulation of the classic jokenpo game, where the player plays against the computer

## Running

```sh
make run
```

## How to play?

Below are examples of curl to play:

* Paper
```ssh
curl http://0.0.0.0:8000/move \
    --include \
    --header "Content-Type: application/json" \
    --request "GET" \
    --data '{"move": "paper"}'
```

* Rock
```ssh
curl http://0.0.0.0:8000/move \
    --include \
    --header "Content-Type: application/json" \
    --request "GET" \
    --data '{"move": "rock"}'
```

* Sissors
```ssh
curl http://0.0.0.0:8000/move \
    --include \
    --header "Content-Type: application/json" \
    --request "GET" \
    --data '{"move": "sissors"}'
```


# shortURL
To start the server, navigate to the project's root directory, and type
```bash
go run cmd/main.go
```

You can verify TWO parts of a simple URL-shorten system, 
by using following scripts:
```bash
# remember to replace <RawURL> with your real URL
shortURL=$(
    curl -X POST \
    http://localhost:8080/ \
    -H 'content-type: application/json' \
    -d '{"long": "<RawURL>"}'
)

curl -L -s ${shortURL}
```
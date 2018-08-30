# fcm-go

Firebase Cloud Messaging send tool

## Installation

```
go get -u github.com/satooon/fcm-go
```

## Use

```
fcm-go help
fcm-go -c ./serviceAccountKey.json --dry-run topic --title test --body test --name topic
fcm-go -c ./serviceAccountKey.json tokens --title test --body test --tokens XXXX --tokens YYYY --tokens ZZZZ
```

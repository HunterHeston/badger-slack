# Honeybadger Assignment

## Dependencies

- Just the [echo](https://echo.labstack.com/) framework for Go.

## Testing

Quick and dirty testing to start:

Payload that sends the notification:

```curl
http://localhost:3000/record \
  -H 'Content-Type: application/json' \
  -d '{
  "RecordType": "Bounce",
  "MessageStream": "outbound",
  "Type": "SpamNotification",
  "TypeCode": 512,
  "Name": "Spam notification",
  "Tag": "",
  "Description": "The message was delivered, but was either blocked by the user, or classified as spam, bulk mail, or had rejected content.",
  "Email": "zaphod@example.com",
  "From": "notifications@honeybadger.io",
  "BouncedAt": "2023-02-27T21:41:30Z"
}'
```

Payload that **does not** sends the notification:

```curl
http://localhost:3000/record \
  -H 'Content-Type: application/json' \
  -d '{
  "RecordType": "Bounce",
  "MessageStream": "outbound",
  "Type": "HardBounce",
  "TypeCode": 1,
  "Name": "Hard bounce",
  "Tag": "Test",
  "Description": "The server was unable to deliver your message (ex: unknown user, mailbox not found).",
  "Email": "arthur@example.com",
  "From": "notifications@honeybadger.io",
  "BouncedAt": "2019-11-05T16:33:54.9070259Z",
}'
```

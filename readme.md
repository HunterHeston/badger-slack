# Honeybadger Assignment

## Code of Interest

1. [main.go](main.go) - Server setup and route handling.
2. [slack/sendMessage.go](slack/sendMessage.go) - Interactions with the slack api
3. [experiment/main.go](experiment/main.go) - Just a test of the slack api

## How to setup and run

1. Run `git clone git@github.com:HunterHeston/badger-slack.git`.
2. Create an `.env` file based on `.env.example` at the project root.
3. You may need to run `go get` to install dependencies.
4. Run `go run .` from the root directory.
5. In another terminal run either of the commands from the "Testing" section.

### Slack api key setup

I found these steps and was able to get it working. Your milage may very.

1. Go to the [Slack API website](https://api.slack.com/).
2. Click on the "Create New App" button to create a new Slack app.
3. Choose a name and a workspace for your app.
4. In the "Add features and functionality" section, click on "Bots" and add a new bot user to your app.
5. In the "Install App" section, click on the "Install App to Workspace" button to install your app in your workspace.
6. After installation, you'll be redirected to a page with your Slack API token. Copy this token and store it securely.

## Dependencies

- [echo](https://echo.labstack.com/) framework for Go.
- [slack-go](https://github.com/slack-go/slack) a community maintained slack API as a library.
- [dotenv](github.com/joho/godotenv) parse and process .env files.

## Testing

Quick and dirty testing to start:

### Payload that sends the notification:

```curl
curl -X POST \
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

### Payload that **does not** sends the notification:

```curl
curl -X POST \
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
    "BouncedAt": "2019-11-05T16:33:54.9070259Z"
}'
```

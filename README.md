# MrBot

A really smart Jabber chat bot underlying on the Dialogflow AI engine and written in Go.

MrBot is a chat bot made to answer "@mrbot" notifications using AI.
It basically connect to your Jabber server (Hipchat here) using XMPP protocol,
listen to "@mrbot" notification and answer using the Dialogflow AI.

## Context

Dialogflow comes with a lot of pre-made integration such as Twitter, Slack of Facebook Messenger. But there is no integration with a classical Jabber server.

MrBot is a boilerplate example of connecting your Dialogfow agent with a Jabber server.

## How it works

Once running, MrBot will connect to your Jabber (Hipchat) server and then join one or more room.
Then, MrBot will listen all the messages sent in the room and only catch the one that begin with "@mrbot".

MrBot will then catch those *notifications* messages and send them to Dialogflow.
Dialogflow will take this message and match it to a corresponding [intent](https://dialogflow.com/docs/intents) you've defined. Then, Dialogflow will return the corresponding response depending on the *context* and your intent configuration.

To be able to do that, MrBot use the [DialogFlow REST API](https://dialogflow.com/docs/reference/agent/).

## Configuration

Configuration is handle by env variable. For development purpose, you can create
a `.env` file containing the following configuration. MrBot will automatically read
the file at start up.

```bash
MRBOT_HIPCHAT_USERNAME="123456"
MRBOT_HIPCHAT_PASSWORD="XxxXxxxX"
MRBOT_HIPCHAT_ROOMJID="test@conf.hipchat.com"
MRBOT_DIALOGFLOW_TOKEN="your_dialogflow_token"
```

## Usage

### Getting dependencies (only once)

```bash
make dep
```

#### Build

```bash
make build
```

### Run the chat bot

````bash
./mrbot
````

### Deployment

MrBot is Heroku ready. Everything is set up for you to deploy the bot on your
own Heroku instance.
# test sending message using webhook
- might be too easy to create a repository

## setups

### slack
1. create slack app
2. invite to slack
3. activate and create webhook
4. copy and set environmental variable `SLACK_WEBHOOK_URL` 
   
### discord
1. `server settings>integrations` and create webhooks
2. copy and set environmental variable `DISCORD_WEBHOOK_URL` 


### others
1. you can handle slack webhook in discord, if you add `/slack` to `DISCORD_WEBHOOK_URL` 
    - [ref](https://discord.com/developers/docs/resources/webhook#execute-slackcompatible-webhook)


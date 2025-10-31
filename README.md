# Squad Rotation Bot

This is a bot that you can integrate into your gchat spaces to send a message
to indicate who's turn is it to facilitate the daily.

## Configuring

Make sure to configure the following according to your needs:

- **The time at which to run the bot** is configured in `scripts/crontab`.
  Follow the indications in that file.
- **Set the webhook URL** by setting the `WEB_HOOK_URL` environment variable to
  the value of the gchat webhook URL. You can generate this in the integrations
  settings tab of your gchat space.
- **Optionally configure the database connection details** by setting the
  environment variables used in the compose file

## Usage

### With docker-compose

1. Clone this repository:

```shell
git clone https://github.com/aloussase/squad-rotation-bot
cd squad-rotation-bot
```

2. Run the services:

```shell
export WEB_HOOK_URL='...'
docker compose up -d
```

## Contributing

Contributing ideas:

- More messaging backends (e.g.: discord, slack, carrier pigeon; etc)

## License
MIT

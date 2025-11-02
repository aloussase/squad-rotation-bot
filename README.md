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

## Managing squad members

You basically have 3 options for managing squad members.

### 1. Edit the `databse/init.sql` file

You can edit the database initialization script and add `INSERT` statements to
insert your squad members.

### 2. Exec into the Postgres container

You can also exec into the running Postgres container to use `psql` to `INSERT` the
squad members:

```shell
docker compose exec db psql -U postgres
```

That command will start a `psql` console where you can issue your database commands.

Make sure to replace `postgres` by whatever value you've configured as the
database user, if any. `postgres` is the default one.

### 3. Use the provided CLI

There is a Docker image that packages a CLI tool to manage your squad members.
You can run it with the following command:

```shell
docker run -it --network host aloussase69/squad-rotation-bot-cli
```

![Squad Rotation Bot CLI Demo](./assets/cli_demo.mp4)

## Contributing

Contributing ideas:

- More messaging backends (e.g.: discord, slack, carrier pigeon; etc)

## License
MIT

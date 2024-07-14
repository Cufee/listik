# Listik

## Dependencies

- [Gomponents](https://www.gomponents.com/)
  - [Gomponents+](https://www.gomponents.com/) should be used where possible
- [Bulma](https://bulma.io/)

## Development
In order to use auth, you need to do the following
  - Create a [Google Cloud](https://developers.google.com/identity/gsi/web/guides/get-google-api-clientid) account and generate a new Client ID, auth scopes can stay default - no need for any additional scopes
  - Sign up for [ngrok](https://ngrok.com/docs/getting-started/) and install the cli

### Requirements
- [go](https://go.dev/)
  - check the version in `go.mod`
- [cosmtrek/air](https://github.com/cosmtrek/air)
- [Task](https://taskfile.dev/installation/)
- Setup `.env` by running `cp .env.example .env` and adjust values where needed

### Start the app locally

```
task dev
```


# slack-status-cli

## Usage

```bash
-> % go run main.go --help
NAME:
   slack-status - A new cli application

USAGE:
   slack-status [global options] command [command options] 

COMMANDS:
   get      
   set      
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Configure systemd timer

Create `.env` under `/etc/slack-status-cli` like this:

```ini
SLACK_TOKEN=<your-api-token>
SLACK_USER_ID=<user-id>
```

Put service and timer file under `/etc/systemd/system`

```
systemctl enable slack-status-cli.timer
```

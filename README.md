# slack-emoji-manager

Slack Emoji Manager is a simple command line tool to manage emojis in Slack.

# Usage

## User token

A user token is required for all requests, ie an `xoxs` token

## Get emoji list

Return a list of emojis for a Slack insance
```bash
./slack-emoji-manager get {userToken}
```

## Download all emojis

Download all emojis for a Slack instance into `/emojis`
```bash
./slack-emoji-manager download {userToken}
```

## Upload emojis

The `emoji.add` request is currently not documented so may stop working at any point.

### Single emoji upload

Upload a single emoji to a Slack instance. 

_png, gif, jpg, jpeg_

```bash
./slack-emoji-manager upload {fileName} {userToken}
```

### Multi emoji upload

Upload all emojis in a folder

_png, gif, jpg, jpeg_

```bash
./slack-emoji-manager upload-all {folder} {userToken}
```


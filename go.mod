module slack-emoji-manager

go 1.12

require (
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/nlopes/slack v0.5.0
	github.com/pkg/errors v0.8.1 // indirect
	github.com/trevorsmiley/fileutils v0.0.0-20190713005028-ca4041d832de
)

replace github.com/nlopes/slack => github.com/trevorsmiley/slack v0.5.1-0.20190705141651-c8017f80f074

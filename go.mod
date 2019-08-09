module github.com/lampjaw/mutterblack

go 1.12

require (
	github.com/bwmarrin/discordgo v0.19.0
	github.com/dustin/go-humanize v1.0.0
	github.com/lampjaw/discordgobot v0.0.0-20190809181220-b17aec24090a
	github.com/lampjaw/mutterblack/pkg/plugins/invite v0.0.0
	github.com/lampjaw/mutterblack/pkg/plugins/planetsidetwo v0.0.0
	github.com/lampjaw/mutterblack/pkg/plugins/stats v0.0.0
	github.com/lampjaw/mutterblack/pkg/plugins/translator v0.0.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)

replace (
	github.com/lampjaw/mutterblack/pkg/plugins/invite => ./pkg/plugins/invite
	github.com/lampjaw/mutterblack/pkg/plugins/planetsidetwo => ./pkg/plugins/planetsidetwo
	github.com/lampjaw/mutterblack/pkg/plugins/stats => ./pkg/plugins/stats
	github.com/lampjaw/mutterblack/pkg/plugins/translator => ./pkg/plugins/translator
)

module github.com/lampjaw/mutterblack

go 1.12

replace (
	github.com/bwmarrin/discordgo => github.com/bwmarrin/discordgo v0.19.0
	github.com/dustin/go-humanize => github.com/dustin/go-humanize v1.0.0
	github.com/lampjaw/discordgobot => github.com/lampjaw/discordgobot v0.2.3
	github.com/lampjaw/mutterblack/cmd/mutterblack => ./cmd/mutterblack
	github.com/lampjaw/mutterblack/pkg/plugins/command => ./pkg/plugins/command
	github.com/lampjaw/mutterblack/pkg/plugins/invite => ./pkg/plugins/invite
	github.com/lampjaw/mutterblack/pkg/plugins/planetsidetwo => ./pkg/plugins/planetsidetwo
	github.com/lampjaw/mutterblack/pkg/plugins/stats => ./pkg/plugins/stats
	github.com/lampjaw/mutterblack/pkg/plugins/translator => ./pkg/plugins/translator
	golang.org/x/oauth2 => golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)

require (
	github.com/lampjaw/mutterblack/pkg/plugins/command v0.0.0-00010101000000-000000000000 // indirect
	github.com/lampjaw/mutterblack/pkg/plugins/invite v0.0.0-00010101000000-000000000000 // indirect
	github.com/lampjaw/mutterblack/pkg/plugins/planetsidetwo v0.0.0-00010101000000-000000000000 // indirect
	github.com/lampjaw/mutterblack/pkg/plugins/stats v0.0.0-00010101000000-000000000000 // indirect
	github.com/lampjaw/mutterblack/pkg/plugins/translator v0.0.0-00010101000000-000000000000 // indirect
	github.com/lampjaw/mutterblack/pkg/plugins/weather v0.0.0-20190814203813-f0a186eec301 // indirect
)

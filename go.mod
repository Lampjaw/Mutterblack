module github.com/lampjaw/mutterblack

go 1.12

require (
	github.com/bwmarrin/discordgo v0.19.0
	github.com/dustin/go-humanize v1.0.0
	github.com/lampjaw/mutterblack/internal/pkg/plugin v0.0.0
	github.com/lampjaw/mutterblack/internal/pkg/plugin/invite v0.0.0
	github.com/lampjaw/mutterblack/internal/pkg/plugin/planetsidetwo v0.0.0
	github.com/lampjaw/mutterblack/internal/pkg/plugin/stats v0.0.0
	github.com/lampjaw/mutterblack/internal/pkg/plugin/translator v0.0.0
	github.com/lampjaw/mutterblack/pkg/bot v0.0.0
	github.com/lampjaw/mutterblack/pkg/command v0.0.0
	github.com/lampjaw/mutterblack/pkg/discord v0.0.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)

replace (
	github.com/lampjaw/mutterblack/internal/pkg/plugin => ./internal/pkg/plugin
	github.com/lampjaw/mutterblack/internal/pkg/plugin/invite => ./internal/pkg/plugin/invite
	github.com/lampjaw/mutterblack/internal/pkg/plugin/planetsidetwo => ./internal/pkg/plugin/planetsidetwo
	github.com/lampjaw/mutterblack/internal/pkg/plugin/stats => ./internal/pkg/plugin/stats
	github.com/lampjaw/mutterblack/internal/pkg/plugin/translator => ./internal/pkg/plugin/translator
	github.com/lampjaw/mutterblack/pkg/bot => ./pkg/bot
	github.com/lampjaw/mutterblack/pkg/command => ./pkg/command
	github.com/lampjaw/mutterblack/pkg/discord => ./pkg/discord
)

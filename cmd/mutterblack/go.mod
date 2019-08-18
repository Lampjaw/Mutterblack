module github.com/lampjaw/mutterblack/cmd/mutterblack

go 1.12

require (
	github.com/lampjaw/discordgobot v0.2.3
	github.com/lampjaw/mutterblack/pkg/plugins/command v0.0.0
	github.com/lampjaw/mutterblack/pkg/plugins/invite v0.0.0
	github.com/lampjaw/mutterblack/pkg/plugins/planetsidetwo v0.0.0
	github.com/lampjaw/mutterblack/pkg/plugins/stats v0.0.0
	github.com/lampjaw/mutterblack/pkg/plugins/translator v0.0.0
	github.com/lampjaw/mutterblack/pkg/plugins/weather v0.0.0
)

replace (
	github.com/lampjaw/mutterblack/pkg/plugins/command => ../../pkg/plugins/command
	github.com/lampjaw/mutterblack/pkg/plugins/invite => ../../pkg/plugins/invite
	github.com/lampjaw/mutterblack/pkg/plugins/planetsidetwo => ../../pkg/plugins/planetsidetwo
	github.com/lampjaw/mutterblack/pkg/plugins/stats => ../../pkg/plugins/stats
	github.com/lampjaw/mutterblack/pkg/plugins/translator => ../../pkg/plugins/translator
	github.com/lampjaw/mutterblack/pkg/plugins/weather => ../../pkg/plugins/weather
)

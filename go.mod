module github.com/lampjaw/mutterblack

go 1.12

replace (
	github.com/lampjaw/mutterblack/pkg/plugins/invite => ./pkg/plugins/invite
	github.com/lampjaw/mutterblack/pkg/plugins/planetsidetwo => ./pkg/plugins/planetsidetwo
	github.com/lampjaw/mutterblack/pkg/plugins/stats => ./pkg/plugins/stats
	github.com/lampjaw/mutterblack/pkg/plugins/translator => ./pkg/plugins/translator
	github.com/lampjaw/mutterblack/pkg/plugins/weather => ./pkg/plugins/weather
)

require (
	github.com/lampjaw/discordgobot v0.1.1
	github.com/lampjaw/mutterblack/cmd/mutterblack v0.0.0-20190809183334-6f42ef28c2e8
	github.com/lampjaw/mutterblack/pkg/plugins/weather v0.0.0-00010101000000-000000000000 // indirect
)

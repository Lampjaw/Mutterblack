module github.com/lampjaw/mutterblack

go 1.12

replace (
	github.com/lampjaw/mutterblack/pkg/plugins/invite => ./pkg/plugins/invite
	github.com/lampjaw/mutterblack/pkg/plugins/planetsidetwo => ./pkg/plugins/planetsidetwo
	github.com/lampjaw/mutterblack/pkg/plugins/stats => ./pkg/plugins/stats
	github.com/lampjaw/mutterblack/pkg/plugins/translator => ./pkg/plugins/translator
)

require github.com/lampjaw/mutterblack/cmd/mutterblack v0.0.0-20190809183334-6f42ef28c2e8 // indirect

module github.com/lampjaw/mutterblack/pkg/bot

go 1.12

require (
	github.com/lampjaw/mutterblack/internal/pkg/plugin v0.0.0
	github.com/lampjaw/mutterblack/pkg/command v0.0.0
	github.com/lampjaw/mutterblack/pkg/discord v0.0.0
)

replace (
	github.com/lampjaw/mutterblack/internal/pkg/plugin => ../../internal/pkg/plugin
	github.com/lampjaw/mutterblack/pkg/command => ../command
	github.com/lampjaw/mutterblack/pkg/discord => ../discord
)

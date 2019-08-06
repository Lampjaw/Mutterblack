module github.com/lampjaw/mutterblack/internal/pkg/plugin/planetsidetwo

go 1.12

require (
	github.com/bwmarrin/discordgo v0.19.0
	github.com/lampjaw/mutterblack/internal/pkg/plugin v0.0.0
	github.com/lampjaw/mutterblack/pkg/command v0.0.0
	github.com/lampjaw/mutterblack/pkg/discord v0.0.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
)

replace (
	github.com/lampjaw/mutterblack/internal/pkg/plugin => ../
	github.com/lampjaw/mutterblack/pkg/command => ../../../../pkg/command
	github.com/lampjaw/mutterblack/pkg/discord => ../../../../pkg/discord
)

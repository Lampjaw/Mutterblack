using Discord;
using Discord.Interactions;

namespace Mutterblack.Bot.Modules
{
    [EnabledInDm(true)]
    public class FunModule : InteractionModuleBase<SocketInteractionContext>
    {
        [SlashCommand("twanswate", "Twanswate the previous comment")]
        public async Task TwanswateAsync()
        {
            var messages = await Context.Channel.GetMessagesAsync(1, CacheMode.AllowDownload).FlattenAsync();
            var lastMessage = messages.FirstOrDefault();

            if (lastMessage == null || string.IsNullOrWhiteSpace(lastMessage.Content))
            {
                await RespondAsync("Unable to find a message to translate.", ephemeral: true);
                return;
            }
            else if (lastMessage.Author.IsBot)
            {
                await RespondAsync("Sorry, I can't translate bots.", ephemeral: true);
                return;
            }

            var content = lastMessage.Content;

            var selfPrefix = $"<@{Context.Client.CurrentUser.Id}>";
            if (content.StartsWith(selfPrefix))
            {
                content = content.Substring(selfPrefix.Length + 1);
            }

            content = content
                .Replace('r', 'w')
                .Replace('R', 'W')
                .Replace('l', 'w')
                .Replace('L', 'W');

            var authorBuilder = new EmbedAuthorBuilder()
                .WithName(lastMessage.Author.Username)
                .WithIconUrl(lastMessage.Author.GetAvatarUrl());

            var embed = new EmbedBuilder()
                .WithAuthor(authorBuilder)
                .WithColor(Constants.DefaultEmbedColor)
                .WithDescription(content)
                .WithTimestamp(lastMessage.Timestamp)
                .WithFooter(string.Format("in #{0} at {1}", Context.Channel.Name, Context.Guild.Name))
                .Build();

            await RespondAsync(embed: embed);
        }
    }
}

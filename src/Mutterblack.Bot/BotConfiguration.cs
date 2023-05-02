namespace Mutterblack.Bot
{
    public class BotConfiguration
    {
        public string DiscordToken { get; set; }
        public string DiscordClientId { get; set; }
        public ulong? GuildId { get; set; }
        public ulong? OwnerId { get; set; }
        public string VoidwellClientId { get; set; }
        public string VoidwellClientSecret { get; set; }
    }
}

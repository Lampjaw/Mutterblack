using Discord;
using Discord.Addons.Hosting;
using Discord.WebSocket;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Mutterblack.Bot;
using Mutterblack.Bot.Services;
using Serilog;
using Serilog.Events;
using Voidwell.Microservice.Http.AuthenticatedHttpClient;

const LogSeverity DiscordLogLevel = LogSeverity.Info;

var logger = new LoggerConfiguration()
    .MinimumLevel.Information()
    .MinimumLevel.Override("Microsoft", LogEventLevel.Error)
    .MinimumLevel.Override("System.Net.Http.HttpClient", LogEventLevel.Error);

if (string.Equals(Environment.GetEnvironmentVariable("LoggingOutput"), "flat", StringComparison.OrdinalIgnoreCase))
{
    logger.WriteTo.Console(outputTemplate: "[{Level:u4} {Timestamp:HH:mm:ss.fff}] {SourceContext}{NewLine}{Message} {Exception}{NewLine}");
}
else
{
    logger.WriteTo.Sink<GraylogConsoleSink>();
}

Log.Logger = logger.CreateLogger();

IHost host = Host.CreateDefaultBuilder(args)
    .UseSerilog()
    .ConfigureAppConfiguration((hostingContext, builder) =>
    {
        builder.Sources.Clear();

        builder.AddEnvironmentVariables();
    })
    .ConfigureDiscordHost((context, config) =>
    {
        config.SocketConfig = new DiscordSocketConfig
        {
            LogLevel = DiscordLogLevel,
            MessageCacheSize = 10
        };

        config.Token = context.Configuration.Get<BotConfiguration>().DiscordToken;

        config.SocketConfig.GatewayIntents = GatewayIntents.Guilds | GatewayIntents.GuildMessages | GatewayIntents.DirectMessages;

        config.LogFormat = (message, exception) => $"{message.Source}: {message.Message}";
    })
    .UseInteractionService((context, config) =>
    {
        config.LogLevel = DiscordLogLevel;
        config.UseCompiledLambda = true;
    })
    .UseCommandService((context, config) =>
    {
        config.LogLevel = DiscordLogLevel;
    })
    .ConfigureServices((hostContext, services) =>
    {
        services.Configure<BotConfiguration>(hostContext.Configuration);

        var botConfiguration = hostContext.Configuration.Get<BotConfiguration>();

        services.AddAuthenticatedHttpClient<VoidwellClient>(options =>
        {
            options.TokenServiceAddress = "https://auth.voidwell.com/connect/token";
            options.ClientId = botConfiguration.VoidwellClientId;
            options.ClientSecret = botConfiguration.VoidwellClientSecret;
            options.Scopes = new List<string>
            {
                "voidwell-daybreakgames",
                "voidwell-api"
            };
        });

        services.AddHostedService<InteractionHandler>();
        services.AddHostedService<CommandHandler>();
    })
    .Build();

await host.RunAsync();
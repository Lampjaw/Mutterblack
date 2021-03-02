# Mutterblack

Mutterblack is a discord bot primarily repsonsible for providing planetside 2 stats. It relies on the [Voidwell](https://voidwell.com) API and requires credentials to access its endpoints. As it stands, it is not possible to run a version of this bot locally as these endpoints are restricted.

### Adding Mutterblack to your Discord server
If you would like to use Mutterblack, an admin of your server needs to go to this link:    
https://discordapp.com/oauth2/authorize?client_id=439194558270537728&scope=bot
 
### Commands
The following commands are available:
* `?invite` - Get an invite link to add this bot to your server!
* `?ps2c <character name>` - Get stats for a player.
* `?ps2c <character name> <weapon name>` - Get weapon stats for a player.
* `?ps2c-ps4eu <character name>` - Get stats for a player.
* `?ps2c-ps4eu <character name> <weapon name>` - Get weapon stats for a player.
* `?ps2c-ps4us <character name>` - Get stats for a player.
* `?ps2c-ps4us <character name> <weapon name>` - Get weapon stats for a player.
* `?ps2o <outfit tag>` - Get outfit stats
* `?ps2o-ps4eu <outfit tag>` - Get outfit stats
* `?ps2o-ps4us <outfit tag>` - Get outfit stats
* `?ps2w <weapon name>` - Get weapon stats
* `?setprefix <prefix>` - Set the command prefix for this server
* `?stats` - Get bot stats
* `?twanswate` - Twanswate the previous comment

Notes:
* In some cases if you're getting a bad match and you know the ID of the character or weapon you're trying to look up you may use that instead of a name.    
* Queries are not case sensitive.
* Character names and outfit tags must be the full name.
* Partial weapon names are allowed and it will try to find the best match (i.e "msw" will return results for the "MSW-R").

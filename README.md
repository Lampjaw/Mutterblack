# Mutterblack

Mutterblack is a discord bot primarily repsonsible for providing planetside 2 stats. It relies on the [Voidwell](https://voidwell.com) API and requires credentials to access its endpoints. As it stands, it is not possible to run a version of this bot locally as these endpoints are restricted.

### Adding Mutterblack to your Discord server
If you would like to use Mutterblack, an admin of your server needs to go to this link:    
https://discordapp.com/oauth2/authorize?client_id=439194558270537728&scope=bot
 
### Slash Commands
The following commands are available:
* `/invite` - Get an invite link to add this bot to your server!
* `/ps2 player <character name>` - Get stats for a player.
* `/ps2 player <character name> <weapon name>` - Get weapon stats for a player.
* `/ps2 outfit <outfit tag>` - Get outfit stats
* `/ps2 weapon <weapon name>` - Get weapon stats

The Player and Outfit commands also support an optional Platform Type argument with possible values PC, PS4-US, and PS4-EU

Notes:
* In some cases if you're getting a bad match and you know the ID of the character or weapon you're trying to look up you may use that instead of a name.    
* Queries are not case sensitive.
* Character names and outfit tags must be the full name.
* Partial weapon names are allowed and it will try to find the best match (i.e "msw" will return results for the "MSW-R").

package planetsidetwoplugin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lampjaw/discordgobot"
	"golang.org/x/oauth2/clientcredentials"
)

const CENSUS_IMAGEBASE_URI = "http://census.daybreakgames.com/files/ps2/images/static/"
const VOIDWELL_URI = "https://voidwell.com/"

var voidwellClientConfig clientcredentials.Config
var voidwellClient *http.Client

type planetsidetwoPlugin struct {
	discordgobot.Plugin
}

func New() discordgobot.IPlugin {
	return &planetsidetwoPlugin{}
}

func (p *planetsidetwoPlugin) Commands() []*discordgobot.CommandDefinition {
	return []*discordgobot.CommandDefinition{
		&discordgobot.CommandDefinition{
			CommandID: "ps2-character",
			Triggers: []string{
				"ps2c",
				"ps2c-ps4us",
				"ps2c-ps4eu",
			},
			Arguments: []discordgobot.CommandDefinitionArgument{
				discordgobot.CommandDefinitionArgument{
					Pattern: "[a-zA-Z0-9]*",
					Alias:   "characterName",
				},
			},
			Description: "Get stats for a player.",
			Callback:    p.runCharacterStatsCommand,
		},
		&discordgobot.CommandDefinition{
			CommandID: "ps2-character-weapons",
			Triggers: []string{
				"ps2c",
				"ps2c-ps4us",
				"ps2c-ps4eu",
			},
			Arguments: []discordgobot.CommandDefinitionArgument{
				discordgobot.CommandDefinitionArgument{
					Pattern: "[a-zA-Z0-9]*",
					Alias:   "characterName",
				},
				discordgobot.CommandDefinitionArgument{
					Pattern: ".*",
					Alias:   "weaponName",
				},
			},
			Description: "Get weapon stats for a player.",
			Callback:    p.runCharacterWeaponStatsCommand,
		},
		&discordgobot.CommandDefinition{
			CommandID: "ps2-outfit",
			Triggers: []string{
				"ps2o",
				"ps2o-ps4us",
				"ps2o-ps4eu",
			},
			Arguments: []discordgobot.CommandDefinitionArgument{
				discordgobot.CommandDefinitionArgument{
					Pattern: "[a-zA-Z0-9]{1,4}",
					Alias:   "outfitAlias",
				},
			},
			Description: "Get outfit stats by outfit tag.",
			Callback:    p.runOutfitStatsCommand,
		},
		&discordgobot.CommandDefinition{
			CommandID: "ps2-weapon",
			Triggers: []string{
				"ps2w",
			},
			Arguments: []discordgobot.CommandDefinitionArgument{
				discordgobot.CommandDefinitionArgument{
					Pattern: ".*",
					Alias:   "weaponName",
				},
			},
			Description: "Get weapon stats by weapon name.",
			Callback:    p.runWeaponStatsCommand,
		},
	}
}

func (p *planetsidetwoPlugin) Help(bot *discordgobot.Gobot, client *discordgobot.DiscordClient, message discordgobot.Message, detailed bool) []string {
	commandPrefix := bot.GetCommandPrefix(message)

	return []string{
		discordgobot.CommandHelp(client, "ps2c", []string{"character name"}, "Get stats for a player.", commandPrefix),
		discordgobot.CommandHelp(client, "ps2c-ps4us", []string{"character name"}, "Get stats for a player.", commandPrefix),
		discordgobot.CommandHelp(client, "ps2c-ps4eu", []string{"character name"}, "Get stats for a player.", commandPrefix),
		discordgobot.CommandHelp(client, "ps2c", []string{"character name", "weapon name"}, "Get weapon stats for a player.", commandPrefix),
		discordgobot.CommandHelp(client, "ps2c-ps4us", []string{"character name", "weapon name"}, "Get weapon stats for a player.", commandPrefix),
		discordgobot.CommandHelp(client, "ps2c-ps4eu", []string{"character name", "weapon name"}, "Get weapon stats for a player.", commandPrefix),
		discordgobot.CommandHelp(client, "ps2o", []string{"outfit name"}, "Get outfit stats", commandPrefix),
		discordgobot.CommandHelp(client, "ps2o-ps4us", []string{"outfit name"}, "Get outfit stats", commandPrefix),
		discordgobot.CommandHelp(client, "ps2o-ps4eu", []string{"outfit name"}, "Get outfit stats", commandPrefix),
		discordgobot.CommandHelp(client, "ps2w", []string{"weapon name"}, "Get weapon stats", commandPrefix),
	}
}

func (p *planetsidetwoPlugin) Name() string {
	return "PS2Stats"
}

func (p *planetsidetwoPlugin) runCharacterStatsCommand(bot *discordgobot.Gobot, client *discordgobot.DiscordClient, payload discordgobot.CommandPayload) {
	trigger, args, message := payload.Trigger, payload.Arguments, payload.Message

	if trigger == "ps2c-ps4us" {
		args["platform"] = "ps4us"
	} else if trigger == "ps2c-ps4eu" {
		args["platform"] = "ps4eu"
	} else {
		args["platform"] = "pc"
	}

	resp, err := voidwellAPIGet(fmt.Sprintf("https://api.voidwell.com/ps2/character/byname/%s?platform=%s", args["characterName"], args["platform"]))

	if err != nil {
		p.RLock()
		client.SendMessage(message.Channel(), fmt.Sprintf("%s", err))
		p.RUnlock()
		return
	}

	var character PlanetsideCharacter
	json.Unmarshal(resp, &character)

	lastSaved, _ := time.Parse(time.RFC3339, character.LastSaved)

	fields := []*discordgo.MessageEmbedField{
		&discordgo.MessageEmbedField{
			Name:   "Last Seen",
			Value:  fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d UTC", lastSaved.Year(), lastSaved.Month(), lastSaved.Day(), lastSaved.Hour(), lastSaved.Minute(), lastSaved.Second()),
			Inline: false,
		},
		&discordgo.MessageEmbedField{
			Name:   "Server",
			Value:  character.World,
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Battle Rank",
			Value:  fmt.Sprintf("%d", character.BattleRank),
			Inline: false,
		},
		&discordgo.MessageEmbedField{
			Name:   "Kills",
			Value:  fmt.Sprintf("%d", character.Kills),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Play Time",
			Value:  fmt.Sprintf("%0.1f (%0.1f) Hours", float32(character.PlayTime)/3600.0, float32(character.TotalPlayTimeMinutes)/60.0),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "KDR",
			Value:  fmt.Sprintf("%0.2f", character.KillDeathRatio),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "HSR",
			Value:  fmt.Sprintf("%0.2f%%", character.HeadshotRatio*100),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "KpH",
			Value:  fmt.Sprintf("%0.2f (%0.2f)", character.KillsPerHour, character.TotalKillsPerHour),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Siege Level",
			Value:  fmt.Sprintf("%0.1f", character.SiegeLevel),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "IVI Score",
			Value:  fmt.Sprintf("%d", character.IVIScore),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "IVI KDR",
			Value:  fmt.Sprintf("%0.2f", character.IVIKillDeathRatio),
			Inline: true,
		},
	}

	if len(character.OutfitName) > 0 {
		outfitValue := character.OutfitName
		if len(character.OutfitAlias) > 0 {
			outfitValue = "[" + character.OutfitAlias + "] " + character.OutfitName
		}

		outfitField := &discordgo.MessageEmbedField{
			Name:   "Outfit",
			Value:  outfitValue,
			Inline: true,
		}

		fields = insertSlice(fields, outfitField, 2)
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: character.Name,
		},
		Title: "Click here for full stats",
		URL:   VOIDWELL_URI + "ps2/player/" + character.CharacterId,
		Color: 0x070707,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: createCensusImageURI(character.FactionImageId),
		},
		Fields: fields,
	}

	p.RLock()
	client.SendEmbedMessage(message.Channel(), embed)
	p.RUnlock()
}

func (p *planetsidetwoPlugin) runCharacterWeaponStatsCommand(bot *discordgobot.Gobot, client *discordgobot.DiscordClient, payload discordgobot.CommandPayload) {
	trigger, args, message := payload.Trigger, payload.Arguments, payload.Message

	if trigger == "ps2c-ps4us" {
		args["platform"] = "ps4us"
	} else if trigger == "ps2c-ps4eu" {
		args["platform"] = "ps4eu"
	} else {
		args["platform"] = "pc"
	}

	resp, err := voidwellAPIGet(fmt.Sprintf("https://api.voidwell.com/ps2/character/byname/%s/weapon/%s?platform=%s", args["characterName"], args["weaponName"], args["platform"]))

	if err != nil {
		p.RLock()
		client.SendMessage(message.Channel(), fmt.Sprintf("%s", err))
		p.RUnlock()
		return
	}

	var weapon PlanetsideCharacterWeapon
	json.Unmarshal(resp, &weapon)

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: weapon.CharacterName + " [" + weapon.WeaponName + "]",
		},
		Title: "Click here for full stats",
		URL:   VOIDWELL_URI + "ps2/player/" + weapon.CharacterId,
		Color: 0x070707,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: createCensusImageURI(weapon.WeaponImageId),
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Kills",
				Value:  fmt.Sprintf("%d", weapon.Kills),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Deaths",
				Value:  fmt.Sprintf("%d", weapon.Deaths),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Play Time",
				Value:  fmt.Sprintf("%d Minutes", weapon.PlayTime/60),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Score",
				Value:  fmt.Sprintf("%d", weapon.Score),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "KpH",
				Value:  fmt.Sprintf("%0.2f", weapon.KillsPerHour),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Δ",
				Value:  weapon.KillsPerHourGrade,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "KDR",
				Value:  fmt.Sprintf("%0.2f", weapon.KillDeathRatio),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Δ",
				Value:  weapon.KillDeathRatioGrade,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "HSR",
				Value:  fmt.Sprintf("%0.2f%%", weapon.HeadshotRatio*100),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Δ",
				Value:  weapon.HeadshotRatioGrade,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Accuracy",
				Value:  fmt.Sprintf("%0.2f%%", weapon.Accuracy*100),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Δ",
				Value:  weapon.AccuracyGrade,
				Inline: true,
			},
		},
	}

	p.RLock()
	client.SendEmbedMessage(message.Channel(), embed)
	p.RUnlock()
}

func (p *planetsidetwoPlugin) runOutfitStatsCommand(bot *discordgobot.Gobot, client *discordgobot.DiscordClient, payload discordgobot.CommandPayload) {
	trigger, args, message := payload.Trigger, payload.Arguments, payload.Message

	if trigger == "ps2c-ps4us" {
		args["platform"] = "ps4us"
	} else if trigger == "ps2c-ps4eu" {
		args["platform"] = "ps4eu"
	} else {
		args["platform"] = "pc"
	}

	resp, err := voidwellAPIGet(fmt.Sprintf("https://api.voidwell.com/ps2/outfit/byalias/%s?platform=%s", args["outfitAlias"], args["platform"]))

	if err != nil {
		p.RLock()
		client.SendMessage(message.Channel(), fmt.Sprintf("%s", err))
		p.RUnlock()
		return
	}

	var outfit PlanetsideOutfit
	json.Unmarshal(resp, &outfit)

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: "[" + outfit.Alias + "] " + outfit.Name,
		},
		Title: "Click here for full stats",
		URL:   VOIDWELL_URI + "ps2/outfit/" + outfit.OutfitId,
		Color: 0x070707,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: createCensusImageURI(outfit.FactionImageId),
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Server",
				Value:  outfit.WorldName,
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Leader",
				Value:  outfit.LeaderName,
				Inline: false,
			},
			&discordgo.MessageEmbedField{
				Name:   "Member Count",
				Value:  fmt.Sprintf("%d", outfit.MemberCount),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Activity 7 Days",
				Value:  fmt.Sprintf("%d", outfit.Activity7Days),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Activity 30 Days",
				Value:  fmt.Sprintf("%d", outfit.Activity30Days),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "Activity 90 Days",
				Value:  fmt.Sprintf("%d", outfit.Activity90Days),
				Inline: true,
			},
		},
	}

	p.RLock()
	client.SendEmbedMessage(message.Channel(), embed)
	p.RUnlock()
}

func (p *planetsidetwoPlugin) runWeaponStatsCommand(bot *discordgobot.Gobot, client *discordgobot.DiscordClient, payload discordgobot.CommandPayload) {
	args, message := payload.Arguments, payload.Message

	resp, err := voidwellAPIGet(fmt.Sprintf("https://api.voidwell.com/ps2/weaponinfo/byname/%s", args["weaponName"]))

	if err != nil {
		p.RLock()
		client.SendMessage(message.Channel(), fmt.Sprintf("%s", err))
		p.RUnlock()
		return
	}

	var weapon PlanetsideWeapon
	json.Unmarshal(resp, &weapon)

	fields := make([]*discordgo.MessageEmbedField, 0)

	factionRestriction := "None"
	if weapon.FactionID > 0 {
		factionRestriction = getFactionName(weapon.FactionID)
	}

	fields = append(fields,
		&discordgo.MessageEmbedField{
			Name:   "Type",
			Value:  weapon.Category,
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Faction restriction",
			Value:  factionRestriction,
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Range",
			Value:  weapon.Range,
			Inline: true,
		},
	)

	if weapon.FireRateMs > 0 {
		fields = append(fields,
			&discordgo.MessageEmbedField{
				Name:   "Fire rate",
				Value:  fmt.Sprintf("%d RPM (%0.2f s)", 60000/weapon.FireRateMs, float32(weapon.FireRateMs)/1000),
				Inline: true,
			},
		)
	}

	if weapon.DamageRadius > 0 {
		fields = append(fields,
			&discordgo.MessageEmbedField{
				Name:   "Damage radius",
				Value:  fmt.Sprintf("%d", weapon.DamageRadius),
				Inline: true,
			},
		)
	}

	fields = append(fields,
		&discordgo.MessageEmbedField{
			Name:   "Muzzle velocity",
			Value:  fmt.Sprintf("%d m/sec", weapon.MuzzleVelocity),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Reload speed",
			Value:  fmt.Sprintf("%0.3f sec / %0.3f sec", float32(weapon.MinReloadSpeed)/1000, float32(weapon.MaxReloadSpeed)/1000),
			Inline: true,
		},
		&discordgo.MessageEmbedField{
			Name:   "Ammunition",
			Value:  fmt.Sprintf("%d / %d", weapon.ClipSize, weapon.Capacity),
			Inline: true,
		},
	)

	if weapon.IronSightZoom > 0 {
		fields = append(fields,
			&discordgo.MessageEmbedField{
				Name:   "Iron sight zoom",
				Value:  fmt.Sprintf("%0.2f", weapon.IronSightZoom),
				Inline: true,
			},
		)
	}

	fields = append(fields,
		&discordgo.MessageEmbedField{
			Name:   "Fire modes",
			Value:  strings.Join(weapon.FireModes, " / "),
			Inline: false,
		},
	)

	if !weapon.IsVehicleWeapon {
		fields = append(fields,
			&discordgo.MessageEmbedField{
				Name:   "Damage",
				Value:  fmt.Sprintf("%d / %dm / %d / %dm", weapon.MaxDamage, weapon.MaxDamageRange, weapon.MinDamage, weapon.MinDamageRange),
				Inline: false,
			},
		)
	}

	if !weapon.IsVehicleWeapon && weapon.IndirectMaxDamage > 0 {
		fields = append(fields,
			&discordgo.MessageEmbedField{
				Name:   "Indirect damage",
				Value:  fmt.Sprintf("%d / %0.fm / %d / %0.fm", weapon.IndirectMaxDamage, weapon.IndirectMaxDamageRange, weapon.IndirectMinDamage, weapon.IndirectMinDamageRange),
				Inline: false,
			},
		)
	}

	if !weapon.IsVehicleWeapon && weapon.HipAcc != nil {
		fields = append(fields,
			&discordgo.MessageEmbedField{
				Name:   "Hip accuracy",
				Value:  fmt.Sprintf("%0.2f / %0.2f / %0.2f / %0.2f / %0.2f", weapon.HipAcc.Crouching, weapon.HipAcc.CrouchWalking, weapon.HipAcc.Standing, weapon.HipAcc.Running, weapon.HipAcc.Cof),
				Inline: false,
			},
		)
	}

	if !weapon.IsVehicleWeapon && weapon.AimAcc != nil {
		fields = append(fields,
			&discordgo.MessageEmbedField{
				Name:   "Aim accuracy",
				Value:  fmt.Sprintf("%0.2f / %0.2f / %0.2f / %0.2f / %0.2f", weapon.AimAcc.Crouching, weapon.AimAcc.CrouchWalking, weapon.AimAcc.Standing, weapon.AimAcc.Running, weapon.AimAcc.Cof),
				Inline: false,
			},
		)
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: weapon.Name,
		},
		Title: "Click here for full stats",
		URL:   fmt.Sprintf("%sps2/item/%d", VOIDWELL_URI, weapon.ItemID),
		Color: 0x070707,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: createCensusImageURI(weapon.ImageID),
		},
		Description: weapon.Description,
		Fields:      fields,
	}

	p.RLock()
	client.SendEmbedMessage(message.Channel(), embed)
	p.RUnlock()
}

func createCensusImageURI(imageID int) string {
	return CENSUS_IMAGEBASE_URI + fmt.Sprintf("%v", imageID) + ".png"
}

func insertSlice(arr []*discordgo.MessageEmbedField, value *discordgo.MessageEmbedField, index int) []*discordgo.MessageEmbedField {
	return append(arr[:index], append([]*discordgo.MessageEmbedField{value}, arr[index:]...)...)
}

func voidwellAPIGet(uri string) (json.RawMessage, error) {
	if voidwellClient == nil {
		voidwellClientConfig := clientcredentials.Config{
			ClientID:     os.Getenv("VoidwellClientId"),
			ClientSecret: os.Getenv("VoidwellClientSecret"),
			TokenURL:     "https://auth.voidwell.com/connect/token",
			Scopes:       []string{"voidwell-daybreakgames", "voidwell-api"},
		}

		ctx := context.Background()
		voidwellClient = voidwellClientConfig.Client(ctx)
	}

	resp, err := voidwellClient.Get(uri)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	var jsonResponse json.RawMessage
	err = json.Unmarshal(body, &jsonResponse)

	if err != nil {
		log.Println(fmt.Sprintf("Failed to unmarshal for %v: %v", uri, err))
		return nil, err
	}

	return jsonResponse, nil
}

func getFactionName(factionID int) string {
	switch factionID {
	case 1:
		return "Vanu Sovereignty"
	case 2:
		return "New Conglomerate"
	case 3:
		return "Terran Republic"
	case 4:
		return "NS Operatives"
	}

	return "Unknown"
}

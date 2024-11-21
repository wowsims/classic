package database

import (
	"regexp"

	"github.com/wowsims/classic/sim/core/proto"
)

var OtherItemIdsToFetch = []string{}

var ItemOverrides = []*proto.UIItem{
	// Valentine's day event rewards
	// {Id: 51804, Phase: 2},

	// SOD Items
	{Id: 10019, Sources: []*proto.UIItemSource{{
		Source: &proto.UIItemSource_Crafted{
			Crafted: &proto.CraftedSource{
				Profession: proto.Profession_Tailoring, SpellId: 3759,
			},
		},
	}}},

	// Updated profession items not updated in the AtlasLoot DB
	// Crimson Silk Robe
	{Id: 217245, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Tailoring, SpellId: 439085}}}}},
	// Black Mageweave Vest
	{Id: 217246, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Tailoring, SpellId: 439086}}}}},
	// Long Silken Cloak
	{Id: 217252, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Tailoring, SpellId: 439094}}}}},
	// Enchanter's Cowl
	{Id: 217257, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Tailoring, SpellId: 439102}}}}},
	// Big Voodoo Mask
	{Id: 217259, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Leatherworking, SpellId: 439105}}}}},
	// Big Voodoo Robe
	{Id: 217261, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Leatherworking, SpellId: 439108}}}}},
	// Turtle Scale Breastplate
	{Id: 217268, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Leatherworking, SpellId: 439116}}}}},
	// Turtle Scale Gloves
	{Id: 217270, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Leatherworking, SpellId: 439118}}}}},
	// Golden Scale Cuirass
	{Id: 217277, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Blacksmithing, SpellId: 439124}}}}},
	// Golden Scale Coif
	{Id: 217279, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Blacksmithing, SpellId: 439126}}}}},
	// Golden Scale Leggings
	{Id: 217285, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Blacksmithing, SpellId: 439132}}}}},

	// The item tooltip is missing the usual Libram tag
	{Id: 221457, RangedWeaponType: proto.RangedWeaponType_RangedWeaponTypeLibram},

	// The item tooltip is missing the usual Totem tag
	{Id: 221464, RangedWeaponType: proto.RangedWeaponType_RangedWeaponTypeTotem},

	// SoD Gnomeregan Quest Necklaces are missing quest info from the gear planner DB
	{Id: 213343, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80324, Name: "The Mad King"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80325, Name: "The Mad King"}}},
	}},
	{Id: 213344, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80324, Name: "The Mad King"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80325, Name: "The Mad King"}}},
	}},
	{Id: 213345, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80324, Name: "The Mad King"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80325, Name: "The Mad King"}}},
	}},
	{Id: 213346, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80324, Name: "The Mad King"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80325, Name: "The Mad King"}}},
	}},

	// SoD Sunken Temple Drakeclaw Bands are missing quest info from the gear planner DB
	{Id: 220626, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82081, Name: "A Broken Ritual"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82083, Name: "A Broken Ritual"}}},
	}},
	{Id: 220627, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82081, Name: "A Broken Ritual"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82083, Name: "A Broken Ritual"}}},
	}},
	{Id: 220628, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82081, Name: "A Broken Ritual"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82083, Name: "A Broken Ritual"}}},
	}},
	{Id: 220629, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82081, Name: "A Broken Ritual"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82083, Name: "A Broken Ritual"}}},
	}},
	{Id: 220630, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82081, Name: "A Broken Ritual"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82083, Name: "A Broken Ritual"}}},
	}},
}

// Keep these sorted by item ID.
var ItemAllowList = map[int32]struct{}{
	14637:  {}, // https://www.wowhead.com/classic/item=14637/cadaverous-armor
	19099:  {}, // https://www.wowhead.com/classic/item=19099/glacial-blade filtered by temp naxx Glacial gear filters
	22335:  {}, // https://www.wowhead.com/classic/item=22335/lord-valthalaks-staff-of-command accidentally left in the loot pool for a while. Allowing for compatibility
	22395:  {}, // https://www.wowhead.com/classic/item=22395/totem-of-rage
	221783: {}, // https://www.wowhead.com/classic/item=221783/lawbringer-spaulders

	// These are all filtered out by the SoD duplicates filter because of new versions added in SoD
	20425: {}, // https://www.wowhead.com/classic/item=20425/advisors-gnarled-staff
	20430: {}, // https://www.wowhead.com/classic/item=20430/legionnaires-sword
	20434: {}, // https://www.wowhead.com/classic/item=20434/lorekeepers-staff
	20437: {}, // https://www.wowhead.com/classic/item=20437/outriders-bow
	20438: {}, // https://www.wowhead.com/classic/item=20438/outrunners-bow
	20440: {}, // https://www.wowhead.com/classic/item=20440/protectors-sword
	20441: {}, // https://www.wowhead.com/classic/item=20441/scouts-blade
	20443: {}, // https://www.wowhead.com/classic/item=20443/sentinels-blade
}

// Keep these sorted by item ID.
var ItemDenyList = map[int32]struct{}{
	9653:   {}, // Speedy Racer Goggles
	12104:  {}, // Brindlethorn Tunic
	12805:  {}, // Orb of Fire
	17782:  {}, // talisman of the binding shard
	17783:  {}, // talisman of the binding fragment
	17802:  {}, // Deprecated version of Thunderfury
	20522:  {}, // Feral Staff
	22736:  {}, // Andonisus, Reaper of Souls
	34576:  {}, // Battlemaster's Cruelty
	34577:  {}, // Battlemaster's Depreavity
	34578:  {}, // Battlemaster's Determination
	34579:  {}, // Battlemaster's Audacity
	34580:  {}, // Battlemaster's Perseverence
	206382: {}, // Tempest Icon
	206387: {}, // Kajaric Icon
	206954: {}, // Idol of Ursine Rage
	208689: {}, // Ferocious Idol
	208849: {}, // Libram of Blessings
	208851: {}, // Libram of Justice
	210195: {}, // Unbalanced Idol
	210534: {}, // Idol of the Wild
	211472: {}, // Libram of Banishment
	211501: {}, // https://www.wowhead.com/classic/item=211501/chestguard-of-might
	213513: {}, // Libram of Deliverance
	213594: {}, // Idol of the Heckler
	215116: {}, // UNUSED - Hyperconductive Speed Belt
	220915: {}, // Idol of the Raging Shambler
	227444: {}, // Idol of the Huntress
	227843: {}, // https://www.wowhead.com/classic/item=227843/reaving-nightfall Removed from SoD
	227954: {}, // https://www.wowhead.com/classic/item=227954/boreal-mantle unused item
	227966: {}, // https://www.wowhead.com/classic/item=227966/naglering unused item
	227977: {}, // https://www.wowhead.com/classic/item=227977/totem-of-rage unused item
	227989: {}, // https://www.wowhead.com/classic/item=227989/hand-of-justice unused item
	227995: {}, // https://www.wowhead.com/classic/item=227995/cadaverous-armor unused item
	228498: {}, // Unused Dreadblade of the Destructor
}

// Item icons to include in the DB, so they don't need to be separately loaded in the UI.
var ExtraItemIcons = []int32{
	// Demonic Rune
	12662,

	// Explosives
	13180,
	11566,
	8956,
	10646,
	18641,
	15993,
	16040,

	// Food IDs
	13928,
	20452,
	13931,
	18254,
	21023,
	13813,
	13810,

	// Flask IDs
	13510,
	13511,
	13512,
	13513,

	// Zanza
	20079,

	// Blasted Lands
	8412,
	8423,
	8424,
	8411,

	// Agility Elixer IDs
	13452,
	9187,

	// Single Elixirs
	20007, // Mana Regen Elixir
	20004, // Major Troll's Blood Potion
	9088,  // Gift of Arthas

	// Armor Elixirs
	3389,  // Defense
	8951,  // Greater
	13445, // Superior Defense

	// Health Elixirs
	2458, // Minor Fortitude
	3825, // Fortitude

	// Strength
	12451,
	9206,

	// AP
	12460,
	12820,

	// Random
	5206, // Bogling Root

	// SP
	13454,
	9264,
	21546,
	17708,

	// Crystal
	11564, // Armor

	// Alcohol Buff
	18284,
	18269,
	20709,
	21114,
	21151,

	// Potions / In Battle Consumes
	13444,

	// Thistle Tea
	7676,

	// Weapon Oils
	20748,
	20749,
	12404,
	18262,
}

var SpellIconoverrides = []*proto.IconData{}

// Raid buffs / debuffs
var SharedSpellsIcons = []int32{
	// World Buffs
	22888, // Ony / Nef
	24425, // Spirit
	16609, // Warchief
	23768, // DMF Damage
	23736, // DMF Agi
	23766, // DMF Int
	23738, // DMF Spirit
	23737, // DMF Stam

	22818, // DM Stam
	22820, // DM Spell Crit
	22817, // DM AP

	15366, // Songflower

	29534, // Silithus

	18264, // Headmasters

	// Registered CD's
	10060, // Power Infusion
	29166, // Innervate

	// Mark
	1126,
	5232,
	6756,
	5234,
	8907,
	9884,
	9885,
	17055,

	20217, // Kings (Talent)
	25898, // Greater Kings
	25899, // Sanctuary

	10293, // Devo Aura
	20142, // Imp. Devo

	// Stoneskin Totem
	10408,
	16293,

	// Fort
	1243,
	1244,
	1245,
	2791,
	10937,
	10938,
	14767,

	// Spirit
	14752,
	14818,
	14819,
	27841,

	// Might
	19740,
	19834,
	19835,
	19836,
	19837,
	19838,
	25291,
	20048,

	// Commanding Shout
	6673,
	5242,
	6192,
	11549,
	11550,
	11551,
	25289,
	12861,

	// AP
	30811, // Unleashed Rage
	19506, // Trueshot

	// Battle Shout
	6673,
	5242,
	6192,
	11549,
	11550,
	11551,
	25289,
	12861, // Imp

	// Wisdom
	19742,
	19850,
	19852,
	19853,
	19854,
	25290,
	20245,

	// Mana Spring
	5675,
	10495,
	10496,
	10497,

	17007, // Leader of the Pack
	24858, // Moonkin

	// Windfury
	8512,
	10613,
	10614,
	29193, // Imp WF

	// Raid Debuffs
	8647,
	7386,
	7405,
	8380,
	11596,
	11597,

	770,
	778,
	9749,
	9907,
	11708,
	18181,

	26016,
	12879,
	9452,
	26021,
	16862,
	9747,
	9898,

	3043,
	14275,
	14276,
	14277,

	17800,
	17803,
	12873,
	28593,

	11374,
	15235,

	24977,
}

// If any of these match the item name, don't include it.
var DenyListNameRegexes = []*regexp.Regexp{
	regexp.MustCompile(`30 Epic`),
	regexp.MustCompile(`63 Blue`),
	regexp.MustCompile(`63 Green`),
	regexp.MustCompile(`66 Epic`),
	regexp.MustCompile(`90 Epic`),
	regexp.MustCompile(`90 Green`),
	regexp.MustCompile(`Boots 1`),
	regexp.MustCompile(`Boots 2`),
	regexp.MustCompile(`Boots 3`),
	regexp.MustCompile(`Bracer 1`),
	regexp.MustCompile(`Bracer 2`),
	regexp.MustCompile(`Bracer 3`),
	regexp.MustCompile(`DB\d`),
	regexp.MustCompile(`DEPRECATED`),
	regexp.MustCompile(`Deprecated: Keanna`),
	regexp.MustCompile(`Indalamar`),
	regexp.MustCompile(`Monster -`),
	regexp.MustCompile(`NEW`),
	regexp.MustCompile(`PH`),
	regexp.MustCompile(`QR XXXX`),
	regexp.MustCompile(`TEST`),
	regexp.MustCompile(`Test`),
	regexp.MustCompile(`zOLD`),

	// PVP Gear
	// regexp.MustCompile(`Grand Marshal's [a-zA-z\s]+`),
	// regexp.MustCompile(`High Warlord's [a-zA-z\s]+`),

	// AQ
	// regexp.MustCompile(`Qiraji`),
	// regexp.MustCompile(`[A-Za-z\s]+ of the Bronze Dragonflight`),
	// regexp.MustCompile(`[A-Za-z\s]+ of the Fallen God`),
	// regexp.MustCompile(`Belt of [A-Za-z]+ Heads`),

	// Naxx
	regexp.MustCompile(`Icebane`),
	regexp.MustCompile(`Icy Scale`),
	regexp.MustCompile(`Polar`),
	regexp.MustCompile(`Glacial`),
	regexp.MustCompile(`Mark of the Champion`),
	regexp.MustCompile(`Atiesh`),
}

// Data can easily be found here:
// https://www.wowhead.com/classic/item-sets#item-sets
var DenyItemSetIds = []int{}

package dpsrogue

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterDpsRogue()
}

func TestDaggers(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassRogue,
			Level:      60,
			Phase:      5,
			Race:       proto.Race_RaceHuman,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     CombatDaggers60Talents,
			GearSet:     core.GetGearSet("../../../ui/rogue/gear_sets", "combat_daggers_60"),
			Rotation:    core.GetAplRotation("../../../ui/rogue/apls", "combat_daggers_60"),
			Buffs:       core.FullBuffsPhase1,
			Consumes:    Phase1Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "No Poisons", SpecOptions: DefaultCombatRogue},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,
		},
	}))
}

func TestSwords(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassRogue,
			Level:      60,
			Phase:      5,
			Race:       proto.Race_RaceHuman,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents:     CombatSwords60Talents,
			GearSet:     core.GetGearSet("../../../ui/rogue/gear_sets", "combat_swords_60"),
			Rotation:    core.GetAplRotation("../../../ui/rogue/apls", "combat_swords_60"),
			Buffs:       core.FullBuffsPhase1,
			Consumes:    Phase1Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "No Poisons", SpecOptions: DefaultCombatRogue},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,
		},
	}))
}

var CombatSwords60Talents = "005323105-0230251020050150231"
var CombatDaggers60Talents = "005323105-0253051020550100201"

var ItemFilters = core.ItemFilter{
	ArmorType: proto.ArmorType_ArmorTypeLeather,
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeDagger,
		proto.WeaponType_WeaponTypeFist,
		proto.WeaponType_WeaponTypeSword,
		proto.WeaponType_WeaponTypeMace,
	},
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeBow,
		proto.RangedWeaponType_RangedWeaponTypeCrossbow,
		proto.RangedWeaponType_RangedWeaponTypeGun,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatAttackPower,
	proto.Stat_StatAgility,
	proto.Stat_StatStrength,
	proto.Stat_StatMeleeHit,
	proto.Stat_StatMeleeCrit,
}

var DefaultAssassinationRogue = &proto.Player_Rogue{
	Rogue: &proto.Rogue{
		Options: DefaultDeadlyBrewOptions,
	},
}

var DefaultCombatRogue = &proto.Player_Rogue{
	Rogue: &proto.Rogue{
		Options: DefaultDeadlyBrewOptions,
	},
}

var DefaultDeadlyBrewOptions = &proto.RogueOptions{}

var Phase1Consumes = core.ConsumesCombo{
	Label: "Classic Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir: proto.AgilityElixir_ElixirOfTheMongoose,
		AttackPowerBuff: proto.AttackPowerBuff_JujuMight,
		DefaultConjured: proto.Conjured_ConjuredRogueThistleTea,
		Flask: proto.Flask_FlaskOfSupremePower,
		Food: proto.Food_FoodGrilledSquid,
		MainHandImbue: proto.WeaponImbue_Windfury,
		OffHandImbue: proto.WeaponImbue_InstantPoison,
		StrengthBuff: proto.StrengthBuff_JujuPower,
		ZanzaBuff: proto.ZanzaBuff_GroundScorpokAssay,
	},
}
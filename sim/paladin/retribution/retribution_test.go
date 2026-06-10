package retribution

import (
	"testing"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterRetributionPaladin()
}

func TestRetribution(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassPaladin,
			Phase:      5,
			Race:       proto.Race_RaceHuman,
			OtherRaces: []proto.Race{proto.Race_RaceDwarf},

			Talents:        Phase45RetTalents,
			GearSet:        core.GetGearSet("../../../ui/retribution_paladin/gear_sets", "blank"),
			Rotation:       core.GetAplRotation("../../../ui/retribution_paladin/apls", "basic_ret"),
			OtherRotations: []core.RotationCombo{core.GetAplRotation("../../../ui/retribution_paladin/apls", "basic_ret")},
			Buffs:          core.FullBuffs,
			Consumes:       Phase5Consumes,
			SpecOptions:    core.SpecOptionsCombo{Label: "P5 Seal of Righteousness Ret", SpecOptions: PlayerOptionsSealofRighteousness},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,
		},
	}))
}

var Phase45RetTalents = "500501-503-52230351200315"

var Phase1Consumes = core.ConsumesCombo{
	Label: "P1-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir: proto.AgilityElixir_ElixirOfLesserAgility,
		DefaultPotion: proto.Potions_ManaPotion,
		FirePowerBuff: proto.FirePowerBuff_ElixirOfFirepower,
		//MainHandImbue: proto.WeaponImbue_WildStrikes,
		StrengthBuff: proto.StrengthBuff_ElixirOfOgresStrength,
	},
}

var Phase2Consumes = core.ConsumesCombo{
	Label: "P2-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir:     proto.AgilityElixir_ElixirOfAgility,
		DefaultPotion:     proto.Potions_ManaPotion,
		DragonBreathChili: true,
		FirePowerBuff:     proto.FirePowerBuff_ElixirOfFirepower,
		Food:              proto.Food_FoodSagefishDelight,
		MainHandImbue:     proto.WeaponImbue_WindfuryWeapon,
		SpellPowerBuff:    proto.SpellPowerBuff_LesserArcaneElixir,
		StrengthBuff:      proto.StrengthBuff_ElixirOfOgresStrength,
	},
}

var Phase5Consumes = core.ConsumesCombo{
	Label: "P5-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir:     proto.AgilityElixir_ElixirOfTheMongoose,
		AttackPowerBuff:   proto.AttackPowerBuff_JujuMight,
		DefaultPotion:     proto.Potions_MajorManaPotion,
		DragonBreathChili: true,
		Flask:             proto.Flask_FlaskOfSupremePower,
		FirePowerBuff:     proto.FirePowerBuff_ElixirOfFirepower,
		Food:              proto.Food_FoodSmokedDesertDumpling,
		//MainHandImbue:     proto.WeaponImbue_WildStrikes,
		//OffHandImbue:      proto.WeaponImbue_MagnificentTrollshine,
		SpellPowerBuff: proto.SpellPowerBuff_GreaterArcaneElixir,
		StrengthBuff:   proto.StrengthBuff_JujuPower,
	},
}

var PlayerOptionsSealofCommand = &proto.Player_RetributionPaladin{
	RetributionPaladin: &proto.RetributionPaladin{
		Options: optionsSealOfCommand,
	},
}

var PlayerOptionsSealofRighteousness = &proto.Player_RetributionPaladin{
	RetributionPaladin: &proto.RetributionPaladin{
		Options: optionsSealOfRighteousness,
	},
}

var optionsSealOfCommand = &proto.PaladinOptions{
	PrimarySeal: proto.PaladinSeal_Command,
}

var optionsSealOfRighteousness = &proto.PaladinOptions{
	PrimarySeal: proto.PaladinSeal_Righteousness,
}

var ItemFilters = core.ItemFilter{
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeAxe,
		proto.WeaponType_WeaponTypeSword,
		proto.WeaponType_WeaponTypeMace,
		proto.WeaponType_WeaponTypePolearm,
		proto.WeaponType_WeaponTypeShield,
	},
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeLibram,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatStrength,
	proto.Stat_StatAgility,
	proto.Stat_StatAttackPower,
	proto.Stat_StatMeleeHit,
	proto.Stat_StatMeleeCrit,
	proto.Stat_StatSpellPower,
	proto.Stat_StatSpellHit,
	proto.Stat_StatSpellCrit,
}

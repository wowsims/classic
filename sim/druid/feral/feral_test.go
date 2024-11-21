package feral

import (
	"testing"

	_ "github.com/wowsims/classic/sim/common"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterFeralDruid()
}

func TestFeral(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassDruid,
			Phase:      5,
			Level:      60,
			Race:       proto.Race_RaceTauren,
			OtherRaces: []proto.Race{proto.Race_RaceNightElf},

			Talents:     Phase4Talents,
			GearSet:     core.GetGearSet("../../../ui/feral_druid/gear_sets", "blank"),
			Rotation:    core.GetAplRotation("../../../ui/feral_druid/apls", "phase_5"),
			Buffs:       core.FullBuffsPhase5,
			Consumes:    Phase4Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Default", SpecOptions: PlayerOptionsMonoCat},
			OtherSpecOptions: []core.SpecOptionsCombo{
				{Label: "Default-NoBleed", SpecOptions: PlayerOptionsMonoCatNoBleed},
				{Label: "Flower-Aoe", SpecOptions: PlayerOptionsFlowerCatAoe},
			},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,
		},
	}))
}

var Phase1Talents = "500005001--05"
var Phase2Talents = "-550002032320211-05"
var Phase3Talents = "500005301-5500020323002-05"
var Phase4Talents = "500005301-5500020323202151-15"

var PlayerOptionsMonoCat = &proto.Player_FeralDruid{
	FeralDruid: &proto.FeralDruid{
		Options: &proto.FeralDruid_Options{
			InnervateTarget:   &proto.UnitReference{}, // no Innervate
			LatencyMs:         100,
			AssumeBleedActive: true,
		},
	},
}

var PlayerOptionsMonoCatNoBleed = &proto.Player_FeralDruid{
	FeralDruid: &proto.FeralDruid{
		Options: &proto.FeralDruid_Options{
			InnervateTarget:   &proto.UnitReference{}, // no Innervate
			LatencyMs:         100,
			AssumeBleedActive: false,
		},
	},
}

var PlayerOptionsFlowerCatAoe = &proto.Player_FeralDruid{
	FeralDruid: &proto.FeralDruid{
		Options: &proto.FeralDruid_Options{
			InnervateTarget:   &proto.UnitReference{}, // no Innervate
			LatencyMs:         100,
			AssumeBleedActive: false,
		},
	},
}

var Phase1Consumes = core.ConsumesCombo{
	Label: "P1-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir:   proto.AgilityElixir_ElixirOfLesserAgility,
		DefaultConjured: proto.Conjured_ConjuredMinorRecombobulator,
		DefaultPotion:   proto.Potions_ManaPotion,
		Food:            proto.Food_FoodSmokedSagefish,
		MainHandImbue:   proto.WeaponImbue_WildStrikes,
		StrengthBuff:    proto.StrengthBuff_ElixirOfOgresStrength,
	},
}

var Phase2Consumes = core.ConsumesCombo{
	Label: "P2-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir:     proto.AgilityElixir_ElixirOfAgility,
		DefaultPotion:     proto.Potions_GreaterManaPotion,
		DragonBreathChili: true,
		Food:              proto.Food_FoodSagefishDelight,
		MainHandImbue:     proto.WeaponImbue_WildStrikes,
		StrengthBuff:      proto.StrengthBuff_ElixirOfOgresStrength,
	},
}

var Phase3Consumes = core.ConsumesCombo{
	Label: "P3-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir:     proto.AgilityElixir_ElixirOfTheMongoose,
		DefaultPotion:     proto.Potions_MajorManaPotion,
		DragonBreathChili: true,
		Food:              proto.Food_FoodSmokedDesertDumpling,
		MainHandImbue:     proto.WeaponImbue_WildStrikes,
		MiscConsumes: &proto.MiscConsumes{
			Catnip: true,
		},
		StrengthBuff: proto.StrengthBuff_ElixirOfGiants,
	},
}

var Phase4Consumes = core.ConsumesCombo{
	Label: "P4-Consumes",
	Consumes: &proto.Consumes{
		AgilityElixir:     proto.AgilityElixir_ElixirOfTheMongoose,
		AttackPowerBuff:   proto.AttackPowerBuff_JujuMight,
		DefaultConjured:   proto.Conjured_ConjuredDemonicRune,
		DefaultPotion:     proto.Potions_MajorManaPotion,
		DragonBreathChili: true,
		Flask:             proto.Flask_FlaskOfDistilledWisdom,
		Food:              proto.Food_FoodSmokedDesertDumpling,
		MainHandImbue:     proto.WeaponImbue_ElementalSharpeningStone,
		MiscConsumes: &proto.MiscConsumes{
			Catnip: true,
		},
		StrengthBuff: proto.StrengthBuff_JujuPower,
	},
}

var ItemFilters = core.ItemFilter{
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeDagger,
		proto.WeaponType_WeaponTypeMace,
		proto.WeaponType_WeaponTypeOffHand,
		proto.WeaponType_WeaponTypeStaff,
		proto.WeaponType_WeaponTypePolearm,
	},
	ArmorType: proto.ArmorType_ArmorTypeLeather,
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeIdol,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatStrength,
	proto.Stat_StatAgility,
	proto.Stat_StatAttackPower,
	proto.Stat_StatMeleeCrit,
	proto.Stat_StatMeleeHit,
}

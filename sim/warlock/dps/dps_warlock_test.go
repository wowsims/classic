package dps

import (
	"testing"

	_ "github.com/wowsims/classic/sim/common"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterDpsWarlock()
}

func TestAffliction(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassWarlock,
			Phase: 4,
			Level: 60,
			Race:  proto.Race_RaceOrc,

			Talents:     Phase4AffTalents,
			GearSet:     core.GetGearSet("../../../ui/warlock/gear_sets", "placeholder"),
			Rotation:    core.GetAplRotation("../../../ui/warlock/apls/p4", "affliction"),
			Buffs:       core.FullBuffsPhase4,
			Consumes:    Phase4Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Affliction Warlock", SpecOptions: DefaultAfflictionWarlock},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
	}))
}

func TestDemonology(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassWarlock,
			Level: 40,
			Race:  proto.Race_RaceOrc,

			Talents:     Phase2DemonologyTalents,
			GearSet:     core.GetGearSet("../../../ui/warlock/gear_sets", "placeholder"),
			Rotation:    core.GetAplRotation("../../../ui/warlock/apls/p2", "demonology"),
			Buffs:       core.FullBuffsPhase2,
			Consumes:    Phase2Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Demonology Warlock", SpecOptions: DefaultDemonologyWarlock},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
	}))
}

func TestDestruction(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassWarlock,
			Phase: 4,
			Level: 60,
			Race:  proto.Race_RaceOrc,

			Talents:     Phase4DestroTalents,
			GearSet:     core.GetGearSet("../../../ui/warlock/gear_sets", "placeholder"),
			Rotation:    core.GetAplRotation("../../../ui/warlock/apls/p4", "destruction"),
			Buffs:       core.FullBuffsPhase4,
			Consumes:    Phase4Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "Destruction Warlock", SpecOptions: DefaultDestroWarlock},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
	}))
}

var Phase1DestructionTalents = "-03-0550201"

var Phase2AfflictionTalents = "3500253012201105--1"
var Phase2DemonologyTalents = "-2050033132501051"
var Phase2DestructionTalents = "-01-055020512000415"

var Phase3BackdraftTalents = "-032004-5050205102005151"
var Phase3NFRuinTalents = "25002500102-03-50502051020001"

var Phase4AffTalents = "4500253012201005--50502051020001"
var Phase4DestroTalents = "05002-035004-5050205102005151"

var DefaultDestroWarlock = &proto.Player_Warlock{
	Warlock: &proto.Warlock{
		Options: &proto.WarlockOptions{
			Armor:       proto.WarlockOptions_FelArmor,
			Summon:      proto.WarlockOptions_Imp,
			WeaponImbue: proto.WarlockOptions_NoWeaponImbue,
		},
	},
}

var DefaultAfflictionWarlock = &proto.Player_Warlock{
	Warlock: &proto.Warlock{
		Options: &proto.WarlockOptions{
			Armor:       proto.WarlockOptions_FelArmor,
			Summon:      proto.WarlockOptions_Imp,
			WeaponImbue: proto.WarlockOptions_NoWeaponImbue,
		},
	},
}

var DefaultDemonologyWarlock = &proto.Player_Warlock{
	Warlock: &proto.Warlock{
		Options: &proto.WarlockOptions{
			Armor:       proto.WarlockOptions_FelArmor,
			Summon:      proto.WarlockOptions_Felguard,
			WeaponImbue: proto.WarlockOptions_NoWeaponImbue,
		},
	},
}

var Phase1Consumes = core.ConsumesCombo{
	Label: "P1-Consumes",
	Consumes: &proto.Consumes{
		DefaultPotion: proto.Potions_ManaPotion,
		FirePowerBuff: proto.FirePowerBuff_ElixirOfFirepower,
		Food:          proto.Food_FoodSmokedSagefish,
		MainHandImbue: proto.WeaponImbue_BlackfathomManaOil,
	},
}

var Phase2Consumes = core.ConsumesCombo{
	Label: "P2-Consumes",
	Consumes: &proto.Consumes{
		DefaultPotion:  proto.Potions_ManaPotion,
		FirePowerBuff:  proto.FirePowerBuff_ElixirOfFirepower,
		Food:           proto.Food_FoodSagefishDelight,
		MainHandImbue:  proto.WeaponImbue_LesserWizardOil,
		SpellPowerBuff: proto.SpellPowerBuff_LesserArcaneElixir,
	},
}

var Phase3Consumes = core.ConsumesCombo{
	Label: "P3-Consumes",
	Consumes: &proto.Consumes{
		DefaultPotion:   proto.Potions_SuperiorManaPotion,
		FirePowerBuff:   proto.FirePowerBuff_ElixirOfFirepower,
		ShadowPowerBuff: proto.ShadowPowerBuff_ElixirOfShadowPower,
		Food:            proto.Food_FoodSagefishDelight,
		MainHandImbue:   proto.WeaponImbue_LesserWizardOil,
		SpellPowerBuff:  proto.SpellPowerBuff_GreaterArcaneElixir,
	},
}

var Phase4Consumes = core.ConsumesCombo{
	Label: "P4-Consumes",
	Consumes: &proto.Consumes{
		DefaultPotion:   proto.Potions_MajorManaPotion,
		Flask:           proto.Flask_FlaskOfSupremePower,
		FirePowerBuff:   proto.FirePowerBuff_ElixirOfGreaterFirepower,
		ShadowPowerBuff: proto.ShadowPowerBuff_ElixirOfShadowPower,
		Food:            proto.Food_FoodTenderWolfSteak,
		MainHandImbue:   proto.WeaponImbue_WizardOil,
		SpellPowerBuff:  proto.SpellPowerBuff_GreaterArcaneElixir,
	},
}

var ItemFilters = core.ItemFilter{
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeSword,
		proto.WeaponType_WeaponTypeDagger,
	},
	HandTypes: []proto.HandType{
		proto.HandType_HandTypeOffHand,
	},
	ArmorType: proto.ArmorType_ArmorTypeCloth,
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeWand,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatIntellect,
	proto.Stat_StatSpellPower,
	proto.Stat_StatSpellHit,
	proto.Stat_StatSpellCrit,
}

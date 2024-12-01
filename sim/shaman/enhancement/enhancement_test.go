package enhancement

import (
	"testing"

	_ "github.com/wowsims/classic/sim/common" // imported to get item effects included.
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func init() {
	RegisterEnhancementShaman()
}

func TestEnhancement(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassShaman,
			Phase:      5,
			Race:       proto.Race_RaceTroll,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},

			Talents: Phase4Talents,
			GearSet: core.GetGearSet("../../../ui/enhancement_shaman/gear_sets", "blank"),
			// OtherGearSets: []core.GearSetCombo{
			// 	core.GetGearSet("../../../ui/enhancement_shaman/gear_sets", "phase_5_2h"),
			// },
			Rotation:    core.GetAplRotation("../../../ui/enhancement_shaman/apls", "phase_5"),
			Buffs:       core.FullBuffs,
			Consumes:    Phase4ConsumesWFWF,
			SpecOptions: core.SpecOptionsCombo{Label: "Sync Auto", SpecOptions: PlayerOptionsSyncAuto},
			OtherSpecOptions: []core.SpecOptionsCombo{
				{Label: "Sync Delay OH", SpecOptions: PlayerOptionsSyncDelayOH},
			},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatAttackPower,
			StatsToWeigh:    Stats,
		},
	}))
}

var Phase1Talents = "-5005202101"
var Phase2Talents = "-5005202105023051"
var Phase3Talents = "05003-5005132105023051"
var Phase4Talents = "25003105003-5005032105023051"

var PlayerOptionsSyncDelayOH = &proto.Player_EnhancementShaman{
	EnhancementShaman: &proto.EnhancementShaman{
		Options: optionsSyncDelayOffhand,
	},
}

var PlayerOptionsSyncAuto = &proto.Player_EnhancementShaman{
	EnhancementShaman: &proto.EnhancementShaman{
		Options: optionsSyncAuto,
	},
}

var optionsSyncDelayOffhand = &proto.EnhancementShaman_Options{
	SyncType: proto.ShamanSyncType_DelayOffhandSwings,
}

var optionsSyncAuto = &proto.EnhancementShaman_Options{
	SyncType: proto.ShamanSyncType_Auto,
}

var Phase4ConsumesWFWF = core.ConsumesCombo{
	Label: "P4-Consumes WF/WF",
	Consumes: &proto.Consumes{
		AttackPowerBuff:   proto.AttackPowerBuff_JujuMight,
		AgilityElixir:     proto.AgilityElixir_ElixirOfTheMongoose,
		DefaultPotion:     proto.Potions_MajorManaPotion,
		DragonBreathChili: true,
		FirePowerBuff:     proto.FirePowerBuff_ElixirOfGreaterFirepower,
		Flask:             proto.Flask_FlaskOfSupremePower,
		Food:              proto.Food_FoodBlessSunfruit,
		MainHandImbue:     proto.WeaponImbue_WindfuryWeapon,
		OffHandImbue:      proto.WeaponImbue_WindfuryWeapon,
		SpellPowerBuff:    proto.SpellPowerBuff_GreaterArcaneElixir,
		StrengthBuff:      proto.StrengthBuff_JujuPower,
	},
}

var ItemFilters = core.ItemFilter{
	WeaponTypes: []proto.WeaponType{
		proto.WeaponType_WeaponTypeAxe,
		proto.WeaponType_WeaponTypeDagger,
		proto.WeaponType_WeaponTypeFist,
		proto.WeaponType_WeaponTypeMace,
		proto.WeaponType_WeaponTypeOffHand,
		proto.WeaponType_WeaponTypeShield,
		proto.WeaponType_WeaponTypeStaff,
	},
	ArmorType: proto.ArmorType_ArmorTypeMail,
	RangedWeaponTypes: []proto.RangedWeaponType{
		proto.RangedWeaponType_RangedWeaponTypeTotem,
	},
}

var Stats = []proto.Stat{
	proto.Stat_StatStrength,
	proto.Stat_StatAgility,
	proto.Stat_StatAttackPower,
	proto.Stat_StatMeleeHit,
	proto.Stat_StatMeleeCrit,
	proto.Stat_StatSpellPower,
}

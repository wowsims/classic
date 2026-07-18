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

func TestWarlockSMRuin(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassWarlock,
			Phase: 1,
			Race:  proto.Race_RaceOrc,

			Talents:     TalentsSMRuin,
			GearSet:     core.GetGearSet("../../../ui/warlock/gear_sets", "mc"),
			Rotation:    core.GetAplRotation("../../../ui/warlock/apls/", "rotation"),
			Buffs:       core.FullBuffs,
			Consumes:    Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "SM/Ruin Warlock", SpecOptions: DefaultDestroWarlock},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
	}))
}

func TestWarlockDSRuin(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class: proto.Class_ClassWarlock,
			Phase: 1,
			Race:  proto.Race_RaceOrc,

			Talents:     TalentsDSRuin,
			GearSet:     core.GetGearSet("../../../ui/warlock/gear_sets", "mc"),
			Rotation:    core.GetAplRotation("../../../ui/warlock/apls/", "rotation"),
			Buffs:       core.FullBuffs,
			Consumes:    Consumes,
			SpecOptions: core.SpecOptionsCombo{Label: "DS/Ruin Warlock", SpecOptions: DefaultDestroWarlock},

			ItemFilter:      ItemFilters,
			EPReferenceStat: proto.Stat_StatSpellPower,
			StatsToWeigh:    Stats,
		},
	}))
}

func TestLifeTapAssumedHealing(t *testing.T) {
	runWithHPS := func(hps float64) *proto.RaidSimResult {
		warlockOptions := *DefaultDestroWarlock.Warlock.Options
		warlockOptions.AssumedLifeTapHps = hps
		player := core.WithSpec(&proto.Player{
			Class:         proto.Class_ClassWarlock,
			Race:          proto.Race_RaceOrc,
			Equipment:     core.GetGearSet("../../../ui/warlock/gear_sets", "mc").GearSet,
			TalentsString: TalentsSMRuin,
			Rotation:      core.GetAplRotation("../../../ui/warlock/apls", "rotation").Rotation,
		}, &proto.Player_Warlock{Warlock: &proto.Warlock{Options: &warlockOptions}})

		result := core.RunRaidSim(&proto.RaidSimRequest{
			Raid:       core.SinglePlayerRaidProto(player, nil, nil, nil),
			Encounter:  core.MakeSingleTargetEncounter(0),
			SimOptions: &proto.SimOptions{Iterations: 1, IsTest: true, RandomSeed: 1},
		})
		if result.Error != nil {
			t.Fatalf("simulation failed: %s", result.Error.Message)
		}
		return result
	}

	withoutHealing := runWithHPS(0)
	withHealing := runWithHPS(20000)
	withoutHealingDPS := withoutHealing.RaidMetrics.Parties[0].Players[0].Dps.Avg
	withHealingDPS := withHealing.RaidMetrics.Parties[0].Players[0].Dps.Avg
	if withHealingDPS <= withoutHealingDPS {
		t.Fatalf("assumed healing did not improve DPS: without healing %.1f, with healing %.1f", withoutHealingDPS, withHealingDPS)
	}

	for _, resource := range withHealing.RaidMetrics.Parties[0].Players[0].Resources {
		if resource.Type == proto.ResourceType_ResourceTypeHealth &&
			resource.Id.GetOtherId() == proto.OtherAction_OtherActionHealingModel &&
			resource.ActualGain > 0 {
			t.Logf("DPS without healing %.1f, with healing %.1f; restored %.1f health", withoutHealingDPS, withHealingDPS, resource.ActualGain)
			return
		}
	}

	t.Fatal("assumed healing produced no effective health gain")
}

var TalentsSMRuin = "5502203112201105--52500051020001"
var TalentsDSRuin = "25002-2050300152201-52500051020001"

var DefaultDestroWarlock = &proto.Player_Warlock{
	Warlock: &proto.Warlock{
		Options: &proto.WarlockOptions{
			Armor:       proto.WarlockOptions_DemonArmor,
			Summon:      proto.WarlockOptions_Succubus,
			WeaponImbue: proto.WarlockOptions_NoWeaponImbue,
		},
	},
}

var Consumes = core.ConsumesCombo{
	Label: "Consumes",
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

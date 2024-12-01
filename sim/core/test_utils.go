package core

import (
	"log"
	"os"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	googleProto "google.golang.org/protobuf/proto"
)

var DefaultSimTestOptions = &proto.SimOptions{
	Iterations: 20,
	IsTest:     true,
	Debug:      false,
	RandomSeed: 101,
}
var StatWeightsDefaultSimTestOptions = &proto.SimOptions{
	Iterations: 300,
	IsTest:     true,
	Debug:      false,
	RandomSeed: 101,
}
var AverageDefaultSimTestOptions = &proto.SimOptions{
	Iterations: 2000,
	IsTest:     true,
	Debug:      false,
	RandomSeed: 101,
}

const ShortDuration = 60
const LongDuration = 300

// TODO: Update
var DefaultTargetProtoLvl60 = &proto.Target{
	Level: 63,
	Stats: stats.Stats{
		stats.Armor:       1104,
		stats.AttackPower: 320,
	}.ToFloatArray(),
	MobType: proto.MobType_MobTypeDemon,

	SwingSpeed:    2,
	MinBaseDamage: 4192.05,
	ParryHaste:    true,
	DamageSpread:  0.3333,
}

///////////////////////////////////////////////////////////////////////////
//                                 Raid Buffs
///////////////////////////////////////////////////////////////////////////

var FullRaidBuffs = &proto.RaidBuffs{
	ArcaneBrilliance:      true,
	BattleShout:           proto.TristateEffect_TristateEffectImproved,
	BloodPact:             proto.TristateEffect_TristateEffectImproved,
	DevotionAura:          proto.TristateEffect_TristateEffectImproved,
	DivineSpirit:          true,
	FireResistanceAura:    true,
	FireResistanceTotem:   true,
	FrostResistanceAura:   true,
	FrostResistanceTotem:  true,
	GiftOfTheWild:         proto.TristateEffect_TristateEffectImproved,
	GraceOfAirTotem:       proto.TristateEffect_TristateEffectImproved,
	LeaderOfThePack:       true,
	ManaSpringTotem:       proto.TristateEffect_TristateEffectImproved,
	MoonkinAura:           true,
	NatureResistanceTotem: true,
	PowerWordFortitude:    proto.TristateEffect_TristateEffectImproved,
	RetributionAura:       proto.TristateEffect_TristateEffectImproved,
	ShadowProtection:      true,
	StrengthOfEarthTotem:  proto.TristateEffect_TristateEffectImproved,
	Thorns:                proto.TristateEffect_TristateEffectImproved,
	TrueshotAura:          true,
	SanctityAura:          true,
}

///////////////////////////////////////////////////////////////////////////
//                                 Party Buffs
///////////////////////////////////////////////////////////////////////////

var FullPartyBuffs = &proto.PartyBuffs{}

///////////////////////////////////////////////////////////////////////////
//                                 Individual Buffs
///////////////////////////////////////////////////////////////////////////

var FullIndividualBuffs = &proto.IndividualBuffs{
	BlessingOfKings:              true,
	BlessingOfMight:              proto.TristateEffect_TristateEffectImproved,
	BlessingOfSanctuary:          true,
	BlessingOfWisdom:             proto.TristateEffect_TristateEffectImproved,
	FengusFerocity:               true,
	MoldarsMoxie:                 true,
	RallyingCryOfTheDragonslayer: true,
	SaygesFortune:                proto.SaygesFortune_SaygesDamage,
	SlipkiksSavvy:                true,
	SongflowerSerenade:           true,
	SpiritOfZandalar:             true,
	WarchiefsBlessing:            true,
}

///////////////////////////////////////////////////////////////////////////
//                                 Debuffs
///////////////////////////////////////////////////////////////////////////

var FullDebuffs = &proto.Debuffs{
	CurseOfElements:        true,
	CurseOfShadow:          true,
	CurseOfRecklessness:    true,
	CurseOfWeakness:        proto.TristateEffect_TristateEffectImproved,
	DemoralizingRoar:       proto.TristateEffect_TristateEffectImproved,
	DemoralizingShout:      proto.TristateEffect_TristateEffectImproved,
	ExposeArmor:            proto.TristateEffect_TristateEffectImproved,
	FaerieFire:             true,
	ImprovedScorch:         true,
	ImprovedShadowBolt:     true,
	InsectSwarm:            true,
	JudgementOfLight:       true,
	JudgementOfWisdom:      true,
	JudgementOfTheCrusader: proto.TristateEffect_TristateEffectImproved,
	ScorpidSting:           true,
	ShadowWeaving:          true,
	Stormstrike:            true,
	SunderArmor:            true,
	ThunderClap:            proto.TristateEffect_TristateEffectImproved,
	WintersChill:           true,
}

///////////////////////////////////////////////////////////////////////////
//                                 Full Buffs
///////////////////////////////////////////////////////////////////////////

var FullBuffs = BuffsCombo{
	Label: "Phase 5 Buffs",

	Debuffs: FullDebuffs,
	Party:   FullPartyBuffs,
	Player:  FullIndividualBuffs,
	Raid:    FullRaidBuffs,
}

func NewDefaultTarget() *proto.Target {
	return DefaultTargetProtoLvl60
}

func MakeDefaultEncounterCombos() []EncounterCombo {
	var DefaultTarget = NewDefaultTarget()

	multipleTargets := make([]*proto.Target, 20)
	for i := range multipleTargets {
		multipleTargets[i] = DefaultTarget
	}

	return []EncounterCombo{
		{
			Label: "ShortSingleTarget",
			Encounter: &proto.Encounter{
				Duration:             ShortDuration,
				ExecuteProportion_20: 0.2,
				ExecuteProportion_25: 0.25,
				ExecuteProportion_35: 0.35,
				Targets: []*proto.Target{
					DefaultTarget,
				},
			},
		},
		{
			Label: "LongSingleTarget",
			Encounter: &proto.Encounter{
				Duration:             LongDuration,
				ExecuteProportion_20: 0.2,
				ExecuteProportion_25: 0.25,
				ExecuteProportion_35: 0.35,
				Targets: []*proto.Target{
					DefaultTarget,
				},
			},
		},
		{
			Label: "LongMultiTarget",
			Encounter: &proto.Encounter{
				Duration:             LongDuration,
				ExecuteProportion_20: 0.2,
				ExecuteProportion_25: 0.25,
				ExecuteProportion_35: 0.35,
				Targets:              multipleTargets,
			},
		},
	}
}

func MakeSingleTargetEncounter(variation float64) *proto.Encounter {
	return &proto.Encounter{
		Duration:             LongDuration,
		DurationVariation:    variation,
		ExecuteProportion_20: 0.2,
		ExecuteProportion_25: 0.25,
		ExecuteProportion_35: 0.35,
		Targets: []*proto.Target{
			NewDefaultTarget(),
		},
	}
}

func CharacterStatsTest(label string, t *testing.T, raid *proto.Raid, expectedStats stats.Stats) {
	csr := &proto.ComputeStatsRequest{
		Raid: raid,
	}

	result := ComputeStats(csr)
	finalStats := stats.FromFloatArray(result.RaidStats.Parties[0].Players[0].FinalStats.Stats)

	const tolerance = 0.5
	if !finalStats.EqualsWithTolerance(expectedStats, tolerance) {
		t.Fatalf("%s failed: CharacterStats() = %v, expected %v", label, finalStats, expectedStats)
	}
}

func StatWeightsTest(label string, t *testing.T, _swr *proto.StatWeightsRequest, expectedStatWeights stats.Stats) {
	// Make a copy so we can safely change fields.
	swr := googleProto.Clone(_swr).(*proto.StatWeightsRequest)

	swr.Encounter.Duration = LongDuration
	swr.SimOptions.Iterations = 5000

	result := StatWeights(swr)
	resultWeights := stats.FromFloatArray(result.Dps.Weights.Stats)

	const tolerance = 0.05
	if !resultWeights.EqualsWithTolerance(expectedStatWeights, tolerance) {
		t.Fatalf("%s failed: CalcStatWeight() = %v, expected %v", label, resultWeights, expectedStatWeights)
	}
}

func RaidSimTest(label string, t *testing.T, rsr *proto.RaidSimRequest, expectedDps float64) {
	result := RunRaidSim(rsr)
	if result.Error != nil {
		t.Fatalf("Sim failed with error: %s", result.Error.Message)
	}
	tolerance := 0.5
	if result.RaidMetrics.Dps.Avg < expectedDps-tolerance || result.RaidMetrics.Dps.Avg > expectedDps+tolerance {
		// Automatically print output if we had debugging enabled.
		if rsr.SimOptions.Debug {
			log.Printf("LOGS:\n%s\n", result.Logs)
		}
		t.Fatalf("%s failed: expected %0f dps from sim but was %0f", label, expectedDps, result.RaidMetrics.Dps.Avg)
	}
}

func RaidBenchmark(b *testing.B, rsr *proto.RaidSimRequest) {
	rsr.Encounter.Duration = LongDuration
	rsr.SimOptions.Iterations = 1

	// Set to false because IsTest adds a lot of computation.
	rsr.SimOptions.IsTest = false

	for i := 0; i < b.N; i++ {
		result := RunRaidSim(rsr)
		if result.Error != nil {
			b.Fatalf("RaidBenchmark() at iteration %d failed: %v", i, result.Error.Message)
		}
	}
}

func GetAplRotation(dir string, file string) RotationCombo {
	filePath := dir + "/" + file + ".apl.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to load apl json file: %s, %s", filePath, err)
	}

	return RotationCombo{Label: file, Rotation: APLRotationFromJsonString(string(data))}
}

func GetGearSet(dir string, file string) GearSetCombo {
	filePath := dir + "/" + file + ".gear.json"
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to load gear json file: %s, %s", filePath, err)
	}

	return GearSetCombo{Label: file, GearSet: EquipmentSpecFromJsonString(string(data))}
}

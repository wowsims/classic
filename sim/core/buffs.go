package core

import (
	"math"
	"time"

	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
	googleProto "google.golang.org/protobuf/proto"
)

type BuffName int32

const (
	// General Buffs
	ArcaneIntellect BuffName = iota
	BattleShout
	BlessingOfMight
	BlessingOfWisdom
	HornOfLordaeron
	BloodPact
	CommandingShout
	DevotionAura
	DivineSpirit
	GraceOfAir
	ManaSpring
	MarkOfTheWild
	PowerWordFortitude
	StrengthOfEarth
	Windfury
	SanctityAura
	BattleSquawk

	// Resistance
	AspectOfTheWild
	FireResistanceAura
	FireResistanceTotem
	FrostResistanceTotem
	FrostResistanceAura
	NatureResistanceTotem
	ShadowProtection
	ShadowResistanceAura

	// Scrolls
	ScrollOfAgility
	ScrollOfIntellect
	ScrollOfSpirit
	ScrollOfStrength
	ScrollOfStamina
	ScrollOfProtection
)

// Stats from buffs pre-tristate buffs
var BuffSpellValues = map[BuffName]stats.Stats{
	ArcaneIntellect: {
		stats.Intellect: 31,
	},
	DivineSpirit: {
		stats.Spirit: 40,
	},
	AspectOfTheWild: {
		stats.NatureResistance: 60,
	},
	BattleShout: {
		stats.AttackPower: TernaryFloat64(IncludeAQ, 232, 193),
	},
	BlessingOfMight: {
		stats.AttackPower: TernaryFloat64(IncludeAQ, 185, 155),
	},
	BlessingOfWisdom: {
		stats.MP5: TernaryFloat64(IncludeAQ, 33, 30),
	},
	HornOfLordaeron: {
		stats.Strength: TernaryFloat64(IncludeAQ, 89, 70.15),
		stats.Agility:  TernaryFloat64(IncludeAQ, 89, 70.15),
	},
	BloodPact: {
		stats.Stamina: 42,
	},
	CommandingShout: {
		stats.Stamina: 42,
	},
	DevotionAura: {
		stats.BonusArmor: 735,
	},
	GraceOfAir: {
		stats.Agility: TernaryFloat64(IncludeAQ, 77, 67),
	},
	FireResistanceAura: {
		stats.FireResistance: 60,
	},
	FireResistanceTotem: {
		stats.FireResistance: 60,
	},
	FrostResistanceAura: {
		stats.FrostResistance: 60,
	},
	FrostResistanceTotem: {
		stats.FrostResistance: 60,
	},
	ManaSpring: {
		stats.MP5: 25,
	},
	MarkOfTheWild: {
		stats.BonusArmor:       285,
		stats.Stamina:          12,
		stats.Agility:          12,
		stats.Strength:         12,
		stats.Intellect:        12,
		stats.Spirit:           12,
		stats.ArcaneResistance: 20,
		stats.ShadowResistance: 20,
		stats.NatureResistance: 20,
		stats.FireResistance:   20,
		stats.FrostResistance:  20,
	},
	NatureResistanceTotem: {
		stats.NatureResistance: 60,
	},
	PowerWordFortitude: {
		stats.Stamina: 54,
	},
	ShadowProtection: {
		stats.ShadowResistance: 60,
	},
	ShadowResistanceAura: {
		stats.ShadowResistance: 60,
	},
	StrengthOfEarth: {
		stats.Strength: TernaryFloat64(IncludeAQ, 77, 61),
	},
	ScrollOfAgility: {
		stats.Agility: 17,
	},
	ScrollOfIntellect: {
		stats.Intellect: 16,
	},
	ScrollOfSpirit: {
		stats.Spirit: 15,
	},
	ScrollOfStamina: {
		stats.Stamina: 16,
	},
	ScrollOfStrength: {
		stats.Strength: 17,
	},
	ScrollOfProtection: {
		stats.BonusArmor: 240,
	},
}

type ExtraOnGain func(aura *Aura, sim *Simulation)
type ExtraOnExpire func(aura *Aura, sim *Simulation)

type BuffConfig struct {
	Category string
	Stats    []StatConfig
	// Hacky way to allow Pseudostat mods
	ExtraOnGain   ExtraOnGain
	ExtraOnExpire ExtraOnExpire
}

type StatConfig struct {
	Stat             stats.Stat
	Amount           float64
	IsMultiplicative bool
}

// Create an exclusive effect that tries to determine within-category priority based on the value of stats provided.
func makeExclusiveBuff(aura *Aura, config BuffConfig) {
	aura.BuildPhase = CharacterBuildPhaseBuffs

	startingStats := aura.Unit.GetStats()
	bonusStats := stats.Stats{}
	statDeps := []*stats.StatDependency{}
	for _, statConfig := range config.Stats {
		if statConfig.IsMultiplicative {
			startingStats[statConfig.Stat] *= statConfig.Amount
			statDeps = append(statDeps, aura.Unit.NewDynamicMultiplyStat(statConfig.Stat, statConfig.Amount))
		} else {
			startingStats[statConfig.Stat] += statConfig.Amount
			bonusStats[statConfig.Stat] += statConfig.Amount
		}
	}

	totalStats := 0.0
	for _, amount := range startingStats {
		totalStats += amount
	}

	aura.NewExclusiveEffect(config.Category, false, ExclusiveEffect{
		Priority: totalStats,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			aura.Unit.AddBuildPhaseStatsDynamic(sim, bonusStats)

			for _, dep := range statDeps {
				ee.Aura.Unit.EnableBuildPhaseStatDep(sim, dep)
			}

			if config.ExtraOnGain != nil {
				config.ExtraOnGain(ee.Aura, sim)
			}
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			aura.Unit.AddBuildPhaseStatsDynamic(sim, bonusStats.Multiply(-1))

			for _, dep := range statDeps {
				ee.Aura.Unit.DisableBuildPhaseStatDep(sim, dep)
			}

			if config.ExtraOnExpire != nil {
				config.ExtraOnExpire(ee.Aura, sim)
			}
		},
	})
}

// Applies buffs that affect individual players.
func applyBuffEffects(agent Agent, playerFaction proto.Faction, raidBuffs *proto.RaidBuffs, partyBuffs *proto.PartyBuffs, individualBuffs *proto.IndividualBuffs) {
	character := agent.GetCharacter()
	isAlliance := playerFaction == proto.Faction_Alliance
	isHorde := playerFaction == proto.Faction_Horde
	bonusResist := float64(0)

	if raidBuffs.ArcaneBrilliance {
		character.AddStats(BuffSpellValues[ArcaneIntellect])
	} else if raidBuffs.ScrollOfIntellect {
		character.AddStats(BuffSpellValues[ScrollOfIntellect])
	}

	if raidBuffs.GiftOfTheWild > 0 {
		updateStats := BuffSpellValues[MarkOfTheWild]
		if raidBuffs.GiftOfTheWild == proto.TristateEffect_TristateEffectImproved {
			updateStats = updateStats.Multiply(1.35).Floor()
		}
		character.AddStats(updateStats)
		bonusResist = updateStats[stats.NatureResistance]
	}

	if raidBuffs.NatureResistanceTotem {
		updateStats := BuffSpellValues[NatureResistanceTotem]
		updateStats[stats.NatureResistance] = updateStats[stats.NatureResistance] - bonusResist
		character.AddStats(updateStats)
	} else if raidBuffs.AspectOfTheWild {
		updateStats := BuffSpellValues[AspectOfTheWild]
		updateStats[stats.NatureResistance] = updateStats[stats.NatureResistance] - bonusResist
		character.AddStats(updateStats)
	}

	if raidBuffs.FireResistanceAura {
		updateStats := BuffSpellValues[FireResistanceAura]
		updateStats[stats.FireResistance] = updateStats[stats.FireResistance] - bonusResist
		character.AddStats(updateStats)
	} else if raidBuffs.FireResistanceTotem {
		updateStats := BuffSpellValues[FireResistanceTotem]
		updateStats[stats.FireResistance] = updateStats[stats.FireResistance] - bonusResist
		character.AddStats(updateStats)
	}

	if raidBuffs.FrostResistanceAura {
		updateStats := BuffSpellValues[FrostResistanceAura]
		updateStats[stats.FrostResistance] = updateStats[stats.FrostResistance] - bonusResist
		character.AddStats(updateStats)
	} else if raidBuffs.FrostResistanceTotem {
		updateStats := BuffSpellValues[FrostResistanceTotem]
		updateStats[stats.FrostResistance] = updateStats[stats.FrostResistance] - bonusResist
		character.AddStats(updateStats)
	}

	if raidBuffs.Thorns != proto.TristateEffect_TristateEffectMissing {
		ThornsAura(character, GetTristateValueInt32(raidBuffs.Thorns, 0, 3))
	}

	if raidBuffs.MoonkinAura {
		character.AddStat(stats.SpellCrit, 3*SpellCritRatingPerCritChance)
	}

	if raidBuffs.LeaderOfThePack {
		character.AddStats(stats.Stats{
			stats.MeleeCrit: 3 * CritRatingPerCritChance,
		})
	}

	if raidBuffs.TrueshotAura {
		TrueshotAura(&character.Unit)
	}

	if raidBuffs.PowerWordFortitude > 0 {
		updateStats := BuffSpellValues[PowerWordFortitude]
		if raidBuffs.PowerWordFortitude == proto.TristateEffect_TristateEffectImproved {
			updateStats = updateStats.Multiply(1.3).Floor()
		}
		character.AddStats(updateStats)
	} else if raidBuffs.ScrollOfStamina {
		character.AddStats(BuffSpellValues[ScrollOfStamina])
	}

	if raidBuffs.BloodPact > 0 {
		updateStats := BuffSpellValues[BloodPact]
		if raidBuffs.BloodPact == proto.TristateEffect_TristateEffectImproved {
			updateStats = updateStats.Multiply(1.3).Floor()
		}
		character.AddStats(updateStats)
	}

	if raidBuffs.ShadowResistanceAura {
		updateStats := BuffSpellValues[ShadowResistanceAura]
		updateStats[stats.ShadowResistance] = updateStats[stats.ShadowResistance] - bonusResist
		character.AddStats(updateStats)
	} else if raidBuffs.ShadowProtection {
		updateStats := BuffSpellValues[ShadowProtection]
		updateStats[stats.ShadowResistance] = updateStats[stats.ShadowResistance] - bonusResist
		character.AddStats(updateStats)
	}

	if raidBuffs.DivineSpirit {
		character.AddStats(BuffSpellValues[DivineSpirit])
	} else if raidBuffs.ScrollOfSpirit {
		character.AddStats(BuffSpellValues[ScrollOfSpirit])
	}

	if individualBuffs.BlessingOfKings && isAlliance {
		MakePermanent(BlessingOfKingsAura(character))
	}

	if raidBuffs.SanctityAura && isAlliance {
		MakePermanent(SanctityAuraAura(character))
	}

	// TODO: Classic
	/*	if individualBuffs.BlessingOfSanctuary {
			character.PseudoStats.DamageTakenMultiplier *= 0.97
			BlessingOfSanctuaryAura(character)
		}
	*/

	if raidBuffs.DevotionAura != proto.TristateEffect_TristateEffectMissing && isAlliance {
		MakePermanent(DevotionAuraAura(&character.Unit, GetTristateValueInt32(raidBuffs.DevotionAura, 0, 2)))
	}

	if raidBuffs.StoneskinTotem != proto.TristateEffect_TristateEffectMissing && isHorde {
		MakePermanent(StoneskinTotemAura(&character.Unit, GetTristateValueInt32(raidBuffs.StoneskinTotem, 0, 2)))
	}

	if raidBuffs.RetributionAura != proto.TristateEffect_TristateEffectMissing && isAlliance {
		RetributionAura(character, GetTristateValueInt32(raidBuffs.RetributionAura, 0, 2))
	}

	if raidBuffs.BattleShout != proto.TristateEffect_TristateEffectMissing {
		MakePermanent(BattleShoutAura(&character.Unit, GetTristateValueInt32(raidBuffs.BattleShout, 0, 5), 0, false)) // Do we implement 3pc wrath for the other sims?
	}

	if individualBuffs.BlessingOfMight != proto.TristateEffect_TristateEffectMissing && isAlliance {
		MakePermanent(BlessingOfMightAura(&character.Unit, GetTristateValueInt32(individualBuffs.BlessingOfMight, 0, 5)))
	}

	if raidBuffs.StrengthOfEarthTotem != proto.TristateEffect_TristateEffectMissing && isHorde {
		multiplier := GetTristateValueFloat(raidBuffs.StrengthOfEarthTotem, 1, 1.15)
		MakePermanent(StrengthOfEarthTotemAura(&character.Unit, multiplier))
	}

	if raidBuffs.GraceOfAirTotem > 0 && isHorde {
		multiplier := GetTristateValueFloat(raidBuffs.GraceOfAirTotem, 1, 1.15)
		MakePermanent(GraceOfAirTotemAura(&character.Unit, multiplier))
	}

	if individualBuffs.BlessingOfWisdom > 0 && isAlliance {
		updateStats := BuffSpellValues[BlessingOfWisdom]
		if individualBuffs.BlessingOfWisdom == proto.TristateEffect_TristateEffectImproved {
			updateStats = updateStats.Multiply(1.2)
		}
		character.AddStats(updateStats)
	} else if raidBuffs.ManaSpringTotem > 0 && isHorde {
		updateStats := BuffSpellValues[ManaSpring]
		if raidBuffs.ManaSpringTotem == proto.TristateEffect_TristateEffectImproved {
			updateStats = updateStats.Multiply(1.25)
		}
		character.AddStats(updateStats)
	}

	if raidBuffs.BattleSquawk > 0 {
		numBattleSquawks := raidBuffs.BattleSquawk
		BattleSquawkAura(&character.Unit, numBattleSquawks)
	}

	// World Buffs
	ApplyDragonslayerBuffs(&character.Unit, individualBuffs)

	if individualBuffs.SpiritOfZandalar {
		ApplySpiritOfZandalar(&character.Unit)
	}

	if individualBuffs.SongflowerSerenade {
		ApplySongflowerSerenade(&character.Unit)
	}

	ApplyWarchiefsBuffs(&character.Unit, individualBuffs, isAlliance, isHorde)

	// Dire Maul Buffs
	if individualBuffs.FengusFerocity {
		ApplyFengusFerocity(&character.Unit)
	}

	if individualBuffs.MoldarsMoxie {
		ApplyMoldarsMoxie(&character.Unit)
	}

	if individualBuffs.SlipkiksSavvy {
		ApplySlipkiksSavvy(&character.Unit)
	}

	// Darkmoon Faire Buffs
	if individualBuffs.SaygesFortune != proto.SaygesFortune_SaygesUnknown {
		ApplySaygesFortunes(character, individualBuffs.SaygesFortune)
	}

	// TODO: Classic provide in APL?
	registerPowerInfusionCD(agent, individualBuffs.PowerInfusions)
	registerManaTideTotemCD(agent, partyBuffs.ManaTideTotems)
	registerInnervateCD(agent, individualBuffs.Innervates)

	character.AddStats(stats.Stats{
		stats.SpellCrit: 2 * SpellCritRatingPerCritChance * float64(partyBuffs.AtieshMage),
	})
	character.AddStats(stats.Stats{
		stats.SpellPower: 33 * float64(partyBuffs.AtieshWarlock),
	})
}

// Applies buffs to pets.
func applyPetBuffEffects(petAgent PetAgent, playerFaction proto.Faction, raidBuffs *proto.RaidBuffs, partyBuffs *proto.PartyBuffs, individualBuffs *proto.IndividualBuffs) {
	// Summoned pets, like Mage Water Elemental, aren't around to receive raid buffs.
	// Also assume that applicable world buffs are applied to the starting pet only
	if petAgent.GetPet().IsGuardian() || !petAgent.GetPet().enabledOnStart {
		return
	}

	raidBuffs = googleProto.Clone(raidBuffs).(*proto.RaidBuffs)
	partyBuffs = googleProto.Clone(partyBuffs).(*proto.PartyBuffs)
	individualBuffs = googleProto.Clone(individualBuffs).(*proto.IndividualBuffs)

	// We need to modify the buffs a bit because some things are applied to pets by
	// the owner during combat or don't make sense for a pet.
	individualBuffs.Innervates = 0
	individualBuffs.PowerInfusions = 0

	// Pets only receive Onyxia, Rend, and ZG buffs because they're globally applied in their respective zones
	// SoD versions were removed from pets though
	individualBuffs.FengusFerocity = false
	individualBuffs.MoldarsMoxie = false
	individualBuffs.SaygesFortune = proto.SaygesFortune_SaygesUnknown
	individualBuffs.SongflowerSerenade = false
	individualBuffs.SlipkiksSavvy = false

	applyBuffEffects(petAgent, playerFaction, raidBuffs, partyBuffs, individualBuffs)
}

func SanctityAuraAura(character *Character) *Aura {
	return character.GetOrRegisterAura(Aura{
		Label:    "Sanctity Aura",
		ActionID: ActionID{SpellID: 20218},
		Duration: NeverExpires,
		OnReset: func(aura *Aura, sim *Simulation) {
			aura.Activate(sim)
		},
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexHoly] *= 1.1
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexHoly] /= 1.1
		},
	})
}

func BlessingOfKingsAura(character *Character) *Aura {
	statDeps := []*stats.StatDependency{
		character.NewDynamicMultiplyStat(stats.Stamina, 1.10),
		character.NewDynamicMultiplyStat(stats.Agility, 1.10),
		character.NewDynamicMultiplyStat(stats.Strength, 1.10),
		character.NewDynamicMultiplyStat(stats.Intellect, 1.10),
		character.NewDynamicMultiplyStat(stats.Spirit, 1.10),
	}

	return MakePermanent(character.RegisterAura(Aura{
		Label:      "Blessing of Kings",
		ActionID:   ActionID{SpellID: 20217},
		BuildPhase: CharacterBuildPhaseBuffs,
		OnGain: func(aura *Aura, sim *Simulation) {
			if aura.Unit.Env.MeasuringStats && aura.Unit.Env.State != Finalized {
				for _, dep := range statDeps {
					aura.Unit.StatDependencyManager.EnableDynamicStatDep(dep)
				}
			} else {
				for _, dep := range statDeps {
					aura.Unit.EnableDynamicStatDep(sim, dep)
				}
			}
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			if aura.Unit.Env.MeasuringStats && aura.Unit.Env.State != Finalized {
				for _, dep := range statDeps {
					aura.Unit.StatDependencyManager.DisableDynamicStatDep(dep)
				}
			} else {
				for _, dep := range statDeps {
					aura.Unit.DisableDynamicStatDep(sim, dep)
				}
			}
		},
	}))
}

// TODO: Classic
func InspirationAura(unit *Unit, points int32) *Aura {
	multiplier := 1 - []float64{0, .03, .07, .10}[points]

	return unit.GetOrRegisterAura(Aura{
		Label:    "Inspiration",
		ActionID: ActionID{SpellID: 15363},
		Duration: time.Second * 15,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexPhysical] *= multiplier
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.SchoolDamageTakenMultiplier[stats.SchoolIndexPhysical] /= multiplier
		},
	})
}

func ApplyInspiration(character *Character, uptime float64) {
	if uptime <= 0 {
		return
	}
	uptime = min(1, uptime)

	inspirationAura := InspirationAura(&character.Unit, 3)

	ApplyFixedUptimeAura(inspirationAura, uptime, time.Millisecond*2500, 1)
}

func DevotionAuraAura(unit *Unit, points int32) *Aura {
	updateStats := BuffSpellValues[DevotionAura]
	updateStats = updateStats.Multiply(1 + .125*float64(points))

	return unit.RegisterAura(Aura{
		Label:    "Devotion Aura",
		ActionID: ActionID{SpellID: 10293},
		Duration: NeverExpires,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.AddStatsDynamic(sim, updateStats)
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.AddStatsDynamic(sim, updateStats.Multiply(-1))
		},
	})
}

func StoneskinTotemAura(unit *Unit, points int32) *Aura {
	meleeDamageReduction := -30.0
	meleeDamageReduction *= 1 + .1*float64(points)
	meleeDamageReduction = math.Floor(meleeDamageReduction)

	return unit.GetOrRegisterAura(Aura{
		Label:    "Stoneskin",
		ActionID: ActionID{SpellID: 10408},
		Duration: NeverExpires,
		OnReset: func(aura *Aura, sim *Simulation) {
			aura.Activate(sim)
		},
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusDamageTakenAfterModifiers[DefenseTypeMelee] += meleeDamageReduction
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.BonusDamageTakenAfterModifiers[DefenseTypeMelee] += meleeDamageReduction
		},
	})
}

func RetributionAura(character *Character, points int32) *Aura {
	baseDamage := 20.0

	actionID := ActionID{SpellID: 10301}

	damage := float64(baseDamage) * (1 + 0.25*float64(points))

	procSpell := character.RegisterSpell(SpellConfig{
		ActionID:    actionID,
		SpellSchool: SpellSchoolHoly,
		ProcMask:    ProcMaskEmpty,
		Flags:       SpellFlagBinary | SpellFlagNoOnCastComplete | SpellFlagPassiveSpell,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMagicHit)
		},
	})

	return character.RegisterAura(Aura{
		Label:    "Retribution Aura",
		ActionID: actionID,
		Duration: NeverExpires,
		OnReset: func(aura *Aura, sim *Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
			if result.Landed() && spell.ProcMask.Matches(ProcMaskMelee) {
				procSpell.Cast(sim, spell.Unit)
			}
		},
	})
}

func ThornsAura(character *Character, points int32) *Aura {
	baseDamage := 18.0

	actionID := ActionID{SpellID: 9910}
	damage := float64(baseDamage) * (1 + 0.25*float64(points))

	procSpell := character.RegisterSpell(SpellConfig{
		ActionID:    actionID,
		SpellSchool: SpellSchoolNature,
		ProcMask:    ProcMaskEmpty,
		Flags:       SpellFlagBinary | SpellFlagNoOnCastComplete | SpellFlagPassiveSpell,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *Simulation, target *Unit, spell *Spell) {
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMagicHit)
		},
	})

	return MakePermanent(character.RegisterAura(Aura{
		Label:    "Thorns",
		ActionID: actionID,
		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
			if result.Landed() && spell.ProcMask.Matches(ProcMaskMelee) {
				procSpell.Cast(sim, spell.Unit)
			}
		},
	}))
}

// func BlessingOfSanctuaryAura(character *Character) {
// 	physReduction := 24.0
// 	blockDamage := 35
// 	actionID := ActionID{SpellID: 20914}
// 	character.RegisterAura(Aura{
// 		Label:    "Blessing of Sanctuary",
// 		ActionID: actionID,
// 		Duration: NeverExpires,
// 		OnReset: func(aura *Aura, sim *Simulation) {
// 			aura.Activate(sim)
// 		},
// 		OnSpellHitTaken: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
// 			if result.Outcome.Matches(OutcomeBlock | OutcomeDodge | OutcomeParry) {
// 			}
// 		},
// 	})
// }

// Used for approximating cooldowns applied by other players to you, such as
// bloodlust, innervate, power infusion, etc. This is specifically for buffs
// which can be consecutively applied multiple times to a single player.
type externalConsecutiveCDApproximation struct {
	ActionID         ActionID
	AuraTag          string
	CooldownPriority int32
	Type             CooldownType
	AuraDuration     time.Duration
	AuraCD           time.Duration

	// Callback for extra activation conditions.
	ShouldActivate CooldownActivationCondition

	// Applies the buff.
	AddAura CooldownActivation
}

// numSources is the number of other players assigned to apply the buff to this player.
// E.g. the number of other shaman in the group using bloodlust.
func registerExternalConsecutiveCDApproximation(agent Agent, config externalConsecutiveCDApproximation, numSources int32) {
	if numSources == 0 {
		panic("Need at least 1 source!")
	}
	character := agent.GetCharacter()

	var nextExternalIndex int

	externalTimers := make([]*Timer, numSources)
	for i := 0; i < int(numSources); i++ {
		externalTimers[i] = character.NewTimer()
	}
	sharedTimer := character.NewTimer()

	spell := character.RegisterSpell(SpellConfig{
		ActionID: config.ActionID,
		Flags:    SpellFlagNoOnCastComplete | SpellFlagNoMetrics | SpellFlagNoLogs,

		Cast: CastConfig{
			CD: Cooldown{
				Timer:    sharedTimer,
				Duration: config.AuraDuration, // Assumes that multiple buffs are different sources.
			},
		},
		ExtraCastCondition: func(sim *Simulation, target *Unit) bool {
			if !externalTimers[nextExternalIndex].IsReady(sim) {
				return false
			}

			if character.HasActiveAuraWithTag(config.AuraTag) {
				return false
			}

			return true
		},

		ApplyEffects: func(sim *Simulation, _ *Unit, _ *Spell) {
			config.AddAura(sim, character)
			externalTimers[nextExternalIndex].Set(sim.CurrentTime + config.AuraCD)

			nextExternalIndex = (nextExternalIndex + 1) % len(externalTimers)

			if externalTimers[nextExternalIndex].IsReady(sim) {
				sharedTimer.Set(sim.CurrentTime + config.AuraDuration)
			} else {
				sharedTimer.Set(sim.CurrentTime + externalTimers[nextExternalIndex].TimeToReady(sim))
			}
		},
	})

	character.AddMajorCooldown(MajorCooldown{
		Spell:    spell,
		Priority: config.CooldownPriority,
		Type:     config.Type,

		ShouldActivate: config.ShouldActivate,
	})
}

// var BloodlustActionID = ActionID{SpellID: 2825}

// const SatedAuraLabel = "Sated"
// const BloodlustAuraTag = "Bloodlust"
// const BloodlustDuration = time.Second * 40
// const BloodlustCD = time.Minute * 10

// func registerBloodlustCD(agent Agent) {
// 	character := agent.GetCharacter()
// 	bloodlustAura := BloodlustAura(character, -1)

// 	spell := character.RegisterSpell(SpellConfig{
// 		ActionID: bloodlustAura.ActionID,
// 		Flags:    SpellFlagNoOnCastComplete | SpellFlagNoMetrics | SpellFlagNoLogs,

// 		Cast: CastConfig{
// 			CD: Cooldown{
// 				Timer:    character.NewTimer(),
// 				Duration: BloodlustCD,
// 			},
// 		},

// 		ApplyEffects: func(sim *Simulation, target *Unit, _ *Spell) {
// 			if !target.HasActiveAura(SatedAuraLabel) {
// 				bloodlustAura.Activate(sim)
// 			}
// 		},
// 	})

// 	character.AddMajorCooldown(MajorCooldown{
// 		Spell:    spell,
// 		Priority: CooldownPriorityBloodlust,
// 		Type:     CooldownTypeDPS,
// 		ShouldActivate: func(sim *Simulation, character *Character) bool {
// 			// Haste portion doesn't stack with Power Infusion, so prefer to wait.
// 			return !character.HasActiveAuraWithTag(PowerInfusionAuraTag) && !character.HasActiveAura(SatedAuraLabel)
// 		},
// 	})
// }

// func BloodlustAura(character *Character, actionTag int32) *Aura {
// 	actionID := BloodlustActionID.WithTag(actionTag)

// 	sated := character.GetOrRegisterAura(Aura{
// 		Label:    SatedAuraLabel,
// 		ActionID: ActionID{SpellID: 57724},
// 		Duration: time.Minute * 10,
// 	})

// 	aura := character.GetOrRegisterAura(Aura{
// 		Label:    "Bloodlust-" + actionID.String(),
// 		Tag:      BloodlustAuraTag,
// 		ActionID: actionID,
// 		Duration: BloodlustDuration,
// 		OnGain: func(aura *Aura, sim *Simulation) {
// 			character.MultiplyAttackSpeed(sim, 1.3)
// 			for _, pet := range character.Pets {
// 				if pet.IsEnabled() && !pet.IsGuardian() {
// 					BloodlustAura(&pet.Character, actionTag).Activate(sim)
// 				}
// 			}

// 			if character.HasActiveAura(SatedAuraLabel) {
// 				aura.Deactivate(sim) // immediately remove it person already has sated.
// 				return
// 			}
// 		},
// 		OnExpire: func(aura *Aura, sim *Simulation) {
// 			character.MultiplyAttackSpeed(sim, 1.0/1.3)
// 			sated.Activate(sim)
// 		},
// 	})
// 	multiplyCastSpeedEffect(aura, 1.3)
// 	return aura
// }

var PowerInfusionActionID = ActionID{SpellID: 10060}
var PowerInfusionAuraTag = "PowerInfusion"

const PowerInfusionDuration = time.Second * 15
const PowerInfusionCD = time.Minute * 3

func registerPowerInfusionCD(agent Agent, numPowerInfusions int32) {
	if numPowerInfusions == 0 {
		return
	}

	piAura := PowerInfusionAura(&agent.GetCharacter().Unit, -1)

	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			ActionID:         PowerInfusionActionID.WithTag(-1),
			AuraTag:          PowerInfusionAuraTag,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     PowerInfusionDuration,
			AuraCD:           PowerInfusionCD,
			Type:             CooldownTypeDPS,

			AddAura: func(sim *Simulation, character *Character) { piAura.Activate(sim) },
		},
		numPowerInfusions)
}

func PowerInfusionAura(character *Unit, actionTag int32) *Aura {
	actionID := ActionID{SpellID: 10060, Tag: actionTag}
	aura := character.GetOrRegisterAura(Aura{
		Label:    "PowerInfusion-" + actionID.String(),
		Tag:      PowerInfusionAuraTag,
		ActionID: actionID,
		Duration: PowerInfusionDuration,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.SchoolDamageDealtMultiplier.MultiplyMagicSchools(1.2)

		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.SchoolDamageDealtMultiplier.MultiplyMagicSchools(1 / 1.2)
		},
	})
	return aura
}

// func multiplyCastSpeedEffect(aura *Aura, multiplier float64) *ExclusiveEffect {
// 	return aura.NewExclusiveEffect("MultiplyCastSpeed", false, ExclusiveEffect{
// 		Priority: multiplier,
// 		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
// 			ee.Aura.Unit.MultiplyCastSpeed(multiplier)
// 		},
// 		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
// 			ee.Aura.Unit.MultiplyCastSpeed(1 / multiplier)
// 		},
// 	})
// }

// var TricksOfTheTradeAuraTag = "TricksOfTheTrade"

// const TricksOfTheTradeCD = time.Second * 3600 // CD is 30s from the time buff ends (so 40s with glyph) but that's in order to be able to set the number of TotT you'll have during the fight

// func registerTricksOfTheTradeCD(agent Agent, numTricksOfTheTrades int32) {
// 	if numTricksOfTheTrades == 0 {
// 		return
// 	}

// 	TotTAura := TricksOfTheTradeAura(&agent.GetCharacter().Unit, -1)

// 	registerExternalConsecutiveCDApproximation(
// 		agent,
// 		externalConsecutiveCDApproximation{
// 			ActionID:         ActionID{SpellID: 57933, Tag: -1},
// 			AuraTag:          TricksOfTheTradeAuraTag,
// 			CooldownPriority: CooldownPriorityDefault,
// 			AuraDuration:     TotTAura.Duration,
// 			AuraCD:           TricksOfTheTradeCD,
// 			Type:             CooldownTypeDPS,

// 			ShouldActivate: func(sim *Simulation, character *Character) bool {
// 				return !agent.GetCharacter().GetExclusiveEffectCategory("PercentDamageModifier").AnyActive()
// 			},
// 			AddAura: func(sim *Simulation, character *Character) { TotTAura.Activate(sim) },
// 		},
// 		numTricksOfTheTrades)
// }

// func TricksOfTheTradeAura(character *Unit, actionTag int32) *Aura {
// 	actionID := ActionID{SpellID: 57933, Tag: actionTag}

// 	aura := character.GetOrRegisterAura(Aura{
// 		Label:    "TricksOfTheTrade-" + actionID.String(),
// 		Tag:      TricksOfTheTradeAuraTag,
// 		ActionID: actionID,
// 		Duration: time.Second * 6,
// 		OnGain: func(aura *Aura, sim *Simulation) {
// 			character.PseudoStats.DamageDealtMultiplier *= 1.15
// 		},
// 		OnExpire: func(aura *Aura, sim *Simulation) {
// 			character.PseudoStats.DamageDealtMultiplier /= 1.15
// 		},
// 	})

// 	RegisterPercentDamageModifierEffect(aura, 1.15)
// 	return aura
// }

func RegisterPercentDamageModifierEffect(aura *Aura, percentDamageModifier float64) *ExclusiveEffect {
	return aura.NewExclusiveEffect("PercentDamageModifier", true, ExclusiveEffect{
		Priority: percentDamageModifier,
	})
}

// var DivineGuardianAuraTag = "DivineGuardian"

// const DivineGuardianDuration = time.Second * 6
// const DivineGuardianCD = time.Minute * 2

// var HandOfSacrificeAuraTag = "HandOfSacrifice"

// const HandOfSacrificeDuration = time.Millisecond * 10500 // subtract Divine Shield GCD
// const HandOfSacrificeCD = time.Minute * 5                // use Divine Shield CD here

// func registerHandOfSacrificeCD(agent Agent, numSacs int32) {
// 	if numSacs == 0 {
// 		return
// 	}

// 	hosAura := HandOfSacrificeAura(agent.GetCharacter(), -1)

// 	registerExternalConsecutiveCDApproximation(
// 		agent,
// 		externalConsecutiveCDApproximation{
// 			ActionID:         ActionID{SpellID: 6940, Tag: -1},
// 			AuraTag:          HandOfSacrificeAuraTag,
// 			CooldownPriority: CooldownPriorityLow,
// 			AuraDuration:     HandOfSacrificeDuration,
// 			AuraCD:           HandOfSacrificeCD,
// 			Type:             CooldownTypeSurvival,

// 			ShouldActivate: func(sim *Simulation, character *Character) bool {
// 				return true
// 			},
// 			AddAura: func(sim *Simulation, character *Character) {
// 				hosAura.Activate(sim)
// 			},
// 		},
// 		numSacs)
// }

// func HandOfSacrificeAura(character *Character, actionTag int32) *Aura {
// 	actionID := ActionID{SpellID: 6940, Tag: actionTag}

// 	return character.GetOrRegisterAura(Aura{
// 		Label:    "HandOfSacrifice-" + actionID.String(),
// 		Tag:      HandOfSacrificeAuraTag,
// 		ActionID: actionID,
// 		Duration: HandOfSacrificeDuration,
// 		OnGain: func(aura *Aura, sim *Simulation) {
// 			character.PseudoStats.DamageTakenMultiplier *= 0.7
// 		},
// 		OnExpire: func(aura *Aura, sim *Simulation) {
// 			character.PseudoStats.DamageTakenMultiplier /= 0.7
// 		},
// 	})
// }

// var PainSuppressionAuraTag = "PainSuppression"

// const PainSuppressionDuration = time.Second * 8
// const PainSuppressionCD = time.Minute * 3

// func registerPainSuppressionCD(agent Agent, numPainSuppressions int32) {
// 	if numPainSuppressions == 0 {
// 		return
// 	}

// 	psAura := PainSuppressionAura(agent.GetCharacter(), -1)

// 	registerExternalConsecutiveCDApproximation(
// 		agent,
// 		externalConsecutiveCDApproximation{
// 			ActionID:         ActionID{SpellID: 33206, Tag: -1},
// 			AuraTag:          PainSuppressionAuraTag,
// 			CooldownPriority: CooldownPriorityDefault,
// 			AuraDuration:     PainSuppressionDuration,
// 			AuraCD:           PainSuppressionCD,
// 			Type:             CooldownTypeSurvival,

// 			ShouldActivate: func(sim *Simulation, character *Character) bool {
// 				return true
// 			},
// 			AddAura: func(sim *Simulation, character *Character) { psAura.Activate(sim) },
// 		},
// 		numPainSuppressions)
// }

// func PainSuppressionAura(character *Character, actionTag int32) *Aura {
// 	actionID := ActionID{SpellID: 33206, Tag: actionTag}

// 	return character.GetOrRegisterAura(Aura{
// 		Label:    "PainSuppression-" + actionID.String(),
// 		Tag:      PainSuppressionAuraTag,
// 		ActionID: actionID,
// 		Duration: PainSuppressionDuration,
// 		OnGain: func(aura *Aura, sim *Simulation) {
// 			character.PseudoStats.DamageTakenMultiplier *= 0.6
// 		},
// 		OnExpire: func(aura *Aura, sim *Simulation) {
// 			character.PseudoStats.DamageTakenMultiplier /= 0.6
// 		},
// 	})
// }

// var GuardianSpiritAuraTag = "GuardianSpirit"

// const GuardianSpiritDuration = time.Second * 10
// const GuardianSpiritCD = time.Minute * 3

// func registerGuardianSpiritCD(agent Agent, numGuardianSpirits int32) {
// 	if numGuardianSpirits == 0 {
// 		return
// 	}

// 	character := agent.GetCharacter()
// 	gsAura := GuardianSpiritAura(character, -1)
// 	healthMetrics := character.NewHealthMetrics(ActionID{SpellID: 47788})

// 	character.AddDynamicDamageTakenModifier(func(sim *Simulation, _ *Spell, result *SpellResult) {
// 		if (result.Damage >= character.CurrentHealth()) && gsAura.IsActive() {
// 			result.Damage = character.CurrentHealth()
// 			character.GainHealth(sim, 0.5*character.MaxHealth(), healthMetrics)
// 			gsAura.Deactivate(sim)
// 		}
// 	})

// 	registerExternalConsecutiveCDApproximation(
// 		agent,
// 		externalConsecutiveCDApproximation{
// 			ActionID:         ActionID{SpellID: 47788, Tag: -1},
// 			AuraTag:          GuardianSpiritAuraTag,
// 			CooldownPriority: CooldownPriorityLow,
// 			AuraDuration:     GuardianSpiritDuration,
// 			AuraCD:           GuardianSpiritCD,
// 			Type:             CooldownTypeSurvival,

// 			ShouldActivate: func(sim *Simulation, character *Character) bool {
// 				return true
// 			},
// 			AddAura: func(sim *Simulation, character *Character) {
// 				gsAura.Activate(sim)
// 			},
// 		},
// 		numGuardianSpirits)
// }

// func GuardianSpiritAura(character *Character, actionTag int32) *Aura {
// 	actionID := ActionID{SpellID: 47788, Tag: actionTag}

// 	return character.GetOrRegisterAura(Aura{
// 		Label:    "GuardianSpirit-" + actionID.String(),
// 		Tag:      GuardianSpiritAuraTag,
// 		ActionID: actionID,
// 		Duration: GuardianSpiritDuration,
// 		OnGain: func(aura *Aura, sim *Simulation) {
// 			character.PseudoStats.HealingTakenMultiplier *= 1.4
// 		},
// 		OnExpire: func(aura *Aura, sim *Simulation) {
// 			character.PseudoStats.HealingTakenMultiplier /= 1.4
// 		},
// 	})
// }

// func registerRevitalizeHotCD(agent Agent, label string, hotID ActionID, ticks int, tickPeriod time.Duration, uptimeCount int32) {
// 	if uptimeCount == 0 {
// 		return
// 	}

// 	character := agent.GetCharacter()
// 	revActionID := ActionID{SpellID: 48545}

// 	manaMetrics := character.NewManaMetrics(revActionID)
// 	energyMetrics := character.NewEnergyMetrics(revActionID)
// 	rageMetrics := character.NewRageMetrics(revActionID)

// 	// Calculate desired downtime based on selected uptimeCount (1 count = 10% uptime, 0%-100%)
// 	totalDuration := time.Duration(ticks) * tickPeriod
// 	uptimePercent := float64(uptimeCount) / 100.0

// 	aura := character.GetOrRegisterAura(Aura{
// 		Label:    "Revitalize-" + label,
// 		ActionID: hotID,
// 		Duration: totalDuration,
// 		OnGain: func(aura *Aura, sim *Simulation) {
// 			pa := NewPeriodicAction(sim, PeriodicActionOptions{
// 				Period:   tickPeriod,
// 				NumTicks: ticks,
// 				OnAction: func(s *Simulation) {
// 					if s.RandomFloat("Revitalize Proc") < 0.15 {
// 						cpb := aura.Unit.GetCurrentPowerBar()
// 						if cpb == ManaBar {
// 							aura.Unit.AddMana(s, 0.01*aura.Unit.MaxMana(), manaMetrics)
// 						} else if cpb == EnergyBar {
// 							aura.Unit.AddEnergy(s, 8, energyMetrics)
// 						} else if cpb == RageBar {
// 							aura.Unit.AddRage(s, 4, rageMetrics)
// 						}
// 					}
// 				},
// 			})
// 			sim.AddPendingAction(pa)
// 		},
// 	})

// 	ApplyFixedUptimeAura(aura, uptimePercent, totalDuration, 1)
// }

const ShatteringThrowCD = time.Minute * 5

var InnervateAuraTag = "Innervate"

const InnervateDuration = time.Second * 20
const InnervateCD = time.Minute * 6

func InnervateManaThreshold(character *Character) float64 {
	if character.Class == proto.Class_ClassMage {
		// Mages burn mana really fast so they need a higher threshold.
		return character.MaxMana() * 0.7
	} else {
		return 1000
	}
}

func registerInnervateCD(agent Agent, numInnervates int32) {
	if numInnervates == 0 {
		return
	}

	character := agent.GetCharacter()
	innervateThreshold := 0.0
	innervateAura := InnervateAura(character, -1)

	character.Env.RegisterPostFinalizeEffect(func() {
		innervateThreshold = InnervateManaThreshold(character)
	})

	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			ActionID:         ActionID{SpellID: 29166, Tag: -1},
			AuraTag:          InnervateAuraTag,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     InnervateDuration,
			AuraCD:           InnervateCD,
			Type:             CooldownTypeMana,
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// Only cast innervate when very low on mana, to make sure all other mana CDs are prioritized.
				return character.CurrentMana() <= innervateThreshold
			},
			AddAura: func(sim *Simulation, character *Character) {
				innervateAura.Activate(sim)
			},
		},
		numInnervates)
}

func InnervateAura(character *Character, actionTag int32) *Aura {
	actionID := ActionID{SpellID: 29166, Tag: actionTag}
	// TODO: Add metrics for increased regen from spirit (either add here and align ticks to mana tick or create mana tick hook?)
	// manaMetrics := character.NewManaMetrics(actionID)
	return character.GetOrRegisterAura(Aura{
		Label:    "Innervate-" + actionID.String(),
		Tag:      InnervateAuraTag,
		ActionID: actionID,
		Duration: InnervateDuration,
		OnGain: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.SpiritRegenMultiplier += 4
			character.PseudoStats.ForceFullSpiritRegen = true
			character.UpdateManaRegenRates()
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			character.PseudoStats.SpiritRegenMultiplier -= 4
			character.PseudoStats.ForceFullSpiritRegen = false
			character.UpdateManaRegenRates()
		},
	})
}

var ManaTideTotemActionID = ActionID{SpellID: 16190}
var ManaTideTotemAuraTag = "ManaTideTotem"

const ManaTideTotemDuration = time.Second * 12
const ManaTideTotemCD = time.Minute * 5

func registerManaTideTotemCD(agent Agent, numManaTideTotems int32) {
	if numManaTideTotems == 0 {
		return
	}

	character := agent.GetCharacter()
	initialDelay := time.Duration(0)
	mttAura := ManaTideTotemAura(character, -1)

	character.Env.RegisterPostFinalizeEffect(func() {
		// Use first MTT at 60s, or halfway through the fight, whichever comes first.
		initialDelay = min(character.Env.BaseDuration/2, time.Second*60)
	})

	registerExternalConsecutiveCDApproximation(
		agent,
		externalConsecutiveCDApproximation{
			ActionID:         ManaTideTotemActionID.WithTag(-1),
			AuraTag:          ManaTideTotemAuraTag,
			CooldownPriority: CooldownPriorityDefault,
			AuraDuration:     ManaTideTotemDuration,
			AuraCD:           ManaTideTotemCD,
			Type:             CooldownTypeMana,
			ShouldActivate: func(sim *Simulation, character *Character) bool {
				// A normal resto shaman would wait to use MTT.
				return sim.CurrentTime >= initialDelay
			},
			AddAura: func(sim *Simulation, character *Character) {
				mttAura.Activate(sim)
			},
		},
		numManaTideTotems)
}

func ManaTideTotemAura(character *Character, actionTag int32) *Aura {
	actionID := ManaTideTotemActionID.WithTag(actionTag)

	metrics := make([]*ResourceMetrics, len(character.Party.Players))
	for i, player := range character.Party.Players {
		char := player.GetCharacter()
		if char.HasManaBar() {
			metrics[i] = char.NewManaMetrics(actionID)
		}
	}

	manaPerTick := 290.0

	return character.GetOrRegisterAura(Aura{
		Label:    "ManaTideTotem-" + actionID.String(),
		Tag:      ManaTideTotemAuraTag,
		ActionID: actionID,
		Duration: ManaTideTotemDuration,
		OnGain: func(aura *Aura, sim *Simulation) {
			StartPeriodicAction(sim, PeriodicActionOptions{
				Period:   ManaTideTotemDuration / 4,
				NumTicks: 4,
				OnAction: func(sim *Simulation) {
					for i, player := range character.Party.Players {
						if metrics[i] != nil {
							char := player.GetCharacter()
							char.AddMana(sim, manaPerTick, metrics[i])
						}
					}
				},
			})
		},
	})
}

const ReplenishmentAuraDuration = time.Second * 15

// Creates the actual replenishment aura for a unit.
// func replenishmentAura(unit *Unit, _ ActionID) *Aura {
// 	if unit.ReplenishmentAura != nil {
// 		return unit.ReplenishmentAura
// 	}

// 	replenishmentDep := unit.NewDynamicStatDependency(stats.Mana, stats.MP5, 0.01)

// 	unit.ReplenishmentAura = unit.RegisterAura(Aura{
// 		Label:    "Replenishment",
// 		ActionID: ActionID{SpellID: 57669},
// 		Duration: ReplenishmentAuraDuration,
// 		OnGain: func(aura *Aura, sim *Simulation) {
// 			aura.Unit.EnableDynamicStatDep(sim, replenishmentDep)
// 		},
// 		OnExpire: func(aura *Aura, sim *Simulation) {
// 			aura.Unit.DisableDynamicStatDep(sim, replenishmentDep)
// 		},
// 	})

// 	return unit.ReplenishmentAura
// }

/* func spellPowerBonusEffect(aura *Aura, spellPowerBonus float64) *ExclusiveEffect {
	return aura.NewExclusiveEffect("SpellPowerBonus", false, ExclusiveEffect{
		Priority: spellPowerBonus,
		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.AddStatsDynamic(sim, stats.Stats{
				stats.SpellPower: ee.Priority,
			})
		},
		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
			ee.Aura.Unit.AddStatsDynamic(sim, stats.Stats{
				stats.SpellPower: -ee.Priority,
			})
		},
	})
} */

func StrengthOfEarthTotemAura(unit *Unit, multiplier float64) *Aura {
	rank := TernaryInt32(IncludeAQ, 5, 4)
	spellID := []int32{0, 8075, 8160, 8161, 10442, 25361}[rank]
	duration := time.Minute * 2
	updateStats := BuffSpellValues[StrengthOfEarth].Multiply(multiplier).Floor()

	aura := unit.GetOrRegisterAura(Aura{
		Label:      "Strength of Earth Totem",
		ActionID:   ActionID{SpellID: spellID},
		Duration:   duration,
		BuildPhase: CharacterBuildPhaseBuffs,
		OnGain: func(aura *Aura, sim *Simulation) {
			if aura.Unit.Env.MeasuringStats && aura.Unit.Env.State != Finalized {
				unit.AddStats(updateStats)
			} else {
				unit.AddStatsDynamic(sim, updateStats)
			}
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			if aura.Unit.Env.MeasuringStats && aura.Unit.Env.State != Finalized {
				unit.AddStats(updateStats.Multiply(-1))
			} else {
				unit.AddStatsDynamic(sim, updateStats.Multiply(-1))
			}
		},
	})
	return aura
}

func GraceOfAirTotemAura(unit *Unit, multiplier float64) *Aura {
	rank := TernaryInt32(IncludeAQ, 3, 2)
	spellID := []int32{0, 8835, 10627, 25359}[rank]
	duration := time.Minute * 2
	updateStats := BuffSpellValues[GraceOfAir].Multiply(multiplier).Floor()

	aura := unit.GetOrRegisterAura(Aura{
		Label:      "Grace of Air Totem",
		ActionID:   ActionID{SpellID: spellID},
		Duration:   duration,
		BuildPhase: CharacterBuildPhaseBuffs,
		OnGain: func(aura *Aura, sim *Simulation) {
			if aura.Unit.Env.MeasuringStats && aura.Unit.Env.State != Finalized {
				unit.AddStats(updateStats)
			} else {
				unit.AddStatsDynamic(sim, updateStats)
			}
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			if aura.Unit.Env.MeasuringStats && aura.Unit.Env.State != Finalized {
				unit.AddStats(updateStats.Multiply(-1))
			} else {
				unit.AddStatsDynamic(sim, updateStats.Multiply(-1))
			}
		},
	})
	return aura
}

const BattleShoutRanks = 7

var BattleShoutSpellId = [BattleShoutRanks + 1]int32{0, 6673, 5242, 6192, 11549, 11550, 11551, 25289}
var BattleShoutBaseAP = [BattleShoutRanks + 1]float64{0, 20, 40, 57, 93, 138, 193, 232}
var BattleShoutLevel = [BattleShoutRanks + 1]int{0, 1, 12, 22, 32, 42, 52, 60}

func BattleShoutAura(unit *Unit, impBattleShout int32, boomingVoicePts int32, has3pcWrath bool) *Aura {
	rank := TernaryInt32(IncludeAQ, 7, 6)
	spellId := BattleShoutSpellId[rank]
	baseAP := BattleShoutBaseAP[rank]

	return unit.GetOrRegisterAura(Aura{
		Label:      "Battle Shout",
		ActionID:   ActionID{SpellID: spellId},
		Duration:   time.Duration(float64(time.Minute*2) * (1 + 0.1*float64(boomingVoicePts))),
		BuildPhase: CharacterBuildPhaseBuffs,
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.AddStatsDynamic(sim, stats.Stats{
				stats.AttackPower: math.Floor(baseAP*(1+0.05*float64(impBattleShout)) + TernaryFloat64(has3pcWrath, 30, 0)),
			})
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.AddStatsDynamic(sim, stats.Stats{
				stats.AttackPower: -1 * math.Floor(baseAP*(1+0.05*float64(impBattleShout))+TernaryFloat64(has3pcWrath, 30, 0)),
			})
		},
	})
}

func TrueshotAura(unit *Unit) *Aura {
	rangedAP := 100.0
	meleeAP := 100.0

	aura := MakePermanent(unit.RegisterAura(Aura{
		Label:    "Trueshot Aura",
		ActionID: ActionID{SpellID: 20906},
	}))

	makeExclusiveBuff(aura, BuffConfig{
		Category: "TrueshotAura",
		Stats: []StatConfig{
			{stats.AttackPower, meleeAP, false},
			{stats.RangedAttackPower, rangedAP, false},
		},
	})

	return aura
}

func BlessingOfMightAura(unit *Unit, impBomPts int32) *Aura {
	spellID := TernaryInt32(IncludeAQ, 25291, 19838)

	bonusAP := math.Floor(BuffSpellValues[BlessingOfMight][stats.AttackPower] * (1 + 0.04*float64(impBomPts)))

	aura := MakePermanent(unit.GetOrRegisterAura(Aura{
		Label:      "Blessing of Might",
		ActionID:   ActionID{SpellID: spellID},
		Duration:   NeverExpires,
		BuildPhase: CharacterBuildPhaseBuffs,
	}))

	makeExclusiveBuff(aura, BuffConfig{
		Category: "Paladin Physical Buffs",
		Stats: []StatConfig{
			{stats.AttackPower, bonusAP, false},
		},
	})

	return aura
}

// TODO: Are there exclusive AP buffs in SoD?
// func attackPowerBonusEffect(aura *Aura, apBonus float64) *ExclusiveEffect {
// 	return aura.NewExclusiveEffect("AttackPowerBonus", false, ExclusiveEffect{
// 		Priority: apBonus,
// 		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
// 			ee.Aura.Unit.AddStatsDynamic(sim, stats.Stats{
// 				stats.AttackPower:       ee.Priority,
// 				stats.RangedAttackPower: ee.Priority,
// 			})
// 		},
// 		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
// 			ee.Aura.Unit.AddStatsDynamic(sim, stats.Stats{
// 				stats.AttackPower:       -ee.Priority,
// 				stats.RangedAttackPower: -ee.Priority,
// 			})
// 		},
// 	})
// }

func BattleSquawkAura(character *Unit, stackcount int32) *Aura {
	aura := character.GetOrRegisterAura(Aura{
		Label:      "Battle Squawk",
		ActionID:   ActionID{SpellID: 23060},
		Duration:   time.Minute * 4,
		MaxStacks:  5,
		BuildPhase: CharacterBuildPhaseBuffs,
		OnReset: func(aura *Aura, sim *Simulation) {
			aura.Activate(sim)
		},
		OnGain: func(aura *Aura, sim *Simulation) {
			aura.SetStacks(sim, stackcount)
		},
		OnStacksChange: func(aura *Aura, sim *Simulation, oldStacks, newStacks int32) {
			character.MultiplyMeleeSpeed(sim, math.Pow(1.05, float64(newStacks-oldStacks)))
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.SetStacks(sim, 0)
		},
	})
	return aura
}

// func healthBonusEffect(aura *Aura, healthBonus float64) *ExclusiveEffect {
// 	return aura.NewExclusiveEffect("HealthBonus", false, ExclusiveEffect{
// 		Priority: healthBonus,
// 		OnGain: func(ee *ExclusiveEffect, sim *Simulation) {
// 			ee.Aura.Unit.AddStatsDynamic(sim, stats.Stats{
// 				stats.Health: ee.Priority,
// 			})
// 		},
// 		OnExpire: func(ee *ExclusiveEffect, sim *Simulation) {
// 			ee.Aura.Unit.AddStatsDynamic(sim, stats.Stats{
// 				stats.Health: -ee.Priority,
// 			})
// 		},
// 	})
// }

func CreateExtraAttackAuraCommon(character *Character, buffActionID ActionID, auraLabel string, rank int32, getBonusAP func(aura *Aura, rank int32) float64) *Aura {
	var bonusAP float64

	apBuffAura := character.GetOrRegisterAura(Aura{
		Label:     auraLabel + " Buff",
		ActionID:  buffActionID,
		Duration:  time.Millisecond * 1500,
		MaxStacks: 2,
		OnGain: func(aura *Aura, sim *Simulation) {
			bonusAP = getBonusAP(aura, rank)
			aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.AttackPower: bonusAP})
		},
		OnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.AttackPower: -bonusAP})
		},
	})

	MakePermanent(character.GetOrRegisterAura(Aura{
		Label:     "Extra Attacks  (Main Hand)", // Tracks Stored Extra Attacks from all sources
		ActionID:  ActionID{SpellID: 21919},     // Thrash ID
		Duration:  NeverExpires,
		MaxStacks: 4, // Max is 4 extra attacks stored - more can proc after
		OnInit: func(aura *Aura, sim *Simulation) {
			aura.Unit.AutoAttacks.mh.extraAttacksAura = aura
		},
	}))

	icd := Cooldown{
		Timer:    character.NewTimer(),
		Duration: time.Millisecond * 1500,
	}

	apBuffAura.Icd = &icd

	MakePermanent(character.GetOrRegisterAura(Aura{
		Label: auraLabel,
		OnSpellHitDealt: func(aura *Aura, sim *Simulation, spell *Spell, result *SpellResult) {
			// charges are removed by every auto or next melee, whether it lands or not
			//  this directly contradicts https://github.com/magey/classic-warrior/wiki/Windfury-Totem#triggered-by-melee-spell-while-an-on-next-swing-attack-is-queued
			//  but can be seen in both "vanilla" and "sod" era logs
			if apBuffAura.IsActive() && spell.ProcMask.Matches(ProcMaskMeleeWhiteHit) {
				apBuffAura.RemoveStack(sim)
			}

			if !result.Landed() || !spell.ProcMask.Matches(ProcMaskMeleeMH) || spell.Flags.Matches(SpellFlagSuppressEquipProcs) {
				return
			}

			if icd.IsReady(sim) && sim.RandomFloat(auraLabel) < 0.2 {
				icd.Use(sim)
				apBuffAura.Activate(sim)
				// aura is up _before_ the triggering swing lands, so if triggered by an auto attack, the aura fades right after the extra attack lands.
				if spell.ProcMask == ProcMaskMeleeMHAuto {
					apBuffAura.SetStacks(sim, 1)
				} else {
					apBuffAura.SetStacks(sim, 2)
				}

				aura.Unit.AutoAttacks.ExtraMHAttackProc(sim, 1, buffActionID, spell)
			}
		},
	}))

	return apBuffAura
}

func GetWildStrikesAP(aura *Aura, rank int32) float64 {
	return 0.2 * aura.Unit.GetStat(stats.AttackPower)
}

const WindfuryRanks = 3

var (
	WindfuryBuffSpellId = [WindfuryRanks + 1]int32{0, 8516, 10608, 10610}
	WindfuryBuffBonusAP = [WindfuryRanks + 1]float64{0, 122, 229, 315}
)

func GetWindfuryAP(aura *Aura, rank int32) float64 {
	return WindfuryBuffBonusAP[rank]
}

func ApplyWindfury(character *Character) *Aura {
	rank := int32(3)
	spellId := WindfuryBuffSpellId[rank]
	buffActionID := ActionID{SpellID: spellId}

	return CreateExtraAttackAuraCommon(character, buffActionID, "Windfury", rank, GetWindfuryAP)

}

///////////////////////////////////////////////////////////////////////////
//                            World Buffs
///////////////////////////////////////////////////////////////////////////

func ApplyDragonslayerBuffs(unit *Unit, buffs *proto.IndividualBuffs) {
	eeCategory := "DragonslayerBuff"
	if buffs.RallyingCryOfTheDragonslayer {
		ApplyRallyingCryOfTheDragonslayer(unit, eeCategory)
	}
}

func ApplyRallyingCryOfTheDragonslayer(unit *Unit, category string) {
	aura := MakePermanent(unit.RegisterAura(Aura{
		Label:      "Rallying Cry of the Dragonslayer",
		ActionID:   ActionID{SpellID: 22888},
		BuildPhase: CharacterBuildPhaseBuffs,
	}))

	makeExclusiveBuff(aura, BuffConfig{
		Category: category,
		Stats: []StatConfig{
			{stats.SpellCrit, 10 * SpellCritRatingPerCritChance, false},
			{stats.MeleeCrit, 5 * CritRatingPerCritChance, false},
			// TODO: {stats.RangedCrit, 5*CritRatingPerCritChance, false},
			{stats.AttackPower, 140, false},
			{stats.RangedAttackPower, 140, false},
		},
	})
}

func ApplySpiritOfZandalar(unit *Unit) {
	aura := MakePermanent(unit.RegisterAura(Aura{
		Label:      "Spirit of Zandalar",
		ActionID:   ActionID{SpellID: 24425},
		BuildPhase: CharacterBuildPhaseBuffs,
		OnInit: func(aura *Aura, sim *Simulation) {
			unit.AddMoveSpeedModifier(&aura.ActionID, 1.10)
		},
	}))

	makeExclusiveBuff(aura, BuffConfig{
		Category: "ZandalarBuff",
		Stats: []StatConfig{
			{stats.Agility, 1.15, true},
			{stats.Intellect, 1.15, true},
			{stats.Spirit, 1.15, true},
			{stats.Stamina, 1.15, true},
			{stats.Strength, 1.15, true},
		},
	})
}

func ApplySongflowerSerenade(unit *Unit) {
	aura := MakePermanent(unit.RegisterAura(Aura{
		Label:      "Songflower Serenade",
		ActionID:   ActionID{SpellID: 15366},
		BuildPhase: CharacterBuildPhaseBuffs,
	}))

	makeExclusiveBuff(aura, BuffConfig{
		Category: "SongflowerSerenade",
		Stats: []StatConfig{
			{stats.Agility, 15, false},
			{stats.Intellect, 15, false},
			{stats.Spirit, 15, false},
			{stats.Stamina, 15, false},
			{stats.Strength, 15, false},
			{stats.MeleeCrit, 5, false},
			// TODO: {stats.RangedCrit, 5, false},
			{stats.SpellCrit, 5, false},
		},
	})
}

func ApplyWarchiefsBuffs(unit *Unit, buffs *proto.IndividualBuffs, isAlliance bool, isHorde bool) {
	if buffs.WarchiefsBlessing /* && isHorde */ {
		ApplyWarchiefsBlessing(unit, "WarchiefsBuff")
	}
}

func ApplyWarchiefsBlessing(unit *Unit, category string) {
	aura := MakePermanent(unit.RegisterAura(Aura{
		Label:      "Warchief's Blessing",
		ActionID:   ActionID{SpellID: 16609},
		BuildPhase: CharacterBuildPhaseBuffs,
	}))

	makeExclusiveBuff(aura, BuffConfig{
		Category: category,
		Stats: []StatConfig{
			{stats.Health, 300, false},
			{stats.MP5, 10, false},
		},
		ExtraOnGain: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.MeleeSpeedMultiplier *= 1.15
		},
		ExtraOnExpire: func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.MeleeSpeedMultiplier /= 1.15
		},
	})
}

func ApplyFengusFerocity(unit *Unit) {
	aura := MakePermanent(unit.RegisterAura(Aura{
		Label:      "Fengus' Ferocity",
		ActionID:   ActionID{SpellID: 22817},
		BuildPhase: CharacterBuildPhaseBuffs,
	}))

	makeExclusiveBuff(aura, BuffConfig{
		Category: "FengusFerocity",
		Stats: []StatConfig{
			{stats.AttackPower, 200, false},
			{stats.RangedAttackPower, 200, false},
		},
	})
}

func ApplyMoldarsMoxie(unit *Unit) {
	aura := MakePermanent(unit.RegisterAura(Aura{
		Label:      "Moldar's Moxie",
		ActionID:   ActionID{SpellID: 22818},
		BuildPhase: CharacterBuildPhaseBuffs,
	}))

	makeExclusiveBuff(aura, BuffConfig{
		Category: "MoldarsMoxie",
		Stats: []StatConfig{
			{stats.Stamina, 1.15, true},
		},
	})
}

func ApplySlipkiksSavvy(unit *Unit) {
	aura := MakePermanent(unit.RegisterAura(Aura{
		Label:      "Slip'kik's Savvy",
		ActionID:   ActionID{SpellID: 22820},
		BuildPhase: CharacterBuildPhaseBuffs,
	}))

	makeExclusiveBuff(aura, BuffConfig{
		Category: "SlipkiksSavvy",
		Stats: []StatConfig{
			{stats.SpellCrit, 3 * SpellCritRatingPerCritChance, false},
		},
	})
}

func ApplySaygesFortunes(character *Character, fortune proto.SaygesFortune) {
	var label string
	var spellID int32

	config := BuffConfig{
		Category: "SaygesFortune",
	}

	switch fortune {
	case proto.SaygesFortune_SaygesDamage:
		label = "Sayge's Dark Fortune of Damage"
		spellID = 23768
		config.ExtraOnGain = func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.DamageDealtMultiplier *= 1.10
		}
		config.ExtraOnExpire = func(aura *Aura, sim *Simulation) {
			aura.Unit.PseudoStats.DamageDealtMultiplier /= 1.10
		}
	case proto.SaygesFortune_SaygesAgility:
		label = "Sayge's Dark Fortune of Agility"
		spellID = 23736
		addAgility := character.GetBaseStats()[stats.Agility] * 0.1
		config.Stats = []StatConfig{
			{stats.Agility, addAgility, false},
		}
	case proto.SaygesFortune_SaygesIntellect:
		label = "Sayge's Dark Fortune of Intellect"
		spellID = 23766
		addIntellect := character.GetBaseStats()[stats.Intellect] * 0.1
		config.Stats = []StatConfig{
			{stats.Intellect, addIntellect, false},
		}
	case proto.SaygesFortune_SaygesSpirit:
		label = "Sayge's Dark Fortune of Spirit"
		spellID = 23738
		addSpirit := character.GetBaseStats()[stats.Spirit] * 0.1
		config.Stats = []StatConfig{
			{stats.Spirit, addSpirit, false},
		}
	case proto.SaygesFortune_SaygesStamina:
		label = "Sayge's Dark Fortune of Stamina"
		spellID = 23737
		addStamina := character.GetBaseStats()[stats.Stamina] * 0.1
		config.Stats = []StatConfig{
			{stats.Stamina, addStamina, false},
		}
	}

	aura := MakePermanent(character.RegisterAura(Aura{
		Label:      label,
		ActionID:   ActionID{SpellID: spellID},
		BuildPhase: CharacterBuildPhaseBuffs,
	}))

	makeExclusiveBuff(aura, config)
}

///////////////////////////////////////////////////////////////////////////
//                            Misc Other Buffs
///////////////////////////////////////////////////////////////////////////

// Equip: Increases healing done by magical spells and effects of all party members within 30 yards by up to 62.
func AtieshHealingEffect(unit *Unit) *Aura {
	label := "Atiesh Greatstaff of the Guardian (Healing)"

	if unit.HasAura(label) {
		return unit.GetAura(label)
	}

	stats := stats.Stats{
		stats.HealingPower: 62,
	}

	return MakePermanent(unit.RegisterAura(Aura{
		ActionID:   ActionID{SpellID: 28144},
		Label:      label,
		BuildPhase: CharacterBuildPhaseBuffs,
	}).AttachStatsBuff(stats))
}

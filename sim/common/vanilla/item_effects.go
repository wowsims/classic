package vanilla

import (
	"fmt"
	"time"

	"github.com/wowsims/classic/sim/common/itemhelpers"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

// Ordered by ID
const (
	ShortswordOfVengeance    = 754
	FieryWarAxe              = 870
	Bloodrazor               = 809
	HammerOfTheNorthernWind  = 810
	FlurryAxe                = 871
	ManualCrowdPummeler      = 9449
	SkullflameShield         = 1168
	Nightblade               = 1982
	Shadowblade              = 2163
	GutRipper                = 2164
	HandOfEdwardTheOdd       = 2243
	BowOfSearingArrows       = 2825
	Gutwrencher              = 5616
	Ravager                  = 7717
	HanzoSword               = 8190
	TheJackhammer            = 9423
	PendulumOfDoom           = 9425
	BloodletterScalpel       = 9511
	TheHandOfAntusul         = 9639
	GryphonRidersStormhammer = 9651
	Firebreather             = 10797
	VilerendSlicer           = 11603
	HookfangShanker          = 11635
	LinkensSwordOfMastery    = 11902
	SearingNeedle            = 12531
	PipsSkinner              = 12709
	ArcaniteChampion         = 12790
	MasterworkStormhammer    = 12794
	Frostguard               = 12797
	SerpentSlicer            = 13035
	TheNeedler               = 13060
	SealOfTheDawn            = 13209
	JoonhosMercy             = 17054
	Deathbringer             = 17068
	ViskagTheBloodletter     = 17075
	ThrashBlade              = 17705
	SatyrsLash               = 17752
	MarkOfTheChosen          = 17774
	Nightfall                = 19169
	EbonHand                 = 19170
	RuneOfTheDawn            = 19812
	ZandalariHeroBadge       = 19948
	ZandalariHeroMedallion   = 19949
	ZandalariHeroCharm       = 19950
	MarkOfTheChampionPhys    = 23206
	MarkOfTheChampionSpell   = 23207
	Ragehammer               = 10626
	SulfurasHandOfRagnaros   = 17182
	SulfuronHammer           = 17193
	FlameWrath               = 11809
	LordGeneralsSword        = 11817
	WraithScythe             = 11920
	SecondWind               = 11819
	BurstOfKnowledge         = 11832
	Ironfoe                  = 11684
	EbonHiltOfMarduk         = 14576
	FrightskullShaft         = 14531
	BarovianFamilySword      = 14541
	Frightalon               = 14024
	// HeadmastersCharge         = 13937
	GravestoneWarAxe          = 13983
	FiendishMachete           = 18310
	TalismanOfEphemeralPower  = 18820
	GutgoreRipper             = 17071
	Shadowstrike              = 17074
	Thunderstrike             = 17223
	BonereaversEdge           = 17076
	EssenceOfThePureFlame     = 18815
	PerditionsBlade           = 18816
	EskhandarsLeftClaw        = 18202
	EskhandarsRightClaw       = 18203
	BlazefuryMedallion        = 17111
	EmpyreanDemolisher        = 17112
	SkullforgeReaver          = 13361
	RunebladeOfBaronRivendare = 13505
	HeartOfWyrmthalak         = 22321
	Venomspitter              = 13183
	SmolderwebsEye            = 13213
	Chillpike                 = 13148
	FangOfTheCrystalSpider    = 13218
	BlackhandDoomsaw          = 12583
	BlackbladeOfShahram       = 12592
	SeepingWillow             = 12969
	DraconicInfusedEmblem     = 22268
	QuelSerrar                = 18348
	HandOfJustice             = 11815
	Felstriker                = 12590
	Thunderfury               = 19019
	TheUntamedBlade           = 19334
	DrakeTalonCleaver         = 19353
	JekliksCrusher            = 19918
	ZulianSlicer              = 19901
	HalberdOfSmiting          = 19874
	NatPaglesBrokenReel       = 19947
	TigulesHarpoon            = 19946
	GrileksCarver             = 19962
	GrileksGrinder            = 19961
	PitchforkOfMadness        = 19963
	ScarabBrooch              = 21625
	KalimdorsRevenge          = 21679
	JomGabbar                 = 23570
	EyeOfMoam                 = 21473
	DarkmoonCardHeroism       = 19287
	DarkmoonCardBlueDragon    = 19288
	DarkmoonCardMaelstrom     = 19289
	Earthstrike               = 21180
	WrathOfCenarius           = 21190
)

func init() {
	core.AddEffectsToTest = false

	// ! Please keep items ordered alphabetically within a given category !

	// Many referenced from Classic WoW Armaments - https://discord.gg/w95B2hXfBF

	///////////////////////////////////////////////////////////////////////////
	//                                 Weapons
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/spell=16916/strength-of-the-champion
	// Chance on hit: Heal self for 270 to 450 and Increases Strength by 120 for 30 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(ArcaniteChampion, "Arcanite Champion", 1.0, StrengthOfTheChampionAura)

	// https://www.wowhead.com/classic/item=14541/barovian-family-sword
	// Chance on hit: Deals 30 Shadow damage every 3 sec for 15 sec. All damage done is then transferred to the caster.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(BarovianFamilySword, "Barovian Family Sword", 0.5, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 18652}

		// Keep track of damage taken by each enemy
		enemyDamageTaken := map[int32]float64{}
		for _, target := range character.Env.Encounter.TargetUnits {
			enemyDamageTaken[target.UnitIndex] = 0
		}

		healthMetrics := character.NewHealthMetrics(actionID)

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,

			Dot: core.DotConfig{
				NumberOfTicks: 5,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Siphon Health (Barovian Family Sword)",
				},
				OnSnapshot: func(_ *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					enemyDamageTaken[target.UnitIndex] = 0
					dot.Snapshot(target, 30, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					result := dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
					enemyDamageTaken[target.UnitIndex] += result.Damage
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				enemyDamageTaken[target.UnitIndex] = 0
				spell.Dot(target).Apply(sim)
			},
		})

		// The healing is applied at the end of the DoT and can crit according to old comments
		for _, dot := range spell.Dots() {
			if dot != nil {
				unit := dot.Unit
				dot.ApplyOnExpire(func(aura *core.Aura, sim *core.Simulation) {
					// TODO: This may not be quite correct but it's close enough
					result := spell.CalcDamage(sim, unit, enemyDamageTaken[unit.UnitIndex], spell.OutcomeHealingCrit)
					character.GainHealth(sim, result.Damage, healthMetrics)
				})
			}
		}

		return spell
	})

	// https://www.wowhead.com/classic/item=12592/blackblade-of-shahram
	// Chance on hit: Summons the infernal spirit of Shahram.
	// Summons an NPC "Shahram" who has an equal chance to cast one of 6 spells:
	// Curse of Shahram: -50% movement speed and -25% attack speed on all enemies within 10 yards of Shahram for 10 seconds.
	// Might of Shahram: 5-second stun on all enemies within 10 yards of Shahram.
	// Fist of Shahram: +30% Melee Attack Speed for all party members within 30 yards of Shahram for 8 seconds.
	// Blessing of Shahram: Restores 50 health and mana every 5 seconds for all party members within 30 yards of Shahram for 20 seconds. The Healing portion of this effect scales at 100% of self-healing buffs such as Amplify Magic.
	// Will of Shahram: +50 all stats for yourself for 20 seconds.
	// Flames of Shahram: Deals 100-150 Fire damage to all enemies within 10 yards of Shahram. Damage scales at 100% with +spelldmg debuffs placed on enemies such as Flame Buffet.
	//
	// Implementing this without the guardian as it seems to just cast a spell and depart and guardians are expensive
	core.NewItemEffect(BlackbladeOfShahram, func(agent core.Agent) {
		character := agent.GetCharacter()

		curseOfShahramAuras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			aura := target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16597},
				Label:    "Curse of Shahram",
				Duration: time.Second * 10,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.MultiplyAttackSpeed(sim, 1/1.25)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.MultiplyAttackSpeed(sim, 1.25)
				},
			})
			core.AtkSpeedReductionEffect(aura, 1.25)
			return aura
		})
		curseOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16597},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				curseOfShahramAuras.Get(target).Activate(sim)
			},
		})

		mightOfShahramAuras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16600},
				Label:    "Might of Shahram",
				Duration: time.Second * 5,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.Stunned = true
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.PseudoStats.Stunned = false
				},
			})
		})

		mightOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16600},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					mightOfShahramAuras.Get(aoeTarget).Activate(sim)
				}
			},
		})

		fistOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16601},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				counter := 0

				for counter < 10 {
					fistOfShahramAura := character.GetOrRegisterAura(core.Aura{
						ActionID: core.ActionID{SpellID: 16601},
						Label:    fmt.Sprintf("Fist of Shahram (%d)", counter),
						Duration: time.Second * 8,
						OnGain: func(aura *core.Aura, sim *core.Simulation) {
							character.MultiplyAttackSpeed(sim, 1.3)
						},
						OnExpire: func(aura *core.Aura, sim *core.Simulation) {
							character.MultiplyAttackSpeed(sim, 1/(1.3))
						},
					})

					if !fistOfShahramAura.IsActive() {
						fistOfShahramAura.Activate(sim)
						break
					}

					counter += 1

				}
			},
		})

		blessingOfShahramManaMetrics := character.NewPartyManaMetrics(core.ActionID{SpellID: 16599})
		blessingOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16599},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagIgnoreAttackerModifiers | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,
			Hot: core.DotConfig{
				Aura: core.Aura{
					Label: "Blessing of Shahram",
				},
				NumberOfTicks: 4,
				TickLength:    time.Second * 5,
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, _ bool) {
					dot.SnapshotBaseDamage = 50
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotHealing(sim, target, dot.OutcomeTick)
					if target.HasManaBar() {
						target.AddMana(sim, 50, blessingOfShahramManaMetrics[target.UnitIndex])
					}
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, agent := range character.Party.PlayersAndPets {
					spell.Hot(&agent.GetCharacter().Unit).Apply(sim)
				}
			},
		})

		willOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 16598},
			SpellSchool: core.SpellSchoolArcane,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				counter := 0

				stats := stats.Stats{
					stats.Agility:   50,
					stats.Intellect: 50,
					stats.Stamina:   50,
					stats.Spirit:    50,
					stats.Strength:  50,
				}

				for counter < 10 {
					willOfShahram := character.GetOrRegisterAura(core.Aura{
						ActionID: core.ActionID{SpellID: 16598},
						Label:    fmt.Sprintf("Will of Shahram (%d)", counter),
						Duration: time.Second * 20,
						OnGain: func(aura *core.Aura, sim *core.Simulation) {
							character.AddStatsDynamic(sim, stats)
						},
						OnExpire: func(aura *core.Aura, sim *core.Simulation) {
							character.AddStatsDynamic(sim, stats.Multiply(-1.0))
						},
					})

					if !willOfShahram.IsActive() {
						willOfShahram.Activate(sim)
						break
					}

					counter += 1

				}
			},
		})

		flamesOfShahram := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16596},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagIgnoreAttackerModifiers | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, 90, spell.OutcomeMagicCrit)
				}
			},
		})

		castableSpells := []*core.Spell{curseOfShahram, mightOfShahram, fistOfShahram, blessingOfShahram, willOfShahram, flamesOfShahram}
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Summon Shahram",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				spellIdx := int32(sim.Roll(0, 6))
				castableSpells[spellIdx].Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=12583/blackhand-doomsaw
	// Chance on hit: Wounds the target for 324 to 540 damage.
	// TODO: Proc rate based on the original item
	itemhelpers.CreateWeaponCoHProcDamage(BlackhandDoomsaw, "Blackhand Doomsaw", 0.4, 16549, core.SpellSchoolPhysical, 324, 216, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=10626/ragehammer
	// Chance on hit: Increases damage done by 20 and attack speed by 5% for 15 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(Ragehammer, "Ragehammer", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 12686},
			Label:    "Enrage (12686)",
			Duration: time.Second * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.BonusPhysicalDamage += 20
				character.MultiplyAttackSpeed(sim, 1.05)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.BonusPhysicalDamage -= 20
				character.MultiplyAttackSpeed(sim, 1/1.05)
			},
		})
	})

	// https://www.wowhead.com/classic/item=9511/bloodletter-scalpel
	itemhelpers.CreateWeaponCoHProcDamage(BloodletterScalpel, "Bloodletter Scalpel", 1.0, 18081, core.SpellSchoolPhysical, 60, 10, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=809/bloodrazor
	itemhelpers.CreateWeaponProcSpell(Bloodrazor, "Bloodrazor", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 17504},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Rend (Bloodrazor)",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 12, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=17076/bonereavers-edge
	// Chance on hit: Your attacks ignore 700 of your enemies' armor for 10 sec. This effect stacks up to 3 times.
	itemhelpers.CreateWeaponProcSpell(BonereaversEdge, "Bonereaver's Edge", 2.0, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 21153}
		buffAura := character.RegisterAura(core.Aura{
			ActionID:  actionID,
			Label:     "Bonereaver's Edge",
			Duration:  time.Second * 10,
			MaxStacks: 3,
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				for _, target := range sim.Encounter.TargetUnits {
					target.AddStatDynamic(sim, stats.Armor, 700*float64(oldStacks))
					target.AddStatDynamic(sim, stats.Armor, -700*float64(newStacks))
				}
			},
		})
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			ProcMask:    core.ProcMaskEmpty,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
				buffAura.AddStack(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=2825/bow-of-searing-arrows
	// Equip: Chance to strike your ranged target with a Searing Arrow for 18 to 26 Fire damage.
	itemhelpers.CreateWeaponProcSpell(BowOfSearingArrows, "Bow of Searing Arrows", 3.35, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 29638},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeRanged,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				dmg := sim.Roll(18, 26)
				spell.CalcAndDealDamage(sim, target, dmg, spell.OutcomeRangedCritOnly)
			},
		})
	})

	// https://www.wowhead.com/classic/item=13148/chillpike
	// Chance on hit: Blasts a target for 160 to 250 Frost damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponCoHProcDamage(Chillpike, "Chillpike", 1.0, 19260, core.SpellSchoolFrost, 160, 90, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=17068/deathbringer
	// Chance on hit: Sends a shadowy bolt at the enemy causing 110 to 140 Shadow damage.
	itemhelpers.CreateWeaponCoHProcDamage(Deathbringer, "Deathbringer", 1.0, 18138, core.SpellSchoolShadow, 110, 30, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=19353/drake-talon-cleaver
	// Chance on hit: Delivers a fatal wound for 240 damage.
	// Original proc rate 1.0 increased to approximately 1.60 in SoD phase 5
	itemhelpers.CreateWeaponCoHProcDamage(DrakeTalonCleaver, "Drake Talon Cleaver", 1.0, 467167, core.SpellSchoolPhysical, 240, 0, 0.0, core.DefenseTypeMelee) // TBD confirm 1 ppm in SoD

	// https://www.wowhead.com/classic/item=19170/ebon-hand
	// Chance on hit: Sends a shadowy bolt at the enemy causing 125 to 275 Shadow damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponCoHProcDamage(EbonHand, "Ebon Hand", 1.0, 18211, core.SpellSchoolShadow, 125, 150, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=14576/ebon-hilt-of-marduk
	// Chance on hit: Corrupts the target, causing 210 damage over 3 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(EbonHiltOfMarduk, "Ebon Hilt of Marduk", 1.0, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18656},
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Corruption (Ebon Hilt of Marduk)",
				},
				TickLength:    time.Second,
				NumberOfTicks: 3,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 70, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=17112/empyrean-demolisher
	// Chance on hit: Increases your attack speed by 20% for 10 sec.
	itemhelpers.CreateWeaponProcAura(EmpyreanDemolisher, "Empyrean Demolisher", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			Label:    "Empyrean Demolisher Haste Aura",
			ActionID: core.ActionID{SpellID: 21165},
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1.2)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1/1.2)
			},
		})
	})

	// https://www.wowhead.com/classic/item=18202/eskhandars-left-claw
	// Chance on hit: Slows enemy's movement by 60% and causes them to bleed for 150 damage over 30 sec.
	// TODO: Proc rate untested
	itemhelpers.CreateWeaponProcSpell(EskhandarsLeftClaw, "Eskhandar's Left Claw", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 22639},
			SpellSchool: core.SpellSchoolPhysical,
			DefenseType: core.DefenseTypeMelee,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Eskhandar's Rake",
				},
				TickLength:    time.Second * 3,
				NumberOfTicks: 10,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 15, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=18203/eskhandars-right-claw
	// Chance on hit: Increases your attack speed by 30% for 5 sec.
	itemhelpers.CreateWeaponProcAura(EskhandarsRightClaw, "Eskhandar's Right Claw", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			Label:    "Eskhandar's Rage",
			ActionID: core.ActionID{SpellID: 22640},
			Duration: time.Second * 5,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1.3)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1/1.3)
			},
		})
	})

	// https://www.wowhead.com/classic/item=13218/fang-of-the-crystal-spider
	// Chance on hit: Slows target enemy's casting speed and increases the time between melee and ranged attacks by 10% for 10 sec.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(FangOfTheCrystalSpider, func(agent core.Agent) {
		character := agent.GetCharacter()

		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit) *core.Aura {
			aura := unit.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 17331},
				Label:    "Fang of the Crystal Spider",
				Duration: time.Second * 10,
			})
			core.AtkSpeedReductionEffect(aura, 1.10)
			return aura
		})

		procMask := character.GetProcMaskForItem(FangOfTheCrystalSpider)
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Fang of the Crystal Spider Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          procMask,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				debuffAuras.Get(result.Target).Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=12590/felstriker
	// Chance on hit: All attacks are guaranteed to land and will be critical strikes for the next 3 sec.
	core.NewItemEffect(Felstriker, func(agent core.Agent) {
		character := agent.GetCharacter()

		effectAura := character.NewTemporaryStatsAura("Felstriker", core.ActionID{SpellID: 16551}, stats.Stats{stats.MeleeCrit: 100 * core.CritRatingPerCritChance, stats.MeleeHit: 100 * core.MeleeHitRatingPerHitChance}, time.Second*3)
		procMask := character.GetProcMaskForItem(Felstriker)
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Felstriker Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          procMask,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				effectAura.Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=18310/fiendish-machete
	core.NewItemEffect(FiendishMachete, func(agent core.Agent) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeElemental {
			character.PseudoStats.MobTypeAttackPower += 36
		}
	})

	// https://www.wowhead.com/classic/item=870/fiery-war-axe
	itemhelpers.CreateWeaponProcSpell(FieryWarAxe, "Fiery War Axe", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18796},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Fiery War Axe Fireball",
				},
				TickLength:    2 * time.Second,
				NumberOfTicks: 3,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				dmg := sim.Roll(155, 197)
				result := spell.CalcAndDealDamage(sim, target, dmg, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(Firebreather, "Firebreather", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16413},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 70, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 3,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Fireball",
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 3, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=11809/flame-wrath
	// Chance on hit: Envelops the caster with a Fire shield for 15 sec and shoots a ring of fire dealing 130 to 170 damage to all nearby enemies.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(FlameWrath, "Flame Wrath", 1.0, func(character *core.Character) *core.Spell {
		shieldActionID := core.ActionID{SpellID: 461152}
		shieldSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:         shieldActionID,
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			BonusCoefficient: 1, // Only the shield portion has scaling
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 10, spell.OutcomeAlwaysHit)
			},
		})
		shieldAura := character.RegisterAura(core.Aura{
			ActionID: shieldActionID,
			Label:    "Flame Wrath",
			Duration: time.Second * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.FireResistance, 30)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.FireResistance, -30)
			},
			OnSpellHitTaken: func(_ *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() {
					shieldSpell.Cast(sim, spell.Unit)
				}
			},
		})
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 461151},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				shieldAura.Activate(sim)

				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, sim.Roll(130, 170), spell.OutcomeMagicHit)
				}
			},
		})
	})

	itemhelpers.CreateWeaponProcSpell(FlurryAxe, "Flurry Axe", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 18797},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttackProc(sim, 1, core.ActionID{SpellID: 18797}, spell)
			},
		})
	})

	// https://www.wowhead.com/classic/item=14024/frightalon
	// Chance on hit: Lowers all attributes of target by 10 for 1 min.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(Frightalon, func(agent core.Agent) {
		character := agent.GetCharacter()
		procMask := character.GetProcMaskForItem(Frightalon)

		debuffAuraArray := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 19755},
				Label:    "Frightalon",
				Duration: time.Minute * 1,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   -10,
						stats.Intellect: -10,
						stats.Stamina:   -10,
						stats.Spirit:    -10,
						stats.Strength:  -10,
					})
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   10,
						stats.Intellect: 10,
						stats.Stamina:   10,
						stats.Spirit:    10,
						stats.Strength:  10,
					})
				},
			})
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Frightalon Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          procMask,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1.0,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				debuffAuraArray.Get(result.Target).Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=14531/frightskull-shaft
	// Chance on hit: Deals 8 Shadow damage every 2 sec for 30 sec and lowers their Strength for the duration of the disease.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(FrightskullShaft, "Frightskull Shaft", 0.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18633},
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPureDot | core.SpellFlagDisease,

			Dot: core.DotConfig{
				NumberOfTicks: 15,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Weakening Disease",
					OnGain: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatDynamic(sim, stats.Strength, -50)
					},
					OnExpire: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatDynamic(sim, stats.Strength, 50)
					},
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=12797/frostguard#comments
	// Chance on hit: Target's movement slowed by 30% and increasing the time between attacks by 25% for 5 sec.
	// TODO: Proc rate assumed and needs testing
	core.NewItemEffect(Frostguard, func(agent core.Agent) {
		character := agent.GetCharacter()
		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit) *core.Aura {
			aura := unit.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 16927},
				Label:    "Chilled (Frostguard)",
				Duration: time.Second * 5,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddMoveSpeedModifier(&aura.ActionID, 0.30)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.RemoveMoveSpeedModifier(&aura.ActionID)
				},
			})
			core.AtkSpeedReductionEffect(aura, 1.25)
			return aura
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Frostguard",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMeleeMH,
			SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
			PPM:               1, // Estimated based on data from WoW Armaments Discord
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				debuffAuras.Get(result.Target).Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=13983/gravestone-war-axe
	// Chance on hit: Diseases target enemy for 55 Nature damage every 3 sec for 15 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(GravestoneWarAxe, "Gravestone War Axe", 0.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18289},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagDisease | core.SpellFlagPureDot,

			Dot: core.DotConfig{
				NumberOfTicks: 15,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Creeping Mold",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 55, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19962/grileks-carver
	// +117 Attack Power when fighting Dragonkin.
	core.NewItemEffect(GrileksCarver, func(agent core.Agent) {
		character := agent.GetCharacter()
		if character.CurrentTarget.MobType == proto.MobType_MobTypeDragonkin {
			character.PseudoStats.MobTypeAttackPower += 117
		}
	})

	// https://www.wowhead.com/classic/item=19961/grileks-grinder
	// +48 Attack Power when fighting Dragonkin.
	core.NewItemEffect(GrileksGrinder, func(agent core.Agent) {
		character := agent.GetCharacter()
		if character.CurrentTarget.MobType == proto.MobType_MobTypeDragonkin {
			character.PseudoStats.MobTypeAttackPower += 48
		}
	})

	// https://www.wowhead.com/classic/item=9651/gryphon-riders-stormhammer
	itemhelpers.CreateWeaponCoHProcDamage(GryphonRidersStormhammer, "Gryphon Rider's Stormhammer", 1.0, 18081, core.SpellSchoolNature, 91, 34, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=17071/gutgore-ripper
	// Chance on hit: Sends a shadowy bolt at the enemy causing 75 Shadow damage and lowering all stats by 25 for 30 sec.
	itemhelpers.CreateWeaponProcSpell(GutgoreRipper, "Gutgore Ripper", 1.0, func(character *core.Character) *core.Spell {
		procAuras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 461682},
				Label:    "Gutgore Ripper",
				Duration: time.Second * 30,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   -25,
						stats.Intellect: -25,
						stats.Stamina:   -25,
						stats.Spirit:    -25,
						stats.Strength:  -25,
					})
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					aura.Unit.AddStatsDynamic(sim, stats.Stats{
						stats.Agility:   25,
						stats.Intellect: 25,
						stats.Stamina:   25,
						stats.Spirit:    25,
						stats.Strength:  25,
					})
				},
			})
		})

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 461682},
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 75, spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					procAuras.Get(target).Activate(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=5616/gutwrencher
	itemhelpers.CreateWeaponProcSpell(Gutwrencher, "Gutwrencher", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 16406},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeSpecialHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Rend (Gutwrencher)",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=2164/gut-ripper
	itemhelpers.CreateWeaponCoHProcDamage(GutRipper, "Gut Ripper", 1.0, 18107, core.SpellSchoolPhysical, 95, 26, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=19874/halberd-of-smiting
	// Equip: Chance to decapitate the target on a melee swing, causing 452 to 676 damage.
	itemhelpers.CreateWeaponEquipProcDamage(HalberdOfSmiting, "Halberd of Smiting", 2.1, 467819, core.SpellSchoolPhysical, 452, 224, 0.0, core.DefenseTypeMelee) // Works as phantom strike

	// https://www.wowhead.com/classic/item=810/hammer-of-the-northern-wind
	itemhelpers.CreateWeaponCoHProcDamage(HammerOfTheNorthernWind, "Hammer of the Northern Wind", 3.5, 13439, core.SpellSchoolFrost, 20, 10, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=2243/hand-of-edward-the-odd
	// Chance on hit: Next spell cast within 4 sec will cast instantly.
	itemhelpers.CreateWeaponProcAura(HandOfEdwardTheOdd, "Hand of Edward the Odd", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 18803},
			Label:    "Focus (Hand of Edward the Odd)",
			Duration: time.Second * 4,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyCastSpeed(100000)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyCastSpeed(1 / 100000.0)
			},
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				aura.Deactivate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=8190/hanzo-sword
	itemhelpers.CreateWeaponCoHProcDamage(HanzoSword, "Hanzo Sword", 1.0, 16405, core.SpellSchoolPhysical, 75, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=13937/headmasters-charge
	// Use: Gives 20 additional intellect to party members within 30 yards. (10 Min Cooldown)
	// Originally did not stack with Arcane Intellect, but is reported to stack in SoD
	/* core.NewItemEffect(HeadmastersCharge, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 18264}

		buffAura := character.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Headmaster's Charge",
			Duration: time.Minute * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.Intellect, 25)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatDynamic(sim, stats.Intellect, -25)
			},
		})
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 10,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	}) */

	// https://www.wowhead.com/classic/item=11635/hookfang-shanker
	itemhelpers.CreateWeaponProcSpell(HookfangShanker, "Hookfang Shanker", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 13526},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPoison | core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				Aura: core.Aura{
					Label: "Corrosive Poison",
					OnGain: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Armor: -50})
					},
					OnExpire: func(aura *core.Aura, sim *core.Simulation) {
						aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Armor: 50})
					},
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.Spell.CalcAndDealPeriodicDamage(sim, target, 7, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=11684/ironfoe
	// Chance on hit: Grants 2 extra attacks on your next swing.
	itemhelpers.CreateWeaponProcSpell(Ironfoe, "Ironfoe", 0.8, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 15494},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttackProc(sim, 2, core.ActionID{SpellID: 15494}, spell)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19918/jekliks-crusher
	// Chance on hit: Wounds the target for 200 to 220 damage.
	// Original proc rate 4.0 lowered to 1.5 in SoD phase 5
	itemhelpers.CreateWeaponCoHProcDamage(JekliksCrusher, "Jeklik's Crusher", 4.0, 467642, core.SpellSchoolPhysical, 200, 20, 0.0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=17054/joonhos-mercy
	itemhelpers.CreateWeaponCoHProcDamage(JoonhosMercy, "Joonho's Mercy", 1.0, 20883, core.SpellSchoolArcane, 70, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=21679/kalimdors-revenge
	itemhelpers.CreateWeaponCoHProcDamage(KalimdorsRevenge, "Kalimdor's Revenge", 1.25, 26415, core.SpellSchoolNature, 239, 38, 0, core.DefenseTypeMagic) // TODO Update PPM/scaling from PTR

	// https://www.wowhead.com/classic/item=11902/linkens-sword-of-mastery
	itemhelpers.CreateWeaponCoHProcDamage(LinkensSwordOfMastery, "Linken's Sword of Mastery", 1.0, 18089, core.SpellSchoolNature, 45, 30, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=11817/lord-generals-sword
	// Chance on hit: Increases attack power by 50 for 30 sec.
	// // TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcAura(LordGeneralsSword, "Lord General's Sword", 1.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 15602},
			Label:    "Lord General's Sword",
			Duration: time.Second * 30,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.AttackPower:       50,
					stats.RangedAttackPower: 50,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.AttackPower:       -50,
					stats.RangedAttackPower: -50,
				})
			},
		})
	})

	// https://www.wowhead.com/classic/item=9449/manual-crowd-pummeler

	core.NewItemEffect(ManualCrowdPummeler, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 13494}
		duration := time.Second * 30

		mcpAura := character.GetOrRegisterAura(core.Aura{
			Label:    "Haste",
			ActionID: actionID,
			Duration: duration,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1.5)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1.0/1.5)
			},
		})

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Second * 30,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				mcpAura.Activate(sim)
			},
		})
		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=12794/masterwork-stormhammer
	// Chance on hit: Blasts up to 3 targets for 105 to 145 Nature damage.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(MasterworkStormhammer, "Masterwork Stormhammer", 0.5, func(character *core.Character) *core.Spell {
		maxHits := int(min(3, character.Env.GetNumTargets()))
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 463946},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for numHits := 0; numHits < maxHits; numHits++ {
					spell.CalcAndDealDamage(sim, target, sim.Roll(105, 145), spell.OutcomeMagicHitAndCrit)
					target = character.Env.NextTargetUnit(target)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=1982/nightblade
	itemhelpers.CreateWeaponCoHProcDamage(Nightblade, "Nightblade", 1.0, 18211, core.SpellSchoolShadow, 125, 150, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=19169/nightfall
	core.NewItemEffect(Nightfall, func(agent core.Agent) {
		makeNightfallProc(agent.GetCharacter(), "Nightfall")
	})

	// https://www.wowhead.com/classic/item=9425/pendulum-of-doom
	itemhelpers.CreateWeaponCoHProcDamage(PendulumOfDoom, "Pendulum of Doom", 0.5, 10373, core.SpellSchoolPhysical, 250, 100, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=12709/pips-skinner
	core.NewItemEffect(PipsSkinner, func(agent core.Agent) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeBeast {
			character.PseudoStats.MobTypeAttackPower += 45
		}
	})

	// https://www.wowhead.com/classic/item=18816/perditions-blade
	// Chance on hit: Blasts a target for 40 to 56 Fire damage.
	itemhelpers.CreateWeaponCoHProcDamage(PerditionsBlade, "Perdition's Blade", 2.8, 23267, core.SpellSchoolFire, 40, 56, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=19963/pitchfork-of-madness
	// +117 Attack Power when fighting Demons.
	core.NewItemEffect(PitchforkOfMadness, func(agent core.Agent) {
		character := agent.GetCharacter()
		if character.CurrentTarget.MobType == proto.MobType_MobTypeDemon {
			character.PseudoStats.MobTypeAttackPower += 117
		}
	})

	// https://www.wowhead.com/classic/item=18348/quelserrar
	// Chance on hit: When active, grants the wielder 13 defense and 300 armor for 10 sec.
	// Proc rate estimated based on data from WoW Armaments Discord for the original item
	itemhelpers.CreateWeaponProcAura(QuelSerrar, "Quel'Serrar", 2.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 463105},
			Label:    "Sanctuary",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Defense:    13,
					stats.BonusArmor: 300,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Defense:    -13,
					stats.BonusArmor: -300,
				})
			},
		})
	})

	// https://www.wowhead.com/classic/item=7717/ravager
	itemhelpers.CreateWeaponProcAura(Ravager, "Ravager", 1.0, func(character *core.Character) *core.Aura {
		tickActionID := core.ActionID{SpellID: 9633}
		procActionID := core.ActionID{SpellID: 9632}
		//Used as part of a canceling ravager APL
		auraActionID := core.ActionID{SpellID: 433801}

		ravegerBladestormTickSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    tickActionID,
			SpellSchool: core.SpellSchoolPhysical,
			DefenseType: core.DefenseTypeMelee,
			ProcMask:    core.ProcMaskMeleeMHSpecial,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			BonusCoefficient: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				damage := 5.0 + spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower())
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, damage, spell.OutcomeMeleeSpecialHitAndCrit)
				}
			},
		})

		character.GetOrRegisterSpell(core.SpellConfig{
			SpellSchool: core.SpellSchoolPhysical,
			ActionID:    procActionID,
			ProcMask:    core.ProcMaskMeleeMHSpecial,
			Flags:       core.SpellFlagChanneled,
			Dot: core.DotConfig{
				IsAOE: true,
				Aura: core.Aura{
					Label: "Ravager Whirlwind",
				},
				NumberOfTicks:       3,
				TickLength:          time.Second * 3,
				AffectedByCastSpeed: false,
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					ravegerBladestormTickSpell.Cast(sim, target)
				},
			},
		})

		return character.GetOrRegisterAura(core.Aura{
			Label:    "Ravager Bladestorm",
			ActionID: auraActionID,
			Duration: time.Second * 9,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AutoAttacks.CancelAutoSwing(sim)
				dotSpell := character.GetSpell(procActionID)
				dotSpell.AOEDot().Apply(sim)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AutoAttacks.EnableAutoSwing(sim)
				dotSpell := character.GetSpell(procActionID)
				dotSpell.AOEDot().Cancel(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=13505/runeblade-of-baron-rivendare
	// Equip: Increases movement speed and life regeneration rate.
	// TODO: Movement speed not implemented
	core.NewItemEffect(RunebladeOfBaronRivendare, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 17625}
		healthMetrics := character.NewHealthMetrics(actionID)
		character.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Unholy Aura",
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				core.StartPeriodicAction(sim, core.PeriodicActionOptions{
					Period:   time.Second * 5,
					Priority: core.ActionPriorityAuto,
					OnAction: func(sim *core.Simulation) {
						character.GainHealth(sim, 20, healthMetrics)
					},
				})
			},
		})
	})

	// https://www.wowhead.com/classic/item=17752/satyrs-lash
	itemhelpers.CreateWeaponCoHProcDamage(SatyrsLash, "Satyr's Lash", 1.0, 18205, core.SpellSchoolShadow, 55, 30, 0, core.DefenseTypeMagic)

	// TODO Searing Needle adds an "Apply Aura: Mod Damage Done (Fire): 10" aura to the /target/, buffing it; not currently modelled
	itemhelpers.CreateWeaponCoHProcDamage(SearingNeedle, "Searing Needle", 1.0, 16454, core.SpellSchoolFire, 60, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=12969/seeping-willow
	// Chance on hit: Lowers all stats by 20 and deals 20 Nature damage every 3 sec to all enemies within an 8 yard radius of the caster for 30 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(SeepingWillow, "Seeping Willow", 0.5, func(character *core.Character) *core.Spell {
		stats := stats.Stats{
			stats.Agility:   20,
			stats.Intellect: 20,
			stats.Stamina:   20,
			stats.Spirit:    20,
			stats.Strength:  20,
		}
		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit) *core.Aura {
			return unit.GetOrRegisterAura(core.Aura{
				ActionID: core.ActionID{SpellID: 17196},
				Label:    "Seeping Willow",
				Duration: time.Second * 30,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					unit.AddStatsDynamic(sim, stats.Multiply(-1))
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					unit.AddStatsDynamic(sim, stats)
				},
			})
		})

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 17196},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPoison | core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Seeping Willow Poison",
				},
				NumberOfTicks: 10,
				TickLength:    time.Second * 3,
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 20, isRollover)
					debuffAuras.Get(target).Activate(sim)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					result := spell.CalcAndDealOutcome(sim, aoeTarget, spell.OutcomeMagicHit)
					if result.Landed() {
						spell.Dot(aoeTarget).Apply(sim)
					}
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=13035/serpent-slicer
	itemhelpers.CreateWeaponProcSpell(SerpentSlicer, "Serpent Slicer", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 17511},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			Flags:            core.SpellFlagPoison | core.SpellFlagPureDot,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Poison (Serpent Slicer)",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 8, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
		})
	})

	// https://www.wowhead.com/classic/item=2163/shadowblade
	itemhelpers.CreateWeaponCoHProcDamage(Shadowblade, "Shadowblade", 1.0, 18138, core.SpellSchoolShadow, 110, 30, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=17074/shadowstrike
	// Chance on hit: Steals 100 to 180 life from target enemy.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(Shadowstrike, "Shadowstrike", 2.2, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 461683}
		healthMetrics := character.NewHealthMetrics(actionID)
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			BonusCoefficient: 1.0,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(100, 180), spell.OutcomeMagicHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})
	})

	// https://www.wowhead.com/classic/item=754/shortsword-of-vengeance
	itemhelpers.CreateWeaponCoHProcDamage(ShortswordOfVengeance, "Shortsword of Vengeance", 1.0, 13519, core.SpellSchoolHoly, 30, 0, 0, core.DefenseTypeMagic)

	// https://www.wowhead.com/classic/item=13361/skullforge-reaver
	// Equip: Drains target for 2 Shadow damage every 1 sec and transfers it to the caster. Lasts for 30 sec.
	// Estimated based on data from WoW Armaments Discord
	itemhelpers.CreateWeaponProcSpell(SkullforgeReaver, "Skullforge Reaver", 1.7, func(character *core.Character) *core.Spell {
		procMask := character.GetProcMaskForItem(SkullforgeReaver)
		actionID := core.ActionID{SpellID: 17484}
		healthMetrics := character.NewHealthMetrics(actionID)
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolShadow,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    procMask,
			Flags:       core.SpellFlagPureDot,
			Dot: core.DotConfig{
				NumberOfTicks: 30,
				TickLength:    time.Second,
				Aura: core.Aura{
					Label: "Skullforge Brand",
				},
				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 2, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					result := dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
					character.GainHealth(sim, result.Damage, healthMetrics)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=17182/sulfuras-hand-of-ragnaros
	// Chance on hit: Hurls a fiery ball that causes 273 to 333 Fire damage and an additional 75 damage over 10 sec.
	// Equip: Deals 5 Fire damage to anyone who strikes you with a melee attack.
	core.NewItemEffect(SulfurasHandOfRagnaros, func(agent core.Agent) {
		character := agent.GetCharacter()

		immolationActionID := core.ActionID{SpellID: 21142}

		immolationSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    immolationActionID,
			SpellSchool: core.SpellSchoolFire,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 5, spell.OutcomeMagicHit)
			},
		})

		character.GetOrRegisterAura(core.Aura{
			ActionID: immolationActionID,
			Label:    "Immolation (Hand of Ragnaros)",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) {
					immolationSpell.Cast(sim, spell.Unit)
				}
			},
		})

		fireballSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 21162},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Fireball (Hand of Ragnaros)",
				},
				TickLength:    2 * time.Second,
				NumberOfTicks: 5,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 15, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(273, 333), spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Hand of Ragnaros Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      1, // Estimated based on data from WoW Armaments Discord
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				fireballSpell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=17193/sulfuron-hammer
	// Chance on hit: Hurls a fiery ball that causes 83 to 101 Fire damage and an additional 16 damage over 8 sec.
	core.NewItemEffect(SulfuronHammer, func(agent core.Agent) {
		character := agent.GetCharacter()

		fireballSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 21159},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Fireball (Sulfuron Hammer)",
				},
				TickLength:    2 * time.Second,
				NumberOfTicks: 4,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 4, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(83, 101), spell.OutcomeMagicHitAndCrit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Sulfuron Hammer Trigger",
			Callback: core.CallbackOnSpellHitDealt,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee,
			PPM:      1, // TODO: Armaments Discord didn't have any data on Sulfuron Hammer
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				fireballSpell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=9639/the-hand-of-antusul
	itemhelpers.CreateWeaponProcSpell(TheHandOfAntusul, "The Hand of Antu'sul", 1.0, func(character *core.Character) *core.Spell {
		debuffAuras := character.NewEnemyAuraArray(func(unit *core.Unit) *core.Aura {
			aura := unit.GetOrRegisterAura(core.Aura{
				Label:    "ThunderClap-Antu'sul",
				ActionID: core.ActionID{SpellID: 13532},
				Duration: time.Second * 10,
			})
			core.AtkSpeedReductionEffect(aura, 1.11)
			return aura
		})

		results := make([]*core.SpellResult, min(4, character.Env.GetNumTargets()))

		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 13532},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for idx := range results {
					results[idx] = spell.CalcDamage(sim, target, 7, spell.OutcomeMagicHitAndCrit)
					target = character.Env.NextTargetUnit(target)
				}
				for _, result := range results {
					spell.DealDamage(sim, result)
					if result.Landed() {
						debuffAuras.Get(result.Target).Activate(sim)
					}
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=9423/the-jackhammer
	itemhelpers.CreateWeaponProcAura(TheJackhammer, "The Jackhammer", 1.0, func(character *core.Character) *core.Aura {
		return character.GetOrRegisterAura(core.Aura{
			Label:    "The Jackhammer Haste Aura",
			ActionID: core.ActionID{SpellID: 13533},
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1.3)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.MultiplyAttackSpeed(sim, 1/1.3)
			},
		})
	})

	// https://www.wowhead.com/classic/item=13060/the-needler
	itemhelpers.CreateWeaponCoHProcDamage(TheNeedler, "The Needler", 3.0, 13060, core.SpellSchoolPhysical, 75, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=19334/the-untamed-blade
	// Chance on hit: Increases Strength by 300 for 8 sec.
	// Estimated based on data from WoW Armaments Discord
	// Original proc rate 1.0 lowered to approximately 0.55 in SoD phase 5
	itemhelpers.CreateWeaponProcAura(TheUntamedBlade, "The Untamed Blade", 1.0, func(character *core.Character) *core.Aura {
		return character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 23719},
			Label:    "Untamed Fury",
			Duration: time.Second * 8,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Strength: 300})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.AddStatsDynamic(sim, stats.Stats{stats.Strength: -300})
			},
		})
	})

	// https://www.wowhead.com/classic/item=17705/thrash-blade
	itemhelpers.CreateWeaponProcSpell(ThrashBlade, "Thrash Blade", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 21919},
			SpellSchool:      core.SpellSchoolPhysical,
			DefenseType:      core.DefenseTypeMelee,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AutoAttacks.ExtraMHAttackProc(sim, 1, core.ActionID{SpellID: 21919}, spell)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19019/thunderfury-blessed-blade-of-the-windseeker
	// Chance on hit: Blasts your enemy with lightning, dealing 300 Nature damage and then jumping to additional nearby enemies.
	// Each jump reduces that victim's Nature resistance by 25. Affects 5 targets.
	// Your primary target is also consumed by a cyclone, slowing its attack speed by 20% for 12 sec.
	core.NewItemEffect(Thunderfury, func(agent core.Agent) {
		character := agent.GetCharacter()

		procMask := character.GetProcMaskForItem(Thunderfury)
		ppmm := character.AutoAttacks.NewPPMManager(6.0, procMask)
		thunderfuryASAuras := character.NewEnemyAuraArray(core.ThunderfuryASAura)
		procActionID := core.ActionID{SpellID: 21992}

		singleTargetSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    procActionID.WithTag(1),
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskSpellProc | core.ProcMaskSpellDamageProc,
			Flags:       core.SpellFlagIgnoreAttackerModifiers,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			FlatThreatBonus:  126,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 300, spell.OutcomeMagicHitAndCrit)
				thunderfuryASAuras.Get(target).Activate(sim)
			},
		})

		debuffAuras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
			return target.GetOrRegisterAura(core.Aura{
				Label:    "Thunderfury",
				ActionID: procActionID,
				Duration: time.Second * 12,
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					target.AddStatDynamic(sim, stats.NatureResistance, -25)
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					target.AddStatDynamic(sim, stats.NatureResistance, 25)
				},
			})
		})

		results := make([]*core.SpellResult, min(5, character.Env.GetNumTargets()))

		bounceSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    procActionID.WithTag(2),
			SpellSchool: core.SpellSchoolNature,
			ProcMask:    core.ProcMaskEmpty,

			ThreatMultiplier: 1,
			FlatThreatBonus:  126,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for idx := range results {
					results[idx] = spell.CalcDamage(sim, target, 0, spell.OutcomeMagicHit)
					target = sim.Environment.NextTargetUnit(target)
				}
				for _, result := range results {
					if result.Landed() {
						debuffAuras[result.Target.Index].Activate(sim)
					}
					spell.DealDamage(sim, result)
				}
			},
		})

		core.MakePermanent(character.GetOrRegisterAura(core.Aura{
			Label: "Thunderfury Trigger",
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && ppmm.Proc(sim, spell.ProcMask, "Thunderfury") {
					singleTargetSpell.Cast(sim, result.Target)
					bounceSpell.Cast(sim, result.Target)
				}
			},
		}))
	})

	// https://www.wowhead.com/classic/item=17223/thunderstrike
	// Chance on hit: Blasts up to 3 targets for 150 to 250 Nature damage. Each target after the first takes less damage.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(Thunderstrike, "Thunderstrike", 1.5, func(character *core.Character) *core.Spell {
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 461686},
			SpellSchool:      core.SpellSchoolNature,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				initialResult := spell.CalcOutcome(sim, target, spell.OutcomeMagicHit)
				// Only the initial hit can be fully resisted according to a wowhead comment
				if initialResult.Landed() {
					damageMultiplier := 1.0
					for numHits := 0; numHits < 3; numHits++ {
						spell.CalcAndDealDamage(sim, target, sim.Roll(150, 250)*damageMultiplier, spell.OutcomeMagicCrit)
						numHits++
						target = character.Env.NextTargetUnit(target)
						// TODO: Couldn't find information on what the multiplier actually is
						damageMultiplier *= .65
					}
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=19946/tigules-harpoon
	// +60 Attack Power when fighting Beasts.
	core.NewItemEffect(TigulesHarpoon, func(agent core.Agent) {
		character := agent.GetCharacter()
		if character.CurrentTarget.MobType == proto.MobType_MobTypeBeast {
			character.PseudoStats.MobTypeAttackPower += 60
		}
	})

	// https://www.wowhead.com/classic/item=13183/venomspitter
	// Chance on hit: Poisons target for 7 Nature damage every 2 sec for 30 sec.
	// TODO: Proc rate assumed and needs testing
	itemhelpers.CreateWeaponProcSpell(Venomspitter, "Venomspitter", 1.0, func(character *core.Character) *core.Spell {
		return character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 18203},
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPoison | core.SpellFlagPureDot,
			Dot: core.DotConfig{
				Aura: core.Aura{
					Label: "Poison (Venomspitter)",
				},
				TickLength:    time.Second * 2,
				NumberOfTicks: 15,

				OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 7, isRollover)
				},

				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
				if result.Landed() {
					spell.Dot(target).Apply(sim)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=11603/vilerend-slicer
	itemhelpers.CreateWeaponCoHProcDamage(VilerendSlicer, "Vilerend Slicer", 1.0, 16405, core.SpellSchoolPhysical, 75, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=17075/viskag-the-bloodletter
	itemhelpers.CreateWeaponCoHProcDamage(ViskagTheBloodletter, "Vis'kag the Bloodletter", 0.6, 21140, core.SpellSchoolPhysical, 240, 0, 0, core.DefenseTypeMelee)

	// https://www.wowhead.com/classic/item=11920/wraith-scythe
	// Chance on hit: Steals 45 life from target enemy.
	itemhelpers.CreateWeaponProcSpell(WraithScythe, "Wraith Scythe", 1.0, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 16414}
		healthMetrics := character.NewHealthMetrics(actionID)

		return character.RegisterSpell(core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			BonusCoefficient: 0.3,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 45, spell.OutcomeAlwaysHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19901/zulian-slicer
	// Chance on hit: Slices the enemy for 72 to 96 Nature damage.
	itemhelpers.CreateWeaponCoHProcDamage(ZulianSlicer, "Zulian Slicer", 1.2, 467738, core.SpellSchoolNature, 72, 24, 0.35, core.DefenseTypeMelee)

	///////////////////////////////////////////////////////////////////////////
	//                                 Trinkets
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/item=11832/burst-of-knowledge
	// Use: Reduces mana cost of all spells by 100 for 10 sec. (5 Min Cooldown)
	core.NewItemEffect(BurstOfKnowledge, func(agent core.Agent) {
		character := agent.GetCharacter()

		aura := character.GetOrRegisterAura(core.Aura{
			ActionID: core.ActionID{ItemID: BurstOfKnowledge},
			Label:    "Burst of Knowledge",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range aura.Unit.Spellbook {
					if spell.Cost != nil && spell.Cost.CostType() == core.CostTypeMana {
						spell.Cost.FlatModifier -= 100
					}
				}
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range aura.Unit.Spellbook {
					if spell.Cost != nil && spell.Cost.CostType() == core.CostTypeMana {
						spell.Cost.FlatModifier += 100
					}
				}
			},
		})

		spell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID: core.ActionID{ItemID: BurstOfKnowledge},
			ProcMask: core.ProcMaskEmpty,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 15,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				aura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeMana,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=22268/draconic-infused-emblem
	// Use: Increases your spell damage by up to 100 and your healing by up to 190 for 15 sec. (1 Min, 30 Sec Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(DraconicInfusedEmblem, stats.Stats{stats.SpellDamage: 100, stats.HealingPower: 190}, time.Second*15, time.Second*90)

	// https://www.wowhead.com/classic/item=19288/darkmoon-card-blue-dragon
	// Equip: 2% chance on successful spellcast to allow 100% of your Mana regeneration to continue while casting for 15 sec.
	core.NewItemEffect(DarkmoonCardBlueDragon, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 23688}

		procAura := character.GetOrRegisterAura(core.Aura{
			Label:    "Aura of the Blue Dragon",
			ActionID: actionID,
			Duration: time.Second * 15,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.SpiritRegenRateCasting += 1
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.PseudoStats.SpiritRegenRateCasting -= 1
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Aura of the Blue Dragon Trigger",
			Callback:   core.CallbackOnCastComplete,
			ProcMask:   core.ProcMaskSpellDamage | core.ProcMaskSpellHealing,
			ProcChance: .02,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procAura.Activate(sim)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19287/darkmoon-card-heroism
	// Equip: Sometimes heals bearer of 120 to 180 damage when damaging an enemy in melee.
	core.NewItemEffect(DarkmoonCardHeroism, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 23689}
		healthMetrics := character.NewHealthMetrics(actionID)

		procSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolHoly,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.GainHealth(sim, sim.Roll(120, 180), healthMetrics)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Heroism Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			PPM:               2,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, spell.Unit)
			},
		})
	})

	// https://www.wowhead.com/classic/item=19289/darkmoon-card-maelstrom
	// Equip: Chance to strike your melee target with lightning for 200 to 300 Nature damage.
	core.NewItemEffect(DarkmoonCardMaelstrom, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{SpellID: 23686}

		procSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolNature,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(200, 300), spell.OutcomeMagicHitAndCrit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Lightning Strike Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			PPM:               1.0,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=21180/earthstrike
	// Use: Increases your melee and ranged attack power by 280.  Effect lasts for 20 sec. (2 Min Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(Earthstrike, stats.Stats{stats.AttackPower: 280, stats.RangedAttackPower: 280}, time.Second*20, time.Second*120)

	// https://www.wowhead.com/classic/item=18815/essence-of-the-pure-flame
	// Equip: When struck in combat inflicts 13 Fire damage to the attacker.
	core.NewItemEffect(EssenceOfThePureFlame, func(agent core.Agent) {
		character := agent.GetCharacter()

		procSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 23266},
			SpellSchool: core.SpellSchoolFire,
			DefenseType: core.DefenseTypeMagic,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, 13, spell.OutcomeAlwaysHit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Fiery Aura Proc",
			Callback: core.CallbackOnSpellHitTaken,
			Outcome:  core.OutcomeLanded,
			ProcMask: core.ProcMaskMelee, // TODO: Unsure if this means melee attacks or all attacks
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				procSpell.Cast(sim, spell.Unit)
			},
		})
	})

	// https://www.wowhead.com/classic/item=21473/eye-of-moam
	// Use: Increases damage done by magical spells and effects by up to 50, and decreases the magical resistances of your spell targets by 100 for 30 sec. (3 Min Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(EyeOfMoam, stats.Stats{stats.SpellDamage: 50, stats.SpellPenetration: 100}, time.Second*30, time.Minute*3)

	core.NewItemEffect(HandOfJustice, func(agent core.Agent) {
		character := agent.GetCharacter()
		if !character.AutoAttacks.AutoSwingMelee {
			return
		}

		icd := core.Cooldown{
			Timer:    character.NewTimer(),
			Duration: time.Second * 2,
		}

		character.GetOrRegisterAura(core.Aura{
			Label:    "Hand of Justice",
			Duration: core.NeverExpires,
			OnReset: func(aura *core.Aura, sim *core.Simulation) {
				aura.Activate(sim)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if spell.Flags.Matches(core.SpellFlagSuppressEquipProcs) {
					return
				}
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) && icd.IsReady(sim) && sim.Proc(0.02, "HandOfJustice") {
					icd.Use(sim)
					aura.Unit.AutoAttacks.ExtraMHAttackProc(sim, 1, core.ActionID{SpellID: 15600}, spell)
				}
			},
		})
	})

	// https://www.wowhead.com/classic/item=22321/heart-of-wyrmthalak
	// Equip: Chance to bathe your melee target in flames for 120 to 180 Fire damage.
	// TODO: Proc rate assumed from a wowhead comment and needs testing
	core.NewItemEffect(HeartOfWyrmthalak, func(agent core.Agent) {
		character := agent.GetCharacter()
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 27656},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(120, 180), spell.OutcomeMagicHitAndCrit)
			},
		})
		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:              "Heart of Wyrmthalak Trigger",
			Callback:          core.CallbackOnSpellHitDealt,
			Outcome:           core.OutcomeLanded,
			ProcMask:          core.ProcMaskMelee,
			SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
			PPM:               0.4,
			Handler: func(sim *core.Simulation, _ *core.Spell, result *core.SpellResult) {
				spell.Cast(sim, result.Target)
			},
		})
	})

	// https://www.wowhead.com/classic/item=23570/jom-gabbar
	// Use: Increases attack power by 65 and an additional 65 every 2 sec.  Lasts 20 sec. (2 Min Cooldown)
	core.NewItemEffect(JomGabbar, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 29602}
		duration := time.Second * 20
		bonusPerStack := stats.Stats{
			stats.AttackPower:       65,
			stats.RangedAttackPower: 65,
		}

		jomGabbarAura := character.GetOrRegisterAura(core.Aura{
			Label:     "Jom Gabbar",
			ActionID:  actionID,
			Duration:  duration,
			MaxStacks: 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				core.StartPeriodicAction(sim, core.PeriodicActionOptions{
					Period:          time.Second * 2,
					NumTicks:        10,
					Priority:        core.ActionPriorityAuto,
					TickImmediately: true,
					OnAction: func(sim *core.Simulation) {
						aura.AddStack(sim)
					},
				})
			},
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				bonusStats := bonusPerStack.Multiply(float64(newStacks - oldStacks))
				character.AddStatsDynamic(sim, bonusStats)
			},
		})
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
				SharedCD: core.Cooldown{
					Timer:    character.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				jomGabbarAura.Activate(sim)
			},
		})
		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	/* core.NewItemEffect(MarkOfTheChampionPhys, func(agent core.Agent) {
		character := agent.GetCharacter()
		if character.CurrentTarget.MobType == proto.MobType_MobTypeUndead || character.CurrentTarget.MobType == proto.MobType_MobTypeDemon {
			character.PseudoStats.MobTypeAttackPower += 150
		}
	})

	core.NewItemEffect(MarkOfTheChampionSpell, func(agent core.Agent) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeUndead || character.CurrentTarget.MobType == proto.MobType_MobTypeDemon {
			character.PseudoStats.MobTypeSpellPower += 85
		}
	}) */

	// https://www.wowhead.com/classic/item=17774/mark-of-the-chosen
	core.NewItemEffect(MarkOfTheChosen, func(agent core.Agent) {
		character := agent.GetCharacter()
		statIncrease := float64(25)
		markProcChance := 0.02

		procAura := character.GetOrRegisterAura(core.Aura{
			Label:    "Mark of the Chosen Effect",
			ActionID: core.ActionID{SpellID: 21970},
			Duration: time.Minute,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:   statIncrease,
					stats.Agility:   statIncrease,
					stats.Strength:  statIncrease,
					stats.Intellect: statIncrease,
					stats.Spirit:    statIncrease,
				})
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatsDynamic(sim, stats.Stats{
					stats.Stamina:   -statIncrease,
					stats.Agility:   -statIncrease,
					stats.Strength:  -statIncrease,
					stats.Intellect: -statIncrease,
					stats.Spirit:    -statIncrease,
				})
			},
		})

		core.MakePermanent(character.GetOrRegisterAura(core.Aura{
			Label: "Mark of the Chosen",
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMelee) && sim.RandomFloat("Mark of the Chosen") < markProcChance {
					procAura.Activate(sim)
				}
			},
		}))
	})

	// https://www.wowhead.com/classic/item=19947/nat-pagles-broken-reel
	core.NewSimpleStatOffensiveTrinketEffect(NatPaglesBrokenReel, stats.Stats{
		stats.SpellHit: 10 * core.SpellHitRatingPerHitChance,
		stats.MeleeHit: 10 * core.MeleeHitRatingPerHitChance,
	}, time.Second*15, time.Second*90)

	// https://www.wowhead.com/classic/item=19812/rune-of-the-dawn
	// Equip: Increases damage done to Undead by magical spells and effects by up to 48.
	core.NewItemEffect(RuneOfTheDawn, func(agent core.Agent) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeUndead {
			character.AddStat(stats.SpellDamage, 48)
		}
	})

	// https://www.wowhead.com/classic/item=21625/scarab-brooch
	core.NewItemEffect(ScarabBrooch, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{ItemID: ScarabBrooch}

		shieldSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 26470},
			SpellSchool: core.SpellSchoolNature,
			ProcMask:    core.ProcMaskSpellHealing,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | core.SpellFlagHelpful,

			DamageMultiplier: 1,
			ThreatMultiplier: 1,

			Shield: core.ShieldConfig{
				Aura: core.Aura{
					Label:    "Scarab Brooch Shield",
					Duration: time.Second * 30,
				},
			},
		})

		activeAura := core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:     "Persistent Shield",
			Callback: core.CallbackOnHealDealt,
			Duration: time.Second * 30,
			Handler: func(sim *core.Simulation, _ *core.Spell, result *core.SpellResult) {
				shieldSpell.Shield(result.Target).Apply(sim, result.Damage*0.15)
			},
		})

		spell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolPhysical,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 3,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				activeAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=11819/second-wind
	// Use: Restores 30 mana every 1 sec for 10 sec. (15 Min Cooldown)
	core.NewItemEffect(SecondWind, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{SpellID: 15604}
		manaMetrics := character.NewManaMetrics(actionID)
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			ProcMask: core.ProcMaskEmpty,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 15,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				core.StartPeriodicAction(sim, core.PeriodicActionOptions{
					Period:   time.Second * 1,
					NumTicks: 10,
					Priority: core.ActionPriorityAuto,
					OnAction: func(sim *core.Simulation) {
						character.AddMana(sim, 30, manaMetrics)
					},
				})
			},
		})
		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=13213/smolderwebs-eye
	// Use: Poisons target for 20 Nature damage every 2 sec for 20 sec. (2 Min Cooldown)
	core.NewItemEffect(SmolderwebsEye, func(agent core.Agent) {
		character := agent.GetCharacter()
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{SpellID: 17330},
			SpellSchool: core.SpellSchoolNature,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagPoison | core.SpellFlagPureDot | core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
			},
			Dot: core.DotConfig{
				NumberOfTicks: 10,
				TickLength:    time.Second * 2,
				Aura: core.Aura{
					Label: "Poison (Smolderweb's Eye)",
				},
				OnSnapshot: func(_ *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
					dot.Snapshot(target, 20, isRollover)
				},
				OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
					dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
				},
			},
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.Dot(target).Apply(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=13209/seal-of-the-dawn
	// Equip: +81 Attack Power when fighting Undead.
	core.NewItemEffect(SealOfTheDawn, func(agent core.Agent) {
		character := agent.GetCharacter()

		if character.CurrentTarget.MobType == proto.MobType_MobTypeUndead {
			character.AddStat(stats.AttackPower, 81)
			character.AddStat(stats.AttackPower, 81)
		}
	})

	// https://www.wowhead.com/classic/item=18820/talisman-of-ephemeral-power
	// Use: Increases damage and healing done by magical spells and effects by up to 175 for 15 sec. (1 Min, 30 Sec Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(TalismanOfEphemeralPower, stats.Stats{stats.SpellPower: 175}, time.Second*15, time.Second*90)

	// https://www.wowhead.com/classic/item=19948/zandalarian-hero-badge
	// Increases your armor by 2000 and defense skill by 30 for 20 sec.
	// Every time you take melee or ranged damage, this bonus is reduced by 200 armor and 3 defense.
	core.NewItemEffect(ZandalariHeroBadge, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{ItemID: ZandalariHeroBadge}
		duration := time.Second * 20
		bonusPerStack := stats.Stats{
			stats.Armor:   200,
			stats.Defense: 3,
		}

		buffAura := character.GetOrRegisterAura(core.Aura{
			Label:     "Fragile Armor",
			ActionID:  actionID,
			Duration:  duration,
			MaxStacks: 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.SetStacks(sim, aura.MaxStacks)
			},
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				bonusStats := bonusPerStack.Multiply(float64(newStacks - oldStacks))
				character.AddStatsDynamic(sim, bonusStats)
			},
			OnSpellHitTaken: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if spell.ProcMask.Matches(core.ProcMaskMeleeOrRanged) && result.Landed() {
					aura.RemoveStack(sim)
				}
			},
		})

		cdSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Spell: cdSpell,
			Type:  core.CooldownTypeSurvival,
		})
	})

	// https://www.wowhead.com/classic/item=19950/zandalarian-hero-charm
	// Increases your spell damage by up to 204 and your healing by up to 408 for 20 sec.
	// Every time you cast a spell, the bonus is reduced by 17 spell damage and 34 healing.
	core.NewItemEffect(ZandalariHeroCharm, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{ItemID: ZandalariHeroCharm}
		duration := time.Second * 20
		bonusPerStack := stats.Stats{
			stats.SpellDamage:  17,
			stats.HealingPower: 34,
		}

		buffAura := character.GetOrRegisterAura(core.Aura{
			ActionID:  actionID,
			Label:     "Unstable Power",
			Duration:  duration,
			MaxStacks: 12,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.SetStacks(sim, aura.MaxStacks)
			},
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				bonusStats := bonusPerStack.Multiply(float64(newStacks - oldStacks))
				character.AddStatsDynamic(sim, bonusStats)
			},
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				if !spell.ProcMask.Matches(core.ProcMaskSpellDamage) {
					return
				}
				aura.RemoveStack(sim)
			},
		})

		cdSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
				SharedCD: core.Cooldown{
					Timer:    character.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Spell: cdSpell,
			Type:  core.CooldownTypeDPS,
		})
	})

	// https://www.wowhead.com/classic/item=19949/zandalarian-hero-medallion
	core.NewItemEffect(ZandalariHeroMedallion, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{ItemID: ZandalariHeroMedallion}
		duration := time.Second * 20

		buffAura := character.GetOrRegisterAura(core.Aura{
			ActionID:  actionID,
			Label:     "Restless Strength",
			Duration:  duration,
			MaxStacks: 20,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.SetStacks(sim, aura.MaxStacks)
			},
			OnStacksChange: func(aura *core.Aura, sim *core.Simulation, oldStacks, newStacks int32) {
				character.PseudoStats.BonusPhysicalDamage += 2 * float64(newStacks-oldStacks)
			},
			OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				if result.Landed() && spell.ProcMask.Matches(core.ProcMaskMeleeOrRanged) {
					aura.RemoveStack(sim)
				}
			},
		})

		cdSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 2,
				},
				SharedCD: core.Cooldown{
					Timer:    character.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Spell: cdSpell,
			Type:  core.CooldownTypeDPS,
		})
	})

	///////////////////////////////////////////////////////////////////////////
	//                                 Other
	///////////////////////////////////////////////////////////////////////////

	// https://www.wowhead.com/classic/item=17111/blazefury-medallion
	// Equip: Adds 2 fire damage to your melee attacks.
	core.NewItemEffect(BlazefuryMedallion, func(agent core.Agent) {
		character := agent.GetCharacter()
		BlazefuryTriggerAura(character, 7712, core.SpellSchoolFire, 2)
	})

	// https://www.wowhead.com/classic/item=1168/skullflame-shield
	// Equip: When struck in combat has a 3% chance of stealing 35 life from target enemy. (Proc chance: 3%)
	// Equip: When struck in combat has a 1% chance of dealing 75 to 125 Fire damage to all targets around you. (Proc chance: 1%)
	core.NewItemEffect(SkullflameShield, func(agent core.Agent) {
		character := agent.GetCharacter()

		drainLifeActionID := core.ActionID{SpellID: 18817}
		healthMetrics := character.NewHealthMetrics(drainLifeActionID)
		drainLifeSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:         drainLifeActionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, 35, spell.OutcomeAlwaysHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})

		flamestrikeSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:         core.ActionID{SpellID: 18818},
			SpellSchool:      core.SpellSchoolFire,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskEmpty,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				for _, aoeTarget := range sim.Encounter.TargetUnits {
					spell.CalcAndDealDamage(sim, aoeTarget, sim.Roll(75, 125), spell.OutcomeMagicHit)
				}
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Drain Life Trigger",
			Callback:   core.CallbackOnSpellHitTaken,
			Outcome:    core.OutcomeLanded,
			ProcMask:   core.ProcMaskMelee,
			ProcChance: 0.03,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				drainLifeSpell.Cast(sim, spell.Unit)
			},
		})

		core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
			Name:       "Flamestrike Trigger",
			Callback:   core.CallbackOnSpellHitTaken,
			Outcome:    core.OutcomeLanded,
			ProcMask:   core.ProcMaskMelee,
			ProcChance: 0.01,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				flamestrikeSpell.Cast(sim, spell.Unit)
			},
		})
	})

	// https://www.wowhead.com/classic/item=21190/wrath-of-cenarius
	// Gives a chance when your harmful spells land to increase the damage of your spells and effects by 132 for 10 sec.
	// (Proc chance: 5%)
	core.NewItemEffect(WrathOfCenarius, func(agent core.Agent) {
		character := agent.GetCharacter()
		buffAura := character.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 25906},
			Label:    "Spell Blasting",
			Duration: time.Second * 10,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatDynamic(sim, stats.SpellDamage, 132)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatDynamic(sim, stats.SpellDamage, -132)
			},
		})

		core.MakeProcTriggerAura(&agent.GetCharacter().Unit, core.ProcTrigger{
			Name:       "Spell Blasting Trigger",
			Callback:   core.CallbackOnSpellHitDealt,
			Outcome:    core.OutcomeLanded,
			ProcMask:   core.ProcMaskSpellDamage,
			ProcChance: 0.05,
			Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
				buffAura.Activate(sim)
			},
		})
	})

	core.AddEffectsToTest = true
}

func BlazefuryTriggerAura(character *core.Character, spellID int32, spellSchool core.SpellSchool, damage float64) {
	if character.GetSpell(core.ActionID{SpellID: spellID}) != nil {
		return
	}

	procSpell := character.RegisterSpell(core.SpellConfig{
		ActionID:         core.ActionID{SpellID: spellID},
		SpellSchool:      spellSchool,
		DefenseType:      core.DefenseTypeMagic,
		ProcMask:         core.ProcMaskMeleeDamageProc,
		Flags:            core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,
		DamageMultiplier: 1,
		ThreatMultiplier: 1,
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMagicCrit)
		},
	})

	core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
		Name:              fmt.Sprintf("Blazefury Trigger (%d)", spellID),
		Callback:          core.CallbackOnSpellHitDealt,
		Outcome:           core.OutcomeLanded,
		ProcMask:          core.ProcMaskMelee,
		SpellFlagsExclude: core.SpellFlagSuppressEquipProcs,
		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if spell.ProcMask.Matches(core.ProcMaskMeleeSpecial) {
				procSpell.ProcMask = core.ProcMaskEmpty
			} else {
				procSpell.ProcMask = core.ProcMaskDamageProc // Both spell and melee procs
			}
			procSpell.Cast(sim, result.Target)
		},
	})
}

// Chance on hit: Spell damage taken by target increased by 15% for 5 sec.
func makeNightfallProc(character *core.Character, itemName string) {
	procAuras := character.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return target.GetOrRegisterAura(core.Aura{
			Label:    fmt.Sprintf("Spell Vulnerability (%s)", itemName),
			ActionID: core.ActionID{SpellID: 23605},
			Duration: time.Second * 5,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexArcane] *= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] *= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFrost] *= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] *= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexNature] *= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexShadow] *= 1.15
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexArcane] /= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFire] /= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexFrost] /= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexHoly] /= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexNature] /= 1.15
				aura.Unit.PseudoStats.SchoolBonusDamageTaken[stats.SchoolIndexShadow] /= 1.15
			},
		})
	})

	core.MakeProcTriggerAura(&character.Unit, core.ProcTrigger{
		Name:              fmt.Sprintf("%s Trigger", itemName),
		Callback:          core.CallbackOnSpellHitDealt,
		Outcome:           core.OutcomeLanded,
		ProcMask:          core.ProcMaskMelee,
		SpellFlagsExclude: core.SpellFlagSuppressWeaponProcs,
		PPM:               2,
		Handler: func(sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			procAuras.Get(result.Target).Activate(sim)
		},
	})
}

func StrengthOfTheChampionAura(character *core.Character) *core.Aura {
	actionID := core.ActionID{SpellID: 16916}
	healthMetrics := character.NewHealthMetrics(actionID)
	return character.GetOrRegisterAura(core.Aura{
		Label:    "Strength of the Champion",
		ActionID: actionID,
		Duration: time.Second * 30,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			character.GainHealth(sim, sim.Roll(270, 450), healthMetrics)
			character.AddStatDynamic(sim, stats.Strength, 120)
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			character.AddStatDynamic(sim, stats.Strength, -120)
		},
	})
}

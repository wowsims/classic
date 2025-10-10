package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

const RaptorStrikeRanks = 8

var RaptorStrikeSpellId = [RaptorStrikeRanks + 1]int32{0, 2973, 14260, 14261, 14262, 14263, 14264, 14265, 14266}
var RaptorStrikeSpellIdMeleeSpecialist = [RaptorStrikeRanks + 1]int32{0, 415335, 415336, 415337, 415338, 415340, 415341, 415342, 415343}
var RaptorStrikeBaseDamage = [RaptorStrikeRanks + 1]float64{0, 5, 11, 21, 34, 50, 80, 110, 140}
var RaptorStrikeManaCost = [RaptorStrikeRanks + 1]float64{0, 15, 25, 35, 45, 55, 70, 80, 100}
var RaptorStrikeLevel = [RaptorStrikeRanks + 1]int{0, 1, 8, 16, 24, 32, 40, 48, 56}

// Returns true if the regular melee swing should be used, false otherwise.
func (hunter *Hunter) TryRaptorStrike(sim *core.Simulation, mhSwingSpell *core.Spell) *core.Spell {
	if hunter.curQueuedAutoSpell != nil && hunter.curQueuedAutoSpell.CanCast(sim, hunter.CurrentTarget) {
		return hunter.curQueuedAutoSpell
	}
	return mhSwingSpell
}

func (hunter *Hunter) getRaptorStrikeConfig(rank int) core.SpellConfig {
	spellID := RaptorStrikeSpellId[rank]
	manaCost := RaptorStrikeManaCost[rank]
	level := RaptorStrikeLevel[rank]

	hunter.RaptorStrikeHit = hunter.newRaptorStrikeHitSpell(rank)

	spellConfig := core.SpellConfig{
		SpellCode:     SpellCode_HunterRaptorStrike,
		ActionID:      core.ActionID{SpellID: spellID},
		SpellSchool:   core.SpellSchoolPhysical,
		DefenseType:   core.DefenseTypeMelee,
		ProcMask:      core.ProcMaskMeleeMHSpecial | core.ProcMaskMeleeMHAuto,
		Flags:         core.SpellFlagMeleeMetrics | SpellFlagStrike,
		Rank:          rank,
		RequiredLevel: level,

		ManaCost: core.ManaCostOptions{
			FlatCost: manaCost,
		},

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: time.Second * 6,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.DistanceFromTarget <= core.MaxMeleeAttackDistance
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			hunter.RaptorStrikeHit.Cast(sim, target)

			if hunter.curQueueAura != nil {
				hunter.curQueueAura.Deactivate(sim)
			}
		},
	}

	return spellConfig
}

func (hunter *Hunter) newRaptorStrikeHitSpell(rank int) *core.Spell {
	spellID := RaptorStrikeSpellId[rank]
	baseDamage := RaptorStrikeBaseDamage[rank]

	return hunter.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_HunterRaptorStrikeHit,
		ActionID:    core.ActionID{SpellID: spellID}.WithTag(1),
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagNoOnCastComplete,

		BonusCritRating:  float64(hunter.Talents.SavageStrikes) * 10 * core.CritRatingPerCritChance,
		CritDamageBonus:  hunter.mortalShots(),
		DamageMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			damage := baseDamage + hunter.MHWeaponDamage(sim, spell.MeleeAttackPower(target))
			spell.CalcAndDealDamage(sim, target, damage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)
		},
	})
}

func (hunter *Hunter) makeQueueSpellsAndAura() *core.Spell {
	queueAura := hunter.RegisterAura(core.Aura{
		Label:    "Raptor Strike Queued",
		ActionID: hunter.RaptorStrike.ActionID,
		Duration: core.NeverExpires,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			if hunter.curQueueAura != nil {
				hunter.curQueueAura.Deactivate(sim)
			}
			hunter.PseudoStats.DisableDWMissPenalty = true
			hunter.curQueueAura = aura
			hunter.curQueuedAutoSpell = hunter.RaptorStrike
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			hunter.PseudoStats.DisableDWMissPenalty = false
			hunter.curQueueAura = nil
			hunter.curQueuedAutoSpell = nil
		},
	})

	queueSpell := hunter.RegisterSpell(core.SpellConfig{
		SpellCode: SpellCode_HunterRaptorStrike,
		ActionID:  hunter.RaptorStrike.WithTag(3),
		Flags:     core.SpellFlagMeleeMetrics | core.SpellFlagAPL,

		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return hunter.curQueueAura != queueAura &&
				hunter.CurrentMana() >= hunter.RaptorStrike.Cost.GetCurrentCost() &&
				!hunter.IsCasting(sim) &&
				hunter.DistanceFromTarget <= core.MaxMeleeAttackDistance &&
				hunter.RaptorStrike.IsReady(sim)
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			queueAura.Activate(sim)
		},
	})
	queueSpell.CdSpell = hunter.RaptorStrike

	return queueSpell
}

func (hunter *Hunter) registerRaptorStrikeSpell() {
	rank := map[int32]int{
		25: 4,
		40: 6,
		50: 7,
		60: 8,
	}[hunter.Level]

	config := hunter.getRaptorStrikeConfig(rank)
	hunter.RaptorStrike = hunter.GetOrRegisterSpell(config)
	hunter.makeQueueSpellsAndAura()
}

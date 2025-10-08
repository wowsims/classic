package rogue

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (rogue *Rogue) registerEviscerate() {
	flatDamage := map[int32]float64{
		25: 10,
		40: 22,
		50: 34,
		60: core.TernaryFloat64(core.IncludeAQ, 54, 48),
	}[rogue.Level]

	comboDamageBonus := map[int32]float64{
		25: 31,
		40: 77,
		50: 110,
		60: core.TernaryFloat64(core.IncludeAQ, 170, 151),
	}[rogue.Level]

	damageVariance := map[int32]float64{
		25: 20,
		40: 44,
		50: 68,
		60: core.TernaryFloat64(core.IncludeAQ, 108, 96),
	}[rogue.Level]

	spellID := map[int32]int32{
		25: 6762,
		40: 8624,
		50: 11299,
		60: core.TernaryInt32(core.IncludeAQ, 31016, 11300),
	}[rogue.Level]

	rogue.Eviscerate = rogue.RegisterSpell(core.SpellConfig{
		SpellCode:    SpellCode_RogueEviscerate,
		ActionID:     core.ActionID{SpellID: spellID},
		SpellSchool:  core.SpellSchoolPhysical,
		DefenseType:  core.DefenseTypeMelee,
		ProcMask:     core.ProcMaskMeleeMHSpecial,
		Flags:        rogue.finisherFlags() | SpellFlagColdBlooded,
		MetricSplits: 6,

		EnergyCost: core.EnergyCostOptions{
			Cost:   35,
			Refund: 0,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: time.Second,
			},
			IgnoreHaste: true,
			ModifyCast: func(sim *core.Simulation, spell *core.Spell, cast *core.Cast) {
				spell.SetMetricsSplit(spell.Unit.ComboPoints())
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return rogue.ComboPoints() > 0
		},

		DamageMultiplier: 1 +
			[]float64{0, 0.05, 0.10, 0.15}[rogue.Talents.ImprovedEviscerate] +
			[]float64{0, 0.02, 0.04, 0.06}[rogue.Talents.Aggression],
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			rogue.BreakStealth(sim)

			comboPoints := rogue.ComboPoints()
			flatBaseDamage := flatDamage + comboDamageBonus*float64(comboPoints)

			baseDamage := sim.Roll(flatBaseDamage, flatBaseDamage+damageVariance) +
				0.03*float64(comboPoints)*spell.MeleeAttackPower(target)

			result := spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)

			if result.Landed() {
				rogue.SpendComboPoints(sim, spell)
			} else {
				spell.IssueRefund(sim)
			}

			spell.DealDamage(sim, result)
		},
	})
	rogue.Finishers = append(rogue.Finishers, rogue.Eviscerate)
}

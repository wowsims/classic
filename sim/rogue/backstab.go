package rogue

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (rogue *Rogue) registerBackstabSpell() {
	flatDamageBonus := map[int32]float64{
		25: 32,
		40: 60,
		50: 90,
		60: core.TernaryFloat64(rogue.Env.UseAQSpellRanks, 210, 225),
	}[rogue.Level]

	spellID := map[int32]int32{
		25: 2590,
		40: 8721,
		50: 11279,
		60: core.TernaryInt32(rogue.Env.UseAQSpellRanks, 25300, 11281),
	}[rogue.Level]

	damageMultiplier := 1.5 * []float64{1, 1.04, 1.08, 1.12, 1.16, 1.2}[rogue.Talents.Opportunity]

	rogue.Backstab = rogue.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_RogueBackstab,
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       rogue.builderFlags(),

		EnergyCost: core.EnergyCostOptions{
			Cost:   60,
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: time.Second,
			},
			IgnoreHaste: true,
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			if !rogue.HasDagger(core.MainHand) {
				return false
			}
			return !rogue.PseudoStats.InFrontOfTarget
		},

		BonusCritRating: 10 * core.CritRatingPerCritChance * float64(rogue.Talents.ImprovedBackstab),

		CritDamageBonus: rogue.lethality(),

		DamageMultiplier: damageMultiplier,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			rogue.BreakStealth(sim)
			baseDamage := (flatDamageBonus + spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower()))
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)

			if result.Landed() {
				rogue.AddComboPoints(sim, 1, target, spell.ComboPointMetrics())
			} else {
				spell.IssueRefund(sim)
			}
		},
	})
}

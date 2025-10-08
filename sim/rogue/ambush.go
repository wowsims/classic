package rogue

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (rogue *Rogue) registerAmbushSpell() {
	flatDamageBonus := map[int32]float64{
		25: 28,
		40: 50,
		50: 92,
		60: 116,
	}[rogue.Level]

	spellID := map[int32]int32{
		25: 8676,
		40: 8725,
		50: 11268,
		60: 11269,
	}[rogue.Level]

	damageMultiplier := 2.5 * []float64{1, 1.04, 1.08, 1.12, 1.16, 1.2}[rogue.Talents.Opportunity]

	rogue.Ambush = rogue.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_RogueAmbush,
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
			if rogue.IsStealthed() {
				return true
			}
			return !rogue.PseudoStats.InFrontOfTarget && rogue.IsStealthed()
		},

		BonusCritRating:  15 * core.CritRatingPerCritChance * float64(rogue.Talents.ImprovedAmbush),
		DamageMultiplier: damageMultiplier,
		ThreatMultiplier: 1,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			rogue.BreakStealth(sim)
			baseDamage := (flatDamageBonus + spell.Unit.MHNormalizedWeaponDamage(sim, spell.MeleeAttackPower(target)))

			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialNoBlockDodgeParry)

			if result.Landed() {
				rogue.AddComboPoints(sim, 1, target, spell.ComboPointMetrics())
			} else {
				spell.IssueRefund(sim)
			}
		},
	})
}

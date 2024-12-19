package warrior

import (
	"github.com/wowsims/classic/sim/core"
)

func (warrior *Warrior) registerExecuteSpell() {

	flatDamage := 600.0
	convertedRageDamage := 15.0
	spellID := int32(20662)

	var rageMetrics *core.ResourceMetrics
	warrior.Execute = warrior.RegisterSpell(BattleStance|BerserkerStance, core.SpellConfig{
		SpellCode:   SpellCode_WarriorExecute,
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL | core.SpellFlagPassiveSpell | SpellFlagOffensive,

		RageCost: core.RageCostOptions{
			Cost:   15 - []float64{0, 2, 5}[warrior.Talents.ImprovedExecute],
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return sim.IsExecutePhase20()
		},

		CritDamageBonus: warrior.impale(),

		DamageMultiplier: 1,
		ThreatMultiplier: 1.25,
		BonusCoefficient: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			extraRage := spell.Unit.CurrentRage()
			warrior.SpendRage(sim, extraRage, rageMetrics)
			// We must count this rage event if the spell itself cost 0,
			// otherwise we could end up with 0 events even though rage was spent.
			if spell.Cost.GetCurrentCost() > 0 {
				rageMetrics.Events--
			}

			baseDamage := flatDamage + convertedRageDamage*(extraRage)

			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)

			if !result.Landed() {
				spell.IssueRefund(sim)
			}
		},
	})
	rageMetrics = warrior.Execute.Cost.SpellCostFunctions.(*core.RageCost).ResourceMetrics
}

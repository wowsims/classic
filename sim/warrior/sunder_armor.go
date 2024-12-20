package warrior

import (
	"github.com/wowsims/classic/sim/core"
)

func (warrior *Warrior) registerSunderArmorSpell() *WarriorSpell {
	warrior.SunderArmorAuras = warrior.NewEnemyAuraArray(core.SunderArmorAura)

	spellID := int32(11597)

	spell_level := 58

	var canApplySunder bool


	return warrior.RegisterSpell(AnyStance, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagAPL | SpellFlagOffensive,

		RageCost: core.RageCostOptions{
			Cost:   15 - float64(warrior.Talents.ImprovedSunderArmor),
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			sa := warrior.SunderArmorAuras.Get(target)
			if sa.IsActive() {
				canApplySunder = true
			} else if sa.ExclusiveEffects[0].Category.AnyActive() {
				canApplySunder = false
			} else {
				canApplySunder = true
			}
			return canApplySunder
		},

		ThreatMultiplier: 1,
		FlatThreatBonus:  2.25 * 2 * float64(spell_level),

		RelatedAuras: []core.AuraArray{warrior.SunderArmorAuras},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMeleeWeaponSpecialNoCrit) // Cannot be blocked

			if !result.Landed() {
				spell.IssueRefund(sim)
				return
			}

			if canApplySunder {
				sa := warrior.SunderArmorAuras.Get(target)
				sa.Activate(sim)
				sa.AddStack(sim)
			}
		},
	})
}

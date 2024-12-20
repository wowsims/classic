package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (warrior *Warrior) registerThunderClapSpell() {
	spellID := int32(11581)
	baseDamage := 103.0
	duration := time.Second * 30
	has5pcConq := warrior.HasSetBonus(ItemSetConquerorsBattleGear, 5)
	attackSpeedReduction := core.TernaryInt32(has5pcConq, 15, 10)
	stanceMask := BattleStance

	warrior.ThunderClapAuras = warrior.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return core.ThunderClapAura(target, spellID, duration, attackSpeedReduction)
	})

	results := make([]*core.SpellResult, min(4, warrior.Env.GetNumTargets()))

	warrior.ThunderClap = warrior.RegisterSpell(stanceMask, core.SpellConfig{
		ActionID:    core.ActionID{SpellID: spellID},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellDamage,
		Flags:       core.SpellFlagAPL | SpellFlagOffensive,

		RageCost: core.RageCostOptions{
			Cost: 20 - []float64{0, 1, 2, 4}[warrior.Talents.ImprovedThunderClap],
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			IgnoreHaste: true,
			CD: core.Cooldown{
				Timer:    warrior.NewTimer(),
				Duration: time.Second * 4,
			},
		},

		CritDamageBonus: warrior.impale(),

		DamageMultiplier: core.TernaryFloat64(has5pcConq, 1.5, 1),
		ThreatMultiplier: 2.5,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			for idx := range results {
				results[idx] = spell.CalcDamage(sim, target, baseDamage, spell.OutcomeMagicHitAndCrit)
				target = sim.Environment.NextTargetUnit(target)
			}

			for _, result := range results {
				spell.DealDamage(sim, result)
				if result.Landed() {
					warrior.ThunderClapAuras.Get(result.Target).Activate(sim)
				}
			}
		},

		RelatedAuras: []core.AuraArray{warrior.ThunderClapAuras},
	})
}

package shaman

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

func (shaman *Shaman) registerStormstrikeSpell() {
	if !shaman.Talents.Stormstrike {
		return
	}

	stormStrikeAuras := shaman.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return core.StormstrikeAura(target)
	})

	shaman.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_ShamanStormstrike,
		ActionID:    core.ActionID{SpellID: 17364},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       SpellFlagShaman | core.SpellFlagAPL | core.SpellFlagMeleeMetrics,

		ManaCost: core.ManaCostOptions{
			BaseCost: .21,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: core.GCDDefault,
			},
			CD: core.Cooldown{
				Timer:    shaman.NewTimer(),
				Duration: time.Second * 20,
			},
		},

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := shaman.MHWeaponDamage(sim, spell.MeleeAttackPower(target))
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)

			if result.Landed() {
				stormStrikeAuras.Get(target).Activate(sim)
			}
		},
	})
}

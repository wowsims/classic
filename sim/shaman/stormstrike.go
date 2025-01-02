package shaman

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// TODO: All of this needs to be refactored for how it works in Vanilla vs SoD.
// Gives an extra attack instead of a true yellow hit
func (shaman *Shaman) registerStormstrikeSpell() {
	if !shaman.Talents.Stormstrike {
		return
	}

	shaman.StormstrikeMH = shaman.newStormstrikeHitSpell()
	shaman.StormstrikeMH.SpellCode = SpellCode_ShamanStormstrike

	shaman.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 17364},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       SpellFlagShaman | core.SpellFlagAPL | core.SpellFlagNoOnCastComplete,

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

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			shaman.StormstrikeMH.Cast(sim, target)
		},
	})
}

// Only the main-hand hit triggers procs / the debuff
func (shaman *Shaman) newStormstrikeHitSpell() *core.Spell {
	stormStrikeAuras := shaman.NewEnemyAuraArray(func(target *core.Unit) *core.Aura {
		return core.StormstrikeAura(target)
	})

	return shaman.RegisterSpell(core.SpellConfig{
		ActionID:    core.ActionID{SpellID: 17364}.WithTag(1),
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := shaman.MHWeaponDamage(sim, spell.MeleeAttackPower())
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeWeaponSpecialHitAndCrit)

			if result.Landed() {
				// TODO: Stormstrike with 2 stacks instead of unlimited
				stormStrikeAuras.Get(target).Activate(sim)
			}
		},
	})
}

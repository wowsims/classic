package mage

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// If two spells proc Ignite at almost exactly the same time, the latter
// overwrites the former.
const IgniteTicks = 2

func (mage *Mage) applyIgnite() {
	if mage.Talents.Ignite == 0 {
		return
	}
	newIgniteDamage := 0.0

	mage.RegisterAura(core.Aura{
		Label:    "Ignite Talent",
		Duration: core.NeverExpires,
		OnReset: func(aura *core.Aura, sim *core.Simulation) {
			aura.Activate(sim)
		},
		OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
			if !spell.ProcMask.Matches(core.ProcMaskSpellDamage) {
				return
			}
			if spell.SpellSchool.Matches(core.SpellSchoolFire) && result.DidCrit() {
				newIgniteDamage = result.Damage * 0.08 * float64(mage.Talents.Ignite)
				mage.Ignite.Cast(sim, result.Target)
			}
		},
	})

	mage.Ignite = mage.RegisterSpell(core.SpellConfig{
		ActionID: core.ActionID{SpellID: 12654},
		Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | SpellFlagMage,

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
//			result := spell.CalcAndDealOutcome(sim, target, spell.OutcomeMagicHit)
	
			dot := mage.igniteTick.Dot(target)
			dot.ApplyOrRefresh(sim)
			if dot.GetStacks() < dot.MaxStacks{
				dot.AddStack(sim)
				dot.TakeSnapshot(sim, true)
}
			
		},
	})

	mage.igniteTick = mage.RegisterSpell(core.SpellConfig{
		SpellCode:   SpellCode_MageIgnite,
		ActionID:    core.ActionID{SpellID: 12654},
		SpellSchool: core.SpellSchoolFire,
		DefenseType: core.DefenseTypeMagic,
		ProcMask:    core.ProcMaskSpellProc,
		Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell | SpellFlagMage,

		DamageMultiplier: 1,
		ThreatMultiplier: 1,

		Cast: core.CastConfig{
			IgnoreHaste: true,
		},

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "Ignite",
				MaxStacks: 5,
				Duration: time.Second * 4,
			},
			NumberOfTicks: IgniteTicks,
			TickLength:    time.Second * 2,

			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, applyStack bool) {
				if !applyStack {
					return
				}

				// only the first stack snapshots the multiplier
				if dot.GetStacks() == 1 {
					attackTable := dot.Spell.Unit.AttackTables[target.UnitIndex][dot.Spell.CastType]
					dot.SnapshotAttackerMultiplier = dot.Spell.AttackerDamageMultiplier(attackTable)
					dot.SnapshotBaseDamage = newIgniteDamage
				} else {
					dot.SnapshotBaseDamage += newIgniteDamage
				}	
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},
	})
}
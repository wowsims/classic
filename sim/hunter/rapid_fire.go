package hunter

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

func (hunter *Hunter) registerRapidFire() {
	if hunter.Level < 26 {
		return
	}

	hasRapidKilling := hunter.HasRune(proto.HunterRune_RuneHelmRapidKilling)

	actionID := core.ActionID{SpellID: 3045}
	cooldown := core.TernaryDuration(hasRapidKilling, time.Minute*1, time.Minute*5)

	hunter.RapidFireAura = hunter.RegisterAura(core.Aura{
		Label:    "Rapid Fire",
		ActionID: actionID,
		Duration: time.Second * 15,

		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.MultiplyRangedSpeed(sim, 1.4)
			if hasRapidKilling {
				aura.Unit.MultiplyMeleeSpeed(sim, 1.4)
			}
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			aura.Unit.MultiplyRangedSpeed(sim, 1/1.4)
			if hasRapidKilling {
				aura.Unit.MultiplyMeleeSpeed(sim, 1/1.4)
			}
		},
	})

	hunter.RapidFire = hunter.RegisterSpell(core.SpellConfig{
		ActionID: actionID,

		ManaCost: core.ManaCostOptions{
			FlatCost: 100,
		},
		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    hunter.NewTimer(),
				Duration: cooldown,
			},
			SharedCD: core.Cooldown{
				Timer:    hunter.GetAttackSpeedBuffCD(),
				Duration: cooldown,
			},
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, _ *core.Spell) {
			hunter.RapidFireAura.Activate(sim)
		},
	})

	hunter.AddMajorCooldown(core.MajorCooldown{
		Spell: hunter.RapidFire,
		Type:  core.CooldownTypeDPS,
	})
}

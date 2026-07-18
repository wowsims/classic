package warlock

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const (
	lifeTapHealingCadenceSeconds          = 3.0
	lifeTapHealingCadenceVariationSeconds = 1.0
)

func (warlock *Warlock) registerLifeTapHealingModel() {
	hps := warlock.Options.AssumedLifeTapHps
	if hps == 0 || warlock.IsTanking() {
		return
	}

	minCadence := max(0.0, lifeTapHealingCadenceSeconds-lifeTapHealingCadenceVariationSeconds)
	cadenceVariationLow := lifeTapHealingCadenceSeconds - minCadence

	healthMetrics := warlock.NewHealthMetrics(core.ActionID{OtherID: proto.OtherAction_OtherActionHealingModel})
	healingModelSpell := warlock.RegisterSpell(core.SpellConfig{
		ActionID: core.ActionID{OtherID: proto.OtherAction_OtherActionHealingModel},
	})

	rollHealingCadence := func(sim *core.Simulation) time.Duration {
		signRoll := sim.RandomFloat("Life Tap Healing Cadence Variation Sign")
		magnitudeRoll := sim.RandomFloat("Life Tap Healing Cadence Variation Magnitude")
		if signRoll < 0.5 {
			return core.DurationFromSeconds(minCadence + magnitudeRoll*cadenceVariationLow)
		}
		return core.DurationFromSeconds(lifeTapHealingCadenceSeconds + magnitudeRoll*lifeTapHealingCadenceVariationSeconds)
	}

	warlock.RegisterResetEffect(func(sim *core.Simulation) {
		timeToNextHeal := rollHealingCadence(sim)
		pa := &core.PendingAction{NextActionAt: sim.CurrentTime + timeToNextHeal}

		pa.OnAction = func(sim *core.Simulation) {
			totalHeal := hps * timeToNextHeal.Seconds() * warlock.PseudoStats.HealingTakenMultiplier
			warlock.GainHealth(sim, totalHeal, healthMetrics)

			result := healingModelSpell.NewResult(&warlock.Unit)
			result.Damage = totalHeal
			warlock.OnHealTaken(sim, healingModelSpell, result)
			healingModelSpell.DisposeResult(result)

			timeToNextHeal = rollHealingCadence(sim)
			pa.NextActionAt = sim.CurrentTime + timeToNextHeal
			sim.AddPendingAction(pa)
		}

		sim.AddPendingAction(pa)
	})
}

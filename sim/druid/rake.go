package druid

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

type RakeRankInfo struct {
	id            int32
	level         int32
	initialDamage float64
	dotTickDamage float64
}

var rakeSpells = []RakeRankInfo{
	{
		id:            1822,
		level:         24,
		initialDamage: 19.0,
		dotTickDamage: 13.0,
	},
	{
		id:            1823,
		level:         34,
		initialDamage: 28.0,
		dotTickDamage: 19.0,
	},
	{

		id:            1824,
		level:         44,
		initialDamage: 43.0,
		dotTickDamage: 25.0,
	},
	{

		id:            9904,
		level:         54,
		initialDamage: 58.0,
		dotTickDamage: 32.0,
	},
}

func (druid *Druid) registerRakeSpell() {
	// Add highest available rake rank for level.
	for rank := len(rakeSpells) - 1; rank >= 0; rank-- {
		if druid.Level >= rakeSpells[rank].level {
			config := druid.newRakeSpellConfig(rakeSpells[rank])
			druid.Rake = druid.RegisterSpell(Cat, config)
			return
		}
	}
}

func (druid *Druid) newRakeSpellConfig(rakeRank RakeRankInfo) core.SpellConfig {
	baseDamageInitial := rakeRank.initialDamage
	baseDamageTick := rakeRank.dotTickDamage
	energyCost := 40 - float64(druid.Talents.Ferocity)

	return core.SpellConfig{
		SpellCode:   SpellCode_DruidRake,
		ActionID:    core.ActionID{SpellID: rakeRank.id},
		SpellSchool: core.SpellSchoolPhysical,
		DefenseType: core.DefenseTypeMelee,
		ProcMask:    core.ProcMaskMeleeMHSpecial,
		Flags:       core.SpellFlagMeleeMetrics | core.SpellFlagIgnoreResists | core.SpellFlagBinary | core.SpellFlagAPL | SpellFlagOmen | SpellFlagBuilder,

		EnergyCost: core.EnergyCostOptions{
			Cost:   energyCost,
			Refund: 0.8,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: time.Second,
			},
			IgnoreHaste: true,
		},

		DamageMultiplierAdditive: 1 + 0.1*float64(druid.Talents.SavageFury),
		DamageMultiplier:         1,
		ThreatMultiplier:         1,

		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "Rake",
			},
			NumberOfTicks: 3,
			TickLength:    time.Second * 3,
			OnSnapshot: func(sim *core.Simulation, target *core.Unit, dot *core.Dot, isRollover bool) {
				damage := baseDamageTick
				dot.Snapshot(target, damage, isRollover)
			},
			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				dot.CalcAndDealPeriodicSnapshotDamage(sim, target, dot.OutcomeTick)
			},
		},

		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			baseDamage := baseDamageInitial
			result := spell.CalcAndDealDamage(sim, target, baseDamage, spell.OutcomeMeleeSpecialHitAndCrit)

			if result.Landed() {
				druid.AddComboPoints(sim, 1, target, spell.ComboPointMetrics())
				spell.Dot(target).Apply(sim)
			} else {
				spell.IssueRefund(sim)
			}
		},

		ExpectedInitialDamage: func(sim *core.Simulation, target *core.Unit, spell *core.Spell, _ bool) *core.SpellResult {
			baseDamage := baseDamageInitial
			initial := spell.CalcPeriodicDamage(sim, target, baseDamage, spell.OutcomeExpectedMagicAlwaysHit)

			attackTable := spell.Unit.AttackTables[target.UnitIndex][spell.CastType]
			critChance := spell.PhysicalCritChance(attackTable)
			critMod := critChance * (spell.CritMultiplier(attackTable) - 1)
			initial.Damage *= 1 + critMod
			return initial
		},
		ExpectedTickDamage: func(sim *core.Simulation, target *core.Unit, spell *core.Spell, _ bool) *core.SpellResult {
			tickBase := baseDamageTick
			ticks := spell.CalcPeriodicDamage(sim, target, tickBase, spell.OutcomeExpectedMagicAlwaysHit)
			return ticks
		},
	}
}

func (druid *Druid) CurrentRakeCost() float64 {
	return druid.Rake.Cost.GetCurrentCost()
}

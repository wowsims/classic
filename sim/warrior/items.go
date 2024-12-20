package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

const (
	DiamondFlask = 20130
	GrileksCharmOfMight = 19951
	RageOfMugamba = 19577
	LifegivingGem = 19341
)

func init() {
	core.AddEffectsToTest = false

	core.NewItemEffect(DiamondFlask, func(agent core.Agent) {
		character := agent.GetCharacter()

		buffAura := character.NewTemporaryStatsAura("Diamond Flask", core.ActionID{SpellID: 24427}, stats.Stats{stats.Strength: 75}, time.Second*60)

		triggerSpell := character.GetOrRegisterSpell(core.SpellConfig{
			ActionID: core.ActionID{SpellID: 24427},
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 6,
				},
				SharedCD: core.Cooldown{
					Timer:    character.GetOffensiveTrinketCD(),
					Duration: time.Second * 60,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Spell:    triggerSpell,
			Priority: core.CooldownPriorityDefault,
			Type:     core.CooldownTypeDPS,
		})
	})

	core.NewItemEffect(GrileksCharmOfMight, func(agent core.Agent) {
		warrior := agent.(WarriorAgent).GetWarrior()
		actionID := core.ActionID{ItemID: GrileksCharmOfMight}
		rageMetrics := warrior.NewRageMetrics(actionID)


		spell := warrior.Character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolPhysical,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    warrior.NewTimer(),
					Duration: time.Minute * 3,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				warrior.AddRage(sim, 30, rageMetrics)
			},
		})

		warrior.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	core.NewItemEffect(RageOfMugamba, func(agent core.Agent) {
		warrior := agent.(WarriorAgent).GetWarrior()

		warrior.RegisterAura(core.Aura{
			Label: "Reduces the cost of your Hamstring ability by 2 rage points.",
			OnInit: func(aura *core.Aura, sim *core.Simulation) {
				warrior.Hamstring.Cost.FlatModifier -= 2
			},
		})
	})

	core.NewItemEffect(LifegivingGem, func(agent core.Agent) {
		warrior := agent.(WarriorAgent).GetWarrior()
		actionID := core.ActionID{ItemID: LifegivingGem}
		healthMetrics := warrior.NewHealthMetrics(actionID)
	
		var bonusHealth float64
		lifegivingGemAura := warrior.RegisterAura(core.Aura{
			Label:    "Gift of Life",
			ActionID: core.ActionID{SpellID: 23725},
			Duration: time.Second * 20,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				bonusHealth = warrior.MaxHealth() * 0.15
				warrior.AddStatsDynamic(sim, stats.Stats{stats.Health: bonusHealth})
				warrior.GainHealth(sim, bonusHealth, healthMetrics)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				warrior.AddStatsDynamic(sim, stats.Stats{stats.Health: -bonusHealth})
			},
		})
	
		lifegivingGemSpell := warrior.RegisterSpell(AnyStance, core.SpellConfig{
			ActionID: actionID,
	
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    warrior.NewTimer(),
					Duration: time.Minute * 5,
				},
			},
	
			ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
				lifegivingGemAura.Activate(sim)
			},
		})
	
		warrior.AddMajorCooldown(core.MajorCooldown{
			Spell: lifegivingGemSpell.Spell,
			Type:  core.CooldownTypeSurvival,
		})
	})

	core.AddEffectsToTest = true
}

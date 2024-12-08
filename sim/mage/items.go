package mage

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

const (
	FireRuby = 20036
	MindQuickeningGem        	    = 19339
	HazzarahsCharmOfMagic 			= 19959
	JewelOfKajaro                	= 19601
)

func init() {
	core.AddEffectsToTest = false

	core.NewItemEffect(FireRuby, func(agent core.Agent) {
		character := agent.GetCharacter()

		actionID := core.ActionID{ItemID: FireRuby}
		manaMetrics := character.NewManaMetrics(actionID)

		damageAura := character.GetOrRegisterAura(core.Aura{
			Label:    "Chaos Fire",
			ActionID: core.ActionID{SpellID: 24389},
			Duration: time.Minute * 1,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatDynamic(sim, stats.FirePower, 100)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				character.AddStatDynamic(sim, stats.FirePower, -100)
			},
			OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
				if spell.SpellSchool.Matches(core.SpellSchoolFire) {
					aura.Deactivate(sim)
				}
			},
		})

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolPhysical,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,

			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Minute * 3,
				},
			},

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AddMana(sim, sim.Roll(1, 500), manaMetrics)
				damageAura.Activate(sim)
			},
		})

		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://www.wowhead.com/classic/item=19959/hazzarahs-charm-of-magic
	// Increases the critical hit chance of your Arcane spells by 5%, and increases the critical hit damage of your Arcane spells by 50% for 20 sec. 
	// (3 Min Cooldown)
	core.NewItemEffect(HazzarahsCharmOfMagic, func(agent core.Agent) {
		mage := agent.(MageAgent).GetMage()

		duration := time.Second * 20
		affectedSpells := []*core.Spell{}

		aura := mage.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 24544},
			Label:    "Arcane Potency",
			Duration: duration,
			OnInit: func(aura *core.Aura, sim *core.Simulation) {
				for spellIdx := range mage.Spellbook {
					if spell := mage.Spellbook[spellIdx]; spell.SpellSchool == core.SpellSchoolArcane {
						affectedSpells = append(affectedSpells, spell)
					}
				}
			},
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range affectedSpells {
					spell.BonusCritRating += 5 * core.SpellCritRatingPerCritChance
					spell.CritDamageBonus += 0.50
				}
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				for _, spell := range affectedSpells {
					spell.BonusCritRating -= 5 * core.SpellCritRatingPerCritChance
					spell.CritDamageBonus -= 0.50
				}
			},
		})

		spell := mage.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{ItemID: HazzarahsCharmOfMagic},
			SpellSchool: core.SpellSchoolArcane,
			Flags:       core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    mage.NewTimer(),
					Duration: time.Minute * 3,
				},
				SharedCD: core.Cooldown{
					Timer:    mage.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				aura.Activate(sim)
			},
		})

		mage.AddMajorCooldown(core.MajorCooldown{
			Spell:    spell,
			Priority: core.CooldownPriorityBloodlust,
			Type:     core.CooldownTypeDPS,
		})
	})

	// https://www.wowhead.com/classic/item=19601/jewel-of-kajaro
	// Equip: Reduces the cooldown of Counterspell by 2 sec.
	core.NewItemEffect(JewelOfKajaro, func(agent core.Agent) {
		mage := agent.(MageAgent).GetMage()

		mage.RegisterAura(core.Aura{
			Label: "Improved Counterspell",
			OnInit: func(aura *core.Aura, sim *core.Simulation) {
				mage.Counterspell.CD.Duration -= time.Second * 2
			},
		})
	})

	// https://www.wowhead.com/classic/item=19339/mind-quickening-gem
	// Use: Quickens the mind, increasing the Mage's casting speed of non-channeled spells by 33% for 20 sec. (2 Min Cooldown)
	core.NewItemEffect(MindQuickeningGem, func(agent core.Agent) {
		mage := agent.(MageAgent).GetMage()

		actionID := core.ActionID{ItemID: MindQuickeningGem}
		duration := time.Second * 20

		buffAura := mage.RegisterAura(core.Aura{
			ActionID: core.ActionID{SpellID: 23723},
			Label:    "Mind Quickening",
			Duration: duration,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				mage.MultiplyCastSpeed(1.33)
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				mage.MultiplyCastSpeed(1 / 1.33)
			},
		})

		spell := mage.RegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    mage.NewTimer(),
					Duration: time.Minute * 5,
				},
				SharedCD: core.Cooldown{
					Timer:    mage.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		mage.AddMajorCooldown(core.MajorCooldown{
			Spell:    spell,
			Priority: core.CooldownPriorityBloodlust,
			Type:     core.CooldownTypeDPS,
		})
	})

	core.AddEffectsToTest = true
}

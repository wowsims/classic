package druid

import (
	"time"

	"github.com/wowsims/classic/sim/core"
)

// Item IDs
const (
	WolfsheadHelm       = 8345
	IdolOfFerocity      = 22397
	IdolOfTheMoon       = 23197
	IdolOfBrutality     = 23198
	RuneOfMetamorphosis = 19340
)

func init() {
	core.AddEffectsToTest = false

	// https://www.wowhead.com/classic/item=22397/idol-of-ferocity
	// Equip: Reduces the energy cost of Claw and Rake by 3.
	core.NewItemEffect(IdolOfFerocity, func(agent core.Agent) {
		druid := agent.(DruidAgent).GetDruid()

		druid.OnSpellRegistered(func(spell *core.Spell) {
			if spell.SpellCode == SpellCode_DruidRake || spell.SpellCode == SpellCode_DruidClaw {
				spell.Cost.FlatModifier -= 3
			}
		})
	})

	// https://www.wowhead.com/classic/item=23197/idol-of-the-moon
	// Equip: Increases the damage of your Moonfire spell by up to 33.
	core.NewItemEffect(IdolOfTheMoon, func(agent core.Agent) {
		druid := agent.(DruidAgent).GetDruid()
		druid.OnSpellRegistered(func(spell *core.Spell) {
			if spell.SpellCode == SpellCode_DruidMoonfire {
				spell.BonusDamage += 33
			}
		})
	})

	// https://www.wowhead.com/classic/item=23198/idol-of-brutality
	// Equip: Reduces the rage cost of Maul and Swipe by 3.
	core.NewItemEffect(IdolOfBrutality, func(agent core.Agent) {
		// Implemented in maul.go and swipe.go
	})

	// https://www.wowhead.com/classic/item=19340/rune-of-metamorphosis
	// Use: Decreases the mana cost of all Druid shapeshifting forms by 100% for 20 sec. (5 Min Cooldown)
	core.NewItemEffect(RuneOfMetamorphosis, func(agent core.Agent) {
		druid := agent.(DruidAgent).GetDruid()

		actionID := core.ActionID{SpellID: 23724}
		duration := time.Second * 20
		cooldown := time.Minute * 5

		buffAura := druid.GetOrRegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Metamorphosis Rune",
			Duration: duration,
			OnGain: func(aura *core.Aura, sim *core.Simulation) {
				druid.CatForm.Cost.Multiplier -= 100
				//druid.BearForm.Cost.Multiplier -= 100
				//druid.MoonkinForm.Cost.Multiplier -= 100
			},
			OnExpire: func(aura *core.Aura, sim *core.Simulation) {
				druid.CatForm.Cost.Multiplier += 100
				//druid.BearForm.Cost.Multiplier += 100
				//druid.MoonkinForm.Cost.Multiplier += 100
			},
		})

		spell := druid.GetOrRegisterSpell(core.SpellConfig{
			ActionID: actionID,
			Flags:    core.SpellFlagNoOnCastComplete | core.SpellFlagOffensiveEquipment,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    druid.NewTimer(),
					Duration: cooldown,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})

		druid.AddMajorCooldown(core.MajorCooldown{
			Spell: spell,
			Type:  core.CooldownTypeDPS,
		})
	})

	core.AddEffectsToTest = true
}

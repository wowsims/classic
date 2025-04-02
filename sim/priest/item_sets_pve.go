package priest

import (
	"slices"
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 1 Item Sets - Molten Core
///////////////////////////////////////////////////////////////////////////

var ItemSetVestmentsOfProphecy = core.NewItemSet(core.ItemSet{
	Name: "Vestments of Prophecy",
	Bonuses: map[int32]core.ApplyEffect{
		// -0.1 sec to the casting time of your Flash Heal spell.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// Improves your chance to get a critical strike with Holy spells by 2%.
		5: func(agent core.Agent) {
			priest := agent.(PriestAgent).GetPriest()
			priest.PseudoStats.SchoolBonusCritChance[stats.SchoolIndexHoly] += 2 * core.SpellCritRatingPerCritChance
		},
		// Increases your chance of a critical hit with Prayer of Healing by 25%.
		8: func(agent core.Agent) {
			// Nothing to do
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 3 Item Sets - BWL
///////////////////////////////////////////////////////////////////////////

var ItemSetVestmentsOfTranscendence = core.NewItemSet(core.ItemSet{
	Name: "Vestments of Transcendence",
	Bonuses: map[int32]core.ApplyEffect{
		// Allows 15% of your Mana regeneration to continue while casting.
		3: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.PseudoStats.SpiritRegenRateCasting += .15
		},
		// When struck in melee there is a 50% chance you will Fade for 4 sec.
		5: func(agent core.Agent) {
			// Nothing to do
		},
		// Your Greater Heals now have a heal over time component equivalent to a rank 5 Renew.
		8: func(agent core.Agent) {
			// Nothing to do
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 4 Item Sets - ZG and AB
///////////////////////////////////////////////////////////////////////////

var ItemSetConfessorsRaiment = core.NewItemSet(core.ItemSet{
	Name: "Confessor's Raiment",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases healing done by spells and effects by up to 22.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.HealingPower, 22)
		},
		// Increase the range of your Smite and Holy Fire spells by 5 yds.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// Reduces the casting time of your Mind Control spell by 0.5 sec.
		5: func(agent core.Agent) {
			// Nothing to do
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 5 Item Sets - AQ
///////////////////////////////////////////////////////////////////////////

var ItemSetGarmentsOfTheOracle = core.NewItemSet(core.ItemSet{
	Name: "Garments of the Oracle",
	Bonuses: map[int32]core.ApplyEffect{
		// 20% chance that your heals on others will also heal you 10% of the amount healed.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// Increases the duration of your Renew spell by 3 sec.
		5: func(agent core.Agent) {
			// Nothing to do
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 6 Item Sets - Naxx
///////////////////////////////////////////////////////////////////////////

var ItemSetVestmentsOfFaith = core.NewItemSet(core.ItemSet{
	Name: "Vestments of Faith",
	Bonuses: map[int32]core.ApplyEffect{
		// Reduces the mana cost of your Renew spell by 12%.
		2: func(agent core.Agent) {
			// Nothing to do
		},
		// On Greater Heal critical hits, your target will gain Armor of Faith, absorbing up to 500 damage.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// Reduces the threat from your healing spells.
		6: func(agent core.Agent) {
			// Nothing to do
		},
		// Each spell you cast can trigger an Epiphany, increasing your mana regeneration by 24 for 30 sec.
		8: func(agent core.Agent) {
			c := agent.GetCharacter()

			procAura := c.NewTemporaryStatsAura("Epiphany", core.ActionID{SpellID: 28802}, stats.Stats{stats.MP5: 24}, time.Second*30)
			core.MakeProcTriggerAura(&c.Unit, core.ProcTrigger{
				Name:       "Item - Epiphany Proc (Spell Cast)",
				Callback:   core.CallbackOnCastComplete,
				ProcMask:   core.ProcMaskSpellDamage,
				SpellFlags:   core.SpellFlagHelpful,
				ProcChance: 0.05,
				Handler: func(sim *core.Simulation, spell *core.Spell, _ *core.SpellResult) {
					procAura.Activate(sim)
				},
			})
		},
	},
})

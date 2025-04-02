package paladin

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 1 Item Sets - Molten Core
///////////////////////////////////////////////////////////////////////////

var ItemSetVestmentsOfProphecy = core.NewItemSet(core.ItemSet{
	Name: "Lawbringer Armor",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases the chance of triggering a Judgement of Light heal by 10%.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// Improves your chance to get a critical strike with spells by 1%.
		// Improves your chance to get a critical strike by 1%.
		5: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()
			paladin.AddStat(stats.MeleeCrit, 1)
			paladin.AddStat(stats.SpellCrit, 1)
		},
		// Gives the Paladin a chance on every melee hit to heal your party for 189 to 211.
		8: func(agent core.Agent) {
			// Nothing to do
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 3 Item Sets - BWL
///////////////////////////////////////////////////////////////////////////

var ItemSetSoulforgeArmor = core.NewItemSet(core.ItemSet{
	Name: "Judgement Armor",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases the radius of a Paladin's auras by 10.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// Increases damage and healing done by magical spells and effects by up to 47.
		5: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 47)
		},
		// Inflicts 60 to 66 additional Holy damage on the target of a Paladin's Judgement.
		8: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()
			paladin.RegisterAura(core.Aura{
				Label: "Judgement - T2 - Paladin - 8P Bonus",
				SpellSchool: core.SpellSchoolHoly,
				OnInit: func(aura *core.MakePermanent, sim *core.Simulation) {
				OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
					if spell.SpellCode == SpellCode_PaladinJudgement && result.Landed() {
						spell.CalcAndDealDamage(sim, target, sim.Roll(60, 66), spell.OutcomeMagicCrit)
					}
				},
			})
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 4 Item Sets - ZG and AB
///////////////////////////////////////////////////////////////////////////

var ItemSetConfessorsRaiment = core.NewItemSet(core.ItemSet{
	Name: "Freethinker's Armor",
	Bonuses: map[int32]core.ApplyEffect{
		// Restores 4 mana per 5 sec.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.MP5, 4)
		},
		// Reduces the casting time of your Holy Light spell by 0.1 sec.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// Increases the duration of all Blessings by 10%.
		5: func(agent core.Agent) {
			// Nothing to do
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 5 Item Sets - AQ
///////////////////////////////////////////////////////////////////////////

var ItemSetGarmentsOfTheOracle = core.NewItemSet(core.ItemSet{
	Name: "Avenger's Battlegear",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases the duration of your Judgements by 20%.
		3: func(agent core.Agent) {
			// Nothing to do
		},
		// Increases damage and healing done by magical spells and effects by up to 71.
		5: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 71)
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 6 Item Sets - Naxx
///////////////////////////////////////////////////////////////////////////

var ItemSetVestmentsOfFaith = core.NewItemSet(core.ItemSet{
	Name: "Redemption Armor",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases the amount healed by your Judgement of Light by 20.
		2: func(agent core.Agent) {
			// Nothing to do
		},
		// Reduces cooldown on your Lay on Hands by 12 min.
		4: func(agent core.Agent) {
			// Nothing to do
		},
		// Your Flash of Light and Holy Light spells have a chance to imbue your target with Holy Power.
		6: func(agent core.Agent) {
			// Nothing to do
		},
		// Your Cleanse spell also heals the target for 200.
		8: func(agent core.Agent) {
			// Nothing to do
		},
	},
})

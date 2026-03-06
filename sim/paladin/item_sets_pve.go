package paladin

import (
	"slices"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 1 Item Sets - Molten Core
///////////////////////////////////////////////////////////////////////////

var ItemSetLawbringerArmor = core.NewItemSet(core.ItemSet{
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

var ItemSetJudgementArmor = core.NewItemSet(core.ItemSet{
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

			spellCodes := []int32{SpellCode_PaladinJudgementOfCommand, SpellCode_PaladinJudgementOfRighteousness}
			paladin.RegisterAura(core.Aura{
				Label: "Judgement - T2 - Paladin - 8P Bonus",
				OnSpellHitDealt: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell, result *core.SpellResult) {
					if slices.Contains(spellCodes, spell.SpellCode) && result.Landed() {
						spell.CalcAndDealDamage(sim, result.Target, sim.Roll(60, 66), spell.OutcomeMagicCrit)
					}
				},
			})
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 4 Item Sets - ZG and AB
///////////////////////////////////////////////////////////////////////////

var ItemSetFreethinkersArmor = core.NewItemSet(core.ItemSet{
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

var ItemSetAvengersBattlegear = core.NewItemSet(core.ItemSet{
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

var ItemSetBattlegearOfEternalJustice = core.NewItemSet(core.ItemSet{
	Name: "Battlegear of Eternal Justice",
	Bonuses: map[int32]core.ApplyEffect{
		// 20% chance to regain 100 mana when you cast a Judgement.
		3: func(agent core.Agent) {
			paladin := agent.(PaladinAgent).GetPaladin()

			manaMetrics := paladin.NewManaMetrics(core.ActionID{SpellID: 0})

			spellCodes := []int32{
				SpellCode_PaladinJudgementOfCommand,
				SpellCode_PaladinJudgementOfRighteousness,
			}

			paladin.RegisterAura(core.Aura{
				Label: "Battlegear of Eternal Justice 3 pc",
				OnCastComplete: func(aura *core.Aura, sim *core.Simulation, spell *core.Spell) {
					if slices.Contains(spellCodes, spell.SpellCode) {
						if sim.Proc(0.20, "Battlegear of Eternal Justice 3 pc") {
							paladin.AddMana(sim, 100, manaMetrics)
						}
					}
				},
			})
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                            Classic Phase 6 Item Sets - Naxx
///////////////////////////////////////////////////////////////////////////

var ItemSetRedemptionArmor = core.NewItemSet(core.ItemSet{
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

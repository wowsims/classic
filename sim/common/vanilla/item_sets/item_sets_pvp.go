package item_sets

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/stats"
)

///////////////////////////////////////////////////////////////////////////
//                                 Cloth
///////////////////////////////////////////////////////////////////////////

var ItemSetTheHighlandersIntent = core.NewItemSet(core.ItemSet{
	Name: "The Highlander's Intent",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Spells.
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.SpellCrit, 1)
		},
	},
})

var ItemSetTheDefilersIntent = core.NewItemSet(core.ItemSet{
	Name: "The Defiler's Intent",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Spells.
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.SpellCrit, 1)
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                                 Leather
///////////////////////////////////////////////////////////////////////////

var ItemSetTheHighlandersPurpose = core.NewItemSet(core.ItemSet{
	Name: "The Highlander's Purpose",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Melee.
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.MeleeCrit, 1)
		},
	},
})

var ItemSetTheHighlandersWill = core.NewItemSet(core.ItemSet{
	Name: "The Highlander's Will",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Spells.
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.SpellCrit, 1)
		},
	},
})

var ItemSetTheDefilersPurpose = core.NewItemSet(core.ItemSet{
	Name: "The Defiler's Purpose",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Melee.
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.MeleeCrit, 1)
		},
	},
})

var ItemSetTheDefilersWill = core.NewItemSet(core.ItemSet{
	Name: "The Defiler's Will",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Spells.
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.SpellCrit, 1)
		},
	},
})


///////////////////////////////////////////////////////////////////////////
//                                 Mail
///////////////////////////////////////////////////////////////////////////

var ItemSetTheHighlandersFortitude = core.NewItemSet(core.ItemSet{
	Name: "The Highlander's Fortitude",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Spells.
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.SpellCrit, 1)
		},
	},
})

var ItemSetTheHighlandersDetermination = core.NewItemSet(core.ItemSet{
	Name: "The Highlander's Determination",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Melee.
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.MeleeCrit, 1)
		},
	},
})

var ItemSetTheDefilersFortitude = core.NewItemSet(core.ItemSet{
	Name: "The Defiler's Fortitude",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Melee.
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.MeleeCrit, 1)
		},
	},
})

var ItemSetTheDefilersDetermination = core.NewItemSet(core.ItemSet{
	Name: "The Defiler's Determination",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Melee.
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.MeleeCrit, 1)
		},
	},
})

///////////////////////////////////////////////////////////////////////////
//                                 Plate
///////////////////////////////////////////////////////////////////////////

var ItemSetTheHighlandersResolve = core.NewItemSet(core.ItemSet{
	Name: "The Highlander's Resolve",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Melee.
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.MeleeCrit, 1)
		},
	},
})

var ItemSetTheHighlandersResolution = core.NewItemSet(core.ItemSet{
	Name: "The Highlander's Resolution",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Melee
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.MeleeCrit, 1)
		},
	},
})

var ItemSetTheDefilersResolution = core.NewItemSet(core.ItemSet{
	Name: "The Defiler's Resolution",
	Bonuses: map[int32]core.ApplyEffect{
		// Increase Stamina +5
		2: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.Stamina, 5)
		},
		// +1 Crit with Melee
		3: func(agent core.Agent) {
			character := agent.GetCharacter()
			character.AddStat(stats.MeleeCrit, 1)
		},
	},
})

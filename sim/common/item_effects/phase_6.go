package item_effects

import (
	"time"

	"github.com/wowsims/classic/sim/common/itemhelpers"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

const (
	CorruptedAshbringer             = 22691
	KissOfTheSpider                 = 22954
	GlyphOfDeflection               = 23040
	SlayersCrest                    = 23041
	TheRestrainedEssenceOfSapphiron = 23046
	MarkOfTheChampionPhys           = 23206
	MarkOfTheChampionSpell          = 23207
	WarmthOfForgiveness             = 23027
)

func init() {
	core.AddEffectsToTest = false

	// https://www.wowhead.com/classic/item=22691/corrupted-ashbringer
	// Chance on hit: Steals 185 to 215 life from target enemy.
	// Proc rate taken from Classic 2019 testing
	// It was reported in Vanilla to scale with Spell Damage but during 2019 it was reported NOT to
	itemhelpers.CreateWeaponProcSpell(CorruptedAshbringer, "Corrupted Ashbringer", 1.6, func(character *core.Character) *core.Spell {
		actionID := core.ActionID{SpellID: 29155}
		healthMetrics := character.NewHealthMetrics(actionID)
		return character.RegisterSpell(core.SpellConfig{
			ActionID:         actionID,
			SpellSchool:      core.SpellSchoolShadow,
			DefenseType:      core.DefenseTypeMagic,
			ProcMask:         core.ProcMaskSpellProc | core.ProcMaskSpellDamageProc,
			Flags:            core.SpellFlagNoOnCastComplete | core.SpellFlagPassiveSpell,
			DamageMultiplier: 1,
			ThreatMultiplier: 1,
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				result := spell.CalcAndDealDamage(sim, target, sim.Roll(185, 215), spell.OutcomeMagicHit)
				character.GainHealth(sim, result.Damage, healthMetrics)
			},
		})
	})

	// https://www.wowhead.com/classic/item=23040/glyph-of-deflection
	// Use: Increases the block value of your shield by 235 for 20 sec. (2 Min Cooldown)
	core.NewSimpleStatDefensiveTrinketEffect(GlyphOfDeflection, stats.Stats{stats.BlockValue: 235}, time.Second*20, time.Minute*2)

	// https://www.wowhead.com/classic/item=22954/kiss-of-the-spider
	// Use: Increases your attack speed by 20% for 15 sec. (2 Min Cooldown)
	core.NewItemEffect(KissOfTheSpider, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{ItemID: KissOfTheSpider}
		duration := time.Second * 15
		buffAura := character.RegisterAura(core.Aura{
			ActionID: actionID,
			Label:    "Kiss of the Spider",
			Duration: duration,
		}).AttachMultiplyAttackSpeed(&character.Unit, 1.20)

		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolPhysical,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(),
					Duration: time.Second * 120,
				},
				SharedCD: core.Cooldown{
					Timer:    character.GetOffensiveTrinketCD(),
					Duration: duration,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				buffAura.Activate(sim)
			},
		})
		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	// https://wowhead.com/classic/item=23206?level=60&rand=0
	// Equip: +150 Attack Power when fighting Undead and Demons.
	core.NewMobTypeAttackPowerEffect(MarkOfTheChampionPhys, []proto.MobType{proto.MobType_MobTypeUndead, proto.MobType_MobTypeDemon}, 150)

	// https://www.wowhead.com/classic/item=23207/mark-of-the-champion
	// Equip: Increases damage done to Undead and Demons by magical spells and effects by up to 85.
	core.NewMobTypeSpellPowerEffect(MarkOfTheChampionSpell, []proto.MobType{proto.MobType_MobTypeUndead, proto.MobType_MobTypeDemon}, 85)

	// https://www.wowhead.com/classic/item=236334/slayers-crest
	// Use: Increases Attack Power by 280 for 20 sec. (2 Min Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(SlayersCrest, stats.Stats{stats.AttackPower: 260, stats.RangedAttackPower: 260}, time.Second*20, time.Minute*2)

	// https://www.wowhead.com/classic/item=23046/the-restrained-essence-of-sapphiron
	// Use: Increases damage and healing done by magical spells and effects by up to 130 for 20 sec. (2 Min Cooldown)
	core.NewSimpleStatOffensiveTrinketEffect(TheRestrainedEssenceOfSapphiron, stats.Stats{stats.SpellPower: 130}, time.Second*20, time.Minute*2)

	// https://www.wowhead.com/classic/item=23027/warmth-of-forgiveness
	// Use: Restores 500 mana. (3 Min Cooldown)
	core.NewItemEffect(WarmthOfForgiveness, func(agent core.Agent) {
		character := agent.GetCharacter()
		actionID := core.ActionID{ItemID: WarmthOfForgiveness}
		manaMetrics := character.NewManaMetrics(actionID)
		spell := character.RegisterSpell(core.SpellConfig{
			ActionID:    actionID,
			SpellSchool: core.SpellSchoolPhysical,
			ProcMask:    core.ProcMaskEmpty,
			Flags:       core.SpellFlagNoOnCastComplete,
			Cast: core.CastConfig{
				CD: core.Cooldown{
					Timer:    character.NewTimer(), // Doesn't share the trinket timer
					Duration: time.Minute * 3,
				},
			},
			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				character.AddMana(sim, 500, manaMetrics)
			},
		})
		character.AddMajorCooldown(core.MajorCooldown{
			Type:  core.CooldownTypeDPS,
			Spell: spell,
		})
	})

	core.AddEffectsToTest = true
}

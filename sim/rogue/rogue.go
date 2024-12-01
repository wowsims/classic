package rogue

import (
	"time"

	"github.com/wowsims/classic/sim/common/guardians"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

const (
	SpellFlagBuilder      = core.SpellFlagAgentReserved1
	SpellFlagColdBlooded  = core.SpellFlagAgentReserved2
	SpellFlagDeadlyBrewed = core.SpellFlagAgentReserved3
	SpellFlagCarnage      = core.SpellFlagAgentReserved4 // for Carnage
	SpellFlagRoguePoison  = core.SpellFlagAgentReserved5 // RogueT1
)

const (
	SpellCode_RogueNone int32 = iota

	SpellCode_RogueAmbush
	SpellCode_RogueAdrenalineRush
	SpellCode_RogueBackstab
	SpellCode_RogueBladeFlurry
	SpellCode_RogueEviscerate
	SpellCode_RogueExposeArmor
	SpellCode_RogueGarrote
	SpellCode_RogueGhostlyStrike
	SpellCode_RogueHemorrhage
	SpellCode_RogueRupture
	SpellCode_RogueSinisterStrike
	SpellCode_RogueSliceandDice
	SpellCode_RogueVanish
)

var TalentTreeSizes = [3]int{15, 19, 17}

type Rogue struct {
	core.Character

	Talents *proto.RogueTalents
	Options *proto.RogueOptions

	sliceAndDiceDurations [6]time.Duration

	AdrenalineRush *core.Spell
	Backstab       *core.Spell
	BladeFlurry    *core.Spell
	Feint          *core.Spell
	Garrote        *core.Spell
	Ambush         *core.Spell
	Hemorrhage     *core.Spell
	GhostlyStrike  *core.Spell
	SinisterStrike *core.Spell
	Shadowstep     *core.Spell
	Preparation    *core.Spell
	Premeditation  *core.Spell
	ColdBlood      *core.Spell
	Vanish         *core.Spell

	Eviscerate   *core.Spell
	ExposeArmor  *core.Spell
	Rupture      *core.Spell
	SliceAndDice *core.Spell
	Finishers    []*core.Spell

	Evasion *core.Spell

	DeadlyPoison     *core.Spell
	deadlyPoisonTick *core.Spell
	InstantPoison    *core.Spell
	WoundPoison      *core.Spell

	additivePoisonBonusChance float64

	AdrenalineRushAura *core.Aura
	BladeFlurryAura    *core.Aura
	ExposeArmorAuras   core.AuraArray
	EvasionAura        *core.Aura
	SliceAndDiceAura   *core.Aura
	StealthAura        *core.Aura
	VanishAura         *core.Aura

	woundPoisonDebuffAuras core.AuraArray
}

func (rogue *Rogue) GetCharacter() *core.Character {
	return &rogue.Character
}

func (rogue *Rogue) GetRogue() *Rogue {
	return rogue
}

func (rogue *Rogue) AddRaidBuffs(_ *proto.RaidBuffs)   {}
func (rogue *Rogue) AddPartyBuffs(_ *proto.PartyBuffs) {}

func (rogue *Rogue) finisherFlags() core.SpellFlag {
	return core.SpellFlagMeleeMetrics | core.SpellFlagAPL
}

func (rogue *Rogue) builderFlags() core.SpellFlag {
	return SpellFlagBuilder | SpellFlagColdBlooded | core.SpellFlagMeleeMetrics | core.SpellFlagAPL
}

func (rogue *Rogue) Initialize() {
	rogue.registerBackstabSpell()
	rogue.registerEviscerate()
	rogue.registerExposeArmorSpell()
	rogue.registerFeintSpell()
	rogue.registerGarrote()
	rogue.registerHemorrhageSpell()
	rogue.registerRupture()
	rogue.registerSinisterStrikeSpell()
	rogue.registerSliceAndDice()
	rogue.registerThistleTeaCD()
	rogue.registerAmbushSpell()

	// Poisons
	rogue.registerInstantPoisonSpell()
	rogue.registerDeadlyPoisonSpell()
	rogue.registerWoundPoisonSpell()

	// Stealth
	rogue.registerStealthAura()
	rogue.registerVanishSpell()
}

func (rogue *Rogue) ApplyEnergyTickMultiplier(multiplier float64) {
	rogue.EnergyTickMultiplier += multiplier
}

func (rogue *Rogue) Reset(_ *core.Simulation) {
	for _, mcd := range rogue.GetMajorCooldowns() {
		mcd.Disable()
	}
}

func NewRogue(character *core.Character, options *proto.Player, rogueOptions *proto.RogueOptions) *Rogue {
	rogue := &Rogue{
		Character: *character,
		Talents:   &proto.RogueTalents{},
		Options:   rogueOptions,
	}
	core.FillTalentsProto(rogue.Talents.ProtoReflect(), options.TalentsString, TalentTreeSizes)

	// Passive rogue threat reduction: https://wotlk.wowhead.com/spell=21184/rogue-passive-dnd
	rogue.PseudoStats.ThreatMultiplier *= 0.71
	// TODO: Be able to Parry based on results
	rogue.PseudoStats.CanParry = true
	maxEnergy := 100.0
	if rogue.Talents.Vigor {
		maxEnergy += 10
	}
	rogue.EnableEnergyBar(maxEnergy)

	rogue.EnableAutoAttacks(rogue, core.AutoAttackOptions{
		MainHand:       rogue.WeaponFromMainHand(),
		OffHand:        rogue.WeaponFromOffHand(),
		Ranged:         rogue.WeaponFromRanged(),
		AutoSwingMelee: true,
	})
	rogue.applyPoisons()

	rogue.AddStatDependency(stats.Strength, stats.AttackPower, core.APPerStrength[character.Class])
	rogue.AddStatDependency(stats.Agility, stats.AttackPower, 1)
	rogue.AddStatDependency(stats.Agility, stats.RangedAttackPower, 1)
	rogue.AddStatDependency(stats.Agility, stats.MeleeCrit, core.CritPerAgiAtLevel[character.Class]*core.CritRatingPerCritChance)
	rogue.AddStatDependency(stats.Agility, stats.Dodge, core.DodgePerAgiAtLevel[character.Class]*core.DodgeRatingPerDodgeChance)
	rogue.AddStatDependency(stats.BonusArmor, stats.Armor, 1)

	guardians.ConstructGuardians(&rogue.Character)

	return rogue
}

// Deactivate Stealth if it is active. This must be added to all abilities that cause Stealth to fade.
func (rogue *Rogue) BreakStealth(sim *core.Simulation) {
	if rogue.StealthAura.IsActive() {
		rogue.StealthAura.Deactivate(sim)
		rogue.AutoAttacks.EnableAutoSwing(sim)
	}
}

// Does the rogue have a dagger equipped in the specified hand (main or offhand)?
func (rogue *Rogue) HasDagger(hand core.Hand) bool {
	if hand == core.MainHand {
		return rogue.MainHand().WeaponType == proto.WeaponType_WeaponTypeDagger
	}
	return rogue.OffHand().WeaponType == proto.WeaponType_WeaponTypeDagger
}

// Check if the rogue is considered in "stealth" for the purpose of casting abilities
func (rogue *Rogue) IsStealthed() bool {
	return rogue.StealthAura.IsActive()
}

// Agent is a generic way to access underlying rogue on any of the agents.
type RogueAgent interface {
	GetRogue() *Rogue
}

func (rogue *Rogue) getImbueProcMask(imbue proto.WeaponImbue) core.ProcMask {
	var mask core.ProcMask
	if rogue.HasMHWeapon() && rogue.Consumes.MainHandImbue == imbue {
		mask |= core.ProcMaskMeleeMH
	}
	if rogue.HasOHWeapon() && rogue.Consumes.OffHandImbue == imbue {
		mask |= core.ProcMaskMeleeOH
	}
	return mask
}

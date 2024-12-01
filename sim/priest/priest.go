package priest

import (
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

var TalentTreeSizes = [3]int{15, 16, 16}

const (
	SpellFlagPriest = core.SpellFlagAgentReserved1
)

const (
	SpellCode_PriestNone int32 = iota

	SpellCode_PriestDevouringPlague
	SpellCode_PriestFlashHeal
	SpellCode_PriestGreaterHeal
	SpellCode_PriestHeal
	SpellCode_PriestHolyFire
	SpellCode_PriestMindBlast
	SpellCode_PriestMindFlay
	SpellCode_PriestShadowWordPain
	SpellCode_PriestSmite
	SpellCode_PriestVampiricTouch
)

type Priest struct {
	core.Character
	Talents *proto.PriestTalents

	Latency float64

	CircleOfHealing   *core.Spell
	DevouringPlague   []*core.Spell
	EmpoweredRenew    *core.Spell
	FlashHeal         []*core.Spell
	GreaterHeal       []*core.Spell
	HolyFire          []*core.Spell
	InnerFocus        *core.Spell
	MindBlast         []*core.Spell
	MindFlay          [][]*core.Spell // 1 entry for each tick for each rank
	PowerWordShield   []*core.Spell
	PrayerOfHealing   []*core.Spell
	PrayerOfMending   *core.Spell
	Renew             []*core.Spell
	Shadowform        *core.Spell
	ShadowWeavingProc *core.Spell
	ShadowWordPain    []*core.Spell
	Smite             []*core.Spell
	VampiricEmbrace   *core.Spell

	InnerFocusAura *core.Aura
	ShadowformAura *core.Aura
	SpiritTapAura  *core.Aura

	ShadowWeavingAuras   core.AuraArray
	VampiricEmbraceAuras core.AuraArray
	WeakenedSouls        core.AuraArray

	ProcPrayerOfMending core.ApplySpellResults
}

func (priest *Priest) GetCharacter() *core.Character {
	return &priest.Character
}

func (priest *Priest) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	raidBuffs.ShadowProtection = true
	raidBuffs.DivineSpirit = true
	raidBuffs.PowerWordFortitude = max(
		raidBuffs.PowerWordFortitude,
		core.MakeTristateValue(true, priest.Talents.ImprovedPowerWordFortitude == 2),
	)
}

func (priest *Priest) AddPartyBuffs(_ *proto.PartyBuffs) {
}

func (priest *Priest) Initialize() {
	priest.registerMindBlast()
	priest.registerMindFlay()
	priest.registerShadowWordPainSpell()
	if priest.GetCharacter().Race == proto.Race_RaceUndead {
		priest.registerDevouringPlagueSpell()
	}
	priest.RegisterSmiteSpell()
	priest.registerHolyFire()

	priest.registerPowerInfusionCD()
}

func (priest *Priest) RegisterHealingSpells() {
	// priest.registerFlashHealSpell()
	// priest.registerGreaterHealSpell()
	// priest.registerPowerWordShieldSpell()
	// priest.registerPrayerOfHealingSpell()
	// priest.registerRenewSpell()
}

func New(character *core.Character, talents string) *Priest {
	priest := &Priest{
		Character: *character,
		Talents:   &proto.PriestTalents{},
	}
	core.FillTalentsProto(priest.Talents.ProtoReflect(), talents, TalentTreeSizes)

	priest.EnableManaBar()

	priest.AddStatDependency(stats.Strength, stats.AttackPower, core.APPerStrength[character.Class])
	priest.AddStatDependency(stats.Intellect, stats.SpellCrit, core.CritPerIntAtLevel[priest.Class]*core.SpellCritRatingPerCritChance)

	// Set mana regen to 12.5 + Spirit/4 each 2s tick
	priest.SpiritManaRegenPerSecond = func() float64 {
		return 6.25 + priest.GetStat(stats.Spirit)/8
	}

	return priest
}

// Agent is a generic way to access underlying priest on any of the agents.
type PriestAgent interface {
	GetPriest() *Priest
}

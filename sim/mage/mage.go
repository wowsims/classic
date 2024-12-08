package mage

import (
	"github.com/wowsims/classic/sim/common/guardians"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

const (
	SpellFlagMage       = core.SpellFlagAgentReserved1
	SpellFlagChillSpell = core.SpellFlagAgentReserved2
)

const (
	SpellCode_MageNone int32 = iota
	SpellCode_MageArcaneExplosion
	SpellCode_MageArcaneMissiles
	SpellCode_MageArcaneMissilesTick
	SpellCode_MageBlastWave
	SpellCode_MageFireball
	SpellCode_MageFireBlast
	SpellCode_MageFlamestrike
	SpellCode_MageFrostbolt
	SpellCode_MageIgnite
	SpellCode_MageScorch
)

var TalentTreeSizes = [3]int{16, 16, 17}

func RegisterMage() {
	core.RegisterAgentFactory(
		proto.Player_Mage{},
		proto.Spec_SpecMage,
		func(character *core.Character, options *proto.Player) core.Agent {
			return NewMage(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Mage)
			if !ok {
				panic("Invalid spec value for Mage!")
			}
			player.Spec = playerSpec
		},
	)
}

type Mage struct {
	core.Character

	Talents *proto.MageTalents
	Options *proto.Mage_Options

	activeBarrier *core.Aura

	ArcaneExplosion         []*core.Spell
	ArcaneMissiles          []*core.Spell
	ArcaneMissilesTickSpell []*core.Spell
	BlastWave               []*core.Spell
	Blizzard                []*core.Spell
	Counterspell			*core.Spell
	Evocation				*core.Spell
	Fireball                []*core.Spell
	FireBlast               []*core.Spell
	Flamestrike             []*core.Spell
	Frostbolt               []*core.Spell
	IceBarrier              []*core.Spell
	Ignite                  *core.Spell
	ManaGem                 []*core.Spell
	PresenceOfMind          *core.Spell
	Pyroblast               []*core.Spell
	Scorch                  []*core.Spell

	ArcanePowerAura     *core.Aura
	ClearcastingAura    *core.Aura
	CombustionAura      *core.Aura
	IceArmorAura        *core.Aura
	IceBarrierAuras     []*core.Aura
	ImprovedScorchAuras core.AuraArray
	MageArmorAura       *core.Aura
}

// Agent is a generic way to access underlying mage on any of the agents.
type MageAgent interface {
	GetMage() *Mage
}

func (mage *Mage) GetCharacter() *core.Character {
	return &mage.Character
}

func (mage *Mage) GetMage() *Mage {
	return mage
}

func (mage *Mage) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
	raidBuffs.ArcaneBrilliance = true
}
func (mage *Mage) AddPartyBuffs(partyBuffs *proto.PartyBuffs) {
}

func (mage *Mage) Initialize() {
	mage.registerArcaneMissilesSpell()
	mage.registerFireballSpell()
	mage.registerFireBlastSpell()
	mage.registerFrostboltSpell()
	mage.registerPyroblastSpell()
	mage.registerScorchSpell()

	mage.registerArcaneExplosionSpell()
	mage.registerBlastWaveSpell()
	mage.registerBlizzardSpell()
	mage.registerFlamestrikeSpell()

	mage.registerEvocationCD()
	mage.registerManaGemCD()
	mage.registerCounterspellSpell()
}

func (mage *Mage) Reset(sim *core.Simulation) {
}

func NewMage(character *core.Character, options *proto.Player) *Mage {
	mageOptions := options.GetMage()

	mage := &Mage{
		Character: *character,
		Talents:   &proto.MageTalents{},
		Options:   mageOptions.Options,
	}
	core.FillTalentsProto(mage.Talents.ProtoReflect(), options.TalentsString, TalentTreeSizes)

	mage.EnableManaBar()

	mage.AddStatDependency(stats.Strength, stats.AttackPower, core.APPerStrength[character.Class])
	mage.AddStatDependency(stats.Intellect, stats.SpellCrit, core.CritPerIntAtLevel[mage.Class]*core.SpellCritRatingPerCritChance)

	switch mage.Options.Armor {
	case proto.Mage_Options_IceArmor:
		mage.applyFrostIceArmor()
	case proto.Mage_Options_MageArmor:
		mage.applyMageArmor()
	}

	// Set mana regen to 12.5 + Spirit/4 each 2s tick
	mage.SpiritManaRegenPerSecond = func() float64 {
		return 6.25 + mage.GetStat(stats.Spirit)/8
	}

	guardians.ConstructGuardians(&mage.Character)

	return mage
}

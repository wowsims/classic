package warrior

import (
	"time"

	"github.com/wowsims/classic/sim/common/guardians"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

const (
	SpellFlagOffensive = core.SpellFlagAgentReserved1
)

const (
	SpellCode_WarriorNone int32 = iota

	SpellCode_WarriorBloodthirst
	SpellCode_WarriorDeepWounds
	SpellCode_WarriorExecute
	SpellCode_WarriorMortalStrike
	SpellCode_WarriorOverpower
	SpellCode_WarriorRend
	SpellCode_WarriorRevenge
	SpellCode_WarriorShieldSlam
	SpellCode_WarriorSlam
	SpellCode_WarriorStanceBattle
	SpellCode_WarriorStanceBerserker
	SpellCode_WarriorStanceDefensive
	SpellCode_WarriorWhirlwind
)

var TalentTreeSizes = [3]int{18, 17, 17}

type WarriorInputs struct {
	StanceSnapshot bool
	Stance         proto.WarriorStance
}

const (
	ArmsTree = 0
	FuryTree = 1
	ProtTree = 2
)

type Warrior struct {
	core.Character

	Talents *proto.WarriorTalents

	WarriorInputs

	// Current state
	Stance          Stance
	PreviousStance  Stance // Used for Warrior T1 DPS 4P
	revengeProcAura *core.Aura
	OverpowerAura   *core.Aura

	lastMeleeAutoTarget *core.Unit

	// Enrage Auras
	BerserkerRageAura      *core.Aura
	BloodrageAura          *core.Aura
	EnrageAura             *core.Aura

	// Reaction time values
	reactionTime time.Duration
	LastAMTick   time.Duration

	BattleShout *WarriorSpell

	BattleStanceSpells    []*WarriorSpell
	DefensiveStanceSpells []*WarriorSpell
	BerserkerStanceSpells []*WarriorSpell

	Stances         []*WarriorSpell
	BattleStance    *WarriorSpell
	DefensiveStance *WarriorSpell
	BerserkerStance *WarriorSpell

	Bloodrage         *WarriorSpell
	BerserkerRage     *WarriorSpell
	Bloodthirst       *WarriorSpell
	DeathWish         *WarriorSpell
	DemoralizingShout *WarriorSpell
	Execute           *WarriorSpell
	MortalStrike      *WarriorSpell
	Overpower         *WarriorSpell
	Rend              *WarriorSpell
	Revenge           *WarriorSpell
	ShieldBlock       *WarriorSpell
	ShieldSlam        *WarriorSpell
	Slam              *WarriorSpell
	SunderArmor       *WarriorSpell
	Devastate         *WarriorSpell
	ThunderClap       *WarriorSpell
	Whirlwind         *WarriorSpell
	DeepWounds        *WarriorSpell
	ConcussionBlow    *WarriorSpell
	Hamstring         *WarriorSpell

	HeroicStrike       *WarriorSpell
	HeroicStrikeQueue  *WarriorSpell
	Cleave             *WarriorSpell
	CleaveQueue        *WarriorSpell
	curQueueAura       *core.Aura
	curQueuedAutoSpell *WarriorSpell

	BattleStanceAura    *core.Aura
	DefensiveStanceAura *core.Aura
	BerserkerStanceAura *core.Aura

	defensiveStanceThreatMultiplier float64

	ShieldBlockAura *core.Aura

	DemoralizingShoutAuras core.AuraArray
	SunderArmorAuras       core.AuraArray
	ThunderClapAuras       core.AuraArray
}

func (warrior *Warrior) GetCharacter() *core.Character {
	return &warrior.Character
}

func (warrior *Warrior) AddRaidBuffs(raidBuffs *proto.RaidBuffs) {
}

func (warrior *Warrior) AddPartyBuffs(_ *proto.PartyBuffs) {
}

func (warrior *Warrior) RegisterSpell(stanceMask Stance, config core.SpellConfig) *WarriorSpell {
	ws := &WarriorSpell{
		StanceMask: stanceMask,
	}

	castConditionOld := config.ExtraCastCondition
	config.ExtraCastCondition = func(sim *core.Simulation, target *core.Unit) bool {
		// Check if we're in a correct stance to cast the spell
		if stance := ws.GetStanceMask(); !ws.stanceOverride && stance != AnyStance && !warrior.StanceMatches(stance) {
			if sim.Log != nil {
				sim.Log("Failed cast to spell %s, wrong stance", ws.ActionID)
			}
			return false
		}
		return castConditionOld == nil || castConditionOld(sim, target)
	}

	ws.Spell = warrior.Unit.RegisterSpell(config)

	if stanceMask.Matches(BattleStance) {
		warrior.BattleStanceSpells = append(warrior.BattleStanceSpells, ws)
	}
	if stanceMask.Matches(DefensiveStance) {
		warrior.DefensiveStanceSpells = append(warrior.DefensiveStanceSpells, ws)
	}
	if stanceMask.Matches(BerserkerStance) {
		warrior.BerserkerStanceSpells = append(warrior.BerserkerStanceSpells, ws)
	}

	return ws
}

func (warrior *Warrior) Initialize() {
	primaryTimer := warrior.NewTimer()
	overpowerRevengeTimer := warrior.NewTimer()

	warrior.reactionTime = time.Millisecond * 500

	warrior.registerShouts()
	warrior.registerStances()
	warrior.registerBerserkerRageSpell()
	warrior.registerBloodthirstSpell(primaryTimer)
	warrior.registerDemoralizingShoutSpell()
	warrior.registerExecuteSpell()
	warrior.registerMortalStrikeSpell(primaryTimer)
	warrior.registerOverpowerSpell(overpowerRevengeTimer)
	warrior.registerRevengeSpell(overpowerRevengeTimer)
	warrior.registerShieldSlamSpell()
	warrior.registerSlamSpell()
	warrior.registerThunderClapSpell()
	warrior.registerWhirlwindSpell()
	warrior.registerRendSpell()
	warrior.registerHamstringSpell()

	// The sim often re-enables heroic strike in an unrealistic amount of time.
	// This can cause an unrealistic immediate double-hit around wild strikes procs
	queuedRealismICD := &core.Cooldown{
		Timer:    warrior.NewTimer(),
		Duration: core.SpellBatchWindow * 10,
	}
	warrior.registerHeroicStrikeSpell(queuedRealismICD)
	warrior.registerCleaveSpell(queuedRealismICD)

	warrior.registerSunderArmorSpell()

	warrior.registerBloodrageCD()
	warrior.RegisterRecklessnessCD()
}

func (warrior *Warrior) Reset(sim *core.Simulation) {
	warrior.curQueueAura = nil
	warrior.curQueuedAutoSpell = nil

	// Reset Stance
	switch warrior.WarriorInputs.Stance {
	case proto.WarriorStance_WarriorStanceBattle:
		warrior.Stance = BattleStance
		warrior.BattleStanceAura.Activate(sim)
	case proto.WarriorStance_WarriorStanceDefensive:
		warrior.Stance = DefensiveStance
		warrior.DefensiveStanceAura.Activate(sim)
	case proto.WarriorStance_WarriorStanceBerserker:
		warrior.Stance = BerserkerStance
		warrior.BerserkerStanceAura.Activate(sim)
	default:
		if warrior.PrimaryTalentTree == ArmsTree {
			warrior.Stance = BattleStance
			warrior.BattleStanceAura.Activate(sim)
		} else if warrior.PrimaryTalentTree == FuryTree {
			warrior.Stance = BerserkerStance
			warrior.BerserkerStanceAura.Activate(sim)
		} else {
			warrior.Stance = DefensiveStance
			warrior.DefensiveStanceAura.Activate(sim)
		}
	}
}

func NewWarrior(character *core.Character, talents string, inputs WarriorInputs) *Warrior {
	warrior := &Warrior{
		Character:     *character,
		Talents:       &proto.WarriorTalents{},
		WarriorInputs: inputs,
	}
	core.FillTalentsProto(warrior.Talents.ProtoReflect(), talents, TalentTreeSizes)

	warrior.PseudoStats.CanParry = true

	warrior.AddStatDependency(stats.Strength, stats.AttackPower, core.APPerStrength[character.Class])
	warrior.PseudoStats.BlockValuePerStrength = .05 // 20 str = 1 block
	warrior.AddStatDependency(stats.Agility, stats.MeleeCrit, core.CritPerAgiAtLevel[character.Class]*core.CritRatingPerCritChance)
	warrior.AddStatDependency(stats.Agility, stats.Dodge, core.DodgePerAgiAtLevel[character.Class]*core.DodgeRatingPerDodgeChance)
	warrior.AddStatDependency(stats.BonusArmor, stats.Armor, 1)

	guardians.ConstructGuardians(&warrior.Character)

	return warrior
}

// Agent is a generic way to access underlying warrior on any of the agents.
type WarriorAgent interface {
	GetWarrior() *Warrior
}

type WarriorSpell struct {
	*core.Spell
	StanceMask     Stance
	stanceOverride bool // Allows the override of the StanceMask so that the spell can be used in any stance
}

func (ws *WarriorSpell) IsReady(sim *core.Simulation) bool {
	if ws == nil {
		return false
	}
	return ws.Spell.IsReady(sim)
}

func (ws *WarriorSpell) CanCast(sim *core.Simulation, target *core.Unit) bool {
	if ws == nil {
		return false
	}
	return ws.Spell.CanCast(sim, target)
}

func (ws *WarriorSpell) IsEqual(s *core.Spell) bool {
	if ws == nil || s == nil {
		return false
	}
	return ws.Spell == s
}

// Returns the StanceMask accounting for a possible override
func (ws *WarriorSpell) GetStanceMask() Stance {
	if ws.stanceOverride {
		return AnyStance
	}

	return ws.StanceMask
}

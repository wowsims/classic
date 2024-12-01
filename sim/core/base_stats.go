package core

import (
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

type BaseStatsKey struct {
	Race  proto.Race
	Class proto.Class
	Level int
}

var BaseStats = map[BaseStatsKey]stats.Stats{}

// To calculate base stats, get a naked toon of desired level of the race/class you want, ideally without any talents to mess up base stats.
//  Basic stats are as-shown (str/agi/stm/int/spirit)

// Base Spell Crit is calculated by
//   1. Take as-shown value (troll shaman have 3.5%)
//   2. Calculate the bonus from int (for troll shaman that would be 104/78.1=1.331% crit)
//   3. Subtract as-shown from int bouns (3.5-1.331=2.169)
//   4. 2.169*22.08 (rating per crit percent) = 47.89 crit rating.

// Base mana can be looked up here: https://wowwiki-archive.fandom.com/wiki/Base_mana

// These are also scattered in various dbc/casc files,
// `octbasempbyclass.txt`, `combatratings.txt`, `chancetospellcritbase.txt`, etc.

var RaceOffsets = map[proto.Race]stats.Stats{
	proto.Race_RaceUnknown: {},
	proto.Race_RaceHuman:   {},
	proto.Race_RaceOrc: {
		stats.Agility:   -3,
		stats.Strength:  3,
		stats.Intellect: -3,
		stats.Spirit:    3,
		stats.Stamina:   2,
	},
	proto.Race_RaceDwarf: {
		stats.Agility:   -4,
		stats.Strength:  2,
		stats.Intellect: -1,
		stats.Spirit:    -1,
		stats.Stamina:   3,
	},
	proto.Race_RaceNightElf: {
		stats.Agility:   5,
		stats.Strength:  -3,
		stats.Intellect: 0,
		stats.Spirit:    0,
		stats.Stamina:   -1,
	},
	proto.Race_RaceUndead: {
		stats.Agility:   -2,
		stats.Strength:  -1,
		stats.Intellect: -2,
		stats.Spirit:    5,
		stats.Stamina:   1,
	},
	proto.Race_RaceTauren: {
		stats.Agility:   -5,
		stats.Strength:  5,
		stats.Intellect: -5,
		stats.Spirit:    2,
		stats.Stamina:   2,
	},
	proto.Race_RaceGnome: {
		stats.Agility:   3,
		stats.Strength:  -5,
		stats.Intellect: 3,
		stats.Spirit:    0,
		stats.Stamina:   -1,
	},
	proto.Race_RaceTroll: {
		stats.Agility:   2,
		stats.Strength:  1,
		stats.Intellect: -4,
		stats.Spirit:    1,
		stats.Stamina:   1,
	},
}

var ClassBaseCrit = map[proto.Class]stats.Stats{
	proto.Class_ClassUnknown: {},
	proto.Class_ClassWarrior: {
		stats.SpellCrit: 0.0000 * CritRatingPerCritChance,
		stats.MeleeCrit: 0.0000 * CritRatingPerCritChance,
		stats.Dodge:     0.0000 * DodgeRatingPerDodgeChance,
	},
	proto.Class_ClassPaladin: {
		stats.SpellCrit: 3.5000 * CritRatingPerCritChance,
		stats.MeleeCrit: 0.7000 * CritRatingPerCritChance,
		stats.Dodge:     0.7000 * DodgeRatingPerDodgeChance,
	},
	proto.Class_ClassHunter: {
		stats.SpellCrit: 3.6000 * CritRatingPerCritChance,
		stats.MeleeCrit: 0.0000 * CritRatingPerCritChance,
		stats.Dodge:     0.0000 * DodgeRatingPerDodgeChance,
	},
	proto.Class_ClassRogue: {
		stats.SpellCrit: 0.0000 * CritRatingPerCritChance,
		stats.MeleeCrit: 0.0000 * CritRatingPerCritChance,
		stats.Dodge:     0.0000 * DodgeRatingPerDodgeChance,
	},
	proto.Class_ClassPriest: {
		stats.SpellCrit: 0.8000 * CritRatingPerCritChance,
		stats.MeleeCrit: 3.0000 * CritRatingPerCritChance,
		stats.Dodge:     3.0000 * DodgeRatingPerDodgeChance,
	},
	proto.Class_ClassShaman: {
		stats.SpellCrit: 2.3000 * CritRatingPerCritChance,
		stats.MeleeCrit: 1.7000 * CritRatingPerCritChance,
		stats.Dodge:     1.7000 * DodgeRatingPerDodgeChance,
	},
	proto.Class_ClassMage: {
		stats.SpellCrit: 0.2000 * CritRatingPerCritChance,
		stats.MeleeCrit: 3.2000 * CritRatingPerCritChance,
		stats.Dodge:     3.2000 * DodgeRatingPerDodgeChance,
	},
	proto.Class_ClassWarlock: {
		stats.SpellCrit: 1.7000 * CritRatingPerCritChance,
		stats.MeleeCrit: 2.0000 * CritRatingPerCritChance,
		stats.Dodge:     2.0000 * DodgeRatingPerDodgeChance,
	},
	proto.Class_ClassDruid: {
		stats.SpellCrit: 1.8000 * CritRatingPerCritChance,
		stats.MeleeCrit: 0.9000 * CritRatingPerCritChance,
		stats.Dodge:     0.9000 * DodgeRatingPerDodgeChance,
	},
}

var APPerStrength = map[proto.Class]float64{
	proto.Class_ClassWarrior: 2,
	proto.Class_ClassPaladin: 2,
	proto.Class_ClassHunter:  1,
	proto.Class_ClassRogue:   1,
	proto.Class_ClassPriest:  1,
	proto.Class_ClassShaman:  2,
	proto.Class_ClassMage:    1,
	proto.Class_ClassWarlock: 1,
	proto.Class_ClassDruid:   2,
}

var APPerAgility = map[proto.Class]float64{
	proto.Class_ClassWarrior: 0,
	proto.Class_ClassPaladin: 0,
	proto.Class_ClassHunter:  0,
	proto.Class_ClassRogue:   1,
	proto.Class_ClassPriest:  0,
	proto.Class_ClassShaman:  0,
	proto.Class_ClassMage:    0,
	proto.Class_ClassWarlock: 0,
	proto.Class_ClassDruid:   1,
}

// Melee/Ranged crit agi scaling
var CritPerAgiAtLevel = map[proto.Class]float64{
	proto.Class_ClassUnknown: 0.0,
	proto.Class_ClassWarrior: 0.0500,
	proto.Class_ClassPaladin: 0.0506,
	proto.Class_ClassHunter:  0.0189,
	proto.Class_ClassRogue:   0.0345,
	proto.Class_ClassPriest:  0.0500,
	proto.Class_ClassShaman:  0.0508,
	proto.Class_ClassMage:    0.0514,
	proto.Class_ClassWarlock: 0.0500,
	proto.Class_ClassDruid:   0.0500,
}

// Spell crit int scaling
var CritPerIntAtLevel = map[proto.Class]float64{
	proto.Class_ClassUnknown: 0.0,
	proto.Class_ClassWarrior: 0.0,
	proto.Class_ClassPaladin: 0.0167,
	proto.Class_ClassHunter:  0.0165,
	proto.Class_ClassRogue:   0.0,
	proto.Class_ClassPriest:  0.0168,
	proto.Class_ClassShaman:  0.0169,
	proto.Class_ClassMage:    0.0168,
	proto.Class_ClassWarlock: 0.0165,
	proto.Class_ClassDruid:   0.0167,
}

// Dodge agility scaling
var DodgePerAgiAtLevel = map[proto.Class]float64{
	proto.Class_ClassUnknown: 0.0,
	proto.Class_ClassWarrior: 0.0500,
	proto.Class_ClassPaladin: 0.0506,
	proto.Class_ClassHunter:  0.0378,
	proto.Class_ClassRogue:   0.0690,
	proto.Class_ClassPriest:  0.0500,
	proto.Class_ClassShaman:  0.0508,
	proto.Class_ClassMage:    0.0514,
	proto.Class_ClassWarlock: 0.0500,
	proto.Class_ClassDruid:   0.0500,
}

var ClassBaseStats = map[proto.Class]stats.Stats{
	proto.Class_ClassUnknown: {},
	proto.Class_ClassWarrior: {
		stats.Health:      1689,
		stats.Mana:        0,
		stats.Agility:     80,
		stats.Strength:    120,
		stats.Intellect:   30,
		stats.Spirit:      45,
		stats.Stamina:     110,
		stats.AttackPower: 60*3 - 20,
	},
	proto.Class_ClassPaladin: {
		stats.Health:      1381,
		stats.Mana:        1512,
		stats.Agility:     65,
		stats.Strength:    105,
		stats.Intellect:   70,
		stats.Spirit:      75,
		stats.Stamina:     100,
		stats.AttackPower: 60*3 - 20,
	},
	proto.Class_ClassHunter: {
		stats.Health:            1467,
		stats.Mana:              1720,
		stats.Agility:           125,
		stats.Strength:          55,
		stats.Intellect:         65,
		stats.Spirit:            70,
		stats.Stamina:           90,
		stats.AttackPower:       60*2 - 20,
		stats.RangedAttackPower: 60*2 - 20,
	},
	proto.Class_ClassRogue: {
		stats.Health:      1523,
		stats.Mana:        0,
		stats.Agility:     130,
		stats.Strength:    80,
		stats.Intellect:   35,
		stats.Spirit:      50,
		stats.Stamina:     75,
		stats.AttackPower: 60*2 - 20,
	},
	proto.Class_ClassPriest: {
		stats.Health:      1397,
		stats.Mana:        1376,
		stats.Agility:     40,
		stats.Strength:    35,
		stats.Intellect:   120,
		stats.Spirit:      125,
		stats.Stamina:     50,
		stats.AttackPower: -10,
	},
	proto.Class_ClassShaman: {
		stats.Health:      1423,
		stats.Mana:        1520,
		stats.Agility:     55,
		stats.Strength:    85,
		stats.Intellect:   90,
		stats.Spirit:      100,
		stats.Stamina:     95,
		stats.AttackPower: 60*2 - 20,
	},
	proto.Class_ClassMage: {
		stats.Health:      1370,
		stats.Mana:        1213,
		stats.Agility:     35,
		stats.Strength:    30,
		stats.Intellect:   125,
		stats.Spirit:      120,
		stats.Stamina:     45,
		stats.AttackPower: -10,
	},
	proto.Class_ClassWarlock: {
		stats.Health:      1414,
		stats.Mana:        1373,
		stats.Agility:     50,
		stats.Strength:    45,
		stats.Intellect:   110,
		stats.Spirit:      115,
		stats.Stamina:     65,
		stats.AttackPower: -10,
	},
	proto.Class_ClassDruid: {
		stats.Health:      1483,
		stats.Mana:        1244,
		stats.Agility:     60,
		stats.Strength:    65,
		stats.Intellect:   100,
		stats.Spirit:      110,
		stats.Stamina:     70,
		stats.AttackPower: -20,
	},
}

// Retrieves base stats, with race offsets, and crit rating adjustments per level
func getBaseStatsCombo(r proto.Race, c proto.Class) stats.Stats {
	starting := ClassBaseStats[c]
	return starting.Add(RaceOffsets[r]).Add(ClassBaseCrit[c])
}

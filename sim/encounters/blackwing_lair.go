package encounters

import (
	"time"

	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/core/stats"
)

// essence of the red updated for classic

func addVaelastraszTheCorrupt(bossPrefix string) {
	core.AddPresetTarget(&core.PresetTarget{
		PathPrefix: bossPrefix,
		Config: &proto.Target{
			Id:        13020,
			Name:      "Blackwing Lair Vaelastrasz the Corrupt",
			Level:     63,
			MobType:   proto.MobType_MobTypeDragonkin,
			TankIndex: 0,

			Stats: stats.Stats{
				stats.Health:      3_331_000,
				stats.Armor:       3731, // TODO:
				stats.AttackPower: 805,  // TODO: Unknown attack power
				// TODO: Resistances
			}.ToFloatArray(),

			SpellSchool:      proto.SpellSchool_SpellSchoolPhysical,
			SwingSpeed:       2,     // TODO: Very slow attack interrupted by spells
			MinBaseDamage:    5000,  // TODO: Minimum unmitigated damage on reviewed log
			DamageSpread:     0.333, // TODO:
			ParryHaste:       true,
			DualWield:        false,
			DualWieldPenalty: false,
		},
		AI: NewVaelastraszTheCorruptAI(),
	})
	core.AddPresetEncounter("Blackwing Lair Vaelastrasz the Corrupt", []string{
		bossPrefix + "/Blackwing Lair Vaelastrasz the Corrupt",
	})
}

type VaelastraszTheCorruptAI struct {
	Target               *core.Target
	essenceOfTheRedSpell *core.Spell
	canAct               bool
}

func NewVaelastraszTheCorruptAI() core.AIFactory {
	return func() core.TargetAI {
		return &VaelastraszTheCorruptAI{}
	}
}

func (ai *VaelastraszTheCorruptAI) Initialize(target *core.Target, config *proto.Target) {
	ai.Target = target
	ai.registerSpells()
	ai.canAct = true
}

func (ai *VaelastraszTheCorruptAI) registerSpells() {
	essenceOfTheRedActionID := core.ActionID{SpellID: 23513}
	//burningAdrenalineActionID := core.ActionID{SpellID: 367987}
	//burningAdrenalineTankActionID := core.ActionID{SpellID: 469261}
	//fireNovaActionID := core.ActionID{SpellID: 23462}
	//flameBreathActionID := core.ActionID{SpellID: 23461}
	//cleaveActionID := core.ActionID{SpellID: 19983}

	target := &ai.Target.Env.Raid.Parties[0].Players[0].GetCharacter().Unit

	essenceOfTheRedManaMetrics := target.NewManaMetrics(essenceOfTheRedActionID)
	essenceOfTheRedEnergyMetrics := target.NewEnergyMetrics(essenceOfTheRedActionID)
	essenceOfTheRedRageMetrics := target.NewRageMetrics(essenceOfTheRedActionID)

	ai.essenceOfTheRedSpell = ai.Target.RegisterSpell(core.SpellConfig{
		ActionID: essenceOfTheRedActionID,
		ProcMask: core.ProcMaskEmpty,

		Cast: core.CastConfig{
			CD: core.Cooldown{
				Timer:    ai.Target.NewTimer(),
				Duration: time.Minute * 4,
			},
		},
		Dot: core.DotConfig{
			Aura: core.Aura{
				Label: "Essemce of the Red",
			},
			NumberOfTicks: 240,
			TickLength:    time.Second * 1,

			OnTick: func(sim *core.Simulation, target *core.Unit, dot *core.Dot) {
				if target.HasManaBar() {
					target.AddMana(sim, 500, essenceOfTheRedManaMetrics)
				}
				if target.HasEnergyBar() {
					target.AddEnergy(sim, 50, essenceOfTheRedEnergyMetrics)
				}
				if target.HasRageBar() {
					target.AddRage(sim, 20, essenceOfTheRedRageMetrics)
				}
			},
		},
		ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
			spell.Dot(target).Apply(sim)
		},
	})
}

func (ai *VaelastraszTheCorruptAI) Reset(*core.Simulation) {
}

const BossGCD = time.Millisecond * 1600

func (ai *VaelastraszTheCorruptAI) ExecuteCustomRotation(sim *core.Simulation) {
	if !ai.canAct {
		ai.Target.WaitUntil(sim, sim.CurrentTime+BossGCD)
		return
	}

	target := ai.Target.CurrentTarget

	if target == nil {
		// For individual non tank sims we still want abilities to work
		target = &ai.Target.Env.Raid.Parties[0].Players[0].GetCharacter().Unit
	}

	if ai.essenceOfTheRedSpell.CanCast(sim, target) {
		ai.essenceOfTheRedSpell.Cast(sim, target)
		ai.Target.WaitUntil(sim, sim.CurrentTime+BossGCD)
		return
	}
}

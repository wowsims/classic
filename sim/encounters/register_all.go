package encounters

import (
	"github.com/wowsims/classic/sim/core"
)

func init() {
	// TODO: Classic encounters?
	// naxxramas.Register()
	addLevel60("SoD")
	addVaelastraszTheCorrupt("SoD")
}

func AddSingleTargetBossEncounter(presetTarget *core.PresetTarget) {
	core.AddPresetTarget(presetTarget)
	core.AddPresetEncounter(presetTarget.Config.Name, []string{
		presetTarget.Path(),
	})
}

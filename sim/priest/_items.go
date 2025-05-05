package priest

import (

	"github.com/wowsims/classic/sim/core"
)

const (
	// Keep these ordered by ID
	AtieshPriest	= 22631
)

func init() {
	core.AddEffectsToTest = false

	// Keep these ordered by name

	// https://www.wowhead.com/classic/item=22631/atiesh-greatstaff-of-the-guardian
	core.NewItemEffect(AtieshPriest, func(agent core.Agent) {
		character := agent.GetCharacter()
		aura := core.AtieshHealingEffect(&character.Unit)
		character.ItemSwap.RegisterProc(AtieshPriest, aura)
	})
	
	core.AddEffectsToTest = true
}

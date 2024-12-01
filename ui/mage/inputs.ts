import * as InputHelpers from '../core/components/input_helpers.js';
import { Spec } from '../core/proto/common.js';
import { Mage_Options_ArmorType as ArmorType } from '../core/proto/mage.js';
import { ActionId } from '../core/proto_utils/action_id.js';
import { TypedEvent } from '../core/typed_event.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const Armor = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecMage, ArmorType>({
	fieldName: 'armor',
	values: [
		{ value: ArmorType.NoArmor, tooltip: 'No Armor' },
		{
			actionId: () => ActionId.fromSpellId(22783),
			value: ArmorType.MageArmor,
		},
		{
			actionId: () => ActionId.fromSpellId(10220),
			value: ArmorType.IceArmor,
		},
	],
	changeEmitter: player => TypedEvent.onAny([player.gearChangeEmitter, player.specOptionsChangeEmitter]),
});

export const MageRotationConfig = {
	inputs: [],
};

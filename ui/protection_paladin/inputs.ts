import * as InputHelpers from '../core/components/input_helpers.js';
import { Player } from '../core/player.js';
import { ItemSlot, Spec } from '../core/proto/common.js';
import { Blessings,PaladinAura, PaladinRune, PaladinSeal } from '../core/proto/paladin.js';
import { ActionId } from '../core/proto_utils/action_id.js';
import { TypedEvent } from '../core/typed_event.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const AuraSelection = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecProtectionPaladin, PaladinAura>({
	fieldName: 'aura',
	values: [
		{ value: PaladinAura.NoPaladinAura, tooltip: 'No Aura' },
		{ actionId: () => ActionId.fromSpellId(20218), value: PaladinAura.SanctityAura },
		//{ actionId: () => ActionId.fromSpellId(10299), value: PaladinAura.DevotionAura },
		//{ actionId: () => ActionId.fromSpellId(10299), value: PaladinAura.RetributionAura },
		//{ actionId: () => ActionId.fromSpellId(19746), value: PaladinAura.ConcentrationAura },
		//{ actionId: () => ActionId.fromSpellId(19888), value: PaladinAura.FrostResistanceAura },
		//{ actionId: () => ActionId.fromSpellId(19892), value: PaladinAura.ShadowResistanceAura },
		//{ actionId: () => ActionId.fromSpellId(19891), value: PaladinAura.FireResistanceAura },
	],
});

export const BlessingSelection = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecProtectionPaladin, Blessings>({
	fieldName: 'personalBlessing',
	values: [
		{ value: Blessings.BlessingUnknown, tooltip: 'No Blessing' },
		{
			actionId: () => ActionId.fromSpellId(20914),
			value: Blessings.BlessingOfSanctuary,
		},
	],
	changeEmitter: (player: Player<Spec.SpecProtectionPaladin>) => TypedEvent.onAny([player.specOptionsChangeEmitter]),
});

export const RighteousFuryToggle = InputHelpers.makeSpecOptionsBooleanIconInput<Spec.SpecProtectionPaladin>({
	fieldName: 'righteousFury',
	actionId: (_player: Player<Spec.SpecProtectionPaladin>) => ActionId.fromSpellId(25780),
	changeEmitter: (player: Player<Spec.SpecProtectionPaladin>) => TypedEvent.onAny([player.gearChangeEmitter, player.specOptionsChangeEmitter]),
});

// The below is used in the custom APL action "Cast Primary Seal".
// Only shows SoC if it's talented.
export const PrimarySealSelection = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecProtectionPaladin, PaladinSeal>({
	fieldName: 'primarySeal',
	values: [
		{
			actionId: () => ActionId.fromSpellId(20293),
			value: PaladinSeal.Righteousness,
		},
		{
			actionId: () => ActionId.fromSpellId(20920),
			value: PaladinSeal.Command,
			showWhen: (player: Player<Spec.SpecProtectionPaladin>) => player.getTalents().sealOfCommand,
		},
		{
			actionId: () => ActionId.fromSpellId(407798),
			value: PaladinSeal.Martyrdom,
		},
	],
	changeEmitter: (player: Player<Spec.SpecProtectionPaladin>) =>
		TypedEvent.onAny([player.gearChangeEmitter, player.talentsChangeEmitter, player.specOptionsChangeEmitter]),
});

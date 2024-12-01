import * as InputHelpers from '../core/components/input_helpers.js';
import { Player } from '../core/player.js';
import { Spec } from '../core/proto/common.js';
import { PaladinAura,PaladinSeal } from '../core/proto/paladin.js';
import { ActionId } from '../core/proto_utils/action_id.js';
import { TypedEvent } from '../core/typed_event.js';

// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.

export const AuraSelection = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecRetributionPaladin, PaladinAura>({
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

// The below is used in the custom APL action "Cast Primary Seal".
// Only shows SoC if it's talented, only shows SoM if the relevant rune is equipped.
export const PrimarySealSelection = InputHelpers.makeSpecOptionsEnumIconInput<Spec.SpecRetributionPaladin, PaladinSeal>({
	fieldName: 'primarySeal',
	values: [
		{
			actionId: () => ActionId.fromSpellId(20293),
			value: PaladinSeal.Righteousness,
		},
		{
			actionId: () => ActionId.fromSpellId(20920),
			value: PaladinSeal.Command,
			showWhen: (player: Player<Spec.SpecRetributionPaladin>) => player.getTalents().sealOfCommand,
		},
		{
			actionId: () => ActionId.fromSpellId(407798),
			value: PaladinSeal.Martyrdom,
		},
	],
	// changeEmitter: (player: Player<Spec.SpecRetributionPaladin>) => player.changeEmitter,
	changeEmitter: (player: Player<Spec.SpecRetributionPaladin>) =>
		TypedEvent.onAny([player.gearChangeEmitter, player.talentsChangeEmitter, player.specOptionsChangeEmitter]),
});

export const CrusaderStrikeStopAttack = InputHelpers.makeSpecOptionsBooleanInput<Spec.SpecRetributionPaladin>({
	fieldName: 'isUsingCrusaderStrikeStopAttack',
	label: 'Using Crusader Strike StopAttack Macro',
	labelTooltip: 'Allows saving of extra attacks',
});

export const DivineStormStopAttack = InputHelpers.makeSpecOptionsBooleanInput<Spec.SpecRetributionPaladin>({
	fieldName: 'isUsingDivineStormStopAttack',
	label: 'Using Divine Storm StopAttack Macro',
	labelTooltip: 'Allows saving of extra attacks',
});

export const JudgementStopAttack = InputHelpers.makeSpecOptionsBooleanInput<Spec.SpecRetributionPaladin>({
	fieldName: 'isUsingJudgementStopAttack',
	label: 'Using Judgement StopAttack Macro',
	labelTooltip: 'Allows saving of extra attacks',
});

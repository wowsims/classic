import { IconEnumPicker } from '../components/icon_enum_picker.js';
import { IndividualSimUI } from '../individual_sim_ui.js';
import { Player } from '../player.js';
import {
	AirTotem,
	EarthTotem,
	FireTotem,
	WaterTotem,
	ShamanTotems,
} from '../proto/shaman.js';
import { ActionId } from '../proto_utils/action_id.js';
import { ShamanSpecs } from '../proto_utils/utils.js';
import { EventID, TypedEvent } from '../typed_event.js';

import { ContentBlock } from './content_block.js';
import { Input } from './input.js';
import { StoneskinTotemInputs, StrengthOfEarthTotemInputs, TremorTotemInput } from './inputs/earth_totems.js';

export function TotemsSection(parentElem: HTMLElement, simUI: IndividualSimUI<ShamanSpecs>): ContentBlock {
	let contentBlock = new ContentBlock(parentElem, 'totems-settings', {
		header: { title: 'Totems' }
	});

	let totemDropdownGroup = Input.newGroupContainer();
	totemDropdownGroup.classList.add('totem-dropdowns-container', 'icon-group');

	contentBlock.bodyElement.appendChild(totemDropdownGroup);

	new IconEnumPicker(totemDropdownGroup, simUI.player, {
		extraCssClasses: [
			'earth-totem-picker',
		],
		numColumns: 1,
		values: [
			{ color: '#ffdfba', value: EarthTotem.NoEarthTotem },
			...StrengthOfEarthTotemInputs,
			...StoneskinTotemInputs,
			TremorTotemInput,
		],
		equals: (a: EarthTotem, b: EarthTotem) => a == b,
		zeroValue: EarthTotem.NoEarthTotem,
		changedEvent: (player: Player<ShamanSpecs>) => TypedEvent.onAny([player.specOptionsChangeEmitter, player.levelChangeEmitter]),
		getValue: (player: Player<ShamanSpecs>) => player.getSpecOptions().totems?.earth || EarthTotem.NoEarthTotem,
		setValue: (eventID: EventID, player: Player<ShamanSpecs>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			if (!newOptions.totems)
				newOptions.totems = ShamanTotems.create();
			newOptions.totems!.earth = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	});

	new IconEnumPicker(totemDropdownGroup, simUI.player, {
		extraCssClasses: [
			'fire-totem-picker',
		],
		numColumns: 1,
		values: [
			{ color: '#ffb3ba', value: FireTotem.NoFireTotem },
			{ actionId: () => ActionId.fromSpellId(58734), value: FireTotem.MagmaTotem },
			{ actionId: () => ActionId.fromSpellId(58704), value: FireTotem.SearingTotem },
			{ actionId: () => ActionId.fromSpellId(58656), value: FireTotem.FlametongueTotem },
		],
		equals: (a: FireTotem, b: FireTotem) => a == b,
		zeroValue: FireTotem.NoFireTotem,
		changedEvent: (player: Player<ShamanSpecs>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<ShamanSpecs>) => player.getSpecOptions().totems?.fire || FireTotem.NoFireTotem,
		setValue: (eventID: EventID, player: Player<ShamanSpecs>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			if (!newOptions.totems)
				newOptions.totems = ShamanTotems.create();
			newOptions.totems!.fire = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	});

	new IconEnumPicker(totemDropdownGroup, simUI.player, {
		extraCssClasses: [
			'water-totem-picker',
		],
		numColumns: 1,
		values: [
			{ color: '#bae1ff', value: WaterTotem.NoWaterTotem },
			{ actionId: () => ActionId.fromSpellId(58774), value: WaterTotem.ManaSpringTotem },
			{ actionId: () => ActionId.fromSpellId(58757), value: WaterTotem.HealingStreamTotem },
		],
		equals: (a: WaterTotem, b: WaterTotem) => a == b,
		zeroValue: WaterTotem.NoWaterTotem,
		changedEvent: (player: Player<ShamanSpecs>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<ShamanSpecs>) => player.getSpecOptions().totems?.water || WaterTotem.NoWaterTotem,
		setValue: (eventID: EventID, player: Player<ShamanSpecs>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			if (!newOptions.totems)
				newOptions.totems = ShamanTotems.create();
			newOptions.totems!.water = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	});

	new IconEnumPicker(totemDropdownGroup, simUI.player, {
		extraCssClasses: [
			'air-totem-picker',
		],
		numColumns: 1,
		values: [
			{ color: '#baffc9', value: AirTotem.NoAirTotem },
			{ actionId: () => ActionId.fromSpellId(8512), value: AirTotem.WindfuryTotem },
			{ actionId: () => ActionId.fromSpellId(8835), value: AirTotem.GraceOfAirTotem },
		],
		equals: (a: AirTotem, b: AirTotem) => a == b,
		zeroValue: AirTotem.NoAirTotem,
		changedEvent: (player: Player<ShamanSpecs>) => player.specOptionsChangeEmitter,
		getValue: (player: Player<ShamanSpecs>) => player.getSpecOptions().totems?.air || AirTotem.NoAirTotem,
		setValue: (eventID: EventID, player: Player<ShamanSpecs>, newValue: number) => {
			const newOptions = player.getSpecOptions();
			if (!newOptions.totems)
				newOptions.totems = ShamanTotems.create();
			newOptions.totems!.air = newValue;
			player.setSpecOptions(eventID, newOptions);
		},
	});

	return contentBlock;
}

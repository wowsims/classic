import { Player } from '../../player.js';
import { ItemSlot } from '../../proto/common.js';
import {
	WarlockOptions_Armor as Armor,
	WarlockOptions_MaxFireboltRank as MaxFireboltRank,
	WarlockOptions_Summon as Summon,
	WarlockOptions_WeaponImbue as WeaponImbue,
} from '../../proto/warlock.js';
import { ActionId } from '../../proto_utils/action_id.js';
import { WarlockSpecs } from '../../proto_utils/utils.js';
import * as InputHelpers from '../input_helpers.js';

export const ArmorInput = <SpecType extends WarlockSpecs>() =>
	InputHelpers.makeSpecOptionsEnumIconInput<SpecType, Armor>({
		fieldName: 'armor',
		values: [
			{ value: Armor.NoArmor, tooltip: 'No Armor' },
			{
				actionId: () => ActionId.fromSpellId(11735),
				value: Armor.DemonArmor,
			},
			{
				actionId: () => ActionId.fromSpellId(403619),
				value: Armor.DemonArmor,
			},
		],
	});

export const WeaponImbueInput = <SpecType extends WarlockSpecs>() =>
	InputHelpers.makeSpecOptionsEnumIconInput<SpecType, WeaponImbue>({
		fieldName: 'weaponImbue',
		values: [
			{ value: WeaponImbue.NoWeaponImbue, tooltip: 'No Weapon Stone' },
			{
				actionId: () => ActionId.fromSpellId(13701),
				value: WeaponImbue.Firestone,
			},
			{
				actionId: () => ActionId.fromSpellId(13603),
				value: WeaponImbue.Spellstone,
			},
		],
		showWhen: player => player.getEquippedItem(ItemSlot.ItemSlotOffHand) == null,
		changeEmitter: (player: Player<SpecType>) => player.changeEmitter,
	});

export const PetInput = <SpecType extends WarlockSpecs>() =>
	InputHelpers.makeSpecOptionsEnumIconInput<SpecType, Summon>({
		fieldName: 'summon',
		values: [
			{ value: Summon.NoSummon, tooltip: 'No Pet' },
			{ actionId: () => ActionId.fromSpellId(688), value: Summon.Imp },
			{ actionId: () => ActionId.fromSpellId(697), value: Summon.Voidwalker },
			{ actionId: () => ActionId.fromSpellId(712), value: Summon.Succubus },
			{ actionId: () => ActionId.fromSpellId(691), value: Summon.Felhunter },
		],
		changeEmitter: (player: Player<SpecType>) => player.changeEmitter,
	});

export const ImpFireboltRank = <SpecType extends WarlockSpecs>() =>
	InputHelpers.makeSpecOptionsEnumIconInput<SpecType, MaxFireboltRank>({
		fieldName: 'maxFireboltRank',
		showWhen: player => player.getSpecOptions().summon == Summon.Imp,
		values: [
			{ value: MaxFireboltRank.NoMaximum, tooltip: 'Max' },
			{ actionId: () => ActionId.fromSpellId(3110), value: MaxFireboltRank.Rank1 },
			{ actionId: () => ActionId.fromSpellId(7799), value: MaxFireboltRank.Rank2 },
			{ actionId: () => ActionId.fromSpellId(7800), value: MaxFireboltRank.Rank3 },
			{ actionId: () => ActionId.fromSpellId(7801), value: MaxFireboltRank.Rank4 },
			{ actionId: () => ActionId.fromSpellId(7802), value: MaxFireboltRank.Rank5 },
			{ actionId: () => ActionId.fromSpellId(11762), value: MaxFireboltRank.Rank6 },
			{ actionId: () => ActionId.fromSpellId(11763), value: MaxFireboltRank.Rank7 },
		],
		changeEmitter: (player: Player<SpecType>) => player.changeEmitter,
	});

export const PetPoolManaInput = <SpecType extends WarlockSpecs>() =>
	InputHelpers.makeSpecOptionsBooleanInput<SpecType>({
		fieldName: 'petPoolMana',
		label: 'No Pet Management',
		labelTooltip: 'Should Pet keep trying to cast on every mana regen instead of waiting for mana',
		changeEmitter: (player: Player<SpecType>) => player.changeEmitter,
	});

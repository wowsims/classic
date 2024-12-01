import { Class, ItemSlot, WeaponImbue } from '../../proto/common.js';
import { ActionId } from '../../proto_utils/action_id';
import { isWeapon } from '../../proto_utils/utils';
import { ConsumableInputConfig } from './consumables';

// Shaman Imbues
export const RockbiterWeaponImbue = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromSpellId(16316),
		value: WeaponImbue.RockbiterWeapon,
		showWhen: player => {
			if (!player.isClass(Class.ClassShaman)) return false;

			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};

export const FlametongueWeaponImbue = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromSpellId(16342),
		value: WeaponImbue.FlametongueWeapon,
		showWhen: player => {
			if (!player.isClass(Class.ClassShaman)) return false;
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};

export const FrostbrandWeaponImbue = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromSpellId(16356),
		value: WeaponImbue.FrostbrandWeapon,
		showWhen: player => {
			if (!player.isClass(Class.ClassShaman)) return false;
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};

export const WindfuryWeaponImbue = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromSpellId(16362),
		value: WeaponImbue.WindfuryWeapon,
		showWhen: player => {
			if (!player.isClass(Class.ClassShaman)) return false;
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};

import { Class, WeaponImbue } from '../../proto/common.js';
import { ActionId } from '../../proto_utils/action_id';
import { ConsumableInputConfig } from './consumables';

// Rogue Imbues
export const InstantPoisonWeaponImbue: ConsumableInputConfig<WeaponImbue> = {
	actionId: () => ActionId.fromItemId(8928),
	value: WeaponImbue.InstantPoison,
	showWhen: player => player.getClass() == Class.ClassRogue,
};

export const DeadlyPoisonWeaponRank4Imbue: ConsumableInputConfig<WeaponImbue> = {
	actionId: () => ActionId.fromItemId(8985),
	value: WeaponImbue.DeadlyPoisonRank4,
	// value: WeaponImbue.DeadlyPoisonRank4,
	showWhen: player => player.getClass() == Class.ClassRogue,
};

export const DeadlyPoisonWeaponRank5Imbue: ConsumableInputConfig<WeaponImbue> = {
	actionId: () => ActionId.fromItemId(20844),
	value: WeaponImbue.DeadlyPoisonRank5,
	// value: WeaponImbue.DeadlyPoisonRank5,
	showWhen: player => player.getClass() == Class.ClassRogue,
};

export const WoundPoisonWeaponImbue: ConsumableInputConfig<WeaponImbue> = {
	actionId: () => ActionId.fromItemId(10922),
	value: WeaponImbue.WoundPoison,
	showWhen: player => player.getClass() == Class.ClassRogue,
};

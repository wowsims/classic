import { AirTotem, EarthTotem, FireTotem, WaterTotem } from '../../proto/shaman.js';
import { ActionId } from '../../proto_utils/action_id';

///////////////////////////////////////////////////////////////////////////
//                                 Earth Totems
///////////////////////////////////////////////////////////////////////////

export const StoneskinTotem = {
	actionId: () => ActionId.fromSpellId(10408),
	value: EarthTotem.StoneskinTotem,
};

export const StrengthOfEarthTotem = {
	actionId: () => ActionId.fromSpellId(25361),
	value: EarthTotem.StrengthOfEarthTotem,
};

export const TremorTotem = {
	actionId: () => ActionId.fromSpellId(8143),
	value: EarthTotem.TremorTotem,
};

///////////////////////////////////////////////////////////////////////////
//                                 Fire Totems
///////////////////////////////////////////////////////////////////////////

export const SearingTotem = {
	actionId: () => ActionId.fromSpellId(10438),
	value: FireTotem.SearingTotem,
};

export const FireNovaTotem = {
	actionId: () => ActionId.fromSpellId(11315),
	value: FireTotem.FireNovaTotem,
};

export const MagmaTotem = {
	actionId: () => ActionId.fromSpellId(10587),
	value: FireTotem.FireNovaTotem,
};

///////////////////////////////////////////////////////////////////////////
//                                 Water Totems
///////////////////////////////////////////////////////////////////////////

export const HealingStreamTotem = {
	actionId: () => ActionId.fromSpellId(10463),
	value: WaterTotem.HealingStreamTotem,
};

export const ManaSpringTotem = {
	actionId: () => ActionId.fromSpellId(10497),
	value: WaterTotem.ManaSpringTotem,
};

///////////////////////////////////////////////////////////////////////////
//                                 Air Totems
///////////////////////////////////////////////////////////////////////////

export const WindfuryTotem = {
	actionId: () => ActionId.fromSpellId(25359),
	value: AirTotem.WindfuryTotem,
};

export const GraceOfAirTotem = {
	actionId: () => ActionId.fromSpellId(25359),
	value: AirTotem.GraceOfAirTotem,
};

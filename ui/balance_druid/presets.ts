import { Phase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	Consumes,
	Debuffs,
	Flask,
	Food,
	IndividualBuffs,
	ManaRegenElixir,
	PartyBuffs,
	Potions,
	Profession,
	RaidBuffs,
	SaygesFortune,
	SpellPowerBuff,
	TristateEffect,
	UnitReference,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { BalanceDruid_Options as BalanceDruidOptions } from '../core/proto/druid.js';
import { SavedTalents } from '../core/proto/ui.js';
import Balance from './apls/balance.apl.json';
import P0BISGear from './gear_sets/p0.bis.gear.json';
import P1BISGear from './gear_sets/p1.bis.gear.json';
import P2BISGear from './gear_sets/p2.bis.gear.json';
import P3BISGear from './gear_sets/p3.bis.gear.json';
import P4BISGear from './gear_sets/p4.bis.gear.json';
import P5BISGear from './gear_sets/p5.bis.gear.json';
// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearP0BIS = PresetUtils.makePresetGear('Pre-BiS', P0BISGear);
export const GearP1BIS = PresetUtils.makePresetGear('P1 BiS', P1BISGear);
export const GearP2BIS = PresetUtils.makePresetGear('P2 BiS', P2BISGear);
export const GearP3BIS = PresetUtils.makePresetGear('P3 BiS', P3BISGear);
export const GearP4BIS = PresetUtils.makePresetGear('P4 BiS', P4BISGear);
export const GearP5BIS = PresetUtils.makePresetGear('P5 BiS', P5BISGear);


export const GearPresets = {
	[Phase.Phase5]: [GearP0BIS, GearP1BIS, GearP2BIS, GearP3BIS, GearP4BIS, GearP5BIS],
};

export const DefaultGear = GearP5BIS;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

// export const APLP1Balance = PresetUtils.makePresetAPLRotation('Balance', P1APL);
export const DefaultBalance = PresetUtils.makePresetAPLRotation('DefaultBalance', Balance);
export const APLPresets = {
	[Phase.Phase1]: [DefaultBalance],
};

export const DefaultAPL = APLPresets[Phase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

export const TalentsP1Balance = PresetUtils.makePresetTalents('Balance', SavedTalents.create({ talentsString: '5000550012551251--5005031' }));

export const TalentPresets = {
	[Phase.Phase1]: [TalentsP1Balance],
};

export const DefaultTalents = TalentPresets[Phase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = BalanceDruidOptions.create({
	innervateTarget: UnitReference.create(),
});

export const DefaultConsumes = Consumes.create({
	defaultPotion: Potions.MajorManaPotion,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodRunnTumTuberSurprise,
	mainHandImbue: WeaponImbue.BrilliantWizardOil,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,

	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	zanzaBuff: ZanzaBuff.CerebralCortexCompound,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	moonkinAura: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	rallyingCryOfTheDragonslayer: true,
	saygesFortune: SaygesFortune.SaygesDamage,
	slipkiksSavvy: true,
	songflowerSerenade: true,
	spiritOfZandalar: true,
	warchiefsBlessing: true,
});

export const DefaultPartyBuffs = PartyBuffs.create({});

export const DefaultDebuffs = Debuffs.create({
	faerieFire: true,
	judgementOfWisdom: true,
	stormstrike: true,
});

export const OtherDefaults = {
	distanceFromTarget: 15,
	profession1: Profession.Engineering,
	profession2: Profession.Tailoring,
};

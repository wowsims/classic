import { Phase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	Conjured,
	Consumes,
	Debuffs,
	Flask,
	Food,
	IndividualBuffs,
	ManaRegenElixir,
	Potions,
	Profession,
	RaidBuffs,
	SaygesFortune,
	ShadowPowerBuff,
	SpellPowerBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { ShadowPriest_Options as Options } from '../core/proto/priest.js';
import { SavedTalents } from '../core/proto/ui.js';
import APLShadowUDJSON from './apls/shadow_ud.apl.json';
import APLShadowJSON from './apls/shadow.apl.json';
import APLDiscJSON from './apls/disc.apl.json';
import P0BISGear from './gear_sets/p0.bis.gear.json';
import P1Shadow from './gear_sets/shadow.p1.json';
import P2Shadow from './gear_sets/shadow.p2.json';
import P3Shadow from './gear_sets/shadow.p3.json';
import P4Shadow from './gear_sets/shadow.p4.json';
import P5Shadow from './gear_sets/shadow.p5.json';
import P6Shadow from './gear_sets/shadow.p6.json';
import P1Disc from './gear_sets/disc.p1.json';
import P2Disc from './gear_sets/disc.p2.json';
import P3Disc from './gear_sets/disc.p3.json';
import P4Disc from './gear_sets/disc.p4.json';
import P5Disc from './gear_sets/disc.p5.json';
import P6Disc from './gear_sets/disc.p6.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearP0BIS = PresetUtils.makePresetGear('Pre-BiS', P0BISGear);
export const GearP1Shadow = PresetUtils.makePresetGear('P1 Shadow', P1Shadow);
export const GearP2Shadow = PresetUtils.makePresetGear('P2 Shadow', P2Shadow);
export const GearP3Shadow = PresetUtils.makePresetGear('P3 Shadow', P3Shadow);
export const GearP4Shadow = PresetUtils.makePresetGear('P4 Shadow', P4Shadow);
export const GearP5Shadow = PresetUtils.makePresetGear('P5 Shadow', P5Shadow);
export const GearP6Shadow = PresetUtils.makePresetGear('P6 Shadow', P6Shadow);
export const GearP1Disc = PresetUtils.makePresetGear('P1 Disc', P1Disc);
export const GearP2Disc = PresetUtils.makePresetGear('P2 Disc', P2Disc);
export const GearP3Disc = PresetUtils.makePresetGear('P3 Disc', P3Disc);
export const GearP4Disc = PresetUtils.makePresetGear('P4 Disc', P4Disc);
export const GearP5Disc = PresetUtils.makePresetGear('P5 Disc', P5Disc);
export const GearP6Disc = PresetUtils.makePresetGear('P6 Disc', P6Disc);

export const GearPresets = {
	[Phase.Phase1]: [GearP0BIS, GearP1Shadow,GearP1Disc],
	[Phase.Phase2]: [GearP2Shadow,GearP2Disc],
	[Phase.Phase3]: [GearP3Shadow,GearP3Disc],
	[Phase.Phase4]: [GearP4Shadow,GearP4Disc],
	[Phase.Phase5]: [GearP5Shadow,GearP5Disc],
	[Phase.Phase6]: [GearP6Shadow,GearP6Disc],
};

export const DefaultGear = GearP0BIS;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLShadowUD = PresetUtils.makePresetAPLRotation('Shadow UD', APLShadowUDJSON);
export const APLShadow = PresetUtils.makePresetAPLRotation('Shadow', APLShadowJSON);
export const APLDisc = PresetUtils.makePresetAPLRotation('Disc', APLDiscJSON);

export const APLPresets = {
	[Phase.Phase1]: [APLShadowUD,APLShadow,APLDisc],
};

export const DefaultAPL = APLPresets[Phase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsP1Shadow = PresetUtils.makePresetTalents('Shadow', SavedTalents.create({ talentsString: '50023013--5002524103511251' }));

export const TalentPresets = {
	[Phase.Phase1]: [TalentsP1Shadow],
};

export const DefaultTalents = TalentPresets[Phase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = Options.create({});

export const DefaultConsumes = Consumes.create({
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodRunnTumTuberSurprise,
	mainHandImbue: WeaponImbue.BrilliantWizardOil,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,

	shadowPowerBuff: ShadowPowerBuff.ElixirOfShadowPower,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	zanzaBuff: ZanzaBuff.CerebralCortexCompound,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	manaSpringTotem: TristateEffect.TristateEffectImproved,
	moonkinAura: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	rallyingCryOfTheDragonslayer: true,
	saygesFortune: SaygesFortune.SaygesDamage,
	slipkiksSavvy: true,
	songflowerSerenade: true,
	spiritOfZandalar: true,
	warchiefsBlessing: true,
});

export const DefaultDebuffs = Debuffs.create({
	improvedShadowBolt: true,
	judgementOfWisdom: true,
	wintersChill: true,
});

export const OtherDefaults = {
	channelClipDelay: 100,
	distanceFromTarget: 30,
	profession1: Profession.Alchemy,
	profession2: Profession.Enchanting,
};

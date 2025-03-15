import { Phase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	AttackPowerBuff,
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
	SapperExplosive,
	SaygesFortune,
	Spec,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { FeralDruid_Options as FeralDruidOptions, FeralDruid_Rotation as FeralDruidRotation } from '../core/proto/druid.js';
import { SavedTalents } from '../core/proto/ui.js';
import P2APL from './apls/p2.apl.json';
import P2BISGear from './gear_sets/p2.bis.gear.json';
import P2PreBISGear from './gear_sets/p2.pre-bis.gear.json';
import P3BISGear from './gear_sets/p3.bis.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearP2BIS = PresetUtils.makePresetGear('P2 BiS', P2BISGear);
export const GearP2PreBIS = PresetUtils.makePresetGear('P2 Pre-BiS', P2PreBISGear);
export const GearP3BIS = PresetUtils.makePresetGear('P3 BiS', P3BISGear)

export const GearPresets = {
	[Phase.Phase2]: [GearP2PreBIS, GearP2BIS, GearP3BIS],
};

export const DefaultGear = GearP2BIS;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLP2Feral = PresetUtils.makePresetAPLRotation('Feral', P2APL);

export const APLPresets = {
	[Phase.Phase2]: [APLP2Feral],
};

export const DefaultAPL = APLPresets[Phase.Phase2][0];

export const DefaultRotation = FeralDruidRotation.create({
	maintainFaerieFire: false,
	minCombosForRip: 3,
	maxWaitTime: 2.0,
	precastTigersFury: false,
	useShredTrick: false,
});

export const SIMPLE_ROTATION_DEFAULT = PresetUtils.makePresetSimpleRotation('Simple Default', Spec.SpecFeralDruid, DefaultRotation);

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

export const TalentsP2Feral = PresetUtils.makePresetTalents('Feral', SavedTalents.create({ talentsString: '500005301-5500021323202151-05' }));

export const TalentPresets = {
	[Phase.Phase2]: [TalentsP2Feral],
};

export const DefaultTalents = TalentPresets[Phase.Phase2][0];

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = FeralDruidOptions.create({
	latencyMs: 100,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	dragonBreathChili: true,
	flask: Flask.FlaskOfDistilledWisdom,
	food: Food.FoodGrilledSquid,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.GroundScorpokAssay,
	sapperExplosive: SapperExplosive.SapperGoblinSapper,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	battleShout: TristateEffect.TristateEffectImproved,
	divineSpirit: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	leaderOfThePack: true,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	fengusFerocity: true,
	rallyingCryOfTheDragonslayer: true,
	saygesFortune: SaygesFortune.SaygesDamage,
	songflowerSerenade: true,
	spiritOfZandalar: false,
	warchiefsBlessing: true,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	exposeArmor: TristateEffect.TristateEffectImproved,
	faerieFire: false,
	sunderArmor: true,
});

export const OtherDefaults = {
	profession1: Profession.Engineering,
	profession2: Profession.Leatherworking,
};

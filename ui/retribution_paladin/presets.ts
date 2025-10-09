import { Phase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	AttackPowerBuff,
	Conjured,
	Consumes,
	Debuffs,
	Explosive,
	FirePowerBuff,
	Flask,
	Food,
	IndividualBuffs,
	Potions,
	Profession,
	RaidBuffs,
	SaygesFortune,
	SpellPowerBuff,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { PaladinAura, PaladinOptions as RetributionPaladinOptions, PaladinSeal } from '../core/proto/paladin.js';
import { SavedTalents } from '../core/proto/ui.js';
import APLP1RetJson from './apls/p1ret.apl.json';
import APLP2RetJson from './apls/p2ret.apl.json';
import APLP3RetJson from './apls/p3ret.apl.json';
import APLP4RetJson from './apls/p4ret.apl.json';
import APLP4RetExodinJson from './apls/p4ret-exodin.apl.json';
import APLP4RetExodin6PcT1Json from './apls/p4ret-exodin-6pcT1.apl.json';
import APLP4RetTwisting6PcT1Json from './apls/p4ret-twisting-6pcT1.apl.json';
import APLPP5ExodinJson from './apls/p5ret-exodin-6CF2DR.apl.json';
import APLPP5TwistingSlowJson from './apls/p5ret-twist-4DR-3.5-3.6.apl.json';
import APLPP5TwistingSlowerJson from './apls/p5ret-twist-4DR-3.7-4.0.apl.json';
import APLPP5ShockadinJson from './apls/p5Shockadin.apl.json';
import BlankGear from './gear_sets/blank.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearBlank = PresetUtils.makePresetGear('Blank', BlankGear);

export const GearPresets = {};

export const DefaultGear = GearBlank;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLP1Ret = PresetUtils.makePresetAPLRotation('P1 Ret', APLP1RetJson);
export const APLP2Ret = PresetUtils.makePresetAPLRotation('P2 Ret/Shockadin', APLP2RetJson);
export const APLP3Ret = PresetUtils.makePresetAPLRotation('P3 Ret/Shockadin', APLP3RetJson);
export const APLP4RetTwist = PresetUtils.makePresetAPLRotation('P4 Ret Twist', APLP4RetJson);
export const APLP4RetTwist6pT1 = PresetUtils.makePresetAPLRotation('P4 Ret Twist 6pT1', APLP4RetTwisting6PcT1Json);
export const APLP4RetExodin = PresetUtils.makePresetAPLRotation('P4 Ret Exodin', APLP4RetExodinJson);
export const APLP4RetExodin6pT1 = PresetUtils.makePresetAPLRotation('P4 Ret Exodin 6pT1', APLP4RetExodin6PcT1Json);
export const APLPP5Twisting4DRSlow = PresetUtils.makePresetAPLRotation('P5 Twist 4DR Slow 3.5-3.6', APLPP5TwistingSlowJson);
export const APLPP5Twisting4DRSlower = PresetUtils.makePresetAPLRotation('P5 Twist 4DR Slower 3.7+', APLPP5TwistingSlowerJson);
export const APLPP5Exodin = PresetUtils.makePresetAPLRotation('P5 Exodin', APLPP5ExodinJson);
export const APLPP5Shockadin = PresetUtils.makePresetAPLRotation('P5 Shockadin', APLPP5ShockadinJson);

export const APLPresets = {
	[Phase.Phase1]: [APLP1Ret],
	[Phase.Phase2]: [APLP2Ret],
	[Phase.Phase3]: [APLP3Ret],
	[Phase.Phase4]: [APLP4RetTwist, APLP4RetTwist6pT1, APLP4RetExodin, APLP4RetExodin6pT1],
	[Phase.Phase5]: [APLPP5Twisting4DRSlow, APLPP5Twisting4DRSlower, APLPP5Exodin, APLPP5Shockadin],
};

export const DefaultAPL = APLPresets[Phase.Phase5][0];

///////////////////////////////////////////////////////////////////////////
//                                 Talent presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const RetTalents = PresetUtils.makePresetTalents('Retribution', SavedTalents.create({ talentsString: '500501-503-52230351200315' }));

export const ShockadinTalents = PresetUtils.makePresetTalents('Shockadin', SavedTalents.create({ talentsString: '55050112501051--0523005122' }));

export const TalentPresets = {
	[Phase.Phase1]: [RetTalents, ShockadinTalents],
};

export const DefaultTalents = TalentPresets[Phase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = RetributionPaladinOptions.create({
	aura: PaladinAura.SanctityAura,
	primarySeal: PaladinSeal.Command,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	boglingRoot: false,
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	dragonBreathChili: true,
	fillerExplosive: Explosive.ExplosiveUnknown,
	firePowerBuff: FirePowerBuff.ElixirOfGreaterFirepower,
	food: Food.FoodBlessSunfruit,
	flask: Flask.FlaskOfSupremePower,
	//mainHandImbue: WeaponImbue.WildStrikes,
	//offHandImbue: WeaponImbue.MagnificentTrollshine,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.ROIDS,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfMight: TristateEffect.TristateEffectImproved,
	blessingOfKings: true,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	fengusFerocity: true,
	moldarsMoxie: true,
	rallyingCryOfTheDragonslayer: true,
	saygesFortune: SaygesFortune.SaygesDamage,
	slipkiksSavvy: true,
	songflowerSerenade: true,
	spiritOfZandalar: true,
	warchiefsBlessing: true,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	battleShout: TristateEffect.TristateEffectImproved,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	sanctityAura: true,
	leaderOfThePack: true,
	moonkinAura: true,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	faerieFire: true,
	giftOfArthas: true,
	sunderArmor: true,
	judgementOfWisdom: true,
	judgementOfTheCrusader: TristateEffect.TristateEffectImproved,
	improvedScorch: true,
});

export const OtherDefaults = {
	profession1: Profession.Blacksmithing,
	profession2: Profession.Enchanting,
};

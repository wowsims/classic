import { Phase } from '../core/constants/other.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	Alcohol,
	AttackPowerBuff,
	Conjured,
	Consumes,
	Debuffs,
	Flask,
	Food,
	HealthElixir,
	IndividualBuffs,
	ManaRegenElixir,
	Potions,
	Profession,
	Race,
	RaidBuffs,
	SapperExplosive,
	SaygesFortune,
	SpellPowerBuff,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import {
	Hunter_Options as HunterOptions,
	Hunter_Options_Ammo as Ammo,
	Hunter_Options_PetAttackSpeed as PetAttackSpeed,
	Hunter_Options_PetType as PetType,
	Hunter_Options_QuiverBonus,
} from '../core/proto/hunter.js';
import { SavedTalents } from '../core/proto/ui.js';
import P1APL from './apls/p1.apl.json';
import P0BISGear from './gear_sets/p0.bis.gear.json';
import P1BISGear from './gear_sets/p1.bis.gear.json';
import P2BISGear from './gear_sets/p2.bis.gear.json';
import P3BISGear from './gear_sets/p3.bis.gear.json';
import P4BISGear from './gear_sets/p4.bis.gear.json';
import P5BISGear from './gear_sets/p5.bis.gear.json';
import P6BISGear from './gear_sets/p6.bis.gear.json';

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
export const GearP6BIS = PresetUtils.makePresetGear('P6 BiS', P6BISGear);

export const GearPresets = {
	[Phase.Phase1]: [GearP0BIS, GearP1BIS],
	[Phase.Phase2]: [GearP2BIS],
	[Phase.Phase3]: [GearP3BIS],
	[Phase.Phase4]: [GearP4BIS],
	[Phase.Phase5]: [GearP5BIS],
	[Phase.Phase6]: [GearP6BIS],
};

export const DefaultGear = GearP0BIS;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLP1 = PresetUtils.makePresetAPLRotation('Marksmanship', P1APL);

export const APLPresets = {
	[Phase.Phase1]: [APLP1],
};

export const DefaultAPL = APLPresets[Phase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

export const TalentsP1 = PresetUtils.makePresetTalents('Marksmanship', SavedTalents.create({ talentsString: '55000000505-05451002503051' }));

export const TalentPresets = {
	[Phase.Phase1]: [TalentsP1],
};

export const DefaultTalents = TalentPresets[Phase.Phase1][0];

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = HunterOptions.create({
	ammo: Ammo.ThoriumHeadedArrow,
	quiverBonus: Hunter_Options_QuiverBonus.Speed15,
	petAttackSpeed: PetAttackSpeed.OneTwo,
	petType: PetType.Cat,
	petUptime: 1,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	dragonBreathChili: true,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodSmokedDesertDumpling,
	healthElixir: HealthElixir.ElixirOfFortitude,
	mainHandImbue: WeaponImbue.Windfury,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,
	offHandImbue: WeaponImbue.ElementalSharpeningStone,
	petAttackPowerConsumable: 1,
	petAgilityConsumable: 1,
	petStrengthConsumable: 1,
	sapperExplosive: SapperExplosive.SapperUnknown,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.GroundScorpokAssay,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	battleShout: TristateEffect.TristateEffectImproved,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	graceOfAirTotem: TristateEffect.TristateEffectImproved,
	leaderOfThePack: false,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	blessingOfWisdom: TristateEffect.TristateEffectImproved,
	fengusFerocity: false,
	moldarsMoxie: false,
	rallyingCryOfTheDragonslayer: true,
	saygesFortune: SaygesFortune.SaygesDamage,
	slipkiksSavvy: false,
	songflowerSerenade: true,
	spiritOfZandalar: false,
	warchiefsBlessing: false,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	exposeArmor: TristateEffect.TristateEffectImproved,
	faerieFire: true,
	huntersMark: TristateEffect.TristateEffectImproved,
	improvedScorch: true,
	judgementOfWisdom: true,
	stormstrike: false,
	sunderArmor: true,
});

export const OtherDefaults = {
	distanceFromTarget: 12,
	profession1: Profession.Enchanting,
	profession2: Profession.Engineering,
	race: Race.RaceTroll,
};

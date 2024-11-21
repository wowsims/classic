import { Phase } from '../core/constants/other.js';
import { Player } from '../core/player.js';
import * as PresetUtils from '../core/preset_utils.js';
import {
	AgilityElixir,
	Alcohol,
	ArmorElixir,
	Conjured,
	Consumes,
	Debuffs,
	Explosive,
	FirePowerBuff,
	Flask,
	Food,
	HealthElixir,
	IndividualBuffs,
	ManaRegenElixir,
	Potions,
	Profession,
	RaidBuffs,
	SapperExplosive,
	SaygesFortune,
	ShadowPowerBuff,
	SpellPowerBuff,
	TristateEffect,
	WeaponImbue,
} from '../core/proto/common.js';
import { SavedTalents } from '../core/proto/ui.js';
import {
	WarlockOptions as WarlockOptions,
	WarlockOptions_Armor as Armor,
	WarlockOptions_Summon as Summon,
	WarlockOptions_WeaponImbue as WarlockWeaponImbue,
} from '../core/proto/warlock.js';
import Phase1AfflictionAPL from './apls/p1.affi.tank.apl.json';
import Phase1DestroTankAPL from './apls/p1.destro.tank.apl.json';
import Phase2DemonologyAPL from './apls/p2.demo.tank.apl.json';
import Phase2DestroTankAPL from './apls/p2.destro.tank.apl.json';
import Phase3TankAPL from './apls/p3.destro.tank.apl.json';
import Phase4DemoTankAPL from './apls/p4_demo_tank.apl.json';
import Phase4DestroAffTankAPL from './apls/p4_destro_aff_tank.apl.json';
import BlankGear from './gear_sets/blank.gear.json';

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearBlank = PresetUtils.makePresetGear('Blank', BlankGear);

export const GearPresets = {};

export const DefaultGear = GearBlank;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets
///////////////////////////////////////////////////////////////////////////

export const APLAfflictionTankPhase1 = PresetUtils.makePresetAPLRotation('P1 Aff', Phase1AfflictionAPL, {
	customCondition: player => player.getLevel() === 25,
});
export const APLDestructionTankPhase1 = PresetUtils.makePresetAPLRotation('P1 Destro', Phase1DestroTankAPL, {
	customCondition: player => player.getLevel() === 25,
});

export const APLDemonologyTankPhase2 = PresetUtils.makePresetAPLRotation('P2 Demo', Phase2DemonologyAPL, {
	customCondition: player => player.getLevel() === 40,
});
export const APLDestructionTankPhase2 = PresetUtils.makePresetAPLRotation('P2 Destro', Phase2DestroTankAPL, {
	customCondition: player => player.getLevel() === 40,
});

export const APLTankPhase3 = PresetUtils.makePresetAPLRotation('P3 Destro', Phase3TankAPL, {
	customCondition: player => player.getLevel() === 50,
});

export const APLDemoTankPhase4 = PresetUtils.makePresetAPLRotation('P4 Demo', Phase4DemoTankAPL, {
	customCondition: player => player.getLevel() === 60,
});
export const APLDestroAffTankPhase4 = PresetUtils.makePresetAPLRotation('P4 Destro/Aff', Phase4DestroAffTankAPL, {
	customCondition: player => player.getLevel() === 60,
});

export const APLPresets = {
	[Phase.Phase1]: [APLAfflictionTankPhase1, APLDestructionTankPhase1],
	[Phase.Phase2]: [APLDemonologyTankPhase2, APLDestructionTankPhase2],
	[Phase.Phase3]: [APLTankPhase3],
	[Phase.Phase4]: [APLDemoTankPhase4, APLDestroAffTankPhase4],
	[Phase.Phase5]: [],
};

// The default APLs for each spec
// 0 = Affliction
// 1 = Demonology
// 2 = Destruction
export const DefaultAPLs: Record<number, Record<number, PresetUtils.PresetRotation>> = {
	25: {
		0: APLPresets[Phase.Phase1][0],
		1: APLPresets[Phase.Phase1][1],
	},
	40: {
		0: APLPresets[Phase.Phase2][0],
		1: APLPresets[Phase.Phase2][1],
	},
	50: {
		0: APLPresets[Phase.Phase3][0],
		1: APLPresets[Phase.Phase3][0],
	},
	60: {
		0: APLPresets[Phase.Phase4][1],
		1: APLPresets[Phase.Phase4][0],
		2: APLPresets[Phase.Phase4][1],
	},
};

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

export const TalentsAfflictionTankPhase1 = PresetUtils.makePresetTalents('25 Aff', SavedTalents.create({ talentsString: '050025001-003' }), {
	customCondition: (player: Player<any>) => player.getLevel() === 25,
});

export const TalentsDestructionTankPhase1 = PresetUtils.makePresetTalents('25 Destro', SavedTalents.create({ talentsString: '-03-0550201' }), {
	customCondition: (player: Player<any>) => player.getLevel() === 25,
});

export const TalentsDemonologyTankPhase2 = PresetUtils.makePresetTalents('40 Demo', SavedTalents.create({ talentsString: '-2050033112501251' }), {
	customCondition: (player: Player<any>) => player.getLevel() === 40,
});

export const TalentsDestructionTankPhase2 = PresetUtils.makePresetTalents('40 Destro', SavedTalents.create({ talentsString: '-035-05500050025001' }), {
	customCondition: (player: Player<any>) => player.getLevel() === 40,
});

export const TalentsTankPhase3 = PresetUtils.makePresetTalents('50 Destro', SavedTalents.create({ talentsString: '05-03-505020500050515' }), {
	customCondition: (player: Player<any>) => player.getLevel() === 50,
});

export const TalentsAffTankPhase4 = PresetUtils.makePresetTalents('60 Aff', SavedTalents.create({ talentsString: '5500253011201002-03-50502051002001' }), {
	customCondition: (player: Player<any>) => player.getLevel() === 60,
});
export const TalentsDemoTankPhase4 = PresetUtils.makePresetTalents('60 Demo', SavedTalents.create({ talentsString: '-205004015250105-50500050005001' }), {
	customCondition: (player: Player<any>) => player.getLevel() === 60,
});
export const TalentsDestroTankPhase4 = PresetUtils.makePresetTalents('60 Destro', SavedTalents.create({ talentsString: '45002400102-03-505020510050115' }), {
	customCondition: (player: Player<any>) => player.getLevel() === 60,
});

export const TalentPresets = {
	[Phase.Phase1]: [TalentsAfflictionTankPhase1, TalentsDestructionTankPhase1],
	[Phase.Phase2]: [TalentsDemonologyTankPhase2, TalentsDestructionTankPhase2],
	[Phase.Phase3]: [TalentsTankPhase3],
	[Phase.Phase4]: [TalentsAffTankPhase4, TalentsDemoTankPhase4, TalentsDestroTankPhase4],
	[Phase.Phase5]: [],
};

export const DefaultTalentsAff = TalentPresets[Phase.Phase4][0];
export const DefaultTalentsDemo = TalentPresets[Phase.Phase4][1];
export const DefaultTalentsDestro = TalentPresets[Phase.Phase4][2];

export const DefaultTalents = DefaultTalentsDestro;

export const PresetBuildAff = PresetUtils.makePresetBuild('Aff', { gear: DefaultGear, talents: DefaultTalentsAff, rotation: DefaultAPLs[60][0] });
export const PresetBuildDemo = PresetUtils.makePresetBuild('Demo', { gear: DefaultGear, talents: DefaultTalentsDemo, rotation: DefaultAPLs[60][1] });
export const PresetBuildDestro = PresetUtils.makePresetBuild('Destro', {
	gear: DefaultGear,
	talents: DefaultTalentsDestro,
	rotation: DefaultAPLs[60][2],
});

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = WarlockOptions.create({
	armor: Armor.FelArmor,
	summon: Summon.Imp,
	weaponImbue: WarlockWeaponImbue.NoWeaponImbue,
});

export const DefaultConsumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfGreaterAgility,
	alcohol: Alcohol.AlcoholRumseyRumBlackLabel,
	armorElixir: ArmorElixir.ElixirOfSuperiorDefense,
	defaultConjured: Conjured.ConjuredDemonicRune,
	defaultPotion: Potions.MajorManaPotion,
	dragonBreathChili: true,
	fillerExplosive: Explosive.ExplosiveDenseDynamite,
	firePowerBuff: FirePowerBuff.ElixirOfGreaterFirepower,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodTenderWolfSteak,
	healthElixir: HealthElixir.ElixirOfFortitude,
	mainHandImbue: WeaponImbue.ShadowOil,
	manaRegenElixir: ManaRegenElixir.MagebloodPotion,

	sapperExplosive: SapperExplosive.SapperGoblinSapper,
	shadowPowerBuff: ShadowPowerBuff.ElixirOfShadowPower,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
});

export const DefaultRaidBuffs = RaidBuffs.create({
	arcaneBrilliance: true,
	battleShout: TristateEffect.TristateEffectImproved,
	bloodPact: TristateEffect.TristateEffectImproved,
	devotionAura: TristateEffect.TristateEffectImproved,
	divineSpirit: true,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	graceOfAirTotem: TristateEffect.TristateEffectRegular,
	manaSpringTotem: TristateEffect.TristateEffectRegular,
	moonkinAura: true,
	powerWordFortitude: TristateEffect.TristateEffectImproved,
	shadowProtection: true,
	strengthOfEarthTotem: TristateEffect.TristateEffectRegular,
	thorns: TristateEffect.TristateEffectImproved,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	blessingOfSanctuary: true,
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

export const DefaultDebuffs = Debuffs.create({
	curseOfElements: true,
	curseOfShadow: true,
	curseOfRecklessness: true,
	demoralizingShout: TristateEffect.TristateEffectImproved,
	faerieFire: true,
	improvedScorch: true,
	judgementOfWisdom: true,
	shadowWeaving: true,
});

export const OtherDefaults = {
	channelClipDelay: 150,
	distanceFromTarget: 5,
	profession1: Profession.Enchanting,
	profession2: Profession.Engineering,
};

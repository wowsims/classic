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
	Profession,
	RaidBuffs,
	SaygesFortune,
	SpellPowerBuff,
	StrengthBuff,
	TristateEffect,
	WeaponImbue,
	ZanzaBuff,
} from '../core/proto/common.js';
import { RogueOptions } from '../core/proto/rogue.js';
import { SavedTalents } from '../core/proto/ui.js';
import BackstabAPL from './apls/combat_backstab.apl.json';
import SinisterStrikeAPL from './apls/combat_sinister_strike.apl.json';
import BlankGear from './gear_sets/blank.gear.json';
import BackstabGearPreBiS from './gear_sets/combat_backstab_prebis.gear.json';
import SinisterStrikeGearPreBiS from './gear_sets/combat_sinister_strike_prebis.gear.json';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

///////////////////////////////////////////////////////////////////////////
//                                 Gear Presets
///////////////////////////////////////////////////////////////////////////

export const GearBlank = PresetUtils.makePresetGear('Blank', BlankGear);
export const GearDaggersPreBiS = PresetUtils.makePresetGear('Backstab PreBiS', BackstabGearPreBiS);
export const GearSwordsPreBiS = PresetUtils.makePresetGear('Swords PreBiS', SinisterStrikeGearPreBiS);

export const GearPresets = {
	[Phase.Phase1]: [GearDaggersPreBiS, GearSwordsPreBiS],
	[Phase.Phase2]: [GearDaggersPreBiS, GearSwordsPreBiS],
	[Phase.Phase3]: [GearDaggersPreBiS, GearSwordsPreBiS],
	[Phase.Phase4]: [GearDaggersPreBiS, GearSwordsPreBiS],
	[Phase.Phase5]: [GearDaggersPreBiS, GearSwordsPreBiS],
};

export const DefaultGear = GearSwordsPreBiS;

///////////////////////////////////////////////////////////////////////////
//                                 APL Presets[]
///////////////////////////////////////////////////////////////////////////

export const ROTATION_PRESET_BACKSTAB = PresetUtils.makePresetAPLRotation('Backstab', BackstabAPL, {});
export const ROTATION_PRESET_SINISTER_STRIKE = PresetUtils.makePresetAPLRotation('Sinister Strike', SinisterStrikeAPL, {});

export const APLPresets = {
	[Phase.Phase1]: [ROTATION_PRESET_BACKSTAB, ROTATION_PRESET_SINISTER_STRIKE],
	[Phase.Phase2]: [ROTATION_PRESET_BACKSTAB, ROTATION_PRESET_SINISTER_STRIKE],
	[Phase.Phase3]: [ROTATION_PRESET_BACKSTAB, ROTATION_PRESET_SINISTER_STRIKE],
	[Phase.Phase4]: [ROTATION_PRESET_BACKSTAB, ROTATION_PRESET_SINISTER_STRIKE],
	[Phase.Phase5]: [ROTATION_PRESET_BACKSTAB, ROTATION_PRESET_SINISTER_STRIKE],
};

//Need to add main hand equip logic or talent/rotation logic to map to Auto APL
export const DefaultAPLs: Record<number, PresetUtils.PresetRotation> = {
	[0]: ROTATION_PRESET_SINISTER_STRIKE,
	[1]: ROTATION_PRESET_BACKSTAB,
};

export const DefaultAPLBackstab = APLPresets[Phase.Phase5][0];
export const DefaultAPLSinisterStrike = APLPresets[Phase.Phase5][1];

///////////////////////////////////////////////////////////////////////////
//                                 Talent Presets
///////////////////////////////////////////////////////////////////////////

// Default talents. Uses the wowhead calculator format, make the talents on
// https://wowhead.com/classic/talent-calc and copy the numbers in the url.

// Preset name must be unique. Ex: 'Backstab DPS' cannot be used as a name more than once

export const CombatBackstabTalents = PresetUtils.makePresetTalents(
	'Combat Backstab',
	SavedTalents.create({ talentsString: '005023104-0233050020550100221-05' }),
);

export const CombatSwordsTalents = PresetUtils.makePresetTalents('Combat Swords', SavedTalents.create({ talentsString: '005323105-0240052020050150231' }));

export const TalentPresets = {
	[Phase.Phase1]: [CombatSwordsTalents, CombatBackstabTalents],
	[Phase.Phase2]: [CombatSwordsTalents, CombatBackstabTalents],
	[Phase.Phase3]: [CombatSwordsTalents, CombatBackstabTalents],
	[Phase.Phase4]: [CombatSwordsTalents, CombatBackstabTalents],
	[Phase.Phase5]: [CombatSwordsTalents, CombatBackstabTalents],
};

export const DefaultTalentsAssassin = TalentPresets[Phase.Phase5][0];
export const DefaultTalentsCombat = TalentPresets[Phase.Phase5][0];
export const DefaultTalentsSubtlety = TalentPresets[Phase.Phase5][0];

export const DefaultTalentsBackstab = TalentPresets[Phase.Phase5][0];
export const DefaultTalentsSinisterStrike = TalentPresets[Phase.Phase5][1];

export const DefaultTalents = DefaultTalentsAssassin;

///////////////////////////////////////////////////////////////////////////
//                                Encounters
///////////////////////////////////////////////////////////////////////////
export const PresetBuildBackstab = PresetUtils.makePresetBuild('Backstab', {
	gear: GearDaggersPreBiS,
	talents: DefaultTalentsBackstab,
	rotation: DefaultAPLBackstab,
});

///////////////////////////////////////////////////////////////////////////
//                                 Options
///////////////////////////////////////////////////////////////////////////

export const DefaultOptions = RogueOptions.create({});

///////////////////////////////////////////////////////////////////////////
//                         Consumes/Buffs/Debuffs
///////////////////////////////////////////////////////////////////////////

export const P5Consumes = Consumes.create({
	agilityElixir: AgilityElixir.ElixirOfTheMongoose,
	attackPowerBuff: AttackPowerBuff.JujuMight,
	defaultConjured: Conjured.ConjuredRogueThistleTea,
	dragonBreathChili: true,
	flask: Flask.FlaskOfSupremePower,
	food: Food.FoodGrilledSquid,
	mainHandImbue: WeaponImbue.Windfury,
	offHandImbue: WeaponImbue.InstantPoison,
	spellPowerBuff: SpellPowerBuff.GreaterArcaneElixir,
	strengthBuff: StrengthBuff.JujuPower,
	zanzaBuff: ZanzaBuff.GroundScorpokAssay,
});

export const DefaultConsumes = {
	[Phase.Phase1]: P5Consumes,
	[Phase.Phase2]: P5Consumes,
	[Phase.Phase3]: P5Consumes,
	[Phase.Phase4]: P5Consumes,
	[Phase.Phase5]: P5Consumes,
};

export const DefaultRaidBuffs = RaidBuffs.create({
	battleShout: TristateEffect.TristateEffectImproved,
	fireResistanceAura: true,
	fireResistanceTotem: true,
	giftOfTheWild: TristateEffect.TristateEffectImproved,
	strengthOfEarthTotem: TristateEffect.TristateEffectImproved,
	graceOfAirTotem: TristateEffect.TristateEffectImproved,
	leaderOfThePack: true,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfMight: TristateEffect.TristateEffectImproved,
	fengusFerocity: true,
	rallyingCryOfTheDragonslayer: true,
	saygesFortune: SaygesFortune.SaygesDamage,
	slipkiksSavvy: true,
	songflowerSerenade: true,
	spiritOfZandalar: true,
	warchiefsBlessing: true,
});

export const DefaultDebuffs = Debuffs.create({
	curseOfRecklessness: true,
	faerieFire: true,
	improvedScorch: true,
	sunderArmor: true,
});

export const OtherDefaults = {
	profession1: Profession.Engineering,
	profession2: Profession.Alchemy,
};

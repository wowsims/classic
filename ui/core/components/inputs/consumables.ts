import { IndividualSimUI } from '../../individual_sim_ui';
import { Player } from '../../player';
import {
	AgilityElixir,
	Alcohol,
	ArmorElixir,
	AttackPowerBuff,
	Class,
	Conjured,
	Consumes,
	Explosive,
	Faction,
	FirePowerBuff,
	Flask,
	Food,
	FrostPowerBuff,
	HealthElixir,
	ItemSlot,
	ManaRegenElixir,
	Potions,
	Profession,
	SapperExplosive,
	ShadowPowerBuff,
	Spec,
	SpellPowerBuff,
	Stat,
	StrengthBuff,
	WeaponImbue,
	ZanzaBuff,
} from '../../proto/common';
import { ActionId } from '../../proto_utils/action_id';
import { isBluntWeaponType, isSharpWeaponType, isWeapon } from '../../proto_utils/utils';
import { EventID, TypedEvent } from '../../typed_event';
import { IconEnumValueConfig } from '../icon_enum_picker';
import { makeBooleanConsumeInput, makeBooleanMiscConsumeInput, makeBooleanPetMiscConsumeInput, makeEnumConsumeInput } from '../icon_inputs';
import { IconPicker, IconPickerDirection } from '../icon_picker';
import * as InputHelpers from '../input_helpers';
import { MultiIconPicker, MultiIconPickerConfig, MultiIconPickerItemConfig } from '../multi_icon_picker';
import { DeadlyPoisonWeaponImbue, InstantPoisonWeaponImbue, WoundPoisonWeaponImbue } from './rogue_imbues';
import { FlametongueWeaponImbue, FrostbrandWeaponImbue, RockbiterWeaponImbue, WindfuryWeaponImbue } from './shaman_imbues';
import { ActionInputConfig, ItemStatOption, PickerStatOptions, StatOptions } from './stat_options';
import { FeralDruid } from '../../proto/druid';

export interface ConsumableInputConfig<T> extends ActionInputConfig<T> {
	value: T;
}

export interface ConsumableStatOption<T> extends ItemStatOption<T> {
	config: ConsumableInputConfig<T>;
}

export interface ConsumeInputFactoryArgs<T extends number> {
	consumesFieldName: keyof Consumes;
	// Additional callback if logic besides syncing consumes is required
	onSet?: (eventactionId: EventID, player: Player<any>, newValue: T) => void;
	showWhen?: (player: Player<any>) => boolean;
}

function makeConsumeInputFactory<T extends number>(
	args: ConsumeInputFactoryArgs<T>,
): (options: ConsumableStatOption<T>[], tooltip?: string) => InputHelpers.TypedIconEnumPickerConfig<Player<any>, T> {
	return (options: ConsumableStatOption<T>[], tooltip?: string) => {
		return {
			type: 'iconEnum',
			tooltip: tooltip,
			numColumns: options.length > 11 ? 4 : options.length > 8 ? 3 : options.length > 5 ? 2 : 1,
			values: [{ value: 0 } as unknown as IconEnumValueConfig<Player<any>, T>].concat(
				options.map(option => {
					return {
						actionId: option.config.actionId,
						value: option.config.value,
						showWhen: (player: Player<any>) => !option.config.showWhen || option.config.showWhen(player),
					} as IconEnumValueConfig<Player<any>, T>;
				}),
			),
			equals: (a: T, b: T) => a == b,
			zeroValue: 0 as T,
			changedEvent: (player: Player<any>) =>
				TypedEvent.onAny([player.consumesChangeEmitter, player.gearChangeEmitter, player.professionChangeEmitter, player.raceChangeEmitter]),
			showWhen: (player: Player<any>) => !args.showWhen || args.showWhen(player),
			getValue: (player: Player<any>) => player.getConsumes()[args.consumesFieldName] as T,
			setValue: (eventID: EventID, player: Player<any>, newValue: number) => {
				const newConsumes = player.getConsumes();

				if (newConsumes[args.consumesFieldName] === newValue) {
					return;
				}

				(newConsumes[args.consumesFieldName] as number) = newValue;
				TypedEvent.freezeAllAndDo(() => {
					player.setConsumes(eventID, newConsumes);
					if (args.onSet) {
						args.onSet(eventID, player, newValue as T);
					}
				});
			},
		};
	};
}

type MultiIconConsumeInputFactoryArg<ModObject> = Omit<MultiIconPickerConfig<ModObject>, 'values'>;

export const makeMultiIconConsumesInputFactory = <ModObject>(
	config: MultiIconConsumeInputFactoryArg<ModObject>,
): ((parent: HTMLElement, modObj: ModObject, simUI: IndividualSimUI<Spec>, options: StatOptions<any, any>) => MultiIconPicker<any>) => {
	return (parent: HTMLElement, modObj: ModObject, simUI: IndividualSimUI<Spec>, options: StatOptions<any, any>) => {
		const pickerConfig = {
			...config,
			values: options.map(option => option.config) as Array<MultiIconPickerItemConfig<ModObject>>,
		};
		return new MultiIconPicker(parent, modObj, pickerConfig, simUI);
	};
};

///////////////////////////////////////////////////////////////////////////
//                                 CONJURED
///////////////////////////////////////////////////////////////////////////

export const ConjuredHealthstone: ConsumableInputConfig<Conjured> = {
	actionId: () => ActionId.fromItemId(5509),
	value: Conjured.ConjuredHealthstone,
};
export const ConjuredGreaterHealthstone: ConsumableInputConfig<Conjured> = {
	actionId: () => ActionId.fromItemId(5510),
	value: Conjured.ConjuredGreaterHealthstone,
};
export const ConjuredMajorHealthstone: ConsumableInputConfig<Conjured> = {
	actionId: () => ActionId.fromItemId(9421),
	value: Conjured.ConjuredMajorHealthstone,
};

export const ConjuredMinorRecombobulator: ConsumableInputConfig<Conjured> = {
	actionId: () => ActionId.fromItemId(4381),
	value: Conjured.ConjuredMinorRecombobulator,
	showWhen: (player: Player<any>) => player.getGear().hasTrinket(4381),
};
export const ConjuredDemonicRune: ConsumableInputConfig<Conjured> = {
	actionId: () => ActionId.fromItemId(12662),
	value: Conjured.ConjuredDemonicRune,
};
export const ConjuredRogueThistleTea: ConsumableInputConfig<Conjured> = {
	actionId: () => ActionId.fromItemId(7676),
	value: Conjured.ConjuredRogueThistleTea,
	showWhen: player => player.getClass() == Class.ClassRogue,
};

export const CONJURED_CONFIG: ConsumableStatOption<Conjured>[] = [
	{ config: ConjuredMajorHealthstone, stats: [Stat.StatArmor] },
	{ config: ConjuredGreaterHealthstone, stats: [Stat.StatArmor] },
	{ config: ConjuredHealthstone, stats: [Stat.StatArmor] },

	{ config: ConjuredDemonicRune, stats: [Stat.StatIntellect] },
	{ config: ConjuredMinorRecombobulator, stats: [Stat.StatIntellect] },

	{ config: ConjuredRogueThistleTea, stats: [] },
];

export const makeConjuredInput = makeConsumeInputFactory({ consumesFieldName: 'defaultConjured' });

///////////////////////////////////////////////////////////////////////////
//                                 EXPLOSIVES
///////////////////////////////////////////////////////////////////////////

export const SapperGoblinSapper: ConsumableInputConfig<SapperExplosive> = {
	actionId: () => ActionId.fromItemId(10646),
	showWhen: player => player.hasProfession(Profession.Engineering),
	value: SapperExplosive.SapperGoblinSapper,
};

export const ExplosiveSolidDynamite: ConsumableInputConfig<Explosive> = {
	actionId: () => ActionId.fromItemId(10507),
	showWhen: player => player.hasProfession(Profession.Engineering),
	value: Explosive.ExplosiveSolidDynamite,
};

export const ExplosiveGoblinLandMine: ConsumableInputConfig<Explosive> = {
	actionId: () => ActionId.fromItemId(4395),
	showWhen: player => player.hasProfession(Profession.Engineering),
	value: Explosive.ExplosiveGoblinLandMine,
};

export const ExplosiveDenseDynamite: ConsumableInputConfig<Explosive> = {
	actionId: () => ActionId.fromItemId(18641),
	showWhen: player => player.hasProfession(Profession.Engineering),
	value: Explosive.ExplosiveDenseDynamite,
};

export const ExplosiveThoriumGrenade: ConsumableInputConfig<Explosive> = {
	actionId: () => ActionId.fromItemId(15993),
	showWhen: player => player.hasProfession(Profession.Engineering),
	value: Explosive.ExplosiveThoriumGrenade,
};

export const EXPLOSIVES_CONFIG: ConsumableStatOption<Explosive>[] = [
	{ config: ExplosiveSolidDynamite, stats: [] },
	{ config: ExplosiveDenseDynamite, stats: [] },
	{ config: ExplosiveThoriumGrenade, stats: [] },
	{ config: ExplosiveGoblinLandMine, stats: [] },
];

export const SAPPER_CONFIG: ConsumableStatOption<SapperExplosive>[] = [{ config: SapperGoblinSapper, stats: [] }];

export const makeExplosivesInput = makeConsumeInputFactory({
	consumesFieldName: 'fillerExplosive',
	//showWhen: (player) => !!player.getProfessions().find(p => p == Profession.Engineering),
});

export const makeSappersInput = makeConsumeInputFactory({
	consumesFieldName: 'sapperExplosive',
});

///////////////////////////////////////////////////////////////////////////
//                                 FLASKS
///////////////////////////////////////////////////////////////////////////

// Original lvl 50 not obtainable in Phase 3
export const FlaskOfTheTitans: ConsumableInputConfig<Flask> = {
	actionId: () => ActionId.fromItemId(13510),
	value: Flask.FlaskOfTheTitans,
};
// Original lvl 50 not obtainable in Phase 3
export const FlaskOfDistilledWisdom: ConsumableInputConfig<Flask> = {
	actionId: () => ActionId.fromItemId(13511),
	value: Flask.FlaskOfDistilledWisdom,
};
// Original lvl 50 not obtainable in Phase 3
export const FlaskOfSupremePower: ConsumableInputConfig<Flask> = {
	actionId: () => ActionId.fromItemId(13512),
	value: Flask.FlaskOfSupremePower,
};
// Original lvl 50 not obtainable in Phase 3
export const FlaskOfChromaticResistance: ConsumableInputConfig<Flask> = {
	actionId: () => ActionId.fromItemId(13513),
	value: Flask.FlaskOfChromaticResistance,
};

export const FLASKS_CONFIG: ConsumableStatOption<Flask>[] = [
	{ config: FlaskOfTheTitans, stats: [] },
	{ config: FlaskOfDistilledWisdom, stats: [Stat.StatIntellect] },
	{ config: FlaskOfSupremePower, stats: [Stat.StatMP5, Stat.StatSpellPower] },
	{ config: FlaskOfChromaticResistance, stats: [] },
];

export const makeFlasksInput = makeConsumeInputFactory({ consumesFieldName: 'flask' });

///////////////////////////////////////////////////////////////////////////
//                                 FOOD
///////////////////////////////////////////////////////////////////////////

export const DirgesKickChimaerokChops: ConsumableInputConfig<Food> = {
	actionId: () => ActionId.fromItemId(21023),
	value: Food.FoodDirgesKickChimaerokChops,
};
export const GrilledSquid: ConsumableInputConfig<Food> = {
	actionId: () => ActionId.fromItemId(13928),
	value: Food.FoodGrilledSquid,
};
// Original lvl 50 not obtainable in Phase 3
export const SmokedDesertDumpling: ConsumableInputConfig<Food> = {
	actionId: () => ActionId.fromItemId(20452),
	value: Food.FoodSmokedDesertDumpling,
};
// Original lvl 45 not obtainable in Phase 3
export const RunnTumTuberSurprise: ConsumableInputConfig<Food> = {
	actionId: () => ActionId.fromItemId(18254),
	value: Food.FoodRunnTumTuberSurprise,
};
export const BlessSunfruit: ConsumableInputConfig<Food> = {
	actionId: () => ActionId.fromItemId(13810),
	value: Food.FoodBlessSunfruit,
};
export const BlessedSunfruitJuice: ConsumableInputConfig<Food> = {
	actionId: () => ActionId.fromItemId(13813),
	value: Food.FoodBlessedSunfruitJuice,
};
export const NightfinSoup: ConsumableInputConfig<Food> = {
	actionId: () => ActionId.fromItemId(13931),
	value: Food.FoodNightfinSoup,
};
export const TenderWolfSteak: ConsumableInputConfig<Food> = {
	actionId: () => ActionId.fromItemId(18045),
	value: Food.FoodTenderWolfSteak,
};
export const SagefishDelight: ConsumableInputConfig<Food> = {
	actionId: () => ActionId.fromItemId(21217),
	value: Food.FoodSagefishDelight,
};
export const HotWolfRibs: ConsumableInputConfig<Food> = {
	actionId: () => ActionId.fromItemId(13851),
	value: Food.FoodHotWolfRibs,
};
export const SmokedSagefish: ConsumableInputConfig<Food> = {
	actionId: () => ActionId.fromItemId(21072),
	value: Food.FoodSmokedSagefish,
};

// Ordered by level
export const FOOD_CONFIG: ConsumableStatOption<Food>[] = [
	{ config: DirgesKickChimaerokChops, stats: [Stat.StatStamina] },
	{ config: GrilledSquid, stats: [Stat.StatAgility] },
	{ config: SmokedDesertDumpling, stats: [Stat.StatStrength] },
	{ config: RunnTumTuberSurprise, stats: [Stat.StatIntellect] },
	{ config: BlessSunfruit, stats: [Stat.StatStrength] },
	{ config: BlessedSunfruitJuice, stats: [Stat.StatSpirit] },
	{ config: NightfinSoup, stats: [Stat.StatMP5] },
	{ config: TenderWolfSteak, stats: [Stat.StatStamina, Stat.StatSpirit] },
	{ config: SagefishDelight, stats: [Stat.StatMP5] },
	{ config: HotWolfRibs, stats: [Stat.StatSpirit] },
	{ config: SmokedSagefish, stats: [Stat.StatMP5] },
];

export const makeFoodInput = makeConsumeInputFactory({ consumesFieldName: 'food' });

export const DragonBreathChili = makeBooleanConsumeInput({
	actionId: () => ActionId.fromItemId(12217),
	fieldName: 'dragonBreathChili',
});

export const RumseyRumBlackLabel: ConsumableInputConfig<Alcohol> = {
	actionId: () => ActionId.fromItemId(21151),
	value: Alcohol.AlcoholRumseyRumBlackLabel,
};
export const GordokGreenGrog: ConsumableInputConfig<Alcohol> = {
	actionId: () => ActionId.fromItemId(18269),
	value: Alcohol.AlcoholGordokGreenGrog,
};
export const RumseyRumDark: ConsumableInputConfig<Alcohol> = {
	actionId: () => ActionId.fromItemId(21114),
	value: Alcohol.AlcoholRumseyRumDark,
};
export const RumseyRumLight: ConsumableInputConfig<Alcohol> = {
	actionId: () => ActionId.fromItemId(20709),
	value: Alcohol.AlcoholRumseyRumLight,
};
export const KreegsStoutBeatdown: ConsumableInputConfig<Alcohol> = {
	actionId: () => ActionId.fromItemId(18284),
	value: Alcohol.AlcoholKreegsStoutBeatdown,
};

export const ALCOHOL_CONFIG: ConsumableStatOption<Alcohol>[] = [
	{ config: RumseyRumBlackLabel, stats: [Stat.StatStamina] },
	{ config: GordokGreenGrog, stats: [Stat.StatStamina] },
	{ config: RumseyRumDark, stats: [Stat.StatStamina] },
	{ config: RumseyRumLight, stats: [Stat.StatStamina] },
	{ config: KreegsStoutBeatdown, stats: [Stat.StatSpirit] },
];

export const makeAlcoholInput = makeConsumeInputFactory({ consumesFieldName: 'alcohol' });

///////////////////////////////////////////////////////////////////////////
//                                 DEFENSIVE CONSUMES
///////////////////////////////////////////////////////////////////////////

// Armor
export const ElixirOfSuperiorDefense: ConsumableInputConfig<ArmorElixir> = {
	actionId: () => ActionId.fromItemId(13445),
	value: ArmorElixir.ElixirOfSuperiorDefense,
};
export const ElixirOfGreaterDefense: ConsumableInputConfig<ArmorElixir> = {
	actionId: () => ActionId.fromItemId(8951),
	value: ArmorElixir.ElixirOfGreaterDefense,
};
export const ElixirOfDefense: ConsumableInputConfig<ArmorElixir> = {
	actionId: () => ActionId.fromItemId(3389),
	value: ArmorElixir.ElixirOfDefense,
};
export const ElixirOfMinorDefense: ConsumableInputConfig<ArmorElixir> = {
	actionId: () => ActionId.fromItemId(5997),
	value: ArmorElixir.ElixirOfMinorDefense,
};
export const ScrollOfProtection: ConsumableInputConfig<ArmorElixir> = {
	actionId: () => ActionId.fromItemId(10305),
	value: ArmorElixir.ScrollOfProtection,
};
export const ARMOR_CONSUMES_CONFIG: ConsumableStatOption<ArmorElixir>[] = [
	{ config: ElixirOfSuperiorDefense, stats: [Stat.StatArmor] },
	{ config: ElixirOfGreaterDefense, stats: [Stat.StatArmor] },
	{ config: ElixirOfDefense, stats: [Stat.StatArmor] },
	{ config: ElixirOfMinorDefense, stats: [Stat.StatArmor] },
	{ config: ScrollOfProtection, stats: [Stat.StatArmor] },
];

export const makeArmorConsumeInput = makeConsumeInputFactory({ consumesFieldName: 'armorElixir' });

// Health
export const ElixirOfFortitude: ConsumableInputConfig<HealthElixir> = {
	actionId: () => ActionId.fromItemId(3825),
	value: HealthElixir.ElixirOfFortitude,
};
export const ElixirOfMinorFortitude: ConsumableInputConfig<HealthElixir> = {
	actionId: () => ActionId.fromItemId(2458),
	value: HealthElixir.ElixirOfMinorFortitude,
};
export const HEALTH_CONSUMES_CONFIG: ConsumableStatOption<HealthElixir>[] = [
	{ config: ElixirOfFortitude, stats: [Stat.StatStamina] },
	{ config: ElixirOfMinorFortitude, stats: [Stat.StatStamina] },
];

export const makeHealthConsumeInput = makeConsumeInputFactory({ consumesFieldName: 'healthElixir' });

///////////////////////////////////////////////////////////////////////////
//                                 PHYSICAL DAMAGE CONSUMES
///////////////////////////////////////////////////////////////////////////

// Attack Power
export const JujuMight: ConsumableInputConfig<AttackPowerBuff> = {
	actionId: () => ActionId.fromItemId(12460),
	value: AttackPowerBuff.JujuMight,
};
export const WinterfallFirewater: ConsumableInputConfig<AttackPowerBuff> = {
	actionId: () => ActionId.fromItemId(12820),
	value: AttackPowerBuff.WinterfallFirewater,
};

export const ATTACK_POWER_CONSUMES_CONFIG: ConsumableStatOption<AttackPowerBuff>[] = [
	{ config: JujuMight, stats: [Stat.StatAttackPower] },
	{ config: WinterfallFirewater, stats: [Stat.StatAttackPower] },
];

export const makeAttackPowerConsumeInput = makeConsumeInputFactory({ consumesFieldName: 'attackPowerBuff' });

// Agility
export const ElixirOfTheMongoose: ConsumableInputConfig<AgilityElixir> = {
	actionId: () => ActionId.fromItemId(13452),
	value: AgilityElixir.ElixirOfTheMongoose,
};
export const ElixirOfGreaterAgility: ConsumableInputConfig<AgilityElixir> = {
	actionId: () => ActionId.fromItemId(9187),
	value: AgilityElixir.ElixirOfGreaterAgility,
};
export const ElixirOfAgility: ConsumableInputConfig<AgilityElixir> = {
	actionId: () => ActionId.fromItemId(8949),
	value: AgilityElixir.ElixirOfAgility,
};
export const ElixirOfLesserAgility: ConsumableInputConfig<AgilityElixir> = {
	actionId: () => ActionId.fromItemId(3390),
	value: AgilityElixir.ElixirOfLesserAgility,
};
export const ScrollOfAgility: ConsumableInputConfig<AgilityElixir> = {
	actionId: () => ActionId.fromItemId(10309),
	value: AgilityElixir.ScrollOfAgility,
};

export const AGILITY_CONSUMES_CONFIG: ConsumableStatOption<AgilityElixir>[] = [
	{ config: ElixirOfTheMongoose, stats: [Stat.StatAgility, Stat.StatMeleeCrit] },
	{ config: ElixirOfGreaterAgility, stats: [Stat.StatAgility] },
	{ config: ElixirOfAgility, stats: [Stat.StatAgility] },
	{ config: ElixirOfLesserAgility, stats: [Stat.StatAgility] },
	{ config: ScrollOfAgility, stats: [Stat.StatAgility] },
];

export const makeAgilityConsumeInput = makeConsumeInputFactory({ consumesFieldName: 'agilityElixir' });

// Strength
export const JujuPower: ConsumableInputConfig<StrengthBuff> = {
	actionId: () => ActionId.fromItemId(12451),
	value: StrengthBuff.JujuPower,
};
export const ElixirOfGiants: ConsumableInputConfig<StrengthBuff> = {
	actionId: () => ActionId.fromItemId(9206),
	value: StrengthBuff.ElixirOfGiants,
};
export const ElixirOfOgresStrength: ConsumableInputConfig<StrengthBuff> = {
	actionId: () => ActionId.fromItemId(3391),
	value: StrengthBuff.ElixirOfOgresStrength,
};
export const ScrollOfStrength: ConsumableInputConfig<StrengthBuff> = {
	actionId: () => ActionId.fromItemId(10310),
	value: StrengthBuff.ScrollOfStrength,
};

export const STRENGTH_CONSUMES_CONFIG: ConsumableStatOption<StrengthBuff>[] = [
	{ config: JujuPower, stats: [Stat.StatStrength] },
	{ config: ElixirOfGiants, stats: [Stat.StatStrength] },
	{ config: ElixirOfOgresStrength, stats: [Stat.StatStrength] },
	{ config: ScrollOfStrength, stats: [Stat.StatStrength] },
];

export const makeStrengthConsumeInput = makeConsumeInputFactory({ consumesFieldName: 'strengthBuff' });

///////////////////////////////////////////////////////////////////////////
//                                 Misc Throughput Consumes
///////////////////////////////////////////////////////////////////////////

// Blasted Lands Consumes
export const ROIDS: ConsumableInputConfig<ZanzaBuff> = {
	actionId: () => ActionId.fromItemId(8410),
	value: ZanzaBuff.ROIDS,
};
export const GroundScorpokAssay: ConsumableInputConfig<ZanzaBuff> = {
	actionId: () => ActionId.fromItemId(8412),
	value: ZanzaBuff.GroundScorpokAssay,
};
export const LungJuiceCocktail: ConsumableInputConfig<ZanzaBuff> = {
	actionId: () => ActionId.fromItemId(8411),
	value: ZanzaBuff.LungJuiceCocktail,
};
export const CerebralCortexCompound: ConsumableInputConfig<ZanzaBuff> = {
	actionId: () => ActionId.fromItemId(8423),
	value: ZanzaBuff.CerebralCortexCompound,
};
export const GizzardGum: ConsumableInputConfig<ZanzaBuff> = {
	actionId: () => ActionId.fromItemId(8424),
	value: ZanzaBuff.GizzardGum,
};

// Zanza Potions
export const SpiritOfZanza: ConsumableInputConfig<ZanzaBuff> = {
	actionId: () => ActionId.fromItemId(20079),
	value: ZanzaBuff.SpiritOfZanza,
};

export const ZANZA_BUFF_CONSUMES_CONFIG: ConsumableStatOption<ZanzaBuff>[] = [
	{ config: SpiritOfZanza, stats: [Stat.StatStamina, Stat.StatSpirit] },
	{ config: ROIDS, stats: [Stat.StatStrength] },
	{ config: GroundScorpokAssay, stats: [Stat.StatAgility] },
	{ config: LungJuiceCocktail, stats: [Stat.StatStamina] },
	{ config: CerebralCortexCompound, stats: [Stat.StatIntellect] },
	{ config: GizzardGum, stats: [Stat.StatSpirit] },
];
export const makeZanzaBuffConsumesInput = makeConsumeInputFactory({ consumesFieldName: 'zanzaBuff' });

export const JujuFlurry = makeBooleanMiscConsumeInput({
	actionId: () => ActionId.fromItemId(12450),
	fieldName: 'jujuFlurry',
});
export const BoglingRoot = makeBooleanMiscConsumeInput({ actionId: () => ActionId.fromItemId(5206), fieldName: 'boglingRoot' });
export const RaptorPunch = makeBooleanMiscConsumeInput({
	actionId: () => ActionId.fromItemId(5342),
	fieldName: 'raptorPunch',
});

export const MISC_OFFENSIVE_CONSUMES_CONFIG: PickerStatOptions[] = [
	{ config: JujuFlurry, picker: IconPicker, stats: [Stat.StatAttackPower] },
	{ config: RaptorPunch, picker: IconPicker, stats: [Stat.StatIntellect] },
	{ config: BoglingRoot, picker: IconPicker, stats: [Stat.StatAttackPower] },
];

export const makeMiscOffensiveConsumesInput = makeMultiIconConsumesInputFactory({
	direction: IconPickerDirection.Vertical,
	tooltip: 'Misc Offensive',
});

export const JujuEmber = makeBooleanMiscConsumeInput({
	actionId: () => ActionId.fromItemId(12455),
	fieldName: 'jujuEmber',
});
export const JujuChill = makeBooleanMiscConsumeInput({
	actionId: () => ActionId.fromItemId(12457),
	fieldName: 'jujuChill',
});
export const JujuEscape = makeBooleanMiscConsumeInput({
	actionId: () => ActionId.fromItemId(12459),
	fieldName: 'jujuEscape',
});

export const MISC_DEFENSIVE_CONSUMES_CONFIG: PickerStatOptions[] = [
	{ config: JujuEmber, picker: IconPicker, stats: [] },
	{ config: JujuChill, picker: IconPicker, stats: [] },
	{ config: JujuEscape, picker: IconPicker, stats: [Stat.StatDodge] },
];

export const makeMiscDefensiveConsumesInput = makeMultiIconConsumesInputFactory({
	direction: IconPickerDirection.Vertical,
	tooltip: 'Misc Defensive',
});

///////////////////////////////////////////////////////////////////////////
//                                 PET
///////////////////////////////////////////////////////////////////////////

export const PetAttackPowerConsumable = makeEnumConsumeInput({
	direction: IconPickerDirection.Vertical,
	values: [
		{ value: 0, tooltip: 'None' },
		{ actionId: () => ActionId.fromItemId(12460), value: 1 },
	],
	fieldName: 'petAttackPowerConsumable',
});

export const PetAgilityConsumable = makeEnumConsumeInput({
	direction: IconPickerDirection.Vertical,
	values: [
		{ value: 0, tooltip: 'None' },
		{ actionId: () => ActionId.fromItemId(10309), value: 1 },
		{ actionId: () => ActionId.fromItemId(4425), value: 2 },
		{ actionId: () => ActionId.fromItemId(1477), value: 3 },
		{ actionId: () => ActionId.fromItemId(3012), value: 4 },
	],
	fieldName: 'petAgilityConsumable',
});

export const PetStrengthConsumable = makeEnumConsumeInput({
	direction: IconPickerDirection.Vertical,
	values: [
		{ value: 0, tooltip: 'None' },
		{ actionId: () => ActionId.fromItemId(12451), value: 1 },
		{ actionId: () => ActionId.fromItemId(10310), value: 2 },
		{ actionId: () => ActionId.fromItemId(4426), value: 3 },
		{ actionId: () => ActionId.fromItemId(2289), value: 4 },
		{ actionId: () => ActionId.fromItemId(954), value: 5 },
	],
	fieldName: 'petStrengthConsumable',
});

export const JujuFlurryPet = makeBooleanPetMiscConsumeInput({
	actionId: () => ActionId.fromItemId(12450),
	fieldName: 'jujuFlurry',
});

export const MISC_PET_CONSUMES: PickerStatOptions[] = [{ config: JujuFlurryPet, picker: IconPicker, stats: [] }];

export const makeMiscPetConsumesInput = makeMultiIconConsumesInputFactory({
	direction: IconPickerDirection.Vertical,
	tooltip: 'Misc Pet Consumes',
});

///////////////////////////////////////////////////////////////////////////
//                                 POTIONS
///////////////////////////////////////////////////////////////////////////

export const GreaterHealingPotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(1710),
	value: Potions.GreaterHealingPotion,
};
export const SuperiorHealingPotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(3928),
	value: Potions.SuperiorHealingPotion,
};
export const MajorHealingPotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(13446),
	value: Potions.MajorHealingPotion,
};

export const ManaPotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(3827),
	value: Potions.ManaPotion,
};
export const GreaterManaPotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(6149),
	value: Potions.GreaterManaPotion,
};
export const SuperiorManaPotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(13443),
	value: Potions.SuperiorManaPotion,
};
export const MajorManaPotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(13444),
	value: Potions.MajorManaPotion,
};

export const MightRagePotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(13442),
	value: Potions.MightyRagePotion,
	showWhen: player => player.getClass() == Class.ClassWarrior,
};
export const GreatRagePotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(5633),
	value: Potions.GreatRagePotion,
	showWhen: player => player.getClass() == Class.ClassWarrior,
};
export const RagePotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(5631),
	value: Potions.RagePotion,
	showWhen: player => player.getClass() == Class.ClassWarrior,
};

export const MagicResistancePotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(9036),
	value: Potions.MagicResistancePotion,
};
// TODO: Not yet implemented in the back-end. Missing school shields and shields don't actually absorb damage right now
// export const GreaterArcaneProtectionPotion: ConsumableInputConfig<Potions> = {
// 	actionId: () => ActionId.fromItemId(13461),
// 	value: Potions.GreaterArcaneProtectionPotion,
// };
// export const GreaterFireProtectionPotion: ConsumableInputConfig<Potions> = {
// 	actionId: () => ActionId.fromItemId(13457),
// 	value: Potions.GreaterFireProtectionPotion,
// };
// export const GreaterFrostProtectionPotion: ConsumableInputConfig<Potions> = {
// 	actionId: () => ActionId.fromItemId(13456),
// 	value: Potions.GreaterFrostProtectionPotion,
// };
// export const GreaterHolyProtectionPotion: ConsumableInputConfig<Potions> = {
// 	actionId: () => ActionId.fromItemId(13460),
// 	value: Potions.GreaterHolyProtectionPotion,
// };
// export const GreaterNatureProtectionPotion: ConsumableInputConfig<Potions> = {
// 	actionId: () => ActionId.fromItemId(13458),
// 	value: Potions.GreaterNatureProtectionPotion,
// };
// export const GreaterShadowProtectionPotion: ConsumableInputConfig<Potions> = {
// 	actionId: () => ActionId.fromItemId(13459),
// 	value: Potions.GreaterShadowProtectionPotion,
// };

export const GreaterStoneshieldPotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(13455),
	value: Potions.GreaterStoneshieldPotion,
};
export const LesserStoneshieldPotion: ConsumableInputConfig<Potions> = {
	actionId: () => ActionId.fromItemId(4623),
	value: Potions.LesserStoneshieldPotion,
};

export const POTIONS_CONFIG: ConsumableStatOption<Potions>[] = [
	{ config: MajorHealingPotion, stats: [Stat.StatArmor] },
	{ config: SuperiorHealingPotion, stats: [Stat.StatArmor] },
	{ config: GreaterHealingPotion, stats: [Stat.StatArmor] },

	{ config: MajorManaPotion, stats: [Stat.StatIntellect] },
	{ config: SuperiorManaPotion, stats: [Stat.StatIntellect] },
	{ config: GreaterManaPotion, stats: [Stat.StatIntellect] },
	{ config: ManaPotion, stats: [Stat.StatIntellect] },

	{ config: MightRagePotion, stats: [] },
	{ config: GreatRagePotion, stats: [] },
	{ config: RagePotion, stats: [] },

	// { config: MagicResistancePotion, stats: [] },
	{ config: GreaterStoneshieldPotion, stats: [Stat.StatArmor] },
	{ config: LesserStoneshieldPotion, stats: [Stat.StatArmor] },
];

export const makePotionsInput = makeConsumeInputFactory({ consumesFieldName: 'defaultPotion' });

///////////////////////////////////////////////////////////////////////////
//                                 SPELL DAMAGE CONSUMES
///////////////////////////////////////////////////////////////////////////

// Arcane
export const GreaterArcaneElixir: ConsumableInputConfig<SpellPowerBuff> = {
	actionId: () => ActionId.fromItemId(13454),
	value: SpellPowerBuff.GreaterArcaneElixir,
};
export const ArcaneElixir: ConsumableInputConfig<SpellPowerBuff> = {
	actionId: () => ActionId.fromItemId(9155),
	value: SpellPowerBuff.ArcaneElixir,
};
export const LesserArcaneElixir: ConsumableInputConfig<SpellPowerBuff> = {
	actionId: () => ActionId.fromItemId(217398),
	value: SpellPowerBuff.LesserArcaneElixir,
};

export const SPELL_POWER_CONFIG: ConsumableStatOption<SpellPowerBuff>[] = [
	{ config: GreaterArcaneElixir, stats: [Stat.StatSpellPower] },
	{ config: ArcaneElixir, stats: [Stat.StatSpellPower] },
	{ config: LesserArcaneElixir, stats: [Stat.StatSpellPower] },
];

export const makeSpellPowerConsumeInput = makeConsumeInputFactory({ consumesFieldName: 'spellPowerBuff' });

// Fire
// Original lvl 40 not obtainable in Phase 3
export const ElixirOfGreaterFirepower: ConsumableInputConfig<FirePowerBuff> = {
	actionId: () => ActionId.fromItemId(21546),
	value: FirePowerBuff.ElixirOfGreaterFirepower,
};
export const ElixirOfFirepower: ConsumableInputConfig<FirePowerBuff> = {
	actionId: () => ActionId.fromItemId(6373),
	value: FirePowerBuff.ElixirOfFirepower,
};

export const FIRE_POWER_CONFIG: ConsumableStatOption<FirePowerBuff>[] = [
	{ config: ElixirOfGreaterFirepower, stats: [Stat.StatFirePower] },
	{ config: ElixirOfFirepower, stats: [Stat.StatFirePower] },
];

export const makeFirePowerConsumeInput = makeConsumeInputFactory({ consumesFieldName: 'firePowerBuff' });

// Frost
export const ElixirOfFrostPower: ConsumableInputConfig<FrostPowerBuff> = {
	actionId: () => ActionId.fromItemId(17708),
	value: FrostPowerBuff.ElixirOfFrostPower,
};

export const FROST_POWER_CONFIG: ConsumableStatOption<FrostPowerBuff>[] = [{ config: ElixirOfFrostPower, stats: [Stat.StatFrostPower] }];

export const makeFrostPowerConsumeInput = makeConsumeInputFactory({ consumesFieldName: 'frostPowerBuff' });

// Shadow
export const ElixirOfShadowPower: ConsumableInputConfig<ShadowPowerBuff> = {
	actionId: () => ActionId.fromItemId(9264),
	value: ShadowPowerBuff.ElixirOfShadowPower,
};

export const SHADOW_POWER_CONFIG: ConsumableStatOption<ShadowPowerBuff>[] = [{ config: ElixirOfShadowPower, stats: [Stat.StatShadowPower] }];

export const makeShadowPowerConsumeInput = makeConsumeInputFactory({ consumesFieldName: 'shadowPowerBuff' });

// MP5
export const MagebloodPotion: ConsumableInputConfig<ManaRegenElixir> = {
	actionId: () => ActionId.fromItemId(20007),
	value: ManaRegenElixir.MagebloodPotion,
};

export const MP5_CONFIG: ConsumableStatOption<ManaRegenElixir>[] = [{ config: MagebloodPotion, stats: [Stat.StatMP5] }];

export const makeMp5ConsumeInput = makeConsumeInputFactory({ consumesFieldName: 'manaRegenElixir' });

///////////////////////////////////////////////////////////////////////////
//                                 Weapon Imbues
///////////////////////////////////////////////////////////////////////////

// Windfury (Buff)
export const Windfury: ConsumableInputConfig<WeaponImbue> = {
	actionId: () => ActionId.fromSpellId(10614),
	value: WeaponImbue.Windfury,
	showWhen: player => {
		return (player.getFaction() === Faction.Horde) && !player.isSpec(Spec.SpecFeralDruid)
	},
};

// Other Imbues

// Wizard Oils
export const BrilliantWizardOil = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(20749),
		value: WeaponImbue.BrilliantWizardOil,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};
export const WizardOil = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(20750),
		value: WeaponImbue.WizardOil,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};
export const LesserWizardOil = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(20746),
		value: WeaponImbue.LesserWizardOil,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};
export const MinorWizardOil = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(20744),
		value: WeaponImbue.MinorWizardOil,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};
export const BlessedWizardOil = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(23123),
		value: WeaponImbue.BlessedWizardOil,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};

// Mana Oils
// Original lvl 45 but not obtainable in Phase 3
export const BrilliantManaOil = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(20748),
		value: WeaponImbue.BrilliantManaOil,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};
export const LesserManaOil = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(20747),
		value: WeaponImbue.LesserManaOil,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};
export const MinorManaOil = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(20745),
		value: WeaponImbue.MinorManaOil,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};

// Sharpening Stones
export const ConsecratedSharpeningStone = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(23122),
		value: WeaponImbue.ConsecratedSharpeningStone,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return !weapon || isWeapon(weapon.item.weaponType);
		},
	};
};
export const ElementalSharpeningStone = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(18262),
		value: WeaponImbue.ElementalSharpeningStone,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return (!weapon || isWeapon(weapon.item.weaponType)) && !player.isSpec(Spec.SpecFeralDruid);
		},
	};
};
export const DenseSharpeningStone = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(12404),
		value: WeaponImbue.DenseSharpeningStone,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return (!weapon || isSharpWeaponType(weapon.item.weaponType)) && !player.isSpec(Spec.SpecFeralDruid);
		},
	};
};
export const SolidSharpeningStone = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(7964),
		value: WeaponImbue.SolidSharpeningStone,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return (!weapon || isSharpWeaponType(weapon.item.weaponType)) && !player.isSpec(Spec.SpecFeralDruid);
		},
	};
};

// Weightstones
export const DenseWeightstone = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(12643),
		value: WeaponImbue.DenseWeightstone,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return (!weapon || isBluntWeaponType(weapon.item.weaponType)) && !player.isSpec(Spec.SpecFeralDruid);
		},
	};
};
export const SolidWeightstone = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(7965),
		value: WeaponImbue.SolidWeightstone,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return (!weapon || isBluntWeaponType(weapon.item.weaponType)) && !player.isSpec(Spec.SpecFeralDruid);
		},
	};
};

// Spell Oils
export const ShadowOil = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(3824),
		value: WeaponImbue.ShadowOil,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return (!weapon || isWeapon(weapon.item.weaponType)) && !player.isSpec(Spec.SpecFeralDruid);
		},
	};
};
export const FrostOil = (slot: ItemSlot): ConsumableInputConfig<WeaponImbue> => {
	return {
		actionId: () => ActionId.fromItemId(3829),
		value: WeaponImbue.FrostOil,
		showWhen: player => {
			const weapon = player.getEquippedItem(slot);
			return (!weapon || isWeapon(weapon.item.weaponType)) && !player.isSpec(Spec.SpecFeralDruid);
		},
	};
};

const SHAMAN_IMBUES = (slot: ItemSlot): ConsumableStatOption<WeaponImbue>[] => [
	{ config: RockbiterWeaponImbue(slot), stats: [] },
	{ config: FlametongueWeaponImbue(slot), stats: [] },
	{ config: FrostbrandWeaponImbue(slot), stats: [] },
	{ config: WindfuryWeaponImbue(slot), stats: [] },
];

const ROGUE_IMBUES: ConsumableStatOption<WeaponImbue>[] = [
	{ config: InstantPoisonWeaponImbue, stats: [] },
	{ config: DeadlyPoisonWeaponImbue, stats: [] },
	{ config: WoundPoisonWeaponImbue, stats: [] },
];

const CONSUMABLES_IMBUES = (slot: ItemSlot): ConsumableStatOption<WeaponImbue>[] => [
	{ config: BrilliantWizardOil(slot), stats: [Stat.StatSpellPower] },
	{ config: WizardOil(slot), stats: [Stat.StatSpellPower] },
	{ config: LesserWizardOil(slot), stats: [Stat.StatSpellPower] },
	{ config: MinorWizardOil(slot), stats: [Stat.StatSpellPower] },
	{ config: BlessedWizardOil(slot), stats: [Stat.StatHealingPower, Stat.StatSpellPower] },

	{ config: BrilliantManaOil(slot), stats: [Stat.StatMP5, Stat.StatHealingPower] },
	{ config: LesserManaOil(slot), stats: [Stat.StatMP5] },
	{ config: MinorManaOil(slot), stats: [Stat.StatMP5] },

	{ config: ConsecratedSharpeningStone(slot), stats: [Stat.StatAttackPower] },
	{ config: ElementalSharpeningStone(slot), stats: [Stat.StatAttackPower] },
	{ config: DenseSharpeningStone(slot), stats: [Stat.StatAttackPower] },
	{ config: SolidSharpeningStone(slot), stats: [Stat.StatAttackPower] },

	{ config: DenseWeightstone(slot), stats: [Stat.StatAttackPower] },
	{ config: SolidWeightstone(slot), stats: [Stat.StatAttackPower] },

	{ config: ShadowOil(slot), stats: [Stat.StatAttackPower] },
	{ config: FrostOil(slot), stats: [Stat.StatAttackPower] },
];

export const WEAPON_IMBUES_OH_CONFIG: ConsumableStatOption<WeaponImbue>[] = [
	...ROGUE_IMBUES,
	...SHAMAN_IMBUES(ItemSlot.ItemSlotOffHand),
	...CONSUMABLES_IMBUES(ItemSlot.ItemSlotOffHand),
];

export const WEAPON_IMBUES_MH_CONFIG: ConsumableStatOption<WeaponImbue>[] = [
	...ROGUE_IMBUES,
	...SHAMAN_IMBUES(ItemSlot.ItemSlotMainHand),
	{ config: Windfury, stats: [Stat.StatMeleeHit] },
	...CONSUMABLES_IMBUES(ItemSlot.ItemSlotMainHand),
];

export const makeMainHandImbuesInput = makeConsumeInputFactory({
	consumesFieldName: 'mainHandImbue',
	showWhen: player => !!player.getGear().getEquippedItem(ItemSlot.ItemSlotMainHand),
});
export const makeOffHandImbuesInput = makeConsumeInputFactory({
	consumesFieldName: 'offHandImbue',
	showWhen: player => !!player.getGear().getEquippedItem(ItemSlot.ItemSlotOffHand),
});

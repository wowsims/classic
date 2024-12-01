import { Faction, SaygesFortune, Stat } from '../../proto/common';
import { ActionId } from '../../proto_utils/action_id';
import {
	makeBooleanDebuffInput,
	makeBooleanIndividualBuffInput,
	makeBooleanRaidBuffInput,
	makeEnumIndividualBuffInput,
	makeMultistateIndividualBuffInput,
	makeMultistateMultiplierDebuffInput,
	makeMultistateRaidBuffInput,
	makeTristateDebuffInput,
	makeTristateIndividualBuffInput,
	makeTristateRaidBuffInput,
	withLabel,
} from '../icon_inputs';
import { IconPicker, IconPickerDirection } from '../icon_picker';
import * as InputHelpers from '../input_helpers';
import { MultiIconPicker } from '../multi_icon_picker';
import { ItemStatOption, PickerStatOptions } from './stat_options';

///////////////////////////////////////////////////////////////////////////
//                                 RAID BUFFS
///////////////////////////////////////////////////////////////////////////

export const AllStatsBuff = withLabel(
	makeTristateRaidBuffInput({
		actionId: () => ActionId.fromSpellId(9885),
		impId: ActionId.fromSpellId(17055),
		fieldName: 'giftOfTheWild',
	}),
	'Mark of the Wild',
);

// Separate Strength buffs allow us to use a boolean pickers for Horde specifically
export const AllStatsPercentBuff = InputHelpers.makeMultiIconInput({
	values: [
		makeBooleanIndividualBuffInput({
			actionId: () => ActionId.fromSpellId(20217),
			fieldName: 'blessingOfKings',
			showWhen: player => player.getFaction() === Faction.Alliance,
		}),
	],
	label: 'Stats %',
});

export const ArmorBuff = withLabel(
	makeTristateRaidBuffInput({
		actionId: () => ActionId.fromSpellId(10293),
		impId: ActionId.fromSpellId(20142),
		showWhen: player => player.getFaction() === Faction.Alliance,
		fieldName: 'devotionAura',
	}),
	'Armor',
);

export const PhysDamReductionBuff = withLabel(
	makeTristateRaidBuffInput({
		actionId: () => ActionId.fromSpellId(10408),
		impId: ActionId.fromSpellId(16293),
		showWhen: player => player.getFaction() === Faction.Horde,
		fieldName: 'stoneskinTotem',
	}),
	'Stoneskin',
);

//export const DamageReductionPercentBuff = withLabel(
//	makeBooleanIndividualBuffInput({
//		actionId: player =>
//			player.getMatchingSpellActionId([
//				{ id: 20911, minLevel: 30, maxLevel: 39 },
//				{ id: 20912, minLevel: 40, maxLevel: 49 },
//				{ id: 20913, minLevel: 50, maxLevel: 59 },
//				{ id: 20914, minLevel: 60 },
//			]),
//		showWhen: player => player.getFaction() === Faction.Alliance,
//		fieldName: 'blessingOfSanctuary',
//	}),
//	'Blessing of Sanctuary',
//);

export const ResistanceBuff = InputHelpers.makeMultiIconInput({
	values: [
		// Shadow
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(10958),
			fieldName: 'shadowProtection',
		}),
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(19896),
			fieldName: 'shadowResistanceAura',
		}),
		// Nature
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(10601),
			fieldName: 'natureResistanceTotem',
			showWhen: player => player.getFaction() === Faction.Horde,
		}),
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(20190),
			fieldName: 'aspectOfTheWild',
		}),
		// Fire
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(19900),
			fieldName: 'fireResistanceAura',
			showWhen: player => player.getFaction() === Faction.Alliance,
		}),
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(10538),
			fieldName: 'fireResistanceTotem',
			showWhen: player => player.getFaction() === Faction.Horde,
		}),
		// Frost
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(19898),
			fieldName: 'frostResistanceAura',
			showWhen: player => player.getFaction() === Faction.Alliance,
		}),
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(10479),
			fieldName: 'frostResistanceTotem',
			showWhen: player => player.getFaction() === Faction.Horde,
		}),
	],
	label: 'Resistances',
});

export const StaminaBuff = InputHelpers.makeMultiIconInput({
	values: [
		makeTristateRaidBuffInput({
			actionId: () => ActionId.fromSpellId(10938),
			impId: ActionId.fromSpellId(14767),
			fieldName: 'powerWordFortitude',
		}),
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(10307),
			fieldName: 'scrollOfStamina',
		}),
	],
	label: 'Stamina',
});

export const BloodPactBuff = InputHelpers.makeMultiIconInput({
	values: [
		makeTristateRaidBuffInput({
			actionId: () => ActionId.fromSpellId(11767),
			impId: ActionId.fromSpellId(18696),
			fieldName: 'bloodPact',
		}),
	],
	label: 'Blood Pact',
});

export const PaladinPhysicalBuff = InputHelpers.makeMultiIconInput({
	values: [
		makeTristateIndividualBuffInput({
			actionId: () => ActionId.fromSpellId(25291),
			impId: ActionId.fromSpellId(20048),
			fieldName: 'blessingOfMight',
			showWhen: player => player.getFaction() === Faction.Alliance,
		}),
	],
	label: 'Paladin Physical',
});

export const StrengthBuffHorde = withLabel(
	makeTristateRaidBuffInput({
		actionId: () => ActionId.fromSpellId(25361),
		impId: ActionId.fromSpellId(16295),
		fieldName: 'strengthOfEarthTotem',
		showWhen: player => player.getFaction() === Faction.Horde,
	}),
	'Strength',
);

export const GraceOfAir = withLabel(
	makeTristateRaidBuffInput({
		actionId: () => ActionId.fromSpellId(25359),
		impId: ActionId.fromSpellId(16295),
		fieldName: 'graceOfAirTotem',
		showWhen: player => player.getFaction() === Faction.Horde,
	}),
	'Agility',
);

export const IntellectBuff = InputHelpers.makeMultiIconInput({
	values: [
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(10157),
			fieldName: 'arcaneBrilliance',
		}),
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(10308),
			fieldName: 'scrollOfIntellect',
		}),
	],
	label: 'Intellect',
});

export const SpiritBuff = InputHelpers.makeMultiIconInput({
	values: [
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(27841),
			fieldName: 'divineSpirit',
		}),
		makeBooleanRaidBuffInput({
			actionId: () => ActionId.fromSpellId(10306),
			fieldName: 'scrollOfSpirit',
		}),
	],
	label: 'Spirit',
});

export const BattleShoutBuff = withLabel(
	makeTristateRaidBuffInput({
		actionId: () => ActionId.fromSpellId(25289),
		impId: ActionId.fromSpellId(12861),
		fieldName: 'battleShout',
	}),
	'Battle Shout',
);

export const TrueshotAuraBuff = withLabel(
	makeBooleanRaidBuffInput({
		actionId: () => ActionId.fromSpellId(20906),
		fieldName: 'trueshotAura',
	}),
	'Trueshot Aura',
);

export const BlessingOfWisdom = withLabel(
	makeTristateIndividualBuffInput({
		actionId: () => ActionId.fromSpellId(25290),
		impId: ActionId.fromSpellId(20245),
		fieldName: 'blessingOfWisdom',
		showWhen: player => player.getFaction() === Faction.Alliance,
	}),
	'Blessing of Wisdom',
);
export const ManaSpringTotem = withLabel(
	makeTristateRaidBuffInput({
		actionId: () => ActionId.fromSpellId(10497),
		impId: ActionId.fromSpellId(16208),
		fieldName: 'manaSpringTotem',
		showWhen: player => player.getFaction() === Faction.Horde,
	}),
	'Mana Spring Totem',
);

export const MeleeCritBuff = withLabel(
	makeBooleanRaidBuffInput({ actionId: () => ActionId.fromSpellId(24932), fieldName: 'leaderOfThePack' }),
	'Leader of the Pack',
);

export const SpellCritBuff = withLabel(
	makeBooleanRaidBuffInput({ actionId: () => ActionId.fromSpellId(24907), fieldName: 'moonkinAura' }),
	'Moonkin Aura',
);

// Misc Buffs
export const RetributionAura = makeTristateRaidBuffInput({
	actionId: () => ActionId.fromSpellId(10301),
	impId: ActionId.fromSpellId(20092),
	fieldName: 'retributionAura',
	showWhen: player => player.getFaction() === Faction.Alliance,
});

export const SanctityAura = makeBooleanRaidBuffInput({
	actionId: () => ActionId.fromSpellId(20218),
	fieldName: 'sanctityAura',
	showWhen: player => player.getFaction() === Faction.Alliance,
});

export const Thorns = makeTristateRaidBuffInput({
	actionId: () => ActionId.fromSpellId(9910),
	impId: ActionId.fromSpellId(16840),
	fieldName: 'thorns',
});

export const Innervate = makeMultistateIndividualBuffInput({
	actionId: () => ActionId.fromSpellId(29166),
	numStates: 11,
	fieldName: 'innervates',
});

export const PowerInfusion = makeMultistateIndividualBuffInput({
	actionId: () => ActionId.fromSpellId(10060),
	numStates: 11,
	fieldName: 'powerInfusions',
});

export const BattleSquawkBuff = makeMultistateRaidBuffInput({
	actionId: () => ActionId.fromSpellId(23060),
	numStates: 6,
	fieldName: 'battleSquawk',
});

///////////////////////////////////////////////////////////////////////////
//                                 WORLD BUFFS
///////////////////////////////////////////////////////////////////////////

export const RallyingCryOfTheDragonslayer = makeBooleanIndividualBuffInput({
	actionId: () => ActionId.fromSpellId(22888),
	fieldName: 'rallyingCryOfTheDragonslayer',
});
export const DragonslayerBuffInput = InputHelpers.makeMultiIconInput({ values: [RallyingCryOfTheDragonslayer], label: 'Dragonslayer Buff' });

export const SpiritOfZandalar = withLabel(
	makeBooleanIndividualBuffInput({
		actionId: () => ActionId.fromSpellId(24425),
		fieldName: 'spiritOfZandalar',
	}),
	'Spirit of Zandalar',
);
export const SongflowerSerenade = withLabel(
	makeBooleanIndividualBuffInput({
		actionId: () => ActionId.fromSpellId(15366),
		fieldName: 'songflowerSerenade',
	}),
	'Songflower Serenade',
);
export const WarchiefsBlessing = withLabel(
	makeBooleanIndividualBuffInput({
		actionId: () => ActionId.fromSpellId(16609),
		fieldName: 'warchiefsBlessing',
		// showWhen: player => player.getFaction() === Faction.Horde,
	}),
	`Warchief's Blessing`,
);

export const SaygesDarkFortune = (inputs: ItemStatOption<SaygesFortune>[]) =>
	makeEnumIndividualBuffInput({
		direction: IconPickerDirection.Horizontal,
		values: [
			{ iconUrl: 'https://wow.zamimg.com/images/wow/icons/large/inv_misc_orb_02.jpg', value: SaygesFortune.SaygesUnknown, text: `Sayge's Dark Fortune` },
			...inputs.map(input => input.config),
		],
		fieldName: 'saygesFortune',
	});

export const SaygesDamage = { actionId: () => ActionId.fromSpellId(23768), value: SaygesFortune.SaygesDamage, text: `Sayge's Damage` };
export const SaygesAgility = { actionId: () => ActionId.fromSpellId(23736), value: SaygesFortune.SaygesAgility, text: `Sayge's Agility` };
export const SaygesIntellect = { actionId: () => ActionId.fromSpellId(23766), value: SaygesFortune.SaygesIntellect, text: `Sayge's Intellect` };
export const SaygesSpirit = { actionId: () => ActionId.fromSpellId(23738), value: SaygesFortune.SaygesSpirit, text: `Sayge's Spirit` };
export const SaygesStamina = { actionId: () => ActionId.fromSpellId(23737), value: SaygesFortune.SaygesStamina, text: `Sayge's Stamina` };

// Dire Maul Buffs
export const FengusFerocity = withLabel(
	makeBooleanIndividualBuffInput({
		actionId: () => ActionId.fromSpellId(22817),
		fieldName: 'fengusFerocity',
	}),
	`Fengus' Ferocity`,
);
export const MoldarsMoxie = withLabel(
	makeBooleanIndividualBuffInput({
		actionId: () => ActionId.fromSpellId(22818),
		fieldName: 'moldarsMoxie',
	}),
	`Moldar's Moxie`,
);
export const SlipKiksSavvy = withLabel(
	makeBooleanIndividualBuffInput({
		actionId: () => ActionId.fromSpellId(22820),
		fieldName: 'slipkiksSavvy',
	}),
	`Slip'kik's Savvy`,
);

///////////////////////////////////////////////////////////////////////////
//                                 DEBUFFS
///////////////////////////////////////////////////////////////////////////

export const MajorArmorDebuff = InputHelpers.makeMultiIconInput({
	values: [
		makeBooleanDebuffInput({
			actionId: () => ActionId.fromSpellId(11597),
			fieldName: 'sunderArmor',
		}),
		makeTristateDebuffInput({
			actionId: () => ActionId.fromSpellId(11198),
			impId: ActionId.fromSpellId(14169),
			fieldName: 'exposeArmor',
		}),
	],
	label: 'Major Armor Penetration',
});

export const CurseOfRecklessness = withLabel(
	makeBooleanDebuffInput({
		actionId: () => ActionId.fromSpellId(11717),
		fieldName: 'curseOfRecklessness',
	}),
	'Curse of Recklessness',
);

export const FaerieFire = withLabel(
	makeBooleanDebuffInput({
		actionId: () => ActionId.fromSpellId(9907),
		fieldName: 'faerieFire',
	}),
	'Faerie Fire',
);

export const curseOfWeaknessDebuff = withLabel(
	makeTristateDebuffInput({
		actionId: () => ActionId.fromSpellId(11708),
		impId: ActionId.fromSpellId(18181),
		fieldName: 'curseOfWeakness',
	}),
	'Curse of Weakness',
);

export const AttackPowerDebuff = InputHelpers.makeMultiIconInput({
	values: [
		makeTristateDebuffInput({
			actionId: () => ActionId.fromSpellId(11556),
			impId: ActionId.fromSpellId(12879),
			fieldName: 'demoralizingShout',
		}),
		makeTristateDebuffInput({
			actionId: () => ActionId.fromSpellId(9898),
			impId: ActionId.fromSpellId(16862),
			fieldName: 'demoralizingRoar',
		}),
	],
	label: 'Attack Power',
});

// TODO: SoD Mangle
//export const BleedDebuff = withLabel(makeBooleanDebuffInput({ actionId: () => ActionId.fromSpellId(409828), fieldName: 'mangle' }), 'Bleed');

export const MeleeAttackSpeedDebuff = InputHelpers.makeMultiIconInput({
	values: [
		makeTristateDebuffInput({
			actionId: () => ActionId.fromSpellId(6343),
			impId: ActionId.fromSpellId(12666),
			fieldName: 'thunderClap',
		}),
		makeBooleanDebuffInput({
			actionId: () => ActionId.fromSpellId(21992),
			fieldName: 'thunderfury',
		}),
	],
	label: 'Attack Speed',
});

export const MeleeHitDebuff = withLabel(
	makeBooleanDebuffInput({
		actionId: () => ActionId.fromSpellId(24977),
		fieldName: 'insectSwarm',
	}),
	'Insect Swarm',
);

export const SpellISBDebuff = withLabel(
	makeBooleanDebuffInput({
		actionId: () => ActionId.fromSpellId(17803),
		fieldName: 'improvedShadowBolt',
	}),
	'Improved Shadow Bolt',
);

export const SpellScorchDebuff = withLabel(
	makeBooleanDebuffInput({
		actionId: () => ActionId.fromSpellId(12873),
		fieldName: 'improvedScorch',
	}),
	'Fire Damage',
);

export const SpellWintersChillDebuff = withLabel(
	makeBooleanDebuffInput({
		actionId: () => ActionId.fromSpellId(28595),
		fieldName: 'wintersChill',
	}),
	'Frost Damage',
);

export const NatureSpellDamageDebuff = InputHelpers.makeMultiIconInput({
	values: [
		makeBooleanDebuffInput({
			actionId: () => ActionId.fromSpellId(17364),
			fieldName: 'stormstrike',
		}),
	],
	label: 'Nature Damage',
});

export const SpellShadowWeavingDebuff = withLabel(
	makeBooleanDebuffInput({
		actionId: () => ActionId.fromSpellId(15334),
		fieldName: 'shadowWeaving',
	}),
	'Shadow Weaving',
);

export const CurseOfElements = makeBooleanDebuffInput({
	actionId: () => ActionId.fromSpellId(11722),
	fieldName: 'curseOfElements',
});

export const CurseOfShadow = makeBooleanDebuffInput({
	actionId: () => ActionId.fromSpellId(17937),
	fieldName: 'curseOfShadow',
});

export const WarlockCursesConfig = InputHelpers.makeMultiIconInput({ values: [CurseOfElements, CurseOfShadow], label: 'Warlock Curses' });

export const HuntersMark = withLabel(
	makeTristateDebuffInput({
		actionId: () => ActionId.fromSpellId(14325),
		impId: ActionId.fromSpellId(19425),
		fieldName: 'huntersMark',
	}),
	`Hunter's Mark`,
);
export const JudgementOfWisdom = withLabel(
	makeBooleanDebuffInput({
		actionId: () => ActionId.fromSpellId(20355),
		fieldName: 'judgementOfWisdom',
		showWhen: player => player.getFaction() === Faction.Alliance,
	}),
	'Judgement of Wisdom',
);
export const JudgementOfTheCrusader = withLabel(
	makeTristateDebuffInput({
		actionId: () => ActionId.fromSpellId(20303),
		impId: ActionId.fromSpellId(20337),
		fieldName: 'judgementOfTheCrusader',
		showWhen: player => player.getFaction() === Faction.Alliance,
	}),
	'Judgement of the Crusader',
);

// Misc Debuffs
export const JudgementOfLight = makeBooleanDebuffInput({
	actionId: () => ActionId.fromSpellId(20346),
	fieldName: 'judgementOfLight',
	showWhen: player => player.getFaction() === Faction.Alliance,
});
export const GiftOfArthas = makeBooleanDebuffInput({
	actionId: () => ActionId.fromSpellId(11374),
	fieldName: 'giftOfArthas',
});
export const CrystalYield = makeBooleanDebuffInput({
	actionId: () => ActionId.fromSpellId(15235),
	fieldName: 'crystalYield',
});

///////////////////////////////////////////////////////////////////////////
//                                 CONFIGS
///////////////////////////////////////////////////////////////////////////

export const RAID_BUFFS_CONFIG = [
	// Core Stat Buffs
	{
		config: AllStatsBuff,
		picker: IconPicker,
		stats: [],
	},
	{
		config: AllStatsPercentBuff,
		picker: MultiIconPicker,
		stats: [],
	},
	{
		config: StaminaBuff,
		picker: MultiIconPicker,
		stats: [Stat.StatStamina],
	},
	{
		config: BloodPactBuff,
		picker: MultiIconPicker,
		stats: [Stat.StatStamina],
	},
	{
		config: IntellectBuff,
		picker: MultiIconPicker,
		stats: [Stat.StatIntellect],
	},
	{
		config: SpiritBuff,
		picker: MultiIconPicker,
		stats: [Stat.StatSpirit],
	},

	// Tank-related Buffs
	{
		config: ArmorBuff,
		picker: IconPicker,
		stats: [Stat.StatArmor],
	},
	{
		config: PhysDamReductionBuff,
		picker: IconPicker,
		stats: [Stat.StatArmor],
	},
	// {
	// 	config: DamageReductionPercentBuff,
	// 	picker: IconPicker,
	// 	stats: [Stat.StatArmor],
	// },
	{
		config: ResistanceBuff,
		picker: MultiIconPicker,
		stats: [Stat.StatNatureResistance, Stat.StatShadowResistance, Stat.StatFireResistance, Stat.StatFrostResistance],
	},

	// Physical Damage Buffs
	{
		config: PaladinPhysicalBuff,
		picker: MultiIconPicker,
		stats: [Stat.StatAttackPower, Stat.StatStrength, Stat.StatAgility],
	},
	{
		config: StrengthBuffHorde,
		picker: IconPicker,
		stats: [Stat.StatStrength],
	},
	{
		config: BattleShoutBuff,
		picker: IconPicker,
		stats: [Stat.StatAttackPower],
	},
	{
		config: GraceOfAir,
		picker: IconPicker,
		stats: [Stat.StatAgility],
	},
	{
		config: TrueshotAuraBuff,
		picker: IconPicker,
		stats: [Stat.StatRangedAttackPower],
	},
	{
		config: MeleeCritBuff,
		picker: IconPicker,
		stats: [Stat.StatMeleeCrit],
	},
	// Threat Buffs

	// Spell Damage Buffs
	{
		config: SpellCritBuff,
		picker: IconPicker,
		stats: [Stat.StatSpellCrit],
	},
	{
		config: BlessingOfWisdom,
		picker: IconPicker,
		stats: [Stat.StatMP5],
	},
	{
		config: ManaSpringTotem,
		picker: IconPicker,
		stats: [Stat.StatMP5],
	},
] as PickerStatOptions[];

export const MISC_BUFFS_CONFIG = [
	{
		config: Thorns,
		picker: IconPicker,
		stats: [Stat.StatArmor],
	},
	{
		config: RetributionAura,
		picker: IconPicker,
		stats: [Stat.StatArmor],
	},
	{
		config: SanctityAura,
		picker: IconPicker,
		stats: [Stat.StatHolyPower],
	},
	{
		config: Innervate,
		picker: IconPicker,
		stats: [Stat.StatMP5],
	},
	{
		config: PowerInfusion,
		picker: IconPicker,
		stats: [Stat.StatMP5, Stat.StatSpellPower],
	},
	{
		config: BattleSquawkBuff,
		picker: IconPicker,
		stats: [Stat.StatMeleeHit],
	},
] as PickerStatOptions[];

export const WORLD_BUFFS_CONFIG = [
	// {
	// 	config: RallyingCryOfTheDragonslayer,
	// 	picker: IconPicker,
	// 	stats: [
	// 		Stat.StatMeleeCrit,
	// 		// TODO: Stat.StatRangedCrit,
	// 		Stat.StatSpellCrit,
	// 		Stat.StatAttackPower,
	// 	],
	// },
	{
		config: DragonslayerBuffInput,
		picker: MultiIconPicker,
		stats: [],
	},
	{
		config: SongflowerSerenade,
		picker: IconPicker,
		stats: [],
	},
	{
		config: SpiritOfZandalar,
		picker: IconPicker,
		stats: [],
	},
	{
		config: WarchiefsBlessing,
		picker: IconPicker,
		stats: [],
	},
	{
		config: FengusFerocity,
		picker: IconPicker,
		stats: [Stat.StatAttackPower],
	},
	{
		config: MoldarsMoxie,
		picker: IconPicker,
		stats: [Stat.StatStamina],
	},
	{
		config: SlipKiksSavvy,
		picker: IconPicker,
		stats: [Stat.StatSpellCrit],
	},
] as PickerStatOptions[];

export const SAYGES_CONFIG = [
	{
		config: SaygesDamage,
		stats: [],
	},
	{
		config: SaygesAgility,
		stats: [Stat.StatAgility],
	},
	{
		config: SaygesIntellect,
		stats: [Stat.StatIntellect],
	},
	{
		config: SaygesSpirit,
		stats: [Stat.StatSpirit, Stat.StatMP5],
	},
	{
		config: SaygesStamina,
		stats: [Stat.StatStamina],
	},
] as ItemStatOption<SaygesFortune>[];

export const DEBUFFS_CONFIG = [
	// Standard Debuffs
	{
		config: MajorArmorDebuff,
		stats: [Stat.StatAttackPower],
		picker: MultiIconPicker,
	},
	{
		config: CurseOfRecklessness,
		picker: IconPicker,
		stats: [Stat.StatAttackPower],
	},
	{
		config: FaerieFire,
		picker: IconPicker,
		stats: [Stat.StatAttackPower],
	},
	/* {
		config: BleedDebuff,
		picker: IconPicker,
		stats: [Stat.StatAttackPower, Stat.StatRangedAttackPower],
	}, */

	// Magic
	{
		config: JudgementOfTheCrusader,
		picker: IconPicker,
		stats: [Stat.StatHolyPower],
	},
	{
		config: SpellISBDebuff,
		picker: IconPicker,
		stats: [Stat.StatShadowPower],
	},
	{
		config: SpellScorchDebuff,
		picker: IconPicker,
		stats: [Stat.StatFirePower],
	},
	{
		config: SpellWintersChillDebuff,
		picker: IconPicker,
		stats: [Stat.StatFrostPower],
	},
	{
		config: NatureSpellDamageDebuff,
		picker: MultiIconPicker,
		stats: [Stat.StatNaturePower, Stat.StatArcanePower],
	},
	{
		config: SpellShadowWeavingDebuff,
		picker: IconPicker,
		stats: [Stat.StatShadowPower],
	},
	{
		config: WarlockCursesConfig,
		picker: MultiIconPicker,
		stats: [Stat.StatSpellPower],
	},

	// Defensive
	{
		config: AttackPowerDebuff,
		picker: MultiIconPicker,
		stats: [Stat.StatArmor],
	},
	{
		config: MeleeAttackSpeedDebuff,
		picker: MultiIconPicker,
		stats: [Stat.StatArmor],
	},
	{
		config: curseOfWeaknessDebuff,
		picker: IconPicker,
		stats: [Stat.StatArmor],
	},
	{
		config: MeleeHitDebuff,
		picker: IconPicker,
		stats: [Stat.StatDodge],
	},

	// Other Debuffs
	{
		config: HuntersMark,
		picker: IconPicker,
		stats: [Stat.StatRangedAttackPower],
	},
	{
		config: JudgementOfWisdom,
		picker: IconPicker,
		stats: [Stat.StatMP5, Stat.StatIntellect],
	},
] as PickerStatOptions[];

export const MISC_DEBUFFS_CONFIG = [
	// Misc Debuffs
	{
		config: GiftOfArthas,
		picker: IconPicker,
		stats: [Stat.StatAttackPower, Stat.StatRangedAttackPower],
	},
	{
		config: CrystalYield,
		picker: IconPicker,
		stats: [Stat.StatAttackPower, Stat.StatRangedAttackPower],
	},
	{
		config: JudgementOfLight,
		picker: IconPicker,
		stats: [Stat.StatStamina],
	},
] as PickerStatOptions[];

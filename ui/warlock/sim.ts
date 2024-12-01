import * as BuffDebuffInputs from '../core/components/inputs/buffs_debuffs';
import * as ConsumablesInputs from '../core/components/inputs/consumables.js';
import * as WarlockInputs from '../core/components/inputs/warlock_inputs';
import * as OtherInputs from '../core/components/other_inputs.js';
import { IndividualSimUI, registerSpecConfig } from '../core/individual_sim_ui.js';
import { Player } from '../core/player.js';
import { Class, Faction, ItemSlot, PartyBuffs, Race, Spec, Stat } from '../core/proto/common.js';
import { Stats } from '../core/proto_utils/stats.js';
import { getSpecIcon } from '../core/proto_utils/utils.js';
import * as Presets from './presets.js';

const SPEC_CONFIG = registerSpecConfig(Spec.SpecWarlock, {
	cssClass: 'warlock-sim-ui',
	cssScheme: 'warlock',
	// List any known bugs / issues here and they'll be shown on the site.
	knownIssues: [],

	// All stats for which EP should be calculated.
	epStats: [
		// Primary
		Stat.StatHealth,
		Stat.StatMana,
		// Attributes
		Stat.StatStrength,
		Stat.StatStamina,
		Stat.StatIntellect,
		Stat.StatSpirit,
		Stat.StatAgility,
		// Physical
		Stat.StatAttackPower,
		Stat.StatMeleeCrit,
		Stat.StatMeleeHit,
		Stat.StatMeleeHaste,
		// Spell
		Stat.StatSpellPower,
		Stat.StatSpellDamage,
		Stat.StatSpellHit,
		Stat.StatSpellCrit,
		Stat.StatSpellHaste,
		Stat.StatFirePower,
		Stat.StatShadowPower,
		Stat.StatMP5,
	],
	// Reference stat against which to calculate EP. DPS classes use either spell power or attack power.
	epReferenceStat: Stat.StatSpellPower,
	// Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
	displayStats: [
		// Primary
		Stat.StatMana,
		// Attributes
		Stat.StatStrength,
		Stat.StatAgility,
		Stat.StatIntellect,
		Stat.StatStamina,
		Stat.StatSpirit,
		// Physical
		Stat.StatAttackPower,
		Stat.StatMeleeCrit,
		Stat.StatMeleeHit,
		Stat.StatMeleeHaste,
		// Spell
		Stat.StatSpellPower,
		Stat.StatSpellDamage,
		Stat.StatFirePower,
		Stat.StatShadowPower,
		Stat.StatSpellHit,
		Stat.StatSpellCrit,
		Stat.StatSpellHaste,
		Stat.StatMP5,
	],
	displayPseudoStats: [],

	defaults: {
		// Default equipped gear.
		gear: Presets.DefaultGear.gear,

		// Default EP weights for sorting gear in the gear picker.
		epWeights: Stats.fromMap({
			[Stat.StatMana]: 0.01,
			[Stat.StatIntellect]: 0.23,
			[Stat.StatSpirit]: 0.0,
			[Stat.StatMP5]: 0.14,
			[Stat.StatSpellPower]: 1,
			[Stat.StatSpellDamage]: 1,
			[Stat.StatFirePower]: 0.1,
			[Stat.StatShadowPower]: 0.9,
			[Stat.StatSpellHit]: 12.79,
			[Stat.StatSpellCrit]: 7.92,
			[Stat.StatSpellHaste]: 7.83,
			[Stat.StatStamina]: 0.01,
			[Stat.StatFireResistance]: 0.0,
		}),
		// Default consumes settings.
		consumes: Presets.DefaultConsumes,
		// Default talents.
		talents: Presets.DefaultTalents.data,
		// Default spec-specific settings.
		specOptions: Presets.DefaultOptions,

		// Default buffs and debuffs settings.
		raidBuffs: Presets.DefaultRaidBuffs,

		partyBuffs: PartyBuffs.create({}),

		individualBuffs: Presets.DefaultIndividualBuffs,

		debuffs: Presets.DefaultDebuffs,

		other: Presets.OtherDefaults,
	},

	// IconInputs to include in the 'Player' section on the settings tab.
	playerIconInputs: [WarlockInputs.PetInput(), WarlockInputs.ImpFireboltRank(), WarlockInputs.ArmorInput(), WarlockInputs.WeaponImbueInput()],

	// Buff and Debuff inputs to include/exclude, overriding the EP-based defaults.
	includeBuffDebuffInputs: [
		// Physical buffs that affect pets
		BuffDebuffInputs.MajorArmorDebuff,
		BuffDebuffInputs.CurseOfRecklessness,
		BuffDebuffInputs.FaerieFire,
		BuffDebuffInputs.PaladinPhysicalBuff,
		BuffDebuffInputs.StrengthBuffHorde,
		BuffDebuffInputs.BattleShoutBuff,
		BuffDebuffInputs.GraceOfAir,
		BuffDebuffInputs.MeleeCritBuff,
		BuffDebuffInputs.BattleSquawkBuff,
		BuffDebuffInputs.GiftOfArthas,
		BuffDebuffInputs.CrystalYield,
	],
	excludeBuffDebuffInputs: [BuffDebuffInputs.SpellWintersChillDebuff, ...ConsumablesInputs.FROST_POWER_CONFIG],
	petConsumeInputs: [ConsumablesInputs.PetAttackPowerConsumable, ConsumablesInputs.PetAgilityConsumable, ConsumablesInputs.PetStrengthConsumable],
	// Inputs to include in the 'Other' section on the settings tab.
	otherInputs: {
		inputs: [WarlockInputs.PetPoolManaInput(), OtherInputs.DistanceFromTarget, OtherInputs.ChannelClipDelay],
	},
	itemSwapConfig: {
		itemSlots: [ItemSlot.ItemSlotMainHand, ItemSlot.ItemSlotOffHand, ItemSlot.ItemSlotRanged],
	},
	encounterPicker: {
		// Whether to include 'Execute Duration (%)' in the 'Encounter' section of the settings tab.
		showExecuteProportion: false,
	},

	presets: {
		// Preset talents that the user can quickly select.
		talents: Presets.TalentPresets,
		// Preset rotations that the user can quickly select.
		rotations: Presets.APLPresets,
		// Preset gear configurations that the user can quickly select.
		gear: Presets.GearPresets,
	},

	autoRotation: _player => {
		return Presets.DefaultAPL.rotation.rotation!;
	},

	raidSimPresets: [
		{
			spec: Spec.SpecWarlock,
			tooltip: 'Destruction DPS',
			defaultName: 'Destruction',
			iconUrl: getSpecIcon(Class.ClassWarlock, 2),

			talents: Presets.DefaultTalents.data,
			specOptions: Presets.DefaultOptions,
			consumes: Presets.DefaultConsumes,
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceHuman,
				[Faction.Horde]: Race.RaceOrc,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: Presets.DefaultGear.gear,
				},
				[Faction.Horde]: {
					1: Presets.DefaultGear.gear,
				},
			},
			otherDefaults: Presets.OtherDefaults,
		},
	],
});

export class WarlockSimUI extends IndividualSimUI<Spec.SpecWarlock> {
	constructor(parentElem: HTMLElement, player: Player<Spec.SpecWarlock>) {
		super(parentElem, player, SPEC_CONFIG);
	}
}

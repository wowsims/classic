import { Stat } from "../proto/common";

export enum Phase {
	Phase1 = 1,
	Phase2,
	Phase3,
	Phase4,
	Phase5,
}

export const CURRENT_PHASE = Phase.Phase1;

// Github pages serves our site under the /classic directory
export const REPO_NAME = 'classic';

// Get 'elemental_shaman', the pathname part after the repo name
const pathnameParts = window.location.pathname.split('/');
const repoPartIdx = pathnameParts.findIndex(part => part == REPO_NAME);
export const SPEC_DIRECTORY = repoPartIdx == -1 ? '' : pathnameParts[repoPartIdx + 1];

export const GLOBAL_DISPLAY_STATS = [
	Stat.StatHealth,
	Stat.StatFireResistance,
	Stat.StatFrostResistance,
	Stat.StatNatureResistance,
];

export const GLOBAL_DISPLAY_PSEUDO_STATS = [
	
];

export const GLOBAL_EP_STATS = [
	Stat.StatFireResistance,
	Stat.StatFrostResistance,
	Stat.StatNatureResistance,
];

export enum SortDirection {
	ASC,
	DESC,
}

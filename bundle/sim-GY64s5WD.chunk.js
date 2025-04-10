import{a2 as e,a1 as a,m as t,k as l,n as s,o as n,s as o,T as r,X as i,Y as c,Z as p,_ as d,a0 as S,$ as u,H as m,x as h}from"./preset_utils-BU2ZGt6X.chunk.js";import{L as f,H as P,$ as I,T as g,aT as y,a7 as v,a8 as w,aH as A,aa as k,ay as O,aP as B,aU as E,ab as T,ad as H,ac as R,at as C,aV as M,ag as x,av as D,ah as b,ar as F,as as L,aO as N,aj as V,a6 as j,ak as G,ai as U,am as W,P as J,an as z,az as _,al as Y,a as Z,S as $,ap as q,C as X,F as K}from"./detailed_results--FDh2HVQ.chunk.js";const Q=e({fieldName:"aura",values:[{value:f.NoPaladinAura,tooltip:"No Aura"},{actionId:()=>P.fromSpellId(20218),value:f.SanctityAura}]}),ee=e({fieldName:"personalBlessing",values:[{value:I.BlessingUnknown,tooltip:"No Blessing"},{actionId:()=>P.fromSpellId(20914),value:I.BlessingOfSanctuary}],changeEmitter:e=>g.onAny([e.specOptionsChangeEmitter])}),ae=a({fieldName:"righteousFury",actionId:e=>P.fromSpellId(25780),changeEmitter:e=>g.onAny([e.gearChangeEmitter,e.specOptionsChangeEmitter])}),te=e({fieldName:"primarySeal",values:[{actionId:()=>P.fromSpellId(20293),value:y.Righteousness},{actionId:()=>P.fromSpellId(20920),value:y.Command,showWhen:e=>e.getTalents().sealOfCommand},{actionId:()=>P.fromSpellId(407798),value:y.Martyrdom}],changeEmitter:e=>g.onAny([e.gearChangeEmitter,e.talentsChangeEmitter,e.specOptionsChangeEmitter])}),le={type:"TypeAPL",prepullActions:[{action:{castPaladinPrimarySeal:{}},doAtValue:{const:{val:"-3.0s"}}},{action:{castSpell:{spellId:{spellId:20928,rank:3}}},doAtValue:{const:{val:"-1.5s"}}},{action:{castSpell:{spellId:{itemId:18641}}},doAtValue:{const:{val:"-1s"}}}],priorityList:[{action:{condition:{cmp:{op:"OpLt",lhs:{currentHealthPercent:{}},rhs:{const:{val:"10%"}}}},castSpell:{spellId:{spellId:10310,rank:3}}}},{action:{condition:{cmp:{op:"OpLt",lhs:{currentHealthPercent:{}},rhs:{const:{val:"40%"}}}},castSpell:{spellId:{spellId:458371}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpGe",lhs:{spellTimeToReady:{spellId:{spellId:458371}}},rhs:{const:{val:"1m"}}}},{not:{val:{auraIsActive:{auraId:{spellId:25771}}}}}]}},castSpell:{spellId:{spellId:407788}}}},{action:{condition:{cmp:{op:"OpLe",lhs:{currentSealRemainingTime:{}},rhs:{const:{val:"1.5s"}}}},castPaladinPrimarySeal:{}}},{action:{castSpell:{spellId:{spellId:20928,rank:3}}}},{action:{castSpell:{spellId:{spellId:407669}}}},{action:{castSpell:{spellId:{spellId:407632}}}},{action:{castSpell:{spellId:{spellId:440658}}}},{action:{castSpell:{spellId:{spellId:415073}}}},{action:{strictSequence:{actions:[{castSpell:{spellId:{spellId:20271}}},{castPaladinPrimarySeal:{}}]}}}]},se={type:"TypeAPL",prepullActions:[{action:{castPaladinPrimarySeal:{}},doAtValue:{const:{val:"-3.0s"}}},{action:{castSpell:{spellId:{spellId:20928,rank:3}}},doAtValue:{const:{val:"-1.5s"}}},{action:{castSpell:{spellId:{itemId:18641}}},doAtValue:{const:{val:"-1s"}}}],priorityList:[{action:{condition:{cmp:{op:"OpLt",lhs:{currentHealthPercent:{}},rhs:{const:{val:"10%"}}}},castSpell:{spellId:{spellId:10310,rank:3}}}},{action:{condition:{cmp:{op:"OpLt",lhs:{currentHealthPercent:{}},rhs:{const:{val:"40%"}}}},castSpell:{spellId:{spellId:458371}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpGe",lhs:{spellTimeToReady:{spellId:{spellId:458371}}},rhs:{const:{val:"1m"}}}},{not:{val:{auraIsActive:{auraId:{spellId:25771}}}}}]}},castSpell:{spellId:{spellId:407788}}}},{action:{condition:{cmp:{op:"OpGt",lhs:{currentSealRemainingTime:{}},rhs:{const:{val:"0s"}}}},castSpell:{spellId:{spellId:20271}}}},{action:{condition:{cmp:{op:"OpLe",lhs:{currentSealRemainingTime:{}},rhs:{const:{val:"1.5s"}}}},castPaladinPrimarySeal:{}}},{action:{castSpell:{spellId:{spellId:20928,rank:3}}}},{action:{castSpell:{spellId:{spellId:407669}}}},{action:{castSpell:{spellId:{spellId:407632}}}},{action:{castSpell:{spellId:{spellId:415073}}}},{action:{castSpell:{spellId:{spellId:440658}}}}]},ne=t("Blank",{items:[]}),oe=l("P5 Prot",se),re=l("P4 Prot",le),ie={[v.Phase1]:[],[v.Phase2]:[],[v.Phase3]:[],[v.Phase4]:[re,oe],[v.Phase5]:[re,oe]},ce=ie[v.Phase5][0],pe={name:"P4 Prot",data:w.create({talentsString:"-053020335001551-0500535"})},de={name:"P5 Prot",data:w.create({talentsString:"-053020335001551-0520335"})},Se={[v.Phase1]:[],[v.Phase2]:[],[v.Phase3]:[],[v.Phase4]:[pe],[v.Phase5]:[de]},ue=Se[v.Phase5][0],me=A.create({aura:f.SanctityAura,primarySeal:y.Martyrdom,personalBlessing:I.BlessingOfSanctuary,righteousFury:!0}),he=k.create({agilityElixir:O.ElixirOfTheMongoose,healthElixir:B.ElixirOfFortitude,armorElixir:E.ElixirOfSuperiorDefense,defaultPotion:T.GreaterStoneshieldPotion,dragonBreathChili:!0,food:H.FoodTenderWolfSteak,flask:R.FlaskOfTheTitans,firePowerBuff:C.ElixirOfGreaterFirepower,fillerExplosive:M.ExplosiveDenseDynamite,spellPowerBuff:x.GreaterArcaneElixir,strengthBuff:D.JujuPower,zanzaBuff:b.ROIDS,attackPowerBuff:F.JujuMight,defaultConjured:L.ConjuredDemonicRune,alcohol:N.AlcoholRumseyRumBlackLabel}),fe=V.create({blessingOfWisdom:j.TristateEffectImproved,fengusFerocity:!0,moldarsMoxie:!0,rallyingCryOfTheDragonslayer:!0,saygesFortune:G.SaygesDamage,slipkiksSavvy:!0,songflowerSerenade:!0,spiritOfZandalar:!0}),Pe=U.create({powerWordFortitude:j.TristateEffectImproved,arcaneBrilliance:!0,battleShout:j.TristateEffectImproved,divineSpirit:!0,giftOfTheWild:j.TristateEffectImproved,sanctityAura:!0}),Ie=W.create({curseOfRecklessness:!0,faerieFire:!0,giftOfArthas:!0,exposeArmor:j.TristateEffectImproved,judgementOfWisdom:!0,judgementOfTheCrusader:j.TristateEffectImproved,improvedScorch:!0}),ge={distanceFromTarget:5,profession1:J.Blacksmithing,profession2:J.Engineering},ye=s($.SpecProtectionPaladin,{cssClass:"protection-paladin-sim-ui",cssScheme:"paladin",knownIssues:["Judgement of the Crusader is currently not implemented; users can manually award themselves the relevant spellpower amount\n\t\tfor a dps gain that will be slightly inflated given JotC does not benefit from source damage modifiers.","Be aware that not all item and weapon enchants are currently implemented in the sim, which make some notable Retribution\n\t\tweapons like Pendulum of Doom and The Jackhammer undervalued."],warnings:[e=>({updateOn:e.player.changeEmitter,getContent:()=>0==e.player.getSpecOptions().primarySeal?"Your previously selected seal is no longer available because of a talent or rune change.\n\t\t\t\t\t\t\tNo seal will be cast with this configuration. Please select an available seal in the Settings>Player menu.":""})],epStats:[z.StatHealth,z.StatMana,z.StatStrength,z.StatStamina,z.StatAgility,z.StatIntellect,z.StatAttackPower,z.StatMeleeHit,z.StatMeleeCrit,z.StatExpertise,z.StatSpellHit,z.StatSpellCrit,z.StatSpellPower,z.StatHolyPower,z.StatHealingPower,z.StatArmor,z.StatBonusArmor,z.StatDefense,z.StatDodge,z.StatParry,z.StatBlock,z.StatBlockValue,z.StatShadowResistance],epPseudoStats:[_.PseudoStatMainHandDps,_.PseudoStatMeleeSpeedMultiplier],epReferenceStat:z.StatAttackPower,displayStats:[z.StatMana,z.StatStrength,z.StatAgility,z.StatIntellect,z.StatAttackPower,z.StatMeleeHit,z.StatMeleeCrit,z.StatExpertise,z.StatSpellHit,z.StatSpellCrit,z.StatSpellPower,z.StatHolyPower,z.StatHealingPower,z.StatArmor,z.StatBonusArmor,z.StatDefense,z.StatDodge,z.StatParry,z.StatBlock,z.StatBlockValue,z.StatShadowResistance,z.StatArcaneResistance],displayPseudoStats:[_.PseudoStatMeleeSpeedMultiplier],defaults:{gear:ne.gear,epWeights:n.fromMap({[z.StatStrength]:3.23,[z.StatAgility]:18.57,[z.StatStamina]:0,[z.StatIntellect]:.05,[z.StatSpellPower]:.38,[z.StatHolyPower]:.29,[z.StatSpellHit]:8.2,[z.StatSpellCrit]:3.35,[z.StatAttackPower]:1,[z.StatMeleeHit]:0,[z.StatMeleeCrit]:39.75,[z.StatMana]:0,[z.StatArmor]:1,[z.StatDefense]:29.97,[z.StatBlock]:0,[z.StatBlockValue]:17.72,[z.StatDodge]:219.45,[z.StatParry]:217.72,[z.StatHealth]:0,[z.StatArcaneResistance]:0,[z.StatFireResistance]:0,[z.StatFrostResistance]:0,[z.StatNatureResistance]:0,[z.StatShadowResistance]:0,[z.StatBonusArmor]:.96,[z.StatHealingPower]:0},{[_.PseudoStatMainHandDps]:10.12,[_.PseudoStatMeleeSpeedMultiplier]:0}),consumes:he,talents:ue.data,specOptions:me,other:ge,raidBuffs:Pe,partyBuffs:Y.create({}),individualBuffs:fe,debuffs:Ie,race:Z.RaceHuman},playerIconInputs:[te,ae,ee,Q],includeBuffDebuffInputs:[o],excludeBuffDebuffInputs:[],otherInputs:{inputs:[r,i,c,p,d,S,u,m]},encounterPicker:{showExecuteProportion:!1},presets:{rotations:[...ie[v.Phase4]],talents:[...Se[v.Phase5],...Se[v.Phase4]],gear:[ne]},autoRotation:e=>ce.rotation.rotation,raidSimPresets:[{spec:$.SpecProtectionPaladin,tooltip:"Protection Paladin",defaultName:"Protection",iconUrl:q(X.ClassPaladin,1),talents:ue.data,specOptions:me,consumes:he,defaultFactionRaces:{[K.Unknown]:Z.RaceUnknown,[K.Alliance]:Z.RaceHuman,[K.Horde]:Z.RaceUnknown},defaultGear:{[K.Unknown]:{},[K.Alliance]:{1:ne.gear},[K.Horde]:{}}}]});class ve extends h{constructor(e,a){super(e,a,ye)}}export{ve as P};

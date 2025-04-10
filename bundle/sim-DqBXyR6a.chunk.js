import{z as e,m as a,k as t,l as s,n,o as i,t as l,G as o,y as r,T as d,H as c,x as p}from"./preset_utils-BU2ZGt6X.chunk.js";import{aw as h,a7 as S,a8 as u,ax as f,aa as P,ay as m,ar as I,ab as g,as as y,at as v,ac as w,ad as k,ae as M,af as A,ag as O,av as D,ah as T,ai as B,a6 as C,aj as E,am as H,P as R,a as x,an as b,az as F,al as N,ao as j,S as W,a1 as G,ap as z,C as L,F as U}from"./detailed_results--FDh2HVQ.chunk.js";const V=e({fieldName:"syncType",label:"Sync/Stagger Setting",labelTooltip:"Choose your sync or stagger option Perfect\n\t\t<ul>\n\t\t\t<li><div>Auto: Will auto pick sync options based on your weapons attack speeds</div></li>\n\t\t\t<li><div>None: No Sync or Staggering, used for mismatched weapon speeds</div></li>\n\t\t\t<li><div>Perfect Sync: Makes your weapons always attack at the same time, for match weapon speeds</div></li>\n\t\t\t<li><div>Delayed Offhand: Adds a slight delay to the offhand attacks while staying within the 0.5s flurry ICD window</div></li>\n\t\t</ul>",values:[{name:"Automatic",value:h.Auto},{name:"None",value:h.NoSync},{name:"Perfect Sync",value:h.SyncMainhandOffhandSwings},{name:"Delayed Offhand",value:h.DelayOffhandSwings}]}),J={type:"TypeAPL",prepullActions:[{action:{castSpell:{spellId:{spellId:10442,rank:4}}},doAtValue:{const:{val:"-6s"}}},{action:{castSpell:{spellId:{spellId:10614,rank:3}}},doAtValue:{const:{val:"-4.5s"}}},{action:{castSpell:{spellId:{spellId:10627,rank:2}}},doAtValue:{const:{val:"-3s"}}},{action:{castSpell:{spellId:{spellId:10438,rank:6}}},doAtValue:{const:{val:"-1.5s"}}}],priorityList:[{action:{condition:{not:{val:{auraIsActive:{auraId:{spellId:10442,rank:4}}}}},castSpell:{spellId:{spellId:10442,rank:4}}}},{action:{condition:{cmp:{op:"OpLe",lhs:{auraRemainingTime:{auraId:{spellId:10611}}},rhs:{const:{val:"1.5s"}}}},castSpell:{spellId:{spellId:10614,rank:3}}}},{action:{condition:{not:{val:{auraIsActive:{auraId:{spellId:10627,rank:2}}}}},castSpell:{spellId:{spellId:10627,rank:2}}}},{action:{autocastOtherCooldowns:{}}},{action:{castSpell:{spellId:{spellId:17364,rank:1}}}},{action:{condition:{and:{vals:[{not:{val:{dotIsActive:{spellId:{spellId:10438,rank:6}}}}},{cmp:{op:"OpGe",lhs:{remainingTime:{}},rhs:{const:{val:"20s"}}}}]}},castSpell:{spellId:{spellId:10438,rank:6}}}},{action:{condition:{cmp:{op:"OpGe",lhs:{currentManaPercent:{}},rhs:{const:{val:"50%"}}}},castSpell:{spellId:{spellId:10414,rank:7}}}}]},_={items:[{id:18817,enchant:1506},{id:18404},{id:12927},{id:13340,enchant:849},{id:11726,enchant:1891},{id:18812,enchant:1885},{id:15063,enchant:927},{id:11686},{id:15062,enchant:1506},{id:13210,enchant:911},{id:19325},{id:17063},{id:11815},{id:13965},{id:17182,enchant:2563},{},{}]},q={items:[{id:18817,enchant:2543},{id:19377},{id:16580},{id:19436,enchant:849},{id:11726,enchant:1891},{id:19587,enchant:1885},{id:19157,enchant:931},{id:19393},{id:22750,enchant:2543},{id:19381,enchant:911},{id:18821},{id:19384},{id:19406},{id:11815},{id:17182,enchant:1900},{},{}]},K={items:[{id:18817,enchant:2543},{id:21664},{id:21665},{id:19436,enchant:849},{id:21680,enchant:1891},{id:21602,enchant:1885},{id:19157,enchant:931},{id:21607},{id:21651,enchant:2543},{id:21493,enchant:911},{id:18821},{id:19384},{id:19406},{id:11815},{id:21134,enchant:1900},{},{id:22395}]},Q=a("Phase 1",{items:[{id:18817,enchant:1506},{id:18404},{id:12927},{id:13340,enchant:849},{id:11726,enchant:1891},{id:18812,enchant:1885},{id:15063,enchant:927},{id:11686},{id:15062,enchant:1506},{id:13210,enchant:911},{id:18821},{id:17063},{id:11815},{id:13965},{id:17182,enchant:2563},{},{}]}),X=a("Phase 2",_),Y=a("Phase 3",q),Z=a("Phase 5",K),$={[S.Phase1]:[Q],[S.Phase2]:[X],[S.Phase3]:[Y],[S.Phase4]:[],[S.Phase5]:[Z],[S.Phase6]:[]},ee=$[S.Phase1][0],ae=t("Default",J),te={[S.Phase1]:[ae],[S.Phase2]:[],[S.Phase3]:[],[S.Phase4]:[],[S.Phase5]:[],[S.Phase6]:[]},se=te[S.Phase1][0],ne=s("Level 60",u.create({talentsString:"05-5025002105023051-05105301"})),ie={[S.Phase1]:[ne],[S.Phase2]:[],[S.Phase3]:[],[S.Phase4]:[],[S.Phase5]:[],[S.Phase6]:[]},le=ie[S.Phase1][0],oe=f.create({syncType:h.Auto}),re=P.create({agilityElixir:m.ElixirOfTheMongoose,attackPowerBuff:I.JujuMight,defaultPotion:g.MajorManaPotion,defaultConjured:y.ConjuredDemonicRune,dragonBreathChili:!0,firePowerBuff:v.ElixirOfFirepower,flask:w.FlaskOfSupremePower,food:k.FoodBlessSunfruit,mainHandImbue:M.WindfuryWeapon,manaRegenElixir:A.MagebloodPotion,offHandImbue:M.WindfuryWeapon,spellPowerBuff:O.GreaterArcaneElixir,strengthBuff:D.JujuPower,zanzaBuff:T.ROIDS}),de=B.create({arcaneBrilliance:!0,battleShout:C.TristateEffectImproved,divineSpirit:!0,fireResistanceAura:!0,fireResistanceTotem:!0,giftOfTheWild:C.TristateEffectImproved,leaderOfThePack:!0,manaSpringTotem:C.TristateEffectRegular,powerWordFortitude:C.TristateEffectImproved}),ce=E.create({fengusFerocity:!0,moldarsMoxie:!0,rallyingCryOfTheDragonslayer:!0,slipkiksSavvy:!0,songflowerSerenade:!0,warchiefsBlessing:!0}),pe=H.create({curseOfRecklessness:!0,exposeArmor:C.TristateEffectImproved,faerieFire:!0,improvedScorch:!0,sunderArmor:!0}),he={profession1:R.Alchemy,profession2:R.Enchanting,race:x.RaceOrc},Se=n(W.SpecEnhancementShaman,{cssClass:"enhancement-shaman-sim-ui",cssScheme:"shaman",knownIssues:[],epStats:[b.StatIntellect,b.StatAgility,b.StatStrength,b.StatAttackPower,b.StatMeleeHit,b.StatMeleeCrit,b.StatExpertise,b.StatSpellPower,b.StatSpellDamage,b.StatFirePower,b.StatNaturePower,b.StatSpellCrit,b.StatSpellHit,b.StatMP5],epPseudoStats:[F.PseudoStatMainHandDps,F.PseudoStatOffHandDps,F.PseudoStatMeleeSpeedMultiplier],epReferenceStat:b.StatAttackPower,displayStats:[b.StatMana,b.StatStrength,b.StatAgility,b.StatIntellect,b.StatAttackPower,b.StatMeleeHit,b.StatMeleeCrit,b.StatExpertise,b.StatSpellDamage,b.StatSpellHit,b.StatSpellCrit,b.StatMP5],displayPseudoStats:[F.PseudoStatMeleeSpeedMultiplier],defaults:{race:he.race,gear:ee.gear,epWeights:i.fromMap({[b.StatIntellect]:.02,[b.StatAgility]:1.12,[b.StatStrength]:2.29,[b.StatSpellPower]:1.15,[b.StatSpellDamage]:1.15,[b.StatFirePower]:.63,[b.StatNaturePower]:.48,[b.StatSpellHit]:.03,[b.StatSpellCrit]:1.94,[b.StatMP5]:.01,[b.StatAttackPower]:1,[b.StatMeleeHit]:9.62,[b.StatMeleeCrit]:14.8,[b.StatFireResistance]:.5},{[F.PseudoStatMainHandDps]:8.15,[F.PseudoStatOffHandDps]:5.81,[F.PseudoStatMeleeSpeedMultiplier]:5.81}),consumes:re,talents:le.data,specOptions:oe,other:he,raidBuffs:de,partyBuffs:N.create({}),individualBuffs:ce,debuffs:pe},playerIconInputs:[],includeBuffDebuffInputs:[l,o,r],excludeBuffDebuffInputs:[],otherInputs:{inputs:[V,d,c]},itemSwapConfig:{itemSlots:[j.ItemSlotMainHand,j.ItemSlotOffHand]},customSections:[],encounterPicker:{showExecuteProportion:!1},presets:{talents:[...ie[S.Phase6],...ie[S.Phase5],...ie[S.Phase4],...ie[S.Phase3],...ie[S.Phase2],...ie[S.Phase1]],rotations:[...te[S.Phase6],...te[S.Phase5],...te[S.Phase4],...te[S.Phase3],...te[S.Phase2],...te[S.Phase1]],gear:[...$[S.Phase6],...$[S.Phase5],...$[S.Phase4],...$[S.Phase3],...$[S.Phase2],...$[S.Phase1]]},autoRotation:()=>se.rotation.rotation,raidSimPresets:[{spec:W.SpecBalanceDruid,tooltip:G[W.SpecBalanceDruid],defaultName:"Balance",iconUrl:z(L.ClassDruid,0),talents:le.data,specOptions:oe,consumes:re,otherDefaults:he,defaultFactionRaces:{[U.Unknown]:x.RaceUnknown,[U.Alliance]:x.RaceNightElf,[U.Horde]:x.RaceTauren},defaultGear:{[U.Unknown]:{},[U.Alliance]:{1:ee.gear},[U.Horde]:{1:ee.gear}}}]});class ue extends p{constructor(e,a){super(e,a,Se)}}export{ue as E};

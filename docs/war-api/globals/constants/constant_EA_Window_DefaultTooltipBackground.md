# EA_Window_DefaultTooltipBackground

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AuctionStats, Aura, BuddyBind, CCTV, CoolDownLine, Crusher, DAoCBuff, DammazKron |
| Files seen in | BuddyBind.xml, CCTV.xml, Code/CombatLog/CombatLogStatsWindow.xml, CoolDownLine.xml, Core/ToolTip/DK_Tooltip.xml, DascoreWin2.xml, Deathblow.xml, Deathblow2.xml |
| Namespaces detected | EA_Window_DefaultTooltipBackground |
| Source kinds | xml_attributes |
| Example locations | AbilityLinkTemplateWindowBackground, AuctionToolTipBackground, AuraShareTooltipBackground, AuraTooltipBackground, BuddyBindWindowBackground, CCTVSettingsBackground |
| XML usage count | 56 |
| XML attribute usage count | 56 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Engine-supplied XML constant or template class referenced by 36 addons.

## Seen In

- AuctionStats
- Aura
- BuddyBind
- CCTV
- CoolDownLine
- Crusher
- DAoCBuff
- DammazKron
- Dascore
- Deathblow
- Deathblow2
- DetauntHelper
- Duel
- EA_OpenPartyWindow
- Enemy
- GetStats
- Killer
- LootAlert
- MapMonster
- MapPin
- Miracle Grow Remix
- MoraleSet
- Motion
- NerfedButtons
- ReliquaryHunter
- RoR_debolster
- Shinies
- SocialWindow 2.0
- Statdoll
- Statdoll Light
- Statdoll Remix
- TomeTracker
- VPBreakdown
- WARRatingBuster
- Warbuilder
- zMailMod

## Used By

- AbilityLinkTemplateWindowBackground
- AuctionToolTipBackground
- AuraShareTooltipBackground
- AuraTooltipBackground
- BuddyBindWindowBackground
- CCTVSettingsBackground
- CCTV_BG_TEMPLATE
- CoolDownLineWndBackground
- CrusherConfigBase
- DAoCBuffCondenseTooltipBackground
- DascoreWin2WindowBackground
- EA_Tooltip_OpenPartyBackground
- EA_Tooltip_OpenPartyWorldBackground
- EA_Window_OpenPartyFlyOutBackground
- EnemyCombatLogStatsWindowTooltipBackground
- GetStatsCompareWindowBackground
- GetStatsWindowBackground
- KillerTooltipBackground
- LootAlertWindowBackground
- MapMonsterTooltipBackground
- MapPinCallTemplateWindowBackground
- MapPinTooltipsBackground
- MiracleGrow2RepeatTipBackground
- MonitorBodyBackground
- MoraleSetTooltipBackground
- MotionConfigBase
- Motion_Tooltip_TwoLineBackground
- NBSBCheckTooltipBackground
- NemesisTemplateWindowBackground
- ProfilTooltipBackground
- RatingToolTip1Background
- RatingToolTip2Background
- RatingToolTip3Background
- RatingToolTip4Background
- RatingToolTip5Background
- ReliquaryHunterMarkerToolTipWindowBackground
- RoR_debolsterWindowBackground
- SettingsWindowBackground
- ShiniesTooltipWindowBackground
- SocialWindowPlayerInformationTooltipWindowBackground
- StatdollMagicWndBackground
- StatdollMeleeWndBackground
- StatdollOptionsBackground
- StatdollRangeWndBackground
- StatdollWnd2Background
- StatdollWnd3Background
- StatdollWnd4Background
- StatdollWndBackground
- StatdollWndLightBackground
- TomeTracker_TooltipBackground
- VPBreakdownWindow.Background
- duelwindowBackground
- zMailModMassMailCraftingList
- zMailModMassMailCraftingListBackground
- zMailModMassMailRarityList
- zMailModMassMailRarityListBackground

## Related APIs

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

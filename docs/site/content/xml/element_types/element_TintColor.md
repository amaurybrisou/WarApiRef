# TintColor

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AggroMeter, Aura, AutoMark, BlackBox, BuffHead, CCTV, CMap |
| Files seen in | Core/ToolTip/DK_Tooltip.xml, Gui/HealGridGui.xml, Gui/HealGridGuiSpellList.xml, Gui/HealGridGuiTabMouseClick.xml, Gui/HealGridGuiTabSpellTrack.xml, Gui/HealGridGuiTemplates.xml, Gui/KwestorGuiTabArea.xml, Modules/UI/Shinies-UI-Auctions.xml |
| Namespaces detected | TintColor |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDFill, AdvancedPetAssist: APAInstantOnlyHUDFill, AdvancedPetAssist: APAKitingHUDFill, AdvancedPetAssist: APAPetTargetHUDBg, AggroMeter: AggroMeterWindow_AggroWindow1BorderCheck, AggroMeter: AggroMeterWindow_AggroWindow1Seperator1 |
| XML usage count | 573 |
| XML attribute usage count | 573 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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

TintColor is a XML UI element. It commonly appears under Button and CircleImage.

## Common Attributes

- a
- b
- g
- r

## Common Inherits

- none

## Common Parent Elements

- [DynamicImage](element_DynamicImage.md) — 374× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 118× (HIGH)
- [Label](element_Label.md) — 30× (HIGH)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 20× (HIGH)
- [CircleImage](element_CircleImage.md) — 15× (HIGH)
- [Button](element_Button.md) — 12× (HIGH)
- [VerticalResizeImage](element_VerticalResizeImage.md) — 4× (MEDIUM)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `b` | **required** | 100% | 0, 110, 130, 50, ... |
| `g` | **required** | 100% | 0, 200, 185, 130, ... |
| `r` | **required** | 100% | 255, 0, 180, 200, ... |
| `a` | optional | 18% | 30, 255, 0.8, 175, ... |
## Seen In

- AdvancedPetAssist
- AggroMeter
- Aura
- AutoMark
- BlackBox
- BuffHead
- CCTV
- CMap
- CleanCastbar
- CoolDownLine
- DammazKron
- Deathblow
- Deathblow2
- DetauntHelper
- EA_LoadingScreen
- EA_OpenPartyWindow
- EA_UiDebugTools
- EZCraftX
- EveryBodyGuard
- FlagCap
- Group Icons SG
- GroupRange
- GuardLine
- GuildWarden
- HealGrid
- Hopper
- InfoScroller
- KeyBar
- KillStreak
- Killer
- Kwestor
- LibAddonButton
- LibGroup
- Map
- MapMonster
- MapPin
- MarkBuff
- MoraleSet
- Moth
- Motion
- NaturalLog
- NerfedButtons
- Obsidian
- PartyCast
- PieTracker
- Pocket Palette
- Pure
- Pure Careerbar
- RVAPI_ColorDialog
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RaidMeter
- ResHelp
- RoR_SoR
- SessionRPs
- Shinies
- Soloq
- SpamBayes
- Squared
- Statdoll
- Statdoll Light
- Statdoll Remix
- Swinger
- Targets
- TaxPayer
- TexturedButtons
- Tortall_DPS
- TurretRange
- Twister
- Warbuilder
- Wikki's Cooldown Bar
- ZonePOP
- minesweep
- zMailMod

## Examples

- AdvancedPetAssist: APAFollowTargetHUDFill -> TintColor in HorizontalResizeImage APAFollowTargetHUDFill
- AdvancedPetAssist: APAInstantOnlyHUDFill -> TintColor in HorizontalResizeImage APAInstantOnlyHUDFill
- AdvancedPetAssist: APAKitingHUDFill -> TintColor in HorizontalResizeImage APAKitingHUDFill
- AdvancedPetAssist: APAPetTargetHUDBg -> TintColor in HorizontalResizeImage APAPetTargetHUDBg
- AggroMeter: AggroMeterWindow_AggroWindow1BorderCheck -> TintColor in FullResizeImage AggroMeterWindow_AggroWindow1BorderCheck
- AggroMeter: AggroMeterWindow_AggroWindow1Seperator1 -> TintColor in FullResizeImage AggroMeterWindow_AggroWindow1Seperator1

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [CircleImage](element_CircleImage.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [HorizontalResizeImage](element_HorizontalResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [VerticalResizeImage](element_VerticalResizeImage.md) (HIGH 100/100) - XML Element Type

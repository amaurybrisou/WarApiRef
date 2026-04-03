# InterfaceCore

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 125

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, ActionFraction, AdjustTheTip, Amethyst, BuffHead, CMap, CleanUnitFrames, Crusher |
| Files seen in | Code/Assist/Assist.lua, Code/CombatLog/CombatLog.lua, Code/Core/Main.lua, Code/GroupIcons/GroupIcons.lua, Code/Guard/Guard.lua, Code/TalismanAlerter/TalismanAlerter.lua, Code/Timer/Timer.lua, Code/UnitFrames/EffectsIndicator.lua |
| Namespaces detected | InterfaceCore |
| Source kinds | lua_calls |
| Example locations | Ace: Scale, ActionFraction: ResetWindow, AdjustTheTip: UpdateCallback, Amethyst: Recreate, Amethyst: Scale, BuffHead: AnchorContainers |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 165 |
| Global usage count | 3 |
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

Shared function table with 3 member functions; the primary API surface for 69 addons.

## Functions

- InterfaceCore.GetResolutionScale
- InterfaceCore.GetScale
- InterfaceCore.ReloadUI

## Observed Members

- none

## Seen In

- Ace
- ActionFraction
- AdjustTheTip
- Amethyst
- BuffHead
- CMap
- CleanUnitFrames
- Crusher
- DAoCBuff
- EA_LoadingScreen
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- Enemy
- GCDsaver
- Group Icons
- GroupRange
- GuardLine
- Hopper
- InfoScroller
- KeyBar
- LibAddonButton
- LibWBToggler
- Map
- MapMonster
- MapPin
- Mass Refine
- Miracle Grow Remix
- Motion
- Obsidian
- PartyCast
- PotionBar
- Preciousss
- Pure
- Pure Careerbar
- RO-Style Combat Text
- RVAPI_ColorDialog
- RVMOD_Manager
- RVMOD_SquaredDistances
- RealmStatus
- RoR_SoR
- SNT_PANEL
- SessionRPs
- Shinies
- Squared
- SquaredHDIndicator
- SquaredHotIndicators
- TargetRing
- TastyButtons
- TexturedButtons
- TidyQueue
- Tokens
- TurretRange
- Vectors
- Vertigo
- WBStutterLess
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- nLootLink
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Scale -> InterfaceCore.GetScale()
- ActionFraction: ResetWindow -> InterfaceCore.GetScale()
- AdjustTheTip: UpdateCallback -> InterfaceCore.GetScale()
- Amethyst: Recreate -> InterfaceCore.GetScale()
- Amethyst: Scale -> InterfaceCore.GetScale()
- BuffHead: AnchorContainers -> InterfaceCore.GetScale()

## Notes

- none

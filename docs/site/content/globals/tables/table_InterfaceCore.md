# InterfaceCore

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 130

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, BuffHead, DAoCBuff, EA_UiDebugTools, Effigy, Enemy, GCDsaver, GuardLine |
| Files seen in | `/workspace/Ace/LibGUI.lua:236`, `/workspace/BuffHead/AdvancedContainers.lua:275`, `/workspace/BuffHead/Container.lua:758`, `/workspace/BuffHead/EffectFrame.lua:52`, `/workspace/BuffHead/Setup/LayoutControlFrame.lua:8`, `/workspace/BuffHead/Setup/LayoutFrame.lua:105`, `/workspace/BuffHead/Setup/LayoutFrame.lua:3`, `/workspace/BuffHead/Setup/LayoutFrame.lua:71` |
| Namespaces detected | InterfaceCore |
| Source kinds | globals, lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Scale, BuffHead: AutoSize, BuffHead: BuffHead.AdvancedContainers.OnLayoutEditorFinished, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.Show, BuffHead: BuffHead.Setup.Layout.Properties.Show, BuffHead: BuffHead.Setup.LayoutFrame:Create |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 68 |
| Global usage count | 3 |
| Local definition count | 9 |
| Documentation references | 1 |
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

Observed shared global table or namespace surfaced in 23 addons.

## Functions

- InterfaceCore.GetResolutionScale
- InterfaceCore.GetScale
- InterfaceCore.ReloadUI

## Observed Members

- none

## Seen In

- Ace
- BuffHead
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- GuardLine
- LibWBToggler
- MapMonster
- MapPin
- Miracle Grow Remix
- PotionBar
- RVAPI_ColorDialog
- RVMOD_Manager
- RoR_SoR
- Shinies
- TexturedButtons
- TidyQueue
- TurretRange
- WarTriage
- WoH-Reticle
- wbLeadHelper

## Examples

- Ace: LIBGUI_ELEMENT:Scale -> InterfaceCore.GetScale()
- BuffHead: AutoSize -> InterfaceCore.GetScale()
- BuffHead: BuffHead.AdvancedContainers.OnLayoutEditorFinished -> InterfaceCore.GetScale()
- BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.Show -> InterfaceCore.GetScale()
- BuffHead: BuffHead.Setup.Layout.Properties.Show -> InterfaceCore.GetScale()
- BuffHead: BuffHead.Setup.LayoutFrame:Create -> InterfaceCore.GetScale()

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none

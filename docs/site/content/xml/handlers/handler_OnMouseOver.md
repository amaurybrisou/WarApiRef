# OnMouseOver

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 88/100

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, referenced by generated docs or reference files, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Atlas, AuctionStats, Aura, AutoBand, BagOMatic |
| Namespaces detected | OnMouseOver |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedRenownTrainer: .OnMouseOver, AggroMeter: .OnMouseOver, AnywhereTrainer: .OnMouseOver, Atlas: .OnMouseOver, AuctionStats: .OnMouseOver, Aura: .OnMouseOver |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 652 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
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

XML handler event observed across 88 addons.

## Expected Lua Binding

```lua
function()
```

## Element Types

- none

## Seen In

- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Atlas
- AuctionStats
- Aura
- AutoBand
- BagOMatic
- BankArkel
- BlackBook
- Brizio's Crappy Computer Medic
- BuffHead
- CM_ClosetGoblin
- CaVES
- Calling
- CastSequence
- CleanUnitFrames
- Crusher
- DAoCBuff
- DammazKron
- DetauntHelper
- DuffTimer
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EZCraftX
- Effigy
- Emojii
- Enemy
- GroupSpotter
- GuildWarden
- HealGrid
- JunkDump
- KeyBar
- Killer
- Kwestor
- LibAddonButton
- LoyalPet
- MapMonster
- MapPin
- MarkBuff
- MiniMapMonster
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- Motion
- NaturalLog
- NerfedButtons
- NoOverheal
- Obsidian
- Pocket Palette
- PotionBar
- Pure
- QuickTacticSwitch
- QuickWarReport
- RVMOD_Manager
- RVMOD_SquaredDistances
- RandomMount
- Refer
- ReliquaryHunter
- RoR_SoR
- RvRContribution
- SOR
- Sequencer
- SessionRPs
- Shinies
- Squared
- TalismanGenie
- TastyButtons
- TexturedButtons
- Tome Titan
- TomeTracker
- TurretRange
- TwisterSet
- WARCommander
- WSCT
- WarBoard
- WarBoard_Loc
- WarBoard_Session
- WarBoard_TaliMon
- XpStatus+G
- ZonePOP
- emotes
- followTheLeader
- nLootLink
- talisman-monitor
- wbLeadHelper

## Examples

- AdvancedRenownTrainer: .OnMouseOver -> AdvancedRenownTraining.AbilityTooltip
- AggroMeter: .OnMouseOver -> AggroMeter.SelectChar
- AggroMeter: .OnMouseOver -> AggroMeter.OnMouseOverStart
- AnywhereTrainer: .OnMouseOver -> AnywhereTrainer.OnMouseOver
- Atlas: .OnMouseOver -> AtlasMap.OnMouseOver
- Atlas: .OnMouseOver -> AtlasMap.OnLegendItemMouseOver

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](../element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [DynamicImage](../element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](../element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [ListBox](../element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [SliderBar](../element_types/element_SliderBar.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [EA_Window_ContextMenu.OnMouseOverDefaultMenuItem](../../globals/functions/global_EA_Window_ContextMenu.OnMouseOverDefaultMenuItem.md) (HIGH 80/100) - Global Function

## Used With

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Affects

- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).

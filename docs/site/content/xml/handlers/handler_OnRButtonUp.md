# OnRButtonUp

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
| Addons seen in | AggroMeter, AnywhereTrainer, AnywhereTrainerAdditions, Atlas, AuctionStats, Aura, AutoBand, Brizio's Crappy Computer Medic |
| Namespaces detected | OnRButtonUp |
| Source kinds | bindings, xml_handlers |
| Example locations | AggroMeter: .OnRButtonUp, AnywhereTrainer: .OnRButtonUp, AnywhereTrainerAdditions: .OnRButtonUp, Atlas: .OnRButtonUp, AuctionStats: .OnRButtonUp, Aura: .OnRButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 218 |
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

XML handler event observed across 80 addons.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- none

## Seen In

- AggroMeter
- AnywhereTrainer
- AnywhereTrainerAdditions
- Atlas
- AuctionStats
- Aura
- AutoBand
- Brizio's Crappy Computer Medic
- BuffHead
- Busted
- CM_ClosetGoblin
- Calling
- CastSequence
- CleanUnitFrames
- DAoCBuff
- DammazKron
- DetauntHelper
- DuffTimer
- EA_UiDebugTools
- Effigy
- Emojii
- Enemy
- GroupSpotter
- HealGrid
- JunkDump
- KeyBar
- Killer
- Kwestor
- LibAddonButton
- Map
- MapMonster
- MapPin
- MarkBuff
- MiniMapMonster
- Minmap
- Miracle Grow Remix
- MiracleGrowLight
- MoraleCircle
- NerfedButtons
- NoOverheal
- Obsidian
- PeaceOut
- PlayEffectsOn
- Pocket Palette
- PotionBar
- Pure
- RandomMount
- Refer
- ReliquaryHunter
- ResHelp
- RoR_SoR
- Rotation
- RvRContribution
- SOR
- ScenarioStats
- Shinies
- Squared
- SquaredClick
- Statdoll
- Statdoll Light
- Statdoll Remix
- TastyButtons
- TexturedButtons
- ThankTheTank
- ThinkOutLoud
- TidyQueue
- Tome Titan
- Trakario
- TurretRange
- Twister
- TwisterSet
- WARCommander
- WarBoard
- WarBoard_Session
- ZonePOP
- emotes
- followTheLeader
- minesweep
- nLootLink
- wbLeadHelper

## Examples

- AggroMeter: .OnRButtonUp -> AggroMeter.OnTabRBU
- AggroMeter: .OnRButtonUp -> AggroMeter.PickedListMenu
- AnywhereTrainer: .OnRButtonUp -> AnywhereTrainer.OnRButtonUp
- AnywhereTrainerAdditions: .OnRButtonUp -> AnywhereTrainerAdditions.OnRButtonUp
- Atlas: .OnRButtonUp -> AtlasMap.OnMapRBU
- AuctionStats: .OnRButtonUp -> AuctionStats.OnRButtonUpItem

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](../element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](../element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EA_ChatWindow.Print](../../globals/functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateDefaultContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateDefaultContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](../element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetMovable](../../window_api/functions/window_WindowGetMovable.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Affects

- [GameData.AbilityType.PET](../../gamedata/fields/gamedata_GameData.AbilityType.PET.md) (HIGH 100/100) - GameData Field
- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [GameData.Player.rvrPermaFlagged](../../gamedata/fields/gamedata_GameData.Player.rvrPermaFlagged.md) (HIGH 100/100) - GameData Field
- [GameData.Player.rvrZoneFlagged](../../gamedata/fields/gamedata_GameData.Player.rvrZoneFlagged.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.BUTTON_CLICK](../../gamedata/fields/gamedata_GameData.Sound.BUTTON_CLICK.md) (HIGH 100/100) - GameData Field
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected binding arguments remain uncertain because contract artifacts capture symbol linkage, not full handler signatures.

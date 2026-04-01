# WindowSetParent

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 16 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, BagOMatic, CM_ClosetGoblin, Effigy, GCDsaver, GetStats, JunkDump, LibWBToggler |
| Files seen in | `/workspace/Ace/LibGUI.lua:132`, `/workspace/ClosetGoblin/ClosetGoblin.lua:83`, `/workspace/Effigy/Elements/EffigyBar.lua:135`, `/workspace/Effigy/LibGUI.lua:132`, `/workspace/GCDsaver/libs/LibGUI.lua:132`, `/workspace/GetStats/GetStats.lua:119`, `/workspace/JunkDump/JunkDump.lua:50`, `/workspace/LibWarBoardToggler/libs/LibGUI.lua:132` |
| Namespaces detected | WindowSetParent |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Parent, BagOMatic: BagOMatic.init, CM_ClosetGoblin: ClosetGoblin.Initialize, Effigy: EffigyBar:setup, Effigy: LIBGUI_ELEMENT:Parent, GCDsaver: LIBGUI_ELEMENT:Parent |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 24 |
| Global usage count | 24 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
WindowSetParent(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "BagOMaticButton", "GetStatsCompareLine"..i, "GetStatsLine"..GetStats.LineNumber |
| arg2 | Observed as a text or wstring payload. | Observed values: "CharacterWindow", "EA_Window_Backpack", "EA_Window_InteractionStore" |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- BagOMatic
- CM_ClosetGoblin
- Effigy
- GCDsaver
- GetStats
- JunkDump
- LibWBToggler
- MapMonster
- RVMOD_Manager
- Shinies
- TidyChat
- TurretRange
- WarTriage
- WoH-Reticle
- nRarity

## Examples

- Ace: LIBGUI_ELEMENT:Parent -> WindowSetParent(self.name, self.parent)
- BagOMatic: BagOMatic.init -> WindowSetParent("BagOMaticButton", "EA_Window_Backpack")
- CM_ClosetGoblin: ClosetGoblin.Initialize -> WindowSetParent("yClosetGoblinButton", "CharacterWindow")
- Effigy: EffigyBar:setup -> WindowSetParent(self.name, invisiName)
- Effigy: LIBGUI_ELEMENT:Parent -> WindowSetParent(self.name, self.parent)
- GCDsaver: LIBGUI_ELEMENT:Parent -> WindowSetParent(self.name, self.parent)

## Related APIs

- none

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [EA_Window_Backpack](../../globals/tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Settings.Language.active](../../systemdata/fields/systemdata_SystemData.Settings.Language.active.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event

## Affects

- [EA_Window_Backpack](../../globals/tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

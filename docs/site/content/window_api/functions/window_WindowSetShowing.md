# WindowSetShowing

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Moth, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:167`, `/workspace/data/raw/Moth/Moth.lua:270`, `/workspace/data/raw/Moth/Moth.lua:421`, `/workspace/data/raw/Moth/Moth.lua:463`, `/workspace/data/raw/Moth/Moth.lua:578`, `/workspace/data/raw/Moth/Moth.lua:591`, `/workspace/data/raw/Moth/Moth.lua:619`, `/workspace/data/raw/Moth/Moth.lua:636` |
| Namespaces detected | WindowSetShowing |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.Clear, Moth: Moth.HealthBar, Moth: Moth.HideBorders, Moth: Moth.HideTemplateElements, Moth: Moth.ShowDefaultMouseOverTargetWindows, Moth: Moth.UpdateCareerIcon |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 50 |
| Global usage count | 50 |
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
WindowSetShowing(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "DefaultTooltip", "EA_Window_LootRoll", "Moth" |
| arg2 | Observed as a boolean toggle. | Observed values: TidyChatLootRoll.itemRollDataDisplayOrder[1]~=nil, bool, chatwindow_background_showing |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Moth
- TidyChat
- TidyRoll

## Examples

- Moth: Moth.Clear -> WindowSetShowing("Moth", false)
- Moth: Moth.Clear -> WindowSetShowing("MothHealthBar", false)
- Moth: Moth.HealthBar -> WindowSetShowing("MothHealthBar", false)
- Moth: Moth.HideBorders -> WindowSetShowing("MothBordertronned", false)
- Moth: Moth.HideBorders -> WindowSetShowing("MothBorderDark", false)
- Moth: Moth.HideBorders -> WindowSetShowing("MothBorderGold", false)

## Related APIs

- none

## Used With

- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 98/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 98/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 90/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 90/100) - Window Function
- [WindowUnregisterCoreEventHandler](window_WindowUnregisterCoreEventHandler.md) (HIGH 90/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 80/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 80/100) - Window Function
- [WindowGetAnchor](window_WindowGetAnchor.md) (HIGH 80/100) - Window Function
- [WindowGetOffsetFromParent](window_WindowGetOffsetFromParent.md) (HIGH 80/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 71/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

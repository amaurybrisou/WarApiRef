# WindowSetDimensions

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 98/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Moth, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:215`, `/workspace/data/raw/Moth/Moth.lua:227`, `/workspace/data/raw/Moth/Moth.lua:421`, `/workspace/data/raw/Moth/Moth.lua:591`, `/workspace/data/raw/TidyChat/TidyChat.lua:980`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136` |
| Namespaces detected | WindowSetDimensions |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.SetCellText, Moth: Moth.SetCellTextIcon, Moth: Moth.UpdateHealthBar, Moth: Moth.UpdateNPCIcon, TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons, TidyRoll: TidyRollOptions.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 14 |
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
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
WindowSetDimensions(windowName, arg2, arg3)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "MothHealthBar", c_CHANNEL_MENU, c_TROLL_BNUM_TBOX |
| arg2 | Observed as a numeric value. | Observed values: 140, 32, 50 |
| arg3 | Observed as a numeric value. | Observed values: 20, 30, 32 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Moth
- TidyChat
- TidyRoll

## Examples

- Moth: Moth.SetCellText -> WindowSetDimensions(cellName, LabelGetTextDimensions(cellText))
- Moth: Moth.SetCellTextIcon -> WindowSetDimensions(cellName, 32, 32)
- Moth: Moth.UpdateHealthBar -> WindowSetDimensions("MothHealthBar", healthWidth, Moth.Global.healthBar)
- Moth: Moth.UpdateNPCIcon -> WindowSetDimensions(cellName, WindowGetDimensions(cellName.."NPCIcon"))
- TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons -> WindowSetDimensions(c_CHANNEL_MENU, x, y+28*3)
- TidyRoll: TidyRollOptions.Initialize -> WindowSetDimensions(c_TROLL_VERSION, 80, 30)

## Related APIs

- none

## Used With

- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 98/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 90/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 90/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 80/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 80/100) - Window Function
- [LabelGetTextDimensions](window_LabelGetTextDimensions.md) (HIGH 80/100) - Window Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 71/100) - Global Function

## Triggered By

- none

## Affects

- [SystemData.ChatLogFilters.ADVICE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ADVICE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.ALLIANCE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ALLIANCE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.SCENARIO](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SCENARIO.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

# WindowGetDimensions

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

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
| Addons seen in | Moth, TidyChat |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:421`, `/workspace/data/raw/Moth/Moth.lua:591`, `/workspace/data/raw/TidyChat/TidyChat.lua:980` |
| Namespaces detected | WindowGetDimensions |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.UpdateHealthBar, Moth: Moth.UpdateNPCIcon, TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
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
WindowGetDimensions(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "MothBackground", c_CHANNEL_MENU, cellName.."NPCIcon" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Moth
- TidyChat

## Examples

- Moth: Moth.UpdateHealthBar -> WindowGetDimensions("MothBackground")
- Moth: Moth.UpdateNPCIcon -> WindowGetDimensions(cellName.."NPCIcon")
- TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons -> WindowGetDimensions(c_CHANNEL_MENU)

## Related APIs

- none

## Used With

- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 98/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 98/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 90/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 80/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 80/100) - Window Function

## Triggered By

- none

## Affects

- [SystemData.ChatLogFilters.ADVICE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ADVICE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.ALLIANCE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ALLIANCE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.SCENARIO](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SCENARIO.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

# WindowSetId

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | InfoScroller, PartyCast, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:262`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:262`, `/workspace/data/raw/TidyChat/TidyChat.lua:980`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136` |
| Namespaces detected | WindowSetId |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_ELEMENT:SetId, PartyCast: LIBGUI_ELEMENT:SetId, TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons, TidyRoll: TidyRollOptions.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
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
WindowSetId(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: c_CHANNEL_MENU.."AdviceButton", c_CHANNEL_MENU.."AllianceButton", c_CHANNEL_MENU.."ScenarioButton" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 1, 2, ChatSettings.Channels[SystemData.ChatLogFilters.ADVICE].id |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- PartyCast
- TidyChat
- TidyRoll

## Examples

- InfoScroller: LIBGUI_ELEMENT:SetId -> WindowSetId(self.name, id)
- PartyCast: LIBGUI_ELEMENT:SetId -> WindowSetId(self.name, id)
- TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons -> WindowSetId(c_CHANNEL_MENU.."AllianceButton", ChatSettings.Channels[SystemData.ChatLogFilters.ALLIANCE].id)
- TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons -> WindowSetId(c_CHANNEL_MENU.."AdviceButton", ChatSettings.Channels[SystemData.ChatLogFilters.ADVICE].id)
- TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons -> WindowSetId(c_CHANNEL_MENU.."ScenarioButton", ChatSettings.Channels[SystemData.ChatLogFilters.SCENARIO].id)
- TidyRoll: TidyRollOptions.Initialize -> WindowSetId(tabButton, c_TROLL_GENERAL_TAB)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [SystemData.ChatLogFilters.ADVICE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ADVICE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.ALLIANCE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ALLIANCE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.SCENARIO](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SCENARIO.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

# LayoutEditor.RegisterWindow

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

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
| Addons seen in | InfoScroller, Moth, Soloq, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:30`, `/workspace/data/raw/Moth/Moth.lua:713`, `/workspace/data/raw/Soloq/Soloq.lua:22`, `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:227` |
| Namespaces detected | LayoutEditor |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.OnInitialize, Moth: Moth.Initialize, Soloq: Soloq.OnInitialize, TidyChat: TidyChatFrames.Initialize, TidyRoll: TidyRoll.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
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
LayoutEditor.RegisterWindow(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
```

## Description

Observed as a window function across 5 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "InfoScrollerMainWindow", "Moth", "SoloqOverviewWindow" |
| arg2 | Observed as a text or wstring payload. | Observed values: L "Chat Text Entry", L "Moth", L "Scroller" |
| arg3 | Observed as a text or wstring payload. | Observed values: L "Mouse Over Target Hover", L "Scroller", L "Soloq" |
| arg4 | Observed as a boolean toggle. | Observed values: false, true |
| arg5 | Observed as a boolean toggle. | Observed values: false, true |
| arg6 | Observed as a boolean toggle. | Observed values: false, true |
| arg7 | Observed as a runtime window or control identifier. | Observed values: nil |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- InfoScroller
- Moth
- Soloq
- TidyChat
- TidyRoll

## Examples

- InfoScroller: InfoScroller.OnInitialize -> LayoutEditor.RegisterWindow("InfoScrollerMainWindow", L "Scroller", L "Scroller", false, true, true, nil)
- Moth: Moth.Initialize -> LayoutEditor.RegisterWindow("Moth", L "Moth", L "Mouse Over Target Hover", false, false, false, nil)
- Soloq: Soloq.OnInitialize -> LayoutEditor.RegisterWindow("SoloqOverviewWindow", L "Soloq", L "Soloq", false, false, true, nil)
- TidyChat: TidyChatFrames.Initialize -> LayoutEditor.RegisterWindow(c_TEXT_ENTRY_ANCHOR, L "Chat Text Entry", L "Where Text Entry are displayed.", true, false, false, nil)
- TidyRoll: TidyRoll.Initialize -> LayoutEditor.RegisterWindow(c_TIDY_ROLL_ANCHOR, L "Tidy Roll", L "Where Tidy Roll are displayed.", false, false, true, nil)

## Related APIs

- none

## Used With

- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_POST_MODE](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_POST_MODE.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

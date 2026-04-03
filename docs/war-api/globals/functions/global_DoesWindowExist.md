# DoesWindowExist

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 83/100
- Seen in: 5 addons

## Confidence Assessment

- Level: HIGH

- Score: 83/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, Soloq, TidyChat, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:95`, `/workspace/data/raw/Soloq/ui/Overview.lua:24`, `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:265`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:853`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:865`, `/workspace/data/raw/minesweep/minesweep.lua:11` |
| Namespaces detected | DoesWindowExist |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.CreateCard, Soloq: Soloq.createOverwiewWindow, TidyChat: TidyChatFrames.Initialize, TidyRoll: TidyRoll.OnLoad, TidyRoll: TidyRollOptions.RadioGetId, TidyRoll: TidyRollOptions.RadioSetId |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
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
DoesWindowExist(windowName)
```

## Description

Observed guarding runtime window creation and cleanup by checking whether a named window exists.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "EA_Window_LootRoll", "InfoScroller_"..i, "MineSweepWindow" |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- InfoScroller
- Soloq
- TidyChat
- TidyRoll
- minesweep

## Examples

- InfoScroller: InfoScroller.CreateCard -> DoesWindowExist("InfoScroller_"..i)
- Soloq: Soloq.createOverwiewWindow -> DoesWindowExist(overviewWindowName)
- TidyChat: TidyChatFrames.Initialize -> DoesWindowExist(c_TEXT_ENTRY_WINDOW.."EntryBoxLanguageButton")
- TidyRoll: TidyRoll.OnLoad -> DoesWindowExist("EA_Window_LootRoll")
- TidyRoll: TidyRollOptions.RadioGetId -> DoesWindowExist(radioGroupName..index)
- TidyRoll: TidyRollOptions.RadioSetId -> DoesWindowExist(radioGroupName..index)

## Related APIs

- none

## Used With

- [DynamicImageSetTextureOrientation](../../window_api/functions/window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](../../window_api/functions/window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 98/100) - Window Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Triggered By

- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event

## Affects

- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.

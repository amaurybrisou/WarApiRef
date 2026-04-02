# OnMButtonUp

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 176

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TidyRoll/TidyRoll.xml:202` |
| Namespaces detected | OnMButtonUp |
| Source kinds | event_page, flows, lua_event_registration, xml_handlers |
| Example locations | TidyChat: TidyChatFrames.Initialize, TidyRoll: TidyRollFrame.OnMButtonUp |
| XML usage count | 1 |
| XML attribute usage count | 1 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 3 |
| Initialization flow references | 1 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
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

Observed as an engine-supplied UI event hook used by 2 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- TidyChat
- TidyRoll

## Registrars And Handlers

- TidyChat.OnChannelButtonMButtonUp
- TidyRollFrame.OnMButtonUp
- WindowRegisterCoreEventHandler
- core

## Examples

- TidyChat: TidyChatFrames.Initialize -> OnMButtonUp -> TidyChat.OnChannelButtonMButtonUp
- TidyRoll: TidyRollFrame -> TidyRollFrame.OnMButtonUp -> TidyRollFrame.OnMButtonUp
- TidyChat: TidyChat.OnChannelButtonMButtonUp -> WindowRegisterCoreEventHandler(OnMButtonUp, TidyChat.OnChannelButtonMButtonUp)

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

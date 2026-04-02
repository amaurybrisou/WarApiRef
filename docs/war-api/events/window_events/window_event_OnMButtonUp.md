# OnMButtonUp

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 188

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
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
| Addons seen in | Enemy, TidyChat, TidyRoll, TurretRange, followTheLeader |
| Files seen in | `/workspace_addons/Enemy/Code/UnitFrames/UnitFrame.xml:93`, `/workspace_addons/TidyChat/TidyChat.lua:930`, `/workspace_addons/TidyRoll/TidyRoll.xml:202`, `/workspace_addons/TurrentRange/Display.xml:167`, `/workspace_addons/followTheLeader/followTheLeader.xml:11` |
| Namespaces detected | OnMButtonUp |
| Source kinds | event_page, flows, lua_event_registration, xml_handlers |
| Example locations | Enemy: EnemyUnitFrame.OnMButtonUp, TidyChat: TidyChatFrames.Initialize, TidyRoll: TidyRollFrame.OnMButtonUp, TurretRange: TurretMapDisplay.OnMButtonUp, followTheLeader: followTheLeaderWindow.OnMButtonUp |
| XML usage count | 4 |
| XML attribute usage count | 4 |
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

Observed as an engine-supplied UI event hook used by 5 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- Enemy
- TidyChat
- TidyRoll
- TurretRange
- followTheLeader

## Registrars And Handlers

- Enemy.UnitFramesUI_UnitFrame_OnMButtonUp
- Map.OnMButtonUp
- TidyChat.OnChannelButtonMButtonUp
- TidyRollFrame.OnMButtonUp
- WindowRegisterCoreEventHandler
- core
- followTheLeader.OnMButtonUp

## Examples

- TidyChat: TidyChatFrames.Initialize -> OnMButtonUp -> TidyChat.OnChannelButtonMButtonUp
- Enemy: EnemyUnitFrame -> EnemyUnitFrame.OnMButtonUp -> Enemy.UnitFramesUI_UnitFrame_OnMButtonUp
- TidyRoll: TidyRollFrame -> TidyRollFrame.OnMButtonUp -> TidyRollFrame.OnMButtonUp
- TurretRange: TurretMapDisplay -> TurretMapDisplay.OnMButtonUp -> Map.OnMButtonUp
- followTheLeader: followTheLeaderWindow -> followTheLeaderWindow.OnMButtonUp -> followTheLeader.OnMButtonUp
- TidyChat: TidyChat.OnChannelButtonMButtonUp -> WindowRegisterCoreEventHandler(OnMButtonUp, TidyChat.OnChannelButtonMButtonUp)

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../../xml/element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [OnMButtonUp](../../xml/handlers/handler_OnMButtonUp.md) (HIGH 100/100) - XML Handler
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- none

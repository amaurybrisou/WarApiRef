# AlertTextWindow.AddLine

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Score: 100/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, referenced by generated docs or reference files, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Aura, MapMonster, Targets, WSCT |
| Files seen in | `/workspace_addons/Aura/Source/Aura.lua:186`, `/workspace_addons/Aura/Source/Aura.lua:282`, `/workspace_addons/MapMonster/Source/MapMonster.lua:54`, `/workspace_addons/MapMonster/Source/MapMonster.lua:65`, `/workspace_addons/MapMonster/Source/MapMonster_EditorWindow.lua:1024`, `/workspace_addons/MapMonster/Source/MapMonster_EditorWindow.lua:763`, `/workspace_addons/MapMonster/Source/MapMonster_PinTypeEditorWindow.lua:682`, `/workspace_addons/targets/Targets.lua:214` |
| Namespaces detected | AlertTextWindow |
| Source kinds | globals, lua_calls |
| Example locations | Aura: Aura:Activate, Aura: Aura:Deactivate, MapMonster: ConfirmCreateSubType, MapMonster: MapMonster.Editor.OnClickLeft, MapMonster: MapMonster.Editor.OpenEditWindow, MapMonster: MapMonster.PinTypeEditor.LeftButton |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 20 |
| Global usage count | 20 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
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
AlertTextWindow.AddLine(arg1, arg2)
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: AuraConstants.AlertText[self:Get("activation-alerttexttype")].type, AuraConstants.AlertText[self:Get("deactivation-alerttexttype")].type, MapMonster.AlertTypes.ERROR |
| arg2 | Observed as a runtime window or control identifier. | Observed values: T["FailedAlert"], T["NotFoundAlert"], T["PinTypeSuccessAlert"] |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Aura
- MapMonster
- Targets
- WSCT

## Examples

- Aura: Aura:Activate -> AlertTextWindow.AddLine(AuraConstants.AlertText[self:Get("activation-alerttexttype")].type, self:Get("activation-alerttext"))
- Aura: Aura:Deactivate -> AlertTextWindow.AddLine(AuraConstants.AlertText[self:Get("deactivation-alerttexttype")].type, self:Get("deactivation-alerttext"))
- MapMonster: ConfirmCreateSubType -> AlertTextWindow.AddLine(MapMonster.AlertTypes.SUCCESS, T["SubTypeCreatedSuccess"])
- MapMonster: ConfirmCreateSubType -> AlertTextWindow.AddLine(MapMonster.AlertTypes.ERROR, T["SubTypeCreatedFailed"])
- MapMonster: ConfirmCreateSubType -> AlertTextWindow.AddLine(MapMonster.AlertTypes.ERROR, errorMessage)
- MapMonster: MapMonster.Editor.OnClickLeft -> AlertTextWindow.AddLine(MapMonster.AlertTypes.SUCCESS, T["SuccessAlert"])

## Related APIs

- none

## Used With

- [DialogManager.MakeTwoButtonDialog](global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](../../window_api/functions/window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](../../window_api/functions/window_WindowStopAlphaAnimation.md) (HIGH 100/100) - Window Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.

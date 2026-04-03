# AlertTextWindow.AddLine

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 13 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Aura, DetauntHelper, EmoteAlert, MapMonster, MegaphonePlus, MegaphonePlusPlus, Sequencer, SessionRPs |
| Files seen in | AAOTracker.lua, EmoteAlert.lua, MegaphonePlus.lua, MegaphonePlusPlus.lua, Sequencer.lua, Source/Aura.lua, Source/Common.lua, Source/MapMonster.lua |
| Namespaces detected | AlertTextWindow |
| Source kinds | lua_calls |
| Example locations | Aura: Activate, Aura: Deactivate, DetauntHelper: Print, EmoteAlert: showOnscreenMessage, MapMonster: ConfirmCreateSubType, MapMonster: LeftButton |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 33 |
| Global usage count | 33 |
| Local definition count | 0 |
| Documentation references | 0 |
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

Observed as a global function across 13 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: AAOTracker.Alerts[AAOTracker.Settings.AlertChannel].id, AuraConstants.AlertText[self:Get("activation-alerttexttype")].type, AuraConstants.AlertText[self:Get("deactivation-alerttexttype")].type |
| arg2 | Observed as a runtime window or control identifier. | Observed values: GetString(StringTables.Default.TEXT_MORALE_DROP_ERROR), GetString(StringTables.Default.TEXT_TACTIC_DROP_ERROR), L "Entry still locked" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Aura
- DetauntHelper
- EmoteAlert
- MapMonster
- MegaphonePlus
- MegaphonePlusPlus
- Sequencer
- SessionRPs
- SquaredClick
- Targets
- TomeTracker
- WSCT
- WarBoard_AAOTracker

## Examples

- Aura: Activate -> AlertTextWindow.AddLine(AuraConstants.AlertText[self:Get("activation-alerttexttype")].type, self:Get("activation-alerttext"))
- Aura: Deactivate -> AlertTextWindow.AddLine(AuraConstants.AlertText[self:Get("deactivation-alerttexttype")].type, self:Get("deactivation-alerttext"))
- DetauntHelper: Print -> AlertTextWindow.AddLine(SystemData.AlertText.Types.DEFAULT, msg)
- EmoteAlert: showOnscreenMessage -> AlertTextWindow.AddLine(alertType, toWide(message))
- MapMonster: ConfirmCreateSubType -> AlertTextWindow.AddLine(MapMonster.AlertTypes.SUCCESS, T["SubTypeCreatedSuccess"])
- MapMonster: ConfirmCreateSubType -> AlertTextWindow.AddLine(MapMonster.AlertTypes.ERROR, T["SubTypeCreatedFailed"])

## Used With

- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](../../window_api/functions/window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.AlertText.Types.DEFAULT](../../systemdata/fields/systemdata_SystemData.AlertText.Types.DEFAULT.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.

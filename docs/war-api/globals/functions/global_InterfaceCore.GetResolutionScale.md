# InterfaceCore.GetResolutionScale

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 143

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | EA_LoadingScreen, Effigy, GuardLine, MapMonster, MapPin, RoR_SoR, Squared |
| Files seen in | Effigy.lua, GuardLine.lua, RoR_SoR.lua, Source/MapMonster_Calibrate.lua, Squared.lua, source/MapPin.lua, source/scenarioenterloadingscreen.lua, source/standardloadingscreen.lua |
| Namespaces detected | InterfaceCore |
| Source kinds | lua_calls |
| Example locations | EA_LoadingScreen: BeginScenarioEnterLoadScreen, EA_LoadingScreen: BeginStandardLoadScreen, Effigy: LoadLayout, GuardLine: init, MapMonster: OnLMouseButton, MapMonster: OnMouseOverEnd |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
InterfaceCore.GetResolutionScale()
```

## Description

Observed as a global function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- EA_LoadingScreen
- Effigy
- GuardLine
- MapMonster
- MapPin
- RoR_SoR
- Squared

## Examples

- EA_LoadingScreen: BeginScenarioEnterLoadScreen -> InterfaceCore.GetResolutionScale()
- EA_LoadingScreen: BeginStandardLoadScreen -> InterfaceCore.GetResolutionScale()
- Effigy: LoadLayout -> InterfaceCore.GetResolutionScale()
- GuardLine: init -> InterfaceCore.GetResolutionScale()
- MapMonster: OnLMouseButton -> InterfaceCore.GetResolutionScale()
- MapMonster: OnMouseOverEnd -> InterfaceCore.GetResolutionScale()

## Related APIs

- [OnMouseOverEnd](../../xml/handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event

## Used With

- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [wstring.upper](global_wstring.upper.md) (HIGH 100/100) - Global Function

## Affects

- [SystemData.MousePosition.x](../../systemdata/fields/systemdata_SystemData.MousePosition.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.MousePosition.y](../../systemdata/fields/systemdata_SystemData.MousePosition.y.md) (HIGH 100/100) - SystemData Field
- [SystemData.screenResolution.x](../../systemdata/fields/systemdata_SystemData.screenResolution.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.screenResolution.y](../../systemdata/fields/systemdata_SystemData.screenResolution.y.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed

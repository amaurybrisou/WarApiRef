# WindowSetShowing

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 90/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin |
| Files seen in | ClosetGoblinCharacterWindow.lua, ClosetGoblinZoneWindow.lua |
| Namespaces detected | WindowSetShowing |
| Source kinds | lua_calls |
| Example locations | CM_ClosetGoblin: Hide, CM_ClosetGoblin: HideCloakOptions, CM_ClosetGoblin: HideShowHelm, CM_ClosetGoblin: OnInitialize, CM_ClosetGoblin: Show, CM_ClosetGoblin: ShowCloakOptions |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 30 |
| Global usage count | 30 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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
WindowSetShowing(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "ClosetGoblinCharacterWindow", "ClosetGoblinCharacterWindowContentsEquipment", "ClosetGoblinCharacterWindowContentsEquipmentShowCloak" |
| arg2 | Observed as a boolean toggle. | Observed values: ClosetGoblin.sortOrder==ClosetGoblin.SORT_ORDER_DOWN, ClosetGoblin.sortOrder==ClosetGoblin.SORT_ORDER_UP, false |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: Hide -> WindowSetShowing("ClosetGoblinZoneWindow", false)
- CM_ClosetGoblin: Hide -> WindowSetShowing("ClosetGoblinCharacterWindow", false)
- CM_ClosetGoblin: HideCloakOptions -> WindowSetShowing("ClosetGoblinCharacterWindowContentsEquipmentShowCloak", false)
- CM_ClosetGoblin: HideCloakOptions -> WindowSetShowing("ClosetGoblinCharacterWindowContentsEquipmentShowCloakHeraldry", false)
- CM_ClosetGoblin: HideShowHelm -> WindowSetShowing("ClosetGoblinCharacterWindowContentsEquipmentShowHelm", false)
- CM_ClosetGoblin: OnInitialize -> WindowSetShowing("ClosetGoblinCharacterWindow", false)

## Used With

- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 90/100) - Window Function

## Notes

- Only one addon surfaced this symbol in the current corpus.

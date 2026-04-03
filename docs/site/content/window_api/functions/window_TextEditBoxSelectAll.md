# TextEditBoxSelectAll

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | Aura, AutoBand, Enemy, Mass Refine, Squared, zMailMod |
| Files seen in | AutoBand.lua, Code/Core/Main.lua, MassRefine.lua, Source/AuraShares.lua, Squared.lua, zMailModMassMail.lua, zMailModOptions.lua, zMailModSend.lua |
| Namespaces detected | TextEditBoxSelectAll |
| Source kinds | lua_calls |
| Example locations | Aura: OnExportAura, Aura: OnImportAura, AutoBand: ShowCopyLink, Enemy: UI_TextEntryDialog_Open, Mass Refine: NewConfirmThenRefine, Squared: ExportSettings |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
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
TextEditBoxSelectAll(arg1)
```

## Description

Observed reading from or writing to edit-box controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "SquaredImportExportFrameData", "zMailModMassMailSubjectEditBox", COPY_LINK_WINDOW_NAME.."UrlInput" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Aura
- AutoBand
- Enemy
- Mass Refine
- Squared
- zMailMod

## Examples

- Aura: OnExportAura -> TextEditBoxSelectAll(exportWindow.."AuraText")
- Aura: OnImportAura -> TextEditBoxSelectAll(importWindow.."AuraText")
- AutoBand: ShowCopyLink -> TextEditBoxSelectAll(COPY_LINK_WINDOW_NAME.."UrlInput")
- Enemy: UI_TextEntryDialog_Open -> TextEditBoxSelectAll(wn.."Value")
- Mass Refine: NewConfirmThenRefine -> TextEditBoxSelectAll(MassRefine.sWindowName.."TextInput")
- Squared: ExportSettings -> TextEditBoxSelectAll("SquaredImportExportFrameData")

## Used With

- [TextEditBoxSetMaxChars](window_TextEditBoxSetMaxChars.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Notes

- none

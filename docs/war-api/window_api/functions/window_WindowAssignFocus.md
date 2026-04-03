# WindowAssignFocus

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 38 addons

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
| Addons seen in | Ace, Aura, AutoBand, Busted, CMap, Crusher, EA_UiDebugTools, EZCraftX |
| Files seen in | AutoBand.lua, Busted.lua, Code/Core/Main.lua, Code/Intercom/Intercom.lua, LibGUI.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, MassRefine.lua |
| Namespaces detected | WindowAssignFocus |
| Source kinds | lua_calls |
| Example locations | Ace: Defocus, Ace: Focus, Aura: OnExportAura, Aura: OnImportAura, AutoBand: ShowCopyLink, Busted: UpdateErrorView |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 108 |
| Global usage count | 108 |
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
WindowAssignFocus(arg1, arg2)
```

## Description

Observed as a window function across 38 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BustedGUIErrorMessage", "DebugWindowTextBox", "EA_TextEntryGroupEntryBoxTextInput" |
| arg2 | Observed as a boolean toggle. | Observed values: false, true |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- Aura
- AutoBand
- Busted
- CMap
- Crusher
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- Enemy
- GCDsaver
- Hopper
- InfoScroller
- Keyset
- LibWBToggler
- Map
- MapMonster
- Mass Refine
- Motion
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- Shinies
- SocialWindow 2.0
- Squared
- TargetRing
- TidyChat
- Tokens
- Twister
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- zMailMod

## Examples

- Ace: Defocus -> WindowAssignFocus(self.name, false)
- Ace: Focus -> WindowAssignFocus(self.name, true)
- Aura: OnExportAura -> WindowAssignFocus(exportWindow.."AuraText", true)
- Aura: OnImportAura -> WindowAssignFocus(importWindow.."AuraText", true)
- AutoBand: ShowCopyLink -> WindowAssignFocus(COPY_LINK_WINDOW_NAME.."UrlInput", true)
- Busted: UpdateErrorView -> WindowAssignFocus("BustedGUIErrorMessage", true)

## Related APIs

- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnKeyEscape](../../xml/handlers/handler_OnKeyEscape.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event

## Used With

- [TextEditBoxSelectAll](window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Notes

- none

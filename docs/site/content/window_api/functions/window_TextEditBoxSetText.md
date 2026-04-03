# TextEditBoxSetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 17 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, Aura, BuffHead, DAoCBuff, Enemy, Killer |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:656`, `/workspace/data/raw/Ace/LibGUI.lua:669`, `/workspace/data/raw/Ace/LibGUI.lua:711`, `/workspace/data/raw/Ace/LibGUI.lua:724`, `/workspace/data/raw/AdvancedPetAssist/APAGui.lua:983`, `/workspace/data/raw/AdvancedPetAssist/APAGuiHUD.lua:147`, `/workspace/data/raw/AdvancedPetAssist/APAGuiHUD.lua:163`, `/workspace/data/raw/AdvancedPetAssist/APAGuiHUD.lua:220` |
| Namespaces detected | TextEditBoxSetText |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_MultiTextbox:Clear, Ace: LIBGUI_MultiTextbox:SetText, Ace: LIBGUI_Textbox:Clear, Ace: LIBGUI_Textbox:SetText, AdvancedPetAssist: APAGui.ApplyHUDColorOff, AdvancedPetAssist: APAGui.ApplyHUDColorOn |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 203 |
| Global usage count | 203 |
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
TextEditBoxSetText(windowName, text)
```

## Description

Observed reading from or writing to edit-box controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: "APAEditHUDOffB", "APAEditHUDOffG", "APAEditHUDOffR" |
| text | Observed as a text or wstring payload. | Observed values: DAoCBuffSettings.TmpFilter[activefilter.index].name, DAoCBuffVar.Frames[activewindow.index].name, ERASE |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BuffHead
- DAoCBuff
- Enemy
- Killer
- LibWBToggler
- PartyCast
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WhoHealedMe
- WoH-Reticle

## Examples

- Ace: LIBGUI_MultiTextbox:Clear -> TextEditBoxSetText(self.name, L "")
- Ace: LIBGUI_MultiTextbox:SetText -> TextEditBoxSetText(self.name, towstring(text))
- Ace: LIBGUI_Textbox:Clear -> TextEditBoxSetText(self.name, L "")
- Ace: LIBGUI_Textbox:SetText -> TextEditBoxSetText(self.name, towstring(text))
- AdvancedPetAssist: APAGui.ApplyHUDColorOff -> TextEditBoxSetText("APAEditHUDOffR", towstring(APA.hudColorOffR))
- AdvancedPetAssist: APAGui.ApplyHUDColorOff -> TextEditBoxSetText("APAEditHUDOffG", towstring(APA.hudColorOffG))

## Related APIs

- none

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [TextEditBoxSelectAll](window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

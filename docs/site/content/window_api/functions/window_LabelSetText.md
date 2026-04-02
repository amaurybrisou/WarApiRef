# LabelSetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 48 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, BankArkel, BuffHead, Busted |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:426`, `/workspace_addons/Ace/LibGUI.lua:439`, `/workspace_addons/AdvancedPetAssist/APAGui.lua:983`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:11`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:181`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:241`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:98`, `/workspace_addons/AggroMeter/AggroMeter.lua:5` |
| Namespaces detected | LabelSetText |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Label:Clear, Ace: LIBGUI_Label:SetText, AdvancedPetAssist: APAGui.OnShown, AdvancedPetAssist: APAGui.UpdateFollowTargetHUD, AdvancedPetAssist: APAGui.UpdateInstantOnlyHUD, AdvancedPetAssist: APAGui.UpdateKitingHUD |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1401 |
| Global usage count | 1401 |
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
LabelSetText(windowName, text)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: "APAFollowTargetHUDLabel", "APAInstantOnlyHUDLabel", "APAKitingHUDLabel" |
| text | Observed as a text or wstring payload. | Observed values: APA.PetTarget.cachedName, AddonToShow.RVCredits, AddonToShow.RVLicense |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- BankArkel
- BuffHead
- Busted
- CM_ClosetGoblin
- CombatTextNames
- Crafting Info Tooltip
- DAoCBuff
- DaemonAssist
- EA_UiDebugTools
- Effigy
- Enemy
- FastInteract
- GCDsaver
- GetStats
- JunkDump
- Killer
- LibGroup
- LibWBToggler
- MapMonster
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- Moth
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WhoHealedMe
- WoH-Reticle
- wbLeadHelper

## Examples

- Ace: LIBGUI_Label:Clear -> LabelSetText(self.name, L "")
- Ace: LIBGUI_Label:SetText -> LabelSetText(self.name, towstring(text))
- AdvancedPetAssist: APAGui.OnShown -> LabelSetText("APAOptionsTitleText", L "AdvancedPetAssist")
- AdvancedPetAssist: APAGui.OnShown -> LabelSetText("APALabelSectionAutoRecall", L "--- Auto Recall ---")
- AdvancedPetAssist: APAGui.OnShown -> LabelSetText("APALabelSectionMouseControls", L "--- Mouse Controls ---")
- AdvancedPetAssist: APAGui.OnShown -> LabelSetText("APALabelSectionLos", L "--- LOS Detection ---")

## Related APIs

- none

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [EA_Window_ContextMenu.AddUserDefinedMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddUserDefinedMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [StatusBarSetCurrentValue](window_StatusBarSetCurrentValue.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

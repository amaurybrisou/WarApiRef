# WindowSetShowing

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 34 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainerAdditions, Aura, AutoMark, BankArkel |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:74`, `/workspace/data/raw/Ace/LibGUI.lua:79`, `/workspace/data/raw/AdvancedPetAssist/APAGui.lua:1063`, `/workspace/data/raw/AdvancedPetAssist/APAGui.lua:1071`, `/workspace/data/raw/AdvancedPetAssist/APAGui.lua:546`, `/workspace/data/raw/AdvancedPetAssist/APAGui.lua:983`, `/workspace/data/raw/AdvancedPetAssist/APAGuiHUD.lua:138`, `/workspace/data/raw/AdvancedPetAssist/APAGuiHUD.lua:181` |
| Namespaces detected | WindowSetShowing |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Hide, Ace: LIBGUI_ELEMENT:Show, AdvancedPetAssist: APAGui.Hide, AdvancedPetAssist: APAGui.HidePetTargetHUD, AdvancedPetAssist: APAGui.OnShown, AdvancedPetAssist: APAGui.Show |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 912 |
| Global usage count | 912 |
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
WindowSetShowing(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "APAInstantOnlyHUD", "APAKitingHUD", "APAOptions" |
| arg2 | Observed as a boolean toggle. | Observed values: (KEEP1_State==4), (KEEP2_State==4), (NameOfTargetBG~=L "")and PartyCast.Settings.ShowTarget |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainerAdditions
- Aura
- AutoMark
- BankArkel
- BuffHead
- CM_ClosetGoblin
- CombatTextNames
- DAoCBuff
- Enemy
- GuardLine
- Killer
- LibGroup
- LibSlash
- LibWBToggler
- MiracleGrowLight
- PartyCast
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- Swift Assist
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- WoH-Reticle
- followTheLeader

## Examples

- Ace: LIBGUI_ELEMENT:Hide -> WindowSetShowing(self.name, false)
- Ace: LIBGUI_ELEMENT:Show -> WindowSetShowing(self.name, true)
- AdvancedPetAssist: APAGui.Hide -> WindowSetShowing("APAOptions", false)
- AdvancedPetAssist: APAGui.HidePetTargetHUD -> WindowSetShowing("APAPetTargetHUD", false)
- AdvancedPetAssist: APAGui.OnShown -> WindowSetShowing("APAOptionsTabsControls", false)
- AdvancedPetAssist: APAGui.Show -> WindowSetShowing("APAOptions", true)

## Related APIs

- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 100/100) - XML Event

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](window_WindowStopAlphaAnimation.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Event
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

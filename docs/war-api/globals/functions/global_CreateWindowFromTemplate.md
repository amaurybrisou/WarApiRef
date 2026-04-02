# CreateWindowFromTemplate

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 31 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, AutoMark, BankArkel, BankWindowFix |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:1046`, `/workspace_addons/Ace/LibGUI.lua:1149`, `/workspace_addons/Ace/LibGUI.lua:1237`, `/workspace_addons/Ace/LibGUI.lua:344`, `/workspace_addons/Ace/LibGUI.lua:399`, `/workspace_addons/Ace/LibGUI.lua:506`, `/workspace_addons/Ace/LibGUI.lua:596`, `/workspace_addons/Ace/LibGUI.lua:620` |
| Namespaces detected | CreateWindowFromTemplate |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:New, Ace: LIBGUI_Checkbox:New, Ace: LIBGUI_Closebutton:New, Ace: LIBGUI_Combobox:New, Ace: LIBGUI_Image:New, Ace: LIBGUI_Label:New |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 249 |
| Global usage count | 249 |
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
CreateWindowFromTemplate(windowName, templateName, parentWindow)
```

## Description

Observed instantiating repeated UI elements from an XML template.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a runtime window instance name. | Observed values: "AggroMeterWindow"..MobID, "BottomBoard", "EffigySwingTimerWindow" |
| templateName | Observed as an XML template name. | Observed values: "AbilityButtonTemplate", "AggroMeterWindow", "AnimatedHighlight" |
| parentWindow | Observed as a parent window name. | Observed values: "ChatChannelSelectionWindow", "EA_Window_EventTextContainer", "EA_Window_WorldMapZoneView" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Creates or instantiates UI objects from XML-backed definitions.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- AutoMark
- BankArkel
- BankWindowFix
- BuffHead
- CombatTextNames
- Crafting Info Tooltip
- DAoCBuff
- Effigy
- Enemy
- GCDsaver
- GetStats
- LibWBToggler
- MapMonster
- Miracle Grow Remix
- PotionBar
- QuickWarReport
- RoR_SoR
- Shinies
- TidyChat
- TidyRoll
- TurretRange
- WarBoard
- WhoHealedMe
- WoH-Reticle
- nRarity
- wbLeadHelper

## Examples

- Ace: LIBGUI_Button:New -> CreateWindowFromTemplate(w.name, base, w.parent)
- Ace: LIBGUI_Checkbox:New -> CreateWindowFromTemplate(w.name, base, w.parent)
- Ace: LIBGUI_Closebutton:New -> CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
- Ace: LIBGUI_Combobox:New -> CreateWindowFromTemplate(w.name, base, w.parent)
- Ace: LIBGUI_Image:New -> CreateWindowFromTemplate(w.name, base, w.parent)
- Ace: LIBGUI_Label:New -> CreateWindowFromTemplate(w.name, base, w.parent)

## Related APIs

- none

## Used With

- [ButtonSetCheckButtonFlag](../../window_api/functions/window_ButtonSetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](../../window_api/functions/window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [EA_Window_ContextMenu.AddUserDefinedMenuItem](global_EA_Window_ContextMenu.AddUserDefinedMenuItem.md) (HIGH 100/100) - Global Function
- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](../../window_api/functions/window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](../../window_api/functions/window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxSelectAll](../../window_api/functions/window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](../../window_api/functions/window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetId](../../window_api/functions/window_WindowSetId.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](../../window_api/functions/window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [TextEditBoxSetMaxChars](../../window_api/functions/window_TextEditBoxSetMaxChars.md) (HIGH 80/100) - Window Function
- [DestroyWindow](global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.

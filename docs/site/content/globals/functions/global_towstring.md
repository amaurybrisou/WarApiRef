# towstring

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 56 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, BagOMatic, BankArkel, BuffHead |
| Files seen in | `/workspace_addons/Ace/AceLocale-3.0.lua:57`, `/workspace_addons/Ace/LibGUI.lua:1068`, `/workspace_addons/Ace/LibGUI.lua:1124`, `/workspace_addons/Ace/LibGUI.lua:426`, `/workspace_addons/Ace/LibGUI.lua:533`, `/workspace_addons/Ace/LibGUI.lua:656`, `/workspace_addons/Ace/LibGUI.lua:711`, `/workspace_addons/AdvancedPetAssist/APACommands.lua:207` |
| Namespaces detected | towstring |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:SetText, Ace: LIBGUI_Combobox:Add, Ace: LIBGUI_Combobox:Select, Ace: LIBGUI_Label:SetText, Ace: LIBGUI_MultiTextbox:SetText, Ace: LIBGUI_Textbox:SetText |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1468 |
| Global usage count | 1468 |
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
towstring(arg1)
```

## Description

Observed as a global function across 56 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: "    Posting COUNT auction(s) for ITEM with a price of PRICE per stack of SIZE.", " Missing roles\n "..stats.." \n", "" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- BagOMatic
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
- GCDsaver
- GetStats
- GuardLine
- HealAll
- JunkDump
- Killer
- LibGroup
- LibGuard
- LibSlash
- LibWBToggler
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- NPC Item Sale Price
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
- TortallDPSCore
- TurretRange
- WSCT
- WhoHealedMe
- WoH-Reticle
- followTheLeader
- wbLeadHelper

## Examples

- Ace: LIBGUI_Button:SetText -> towstring(text)
- Ace: LIBGUI_Combobox:Add -> towstring(itemText)
- Ace: LIBGUI_Combobox:Select -> towstring(value)
- Ace: LIBGUI_Label:SetText -> towstring(text)
- Ace: LIBGUI_MultiTextbox:SetText -> towstring(text)
- Ace: LIBGUI_Textbox:SetText -> towstring(text)

## Related APIs

- [AlertText](../tables/table_AlertText.md) (HIGH 100/100) - Global Table
- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.UnregisterSlashCmd](global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeTextEntryDialog](global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [EA_ChatTabManager.GetTabName](global_EA_ChatTabManager.GetTabName.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [EA_ChatWindowGroups](../tables/table_EA_ChatWindowGroups.md) (HIGH 100/100) - Global Table
- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [StatusBarSetCurrentValue](../../window_api/functions/window_StatusBarSetCurrentValue.md) (HIGH 100/100) - Window Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [wstring.match](global_wstring.match.md) (HIGH 100/100) - Global Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [wstring.format](global_wstring.format.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed

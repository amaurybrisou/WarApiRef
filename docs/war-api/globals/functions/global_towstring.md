# towstring

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 34 addons

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
| Files seen in | `/workspace/data/raw/Ace/AceLocale-3.0.lua:57`, `/workspace/data/raw/Ace/LibGUI.lua:1068`, `/workspace/data/raw/Ace/LibGUI.lua:1124`, `/workspace/data/raw/Ace/LibGUI.lua:426`, `/workspace/data/raw/Ace/LibGUI.lua:533`, `/workspace/data/raw/Ace/LibGUI.lua:656`, `/workspace/data/raw/Ace/LibGUI.lua:711`, `/workspace/data/raw/AdvancedPetAssist/APACommands.lua:207` |
| Namespaces detected | towstring |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:SetText, Ace: LIBGUI_Combobox:Add, Ace: LIBGUI_Combobox:Select, Ace: LIBGUI_Label:SetText, Ace: LIBGUI_MultiTextbox:SetText, Ace: LIBGUI_Textbox:SetText |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 856 |
| Global usage count | 856 |
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

Observed as a global function across 34 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: "    Posting COUNT auction(s) for ITEM with a price of PRICE per stack of SIZE.", "", "%d items remaining" |

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
- CM_ClosetGoblin
- CombatTextNames
- DAoCBuff
- Enemy
- GuardLine
- Killer
- LibGroup
- LibGuard
- LibSlash
- LibWBToggler
- MiracleGrowLight
- MoraleCircle
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
- WhoHealedMe
- WoH-Reticle
- followTheLeader

## Examples

- Ace: LIBGUI_Button:SetText -> towstring(text)
- Ace: LIBGUI_Combobox:Add -> towstring(itemText)
- Ace: LIBGUI_Combobox:Select -> towstring(value)
- Ace: LIBGUI_Label:SetText -> towstring(text)
- Ace: LIBGUI_MultiTextbox:SetText -> towstring(text)
- Ace: LIBGUI_Textbox:SetText -> towstring(text)

## Related APIs

- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.UnregisterSlashCmd](global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](../../window_api/functions/window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [PartyUtils.GetPartyData](global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [wstring.format](global_wstring.format.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed

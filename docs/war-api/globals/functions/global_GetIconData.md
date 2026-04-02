# GetIconData

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 83/100
- Seen in: 23 addons

## Confidence Assessment

- Level: HIGH

- Score: 83/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, Aura, AutoMark, BankArkel, BuffHead, CM_ClosetGoblin, CombatTextNames, DAoCBuff |
| Files seen in | `/workspace_addons/Aura/Source/AuraAddon.lua:238`, `/workspace_addons/AutoMark/Source/AutoMark.lua:7`, `/workspace_addons/BankArkel/BankArkel.lua:242`, `/workspace_addons/BuffHead/EffectFrame.lua:52`, `/workspace_addons/ClosetGoblin/ClosetGoblinCharacterWindow.lua:384`, `/workspace_addons/DAoCBuff/Source/DAoCBuffFrames.lua:159`, `/workspace_addons/Effigy/Elements/EffigyIcons.lua:138`, `/workspace_addons/Effigy/Elements/EffigyIcons.lua:172` |
| Namespaces detected | GetIconData |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTrainer.local.CreateAbilityWindow, AdvancedRenownTrainer: CreateAbilityWindow, Aura: AuraAddon.GetTextureData, AutoMark: AutoMark.local.CreateMarker, AutoMark: CreateMarker, BankArkel: BankArkel.PackShow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 68 |
| Global usage count | 68 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
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
GetIconData(arg1)
```

## Description

Observed as a shared query API across 23 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: 005118, 32000+tarCareerID, 64 |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AdvancedRenownTrainer
- Aura
- AutoMark
- BankArkel
- BuffHead
- CM_ClosetGoblin
- CombatTextNames
- DAoCBuff
- Effigy
- Enemy
- GetStats
- MapPin
- MoraleCircle
- Moth
- PotionBar
- RandomMount
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- TidyRoll
- WSCT
- wbLeadHelper

## Examples

- AdvancedRenownTrainer: AdvancedRenownTrainer.local.CreateAbilityWindow -> GetIconData(t.icon)
- AdvancedRenownTrainer: CreateAbilityWindow -> GetIconData(t.icon)
- Aura: AuraAddon.GetTextureData -> GetIconData(tonumber(id))
- AutoMark: AutoMark.local.CreateMarker -> GetIconData(career_icon_id)
- AutoMark: CreateMarker -> GetIconData(career_icon_id)
- BankArkel: BankArkel.PackShow -> GetIconData(BankArkel.db.Entry[index].Equip[i].iconNum)

## Related APIs

- none

## Used With

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [EA_Window_ContextMenu.AddCascadingMenuItem](global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [LabelSetTextAlign](../../window_api/functions/window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetId](../../window_api/functions/window_WindowSetId.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.

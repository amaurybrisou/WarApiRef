# GetIconData

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 83/100
- Seen in: 73 addons

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
| Addons seen in | AdvancedRenownTrainer, Amethyst, AuctionStats, Aura, AutoMark, BankArkel, BuffHead, CCTV |
| Files seen in | AdvancedRenownTraining.lua, Amethyst.lua, BankArkel.lua, CCTV.lua, CDownFrames.lua, CallingList.lua, CallingNotification.lua, CallingSetup.lua |
| Namespaces detected | GetIconData |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: CreateAbilityWindow, Amethyst: ShowCastBar, Amethyst: ShowDummyCastBar, AuctionStats: UpdateItemOptions, Aura: GetTextureData, AutoMark: CreateMarker |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 229 |
| Global usage count | 229 |
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

Observed as a shared query API across 73 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: 0, 005118, 1 |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedRenownTrainer
- Amethyst
- AuctionStats
- Aura
- AutoMark
- BankArkel
- BuffHead
- CCTV
- CDown
- CM_ClosetGoblin
- CaVES
- Calling
- CastSequence
- CleanUnitFrames
- CleansingBuddy
- CombatTextNames
- Crusher
- DAoCBuff
- DammazKron
- DetauntHelper
- DuffTimer
- EA_ScenarioGroupWindow
- EZCraftX
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- GCDTracker
- GetStats
- Group Icons
- Group Icons SG
- GuildWarden
- HealGrid
- MacroIcons
- MapPin
- MarkBuff
- MoraleCircle
- Moth
- Motion
- NAMBLA
- NaturalLog
- NerfedButtons
- Obsidian
- PartyCast
- PotionBar
- Pure
- RandomMount
- SNT_BUTTONS
- SNT_CASTBAR
- Shinies
- Squared
- SquaredClick
- Swift Assist
- Swinger
- Targets
- TastyButtons
- TexturedButtons
- TheSeeker
- TidyRoll
- TokenMachine
- Tokens
- Trakario
- Twister
- WSCT
- WarBoard_TaliMon
- WhatsCooking
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- talisman-monitor
- wbLeadHelper
- xHUD
- yAssistHelper
- zMailMod

## Examples

- AdvancedRenownTrainer: CreateAbilityWindow -> GetIconData(t.icon)
- Amethyst: ShowCastBar -> GetIconData(abilityData.iconNum)
- Amethyst: ShowDummyCastBar -> GetIconData(123)
- AuctionStats: UpdateItemOptions -> GetIconData(itemData.iconNum)
- Aura: GetTextureData -> GetIconData(tonumber(id))
- AutoMark: CreateMarker -> GetIconData(career_icon_id)

## Related APIs

- [Texture](../../xml/element_types/element_Texture.md) (HIGH 100/100) - XML Element Type
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [DialogManager.MakeTextEntryDialog](global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](../../window_api/functions/window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow.InsertText](global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelGetTextDimensions](../../window_api/functions/window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [WindowSetId](../../window_api/functions/window_WindowSetId.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [EA_Window_WorldMap.ShowZone](global_EA_Window_WorldMap.ShowZone.md) (HIGH 90/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.

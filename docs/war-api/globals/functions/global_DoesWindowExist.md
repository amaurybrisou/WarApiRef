# DoesWindowExist

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 83/100
- Seen in: 41 addons

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
| Addons seen in | AdvancedPetAssist, AggroMeter, Aura, BagOMatic, Busted, CM_ClosetGoblin, Crafting Info Tooltip, DAoCBuff |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.lua:1063`, `/workspace_addons/AdvancedPetAssist/APAGui.lua:546`, `/workspace_addons/AdvancedPetAssist/APAGui.lua:683`, `/workspace_addons/AdvancedPetAssist/APAGui.lua:704`, `/workspace_addons/AdvancedPetAssist/APAGui.lua:983`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:11`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:138`, `/workspace_addons/AdvancedPetAssist/APAGuiHUD.lua:147` |
| Namespaces detected | DoesWindowExist |
| Source kinds | lua_calls |
| Example locations | AdvancedPetAssist: APAGui.ApplyKitingHUDColorOff, AdvancedPetAssist: APAGui.ApplyKitingHUDColorOn, AdvancedPetAssist: APAGui.ApplyPTColor, AdvancedPetAssist: APAGui.ApplyPetTargetHUDLayout, AdvancedPetAssist: APAGui.HidePetTargetHUD, AdvancedPetAssist: APAGui.InitHUDEditBoxes |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 337 |
| Global usage count | 337 |
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
DoesWindowExist(windowName)
```

## Description

Observed guarding runtime window creation and cleanup by checking whether a named window exists.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "APAEditHUDOnR", "APAEditKitingHUDOffR", "APAEditKitingHUDOnR" |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AdvancedPetAssist
- AggroMeter
- Aura
- BagOMatic
- Busted
- CM_ClosetGoblin
- Crafting Info Tooltip
- DAoCBuff
- DaemonAssist
- EA_UiDebugTools
- Effigy
- Enemy
- FastInteract
- GetStats
- GuardLine
- JunkDump
- Killer
- LibSlash
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MoraleCircle
- PetFixWindow
- PotionBar
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Targets
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WhoHealedMe
- WoH-Reticle
- nRarity
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAGui.ApplyKitingHUDColorOff -> DoesWindowExist("APAEditKitingHUDOffR")
- AdvancedPetAssist: APAGui.ApplyKitingHUDColorOn -> DoesWindowExist("APAEditKitingHUDOnR")
- AdvancedPetAssist: APAGui.ApplyPTColor -> DoesWindowExist("APAEditPTR")
- AdvancedPetAssist: APAGui.ApplyPetTargetHUDLayout -> DoesWindowExist("APAPetTargetHUD")
- AdvancedPetAssist: APAGui.HidePetTargetHUD -> DoesWindowExist("APAPetTargetHUD")
- AdvancedPetAssist: APAGui.InitHUDEditBoxes -> DoesWindowExist("APAEditHUDOnR")

## Related APIs

- none

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [DynamicImageSetRotation](../../window_api/functions/window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](../../window_api/functions/window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](../../window_api/functions/window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [SystemData.Events.AUCTION_BID_RESULT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_BID_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_INIT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_INIT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [DestroyWindow](global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.

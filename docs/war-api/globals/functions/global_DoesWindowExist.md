# DoesWindowExist

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 83/100
- Seen in: 122 addons

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
| Addons seen in | AdvancedPetAssist, AggroMeter, Atlas, Aura, BBars - Mechanic Only, BagOMatic, Busted, CM_ClosetGoblin |
| Files seen in | APAGui.lua, APAGuiHUD.lua, AggroMeter.lua, BBars.lua, BBarsPetHP.lua, BBarsPlayerMechanic.lua, BagOMatic.lua, Busted.lua |
| Namespaces detected | DoesWindowExist |
| Source kinds | lua_calls |
| Example locations | AdvancedPetAssist: ApplyKitingHUDColorOff, AdvancedPetAssist: ApplyKitingHUDColorOn, AdvancedPetAssist: ApplyPTColor, AdvancedPetAssist: ApplyPetTargetHUDLayout, AdvancedPetAssist: EnsureSmallEditBox, AdvancedPetAssist: HidePetTargetHUD |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 752 |
| Global usage count | 752 |
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

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedPetAssist
- AggroMeter
- Atlas
- Aura
- BBars - Mechanic Only
- BagOMatic
- Busted
- CM_ClosetGoblin
- CMap
- ChattyCathy
- CleansingBuddy
- CoolDownLine
- Crafting Info Tooltip
- CraftingWillard
- Crusher
- DAoCBuff
- DaemonAssist
- DammazKron
- DeepSleep
- DetauntHelper
- DuffTimer
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EA_UiModWindow
- EZCraftX
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- FastFriends
- FastInteract
- FozAuction
- GetStats
- GroupSpotter
- GuardLine
- HealGrid
- Hopper
- InfoScroller
- JunkDump
- KeyBar
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- LibSlash
- LibSurveyor
- LoyalPet
- Map
- MapMonster
- MapPin
- MiniMapMonster
- Minmap
- Miracle Grow Remix
- MoraleCircle
- MoraleSet
- Motion
- NaturalLog
- Obsidian
- PetFixWindow
- Phantom
- PotionBar
- Pure
- QuickWarReport
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RVMOD_Targets
- RandomMount
- RealmStatus
- Refer
- ReliquaryHunter
- RoR_SoR
- RoR_debolster
- RvRContribution
- SNT_INFO
- ScenarioStats
- Sequencer
- Shinies
- ShowHealthPercent
- SocialWindow 2.0
- Soloq
- Squared
- SquaredClick
- SquaredHotIndicators
- Statdoll Remix
- Swinger
- TacticSetNames
- TargetInfoRing
- TargetRing
- Targets
- TastyButtons
- TexturedButtons
- ThankTheResser
- TidyChat
- TidyQueue
- TidyRoll
- TokenMachine
- Tokens
- TurretRange
- Vectors
- VerticalMorale
- VerticalTactics
- WARCommander
- WSCT
- WTes
- WarBoard_WarWhisperer
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- bigger_MacroWindow
- minesweep
- nRarity
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- AdvancedPetAssist: ApplyKitingHUDColorOff -> DoesWindowExist("APAEditKitingHUDOffR")
- AdvancedPetAssist: ApplyKitingHUDColorOn -> DoesWindowExist("APAEditKitingHUDOnR")
- AdvancedPetAssist: ApplyPTColor -> DoesWindowExist("APAEditPTR")
- AdvancedPetAssist: ApplyPetTargetHUDLayout -> DoesWindowExist("APAPetTargetHUD")
- AdvancedPetAssist: EnsureSmallEditBox -> DoesWindowExist(name)
- AdvancedPetAssist: HidePetTargetHUD -> DoesWindowExist("APAPetTargetHUD")

## Related APIs

- [ButtonSetDisabledFlag](../../window_api/functions/window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [DynamicImageSetRotation](../../window_api/functions/window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](../../window_api/functions/window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAlpha](../../window_api/functions/window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetAnchorCount](../../window_api/functions/window_WindowGetAnchorCount.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetOffsetFromParent](../../window_api/functions/window_WindowGetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetId](../../window_api/functions/window_WindowSetId.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowUnregisterEventHandler](../../window_api/functions/window_WindowUnregisterEventHandler.md) (HIGH 100/100) - Window Function
- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [DestroyWindow](global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Used With

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](../../window_api/functions/window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [DestroyWindow](global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed

# LayoutEditor.RegisterWindow

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 121 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AbilityAlert, AbilityNotifier, ActionFraction, ActionPoints, Amethyst, Atlas, AutoBand, BBars - Mechanic Only |
| Files seen in | AAOTracker.lua, AbilityAlert.lua, AbilityNotifier.lua, ActionPoints.lua, AdvancedContainers.lua, Amethyst.lua, AutoBand.lua, BBarsPlayerMechanic.lua |
| Namespaces detected | LayoutEditor |
| Source kinds | lua_calls |
| Example locations | AbilityAlert: Initialize, AbilityNotifier: Initialize, ActionFraction: Initialize, ActionPoints: Initialize, Amethyst: Recreate, Atlas: LayoutEditorRegistrations |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 189 |
| Global usage count | 189 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
LayoutEditor.RegisterWindow(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
```

## Description

Observed as a window function across 121 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "AAWindow", "AbHelpWindow", "ActionBarLockToggler" |
| arg2 | Observed as a text or wstring payload. | Observed values: Atlas.Strings.AtlasConfigurationFrameTitle, Atlas.Strings.AtlasFrameTitle, Calling.GetLocalized("callingIconCaption") |
| arg3 | Observed as a text or wstring payload. | Observed values: "", "Text", Atlas.Strings.AtlasConfigurationFrameDescription |
| arg4 | Observed as a boolean toggle. | Observed values: false, true |
| arg5 | Observed as a boolean toggle. | Observed values: false, true |
| arg6 | Observed as a boolean toggle. | Observed values: false, true |
| arg7 | Observed as a runtime window or control identifier. | Observed values: nil |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- AbilityAlert
- AbilityNotifier
- ActionFraction
- ActionPoints
- Amethyst
- Atlas
- AutoBand
- BBars - Mechanic Only
- Brizio's Crappy Computer Medic
- BuffHead
- Busted
- CDown
- CM_ClosetGoblin
- CMap
- Calling
- ChattyCathy
- CleanUnitFrames
- CleansingBuddy
- Clock
- Countdown
- CraftingWillard
- Crusher
- DAoCBuff
- DammazKron
- DetauntHelper
- DuffTimer
- EA_OpenPartyWindow
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- FixGit
- FlagCap
- GCDTracker
- Group Icons SG
- GroupRange
- GroupSpotter
- GuardLine
- GuardList
- GuardRange
- GuildWarden
- HealGrid
- Hopper
- InfoScroller
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- Map
- MarkBuff
- Mech
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleSet
- Moth
- Motion
- MyReasons
- NerfedButtons
- NoOverheal
- Obsidian
- Phantom
- Pure
- QuickWarReport
- RVMOD_Manager
- RVMOD_PlayerStatus
- Rangechecker
- RealmStatus
- ResHelp
- RetAlert
- RoR_SoR
- Rotation
- RvRContribution
- SNT_BUTTONS
- SNT_CASTBAR
- SNT_INFO
- SNT_PANEL
- SOR
- SessionRPs
- Soloq
- Squared
- Statdoll
- Statdoll Light
- Swift Assist
- Swinger
- TalismanGenie
- Targets
- TastyButtons
- TexturedButtons
- ThankTheTank
- TidyChat
- TidyRoll
- Tokens
- Tome Titan
- TomeTracker
- Tortall_DPS
- Trakario
- TurretRange
- TurretScrap
- Twister
- TwisterSet
- VPBreakdown
- WARCommander
- WTes
- WarBoard_AAOTracker
- WarBoard_WarWhisperer
- WhatsCooking
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- XpStatus+G
- ZonePOP
- compass
- emotes
- followTheLeader
- fpsbox
- nLootLink
- talisman-monitor
- yAssistHelper
- zMailMod

## Examples

- AbilityAlert: Initialize -> LayoutEditor.RegisterWindow("AAWindow", L "Ability Alert!", L "Alerts when cooldown completed and combat events.", false, false, true, nil)
- AbilityNotifier: Initialize -> LayoutEditor.RegisterWindow("AbHelpWindow", L "Ability Notifier", L "Reminds you to use skills.", false, false, true, nil)
- ActionFraction: Initialize -> LayoutEditor.RegisterWindow(windowName, L "Action Fraction", L "Displays character's current Action Points", true, true, true, nil)
- ActionPoints: Initialize -> LayoutEditor.RegisterWindow("ActionPointsWindow", L "Action Points Indicator", L "Displays your current AP.", false, false, true, nil)
- Amethyst: Recreate -> LayoutEditor.RegisterWindow(C.name, L "Amethyst CastBar", "", true, true, true)
- Atlas: LayoutEditorRegistrations -> LayoutEditor.RegisterWindow("AtlasFrame", Atlas.Strings.AtlasFrameTitle, Atlas.Strings.AtlasFrameDescription, false, false, true, nil)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Affects

- [SystemData.ChatLogFilters.SHOUT](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field

## Notes

- Advanced return analysis: No strong return evidence observed

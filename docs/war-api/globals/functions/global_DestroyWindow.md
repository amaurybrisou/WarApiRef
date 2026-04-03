# DestroyWindow

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 91 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, Aura, AutoMark, BBars - Mechanic Only, BankWindowFix, BuffHead |
| Files seen in | AAOTracker.lua, BBars.lua, ChattyCathy.lua, CleanCastbar.lua, CleansingBuddy.lua, Code/Core/Main.lua, Code/GroupIcons/GroupIcon.lua, Code/Marks/MarkTemplate.lua |
| Namespaces detected | DestroyWindow |
| Source kinds | lua_calls |
| Example locations | Ace: Destroy, ActionBarHide: Destroy, Amethyst: Destroy, Aura: DeleteWindow, AutoMark: DestroyMarker, BBars - Mechanic Only: OnShutdown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 213 |
| Global usage count | 213 |
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
DestroyWindow(windowName)
```

## Description

Observed tearing down runtime-created windows.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "BBarsPetHPBG", "BBarsPlayerMechanic3BG", "BBarsPlayerMechanic3Frontbar" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Destroys a runtime-created window instance.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- Aura
- AutoMark
- BBars - Mechanic Only
- BankWindowFix
- BuffHead
- CMap
- CastSequence
- ChattyCathy
- CleanCastbar
- CleansingBuddy
- CoolDownLine
- CraftingWillard
- Crusher
- DAoCBuff
- DuffTimer
- EA_ThreePartBar
- EZCraftX
- EZGuard
- Effigy
- Enemy
- FastInteract
- GCDsaver
- GetStats
- Group Icons
- Group Icons SG
- GroupRange
- HealGrid
- Hopper
- InfoScroller
- JunkDump
- KillTracker
- Kwestor
- LibWBToggler
- Map
- MapMonster
- MiniMapMonster
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- NoUselessMods-Assist
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- QuickWarReport
- RO-Style Combat Text
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RVMOD_Targets
- RandomMount
- RealmStatus
- RoR_SoR
- Rotation
- RvRContribution
- SNT_BUTTONS
- SNT_CASTBAR
- ScenarioStats
- SessionRPs
- Shinies
- ShowHealthPercent
- Statdoll Remix
- TargetInfoRing
- TargetRing
- TastyButtons
- ThankTheTank
- TidyRoll
- Tokens
- VerticalMorale
- WSCT
- WarBoard_AAOTracker
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- emotes
- minesweep
- nRarity
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Destroy -> DestroyWindow(self.name)
- ActionBarHide: Destroy -> DestroyWindow(self.name)
- Amethyst: Destroy -> DestroyWindow(self.name)
- Aura: DeleteWindow -> DestroyWindow(windowId)
- AutoMark: DestroyMarker -> DestroyWindow(marker.window_name)
- BBars - Mechanic Only: OnShutdown -> DestroyWindow("BBarsPlayerMechanic3BG")

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [UnregisterEventHandler](global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [LayoutEditor.UnregisterWindow](../../window_api/functions/window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [WindowGetOffsetFromParent](../../window_api/functions/window_WindowGetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [UnregisterEventHandler](global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed

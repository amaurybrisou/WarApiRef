# WindowRegisterCoreEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 58 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 153

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, ActionBarHide, ActionFraction, AdjustTheTip, Amethyst, BetterCC, BuddyBind, CMap |
| Files seen in | AdjustTheTip.lua, BetterCC.lua, BuddyBind.lua, CallingKeybinding.lua, CharacterScreenTabFix.lua, CleanTargetWindow.lua, Core.lua, DuffTimer.lua |
| Namespaces detected | WindowRegisterCoreEventHandler |
| Source kinds | lua_calls |
| Example locations | Ace: RegisterEvent, ActionBarHide: RegisterEvent, ActionFraction: Initialize, AdjustTheTip: CreateCheckBoxMenuItem, AdjustTheTip: CreateSliderMenuItem, AdjustTheTip: Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 117 |
| Global usage count | 117 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
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
WindowRegisterCoreEventHandler(windowName, windowEvent, handlerName)
```

## Description

Observed binding On* window events directly to a named window.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "ActionFractionWindowContextAutoHide", "ActionFractionWindowContextColorCodeCurrentAP", "ApothecaryWindowTitleBarText" |
| windowEvent | Observed as an On* window event string. | Observed values: "OnHidden", "OnInitializeCustomSettings", "OnKeyTab" |
| handlerName | Observed as a Lua handler reference. | Observed values: "ActionFractionWindow.ToggleAutoHide", "ActionFractionWindow.ToggleColorCodeCurrentAP", "AdjustTheTip.OnMouseOverTargetWindowClick" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- AdjustTheTip
- Amethyst
- BetterCC
- BuddyBind
- CMap
- Calling
- CharacterScreenTabFix
- CleanUnitFrames
- Crusher
- DuffTimer
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- GCDsaver
- Hopper
- InfoScroller
- LibRange
- LibWBToggler
- MacroIcons
- Map
- MarkBuff
- Minmap
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RVMOD_3DPortrait
- RVMOD_PlayerStatus
- RVMOD_Targets
- RealmStatus
- Rolodex
- Shinies
- TacticSetNames
- TargetRing
- TastyButtons
- TidyChat
- TidyQueue
- TidyRoll
- Tokens
- VPBreakdown
- Vectors
- WarBoard_Menu
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- bigger_MacroWindow
- nRarity
- scenarioInfo
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: RegisterEvent -> WindowRegisterCoreEventHandler(self.name, e, "LIBGUI_ELEMENT.events."..e.."."..self.name)
- ActionBarHide: RegisterEvent -> WindowRegisterCoreEventHandler(self.name, e, "LIBGUI_ELEMENT.events."..e.."."..self.name)
- ActionFraction: Initialize -> WindowRegisterCoreEventHandler("ActionFractionWindowContextColorCodeCurrentAP", "OnLButtonUp", "ActionFractionWindow.ToggleColorCodeCurrentAP")
- ActionFraction: Initialize -> WindowRegisterCoreEventHandler("ActionFractionWindowContextAutoHide", "OnLButtonUp", "ActionFractionWindow.ToggleAutoHide")
- AdjustTheTip: CreateCheckBoxMenuItem -> WindowRegisterCoreEventHandler(itemData.Name, "OnLButtonUp", itemData.Callback)
- AdjustTheTip: CreateSliderMenuItem -> WindowRegisterCoreEventHandler(itemData.Name, "OnLButtonUp", itemData.Callback)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [ButtonSetTextureSlice](window_ButtonSetTextureSlice.md) (HIGH 100/100) - Window Function
- [AlertTextWindow.AddAlert](../../globals/functions/global_AlertTextWindow.AddAlert.md) (HIGH 75/100) - Global Function

## Affects

- [SystemData.ChatLogFilters.SHOUT](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field

## Notes

- none

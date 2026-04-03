# WindowUnregisterCoreEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 44 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, ActionBarHide, Amethyst, BetterCC, BuddyBind, CMap, Calling, Crusher |
| Files seen in | BetterCC.lua, BuddyBind.lua, CallingKeybinding.lua, Core.lua, LibGUI.lua, LibGui.lua, LibRange.lua, Libraries/LibGUI.lua |
| Namespaces detected | WindowUnregisterCoreEventHandler |
| Source kinds | lua_calls |
| Example locations | Ace: UnregisterEvent, ActionBarHide: UnregisterEvent, Amethyst: UnregisterEvent, BetterCC: UpdatePulse, BuddyBind: OnExitBindingMode, BuddyBind: OnRawDeviceInput |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 59 |
| Global usage count | 59 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
WindowUnregisterCoreEventHandler(arg1, arg2)
```

## Description

Observed as a window function across 44 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "EA_Window_LootRoll", "EA_Window_ScenarioLobby", "KeyMappingWindowActionsList" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: "OnHidden", "OnKeyTab", "OnLButtonDown" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Removes a previously registered callback or command binding.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- BetterCC
- BuddyBind
- CMap
- Calling
- Crusher
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- GCDsaver
- Hopper
- InfoScroller
- LibRange
- LibWBToggler
- Map
- Minmap
- Motion
- NaturalLog
- Obsidian
- PartyCast
- Pure
- Pure Careerbar
- RVMOD_3DPortrait
- RVMOD_PlayerStatus
- RealmStatus
- Rolodex
- Shinies
- TargetRing
- TidyQueue
- TidyRoll
- Tokens
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- nRarity
- scenarioInfo
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- ActionBarHide: UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- Amethyst: UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- BetterCC: UpdatePulse -> WindowUnregisterCoreEventHandler(self, "OnUpdate")
- BuddyBind: OnExitBindingMode -> WindowUnregisterCoreEventHandler(BuddyBind.bindingButtonName, "OnRawDeviceInput")
- BuddyBind: OnRawDeviceInput -> WindowUnregisterCoreEventHandler(BuddyBind.bindingButtonName, "OnRawDeviceInput")

## Related APIs

- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Affects

- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.M_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.M_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.R_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.R_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- none

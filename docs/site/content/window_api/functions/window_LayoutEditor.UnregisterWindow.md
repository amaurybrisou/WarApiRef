# LayoutEditor.UnregisterWindow

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 35 addons

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
| Addons seen in | Amethyst, Atlas, CDown, DAoCBuff, Effigy, HealGrid, Hopper, KillTracker |
| Files seen in | Amethyst.lua, CDown.lua, Core.lua, Display.lua, EffectTracker/EffectContainer.lua, Elements/EffigyBar.lua, GlobalCooldown.lua, KillerLifecycle.lua |
| Namespaces detected | LayoutEditor |
| Source kinds | lua_calls |
| Example locations | Amethyst: Recreate, Atlas: Shutdown, CDown: RestartTracker, DAoCBuff: Shutdown, Effigy: destroy, Effigy: setup |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 70 |
| Global usage count | 70 |
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
LayoutEditor.UnregisterWindow(arg1)
```

## Description

Observed as a window function across 35 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: "AtlasFrame", "CDownWindow", "KillerPersonalCounter" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Removes a previously registered callback or command binding.

## Seen In

- Amethyst
- Atlas
- CDown
- DAoCBuff
- Effigy
- HealGrid
- Hopper
- KillTracker
- Killer
- LibAddonButton
- Map
- Miracle Grow Remix
- NerfedButtons
- Obsidian
- Phantom
- Pure
- QuickWarReport
- RVMOD_Manager
- RVMOD_PlayerStatus
- SNT_CASTBAR
- SNT_PANEL
- SOR
- Squared
- TastyButtons
- TexturedButtons
- TidyRoll
- TurretRange
- Twister
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- emotes
- scenarioInfo
- xHUD
- xPanels

## Examples

- Amethyst: Recreate -> LayoutEditor.UnregisterWindow(C.name)
- Atlas: Shutdown -> LayoutEditor.UnregisterWindow("AtlasFrame")
- CDown: RestartTracker -> LayoutEditor.UnregisterWindow("CDownWindow")
- DAoCBuff: Shutdown -> LayoutEditor.UnregisterWindow(self.m_windowName)
- Effigy: destroy -> LayoutEditor.UnregisterWindow(self.name)
- Effigy: setup -> LayoutEditor.UnregisterWindow(self.name)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Used With

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Notes

- none

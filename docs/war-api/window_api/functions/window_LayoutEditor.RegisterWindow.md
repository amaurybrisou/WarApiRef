# LayoutEditor.RegisterWindow

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 23 addons

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
| Addons seen in | BuffHead, Busted, CM_ClosetGoblin, DAoCBuff, Effigy, Enemy, GuardLine, Killer |
| Files seen in | `/workspace_addons/BuffHead/AdvancedContainers.lua:36`, `/workspace_addons/Busted/Busted.lua:333`, `/workspace_addons/ClosetGoblin/ClosetGoblinOptionWindow.lua:28`, `/workspace_addons/DAoCBuff/Source/DAoCBuffFrames.lua:346`, `/workspace_addons/DAoCBuff/Source/DAoCBuffHeadFrames.lua:122`, `/workspace_addons/Effigy/Elements/EffigyBar.lua:58`, `/workspace_addons/Enemy/Code/Assist/Assist.lua:7`, `/workspace_addons/Enemy/Code/CombatLog/CombatLogEpsWindow.lua:12` |
| Namespaces detected | LayoutEditor |
| Source kinds | lua_calls |
| Example locations | BuffHead: BuffHead.local.RegisterLayoutEditor, BuffHead: RegisterLayoutEditor, Busted: BustedGUI.Initialize, CM_ClosetGoblin: ClosetGoblinOptionWindow.OnInitialize, DAoCBuff: DAoCBuffHeadTracker:Create, DAoCBuff: DAoCBuffTracker:Create |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 43 |
| Global usage count | 43 |
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

Observed as a window function across 23 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "BustedMiniGUI", "ClosetGoblinOptionWindow", "EnemyCombatLogEpsWindow" |
| arg2 | Observed as a text or wstring payload. | Observed values: L "", L "Busted", L "Chat Text Entry" |
| arg3 | Observed as a text or wstring payload. | Observed values: L "", L "Alternative Buffwindow. (Head)", L "Alternative Buffwindow." |
| arg4 | Observed as a boolean toggle. | Observed values: false, true |
| arg5 | Observed as a boolean toggle. | Observed values: false, true |
| arg6 | Observed as a boolean toggle. | Observed values: false, true |
| arg7 | Observed as a runtime window or control identifier. | Observed values: nil |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- BuffHead
- Busted
- CM_ClosetGoblin
- DAoCBuff
- Effigy
- Enemy
- GuardLine
- Killer
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Moth
- QuickWarReport
- RVMOD_Manager
- RoR_SoR
- Swift Assist
- Targets
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WhoHealedMe
- followTheLeader

## Examples

- BuffHead: BuffHead.local.RegisterLayoutEditor -> LayoutEditor.RegisterWindow(layoutWindow, displayName, L "", false, false, true)
- BuffHead: RegisterLayoutEditor -> LayoutEditor.RegisterWindow(layoutWindow, displayName, L "", false, false, true)
- Busted: BustedGUI.Initialize -> LayoutEditor.RegisterWindow("BustedMiniGUI", L "Busted", L "Busted Mini Alert", false, false, true, nil)
- CM_ClosetGoblin: ClosetGoblinOptionWindow.OnInitialize -> LayoutEditor.RegisterWindow("ClosetGoblinOptionWindow", L "Closet Goblin", L "Option Button", true, true, true, nil)
- DAoCBuff: DAoCBuffHeadTracker:Create -> LayoutEditor.RegisterWindow(windowName, towstring(windowName), L "Alternative Buffwindow. (Head)", false, false, true, nil)
- DAoCBuff: DAoCBuffTracker:Create -> LayoutEditor.RegisterWindow(windowName, towstring(windowName), L "Alternative Buffwindow.", false, false, true, nil)

## Related APIs

- none

## Used With

- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function

## Triggered By

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

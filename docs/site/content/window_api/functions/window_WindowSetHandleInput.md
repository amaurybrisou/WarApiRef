# WindowSetHandleInput

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 13 addons

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
| Addons seen in | Ace, BuffHead, DAoCBuff, Effigy, Enemy, GCDsaver, LibWBToggler, Miracle Grow Remix |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:104`, `/workspace_addons/Ace/LibGUI.lua:89`, `/workspace_addons/Ace/LibGUI.lua:94`, `/workspace_addons/BuffHead/EffectContainer.lua:75`, `/workspace_addons/BuffHead/EffectFrame.lua:309`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings.lua:112`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings.lua:161`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings.lua:169` |
| Namespaces detected | WindowSetHandleInput |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:CaptureInput, Ace: LIBGUI_ELEMENT:IgnoreInput, Ace: LIBGUI_ELEMENT:MakeMovable, BuffHead: BuffHeadEffectContainer:Create, BuffHead: BuffHeadEffectFrame:Create, DAoCBuff: DAoCBuffSettings.CreateOptionswindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 69 |
| Global usage count | 69 |
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
WindowSetHandleInput(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "DAoCBuff_Settings", "EnemyGuardDistanceIndicator", "EnemyTarget" |
| arg2 | Observed as a boolean toggle. | Observed values: (handleInput==true), chatwindow_tabs_handle_input~=false, false |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- BuffHead
- DAoCBuff
- Effigy
- Enemy
- GCDsaver
- LibWBToggler
- Miracle Grow Remix
- Shinies
- TidyChat
- TidyRoll
- WoH-Reticle
- nRarity

## Examples

- Ace: LIBGUI_ELEMENT:CaptureInput -> WindowSetHandleInput(self.name, true)
- Ace: LIBGUI_ELEMENT:IgnoreInput -> WindowSetHandleInput(self.name, false)
- Ace: LIBGUI_ELEMENT:MakeMovable -> WindowSetHandleInput(self.name, true)
- BuffHead: BuffHeadEffectContainer:Create -> WindowSetHandleInput(frame:GetName(), true)
- BuffHead: BuffHeadEffectFrame:Create -> WindowSetHandleInput(frameName, (handleInput==true))
- BuffHead: BuffHeadEffectFrame:Create -> WindowSetHandleInput(frame:GetName(), (handleInput==true))

## Related APIs

- none

## Used With

- [WindowGetHandleInput](window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetMovable](window_WindowSetMovable.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

# WindowGetScale

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
| Addons seen in | Ace, BuffHead, DAoCBuff, Enemy, GuardLine, LibWBToggler, PartyCast, PotionBar |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:236`, `/workspace/data/raw/BuffHead/AdvancedContainers.lua:36`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:480`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffFrames.lua:653`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffFrames.lua:72`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffHeadFrames.lua:212`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffHeadFrames.lua:796`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffHeadFrames.lua:850` |
| Namespaces detected | WindowGetScale |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Scale, BuffHead: BuffHead.local.RegisterLayoutEditor, BuffHead: RegisterLayoutEditor, DAoCBuff: DAoCBuff.SetSize, DAoCBuff: DAoCBuffFrame:SetScale, DAoCBuff: DAoCBuffHeadTracker:OnBuffsChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 46 |
| Global usage count | 46 |
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
WindowGetScale(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "GuardLineSelfWindowCircle", "PartyCastStaticWindow"..PlayerNumber, "PartyCastStaticWindow0" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- BuffHead
- DAoCBuff
- Enemy
- GuardLine
- LibWBToggler
- PartyCast
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TurretRange
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:Scale -> WindowGetScale(self.name)
- BuffHead: BuffHead.local.RegisterLayoutEditor -> WindowGetScale(container:GetName())
- BuffHead: RegisterLayoutEditor -> WindowGetScale(container:GetName())
- DAoCBuff: DAoCBuff.SetSize -> WindowGetScale(sourceWnd)
- DAoCBuff: DAoCBuffFrame:SetScale -> WindowGetScale(self.m_name)
- DAoCBuff: DAoCBuffFrame:SetScale -> WindowGetScale(self.m_parentname)

## Related APIs

- none

## Used With

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

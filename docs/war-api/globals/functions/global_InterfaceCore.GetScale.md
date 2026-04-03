# InterfaceCore.GetScale

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 13 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 168

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, BuffHead, DAoCBuff, Enemy, GuardLine, LibWBToggler, PartyCast, PotionBar |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:236`, `/workspace/data/raw/BuffHead/AdvancedContainers.lua:275`, `/workspace/data/raw/BuffHead/Container.lua:758`, `/workspace/data/raw/BuffHead/EffectFrame.lua:52`, `/workspace/data/raw/BuffHead/Setup/LayoutControlFrame.lua:8`, `/workspace/data/raw/BuffHead/Setup/LayoutFrame.lua:105`, `/workspace/data/raw/BuffHead/Setup/LayoutFrame.lua:3`, `/workspace/data/raw/BuffHead/Setup/LayoutFrame.lua:71` |
| Namespaces detected | InterfaceCore |
| Source kinds | globals, lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Scale, BuffHead: AutoSize, BuffHead: BuffHead.AdvancedContainers.OnLayoutEditorFinished, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.Show, BuffHead: BuffHead.Setup.Layout.Properties.Show, BuffHead: BuffHead.Setup.LayoutFrame:Create |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 40 |
| Global usage count | 40 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
InterfaceCore.GetScale()
```

## Description

Observed as a global function across 13 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

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

- Ace: LIBGUI_ELEMENT:Scale -> InterfaceCore.GetScale()
- BuffHead: AutoSize -> InterfaceCore.GetScale()
- BuffHead: BuffHead.AdvancedContainers.OnLayoutEditorFinished -> InterfaceCore.GetScale()
- BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.Show -> InterfaceCore.GetScale()
- BuffHead: BuffHead.Setup.Layout.Properties.Show -> InterfaceCore.GetScale()
- BuffHead: BuffHead.Setup.LayoutFrame:Create -> InterfaceCore.GetScale()

## Related APIs

- [InterfaceCore.GetResolutionScale](global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [LabelSetTextAlign](../../window_api/functions/window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetWordWrap](../../window_api/functions/window_LabelSetWordWrap.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [WindowGetAnchor](../../window_api/functions/window_WindowGetAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function

## Used With

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [LabelSetFont](../../window_api/functions/window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](../../window_api/functions/window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.

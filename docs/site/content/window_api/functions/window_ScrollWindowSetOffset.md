# ScrollWindowSetOffset

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | Enemy, RVMOD_Manager, Shinies, wbLeadHelper |
| Files seen in | `/workspace/Enemy/Code/Core/Main.lua:300`, `/workspace/Enemy/Code/UnitFrames/ClickCasting.lua:202`, `/workspace/Enemy/Code/UnitFrames/EffectsIndicator.lua:613`, `/workspace/Enemy/Code/UnitFrames/UnitFramePart.lua:701`, `/workspace/Enemy/Code/UnitFrames/UnitFrames.lua:637`, `/workspace/RVMOD_Manager/RVMOD_Manager.lua:526`, `/workspace/RVMOD_Manager/RVMOD_Manager.lua:709`, `/workspace/RVMOD_Manager/RVMOD_Manager.lua:752` |
| Namespaces detected | ScrollWindowSetOffset |
| Source kinds | lua_calls |
| Example locations | Enemy: Enemy.UI_ConfigDialog_Open, Enemy: Enemy.UnitFramesUI_ConfigDialog_Import, Enemy: Enemy.UnitFramesUI_EffectsIndicatorDialog_Open, Enemy: Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open, Enemy: Enemy.UnitFramesUI_UnitFramePartDialog_Open, RVMOD_Manager: RVMOD_Manager.FilterAddonsList |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
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
ScrollWindowSetOffset(arg1, arg2)
```

## Description

Observed as a window function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "EnemyClickCastingDialogContent", "EnemyConfigDialogSectionsUnitFramesContent", "EnemyEffectsIndicatorDialogContent" |
| arg2 | Observed as a numeric value. | Observed values: 0, scroll |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Enemy
- RVMOD_Manager
- Shinies
- wbLeadHelper

## Examples

- Enemy: Enemy.UI_ConfigDialog_Open -> ScrollWindowSetOffset(config_dlg.currentSection.windowName.."Content", scroll)
- Enemy: Enemy.UnitFramesUI_ConfigDialog_Import -> ScrollWindowSetOffset("EnemyConfigDialogSectionsUnitFramesContent", 0)
- Enemy: Enemy.UnitFramesUI_EffectsIndicatorDialog_Open -> ScrollWindowSetOffset("EnemyEffectsIndicatorDialogContent", 0)
- Enemy: Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open -> ScrollWindowSetOffset("EnemyClickCastingDialogContent", 0)
- Enemy: Enemy.UnitFramesUI_UnitFramePartDialog_Open -> ScrollWindowSetOffset("EnemyUnitFramePartDialogContent", 0)
- RVMOD_Manager: RVMOD_Manager.FilterAddonsList -> ScrollWindowSetOffset(WindowManager.."ContentListBox", 0)

## Related APIs

- none

## Used With

- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [ScrollWindowUpdateScrollRect](window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

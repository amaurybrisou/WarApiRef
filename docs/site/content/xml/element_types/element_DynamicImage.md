# DynamicImage

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Aura, AutoMark, BagOMatic, BankArkel, BuffHead |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.xml:13`, `/workspace_addons/AggroMeter/AggroMeter.xml:193`, `/workspace_addons/AggroMeter/AggroMeter.xml:247`, `/workspace_addons/AggroMeter/AggroMeter.xml:276`, `/workspace_addons/AggroMeter/AggroMeter.xml:304`, `/workspace_addons/AggroMeter/AggroMeter.xml:332`, `/workspace_addons/AggroMeter/AggroMeter.xml:361`, `/workspace_addons/AggroMeter/AggroMeter.xml:64` |
| Namespaces detected | DynamicImage |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedRenownTrainer: AbilityButtonTemplateSquare, AdvancedRenownTrainer: AbilityButtonTemplateSquareFrame, AdvancedRenownTrainer: AdvancedRenownTrainingWindowCornorImage, AggroMeter: AggroMeterWindow_AggroWindow1Tactic, AggroMeter: AggroMeterWindow_AggroWindow2Tactic, AggroMeter: AggroMeterWindow_AggroWindow3Tactic |
| XML usage count | 321 |
| XML attribute usage count | 321 |
| Lua usage count | 5 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
| Consistent role | yes |
| Consistent arguments | no |
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

## Description

Observed XML element type instantiated by 41 addons.

## Common Attributes

- name
- handleinput
- texture
- layer
- textureScale
- popable
- slice
- id
- texturescale
- inherits
- sticky
- alpha

## Common Handlers

- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)

## Common Handler Functions

- Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip
- BuffHead.Setup.AdvancedContainersItem.OnContainerRClick
- BuffHead.Setup.Container.OnContainerClick
- ClosetGoblinOptionWindow.OnLButtonUp
- ClosetGoblinOptionWindow.OnRButtonUp
- MapPin.ShowIcons
- AggroMeter.OnMouseOverStart
- BagOMatic.wnd_on_lbutton_up
- BagOMatic.wnd_on_mouse_over
- CMapWindow.MouseoverMail
- CMapWindow.OnMouseoverRvRIndicator
- Enemy.CombatLogUI_IDS_OnRowLButtonDown


## XML Event Bindings

| Event | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|---------------------|-------------------|-----------------|
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | Enemy.CombatLogUI_IDS_OnRowLButtonDown | `function(...)` | LOW |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | BuffHead.Setup.Container.OnContainerClick, ClosetGoblinOptionWindow.OnLButtonUp, MapPin.ShowIcons, BagOMatic.wnd_on_lbutton_up, Enemy.Guard_GuardIndicator_OnLButtonUp, Enemy.MarksUI_EnemyMarksWindow_OnAddLButtonUp | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip, AggroMeter.OnMouseOverStart, BagOMatic.wnd_on_mouse_over, CMapWindow.MouseoverMail, CMapWindow.OnMouseoverRvRIndicator, Enemy.Guard_GuardIndicator_OnMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOverEnd | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | BuffHead.Setup.AdvancedContainersItem.OnContainerRClick, ClosetGoblinOptionWindow.OnRButtonUp, Enemy.MarksUI_EnemyMarksWindow_OnAddRButtonUp, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick, Enemy.UI_Icon_OnRButtonUp, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionRButtonUp | `function(...)` | LOW |

## Common Inherits

- EA_ListSortDownArrow
- EA_ListSortUpArrow
- Aggro_Tactic_Template
- MapMonsterIconChooserWindow_IconTemplate
- EA_Image_DefaultIcon
- EA_Image_DefaultIconFrame
- GenericEndCap
- RVMOD_ManagerStatusTemplate
- EA_Default_CharacterImage
- EA_Default_CornerImage
- EA_Default_MerchantImage
- EA_Default_TrainingImage

## Common Parent Elements

- [Window](element_Window.md)
- [Button](element_Button.md)
- [DynamicImage](element_DynamicImage.md)
- [Label](element_Label.md)

## Common Named Child Elements

- [DynamicImage](element_DynamicImage.md)
- [Label](element_Label.md)


## Structural Sub-Elements

### [TintColor](element_TintColor.md)

- Observed in 27 parent frames
- Attributes: `a`, `b`, `g`, `r`
  - `a`: `0.5`, `0.8`, `0.7`, `1`
  - `b`: `0`, `50`, `255`, `155`, `40`, `175`
  - `g`: `0`, `50`, `155`, `255`, `90`, `40`, `175`
  - `r`: `255`, `50`, `155`, `90`, `0`, `40`, `175`


## Attribute Reference

| Attribute | Role | Observed Values | Count |
|-----------|------|-----------------|-------|
| `name` | string | `AbilityButtonTemplateSquare`, `AbilityButtonTemplateSquareFrame`, `AdvancedRenownTrainingWindowCornorImage`, `AggroMeterWindow_AggroWindow1Tactic`, … | 321 |
| `handleinput` | boolean | `false`, `true` | 243 |
| `texture` | string | `EA_SquareFrame`, `AggroMeterIcon`, `EA_Abilities01_d5`, `EA_HUD_01`, … | 236 |
| `layer` | string | `background`, `overlay`, `popup`, `default`, … | 202 |
| `textureScale` | number | `1.0`, `0.5`, `1.171`, `1`, … | 80 |
| `popable` | boolean | `false`, `true` | 55 |
| `slice` | string | `Tab-ALL`, `Round-Swatch-Selection-Ring`, `Radio-Button`, `You-Have-Mail`, … | 55 |
| `id` | number | `1`, `2`, `3`, `4`, … | 48 |
| `texturescale` | number | `0.35`, `1.00`, `.5`, `0.6`, … | 46 |
| `inherits` | frame-ref | `EA_Default_TrainingImage`, `Aggro_Tactic_Template`, `EA_ListSortDownArrow`, `EA_ListSortUpArrow`, … | 44 |
| `sticky` | boolean | `false`, `true` | 27 |
| `alpha` | number | `1.0`, `0.0`, `0.5`, `0.9`, … | 16 |
| `savesettings` | boolean | `false`, `true` | 8 |
| `poppable` | boolean | `false` | 6 |
| `filtering` | boolean | `true` | 5 |

## Seen In

- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Aura
- AutoMark
- BagOMatic
- BankArkel
- BuffHead
- CM_ClosetGoblin
- CMap
- CombatTextNames
- DAoCBuff
- EA_ThreePartBar
- EA_UiModWindow
- Effigy
- Enemy
- GetStats
- GuardLine
- JunkDump
- Killer
- LibWBToggler
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Moth
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- TidyRoll
- TurretRange
- WSCT

## Examples

- AdvancedRenownTrainer: AbilityButtonTemplateSquare -> DynamicImage AbilityButtonTemplateSquare
- AdvancedRenownTrainer: AbilityButtonTemplateSquareFrame -> DynamicImage AbilityButtonTemplateSquareFrame
- AdvancedRenownTrainer: AdvancedRenownTrainingWindowCornorImage -> DynamicImage AdvancedRenownTrainingWindowCornorImage
- AggroMeter: AggroMeterWindow_AggroWindow1Tactic -> DynamicImage AggroMeterWindow_AggroWindow1Tactic
- AggroMeter: AggroMeterWindow_AggroWindow2Tactic -> DynamicImage AggroMeterWindow_AggroWindow2Tactic
- AggroMeter: AggroMeterWindow_AggroWindow3Tactic -> DynamicImage AggroMeterWindow_AggroWindow3Tactic

## Related APIs

- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function

## Used With

- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none

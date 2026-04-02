# Label

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, BankArkel, BuffHead, Busted, CM_ClosetGoblin |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:1012`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1034`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1056`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1068`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1090`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1102`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1114`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1126` |
| Namespaces detected | Label |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDLabel, AdvancedPetAssist: APAInstantOnlyHUDLabel, AdvancedPetAssist: APAKitingHUDLabel, AdvancedPetAssist: APALabelAttackBind, AdvancedPetAssist: APALabelAutoReattack, AdvancedPetAssist: APALabelAutoReattackDelay |
| XML usage count | 1572 |
| XML attribute usage count | 1572 |
| Lua usage count | 7 |
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

Observed XML element type instantiated by 56 addons.

## Common Attributes

- name
- textalign
- font
- inherits
- handleinput
- autoresize
- warnOnTextCropped
- maxchars
- wordwrap
- layer
- popable
- skipinput

## Common Handlers

- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md)
- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)

## Common Handler Functions

- LPET.OnMouseOver
- Cheeseboard.SortColumnClick
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip
- Cheeseboard.OnMouseOver
- FrameManager.OnMouseOverEnd
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick
- Enemy.CombatLogUI_EpsWindow_OnMouseOver
- Enemy.CombatLogUI_EpsWindow_OnRButtonUp
- EA_ChatWindow.OnHyperLinkLButtonUp
- EA_ChatWindow.OnHyperLinkRButtonUp
- Enemy.ConfigurationWindow_ShowTooltip


## XML Event Bindings

| Event | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|---------------------|-------------------|-----------------|
| [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) | EA_ChatWindow.OnHyperLinkLButtonUp | `function(...)` | LOW |
| [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) | EA_ChatWindow.OnHyperLinkRButtonUp, MapPin.RButtonUp | `function(...)` | LOW |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | Enemy.CombatLogUI_IDS_OnRowLButtonDown | `function(...)` | LOW |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | Cheeseboard.SortColumnClick, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick, Enemy.KillSpamUI_KillSpamDialog_OnTotalLButtonUp, ClosetGoblinCharacterWindow.OnClickSetRow, ClosetGoblinZoneWindow.OnClickZoneRow, Enemy.CombatLogUI_EpsWindow_OnInDpsLButtonUp | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | LPET.OnMouseOver, Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip, Cheeseboard.OnMouseOver, Enemy.CombatLogUI_EpsWindow_OnMouseOver, Enemy.ConfigurationWindow_ShowTooltip, Enemy.CombatLogUI_StatsWindow_ListRowMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | FrameManager.OnMouseOverEnd, TexturedButtons.Setup.Fonts.OnFontExampleMouseOut, BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOut, BuffHead.Setup.Layout.Properties.OnFontExampleMouseOut, BuffHead.Setup.Layout.Properties.OnPropertyTitleMouseOut, TurretRange.Setup.Display.OnDistanceFontMouseOut | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick, Enemy.CombatLogUI_EpsWindow_OnRButtonUp, Enemy.CombatLogUI_TargetDefenseWindow_OnRButtonUp, ClosetGoblinCharacterWindow.OnSetRowContextMenu, ClosetGoblinZoneWindow.OnSetZoneRowContextMenu, MiracleGrowLight.switchMode | `function(...)` | LOW |

## Common Inherits

- EA_Label_DefaultText
- Aura_Default_Label_SmallFont
- ClosetGoblinTalismanLabel
- Shinies_Default_Label_ClearSmallFont
- ModDetailsTextDef
- ModDetailsLabelDef
- DefaultWindowSmallText
- EA_Settings_ItemTitle
- MapMonsterEditorWindowLabelDefault
- DefaultWindowText
- EA_Label_DefaultText_Small
- IraConfigHeading

## Common Parent Elements

- [Window](element_Window.md)
- [Button](element_Button.md)
- [SliderBar](element_SliderBar.md)
- [DynamicImage](element_DynamicImage.md)
- [FullResizeImage](element_FullResizeImage.md)
- [EditBox](element_EditBox.md)

## Common Named Child Elements

- [DynamicImage](element_DynamicImage.md)

## Common Structural Child Elements

- [Text](element_Text.md)
- [TintColor](element_TintColor.md)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `textalign` | optional | 74% | left, center, topleft, right, ... |
| `font` | optional | 71% | font_default_text, font_chat_text, font_clear_tiny, font_clear_small, ... |
| `inherits` | optional | 31% | MapMonsterEditorWindowLabelDefault, EA_Label_DefaultText, ModDetailsLabelDef, Aura_Default_Label_SmallFont, ... |
| `handleinput` | optional | 27% | false, true |
| `autoresize` | optional | 23% | true, false |
| `warnOnTextCropped` | optional | 21% | false |
| `maxchars` | optional | 21% | 256, 35, 200, 42, ... |
| `wordwrap` | optional | 16% | true, false |
| `layer` | optional | 16% | default, overlay, background, secondary, ... |
| `popable` | optional | 6% | false |
| `skipinput` | optional | 4% | true |
| `autoresizewidth` | optional | 3% | true, false |
| `textAutoFitMinScale` | optional | 3% | 0.3, 0.5, 0.75 |
| `hanldeinput` | optional | 1% | false |
| `sticky` | optional | 1% | true, false |
| `alpha` | optional | 1% | 1, 0, .1 |
| `show` | optional | 1% | false |
| `handlinput` | optional | 0% | false |
| `text` | optional | 0% |  |
| `autoresizeheight` | optional | 0% | true |
| `ellipsisOnCrop` | optional | 0% | true |
| `ignoreFormattingTags` | optional | 0% | true |
| `linespacing` | optional | 0% | 13, 12 |
| `showing` | optional | 0% | false |
| `background` | optional | 0% | EA_FullResizeImage_TanBorder |
| `id` | optional | 0% | 1 |
## Structural Sub-Elements

### [Text](element_Text.md)

Observed 69 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `text` | optional |  |
| `alignment` | optional |  |
### [TintColor](element_TintColor.md)

Observed 14 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** |  |
| `g` | **required** |  |
| `r` | **required** |  |
| `a` | optional |  |
## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Call Count | From Events |
| --- | --- | --- |
| `WindowGetId` | 30 | OnHyperLinkRButtonUp, OnLButtonUp, OnMouseOver, OnRButtonUp |
| `WindowSetShowing` | 18 | OnHyperLinkLButtonUp, OnHyperLinkRButtonUp, OnLButtonUp |
| `LabelSetTextColor` | 12 | OnMouseOver, OnMouseOverEnd |
| `TextEditBoxSetText` | 12 | OnHyperLinkRButtonUp |
| `SystemData.ActiveWindow.name:match` | 8 | OnLButtonDown, OnMouseOver |
| `CreateWindow` | 6 | OnMouseOver |
| `LabelSetText` | 6 | OnMouseOver |
| `SystemData.MouseOverWindow.name:match` | 1 | OnRButtonUp |
| `WindowGetDimensions` | 1 | OnMouseOver |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
## Lua Functions Manipulating This Type

- BankArkel.BankArkel.PackHide
- TidyRoll.TidyRollOptions.Initialize
- EA_UiDebugTools.BustedGUI.ClearAlertFlash
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- Killer.Killer.ShowRowTooltip
- Killer.Killer.ShowTopKillersTooltip
- Killer.Killer.Initialize
- Busted.BustedGUI.ClearAlertFlash
- Enemy.Enemy.CombatLogUI_StatsWindow_UpdateList
- Enemy.Enemy.CombatLogUI_EpsWindow_UpdateLayout
- BankArkel.BankArkel.SetCharInfo
- Killer.Killer.ShowPersonalStatsTooltip
- Enemy.Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- BankArkel
- BuffHead
- Busted
- CM_ClosetGoblin
- CMap
- Cheeseboard
- CombatTextNames
- Crafting Info Tooltip
- DAoCBuff
- DaemonAssist
- EA_UiDebugTools
- EA_UiModWindow
- Effigy
- Enemy
- FastInteract
- GetStats
- GuardLine
- GuardList
- GuardRange
- JunkDump
- Killer
- LibGroup
- LibWBToggler
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- Moth
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- TidyChat
- TidyRoll
- Tortall_DPS
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAFollowTargetHUDLabel -> Label APAFollowTargetHUDLabel
- AdvancedPetAssist: APAInstantOnlyHUDLabel -> Label APAInstantOnlyHUDLabel
- AdvancedPetAssist: APAKitingHUDLabel -> Label APAKitingHUDLabel
- AdvancedPetAssist: APALabelAttackBind -> Label APALabelAttackBind
- AdvancedPetAssist: APALabelAutoReattack -> Label APALabelAutoReattack
- AdvancedPetAssist: APALabelAutoReattackDelay -> Label APALabelAutoReattackDelay

## Related APIs

- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function

## Used With

- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnHyperLinkLButtonUp](../../events/window_events/window_event_OnHyperLinkLButtonUp.md) (HIGH 100/100) - Window Event
- [OnHyperLinkRButtonUp](../../events/window_events/window_event_OnHyperLinkRButtonUp.md) (HIGH 100/100) - Window Event
- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOverEnd](../../events/window_events/window_event_OnMouseOverEnd.md) (HIGH 100/100) - Window Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none

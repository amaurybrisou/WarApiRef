# OverlayTexCoords

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, Aura, Crusher, DaemonAssist, EA_UiDebugTools, GDes, Ges, Hopper |
| Files seen in | Configuration/Config.xml, Configuration/HopperConfig.xml, Configuration/WCDBConfig.xml, Configuration/WCDPConfig.xml, Source/DebugWindowVerticalScroll.xml, Source/PureUIElementTemplates.xml, Source/ShiniesUITemplates.xml, Source/SocialWindow.xml |
| Namespaces detected | OverlayTexCoords |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APA_ComboBoxButton, AdvancedPetAssist: APA_ComboBoxButtonWide, Aura: Aura_Button_DefaultResizableTinyComboBoxSelected, Aura: Aura_ComboBox_DefaultResizableComboBoxSelected, Aura: Aura_ComboBox_DefaultResizableComboBoxSelectedLarge, Crusher: Crusher_Button_DefaultResizableComboBoxSelected |
| XML usage count | 24 |
| XML attribute usage count | 24 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

OverlayTexCoords is a structural XML sub-element. It commonly appears under Button. It is typically used to organize structural children such as Disabled, Normal, NormalHighlit.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 24× (HIGH)

## Common Structural Child Elements

- [Normal](element_Normal.md) — 24× (HIGH)
- [NormalHighlit](element_NormalHighlit.md) — 24× (HIGH)
- [Pressed](element_Pressed.md) — 24× (HIGH)
- [PressedHighlit](element_PressedHighlit.md) — 24× (HIGH)
- [Disabled](element_Disabled.md) — 20× (HIGH)

## Structural Sub-Elements

### [Normal](element_Normal.md)

Observed 24 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-yellow, RightTabFrame, Tactics-Button, SquareButton |
| `b` | optional | 255, 0, 102, 73 |
| `g` | optional | 255, 0, 204, 175 |
| `r` | optional | 255, 155, 222, 226 |
| `x` | optional | 0, 92, 172, 494 |
| `y` | optional | 28, 44, 0, 341 |
| `a` | optional | 255, 1 |
| `texture` | optional | bpKtxt, EA_SquareFrame, CrusherIconBorderNormal, EA_RoundFrame |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_AnimatedImage_DefaultChoiceOverlay, BuffHeadLayoutVerticalButtonNormal, BuffHeadLayoutHorizontalButtonNormal |
### [NormalHighlit](element_NormalHighlit.md)

Observed 24 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-white, RightTabFrame-Rollover, Tactics-Button-Rollover, SquareButton-Rollover |
| `b` | optional | 63, 36, 0, 50 |
| `g` | optional | 213, 57, 85, 192 |
| `r` | optional | 250, 95, 255, 222 |
| `a` | optional | 255, 1 |
| `x` | optional | 27, 105, 0, 201 |
| `y` | optional | 28, 44, 0, 341 |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_AnimatedImage_DefaultChoiceOverlay, EA_FullResizeImage_RedTransparent, EA_Button_GuildRosterRowHighlight |
| `texture` | optional | EA_SquareFrame_Highlight, CrusherIconBorderHighlight, EA_RoundFrame_Pressed, PinBG |
### [Pressed](element_Pressed.md)

Observed 24 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-white, RightTabFrame-Rollover, Tactics-Button-Depressed, SquareButton-Depressed |
| `b` | optional | 63, 36, 102, 235 |
| `g` | optional | 213, 57, 204, 235 |
| `r` | optional | 250, 95, 255, 235 |
| `a` | optional | 255, 1 |
| `x` | optional | 0, 120, 172, 494 |
| `y` | optional | 56, 44, 0, 370 |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_Button_GuildRosterRowPressed, BuffHeadLayoutVerticalButtonPressed, BuffHeadLayoutHorizontalButtonPressed |
| `texture` | optional | EA_SquareFrame_Pressed, CrusherIconBorderHighlight, EA_RoundFrame_Highlight, PinBG |
### [PressedHighlit](element_PressedHighlit.md)

Observed 24 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-white, RightTabFrame-Rollover, Tactics-Button-Depressed, SquareButton-Depressed |
| `b` | optional | 63, 0, 36, 45 |
| `g` | optional | 213, 85, 57, 255 |
| `r` | optional | 250, 255, 95, 226 |
| `a` | optional | 255, 1 |
| `x` | optional | 0, 120, 172, 475 |
| `y` | optional | 56, 44, 370, 420 |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_Button_GuildRosterRowPressed, BuffHeadLayoutVerticalButtonPressed, BuffHeadLayoutHorizontalButtonPressed |
### [Disabled](element_Disabled.md)

Observed 20 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | optional | 92, 36, 128, 102 |
| `g` | optional | 92, 57, 128, 204 |
| `r` | optional | 92, 95, 128, 255 |
| `a` | optional | 255, 0.5 |
| `x` | optional | 27, 92, 0, 230 |
| `y` | optional | 56, 44, 0, 341 |
| `id` | optional | morale-white, RightTabFrame, Tactics-Button-Disabled, Map-Plus-Button-Disabled |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, Crusher_HorizontalResizeImage_DefaultComboBox, EA_Button_TabDisabled, EA_HorizontalResizeImage_DefaultComboBox2 |
| `texture` | optional | EA_SquareFrame, CrusherIconBorderNormal, EA_RoundFrame, PinBG |
## Recursive Hierarchy

- Root: [OverlayTexCoords](element_OverlayTexCoords.md)
- [Disabled](element_Disabled.md) (structural, 20×, HIGH)
- [Normal](element_Normal.md) (structural, 24×, HIGH)
- [NormalHighlit](element_NormalHighlit.md) (structural, 24×, HIGH)
- [Pressed](element_Pressed.md) (structural, 24×, HIGH)
- [PressedHighlit](element_PressedHighlit.md) (structural, 24×, HIGH)

## Seen In

- AdvancedPetAssist
- Aura
- Crusher
- DaemonAssist
- EA_UiDebugTools
- GDes
- Ges
- Hopper
- LoyalPet
- Motion
- Pure
- Shinies
- SocialWindow 2.0
- WSCT
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse

## Examples

- AdvancedPetAssist: APA_ComboBoxButton -> OverlayTexCoords in Button APA_ComboBoxButton
- AdvancedPetAssist: APA_ComboBoxButtonWide -> OverlayTexCoords in Button APA_ComboBoxButtonWide
- Aura: Aura_Button_DefaultResizableTinyComboBoxSelected -> OverlayTexCoords in Button Aura_Button_DefaultResizableTinyComboBoxSelected
- Aura: Aura_ComboBox_DefaultResizableComboBoxSelected -> OverlayTexCoords in Button Aura_ComboBox_DefaultResizableComboBoxSelected
- Aura: Aura_ComboBox_DefaultResizableComboBoxSelectedLarge -> OverlayTexCoords in Button Aura_ComboBox_DefaultResizableComboBoxSelectedLarge
- Crusher: Crusher_Button_DefaultResizableComboBoxSelected -> OverlayTexCoords in Button Crusher_Button_DefaultResizableComboBoxSelected

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [Disabled](element_Disabled.md) (MEDIUM 45/100) - XML Element Type
- [Normal](element_Normal.md) (MEDIUM 45/100) - XML Element Type
- [NormalHighlit](element_NormalHighlit.md) (MEDIUM 45/100) - XML Element Type
- [Pressed](element_Pressed.md) (MEDIUM 45/100) - XML Element Type
- [PressedHighlit](element_PressedHighlit.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [ResizeImages](element_ResizeImages.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type

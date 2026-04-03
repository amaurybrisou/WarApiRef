# TextColors

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
| Addons seen in | AdvancedPetAssist, Assist, Aura, BlackBook, Crusher, DaemonAssist, DammazKron, Dascore |
| Files seen in | Configuration/Config.xml, Configuration/HopperConfig.xml, Configuration/WCDBConfig.xml, Configuration/WCDPConfig.xml, Core/Tome/DK_NFO.xml, Core/Tome/DK_TOK.xml, Source/DebugWindowVerticalScroll.xml, Source/EZCraftX.xml |
| Namespaces detected | TextColors |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APA_ComboBoxButton, AdvancedPetAssist: APA_ComboBoxButtonWide, Assist: AssistWindowBTN, Aura: Aura_Button_DefaultMenuButton, Aura: Aura_Button_DefaultMenuButtonLarge, Aura: Aura_Button_DefaultMenuButtonTiny |
| XML usage count | 84 |
| XML attribute usage count | 84 |
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

TextColors is a structural XML sub-element. It commonly appears under Button. It is typically used to organize structural children such as Disabled, DisabledPressed, Normal.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 85× (HIGH)

## Common Structural Child Elements

- [NormalHighlit](element_NormalHighlit.md) — 83× (HIGH)
- [PressedHighlit](element_PressedHighlit.md) — 83× (HIGH)
- [Normal](element_Normal.md) — 79× (HIGH)
- [Disabled](element_Disabled.md) — 76× (HIGH)
- [Pressed](element_Pressed.md) — 76× (HIGH)
- [DisabledPressed](element_DisabledPressed.md) — 26× (HIGH)

## Structural Sub-Elements

### [NormalHighlit](element_NormalHighlit.md)

Observed 83 times as an unnamed child.

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
### [PressedHighlit](element_PressedHighlit.md)

Observed 83 times as an unnamed child.

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
### [Normal](element_Normal.md)

Observed 79 times as an unnamed child.

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
### [Disabled](element_Disabled.md)

Observed 76 times as an unnamed child.

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
### [Pressed](element_Pressed.md)

Observed 76 times as an unnamed child.

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
### [DisabledPressed](element_DisabledPressed.md)

Observed 26 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-white, RightTabFrame, Map-Plus-Button-Disabled, Map-Minus-Button-Disabled |
| `a` | optional | 255, 0.5 |
| `b` | optional | 36 |
| `g` | optional | 57 |
| `r` | optional | 95 |
## Recursive Hierarchy

- Root: [TextColors](element_TextColors.md)
- [Disabled](element_Disabled.md) (structural, 76×, HIGH)
- [DisabledPressed](element_DisabledPressed.md) (structural, 26×, HIGH)
- [Normal](element_Normal.md) (structural, 79×, HIGH)
- [NormalHighlit](element_NormalHighlit.md) (structural, 83×, HIGH)
- [Pressed](element_Pressed.md) (structural, 76×, HIGH)
- [PressedHighlit](element_PressedHighlit.md) (structural, 83×, HIGH)

## Seen In

- AdvancedPetAssist
- Assist
- Aura
- BlackBook
- Crusher
- DaemonAssist
- DammazKron
- Dascore
- Deathblow
- Deathblow2
- DuffTimer
- EA_UiDebugTools
- EZCraftX
- FozAuction
- Hopper
- ItemRack
- LoyalPet
- Miracle Grow Remix
- MoraleSet
- Motion
- Pure
- RaidMeter
- RoR_SoR
- Rotation
- Sequencer
- Shinies
- SpamBayes
- Tome Titan
- TwisterSet
- WSCT
- WTes
- Warbuilder
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- ZonePOP
- zMailMod

## Examples

- AdvancedPetAssist: APA_ComboBoxButton -> TextColors in Button APA_ComboBoxButton
- AdvancedPetAssist: APA_ComboBoxButtonWide -> TextColors in Button APA_ComboBoxButtonWide
- Assist: AssistWindowBTN -> TextColors in Button AssistWindowBTN
- Aura: Aura_Button_DefaultMenuButton -> TextColors in Button Aura_Button_DefaultMenuButton
- Aura: Aura_Button_DefaultMenuButtonLarge -> TextColors in Button Aura_Button_DefaultMenuButtonLarge
- Aura: Aura_Button_DefaultMenuButtonTiny -> TextColors in Button Aura_Button_DefaultMenuButtonTiny

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [Disabled](element_Disabled.md) (MEDIUM 45/100) - XML Element Type
- [DisabledPressed](element_DisabledPressed.md) (MEDIUM 45/100) - XML Element Type
- [Normal](element_Normal.md) (MEDIUM 45/100) - XML Element Type
- [NormalHighlit](element_NormalHighlit.md) (MEDIUM 45/100) - XML Element Type
- [Pressed](element_Pressed.md) (MEDIUM 45/100) - XML Element Type
- [PressedHighlit](element_PressedHighlit.md) (MEDIUM 45/100) - XML Element Type

# TexSlices

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
| Addons seen in | AggroMeter, AnywhereTrainer, Assist, Atlas, AutoBand, BankArkel, BuffHead, CCTV |
| Files seen in | Conf/DK_Config.xml, Core/DK_Core.xml, Core/Tome/DK_TOK.xml, Modules/Config/Shinies-Config-Price.xml, Modules/UI/Shinies-UI-Auctions.xml, Modules/UI/Shinies-UI-Auto.xml, Modules/UI/Shinies-UI-Browse.xml, Modules/UI/Shinies-UI-Post.xml |
| Namespaces detected | TexSlices |
| Source kinds | xml_frames |
| Example locations | AggroMeter: Aggro_Button_Template, AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage, Assist: AssistWindowBTN, Atlas: TAtlasDirectionButton, AutoBand: AutoBandMapIconButton, BankArkel: PackWinTab1 |
| XML usage count | 150 |
| XML attribute usage count | 150 |
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

TexSlices is a structural XML sub-element. It commonly appears under Button and FullResizeImage. It is typically used to organize structural children such as Bottom, BottomCenter, BottomLeft.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 129× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 11× (HIGH)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 7× (MEDIUM)
- [VerticalResizeImage](element_VerticalResizeImage.md) — 3× (MEDIUM)

## Common Structural Child Elements

- [Normal](element_Normal.md) — 120× (HIGH)
- [Pressed](element_Pressed.md) — 117× (HIGH)
- [NormalHighlit](element_NormalHighlit.md) — 104× (HIGH)
- [PressedHighlit](element_PressedHighlit.md) — 95× (HIGH)
- [Disabled](element_Disabled.md) — 53× (HIGH)
- [DisabledPressed](element_DisabledPressed.md) — 36× (HIGH)
- [BottomCenter](element_BottomCenter.md) — 11× (HIGH)
- [BottomLeft](element_BottomLeft.md) — 11× (HIGH)
- [BottomRight](element_BottomRight.md) — 11× (HIGH)
- [MiddleCenter](element_MiddleCenter.md) — 11× (HIGH)
- [MiddleLeft](element_MiddleLeft.md) — 11× (HIGH)
- [MiddleRight](element_MiddleRight.md) — 11× (HIGH)
- [TopCenter](element_TopCenter.md) — 11× (HIGH)
- [TopLeft](element_TopLeft.md) — 11× (HIGH)
- [TopRight](element_TopRight.md) — 11× (HIGH)
- [Middle](element_Middle.md) — 10× (HIGH)
- [Left](element_Left.md) — 3× (MEDIUM)
- [Right](element_Right.md) — 3× (MEDIUM)
- [Bottom](element_Bottom.md) — 2× (LOW)
- [Top](element_Top.md) — 2× (LOW)

## Structural Sub-Elements

### [Normal](element_Normal.md)

Observed 120 times as an unnamed child.

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
### [Pressed](element_Pressed.md)

Observed 117 times as an unnamed child.

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
### [NormalHighlit](element_NormalHighlit.md)

Observed 104 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-white, RightTabFrame-Rollover, Tactics-Button-Rollover, SquareButton-Rollover |
| `b` | optional | 63, 0, 36, 50 |
| `g` | optional | 213, 85, 57, 192 |
| `r` | optional | 250, 255, 95, 222 |
| `a` | optional | 255, 1 |
| `x` | optional | 27, 105, 0, 201 |
| `y` | optional | 28, 44, 0, 341 |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_AnimatedImage_DefaultChoiceOverlay, EA_FullResizeImage_RedTransparent, EA_Button_GuildRosterRowHighlight |
| `texture` | optional | EA_SquareFrame_Highlight, CrusherIconBorderHighlight, EA_RoundFrame_Pressed, PinBG |
### [PressedHighlit](element_PressedHighlit.md)

Observed 95 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-white, RightTabFrame-Rollover, Tactics-Button-Depressed, SquareButton-Depressed |
| `b` | optional | 63, 36, 0, 45 |
| `g` | optional | 213, 57, 85, 255 |
| `r` | optional | 250, 95, 255, 226 |
| `a` | optional | 255, 1 |
| `x` | optional | 0, 120, 172, 475 |
| `y` | optional | 56, 44, 370, 420 |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_Button_GuildRosterRowPressed, BuffHeadLayoutVerticalButtonPressed, BuffHeadLayoutHorizontalButtonPressed |
### [Disabled](element_Disabled.md)

Observed 53 times as an unnamed child.

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
### [DisabledPressed](element_DisabledPressed.md)

Observed 36 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-white, RightTabFrame, Map-Plus-Button-Disabled, Map-Minus-Button-Disabled |
| `a` | optional | 255, 0.5 |
| `b` | optional | 36 |
| `g` | optional | 57 |
| `r` | optional | 95 |
### [BottomCenter](element_BottomCenter.md)

Observed 11 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 46, 272, 356, 10 |
| `y` | optional | 609, 197, 684, 54 |
| `id` | optional | Action-Bar-Frame-Bottom-Center, Border-Bottom-Center |
### [BottomLeft](element_BottomLeft.md)

Observed 11 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 36, 256, 346, 0 |
| `y` | optional | 609, 197, 684, 54 |
| `id` | optional | Action-Bar-Frame-Bottom-Left, Border-Bottom-Left |
### [BottomRight](element_BottomRight.md)

Observed 11 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 10, 137, 16, 350 |
| `y` | **required** | 9, 609, 15, 197 |
| `id` | optional | Action-Bar-Frame-Bottom-Right, Border-Bottom-Right |
### [MiddleCenter](element_MiddleCenter.md)

Observed 11 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 46, 272, 356, 10 |
| `y` | optional | 581, 172, 680, 10 |
| `id` | optional | Action-Bar-Frame-Middle-Center, Border-Middle-Center |
### [MiddleLeft](element_MiddleLeft.md)

Observed 11 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 36, 256, 346, 0 |
| `y` | optional | 581, 172, 680, 10 |
| `id` | optional | Action-Bar-Frame-Middle-Left, Border-Middle-Left |
### [MiddleRight](element_MiddleRight.md)

Observed 11 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 137, 350, 346, 54 |
| `y` | optional | 581, 172, 680, 10 |
| `id` | optional | Action-Bar-Frame-Middle-Right, Border-Middle-Right |
### [TopCenter](element_TopCenter.md)

Observed 11 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 46, 272, 356, 10 |
| `y` | optional | 571, 157, 669, 0 |
| `id` | optional | Action-Bar-Frame-Top-Center, Border-Top-Center |
### [TopLeft](element_TopLeft.md)

Observed 11 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 10, 36, 16, 256 |
| `y` | **required** | 10, 571, 15, 157 |
| `id` | optional | Action-Bar-Frame-Top-Left, Border-Top-Left |
### [TopRight](element_TopRight.md)

Observed 11 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 137, 350, 400, 54 |
| `y` | optional | 571, 157, 669, 0 |
| `id` | optional | Action-Bar-Frame-Top-Right, Border-Top-Right |
### [Middle](element_Middle.md)

Observed 10 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0, 424, 50, 78 |
| `y` | **required** | 0, 762, 23, 25 |
| `id` | optional | Order-VP-Bar-horiz, Dest-VP-Bar-horiz, Victor-Bar-horiz, Loser-Bar-horiz |
### [Left](element_Left.md)

Observed 3 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0, 6, 72, 346 |
| `y` | **required** | 0, 43, 241, 655 |
| `id` | optional | Dest-VP-Bar-horiz-End-Cap, Loser-Bar-horiz-End-Cap, objectivebar_left |
### [Right](element_Right.md)

Observed 3 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0, 418, 59, 119 |
| `y` | **required** | 30, 24, 43, 0 |
| `id` | optional | Order-VP-Bar-horiz-End-Cap, Victor-Bar-horiz-End-Cap, objectivebar_right |
### [Bottom](element_Bottom.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | Scrollbar-Bottom-ROLLOVER, Scrollbar-Bottom |
| `x` | optional | 172, 0 |
| `y` | optional | 429, 50 |
### [Top](element_Top.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | Scrollbar-Top-ROLLOVER, Scrollbar-Top |
| `x` | optional | 172, 0 |
| `y` | optional | 403, 0 |
## Recursive Hierarchy

- Root: [TexSlices](element_TexSlices.md)
- [Bottom](element_Bottom.md) (structural, 2×, LOW)
- [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
- [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
- [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
- [Disabled](element_Disabled.md) (structural, 53×, HIGH)
- [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
- [Left](element_Left.md) (structural, 3×, MEDIUM)
- [Middle](element_Middle.md) (structural, 10×, HIGH)
- [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
- [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
- [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
- [Normal](element_Normal.md) (structural, 120×, HIGH)
- [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
- [Pressed](element_Pressed.md) (structural, 117×, HIGH)
- [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
- [Right](element_Right.md) (structural, 3×, MEDIUM)
- [Top](element_Top.md) (structural, 2×, LOW)
- [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
- [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
- [TopRight](element_TopRight.md) (structural, 11×, HIGH)

## Seen In

- AggroMeter
- AnywhereTrainer
- Assist
- Atlas
- AutoBand
- BankArkel
- BuffHead
- CCTV
- CMap
- CaVES
- CastSequence
- Crusher
- DammazKron
- Deathblow
- Deathblow2
- EA_ThreePartBar
- FlagCap
- GuardList
- GuardRange
- GuildWarden
- Hopper
- ItemRack
- KillTracker
- Minmap
- Miracle Grow Remix
- MoraleSet
- Motion
- NerfedButtons
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- Queue Queuer
- RVMOD_Manager
- RaidMeter
- RoR_SoR
- SOR
- Sequencer
- Shinies
- SocialWindow 2.0
- SpamBayes
- Statdoll
- Statdoll Remix
- TalismanGenie
- TastyButtons
- TwisterSet
- Vectors
- VerticalMorale
- WARCommander
- WaaaghBar
- Warbuilder
- emotes
- nLootLink
- zMailMod

## Examples

- AggroMeter: Aggro_Button_Template -> TexSlices in Button Aggro_Button_Template
- AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage -> TexSlices in Button AnywhereTrainerTabTemplateInactiveImage
- Assist: AssistWindowBTN -> TexSlices in Button AssistWindowBTN
- Atlas: TAtlasDirectionButton -> TexSlices in Button TAtlasDirectionButton
- AutoBand: AutoBandMapIconButton -> TexSlices in Button AutoBandMapIconButton
- BankArkel: PackWinTab1 -> TexSlices in Button PackWinTab1

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [HorizontalResizeImage](element_HorizontalResizeImage.md) (HIGH 100/100) - XML Element Type
- [VerticalResizeImage](element_VerticalResizeImage.md) (HIGH 100/100) - XML Element Type
- [BottomCenter](element_BottomCenter.md) (MEDIUM 45/100) - XML Element Type
- [BottomLeft](element_BottomLeft.md) (MEDIUM 45/100) - XML Element Type
- [BottomRight](element_BottomRight.md) (MEDIUM 45/100) - XML Element Type
- [Disabled](element_Disabled.md) (MEDIUM 45/100) - XML Element Type
- [DisabledPressed](element_DisabledPressed.md) (MEDIUM 45/100) - XML Element Type
- [Left](element_Left.md) (MEDIUM 45/100) - XML Element Type
- [Middle](element_Middle.md) (MEDIUM 45/100) - XML Element Type
- [MiddleCenter](element_MiddleCenter.md) (MEDIUM 45/100) - XML Element Type
- [MiddleLeft](element_MiddleLeft.md) (MEDIUM 45/100) - XML Element Type
- [MiddleRight](element_MiddleRight.md) (MEDIUM 45/100) - XML Element Type
- [Normal](element_Normal.md) (MEDIUM 45/100) - XML Element Type
- [NormalHighlit](element_NormalHighlit.md) (MEDIUM 45/100) - XML Element Type
- [Pressed](element_Pressed.md) (MEDIUM 45/100) - XML Element Type
- [PressedHighlit](element_PressedHighlit.md) (MEDIUM 45/100) - XML Element Type
- [Right](element_Right.md) (MEDIUM 45/100) - XML Element Type
- [TopCenter](element_TopCenter.md) (MEDIUM 45/100) - XML Element Type
- [TopLeft](element_TopLeft.md) (MEDIUM 45/100) - XML Element Type
- [TopRight](element_TopRight.md) (MEDIUM 45/100) - XML Element Type
- [Bottom](element_Bottom.md) (MEDIUM 30/100) - XML Element Type
- [Top](element_Top.md) (MEDIUM 30/100) - XML Element Type

## Used With

- [Sizes](element_Sizes.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type

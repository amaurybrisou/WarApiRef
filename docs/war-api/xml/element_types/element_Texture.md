# Texture

- Category: XML Element Type
- Confidence level: MEDIUM
- Confidence score: 55/100

## Confidence Assessment

- Level: MEDIUM

- Score: 55/100

- Rationale: unknown

## Evidence Signals

- none

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | none |
| Source kinds | none |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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

Observed XML element type instantiated by 19 addons.

## Common Attributes

- file

## Common Inherits

- none

## Common Parent Elements

- [Assets](element_Assets.md) — 374× (HIGH)

## Common Structural Child Elements

- [Slice](element_Slice.md) — 73× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `file` | **required** | 100% | SharedAssets/Flat.tga, SharedAssets/Minimalist.tga, SharedAssets/Smoothv2.tga, SharedAssets/AceBarFrames.tga, ... |
## Structural Sub-Elements

### [Slice](element_Slice.md)

Observed 73 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `height` | **required** | 3, 58, 64 |
| `id` | **required** | Border-Top-Center, Border-Top-Left, Border-Top-Right, Border-Middle-Center |
| `left` | **required** | 3, 0, 64, 61 |
| `top` | **required** | 0, 3, 61, 64 |
| `width` | **required** | 58, 3, 64 |
## Recursive Hierarchy

- Root: [Texture](element_Texture.md)
- [Slice](element_Slice.md) (structural, 73×, HIGH)

## Lua Functions Manipulating This Type

- BuffHead.BuffHeadEffectFrame:SetLayout
- TurretRange.ShowElement
- WoHReticle.CreateRing
- TurretRange.Display.SetAlpha
- Enemy.EnemyGroupIcon:Attach
- TurretRange.UpdateDisplay
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- MoraleCircle.ColorChanger2
- Enemy.EnemyEffectsIndicator:Update
- MoraleCircle.OnSetCustomColorEmpty
- TurretRange.Display.UpdateLayout
- TurretRange.OnUpdate
- MoraleCircle.ColorChanger4
- TurretRange.Setup.Display.Initialize
- FilterSettings.DisableCondenseSettings
- GuardLine.update
- TurretRange.local.UpdateDisplay
- TurretRange.Setup.Display.LoadSettings
- TurretRange.Setup.Display.OnCircleInvertLUp
- MoraleCircle.ColorChanger1
- Aura.Aura:UpdateWindow
- DAoCBuff.DAoCBuffFrame:SetBuff
- Enemy.UI_Icon_Switch
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- RoR_SoR.T1_UPDATE
- MoraleCircle.ColorChanger3
- TurretRange.Setup.Display.OnCircleModeChanged
- MoraleCircle.OnSetCustomColorFill
- MoraleCircle.OnSetCustomColor
- WoHReticle.UpdateTargets
- BuffHead.BuffHeadEffectFrame:Update
- TurretRange.local.ShowElement
- MoraleCircle.OnSetCustomColorFull

## Seen In

- Ace
- AggroMeter
- AutoMark
- BagOMatic
- BankArkel
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- GuardLine
- LibSlash
- MoraleCircle
- RoR_SoR
- Shinies
- Swift Assist
- TexturedButtons
- TidyRoll
- TurretRange
- WoH-Reticle

## Examples

- none

## Related APIs

- [Assets](element_Assets.md) (MEDIUM 45/100) - XML Element Type
- [Slice](element_Slice.md) (MEDIUM 35/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none

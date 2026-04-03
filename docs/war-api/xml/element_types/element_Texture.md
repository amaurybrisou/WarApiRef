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

- RoR_SoR.T1_UPDATE
- DAoCBuff.DAoCBuffFrame:SetBuff
- TurretRange.local.UpdateDisplay
- MoraleCircle.OnSetCustomColorEmpty
- MoraleCircle.OnSetCustomColor
- MoraleCircle.ColorChanger1
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- TurretRange.UpdateDisplay
- MoraleCircle.ColorChanger2
- WoHReticle.UpdateTargets
- Enemy.UI_Icon_Switch
- TurretRange.Setup.Display.Initialize
- Enemy.EnemyEffectsIndicator:Update
- FilterSettings.DisableCondenseSettings
- TurretRange.Display.UpdateLayout
- TurretRange.Setup.Display.LoadSettings
- MoraleCircle.OnSetCustomColorFull
- MoraleCircle.ColorChanger4
- GuardLine.update
- TurretRange.ShowElement
- Enemy.EnemyGroupIcon:Attach
- Aura.Aura:UpdateWindow
- TurretRange.Setup.Display.OnCircleInvertLUp
- WoHReticle.CreateRing
- BuffHead.BuffHeadEffectFrame:Update
- TurretRange.OnUpdate
- MoraleCircle.ColorChanger3
- TurretRange.Setup.Display.OnCircleModeChanged
- TurretRange.Display.SetAlpha
- TurretRange.local.ShowElement
- BuffHead.BuffHeadEffectFrame:SetLayout
- MoraleCircle.OnSetCustomColorFill
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged

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

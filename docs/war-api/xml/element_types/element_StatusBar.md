# StatusBar

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BlackBox, CCTV, CleanCastbar, CleanUnitFrames, FlagCap, PartyCast, RaidMeter, RoR_SoR |
| Namespaces detected | StatusBar |
| Source kinds | xml_frames |
| Example locations | BlackBox: RespawnTimerWindowBar, CCTV: CCTVRootWindowBar, CCTV: CCTVRootWindowBar2, CCTV: CCTVSnareWindowBar, CCTV: CCTVSnareWindowBar2, CCTV: CCTVStaggerWindowBar |
| XML usage count | 32 |
| XML attribute usage count | 32 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
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

StatusBar is a container-style XML element. It commonly appears under Window. It is typically used to organize structural children such as Anchors, Size and be manipulated from Lua by functions such as BlackBox.OnApplicationTwoButtonDialogHook, BlackBox.UpdateTimer.

## Common Attributes

- background
- foreground
- handleinput
- inherits
- interpolate
- layer
- name
- popable
- reverseFill

## Common Inherits

- EA_StatusBar_DefaultTintable
- EA_StatusBar_GuildXP
- RRQTomeStatusBar
- UnitFrameAPStatusBar
- UnitFrameFriendlyHealthStatusBar
- UnitFrameFriendlyHealthStatusSmallBar
- UnitFrameHostileStatusBar

## Common Parent Elements

- [Windows](element_Windows.md) — 32× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 32× (HIGH)
- [Size](element_Size.md) — 19× (HIGH)

## Common Template Bases

- EA_StatusBar_DefaultTintable
- EA_StatusBar_GuildXP
- RRQTomeStatusBar
- UnitFrameAPStatusBar
- UnitFrameFriendlyHealthStatusBar
- UnitFrameFriendlyHealthStatusSmallBar
- UnitFrameHostileStatusBar


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `foreground` | optional | 68% | EA_StatusBar_DefaultTintableBar, CleanCastbarFG, ObjectiveBarFill, LOL_BAR, ... |
| `handleinput` | optional | 50% | false |
| `popable` | optional | 50% | false |
| `inherits` | optional | 31% | EA_StatusBar_GuildXP, UnitFrameFriendlyHealthStatusBar, UnitFrameHostileStatusBar, UnitFrameAPStatusBar, ... |
| `interpolate` | optional | 31% | false, true |
| `layer` | optional | 15% | background, default |
| `reverseFill` | optional | 15% | true |
| `background` | optional | 6% | CleanCastbarBG, LOL_BAR |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 32 times as an unnamed child.

### [Size](element_Size.md)

Observed 19 times as an unnamed child.

## Recursive Hierarchy

- Root: [StatusBar](element_StatusBar.md)
- [Anchors](element_Anchors.md) (structural, 32×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
      - (cycle)
- [Size](element_Size.md) (structural, 19×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)

## Lua Functions Manipulating This Type

- BlackBox.UpdateTimer
- BlackBox.OnApplicationTwoButtonDialogHook

## Seen In

- BlackBox
- CCTV
- CleanCastbar
- CleanUnitFrames
- FlagCap
- PartyCast
- RaidMeter
- RoR_SoR
- SNT_CASTBAR
- SNT_INFO
- Targets

## Examples

- BlackBox: RespawnTimerWindowBar -> StatusBar RespawnTimerWindowBar
- CCTV: CCTVRootWindowBar -> StatusBar CCTVRootWindowBar
- CCTV: CCTVRootWindowBar2 -> StatusBar CCTVRootWindowBar2
- CCTV: CCTVSnareWindowBar -> StatusBar CCTVSnareWindowBar
- CCTV: CCTVSnareWindowBar2 -> StatusBar CCTVSnareWindowBar2
- CCTV: CCTVStaggerWindowBar -> StatusBar CCTVStaggerWindowBar

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [EA_StatusBar_GuildXP](../../globals/constants/constant_EA_StatusBar_GuildXP.md) (HIGH 100/100) - Constant
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EA_StatusBar_DefaultTintable](../../globals/constants/constant_EA_StatusBar_DefaultTintable.md) (HIGH 90/100) - Constant

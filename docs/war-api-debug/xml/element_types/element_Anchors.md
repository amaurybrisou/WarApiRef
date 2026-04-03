# Anchors

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 98/100

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin, Clock |
| Namespaces detected | Anchors |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot10, CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot2, CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot3, CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot4, CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot5, CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot6 |
| XML usage count | 153 |
| XML attribute usage count | 153 |
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

Anchors is a structural XML sub-element. It commonly appears under Button and DynamicImage. It is typically used to organize structural children such as Anchor.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Label](element_Label.md) — 60× (HIGH)
- [Button](element_Button.md) — 53× (HIGH)
- [Window](element_Window.md) — 20× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 12× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 6× (MEDIUM)
- [ListBox](element_ListBox.md) — 2× (LOW)

## Common Structural Child Elements

- [Anchor](element_Anchor.md) — 167× (HIGH)

## Structural Sub-Elements

### [Anchor](element_Anchor.md)

Observed 167 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `point` | **required** | left, topleft, bottomright, topright |
| `relativePoint` | **required** | left, topleft, bottomright, top |
| `relativeTo` | **required** | $parentName, $parentTactics, $parentZone, $parentSet |
## Recursive Hierarchy

- Root: [Anchors](element_Anchors.md)
- [Anchor](element_Anchor.md) (structural, 167×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)

## Seen In

- CM_ClosetGoblin
- Clock

## Examples

- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot10 -> Anchors in Button CG_ItemRackEQShow1EquipmentSlot10
- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot2 -> Anchors in Button CG_ItemRackEQShow1EquipmentSlot2
- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot3 -> Anchors in Button CG_ItemRackEQShow1EquipmentSlot3
- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot4 -> Anchors in Button CG_ItemRackEQShow1EquipmentSlot4
- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot5 -> Anchors in Button CG_ItemRackEQShow1EquipmentSlot5
- CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot6 -> Anchors in Button CG_ItemRackEQShow1EquipmentSlot6

## Related APIs

- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Button](element_Button.md) (HIGH 90/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 90/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 90/100) - XML Element Type
- [ListBox](element_ListBox.md) (HIGH 90/100) - XML Element Type
- [Anchor](element_Anchor.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [Size](element_Size.md) (HIGH 98/100) - XML Element Type

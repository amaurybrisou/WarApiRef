# ListData

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
| Addons seen in | AggroMeter, Aura, BlackBook, BuffHead, CDown, CM_ClosetGoblin, CMap, CastSequence |
| Files seen in | Code/CombatLog/CombatLogSnapshotWindow.xml, Code/CombatLog/CombatLogStatsWindow.xml, Code/Core/ChooseIconDialog.xml, Code/Core/Debug.xml, Code/KillSpam/KillSpam.xml, Code/ScenarioInfo/ScenarioInfo.xml, Code/UnitFrames/EffectsIndicatorDialog.xml, Code/UnitFrames/UnitFramesConfiguration.xml |
| Namespaces detected | ListData |
| Source kinds | xml_frames |
| Example locations | AggroMeter: AggroMeterGrayListBox, Aura: AuraSettingsAuraList, Aura: AuraSharesAuraList, Aura: AuraTextureIconsIcons, BlackBook: BlackBookWindowList, BuffHead: BuffHeadSetupAdvancedCompressionItemWindowList |
| XML usage count | 110 |
| XML attribute usage count | 110 |
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

ListData is a structural XML sub-element that binds list controls to Lua-backed row data. It commonly appears under ListBox as list data wiring.

## Common Attributes

- populationfunction
- table

## Common Inherits

- none

## Common Parent Elements

- [ListBox](element_ListBox.md) — 110× (HIGH)

## Common Structural Child Elements

- [ListColumns](element_ListColumns.md) — 74× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `table` | **required** | 100% | AggroMeter.Listdata, AuraSettings.listDisplayData, AuraShares.listDisplayData, AuraTexture.listIconDisplayData, ... |
| `populationfunction` | **required** | 98% | AuraSettings.PopulateDisplay, AuraShares.PopulateDisplay, AuraTexture.PopulateIconsListDisplay, BlackBookWindow.Populate, ... |
## Structural Sub-Elements

### [ListColumns](element_ListColumns.md)

Observed 74 times as an unnamed child.

## Recursive Hierarchy

- Root: [ListData](element_ListData.md)
- [ListColumns](element_ListColumns.md) (structural, 74×, HIGH)
  - [ListColumn](element_ListColumn.md) (structural, 192×, HIGH)

## Seen In

- AggroMeter
- Aura
- BlackBook
- BuffHead
- CDown
- CM_ClosetGoblin
- CMap
- CastSequence
- Cheeseboard
- Crusher
- DAoCBuff
- DPSMeter
- Dascore
- Deathblow
- Deathblow2
- DetauntHelper
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiModWindow
- Emojii
- Enemy
- FozAuction
- HealGrid
- Hopper
- Kwestor
- LibAddonButton
- LootAlert
- MapMonster
- MapPin
- Minmap
- Motion
- NerfedButtons
- Obsidian
- PieTracker
- Pocket Palette
- Pure
- QuickTacticSwitch
- RealmStatus
- SessionRPs
- Shinies
- SocialWindow 2.0
- TastyButtons
- TaxPayer
- TexturedButtons
- TidyChat
- TidyRoll
- Tome Titan
- TomeTracker
- Tortall_DPS
- TurretRange
- WTes
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- nLootLink
- wbLeadHelper
- zMailMod

## Examples

- AggroMeter: AggroMeterGrayListBox -> ListData in ListBox AggroMeterGrayListBox
- Aura: AuraSettingsAuraList -> ListData in ListBox AuraSettingsAuraList
- Aura: AuraSharesAuraList -> ListData in ListBox AuraSharesAuraList
- Aura: AuraTextureIconsIcons -> ListData in ListBox AuraTextureIconsIcons
- BlackBook: BlackBookWindowList -> ListData in ListBox BlackBookWindowList
- BuffHead: BuffHeadSetupAdvancedCompressionItemWindowList -> ListData in ListBox BuffHeadSetupAdvancedCompressionItemWindowList

## Related APIs

- [ListBox](element_ListBox.md) (HIGH 100/100) - XML Element Type
- [ListColumns](element_ListColumns.md) (MEDIUM 55/100) - XML Element Type

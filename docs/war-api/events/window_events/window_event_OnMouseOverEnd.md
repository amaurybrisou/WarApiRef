# OnMouseOverEnd

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 118

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy, MiracleGrowLight, Shinies, TexturedButtons, TurretRange |
| Files seen in | `/workspace/data/raw/BuffHead/Display.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainers.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupEffectCache.xml:0` |
| Namespaces detected | OnMouseOverEnd |
| Source kinds | event_page, xml_handlers |
| Example locations | BuffHead: BuffHeadBuffTemplate.OnMouseOverEnd, BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate.OnMouseOverEnd, BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate.OnMouseOverEnd, BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowTitleLabel.OnMouseOverEnd, BuffHead: BuffHeadSetupAdvancedContainersItemRowTemplate.OnMouseOverEnd, BuffHead: BuffHeadSetupAdvancedContainersRowTemplate.OnMouseOverEnd |
| XML usage count | 50 |
| XML attribute usage count | 50 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
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

Observed as an engine-supplied UI event hook used by 9 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- MiracleGrowLight
- Shinies
- TexturedButtons
- TurretRange
- WarBoard

## Registrars And Handlers

- BuffHead.Setup.AdvancedCompression.OnRowMouseOut
- BuffHead.Setup.AdvancedCompressionItem.OnRowMouseOut
- BuffHead.Setup.AdvancedContainers.OnRowMouseOut
- BuffHead.Setup.AdvancedContainersItem.OnRowMouseOut
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOut
- BuffHead.Setup.EffectCache.OnRowMouseOut
- BuffHead.Setup.Filter.OnRowMouseOut
- BuffHead.Setup.Layout.Properties.OnColorExampleMouseOut
- BuffHead.Setup.Layout.Properties.OnFontExampleMouseOut
- BuffHead.Setup.Layout.Properties.OnPropertyTitleMouseOut
- BuffHead.Setup.Layout.Properties.OnTextureButtonMouseOut
- BuffHead.Setup.OnRowMouseOut
- BuffHead.Setup.PriorityEffects.OnRowMouseOut
- BuffHead.Setup.PriorityEffectsItem.OnRowMouseOut
- BuffHead.Setup.SelectTexture.OnTextureRowMouseOut
- ClosetGoblinCharacterWindow.EquipmentMouseOverEnd
- ClosetGoblinCharacterWindow.HideCloakOptions
- ClosetGoblinCharacterWindow.HideShowHelm
- DAoCBuffFrame.OnMouseOverEnd
- Enemy.CombatLogUI_IDSAnchor_OnMouseOverEnd
- Enemy.KillSpamUI_KillSpamDialog_OnMouseOverEnd
- Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOverEnd
- Enemy.MarksUI_EnemyMarksWindow_OnMouseOverEnd
- Enemy.UnitFramesUI_Anchor_OnMouseOverEnd
- FrameManager.OnMouseOverEnd
- MiracleGrowLight.harvestEnd
- ShiniesAuctionsUI.OnMouseOverEnd_Results_ListItem
- ShiniesAutoUI.OnMouseOverEnd_AutoListRow
- ShiniesBrowseUI.OnMouseOverEnd_Results_ListItem
- ShiniesBrowseUI.OnMouseOverEnd_Searches
- ShiniesPostUI.OnMouseOverEnd_Results_ListItem
- TexturedButtons.Setup.Cooldown.OnColorExampleMouseOut
- TexturedButtons.Setup.Fonts.OnColorExampleMouseOut
- TexturedButtons.Setup.Fonts.OnFontExampleMouseOut
- TexturedButtons.Setup.OnRowMouseOut
- TurretRange.Setup.Display.OnDistanceFontMouseOut
- TurretRange.Setup.Distances.OnRowMouseOut
- WarBoard.OnMouseOverEnd
- WarBoard.OnMouseOverEndBottom

## Examples

- BuffHead: BuffHeadBuffTemplate -> BuffHeadBuffTemplate.OnMouseOverEnd -> FrameManager.OnMouseOverEnd
- BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate -> BuffHeadSetupAdvancedCompressionItemRowTemplate.OnMouseOverEnd -> BuffHead.Setup.AdvancedCompressionItem.OnRowMouseOut
- BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate -> BuffHeadSetupAdvancedCompressionRowTemplate.OnMouseOverEnd -> BuffHead.Setup.AdvancedCompression.OnRowMouseOut
- BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowTitleLabel -> BuffHeadSetupAdvancedContainersItemPropertiesWindowTitleLabel.OnMouseOverEnd -> BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOut
- BuffHead: BuffHeadSetupAdvancedContainersItemRowTemplate -> BuffHeadSetupAdvancedContainersItemRowTemplate.OnMouseOverEnd -> BuffHead.Setup.AdvancedContainersItem.OnRowMouseOut
- BuffHead: BuffHeadSetupAdvancedContainersRowTemplate -> BuffHeadSetupAdvancedContainersRowTemplate.OnMouseOverEnd -> BuffHead.Setup.AdvancedContainers.OnRowMouseOut

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none

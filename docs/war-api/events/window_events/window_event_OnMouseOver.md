# OnMouseOver

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
| Addons seen in | AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Aura, BagOMatic, BankArkel, BuffHead, CM_ClosetGoblin |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/BankArkel/BankArkel.xml:0`, `/workspace/data/raw/BuffHead/Display.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0` |
| Namespaces detected | OnMouseOver |
| Source kinds | event_page, xml_handlers |
| Example locations | AdvancedRenownTrainer: AbilityButtonTemplate.OnMouseOver, AggroMeter: AggroMeter_Button.OnMouseOver, AggroMeter: Aggro_Label_Template.OnMouseOver, AggroMeter: Aggro_Tactic_Template.OnMouseOver, AggroMeter: Aggro_Timer_Template.OnMouseOver, AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage.OnMouseOver |
| XML usage count | 230 |
| XML attribute usage count | 230 |
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

Observed as an engine-supplied UI event hook used by 24 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Aura
- BagOMatic
- BankArkel
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Killer
- MiracleGrowLight
- MoraleCircle
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- bigger_MacroWindow
- followTheLeader

## Registrars And Handlers

- ActionBars.OnMouseoverHotbarPageDown
- ActionBars.OnMouseoverHotbarPageUp
- AdvancedRenownTraining.AbilityTooltip
- AggroMeter.OnMouseOverStart
- AggroMeter.SelectChar
- AnywhereTrainer.OnMouseOver
- AuraConfig.OnIconMouseOver
- AuraSettings.OnMouseOverAuraList
- AuraShares.OnMouseOverAuraList
- BagOMatic.wnd_on_mouse_over
- BankArkel.PackIconOver
- BankArkel.PackTabMover
- BuffHead.Setup.AdvancedCompression.OnRowMouseOver
- BuffHead.Setup.AdvancedCompressionItem.OnRowMouseOver
- BuffHead.Setup.AdvancedContainers.OnRowMouseOver
- BuffHead.Setup.AdvancedContainersItem.OnRowMouseOver
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOver
- BuffHead.Setup.EffectCache.OnRowMouseOver
- BuffHead.Setup.EffectCache.OnSortButtonMouseOver
- BuffHead.Setup.Filter.OnRowMouseOver
- BuffHead.Setup.Layout.OnLayerMouseOver
- BuffHead.Setup.Layout.Properties.OnColorExampleMouseOver
- BuffHead.Setup.Layout.Properties.OnFontExampleMouseOver
- BuffHead.Setup.Layout.Properties.OnPropertyTitleMouseOver
- BuffHead.Setup.Layout.Properties.OnTextureButtonMouseOver
- BuffHead.Setup.OnRowMouseOver
- BuffHead.Setup.PriorityEffects.OnRowMouseOver
- BuffHead.Setup.PriorityEffectsItem.OnRowMouseOver
- BuffHead.Setup.SelectTexture.OnTextureRowMouseOver
- CG_ItemRack.EquipmentMouseOver
- ClosetGoblinCharacterWindow.EquipmentMouseOver
- ClosetGoblinCharacterWindow.ShowCloakOptions
- ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- ClosetGoblinCharacterWindow.ShowShowCloakOnly
- ClosetGoblinCharacterWindow.ShowShowHelm
- ClosetGoblinCharacterWindow.ShowShowHelmOnly
- DAoCBuffFrame.OnMouseOver
- DAoCBuffSettings.ShowTooltip
- EA_Window_Macro.DetailIconMouseOver
- EA_Window_Macro.IconMouseOver
- EA_Window_Macro.SelectionIconMouseOver
- Enemy.AssistUI_ConfigDialog_OnMacroMarkMouseOver
- Enemy.AssistUI_ConfigDialog_OnMacroTargetMouseOver
- Enemy.CombatLogUI_EpsWindow_OnMouseOver
- Enemy.CombatLogUI_IDSAnchor_OnMouseOver
- Enemy.CombatLogUI_SnapshotWindow_ListRowMouseOver
- Enemy.CombatLogUI_StatsWindow_ListRowMouseOver
- Enemy.CombatLogUI_StatsWindow_OnEpsMouseOver
- Enemy.CombatLogUI_TargetDefenseTotalWindow_OnMouseOver
- Enemy.CombatLogUI_TargetDefenseWindow_OnMouseOver
- Enemy.ConfigurationWindow_ShowTooltip
- Enemy.GroupsUI_EffectFilterDialog_OnAbilityIdsMouseOver
- Enemy.Guard_GuardIndicator_OnMouseOver
- Enemy.KillSpamUI_KillSpamAreaStatsDialog_OnRowMouseOver
- Enemy.KillSpamUI_KillSpamDialog_OnMouseOver
- Enemy.KillSpamUI_KillSpamDialog_OnRowMouseOver
- Enemy.KillSpamUI_KillSpamDialog_OnTotalMouseOver
- Enemy.KillSpamUI_PlayerKDR_OnMouseOver
- Enemy.MarksUI_EnemyMarkIcon_OnMouseOver
- Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOver
- Enemy.MarksUI_EnemyMarksWindow_OnMouseOver
- Enemy.ScenarioInfoUI_ConfigDialog_OnMacroToggleMouseOver
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip
- Enemy.TalismanAlerter_TalismanAlerterIndicator_OnMouseOver
- Enemy.TimerUI_OnMouseOver
- Enemy.UI_Icon_OnMouseOver
- Enemy.UnitFramesUI_Anchor_OnMouseOver
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionMouseOver
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig2CommandMouseOver
- Enemy.UnitFramesUI_UnitFrame_OnMouseOver
- FrameManager.OnMouseOver
- Killer.OnAllTimeSummaryMouseOver
- Killer.OnFeedRowMouseOver
- Killer.OnPersonalCounterMouseOver
- Killer.OnRecentSummaryMouseOver
- MiracleGrowLight.harvestStart
- MiracleGrowLight.onHover
- MoraleCircle.OnMouseOverStart
- PP.ItemSlotMouseOver
- PotionBarFloating.ActivatorMouseOver
- PotionBarFloating.Autoshow
- PotionBarFloating.MouseOver
- PotionBarSettings.OnMouseoverActivator
- PotionBarSettings.OnMouseoverAutohide
- PotionBarSettings.OnMouseoverBuild
- PotionBarSettings.OnMouseoverDontSplitName
- PotionBarSettings.OnMouseoverInfoTextBR
- PotionBarSettings.OnMouseoverInfoTextTR
- PotionBarSettings.OnMouseoverMethod
- PotionBarSettings.OnMouseoverMoveDown
- PotionBarSettings.OnMouseoverMoveUp
- PotionBarSettings.OnMouseoverQuickActions
- PotionBarSettings.OnMouseoverShow1
- PotionBarSettings.OnMouseoverShowLast
- PotionBarSettings.OnMouseoverUse
- RoR_SoR.MainTooltip
- RoR_SoR.OnMouseOverStart
- RoR_SoR.ShowPopper
- ShiniesAuctionsUI.OnMouseOver_Results_ListItem
- ShiniesAutoUI.OnMouseOver_AutoListRow
- ShiniesBrowseUI.OnMouseOver_Results_ListItem
- ShiniesBrowseUI.OnMouseOver_Searches
- ShiniesPostUI.OnMouseOver_Results_ListItem
- TexturedButtons.Setup.Cooldown.OnColorExampleMouseOver
- TexturedButtons.Setup.Fonts.OnColorExampleMouseOver
- TexturedButtons.Setup.Fonts.OnFontExampleMouseOver
- TexturedButtons.Setup.OnRowMouseOver
- TurretRange.Setup.Display.OnDistanceFontMouseOver
- TurretRange.Setup.Distances.OnRowMouseOver
- WSCT.CritOnMouseOver
- WSCT.FrameComboOnMouseOver
- WSCT.OnMouseOver
- WarBoard.OnMouseOver
- WarBoard.OnMouseOverBottom
- followTheLeader.OnMouseOver

## Examples

- AdvancedRenownTrainer: AbilityButtonTemplate -> AbilityButtonTemplate.OnMouseOver -> AdvancedRenownTraining.AbilityTooltip
- AggroMeter: AggroMeter_Button -> AggroMeter_Button.OnMouseOver -> AggroMeter.OnMouseOverStart
- AggroMeter: Aggro_Label_Template -> Aggro_Label_Template.OnMouseOver -> AggroMeter.SelectChar
- AggroMeter: Aggro_Tactic_Template -> Aggro_Tactic_Template.OnMouseOver -> AggroMeter.OnMouseOverStart
- AggroMeter: Aggro_Timer_Template -> Aggro_Timer_Template.OnMouseOver -> AggroMeter.OnMouseOverStart
- AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage -> AnywhereTrainerTabTemplateInactiveImage.OnMouseOver -> AnywhereTrainer.OnMouseOver

## Related APIs

- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function

## Used With

- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none

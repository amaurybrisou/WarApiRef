# OnHidden

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 188

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BuffHead, CM_ClosetGoblin, Enemy, RoR_SoR, Shinies |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItemEffect.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainers.xml:0` |
| Namespaces detected | OnHidden |
| Source kinds | event_page, flows, lua_event_registration, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD.OnHidden, AdvancedPetAssist: APAInstantOnlyHUD.OnHidden, AdvancedPetAssist: APAKitingHUD.OnHidden, AdvancedPetAssist: APAOptions.OnHidden, AdvancedRenownTrainer: AdvancedRenownTrainingExportWindow.OnHidden, AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindow.OnHidden |
| XML usage count | 79 |
| XML attribute usage count | 79 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 3 |
| Initialization flow references | 1 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
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

Observed as an engine-supplied UI event hook used by 15 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BuffHead
- CM_ClosetGoblin
- Enemy
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WhoHealedMe
- bigger_MacroWindow

## Registrars And Handlers

- APAGui.OnFollowTargetHUDHidden
- APAGui.OnHidden
- APAGui.OnInstantOnlyHUDHidden
- APAGui.OnKitingHUDHidden
- AdvancedRenownTraining.OnExportHidden
- AdvancedRenownTraining.OnHidden
- AdvancedRenownTraining.OnImportHidden
- AdvancedRenownTraining.OnImportNameInputHidden
- AdvancedRenownTraining.OnLinkHidden
- AdvancedRenownTraining.PresetOnHidden
- AuraConfig.OnHidden
- AuraSettings.OnHidden
- AuraShares.OnHidden
- BuffHead.Setup.AdvancedCompression.OnHidden
- BuffHead.Setup.AdvancedCompressionItem.OnHidden
- BuffHead.Setup.AdvancedCompressionItemEffect.OnHidden
- BuffHead.Setup.AdvancedContainers.OnHidden
- BuffHead.Setup.AdvancedContainersItem.OnHidden
- BuffHead.Setup.AdvancedContainersItem.Properties.OnHidden
- BuffHead.Setup.Container.OnHidden
- BuffHead.Setup.Display.OnHidden
- BuffHead.Setup.EffectCache.OnHidden
- BuffHead.Setup.EffectCacheTable.OnHidden
- BuffHead.Setup.Filter.OnHidden
- BuffHead.Setup.General.OnHidden
- BuffHead.Setup.Layout.Manager.OnHidden
- BuffHead.Setup.Layout.OnHidden
- BuffHead.Setup.Layout.Properties.OnHidden
- BuffHead.Setup.OnHidden
- BuffHead.Setup.Performance.OnHidden
- BuffHead.Setup.PriorityEffects.OnHidden
- BuffHead.Setup.PriorityEffectsItem.OnHidden
- BuffHead.Setup.Trackers.OnHidden
- ClosetGoblinCharacterWindow.OnHide
- ClosetGoblinZoneWindow.OnHide
- EA_Window_Macro.OnHidden
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnHidden
- Shinies.OnHidden
- ShiniesAuctionsUI.OnHidden
- ShiniesAutoUI.OnHidden
- ShiniesBrowseUI.OnHidden
- ShiniesPostUI.OnHidden
- TexturedButtons.Setup.Actionbar.OnHidden
- TexturedButtons.Setup.AdvancedTextures.OnHidden
- TexturedButtons.Setup.Cooldown.OnHidden
- TexturedButtons.Setup.Fonts.OnHidden
- TexturedButtons.Setup.Misc.OnHidden
- TexturedButtons.Setup.OnHidden
- TexturedButtons.Setup.Textures.OnHidden
- TexturedButtons.Setup.Tint.OnHidden
- TidyChat.Copy.OnHidden
- TidyChat.LootRoll.OnHidden
- TidyChat.OnEntryBoxUpdateShowing
- TidyChat.Options.OnHidden
- TidyRoll.CustomAutoRoll.OnHidden
- TidyRoll.OnEsc
- TidyRollOptions.OnHidden
- TurretRange.Setup.Display.OnHidden
- TurretRange.Setup.Distance.OnHidden
- TurretRange.Setup.Distances.OnHidden
- TurretRange.Setup.General.OnHidden
- TurretRange.Setup.OnHidden
- WHMGui.OnWindowHidden
- WSCT.OnHidden
- WindowRegisterCoreEventHandler
- WindowUtils.OnHidden
- core

## Examples

- TidyChat: TidyChatFrames.Initialize -> OnHidden -> TidyChat.OnEntryBoxUpdateShowing
- AdvancedPetAssist: APAFollowTargetHUD -> APAFollowTargetHUD.OnHidden -> APAGui.OnFollowTargetHUDHidden
- AdvancedPetAssist: APAInstantOnlyHUD -> APAInstantOnlyHUD.OnHidden -> APAGui.OnInstantOnlyHUDHidden
- AdvancedPetAssist: APAKitingHUD -> APAKitingHUD.OnHidden -> APAGui.OnKitingHUDHidden
- AdvancedPetAssist: APAOptions -> APAOptions.OnHidden -> APAGui.OnHidden
- AdvancedRenownTrainer: AdvancedRenownTrainingExportWindow -> AdvancedRenownTrainingExportWindow.OnHidden -> AdvancedRenownTraining.OnExportHidden

## Related APIs

- [ButtonSetDisabledFlag](../../window_api/functions/window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxSetDisabledFlag](../../window_api/functions/window_ComboBoxSetDisabledFlag.md) (HIGH 100/100) - Window Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none

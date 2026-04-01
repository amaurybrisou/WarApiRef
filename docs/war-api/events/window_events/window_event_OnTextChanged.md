# OnTextChanged

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
| Addons seen in | Aura, BuffHead, EA_UiDebugTools, Enemy, Killer, MapMonster, MapPin, Miracle Grow Remix |
| Files seen in | `/workspace/Aura/Source/AuraConfig.xml:1569`, `/workspace/Aura/Source/AuraConfig.xml:1591`, `/workspace/Aura/Source/AuraConfig.xml:177`, `/workspace/Aura/Source/AuraConfig.xml:199`, `/workspace/BuffHead/Setup/General.xml:156`, `/workspace/BuffHead/Setup/General.xml:196`, `/workspace/BuffHead/Setup/General.xml:235`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:128` |
| Namespaces detected | OnTextChanged |
| Source kinds | event_page, xml_handlers |
| Example locations | Aura: AuraConfigGeneralOffsetX.OnTextChanged, Aura: AuraConfigGeneralOffsetY.OnTextChanged, Aura: AuraConfigTimerOffsetX.OnTextChanged, Aura: AuraConfigTimerOffsetY.OnTextChanged, BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeColumnsEditBox.OnTextChanged, BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeRowsEditBox.OnTextChanged |
| XML usage count | 112 |
| XML attribute usage count | 112 |
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

Observed as an engine-supplied UI event hook used by 14 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- Aura
- BuffHead
- EA_UiDebugTools
- Enemy
- Killer
- MapMonster
- MapPin
- Miracle Grow Remix
- Pocket Palette
- RVAPI_ColorDialog
- Shinies
- TexturedButtons
- TurretRange
- wbLeadHelper

## Registrars And Handlers

- AuraConfig.OnTextureOffsetXChanged
- AuraConfig.OnTextureOffsetYChanged
- AuraConfig.OnTimerOffsetXChanged
- AuraConfig.OnTimerOffsetYChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerPaddingXChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerPaddingYChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeColumnsChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeRowsChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnIndicatorPaddingXChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnIndicatorPaddingYChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnMaximumThresholdChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnMinimumThresholdChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnOffsetXChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnOffsetYChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPlacementFixedOffsetXChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPlacementFixedOffsetYChanged
- BuffHead.Setup.Display.OnOffsetXChanged
- BuffHead.Setup.Display.OnOffsetYChanged
- BuffHead.Setup.EffectCache.OnSearchChanged
- BuffHead.Setup.Layout.Properties.OnCoreSizeSizeHeightChanged
- BuffHead.Setup.Layout.Properties.OnCoreSizeSizeWidthChanged
- BuffHead.Setup.Layout.Properties.OnOffsetXChanged
- BuffHead.Setup.Layout.Properties.OnOffsetYChanged
- BuffHead.Setup.Layout.Properties.OnSizeScaleChanged
- BuffHead.Setup.Layout.Properties.OnSizeSizeHeightChanged
- BuffHead.Setup.Layout.Properties.OnSizeSizeWidthChanged
- BuffHead.Setup.PriorityEffectsItem.OnAbilityIdChanged
- BuffHead.Setup.SelectColor.OnTintChanged
- DebugWindow.AutoSender
- DebugWindow.PreventType
- DevPadWindow.OnCodeChanged
- Enemy.ConfigurationWindow_OnChange
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Killer.OnSettingsEditChanged
- MapMonster.Editor.OnLabelChange
- MapMonster.Editor.OnNoteChange
- MapMonster.Editor.OnPosTextChanged
- MapMonster.PinTypeEditor.OnPinTypeTextChange
- MapMonster.PinTypeEditor.OnRadiusTextChange
- MapMonster.PinTypeEditor.OnSubTypeTextChange
- MapPin.TimeChanged
- MiracleGrow2.ConfigSoundChanged
- MiracleGrow2.ConfigThrobCChanged
- MiracleGrow2.LayoutBarCChanged
- MiracleGrow2.LayoutDimChanged
- MiracleGrow2.LayoutProgDimChanged
- PP.UpdateDyeFilter
- RVAPI_ColorDialog.OnTextChangedEdit
- ShiniesAutoUI.OnPriceChange
- ShiniesConfigGeneral.OnTextChanged_UIScale
- ShiniesPostUI.OnPriceChange
- ShiniesPostUI.OnStackChange
- ShiniesPostUI.OnStackSizeChange
- TexturedButtons.Setup.Actionbar.OnPaddingXChanged
- TexturedButtons.Setup.Actionbar.OnPaddingYChanged
- TexturedButtons.Setup.Actionbar.OnSpacingXChanged
- TexturedButtons.Setup.Actionbar.OnSpacingYChanged
- TexturedButtons.Setup.SelectColor.OnTintChanged
- TexturedButtons.Setup.Tint.OnTintChanged
- TurretRange.Setup.Display.OnDistanceOffsetXChanged
- TurretRange.Setup.Display.OnDistanceOffsetYChanged
- TurretRange.Setup.Display.OnGraphicLimitChanged
- TurretRange.Setup.Display.OnTintChanged
- TurretRange.Setup.Distance.OnTintChanged
- wbLeadHelperConfigTab.OnChanged

## Examples

- Aura: AuraConfigGeneralOffsetX -> AuraConfigGeneralOffsetX.OnTextChanged -> AuraConfig.OnTextureOffsetXChanged
- Aura: AuraConfigGeneralOffsetY -> AuraConfigGeneralOffsetY.OnTextChanged -> AuraConfig.OnTextureOffsetYChanged
- Aura: AuraConfigTimerOffsetX -> AuraConfigTimerOffsetX.OnTextChanged -> AuraConfig.OnTimerOffsetXChanged
- Aura: AuraConfigTimerOffsetY -> AuraConfigTimerOffsetY.OnTextChanged -> AuraConfig.OnTimerOffsetYChanged
- BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeColumnsEditBox -> BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeColumnsEditBox.OnTextChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeColumnsChanged
- BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeRowsEditBox -> BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeRowsEditBox.OnTextChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeRowsChanged

## Related APIs

- [TextEditBoxSetTextColor](../../window_api/functions/window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function

## Used With

- [EditBox](../../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [OnTextChanged](../../xml/handlers/handler_OnTextChanged.md) (HIGH 100/100) - XML Handler
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](../../window_api/functions/window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none

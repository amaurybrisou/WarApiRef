# Label

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, AutoBand, BankArkel, BuffHead, Busted |
| Files seen in | `/workspace/AdvancedPetAssist/APAGui.xml:1012`, `/workspace/AdvancedPetAssist/APAGui.xml:1034`, `/workspace/AdvancedPetAssist/APAGui.xml:1056`, `/workspace/AdvancedPetAssist/APAGui.xml:1068`, `/workspace/AdvancedPetAssist/APAGui.xml:1090`, `/workspace/AdvancedPetAssist/APAGui.xml:1102`, `/workspace/AdvancedPetAssist/APAGui.xml:1114`, `/workspace/AdvancedPetAssist/APAGui.xml:1126` |
| Namespaces detected | Label |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDLabel, AdvancedPetAssist: APAInstantOnlyHUDLabel, AdvancedPetAssist: APAKitingHUDLabel, AdvancedPetAssist: APALabelAttackBind, AdvancedPetAssist: APALabelAutoReattack, AdvancedPetAssist: APALabelAutoReattackDelay |
| XML usage count | 1606 |
| XML attribute usage count | 1606 |
| Lua usage count | 7 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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

Observed XML element type instantiated by 66 addons.

## Common Attributes

- name
- textalign
- font
- inherits
- handleinput
- autoresize
- maxchars
- warnOnTextCropped
- wordwrap
- layer
- popable
- skipinput

## Common Handlers

- OnMouseOver
- OnLButtonUp
- OnRButtonUp
- OnMouseOverEnd
- OnHyperLinkLButtonUp
- OnHyperLinkRButtonUp
- OnLButtonDown

## Common Inherits

- EA_Label_DefaultText
- Aura_Default_Label_SmallFont
- ClosetGoblinTalismanLabel
- Shinies_Default_Label_ClearSmallFont
- ModDetailsTextDef
- EA_Label_DefaultText_Small
- ModDetailsLabelDef
- DefaultWindowSmallText
- EA_Settings_ItemTitle
- MapMonsterEditorWindowLabelDefault
- DefaultWindowText
- IraConfigHeading

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- AutoBand
- BankArkel
- BuffHead
- Busted
- CM_ClosetGoblin
- CMap
- Cheeseboard
- Clock
- CombatTextNames
- Crafting Info Tooltip
- DAoCBuff
- DaemonAssist
- DeepSleep
- EA_UiDebugTools
- EA_UiModWindow
- Effigy
- Enemy
- FastInteract
- GetStats
- GoldTracker
- GuardLine
- GuardList
- GuardRange
- JunkDump
- Killer
- LibGroup
- LibWBToggler
- LoyalPet
- MapMonster
- MapPin
- MegaphonePlusPlus
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- Moth
- PeaceOut
- Pocket Palette
- PotionBar
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyQueue
- TidyRoll
- TimeInQueue
- Tortall_DPS
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAFollowTargetHUDLabel -> Label APAFollowTargetHUDLabel
- AdvancedPetAssist: APAInstantOnlyHUDLabel -> Label APAInstantOnlyHUDLabel
- AdvancedPetAssist: APAKitingHUDLabel -> Label APAKitingHUDLabel
- AdvancedPetAssist: APALabelAttackBind -> Label APALabelAttackBind
- AdvancedPetAssist: APALabelAutoReattack -> Label APALabelAutoReattack
- AdvancedPetAssist: APALabelAutoReattackDelay -> Label APALabelAutoReattackDelay

## Related APIs

- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function

## Used With

- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [OnHyperLinkLButtonUp](../../events/window_events/window_event_OnHyperLinkLButtonUp.md) (HIGH 100/100) - Window Event
- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnHyperLinkRButtonUp](../../events/window_events/window_event_OnHyperLinkRButtonUp.md) (HIGH 100/100) - Window Event
- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOverEnd](../../events/window_events/window_event_OnMouseOverEnd.md) (HIGH 100/100) - Window Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

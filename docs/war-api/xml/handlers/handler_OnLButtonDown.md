# OnLButtonDown

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 88/100

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, referenced by generated docs or reference files, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, BuffHead, CM_ClosetGoblin, CMap, Calling, Crusher, DaemonAssist, DetauntHelper |
| Namespaces detected | OnLButtonDown |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedRenownTrainer: .OnLButtonDown, BuffHead: .OnLButtonDown, CM_ClosetGoblin: .OnLButtonDown, CMap: .OnLButtonDown, Calling: .OnLButtonDown, Crusher: .OnLButtonDown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 161 |
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

XML handler event observed across 48 addons.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- none

## Seen In

- AdvancedRenownTrainer
- BuffHead
- CM_ClosetGoblin
- CMap
- Calling
- Crusher
- DaemonAssist
- DetauntHelper
- EA_LoadingScreen
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- Effigy
- Enemy
- FozAuction
- GetStats
- GuildWarden
- HealGrid
- Hopper
- LibAddonButton
- MarkBuff
- Minmap
- Miracle Grow Remix
- MoraleCircle
- Motion
- NaturalLog
- Obsidian
- PotionBar
- Pure
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RVMOD_SquaredDistances
- Sequencer
- SocialWindow 2.0
- Squared
- Targets
- TexturedButtons
- Tome Titan
- TurretRange
- Twister
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- XpStatus+G
- bigger_MacroWindow
- emotes
- nLootLink

## Examples

- AdvancedRenownTrainer: .OnLButtonDown -> AdvancedRenownTraining.Select
- BuffHead: .OnLButtonDown -> BuffHead.Setup.SelectTexture.OnTextureRowLDown
- BuffHead: .OnLButtonDown -> BuffHead.Setup.AdvancedCompression.OnRowLDown
- BuffHead: .OnLButtonDown -> BuffHead.Setup.AdvancedCompressionItem.OnRowLDown
- BuffHead: .OnLButtonDown -> BuffHead.Setup.AdvancedContainers.OnRowLDown
- BuffHead: .OnLButtonDown -> BuffHead.Setup.EffectCache.OnRowLDown

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [CircleImage](../element_types/element_CircleImage.md) (HIGH 100/100) - XML Element Type
- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [DynamicImage](../element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionTrigger](../../window_api/functions/window_WindowSetGameActionTrigger.md) (HIGH 100/100) - Window Function
- [BroadcastEvent](../../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function
- [WindowGameAction](../../window_api/functions/window_WindowGameAction.md) (HIGH 80/100) - Window Function

## Affects

- [GameData.TargetType.FRIENDLY](../../gamedata/fields/gamedata_GameData.TargetType.FRIENDLY.md) (HIGH 100/100) - GameData Field
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.TARGET_SELF](../../systemdata/fields/systemdata_SystemData.Events.TARGET_SELF.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected binding arguments remain uncertain because contract artifacts capture symbol linkage, not full handler signatures.

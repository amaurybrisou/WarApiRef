# WindowGetId

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 88 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, AdvancedRenownTrainer, AggroMeter, Atlas, Aura, BagOMatic, BankArkel, BlackBook |
| Files seen in | AdvancedRenownTraining.lua, AggroMeter.lua, BagOMatic.lua, BankArkel.lua, BlackBookWindow.lua, CCTV.lua, CDownSettings.lua, ClosetGoblinCharacterWindow.lua |
| Namespaces detected | WindowGetId |
| Source kinds | lua_calls |
| Example locations | Ace: GetId, AdvancedRenownTrainer: AbilityTooltip, AdvancedRenownTrainer: OnLButtonUpTab, AdvancedRenownTrainer: Select, AggroMeter: OnTabLBU, AggroMeter: PickedListMenu |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 385 |
| Global usage count | 385 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
WindowGetId(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: ButtonName, EA_Window_ContextMenu.activeWindow, ParentRow |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- AdvancedRenownTrainer
- AggroMeter
- Atlas
- Aura
- BagOMatic
- BankArkel
- BlackBook
- BuffHead
- CCTV
- CDown
- CM_ClosetGoblin
- CastSequence
- Crusher
- DAoCBuff
- DPSMeter
- DammazKron
- DetauntHelper
- DuffTimer
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- Emojii
- Enemy
- GCDsaver
- HealGrid
- HideHiddenFrames
- Hopper
- InfoScroller
- JunkDump
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- LibWBToggler
- LootAlert
- MacroIcons
- Map
- MapMonster
- MapPin
- MarkBuff
- Minmap
- Miracle Grow Remix
- Motion
- NerfedButtons
- Obsidian
- PartyCast
- Pocket Palette
- Pure
- Pure Careerbar
- QuickTacticSwitch
- RVMOD_Manager
- RVMOD_Targets
- RaidMeter
- RealmStatus
- Refer
- ResHelp
- RoR_SoR
- Sequencer
- SessionRPs
- Shinies
- Statdoll
- TacticSetNames
- TargetRing
- TastyButtons
- TexturedButtons
- TidyChat
- TidyQueue
- TidyRoll
- TokenMachine
- Tokens
- Tome Titan
- TomeTracker
- TurretRange
- Vectors
- WarBoard_Menu
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- minesweep
- nLootLink
- nRarity
- scenarioInfo
- wbLeadHelper
- zMailMod

## Examples

- Ace: GetId -> WindowGetId(self.name)
- AdvancedRenownTrainer: AbilityTooltip -> WindowGetId(SystemData.MouseOverWindow.name)
- AdvancedRenownTrainer: OnLButtonUpTab -> WindowGetId(SystemData.ActiveWindow.name)
- AdvancedRenownTrainer: Select -> WindowGetId(activeWindow)
- AggroMeter: OnTabLBU -> WindowGetId(SystemData.ActiveWindow.name)
- AggroMeter: PickedListMenu -> WindowGetId(SystemData.MouseOverWindow.name)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnSlide](../../xml/handlers/handler_OnSlide.md) (HIGH 88/100) - XML Event
- [OnTextChanged](../../xml/handlers/handler_OnTextChanged.md) (HIGH 88/100) - XML Event

## Used With

- [EA_Window_ContextMenu.HideAll](../../globals/functions/global_EA_Window_ContextMenu.HideAll.md) (HIGH 100/100) - Global Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [WindowGetParent](window_WindowGetParent.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- none

# ComboBoxAddMenuItem

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 71 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, BankArkel, BuffHead, Busted, CCTV, CDown |
| Files seen in | BankArkel.lua, Busted.lua, CCTV.lua, CDownSettings.lua, ChattyCathy.lua, Code/Core/Groups/EnemyEffectFilter.lua, Code/Intercom/Intercom.lua, Code/UnitFrames/ClickCasting.lua |
| Namespaces detected | ComboBoxAddMenuItem |
| Source kinds | lua_calls |
| Example locations | Ace: Add, ActionBarHide: Add, Amethyst: Add, BankArkel: SetupCombos, BuffHead: Initialize, BuffHead: LoadEffectsSettings |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 587 |
| Global usage count | 587 |
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
ComboBoxAddMenuItem(arg1, arg2)
```

## Description

Observed as a window function across 71 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BankArkelBackpackCombo", "BustedGUIAddonSelect", "CCTVSettingsCombobox4" |
| arg2 | Observed as a function or method reference. | Observed values: BankArkel.db.Entry[i].Name, CraftValueTip.GetPhrase("config","langdefault"), GameData.Player.name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- BankArkel
- BuffHead
- Busted
- CCTV
- CDown
- CMap
- CaVES
- CastSequence
- ChattyCathy
- Crafting Info Tooltip
- Crusher
- DAoCBuff
- DPSMeter
- EA_OpenPartyWindow
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- Enemy
- FozAuction
- GCDsaver
- GroupRange
- GuildWarden
- Hopper
- InfoScroller
- KeysetMonsterPlay
- Killer
- LibAddonButton
- LibWBToggler
- Map
- MarkBuff
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- Obsidian
- OverheadFonts
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- RealmStatus
- SOR
- ScenarioStats
- Shinies
- SocialWindow 2.0
- Statdoll Remix
- TalismanGenie
- TargetRing
- TastyButtons
- TexturedButtons
- ThinkOutLoud
- TidyRoll
- TokenMachine
- Tokens
- TurretRange
- TwisterSet
- WBStutterLess
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- ZonePOP
- alertMod
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: Add -> ComboBoxAddMenuItem(self.name, itemText)
- ActionBarHide: Add -> ComboBoxAddMenuItem(self.name, itemText)
- Amethyst: Add -> ComboBoxAddMenuItem(self.name, itemText)
- BankArkel: SetupCombos -> ComboBoxAddMenuItem("BankArkelBackpackCombo", StringToWString(tbNone))
- BankArkel: SetupCombos -> ComboBoxAddMenuItem("BankArkelBackpackCombo", StringToWString(tbSelf))
- BankArkel: SetupCombos -> ComboBoxAddMenuItem("BankArkelBackpackCombo", BankArkel.db.Entry[i].Name)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event

## Used With

- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function

## Notes

- none

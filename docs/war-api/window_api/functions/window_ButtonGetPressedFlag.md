# ButtonGetPressedFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 89 addons

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
| Addons seen in | Ace, ActionBarHide, AdjustTheTip, Amethyst, Atlas, AuctionStats, Bloody Mess, BuffHead |
| Files seen in | AAOTracker.lua, AdjustTheTip.lua, AuctionAssist.lua, Bloody Mess.lua, CDownSettings.lua, CMap.lua, CallingSetup.lua, Code/Core/ConfigurationWindow.lua |
| Namespaces detected | ButtonGetPressedFlag |
| Source kinds | lua_calls |
| Example locations | Ace: GetValue, ActionBarHide: GetValue, AdjustTheTip: OnDelayedFadingEnabled, Amethyst: GetValue, Atlas: OnCheckboxLBU, AuctionStats: AutoUndercutClicked |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 377 |
| Global usage count | 377 |
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
ButtonGetPressedFlag(arg1)
```

## Description

Observed as a window function across 89 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "AuctionHouseCreateAuctionAutoUndercutCheckboxButton", "AuctionHouseCreateAuctionSetBuyOutCheckboxButton", "BloodyMessOptionsEnableCheckBoxButton" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionBarHide
- AdjustTheTip
- Amethyst
- Atlas
- AuctionStats
- Bloody Mess
- BuffHead
- CDown
- CMap
- CaVES
- Calling
- CastSequence
- Crafting Info Tooltip
- Crusher
- DAoCBuff
- DammazKron
- DetauntHelper
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiModWindow
- EZCraftX
- EZGuard
- Effigy
- Enemy
- EveryBodyGuard
- FastFriends
- GCDsaver
- Group Icons SG
- GroupRange
- GroupSpotter
- HealGrid
- Hopper
- InfoScroller
- KeyBar
- Keyset
- KeysetMonsterPlay
- Killer
- LibAddonButton
- LibGroup
- LibWBToggler
- Map
- MapMonster
- MarkBuff
- MegaphonePlusPlus
- Minmap
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- Obsidian
- PartyCast
- Phantom
- PieTracker
- PotionBar
- Pure
- Pure Careerbar
- RealmStatus
- ReliquaryHunter
- Shinies
- SocialWindow 2.0
- Statdoll Remix
- TalismanGenie
- TargetRing
- TastyButtons
- TaxPayer
- TexturedButtons
- TidyChat
- TidyQueue
- TidyRoll
- Tokens
- TurretRange
- Vectors
- WARCommander
- WSCT
- WTes
- WarBoard_AAOTracker
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- nLootLink
- nRarity
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: GetValue -> ButtonGetPressedFlag(self.name)
- ActionBarHide: GetValue -> ButtonGetPressedFlag(self.name)
- AdjustTheTip: OnDelayedFadingEnabled -> ButtonGetPressedFlag(item.Data.Name.."CheckBox")
- Amethyst: GetValue -> ButtonGetPressedFlag(self.name)
- Atlas: OnCheckboxLBU -> ButtonGetPressedFlag(checkboxName)
- AuctionStats: AutoUndercutClicked -> ButtonGetPressedFlag("AuctionHouseCreateAuctionAutoUndercutCheckboxButton")

## Used With

- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [EA_LabelCheckButton.Toggle](../../globals/functions/global_EA_LabelCheckButton.Toggle.md) (HIGH 100/100) - Global Function
- [LayoutEditor.Hide](window_LayoutEditor.Hide.md) (HIGH 90/100) - Window Function
- [LayoutEditor.Show](window_LayoutEditor.Show.md) (HIGH 90/100) - Window Function

## Affects

- [SystemData.MouseOverWindow.id](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.id.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- none

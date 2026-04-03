# ButtonSetCheckButtonFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 46 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, AuctionStats, AutoSalvage, BuddyBind, CMap, Calling |
| Files seen in | BuddyBind.lua, CMap.lua, CallingKeybinding.lua, KillerConfigWindow.lua, LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua |
| Namespaces detected | ButtonSetCheckButtonFlag |
| Source kinds | lua_calls |
| Example locations | Ace: New, ActionBarHide: New, Amethyst: New, AuctionStats: ExitOnComplete, AuctionStats: SetCheckBox, AuctionStats: ToggleCheckBox |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 118 |
| Global usage count | 118 |
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
ButtonSetCheckButtonFlag(arg1, arg2)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "EA_Window_MacroIconSlot"..slot, "PhantomSettingsHideBarLock", "PhantomSettingsHideGroupBuffs" |
| arg2 | Observed as a boolean toggle. | Observed values: Status, false, status |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- AuctionStats
- AutoSalvage
- BuddyBind
- CMap
- Calling
- Crusher
- EA_ScenarioGroupWindow
- EA_UiModWindow
- EZCraftX
- EZGuard
- Effigy
- GCDsaver
- Hopper
- InfoScroller
- Killer
- LibWBToggler
- Map
- MapMonster
- Motion
- NaturalLog
- PartyCast
- Phantom
- Pure
- Pure Careerbar
- RealmStatus
- Shinies
- SocialWindow 2.0
- TalismanGenie
- TargetRing
- TastyButtons
- Tokens
- WARCommander
- WBStutterLess
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- bigger_MacroWindow
- nLootLink
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: New -> ButtonSetCheckButtonFlag(w.name, true)
- ActionBarHide: New -> ButtonSetCheckButtonFlag(w.name, true)
- Amethyst: New -> ButtonSetCheckButtonFlag(w.name, true)
- AuctionStats: ExitOnComplete -> ButtonSetCheckButtonFlag(checkBoxName, Status)
- AuctionStats: SetCheckBox -> ButtonSetCheckButtonFlag(parentControls..name, status)
- AuctionStats: ToggleCheckBox -> ButtonSetCheckButtonFlag(checkBoxName, Status)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event

## Used With

- [ButtonGetCheckButtonFlag](window_ButtonGetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetStayDownFlag](window_ButtonSetStayDownFlag.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.M_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.M_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.R_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.R_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- none

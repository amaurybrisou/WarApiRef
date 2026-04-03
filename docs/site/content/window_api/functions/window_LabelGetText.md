# LabelGetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 43 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, Countdown, Crusher, DAoCBuff, Dascore, EA_UiDebugTools |
| Files seen in | DascoreResu.lua, Fixes.lua, GuildWardenWin.lua, LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, QueueQueuer.lua |
| Namespaces detected | LabelGetText |
| Source kinds | lua_calls |
| Example locations | Ace: GetText, ActionBarHide: GetText, Amethyst: GetText, Countdown: displayCountdown, Crusher: GetText, DAoCBuff: FilterRowOnLButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 67 |
| Global usage count | 67 |
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
LabelGetText(arg1)
```

## Description

Observed as a window function across 43 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "EA_Window_ScenarioJoinPromptBoxName", "GuildWarden.DBWindow.TextN5", "ScenarioSummaryWindowDestructionPoints" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- Countdown
- Crusher
- DAoCBuff
- Dascore
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- FixGit
- GCDsaver
- GuildWarden
- Hopper
- InfoScroller
- LibWBToggler
- Map
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- Queue Queuer
- RealmStatus
- RoR_SoR
- SOR
- Shinies
- Soloq
- TargetRing
- Targets
- ThankTheResser
- Tokens
- Trakario
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- nLootLink
- scenarioInfo
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: GetText -> LabelGetText(self.name)
- ActionBarHide: GetText -> LabelGetText(self.name)
- Amethyst: GetText -> LabelGetText(self.name)
- Countdown: displayCountdown -> LabelGetText(label)
- Crusher: GetText -> LabelGetText(self.name)
- DAoCBuff: FilterRowOnLButtonUp -> LabelGetText(SystemData.ActiveWindow.name.."Name")

## Related APIs

- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- none

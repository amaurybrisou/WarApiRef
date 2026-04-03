# LabelGetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 48 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, ActionBarHide, Amethyst, CMap, Countdown, Crusher, DAoCBuff, Dascore |
| Files seen in | DascoreResu.lua, Fixes.lua, GuardList.lua, GuardRange.lua, GuildWardenWin.lua, LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua |
| Namespaces detected | LabelGetText |
| Source kinds | lua_calls |
| Example locations | Ace: GetText, ActionBarHide: GetText, Amethyst: GetText, CMap: GetText, Countdown: displayCountdown, Crusher: GetText |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 74 |
| Global usage count | 74 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
LabelGetText(arg1)
```

## Description

Observed as a window function across 48 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "EA_Window_ScenarioJoinPromptBoxName", "GuardList_Window0Label", "GuardRange_Window0Label" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- CMap
- Countdown
- Crusher
- DAoCBuff
- Dascore
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- FixGit
- FozAuction
- GCDsaver
- GuardList
- GuardRange
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
- WarBoard_WarWhisperer
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
- CMap: GetText -> LabelGetText(self.name)
- Countdown: displayCountdown -> LabelGetText(label)
- Crusher: GetText -> LabelGetText(self.name)

## Related APIs

- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function

## Notes

- none

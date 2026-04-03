# ButtonSetTextColor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 36 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, CMap, Crusher, EZCraftX, EZGuard, Effigy |
| Files seen in | LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, TTitan_TTButton.lua, TTitan_UI.lua, TastyButtons.lua, TidyChat.lua |
| Namespaces detected | ButtonSetTextColor |
| Source kinds | lua_calls |
| Example locations | Ace: TextColor, ActionBarHide: TextColor, Amethyst: TextColor, CMap: TextColor, Crusher: TextColor, EZCraftX: TextColor |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 58 |
| Global usage count | 58 |
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
ButtonSetTextColor(arg1, arg2, arg3, arg4)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "TTitanSplashTextBestiaryClick", "TTitanSplashTextHistoryClick", "TTitanSplashTextNoteworthyClick" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 0, Button.ButtonState.HIGHLIGHTED, Button.ButtonState.NORMAL |
| arg3 | Observed as a runtime window or control identifier. | Observed values: 0, 255, 84 |
| arg4 | Observed as a runtime window or control identifier. | Observed values: 0, 137, 255 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- CMap
- Crusher
- EZCraftX
- EZGuard
- Effigy
- GCDsaver
- Hopper
- InfoScroller
- LibWBToggler
- Map
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- Shinies
- TacticSetNames
- TargetRing
- TastyButtons
- TidyChat
- Tokens
- Tome Titan
- TwisterSet
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: TextColor -> ButtonSetTextColor(self.name, red, green, blue)
- ActionBarHide: TextColor -> ButtonSetTextColor(self.name, red, green, blue)
- Amethyst: TextColor -> ButtonSetTextColor(self.name, red, green, blue)
- CMap: TextColor -> ButtonSetTextColor(self.name, red, green, blue)
- Crusher: TextColor -> ButtonSetTextColor(self.name, red, green, blue)
- EZCraftX: TextColor -> ButtonSetTextColor(self.name, red, green, blue)

## Used With

- [ButtonGetTextColor](window_ButtonGetTextColor.md) (HIGH 100/100) - Window Function

## Notes

- none

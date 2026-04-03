# WindowSetMovable

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
| Addons seen in | Ace, ActionBarHide, ActionFraction, Amethyst, Atlas, CMap, CharacterScreenTabFix, Crusher |
| Files seen in | CharacterScreenTabFix.lua, Code/Core/Main.lua, Code/Guard/Guard.lua, JunkDumpOptions.lua, LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua |
| Namespaces detected | WindowSetMovable |
| Source kinds | lua_calls |
| Example locations | Ace: MakeFixed, Ace: MakeMovable, ActionBarHide: MakeFixed, ActionBarHide: MakeMovable, ActionFraction: OnLock, ActionFraction: OnUnlock |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 90 |
| Global usage count | 90 |
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
WindowSetMovable(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "AbilitiesWindow", "ApothecaryWindow", "AtlasFrame" |
| arg2 | Observed as a boolean toggle. | Observed values: (...), (settings.Movable==true), WarWhisperer.Settings.moveable |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- Amethyst
- Atlas
- CMap
- CharacterScreenTabFix
- Crusher
- EZCraftX
- EZGuard
- Effigy
- Enemy
- GCDsaver
- GroupRange
- Hopper
- InfoScroller
- JunkDump
- LibWBToggler
- Map
- MapPin
- Motion
- NaturalLog
- NoOverheal
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- RealmStatus
- Shinies
- Statdoll Remix
- TargetRing
- TastyButtons
- TexturedButtons
- Tokens
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WindowMovers
- WoH-Reticle
- XpStatus+G
- emotes
- nLootLink
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: MakeFixed -> WindowSetMovable(self.name, false)
- Ace: MakeMovable -> WindowSetMovable(self.name, true)
- ActionBarHide: MakeFixed -> WindowSetMovable(self.name, false)
- ActionBarHide: MakeMovable -> WindowSetMovable(self.name, true)
- ActionFraction: OnLock -> WindowSetMovable(windowName, false)
- ActionFraction: OnUnlock -> WindowSetMovable(windowName, true)

## Notes

- none

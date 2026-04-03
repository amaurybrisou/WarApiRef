# WindowSetParent

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 56 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, BBars - Mechanic Only, BagOMatic, BetterCC, CM_ClosetGoblin, CMap |
| Files seen in | BBarsPetHP.lua, BBarsPlayerMechanic.lua, BagOMatic.lua, BetterCC.lua, ClosetGoblin.lua, CoolDownLine.lua, Core.lua, Dye.lua |
| Namespaces detected | WindowSetParent |
| Source kinds | lua_calls |
| Example locations | Ace: Parent, ActionBarHide: Parent, Amethyst: Parent, BBars - Mechanic Only: MechDraw, BBars - Mechanic Only: PetHPDraw, BagOMatic: init |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 92 |
| Global usage count | 92 |
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
WindowSetParent(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "BBarsPetHPBack", "BBarsPetHPFront", "BBarsPlayerMechanic1Frontbar" |
| arg2 | Observed as a text or wstring payload. | Observed values: "BBarsPetHPBG", "BBarsPlayerMechanic1BG", "BBarsPlayerMechanic2BG" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- BBars - Mechanic Only
- BagOMatic
- BetterCC
- CM_ClosetGoblin
- CMap
- CoolDownLine
- Crusher
- Dye Preview
- EZCraftX
- EZGuard
- Effigy
- GCDsaver
- GetStats
- Hopper
- InfoScroller
- JunkDump
- LibWBToggler
- MacroIcons
- Map
- MapMonster
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RVMOD_Targets
- RealmStatus
- RoR_debolster
- RvRStats
- RvRStatsTab
- Shinies
- TargetRing
- TidyChat
- TokenMachine
- Tokens
- TurretRange
- WarBoard_Menu
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- bigger_MacroWindow
- nRarity
- scenarioInfo
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Parent -> WindowSetParent(self.name, self.parent)
- ActionBarHide: Parent -> WindowSetParent(self.name, self.parent)
- Amethyst: Parent -> WindowSetParent(self.name, self.parent)
- BBars - Mechanic Only: MechDraw -> WindowSetParent("BBarsPlayerMechanic1Grey", "BBarsPlayerMechanic1BG")
- BBars - Mechanic Only: MechDraw -> WindowSetParent("BBarsPlayerMechanic2Grey", "BBarsPlayerMechanic2BG")
- BBars - Mechanic Only: MechDraw -> WindowSetParent("BBarsPlayerMechanic3Grey", "BBarsPlayerMechanic3BG")

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Notes

- none

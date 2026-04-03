# ScrollWindowSetOffset

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 11 addons

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
| Addons seen in | Crusher, Enemy, Hopper, Motion, Pure, RVMOD_Manager, Shinies, Tome Titan |
| Files seen in | Code/Core/Main.lua, Code/UnitFrames/ClickCasting.lua, Code/UnitFrames/EffectsIndicator.lua, Code/UnitFrames/UnitFramePart.lua, Code/UnitFrames/UnitFrames.lua, Configuration/Config.lua, Configuration/HopperConfig.lua, Configuration/PureConfig.lua |
| Namespaces detected | ScrollWindowSetOffset |
| Source kinds | lua_calls |
| Example locations | Crusher: UpdateDisplayedConfig, Enemy: UI_ConfigDialog_Open, Enemy: UnitFramesUI_ConfigDialog_Import, Enemy: UnitFramesUI_EffectsIndicatorDialog_Open, Enemy: UnitFramesUI_UnitFrameClickCastingDialog_Open, Enemy: UnitFramesUI_UnitFramePartDialog_Open |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 18 |
| Global usage count | 18 |
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
ScrollWindowSetOffset(arg1, arg2)
```

## Description

Observed as a window function across 11 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: "EnemyClickCastingDialogContent", "EnemyConfigDialogSectionsUnitFramesContent", "EnemyEffectsIndicatorDialogContent" |
| arg2 | Observed as a numeric value. | Observed values: 0, scroll |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Crusher
- Enemy
- Hopper
- Motion
- Pure
- RVMOD_Manager
- Shinies
- Tome Titan
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- wbLeadHelper

## Examples

- Crusher: UpdateDisplayedConfig -> ScrollWindowSetOffset(configWindowScroll, 0)
- Enemy: UI_ConfigDialog_Open -> ScrollWindowSetOffset(config_dlg.currentSection.windowName.."Content", scroll)
- Enemy: UnitFramesUI_ConfigDialog_Import -> ScrollWindowSetOffset("EnemyConfigDialogSectionsUnitFramesContent", 0)
- Enemy: UnitFramesUI_EffectsIndicatorDialog_Open -> ScrollWindowSetOffset("EnemyEffectsIndicatorDialogContent", 0)
- Enemy: UnitFramesUI_UnitFrameClickCastingDialog_Open -> ScrollWindowSetOffset("EnemyClickCastingDialogContent", 0)
- Enemy: UnitFramesUI_UnitFramePartDialog_Open -> ScrollWindowSetOffset("EnemyUnitFramePartDialogContent", 0)

## Notes

- none

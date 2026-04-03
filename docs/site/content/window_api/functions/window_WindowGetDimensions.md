# WindowGetDimensions

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 84 addons

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
| Addons seen in | Ace, AdjustTheTip, AdvancedPetAssist, Amethyst, AnywhereTrainer, AuctionStats, Aura, BuffHead |
| Files seen in | APAGuiHUD.lua, AdjustTheTip.lua, Amethyst.lua, AuctionStats.lua, Button.lua, Classes/Display.lua, Code/GroupIcons/GroupIcon.lua, Code/UnitFrames/EffectsIndicator.lua |
| Namespaces detected | WindowGetDimensions |
| Source kinds | lua_calls |
| Example locations | Ace: New, AdjustTheTip: AddTargetHealthToMouseOver, AdjustTheTip: UpdateCallback, AdvancedPetAssist: ApplyPetTargetHUDLayout, Amethyst: New, Amethyst: SavePosition |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 471 |
| Global usage count | 471 |
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
WindowGetDimensions(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "APAPetTargetHUD", "AbilityTooltipBackground", "CharacterWindow" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- AdjustTheTip
- AdvancedPetAssist
- Amethyst
- AnywhereTrainer
- AuctionStats
- Aura
- BuffHead
- CaVES
- CastSequence
- Crafting Info Tooltip
- CraftingWillard
- Crusher
- DAoCBuff
- DammazKron
- DetauntHelper
- DuffTimer
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- Enemy
- EveryBodyGuard
- FlagCap
- GCDTracker
- GCDsaver
- GroupRange
- GroupSpotter
- HealGrid
- Hopper
- InfoScroller
- KeyBar
- Killer
- Kwestor
- LibAddonButton
- LibSurveyor
- LibWBToggler
- Map
- MapMonster
- MapPin
- MiniMapMonster
- Miracle Grow Remix
- MiracleGrowLight
- Moth
- Motion
- MouseHint
- NAMBLA
- NaturalLog
- NerfedButtons
- Obsidian
- PartyCast
- Pocket Palette
- PotionBar
- Pure
- Pure Careerbar
- RVAPI_ColorDialog
- RVMOD_PlayerStatus
- RealmStatus
- ReliquaryHunter
- RoR_SoR
- RoR_debolster
- RvRContribution
- SNT_PANEL
- Shinies
- TacticSetNames
- TargetRing
- TastyButtons
- TexturedButtons
- TidyChat
- Tokens
- TurretRange
- Twister
- VerticalMorale
- VerticalTactics
- WARCommander
- WARRatingBuster
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- nRarity
- scenarioInfo
- xHUD

## Examples

- Ace: New -> WindowGetDimensions(w.name)
- AdjustTheTip: AddTargetHealthToMouseOver -> WindowGetDimensions("MouseOverTargetUnitWindow")
- AdjustTheTip: UpdateCallback -> WindowGetDimensions("AbilityTooltipBackground")
- AdvancedPetAssist: ApplyPetTargetHUDLayout -> WindowGetDimensions("APAPetTargetHUD")
- Amethyst: New -> WindowGetDimensions(w.name)
- Amethyst: SavePosition -> WindowGetDimensions(C.name)

## Related APIs

- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [StatusBarSetBackgroundTint](window_StatusBarSetBackgroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetForegroundTint](window_StatusBarSetForegroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Notes

- Advanced return analysis: No strong return evidence observed

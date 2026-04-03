# WindowSetScale

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 82 addons

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
| Addons seen in | Ace, ActionFraction, Amethyst, Atlas, Aura, BBars - Mechanic Only, BetterCC, BuffHead |
| Files seen in | AdvancedContainers.lua, Amethyst.lua, BBarsPetHP.lua, BBarsPlayerMechanic.lua, Bars/HealGridProgressBar.lua, BetterCC.lua, CCTV.lua, CDownFrames.lua |
| Namespaces detected | WindowSetScale |
| Source kinds | lua_calls |
| Example locations | Ace: Scale, ActionFraction: ResetWindow, Amethyst: Recreate, Amethyst: Scale, Atlas: PopulateInterface, Atlas: UpdateLegend |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 294 |
| Global usage count | 294 |
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
WindowSetScale(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "AtlasConfigurationFrame", "AtlasFrameLegend", "BBarsPetHPBG" |
| arg2 | Observed as a function or method reference. | Observed values: (0.80+(GetBattlegroupMemberData()[i].players[j].healthPercent/100*0.20))*InterfaceCore.GetScale()*GroupIcons.PREF_SCALE, (0.80+(GetGroupData()[i].healthPercent/100*0.20))*InterfaceCore.GetScale()*GroupIcons.PREF_SCALE, (1-cd)*1.5 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionFraction
- Amethyst
- Atlas
- Aura
- BBars - Mechanic Only
- BetterCC
- BuffHead
- CCTV
- CDown
- Calling
- CastSequence
- CleanCastbar
- CleanUnitFrames
- Countdown
- Crusher
- DAoCBuff
- DetauntHelper
- DuffTimer
- EZCraftX
- EZGuard
- Effigy
- Enemy
- GCDsaver
- Group Icons
- Group Icons SG
- GroupRange
- GroupSpotter
- GuardBot
- GuardLine
- HealGrid
- Hopper
- InfoScroller
- Kwestor
- LibWBToggler
- Map
- MapMonster
- MapPin
- MarkBuff
- MiniMapMonster
- Motion
- MouseHint
- Obsidian
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- RO-Style Combat Text
- RVMOD_Manager
- RVMOD_SquaredDistances
- RealmStatus
- RoR_SoR
- SNT_BUTTONS
- SNT_INFO
- SOR
- ScenarioStats
- Shinies
- ShowHealthPercent
- Squared
- Statdoll
- Statdoll Remix
- TacticSetNames
- TargetInfoRing
- TargetRing
- TidyQueue
- Tokens
- TurretRange
- Vectors
- VerticalTactics
- WSCT
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WindowMovers
- WoH-Reticle
- XpStatus+G
- nLootLink
- nRarity
- scenarioInfo
- wbLeadHelper
- xHUD
- zMailMod

## Examples

- Ace: Scale -> WindowSetScale(self.name, scale)
- ActionFraction: ResetWindow -> WindowSetScale(windowName, InterfaceCore.GetScale())
- Amethyst: Recreate -> WindowSetScale(C.name, s.scale)
- Amethyst: Recreate -> WindowSetScale(C.Fill.name, s.scale)
- Amethyst: Recreate -> WindowSetScale(C.Name.name, s.scale)
- Amethyst: Recreate -> WindowSetScale(C.Timer.name, s.scale)

## Related APIs

- [Size](../../xml/element_types/element_Size.md) (HIGH 100/100) - XML Element Type
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function

## Notes

- Advanced return analysis: No strong return evidence observed

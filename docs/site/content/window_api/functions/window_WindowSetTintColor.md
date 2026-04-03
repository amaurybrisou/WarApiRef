# WindowSetTintColor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 104 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, Amethyst, Atlas, Aura, BBars - Mechanic Only, BuffHead |
| Files seen in | APAGuiHUD.lua, AdvancedRenownTraining.lua, Amethyst.lua, BBarsPetHP.lua, BBarsPlayerMechanic.lua, Bar.lua, Bars/HealGridProgressBar.lua, CCTV.lua |
| Namespaces detected | WindowSetTintColor |
| Source kinds | lua_calls |
| Example locations | Ace: Tint, AdvancedPetAssist: UpdateFollowTargetHUD, AdvancedPetAssist: UpdateInstantOnlyHUD, AdvancedPetAssist: UpdateKitingHUD, AdvancedPetAssist: UpdatePetTargetHUD, AdvancedRenownTrainer: SelectAdvantage |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 717 |
| Global usage count | 717 |
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
WindowSetTintColor(windowName, arg2, arg3, arg4)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "APAFollowTargetHUDFill", "APAInstantOnlyHUDFill", "APAKitingHUDFill" |
| arg2 | Observed as a function or method reference. | Observed values: 0, 100, 12 |
| arg3 | Observed as a numeric value. | Observed values: 0, 100, 110 |
| arg4 | Observed as a numeric value. | Observed values: 0, 10, 100 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- Amethyst
- Atlas
- Aura
- BBars - Mechanic Only
- BuffHead
- CCTV
- CDown
- CM_ClosetGoblin
- Calling
- CastSequence
- ChattyCathy
- CleanUnitFrames
- CraftingWillard
- Crusher
- DAoCBuff
- DetauntHelper
- DuffTimer
- EZCraftX
- EZGuard
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- FlagCap
- GCDsaver
- Group Icons
- Group Icons SG
- GroupRange
- GroupSpotter
- GuardBot
- GuardLine
- GuildWarden
- HealGrid
- Hopper
- InfoScroller
- KeyBar
- Killer
- LibAddonButton
- LibWBToggler
- Map
- MapMonster
- MapPin
- MarkBuff
- MiniMapMonster
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MoraleCircle
- Moth
- Motion
- NaturalLog
- NerfedButtons
- Obsidian
- PartyCast
- PeaceOut
- Pocket Palette
- PotionBar
- Pure
- Pure Careerbar
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RandomMount
- RealmStatus
- ResHelp
- RoR_SoR
- SNT_INFO
- SOR
- Sequencer
- Shinies
- Squared
- Swinger
- TacticSetNames
- TargetInfoRing
- TargetRing
- TastyButtons
- TexturedButtons
- TidyChat
- TidyRoll
- TokenMachine
- Tokens
- TomeTracker
- TurretRange
- Twister
- Vectors
- WARCommander
- WSCT
- WarTriage
- WhatsCooking
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- minesweep
- nRarity
- scenarioInfo
- wbLeadHelper
- xHUD
- zMailMod

## Examples

- Ace: Tint -> WindowSetTintColor(self.name, r, g, b)
- AdvancedPetAssist: UpdateFollowTargetHUD -> WindowSetTintColor("APAFollowTargetHUDFill", APA.hudColorOnR, APA.hudColorOnG, APA.hudColorOnB)
- AdvancedPetAssist: UpdateFollowTargetHUD -> WindowSetTintColor("APAFollowTargetHUDFill", APA.hudColorOffR, APA.hudColorOffG, APA.hudColorOffB)
- AdvancedPetAssist: UpdateInstantOnlyHUD -> WindowSetTintColor("APAInstantOnlyHUDFill", 128, 128, 128)
- AdvancedPetAssist: UpdateInstantOnlyHUD -> WindowSetTintColor("APAInstantOnlyHUDFill", 0, 200, 0)
- AdvancedPetAssist: UpdateInstantOnlyHUD -> WindowSetTintColor("APAInstantOnlyHUDFill", 200, 100, 0)

## Related APIs

- [Color](../../xml/element_types/element_Color.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../../xml/handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [Start](../../events/game_events/game_event_Start.md) (MEDIUM 43/100) - Game Event

## Used With

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Notes

- Advanced return analysis: No strong return evidence observed

# LabelSetTextColor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 116 addons

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
| Addons seen in | Ace, ActionBarHide, ActionFraction, AdvancedPetAssist, Amethyst, AuctionStats, Aura, BuddyBind |
| Files seen in | AAOTWarBoard.lua, AAOTracker.lua, APAGuiHUD.lua, AuctionStats.lua, BuddyBind.lua, Busted.lua, CDownSettings.lua, CallingList.lua |
| Namespaces detected | LabelSetTextColor |
| Source kinds | lua_calls |
| Example locations | Ace: Color, ActionBarHide: Color, ActionFraction: SetPercentageBasedColors, ActionFraction: UpdateActionPoints, AdvancedPetAssist: UpdateFollowTargetHUD, AdvancedPetAssist: UpdateInstantOnlyHUD |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 662 |
| Global usage count | 662 |
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
LabelSetTextColor(arg1, arg2, arg3, arg4)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "APAFollowTargetHUDLabel", "APAInstantOnlyHUDLabel", "APAKitingHUDLabel" |
| arg2 | Observed as a numeric value. | Observed values: 0, 100, 107 |
| arg3 | Observed as a numeric value. | Observed values: 0, 10, 100 |
| arg4 | Observed as a numeric value. | Observed values: 0, 10, 100 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- AdvancedPetAssist
- Amethyst
- AuctionStats
- Aura
- BuddyBind
- BuffHead
- Busted
- CDown
- CM_ClosetGoblin
- CMap
- CaVES
- Calling
- CastSequence
- CleanUnitFrames
- Clock
- CombatTextNames
- Countdown
- Crusher
- DAoCBuff
- DPSMeter
- DammazKron
- DeepSleep
- DetauntHelper
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EA_UiModWindow
- EZCraftX
- EZGuard
- Effigy
- Enemy
- FozAuction
- GCDTracker
- GCDsaver
- Group Icons SG
- GroupRange
- Hopper
- InfoScroller
- Killer
- LibAddonButton
- LibWBToggler
- Map
- MapMonster
- MarkBuff
- Mech
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Moth
- Motion
- NaturalLog
- NoOverheal
- Obsidian
- PartyCast
- PeaceOut
- PotionBar
- Pure
- Pure Careerbar
- RO-Style Combat Text
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RVMOD_Targets
- Rangechecker
- RealmStatus
- ReliquaryHunter
- RoR_SoR
- RoR_debolster
- RvRContribution
- SOR
- SessionRPs
- Shinies
- SimpleCombatText
- Squared
- Statdoll
- Statdoll Light
- Statdoll Remix
- Swift Assist
- TacticSetNames
- TalismanGenie
- TargetInfoRing
- TargetRing
- Targets
- TastyButtons
- TexturedButtons
- Tokens
- Tortall_DPS
- Trakario
- TurretRange
- WARCommander
- WSCT
- WarBoard_AAOTracker
- WarBoard_Loc
- WarBoard_Session
- WarBoard_TogglerSquared
- WarBoard_TogglerWARCommander
- WarBoard_WarWhisperer
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- ZonePOP
- emotes
- minesweep
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Color -> LabelSetTextColor(self.name, red, green, blue)
- ActionBarHide: Color -> LabelSetTextColor(self.name, red, green, blue)
- ActionFraction: SetPercentageBasedColors -> LabelSetTextColor(label, 150, 255, 65)
- ActionFraction: SetPercentageBasedColors -> LabelSetTextColor(label, 255, 255, 70)
- ActionFraction: SetPercentageBasedColors -> LabelSetTextColor(label, 255, 170, 70)
- ActionFraction: SetPercentageBasedColors -> LabelSetTextColor(label, 255, 70, 70)

## Related APIs

- [Color](../../xml/element_types/element_Color.md) (HIGH 100/100) - XML Element Type
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../../xml/handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Used With

- [Color](../../xml/element_types/element_Color.md) (HIGH 100/100) - XML Element Type
- [DefaultColor.SetWindowTint](../../globals/functions/global_DefaultColor.SetWindowTint.md) (HIGH 100/100) - Global Function
- [LabelGetTextColor](window_LabelGetTextColor.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [DefaultColor.GetRowColor](../../globals/functions/global_DefaultColor.GetRowColor.md) (HIGH 96/100) - Global Function

## Notes

- Advanced return analysis: No strong return evidence observed

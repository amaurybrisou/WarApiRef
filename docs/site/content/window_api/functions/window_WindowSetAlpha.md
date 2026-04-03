# WindowSetAlpha

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 107 addons

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
| Addons seen in | Ace, ActionBarHide, AdjustTheTip, Amethyst, Atlas, AuctionStats, Aura, AutoMark |
| Files seen in | AdjustTheTip.lua, AuctionStats.lua, BBarsPetHP.lua, BBarsPlayerMechanic.lua, Bar.lua, BetterCC.lua, Bloody Mess.lua, Busted.lua |
| Namespaces detected | WindowSetAlpha |
| Source kinds | lua_calls |
| Example locations | Ace: Alpha, ActionBarHide: Alpha, AdjustTheTip: AbilityTooltipSetShowing, Amethyst: Alpha, Atlas: SetMapTransparency, Atlas: UpdateLegend |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 574 |
| Global usage count | 574 |
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
WindowSetAlpha(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "AtlasFrame", "AtlasFrameLegend", "AtlasFrameLegendBackground" |
| arg2 | Observed as a function or method reference. | Observed values: ((distance-30)/(maxDist)), (chatwindow_tabs_handle_input~=false and 1)or 0, (scrollbar_position==c_SCROLLBAR_POSITION_HIDDEN and 0)or scrollbar_alpha |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- AdjustTheTip
- Amethyst
- Atlas
- AuctionStats
- Aura
- AutoMark
- BBars - Mechanic Only
- BetterCC
- Bloody Mess
- BuffHead
- Busted
- CCTV
- CDown
- CMap
- Calling
- CastSequence
- ChattyCathy
- Cheeseboard
- CoolDownLine
- CraftingWillard
- Crusher
- DAoCBuff
- DetauntHelper
- DuffTimer
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- Enemy
- EveryBodyGuard
- FozAuction
- GCDsaver
- Group Icons
- Group Icons SG
- GroupRange
- GroupSpotter
- GuardBot
- GuardLine
- GuardList
- HealGrid
- Hopper
- InfoScroller
- KeyBar
- Killer
- Kwestor
- LibAddonButton
- LibWBToggler
- Map
- MapMonster
- MarkBuff
- Miracle Grow Remix
- Motion
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
- RandomMount
- RealmStatus
- ReliquaryHunter
- RetAlert
- RoR_SoR
- SNT_BUTTONS
- SNT_CASTBAR
- SNT_INFO
- SOR
- Shinies
- Squared
- SquaredHDIndicator
- SquaredHotIndicators
- Statdoll
- Statdoll Remix
- TargetInfoRing
- TargetRing
- Targets
- TexturedButtons
- TidyChat
- Tokens
- Tome Titan
- TomeTracker
- Tortall_DPS
- TurretRange
- Vectors
- WARRatingBuster
- WSCT
- WTes
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- compass
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Alpha -> WindowSetAlpha(self.name, alpha)
- ActionBarHide: Alpha -> WindowSetAlpha(self.name, alpha)
- AdjustTheTip: AbilityTooltipSetShowing -> WindowSetAlpha(parent, show and 1.0 or 0.0)
- Amethyst: Alpha -> WindowSetAlpha(self.name, alpha)
- Atlas: SetMapTransparency -> WindowSetAlpha("AtlasFrameMapContainerMapDisplay", 1.0)
- Atlas: SetMapTransparency -> WindowSetAlpha("AtlasFrame", alpha)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnSlide](../../xml/handlers/handler_OnSlide.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Used With

- [DynamicImageSetRotation](window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAlpha](window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowSetFontAlpha](window_WindowSetFontAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function

## Notes

- Advanced return analysis: No strong return evidence observed

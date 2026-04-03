# TextEditBoxSetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 90 addons

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
| Addons seen in | Ace, ActionBarHide, AdvancedPetAssist, AdvancedRenownTrainer, Amethyst, AuctionStats, AuctionTweaker, Aura |
| Files seen in | APAGui.lua, APAGuiHUD.lua, AdvancedRenownTrainingImportExport.lua, AuctionAssist.lua, AuctionTweaker.lua, AutoBand.lua, Busted.lua, Code/Core/Groups/EnemyEffectFilter.lua |
| Namespaces detected | TextEditBoxSetText |
| Source kinds | lua_calls |
| Example locations | Ace: Clear, Ace: SetText, ActionBarHide: Clear, ActionBarHide: SetText, AdvancedPetAssist: ApplyHUDColorOff, AdvancedPetAssist: ApplyHUDColorOn |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 722 |
| Global usage count | 722 |
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
TextEditBoxSetText(windowName, text)
```

## Description

Observed reading from or writing to edit-box controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: "APAEditHUDOffB", "APAEditHUDOffG", "APAEditHUDOffR" |
| text | Observed as a text or wstring payload. | Observed values: AddonToShow.RVProjectURL, AuctionSearchWindow.rankNotSetText, DAoCBuffSettings.TmpFilter[activefilter.index].name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- AdvancedPetAssist
- AdvancedRenownTrainer
- Amethyst
- AuctionStats
- AuctionTweaker
- Aura
- AutoBand
- BuffHead
- Busted
- CMap
- CastSequence
- Crusher
- DAoCBuff
- DetauntHelper
- DuffTimer
- EA_UiDebugTools
- EA_UiModWindow
- EZCraftX
- EZGuard
- Effigy
- Enemy
- FastFriends
- FixGit
- FozAuction
- GCDsaver
- GroupRange
- GuildWarden
- HealGrid
- Hopper
- InfoScroller
- JunkDump
- KeyBar
- Keyset
- Killer
- LibAddonButton
- LibWBToggler
- LoyalPet
- Map
- MapMonster
- MapPin
- Mass Refine
- MegaphonePlusPlus
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- ObjectInspector
- Obsidian
- PartyCast
- PieTracker
- Pure
- Pure Careerbar
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RealmStatus
- SNT_BUTTONS
- SNT_PANEL
- SOR
- Sequencer
- Shinies
- SocialWindow 2.0
- Squared
- TacticSetNames
- TargetRing
- TastyButtons
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyRoll
- Tokens
- TurretRange
- Twister
- Vectors
- WarBoard_WarWhisperer
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- bigger_MacroWindow
- nLootLink
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Clear -> TextEditBoxSetText(self.name, L "")
- Ace: SetText -> TextEditBoxSetText(self.name, towstring(text))
- ActionBarHide: Clear -> TextEditBoxSetText(self.name, L "")
- ActionBarHide: SetText -> TextEditBoxSetText(self.name, towstring(text))
- AdvancedPetAssist: ApplyHUDColorOff -> TextEditBoxSetText("APAEditHUDOffR", towstring(APA.hudColorOffR))
- AdvancedPetAssist: ApplyHUDColorOff -> TextEditBoxSetText("APAEditHUDOffG", towstring(APA.hudColorOffG))

## Related APIs

- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetMaxChars](window_TextEditBoxSetMaxChars.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Notes

- none

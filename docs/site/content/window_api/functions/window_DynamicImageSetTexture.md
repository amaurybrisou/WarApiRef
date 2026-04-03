# DynamicImageSetTexture

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 97 addons

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
| Addons seen in | Ace, AdvancedRenownTrainer, Amethyst, AuctionStats, Aura, AutoMark, BBars - Mechanic Only, BankArkel |
| Files seen in | AdvancedRenownTraining.lua, BBarsPetHP.lua, BBarsPlayerMechanic.lua, BankArkel.lua, Bar.lua, Bloody Mess.lua, BuddyBind.lua, Button.lua |
| Namespaces detected | DynamicImageSetTexture |
| Source kinds | lua_calls |
| Example locations | Ace: Texture, AdvancedRenownTrainer: CreateAbilityWindow, Amethyst: Texture, AuctionStats: OnRButtonUpItem, AuctionStats: UpdateItemOptions, Aura: SetDynamicImageTexture |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 474 |
| Global usage count | 474 |
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
DynamicImageSetTexture(windowName, texture, textureX, textureY)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target image control name. | Observed values: "BBarsPetHPBG", "BBarsPetHPBack", "BBarsPetHPFront" |
| texture | Observed as a texture resource name. | Observed values: "", "BlipDes", "BlipOrd" |
| textureX | Observed as a numeric texture coordinate or atlas offset. | Observed values: 0, 115, 17 |
| textureY | Observed as a numeric texture coordinate or atlas offset. | Observed values: 0, 105, 128 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedRenownTrainer
- Amethyst
- AuctionStats
- Aura
- AutoMark
- BBars - Mechanic Only
- BankArkel
- Bloody Mess
- Brizio's Crappy Computer Medic
- BuddyBind
- BuffHead
- CDown
- CM_ClosetGoblin
- CaVES
- Calling
- CastSequence
- CleanUnitFrames
- CleansingBuddy
- CombatTextNames
- CoolDownLine
- Crusher
- DAoCBuff
- DammazKron
- DetauntHelper
- DuffTimer
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- GCDTracker
- GCDsaver
- GetStats
- Group Icons
- GroupRange
- GroupSpotter
- GuildWarden
- Hopper
- InfoScroller
- LibAddonButton
- LibWBToggler
- MacroIcons
- Map
- MapMonster
- MarkBuff
- MiniMapMonster
- Miracle Grow Remix
- Moth
- Motion
- MouseHint
- NaturalLog
- NerfedButtons
- Obsidian
- PartyCast
- PlayEffectsOn
- PotionBar
- Pure
- Pure Careerbar
- RVMOD_Manager
- RandomMount
- RealmStatus
- ReliquaryHunter
- RoR_SoR
- SNT_BUTTONS
- SNT_CASTBAR
- SNT_PANEL
- SOR
- Sequencer
- Shinies
- Squared
- SquaredClick
- Swift Assist
- TargetInfoRing
- TargetRing
- Targets
- TastyButtons
- TexturedButtons
- TokenMachine
- Tokens
- Tome Titan
- Trakario
- Twister
- WSCT
- WarBoard_Menu
- WarBoard_TaliMon
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- talisman-monitor
- xHUD
- yAssistHelper
- zMailMod

## Examples

- Ace: Texture -> DynamicImageSetTexture(self.name, texture, x, y)
- AdvancedRenownTrainer: CreateAbilityWindow -> DynamicImageSetTexture(t.windowName.."Square", GetIconData(t.icon), 0, 0)
- Amethyst: Texture -> DynamicImageSetTexture(self.name, texture, x, y)
- AuctionStats: OnRButtonUpItem -> DynamicImageSetTexture(parentControls.."ItemImageIcon", "", 0, 0)
- AuctionStats: UpdateItemOptions -> DynamicImageSetTexture(parentControls.."ItemImageIcon", texture, x, y)
- Aura: SetDynamicImageTexture -> DynamicImageSetTexture(window, texture, dx, dy)

## Related APIs

- [DynamicImageSetTextureDimensions](window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [Texture](../../xml/element_types/element_Texture.md) (HIGH 100/100) - XML Element Type
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../../xml/handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Used With

- [DynamicImageSetRotation](window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [Texture](../../xml/element_types/element_Texture.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function

## Notes

- Advanced return analysis: No strong return evidence observed

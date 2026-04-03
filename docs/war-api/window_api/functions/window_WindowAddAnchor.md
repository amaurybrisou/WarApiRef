# WindowAddAnchor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 25 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, AnywhereTrainerAdditions, Aura, AutoMark, BagOMatic, BankArkel |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:140`, `/workspace/data/raw/AdvancedPetAssist/APAGui.lua:552`, `/workspace/data/raw/AnywhereTrainerAdditions/AnywhereTrainerAdditions.lua:18`, `/workspace/data/raw/Aura/Source/Aura.lua:505`, `/workspace/data/raw/Aura/Source/Aura.lua:534`, `/workspace/data/raw/Aura/Source/AuraTexture.lua:41`, `/workspace/data/raw/AutoMark/Source/AutoMark.lua:124`, `/workspace/data/raw/BankArkel/BankArkel.lua:172` |
| Namespaces detected | WindowAddAnchor |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:AddAnchor, AdvancedPetAssist: AdvancedPetAssist.local.AnchorInContent, AdvancedPetAssist: AnchorInContent, AdvancedRenownTrainer: AdvancedRenownTrainer.local.CreateAbilityWindow, AdvancedRenownTrainer: CreateAbilityWindow, AnywhereTrainerAdditions: AnywhereTrainerAdditions.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 297 |
| Global usage count | 297 |
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
WindowAddAnchor(windowName, point, relativeTo, relativePoint, offsetX, offsetY)
```

## Description

Observed positioning windows relative to other runtime UI elements.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as the window being positioned. | Observed values: "BagOMaticButton", "BankArkelBackpack", "EnemyCombatLogEpsWindowInDpsLabel" |
| point | Observed as an anchor point string. | Observed values: "LEFT", "RIGHT", "bottom" |
| relativeTo | Observed as a reference window name. | Observed values: "APAOptionsContent", "EA_Window_Backpack", "EA_Window_BackpackTitleBarText" |
| relativePoint | Observed as a reference anchor point string. | Observed values: "LEFT", "RIGHT", "bottom" |
| offsetX | Observed as a numeric horizontal offset. | Observed values: (layout.Distance.Offset.X or 0), (settings.Distance.Offset.X or 0), -(width/2) |
| offsetY | Observed as a numeric vertical offset. | Observed values: (g.settings.combatLogIDSRowSize[2]+g.settings.combatLogIDSRowPadding)*(index-1), (layout.Distance.Offset.Y or 0), (settings.Distance.Offset.Y or 0) |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates runtime window layout state.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- AnywhereTrainerAdditions
- Aura
- AutoMark
- BagOMatic
- BankArkel
- BuffHead
- CombatTextNames
- DAoCBuff
- Enemy
- Killer
- LibWBToggler
- PartyCast
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WhoHealedMe
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:AddAnchor -> WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
- AdvancedPetAssist: AdvancedPetAssist.local.AnchorInContent -> WindowAddAnchor(name, "topleft", "APAOptionsContent", "topleft", x, y)
- AdvancedPetAssist: AnchorInContent -> WindowAddAnchor(name, "topleft", "APAOptionsContent", "topleft", x, y)
- AdvancedRenownTrainer: AdvancedRenownTrainer.local.CreateAbilityWindow -> WindowAddAnchor(t.windowName, t.relativePoint, t.relativeTo, t.point, t.x, t.y)
- AdvancedRenownTrainer: CreateAbilityWindow -> WindowAddAnchor(t.windowName, t.relativePoint, t.relativeTo, t.point, t.x, t.y)
- AnywhereTrainerAdditions: AnywhereTrainerAdditions.Initialize -> WindowAddAnchor(tab.Name, tab.Anchor.Point, tab.Anchor.RelativeTo, tab.Anchor.RelativePoint, tab.Anchor.X, tab.Anchor.Y)

## Related APIs

- none

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [StatusBarGetCurrentValue](window_StatusBarGetCurrentValue.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetId](window_WindowSetId.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [StatusBarGetMaximumValue](window_StatusBarGetMaximumValue.md) (HIGH 80/100) - Window Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [Anchor](../../xml/element_types/element_Anchor.md) (MEDIUM 55/100) - XML Element Type

## Notes

- none

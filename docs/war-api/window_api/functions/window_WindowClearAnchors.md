# WindowClearAnchors

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 22 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AnywhereTrainerAdditions, Aura, AutoMark, BankArkel, BuffHead, CombatTextNames |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:152`, `/workspace/data/raw/AdvancedPetAssist/APAGui.lua:552`, `/workspace/data/raw/AnywhereTrainerAdditions/AnywhereTrainerAdditions.lua:18`, `/workspace/data/raw/Aura/Source/Aura.lua:505`, `/workspace/data/raw/Aura/Source/Aura.lua:534`, `/workspace/data/raw/Aura/Source/AuraTexture.lua:41`, `/workspace/data/raw/AutoMark/Source/AutoMark.lua:124`, `/workspace/data/raw/BankArkel/BankArkel.lua:172` |
| Namespaces detected | WindowClearAnchors |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:ClearAnchors, AdvancedPetAssist: AdvancedPetAssist.local.AnchorInContent, AdvancedPetAssist: AnchorInContent, AnywhereTrainerAdditions: AnywhereTrainerAdditions.Initialize, Aura: Aura:UpdateTimerWindow, Aura: Aura:UpdateWindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 151 |
| Global usage count | 151 |
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
WindowClearAnchors(windowName)
```

## Description

Observed resetting a window layout before applying new runtime anchors.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "AnywhereTrainerBottomBookend", "AnywhereTrainerTopBookend", "BankArkelBackpack" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates runtime window layout state.

## Seen In

- Ace
- AdvancedPetAssist
- AnywhereTrainerAdditions
- Aura
- AutoMark
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
- TurretRange
- WSCT
- WhoHealedMe
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:ClearAnchors -> WindowClearAnchors(self.name)
- AdvancedPetAssist: AdvancedPetAssist.local.AnchorInContent -> WindowClearAnchors(name)
- AdvancedPetAssist: AnchorInContent -> WindowClearAnchors(name)
- AnywhereTrainerAdditions: AnywhereTrainerAdditions.Initialize -> WindowClearAnchors("AnywhereTrainerTopBookend")
- AnywhereTrainerAdditions: AnywhereTrainerAdditions.Initialize -> WindowClearAnchors("AnywhereTrainerBottomBookend")
- Aura: Aura:UpdateTimerWindow -> WindowClearAnchors(windowId)

## Related APIs

- none

## Used With

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [StatusBarGetCurrentValue](window_StatusBarGetCurrentValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetLayer](window_WindowGetLayer.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [StatusBarGetMaximumValue](window_StatusBarGetMaximumValue.md) (HIGH 80/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [Anchor](../../xml/element_types/element_Anchor.md) (MEDIUM 55/100) - XML Element Type

## Notes

- none

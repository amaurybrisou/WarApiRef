# WindowSetDimensions

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 21 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AnywhereTrainer, Aura, BankArkel, BuffHead, DAoCBuff, Enemy |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:1174`, `/workspace/data/raw/Ace/LibGUI.lua:1261`, `/workspace/data/raw/Ace/LibGUI.lua:367`, `/workspace/data/raw/Ace/LibGUI.lua:419`, `/workspace/data/raw/Ace/LibGUI.lua:527`, `/workspace/data/raw/Ace/LibGUI.lua:650`, `/workspace/data/raw/Ace/LibGUI.lua:704`, `/workspace/data/raw/Ace/LibGUI.lua:983` |
| Namespaces detected | WindowSetDimensions |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:Resize, Ace: LIBGUI_Image:Resize, Ace: LIBGUI_Label:Resize, Ace: LIBGUI_MultiTextbox:Resize, Ace: LIBGUI_Scrollbar:Resize, Ace: LIBGUI_Statusbar:Resize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 140 |
| Global usage count | 140 |
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
WindowSetDimensions(windowName, arg2, arg3)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "APAOptions", "APAOptionsContent", "APAOptionsTitleText" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 1, 100, 1000 |
| arg3 | Observed as a runtime window or control identifier. | Observed values: 1, 100, 110 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedPetAssist
- AnywhereTrainer
- Aura
- BankArkel
- BuffHead
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

- Ace: LIBGUI_Button:Resize -> WindowSetDimensions(self.name, width, self.height)
- Ace: LIBGUI_Image:Resize -> WindowSetDimensions(self.name, width, height)
- Ace: LIBGUI_Label:Resize -> WindowSetDimensions(self.name, self.width, self.height)
- Ace: LIBGUI_MultiTextbox:Resize -> WindowSetDimensions(self.name, self.width, self.height)
- Ace: LIBGUI_Scrollbar:Resize -> WindowSetDimensions(self.name, self.width, self.height)
- Ace: LIBGUI_Statusbar:Resize -> WindowSetDimensions(self.name, width, self.height)

## Related APIs

- none

## Used With

- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none

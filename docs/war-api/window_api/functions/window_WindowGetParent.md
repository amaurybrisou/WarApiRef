# WindowGetParent

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 33 addons

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
| Addons seen in | AggroMeter, Aura, DammazKron, EA_OpenPartyWindow, EA_ScenarioGroupWindow, EA_UiModWindow, Emojii, Enemy |
| Files seen in | AggroMeter.lua, Classes/Display.lua, Code/Marks/MarkTemplate.lua, Core/Tome/DK_Tome.lua, Core/ToolTip/DK_Tooltip.lua, Emojii.lua, Gui/HealGridGuiColorSelect.lua, Gui/HealGridGuiTabSpellTrack.lua |
| Namespaces detected | WindowGetParent |
| Source kinds | lua_calls |
| Example locations | AggroMeter: OnMouseOverStart, AggroMeter: SelectChar, Aura: OnIconLButtonUp, DammazKron: DK_OnSelectProfil, DammazKron: DK_OnSelectProfilMini, DammazKron: DestroyWindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 103 |
| Global usage count | 103 |
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
WindowGetParent(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: ButtonName, SystemData.ActiveWindow.name, SystemData.MouseOverWindow.name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AggroMeter
- Aura
- DammazKron
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiModWindow
- Emojii
- Enemy
- HealGrid
- HideHiddenFrames
- KillTracker
- LibAddonButton
- LibSurveyor
- LoyalPet
- Map
- MapMonster
- Miracle Grow Remix
- Motion
- NAMBLA
- Pure
- RVMOD_Manager
- RVMOD_Targets
- RandomMount
- Refer
- RoR_SoR
- SocialWindow 2.0
- TastyButtons
- TokenMachine
- TomeTracker
- Vectors
- WSCT
- WTes
- zMailMod

## Examples

- AggroMeter: OnMouseOverStart -> WindowGetParent(SystemData.MouseOverWindow.name)
- AggroMeter: SelectChar -> WindowGetParent(SystemData.MouseOverWindow.name)
- Aura: OnIconLButtonUp -> WindowGetParent(WindowGetParent(WindowGetParent(SystemData.ActiveWindow.name)))
- Aura: OnIconLButtonUp -> WindowGetParent(WindowGetParent(SystemData.ActiveWindow.name))
- Aura: OnIconLButtonUp -> WindowGetParent(SystemData.ActiveWindow.name)
- DammazKron: DK_OnSelectProfil -> WindowGetParent(SystemData.ActiveWindow.name)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnSelChanged](../../xml/handlers/handler_OnSelChanged.md) (HIGH 88/100) - XML Event
- [OnTextChanged](../../xml/handlers/handler_OnTextChanged.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Used With

- [ButtonGetDisabledFlag](window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [WindowGetId](window_WindowGetId.md) (HIGH 100/100) - Window Function

## Affects

- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Advanced return analysis: No strong return evidence observed

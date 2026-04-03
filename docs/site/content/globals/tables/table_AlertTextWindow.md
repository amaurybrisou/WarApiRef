# AlertTextWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 125

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Aura, BlackBook, DammazKron, DetauntHelper, EmoteAlert, LootAlert, MapMonster, MapPin |
| Files seen in | Core/DK_Core.lua, Source/Aura.lua, Source/Common.lua, Source/MapMonster.lua, Source/MapMonster_EditorWindow.lua, Source/MapMonster_PinTypeEditorWindow.lua, Source/SquaredClick.lua, Source/TomeTracker_Pins.lua |
| Namespaces detected | AlertTextWindow |
| Source kinds | lua_calls |
| Example locations | Aura: Activate, Aura: Deactivate, BlackBook: CheckOdds, BlackBook: PlayerRenownUpdated, DammazKron: AlertRegister, DetauntHelper: Print |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 35 |
| Global usage count | 3 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Shared function table with 3 member functions; the primary API surface for 21 addons.

## Functions

- AlertTextWindow.AddAlert
- AlertTextWindow.AddLine
- AlertTextWindow.SetAlertData

## Observed Members

- none

## Seen In

- Aura
- BlackBook
- DammazKron
- DetauntHelper
- EmoteAlert
- LootAlert
- MapMonster
- MapPin
- MegaphonePlus
- MegaphonePlusPlus
- Minmap
- PlanB
- SNT_BUTTONS
- SNT_INFO
- Sequencer
- SessionRPs
- SquaredClick
- Targets
- TheSeeker
- TomeTracker
- WSCT

## Examples

- Aura: Activate -> AlertTextWindow.AddLine(AuraConstants.AlertText[self:Get("activation-alerttexttype")].type, self:Get("activation-alerttext"))
- Aura: Deactivate -> AlertTextWindow.AddLine(AuraConstants.AlertText[self:Get("deactivation-alerttexttype")].type, self:Get("deactivation-alerttext"))
- BlackBook: CheckOdds -> AlertTextWindow.AddAlert()
- BlackBook: PlayerRenownUpdated -> AlertTextWindow.AddAlert()
- DammazKron: AlertRegister -> AlertTextWindow.AddAlert()
- DetauntHelper: Print -> AlertTextWindow.AddLine(SystemData.AlertText.Types.DEFAULT, msg)

## Notes

- none

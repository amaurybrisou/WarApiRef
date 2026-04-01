# Usage By Addon

## XML To Lua Examples

| Addon | Frame | Event | Lua Function |
| --- | --- | --- | --- |
| AdvancedPetAssist | APAComboAttackBind | OnSelChanged | APAGui.OnComboChanged |
| AdvancedPetAssist | APAComboAutoReattack | OnSelChanged | APAGui.OnComboChanged |
| AdvancedPetAssist | APAComboAutoReattackDelay | OnSelChanged | APAGui.OnComboChanged |
| AdvancedPetAssist | APAComboCastDelay | OnSelChanged | APAGui.OnComboChanged |
| AdvancedPetAssist | APAComboCastOnAcquire | OnSelChanged | APAGui.OnComboChanged |
| AdvancedPetAssist | APAComboCombatExitDelay | OnSelChanged | APAGui.OnComboChanged |
| AdvancedPetAssist | APAComboDebug | OnSelChanged | APAGui.OnComboChanged |
| AdvancedPetAssist | APAComboEnabled | OnSelChanged | APAGui.OnComboChanged |
| AdvancedPetAssist | APAComboFTAutoArmMount | OnSelChanged | APAGui.OnComboChanged |
| AdvancedPetAssist | APAComboFTFollowOnMount | OnSelChanged | APAGui.OnComboChanged |
| AdvancedPetAssist | APAComboFTPendingDelay | OnSelChanged | APAGui.OnComboChanged |
| AdvancedPetAssist | APAComboFTRedirectCooldown | OnSelChanged | APAGui.OnComboChanged |

## Representative Function Usage

- AlertTextWindow.AddAlert: MapPin -> AlertTextWindow.AddAlert()
- AlertTextWindow.AddAlert: PlanB -> AlertTextWindow.AddAlert()
- AlertTextWindow.AddAlert: PlanB -> AlertTextWindow.AddAlert()
- AlertTextWindow.AddLine: Aura -> AlertTextWindow.AddLine(AuraConstants.AlertText[self:Get("activation-alerttexttype")].type, self:Get("activation-alerttext"))
- AlertTextWindow.AddLine: Aura -> AlertTextWindow.AddLine(AuraConstants.AlertText[self:Get("deactivation-alerttexttype")].type, self:Get("deactivation-alerttext"))
- AlertTextWindow.AddLine: MapMonster -> AlertTextWindow.AddLine(MapMonster.AlertTypes.SUCCESS, T["SubTypeCreatedSuccess"])
- AlertTextWindow.AddLine: MapMonster -> AlertTextWindow.AddLine(MapMonster.AlertTypes.ERROR, T["SubTypeCreatedFailed"])
- AlertTextWindow.AddLine: MapMonster -> AlertTextWindow.AddLine(MapMonster.AlertTypes.ERROR, errorMessage)
- AlertTextWindow.AddLine: MapMonster -> AlertTextWindow.AddLine(MapMonster.AlertTypes.SUCCESS, T["SuccessAlert"])
- BankWindow.GetItem: AnywhereTrainerAdditions -> BankWindow.GetItem(slot)
- BankWindow.GetItem: BankWindowFix -> BankWindow.GetItem(slot)
- BankWindow.GetSlotNumberForButtonIndex: AnywhereTrainerAdditions -> BankWindow.GetSlotNumberForButtonIndex(buttonIndex)
- BankWindow.GetSlotNumberForButtonIndex: BankWindowFix -> BankWindow.GetSlotNumberForButtonIndex(buttonIndex)
- BankWindow.Hide: AnywhereTrainer -> BankWindow.Hide()
- BankWindow.Hide: AnywhereTrainerAdditions -> BankWindow.Hide()
- BankWindow.Hide: BankArkel -> BankWindow.Hide()
- BankWindow.IsShowing: BagOMatic -> BankWindow.IsShowing()
- BankWindow.IsShowing: BankWindowFix -> BankWindow.IsShowing()
- BankWindow.Show: AnywhereTrainer -> BankWindow.Show()
- BankWindow.Show: BagOMatic -> BankWindow.Show()

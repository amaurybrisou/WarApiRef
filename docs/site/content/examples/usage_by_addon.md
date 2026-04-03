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

- BankWindow.Hide: AnywhereTrainer -> BankWindow.Hide()
- BankWindow.Hide: AnywhereTrainerAdditions -> BankWindow.Hide()
- BankWindow.Hide: BankArkel -> BankWindow.Hide()
- BankWindow.Show: AnywhereTrainer -> BankWindow.Show()
- BankWindow.Show: BagOMatic -> BankWindow.Show()
- BankWindow.Show: BankArkel -> BankWindow.Show()
- CreateWindow: AdvancedPetAssist -> CreateWindow("APAOptions", true)
- CreateWindow: AdvancedRenownTrainer -> CreateWindow("AdvancedRenownTrainingPresetsWindow", false)
- CreateWindow: AdvancedRenownTrainer -> CreateWindow(ImportWindowName, false)
- CreateWindow: AdvancedRenownTrainer -> CreateWindow(ImportNameInputWindowName, false)
- CreateWindow: AdvancedRenownTrainer -> CreateWindow(ExportWindowName, false)
- CreateWindow: AdvancedRenownTrainer -> CreateWindow(LinkWindowName, false)
- CreateWindowFromTemplate: Ace -> CreateWindowFromTemplate(w.name, base, w.parent)
- CreateWindowFromTemplate: Ace -> CreateWindowFromTemplate(w.name, base, w.parent)
- CreateWindowFromTemplate: Ace -> CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
- CreateWindowFromTemplate: Ace -> CreateWindowFromTemplate(w.name, base, w.parent)
- CreateWindowFromTemplate: Ace -> CreateWindowFromTemplate(w.name, base, w.parent)
- CreateWindowFromTemplate: Ace -> CreateWindowFromTemplate(w.name, base, w.parent)
- Cursor.Clear: AnywhereTrainerAdditions -> Cursor.Clear()
- Cursor.Clear: CM_ClosetGoblin -> Cursor.Clear()

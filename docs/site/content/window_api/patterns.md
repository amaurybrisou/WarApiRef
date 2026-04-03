# Window Patterns

## Window creation

- Confidence: HIGH

- Description: Observed top-level UI windows being created from XML definitions at initialization time.

- Evidence:

- AbilityAlert: CreateWindow("AAWindow", true)
  - AbilityNotifier: CreateWindow("AbHelpWindow", true)
  - ActionFraction: CreateWindow(windowName, true)
  - ActionPoints: CreateWindow("ActionPointsWindow", true)
  - AdvancedPetAssist: CreateWindow("APAOptions", true)
  - AdvancedRenownTrainer: CreateWindow("AdvancedRenownTrainingPresetsWindow", false)

## Template instantiation

- Confidence: HIGH

- Description: Observed repeated UI elements being instantiated from XML templates at runtime.

- Evidence:

- Ace: CreateWindowFromTemplate(w.name, base, w.parent)
  - Ace: CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
  - ActionBarHide: CreateWindowFromTemplate(w.name, "EA_ComboBox_DefaultResizable", w.parent)
  - ActionBarHide: CreateWindowFromTemplate(w.name, "EA_Button_DefaultCheckBox", w.parent)
  - ActionBarHide: CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
  - ActionBarHide: CreateWindowFromTemplate(w.name, "EA_EditBox_DefaultFrame", w.parent)

## Runtime anchoring

- Confidence: HIGH

- Description: Observed layouts being finalized in Lua by clearing and re-adding anchors after window creation.

- Evidence:

- Ace: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
  - ActionBarHide: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
  - ActionFraction: WindowAddAnchor(windowName, "center", "Root", "center", 0, 0)
  - ActionFraction: WindowAddAnchor(windowName, "center", "PlayerWindowStatusContainerAPPercentBar", "center", 2, 6)
  - ActionFraction: WindowAddAnchor(windowName, "center", "Root", "center", 0, 0)
  - AdjustTheTip: WindowAddAnchor(c_HEALTH_BAR_CONTAINER, "bottomright", "MouseOverTargetUnitWindow", "bottomright", -10, -10)

## XML to Lua binding

- Confidence: HIGH

- Description: Observed XML handler events routing directly into addon Lua functions through shared engine bindings.

- Evidence:

- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp

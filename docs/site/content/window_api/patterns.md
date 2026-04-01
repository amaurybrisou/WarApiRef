# Window Patterns

## Window creation

- Confidence: HIGH

- Description: Observed top-level UI windows being created from XML definitions at initialization time.

- Evidence:

- AdvancedPetAssist: CreateWindow("APAOptions", true)
  - AdvancedRenownTrainer: CreateWindow("AdvancedRenownTrainingPresetsWindow", false)
  - AdvancedRenownTrainer: CreateWindow(ImportWindowName, false)
  - AdvancedRenownTrainer: CreateWindow(ImportNameInputWindowName, false)
  - AdvancedRenownTrainer: CreateWindow(ExportWindowName, false)
  - AdvancedRenownTrainer: CreateWindow(LinkWindowName, false)

## Template instantiation

- Confidence: HIGH

- Description: Observed repeated UI elements being instantiated from XML templates at runtime.

- Evidence:

- Ace: CreateWindowFromTemplate(w.name, base, w.parent)
  - Ace: CreateWindowFromTemplate(w.name, base, w.parent)
  - Ace: CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
  - Ace: CreateWindowFromTemplate(w.name, base, w.parent)
  - Ace: CreateWindowFromTemplate(w.name, base, w.parent)
  - Ace: CreateWindowFromTemplate(w.name, base, w.parent)

## Runtime anchoring

- Confidence: HIGH

- Description: Observed layouts being finalized in Lua by clearing and re-adding anchors after window creation.

- Evidence:

- Ace: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
  - AdvancedPetAssist: WindowAddAnchor(name, "topleft", "APAOptionsContent", "topleft", x, y)
  - AdvancedPetAssist: WindowAddAnchor(name, "topleft", "APAOptionsContent", "topleft", x, y)
  - AdvancedRenownTrainer: WindowAddAnchor(t.windowName, t.relativePoint, t.relativeTo, t.point, t.x, t.y)
  - AdvancedRenownTrainer: WindowAddAnchor(t.windowName, t.relativePoint, t.relativeTo, t.point, t.x, t.y)
  - AnywhereTrainerAdditions: WindowAddAnchor(tab.Name, tab.Anchor.Point, tab.Anchor.RelativeTo, tab.Anchor.RelativePoint, tab.Anchor.X, tab.Anchor.Y)

## XML to Lua binding

- Confidence: HIGH

- Description: Observed XML handler events routing directly into addon Lua functions through shared engine bindings.

- Evidence:

- AdvancedPetAssist: APAComboAttackBind.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboAutoReattack.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboAutoReattackDelay.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboCastDelay.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboCastOnAcquire.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboCombatExitDelay.OnSelChanged -> APAGui.OnComboChanged

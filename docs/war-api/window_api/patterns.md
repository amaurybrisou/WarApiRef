# Window Patterns

## Window creation

- Confidence: HIGH

- Description: Observed top-level UI windows being created from XML definitions at initialization time.

- Evidence:

- InfoScroller: CreateWindow("InfoScrollerMainWindow", true)
  - Moth: CreateWindow("Moth", true)
  - Soloq: CreateWindow(overviewWindowName, false)
  - TidyChat: CreateWindow(c_TEXT_ENTRY_ANCHOR, false)
  - TidyRoll: CreateWindow(c_TROLL_AUTO_ROLL_WINDOW, false)
  - TidyRoll: CreateWindow(c_TIDY_ROLL_ANCHOR, false)

## Template instantiation

- Confidence: HIGH

- Description: Observed repeated UI elements being instantiated from XML templates at runtime.

- Evidence:

- InfoScroller: CreateWindowFromTemplate(WindowName, "InfoScrollerTemplate", "InfoScrollerMainWindow")
  - InfoScroller: CreateWindowFromTemplate(w.name, base, w.parent)
  - InfoScroller: CreateWindowFromTemplate(w.name, base, w.parent)
  - InfoScroller: CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
  - InfoScroller: CreateWindowFromTemplate(w.name, base, w.parent)
  - InfoScroller: CreateWindowFromTemplate(w.name, base, w.parent)

## Runtime anchoring

- Confidence: HIGH

- Description: Observed layouts being finalized in Lua by clearing and re-adding anchors after window creation.

- Evidence:

- InfoScroller: WindowAddAnchor(WindowName.."Image", anchor, WindowName, anchor, T_X, T_Y)
  - InfoScroller: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
  - Moth: WindowAddAnchor("Moth", "bottomleft", "CursorWindow", "topleft", 0, 0)
  - Moth: WindowAddAnchor("MothHealthBar", "bottomleft", "MothBackground", "topleft", 0, yOffset)
  - PartyCast: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
  - PartyCast: WindowAddAnchor("PartyCastWindow"..i, Frame_Anchor, "PartyCastWindow_Dynamic"..i, Frame_Anchor, 0, PartyCast.Settings.Offset)

## XML to Lua binding

- Confidence: HIGH

- Description: Observed XML handler events routing directly into addon Lua functions through shared engine bindings.

- Evidence:

- InfoScroller: InfoScrollerTemplateLabel1.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
  - InfoScroller: InfoScrollerTemplateLabel1.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
  - InfoScroller: InfoScrollerTemplateLabel2.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
  - InfoScroller: InfoScrollerTemplateLabel2.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
  - InfoScroller: InfoScrollerTemplateLabel3.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
  - InfoScroller: InfoScrollerTemplateLabel3.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp

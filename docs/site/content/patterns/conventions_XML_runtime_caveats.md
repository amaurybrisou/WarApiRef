# XML runtime caveats

- Category: conventions
- Confidence: MEDIUM

## What is this pattern

This pattern documents runtime behavior peculiarities of XML input and scroll layout that differ from static declarations. Specifically:
- Parent input state affects child input handling
- Scroll layout dimensions depend on ancestor states and outer window sizing

## Why it exists

The engine's layout and input systems have implicit dependencies on ancestor state that are not declared in XML. Understanding these caveats prevents mysterious input and rendering bugs.

## When it appears

- **Input debugging**: Child button clicks don't register despite correct XML
- **Scroll content**: ScrollWindow content size reported incorrectly
- **Layout issues**: Window content appears cut off or mispositioned

## Minimal example

```lua
-- Input caveat: parent must handle input for child to receive clicks
Parent:SetProperty("handleinput", true)
Child:SetProperty("handleinput", true)

-- Scroll caveat: must update scroll rect manually after content changes
ScrollWindowUpdateScrollRect("MyScrollWindow")
```

## Annotated real example

From WhoHealedMe (observed):

```lua
-- Problem: child OnLButtonUp handler never fires despite being clickable in XML
-- Cause: parent window has handleinput=false

-- Solution: ensure parent enables input
if parentWindow then
  parentWindow:SetProperty("handleinput", true)
  childButton:SetProperty("handleinput", true)
end
```

Another case:

```lua
-- Problem: scroll content dimensions initially under-reported
-- Cause: ScrollWindow layout not finalized at creation time

-- Solution: recompute dimensions after content is added
local outerSize = WindowGetProperty(parentWindow, "width")
ScrollWindowUpdateScrollRect("MyScrollWindow")
```

**Key findings**:
- Input handling traverses ancestor chain; if any ancestor has `handleinput=false`, children don't receive events
- Scroll dimensions should be computed from outer parent, not child content
- Visual layout happens at render time; Lua can't detect size changes in the same frame

## Detection signals / evidence

**Observe**:
- Child element has `OnLButtonUp` handler in XML but handler never fires
- Checking `handleinput` property on parent shows `false`
- ScrollWindow content rendered at incorrect size or clipped unexpectedly
- Calling `ScrollWindowUpdateScrollRect` fixes scroll sizing

**Confidence indicators**:
- Setting parent `handleinput=true` immediately fixes child input
- Setting proper parent container size fixes scroll render

## Related patterns

- [Window event hooks](../events/events_Window_event_hooks.md) — how event handlers work
- [Runtime anchoring](./window_Runtime_anchoring.md) — dynamic positioning
- [UI creation pattern](./conventions_UI_creation_pattern.md) — where UI is created

## Common pitfalls

1. **Assuming child input is independent**: Setting child `handleinput=true` without checking parent.
   ```lua
   -- ❌ Wrong: child input enabled but parent blocks propagation
   child:SetProperty("handleinput", true)
   
   -- ✓ Correct: validate parent chain
   if parent:GetProperty("handleinput") then
     child:SetProperty("handleinput", true)
   end
   ```

2. **Computing scroll size from child**: Getting size of scroll content instead of outer container.
   ```lua
   -- ❌ Wrong: uses content size, not constraint size
   local h = child:GetProperty("height")
   
   -- ✓ Correct: use parent container size
   local h = parentWindow:GetProperty("height")
   ```

3. **Relying on immediate layout**: Expecting scroll size updates in same frame as content changes.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [ScrollWindow](../xml/element_types/element_ScrollWindow.md) (HIGH 100/100) - XML Element Type
- [ScrollWindowUpdateScrollRect](../window_api/functions/window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function

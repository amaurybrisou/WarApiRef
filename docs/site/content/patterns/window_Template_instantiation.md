# Template instantiation

- Category: window
- Confidence: HIGH

## What is this pattern

The template instantiation pattern describes creating multiple similar UI elements from a single XML template definition. [CreateWindowFromTemplate](../globals/functions/global_CreateWindowFromTemplate.md) clones a template window multiple times with different names, avoiding duplicate XML declarations while supporting dynamic UI generation.

## Why it exists

Templates enable:
- **Code reuse**: Define once, instantiate many times
- **Parameter variation**: Each instance has same structure but different content
- **Dynamic generation**: Create UI elements in response to data (player lists, presets, tactic rows)
- **String-based naming**: Templates support dynamic naming (name + index)

## When it appears

- **List rows**: Multiple identical ListBox rows, each instance a template clone
- **Tabbed interfaces**: Multiple tabs created from same template
- **Dynamic groups**: Variable number of UI sections (presets, profiles, groups)

## Minimal example

```xml
<!-- Define template (not visible directly) -->
<Window name="MyButtonTemplate" template="true">
  <Button name="button" />
</Window>
```

```lua
-- Instantiate template multiple times
for i = 1, 10 do
  CreateWindowFromTemplate(
    "MyButton" .. i,
    "MyButtonTemplate",
    parentWindow
  )
end
```

## Annotated real example

From Ace (observed):

```lua
-- Iterate over config data and instantiate button template for each
for i, configItem in ipairs(configItems) do
  local buttonName = "ConfigButton_" .. i
  
  CreateWindowFromTemplate(
    buttonName,
    "EA_Button_Default",  -- Template name
    containerWindow        -- Parent
  )
  
  -- Position each instance
  WindowAddAnchor(
    buttonName,
    "topleft",
    containerWindow,
    "topleft",
    0,
    (i - 1) * 25
  )
end
```

Another example from Ace:

```lua
-- Close button template instantiation
CreateWindowFromTemplate(
  windowName,
  "EA_Button_DefaultWindowClose",
  windowName
)
```

**Key observations**:
- Template name references XML `<Window name= template="true">`
- Instance name is unique (usually suffixed with number or ID)
- Parent window must already exist
- Each instance shares structure but has independent name and state

## Detection signals / evidence

**Observe**:
- `CreateWindowFromTemplate(instanceName, templateName, parent)` calls in loops
- Instance names vary (numbered or ID-suffixed) while template name is constant
- XML has `<Window template="true">` declaration
- Multiple windows created from single template definition

**Confidence indicators**:
- Template name matches XML template declaration
- Instance names are unique across all calls
- All instances have consistent structure

## Related patterns

- [Window creation](./window_Window_creation.md) — creating windows from XML
- [UI creation pattern](./conventions_UI_creation_pattern.md) — broader pattern
- [Runtime anchoring](./window_Runtime_anchoring.md) — positioning instances

## Common pitfalls

1. **Template not found**: Template name doesn't match XML `<Window name= template="true">` declaration.
   ```lua
   -- ❌ Wrong: template name typo
   CreateWindowFromTemplate("btn1", "EA_Botton_Default", ...)  -- Typo in "Button"
   
   -- ✓ Correct: exact template name
   CreateWindowFromTemplate("btn1", "EA_Button_Default", ...)
   ```

2. **Duplicate instance names**: Creating multiple instances with same name (only last one persists).
   ```lua
   -- ❌ Wrong: same name for both instances
   CreateWindowFromTemplate("btn", "MyTemplate", ...)
   CreateWindowFromTemplate("btn", "MyTemplate", ...)  -- Overwrites first
   
   -- ✓ Correct: unique instance names
   CreateWindowFromTemplate("btn1", "MyTemplate", ...)
   CreateWindowFromTemplate("btn2", "MyTemplate", ...)
   ```

3. **Parent window missing**: Creating template instances before parent exists.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [CreateWindowFromTemplate](../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

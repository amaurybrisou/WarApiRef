# Template instantiation

- Category: window
- Confidence: MEDIUM

## Description

Observed repeated UI elements being instantiated from XML templates at runtime.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [EA_Button_DefaultWindowClose](../globals/constants/constant_EA_Button_DefaultWindowClose.md) (HIGH 100/100) - Constant
- [EA_Label_DefaultText](../globals/constants/constant_EA_Label_DefaultText.md) (HIGH 100/100) - Constant
- [Label](../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ListBox, Window
```

## Example Code

```lua
TidyChat: CreateWindowFromTemplate(c_CHANNEL_MENU.."AllianceButton", "ChannelMenuButton", "ChatChannelSelectionWindow")
```

## Evidence

- TidyChat: CreateWindowFromTemplate(c_CHANNEL_MENU.."AllianceButton", "ChannelMenuButton", "ChatChannelSelectionWindow")
- TidyChat: CreateWindowFromTemplate(c_CHANNEL_MENU.."AdviceButton", "ChannelMenuButton", "ChatChannelSelectionWindow")
- TidyChat: CreateWindowFromTemplate(c_CHANNEL_MENU.."ScenarioButton", "ChannelMenuButton", "ChatChannelSelectionWindow")
- TidyRoll: CreateWindowFromTemplate(c_TROLL_TITLE, "EA_Label_DefaultText", c_TIDY_ROLL_OPTIONS)
- TidyRoll: CreateWindowFromTemplate(c_TROLL_VERSION, "EA_Label_DefaultText", c_TIDY_ROLL_OPTIONS)
- TidyRoll: CreateWindowFromTemplate(c_TROLL_CLOSE, "EA_Button_DefaultWindowClose", c_TIDY_ROLL_OPTIONS)

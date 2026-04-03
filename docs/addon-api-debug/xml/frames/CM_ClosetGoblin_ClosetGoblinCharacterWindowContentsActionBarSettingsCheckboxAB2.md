# Frame ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB2

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB2
- Raw Name: $parentCheckboxAB2
- Type: Button
- Parent: ClosetGoblinCharacterWindowContentsActionBarSettings
- Parent Type: Window
- Inherits: EA_Button_DefaultCheckBox
- Template: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.xml:0`

## Children

- none

## Child Element Types

- none

## Structural Child Types

- Anchors
- EventHandlers

## Composition Pattern

```xml
<Button inherits="EA_Button_DefaultCheckBox" layer="overlay" name="...">
  <Anchors>
    <Anchor point="bottom" relativePoint="top" relativeTo="$parentLabelAB2">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.H..."/>
  </EventHandlers>
</Button>
```

## Attributes

- inherits: EA_Button_DefaultCheckBox
- layer: overlay
- name: $parentCheckboxAB2

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.HotbarChangeToggled2 |

# Frame ClosetGoblinZoneWindowContentsChangeZoneButton

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinZoneWindowContentsChangeZoneButton
- Raw Name: $parentChangeZoneButton
- Type: Button
- Parent: ClosetGoblinZoneWindowContents
- Parent Type: Window
- Inherits: EA_Button_DefaultResizeable
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
<Button font="font_chat_text" inherits="EA_Button_DefaultResizeable" name="..." textalign="center">
  <Anchors>
    <Anchor point="topright" relativePoint="topleft" relativeTo="$parentZoneOption">
      <AbsPoint x="0" y="-10"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinZoneWindow.OnClic..."/>
  </EventHandlers>
</Button>
```

## Attributes

- font: font_chat_text
- inherits: EA_Button_DefaultResizeable
- name: $parentChangeZoneButton
- textalign: center

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinZoneWindow.OnClickChangeZoneOptionButton |

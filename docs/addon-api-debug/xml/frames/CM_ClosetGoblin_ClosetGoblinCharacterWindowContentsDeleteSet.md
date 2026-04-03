# Frame ClosetGoblinCharacterWindowContentsDeleteSet

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsDeleteSet
- Raw Name: $parentDeleteSet
- Type: Button
- Parent: ClosetGoblinCharacterWindowContents
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
    <Anchor point="right" relativePoint="left" relativeTo="$parentNewSet">
      <AbsPoint x="10" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.O..."/>
  </EventHandlers>
</Button>
```

## Attributes

- font: font_chat_text
- inherits: EA_Button_DefaultResizeable
- name: $parentDeleteSet
- textalign: center

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.OnClickDeleteSetButton |

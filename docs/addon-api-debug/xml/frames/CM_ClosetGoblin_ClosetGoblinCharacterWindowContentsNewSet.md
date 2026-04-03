# Frame ClosetGoblinCharacterWindowContentsNewSet

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsNewSet
- Raw Name: $parentNewSet
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
    <Anchor point="topleft" relativePoint="bottomleft" relativeTo="$parentSetList">
      <AbsPoint x="48" y="-105"/>
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
- name: $parentNewSet
- textalign: center

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.OnClickNewSetButton |

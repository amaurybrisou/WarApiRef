# Frame ClosetGoblinCharacterWindowContentsSortNameButton

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsSortNameButton
- Raw Name: $parentSortNameButton
- Type: Button
- Parent: ClosetGoblinCharacterWindowContents
- Parent Type: Window
- Inherits: EA_Button_ListSort
- Template: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.xml:0`

## Children

- none

## Child Element Types

- none

## Structural Child Types

- Anchors
- EventHandlers
- Size

## Composition Pattern

```xml
<Button id="1" inherits="EA_Button_ListSort" name="...">
  <Size>
    <AbsPoint x="275" y="40"/>
  </Size>
  <Anchors>
    <Anchor point="topleft" relativePoint="topleft" relativeTo="$parentSetListBackground">
      <AbsPoint x="5" y="5"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.O..."/>
  </EventHandlers>
</Button>
```

## Attributes

- id: 1
- inherits: EA_Button_ListSort
- name: $parentSortNameButton

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.OnClickSortNameButton |

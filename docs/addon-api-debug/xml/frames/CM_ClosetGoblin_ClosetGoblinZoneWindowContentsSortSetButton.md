# Frame ClosetGoblinZoneWindowContentsSortSetButton

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinZoneWindowContentsSortSetButton
- Raw Name: $parentSortSetButton
- Type: Button
- Parent: ClosetGoblinZoneWindowContents
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
<Button id="2" inherits="EA_Button_ListSort" name="...">
  <Size>
    <AbsPoint x="366" y="40"/>
  </Size>
  <Anchors>
    <Anchor point="topright" relativePoint="topleft" relativeTo="$parentSortZoneButton">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.O..."/>
  </EventHandlers>
</Button>
```

## Attributes

- id: 2
- inherits: EA_Button_ListSort
- name: $parentSortSetButton

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.OnClickSortSetButton |

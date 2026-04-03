# Frame ClosetGoblinZoneWindowContentsSetList

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinZoneWindowContentsSetList
- Raw Name: $parentSetList
- Type: ListBox
- Parent: ClosetGoblinZoneWindowContents
- Parent Type: Window
- Inherits: none
- Template: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.xml:0`

## Children

- none

## Child Element Types

- none

## Structural Child Types

- Anchors
- ListData: populationfunction, table
- Size

## Composition Pattern

```xml
<ListBox layer="secondary" name="..." rowcount="100" rowdef="ClosetGoblinZoneRow" rowspacing="1" scrollbar="EA_ScrollBar_DefaultVerticalC..." visiblerows="20">
  <Size>
    <AbsPoint x="486" y="540"/>
  </Size>
  <Anchors>
    <Anchor point="bottomleft" relativePoint="topleft" relativeTo="$parentSortZoneButton">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
  <ListData populationfunction="ClosetGoblinZoneWindow.Update..." table="ClosetGoblinZoneWindow.associ...">
    <ListColumns>
      <ListColumn format="wstring" variable="zone" windowname="Zone"/>
      <ListColumn format="wstring" variable="set" windowname="Set"/>
    </ListColumns>
  </ListData>
</ListBox>
```

## Attributes

- layer: secondary
- name: $parentSetList
- rowcount: 100
- rowdef: ClosetGoblinZoneRow
- rowspacing: 1
- scrollbar: EA_ScrollBar_DefaultVerticalChain
- visiblerows: 20

## Handlers

| Event | Function |
| --- | --- |

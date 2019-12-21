// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package service

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - MediaID
//  - Name
//  - UpdateTime
//  - URL
type Item struct {
  MediaID string `thrift:"media_id,1" db:"media_id" json:"media_id"`
  Name string `thrift:"name,2" db:"name" json:"name"`
  UpdateTime int64 `thrift:"update_time,3" db:"update_time" json:"update_time"`
  URL string `thrift:"url,4" db:"url" json:"url"`
}

func NewItem() *Item {
  return &Item{}
}


func (p *Item) GetMediaID() string {
  return p.MediaID
}

func (p *Item) GetName() string {
  return p.Name
}

func (p *Item) GetUpdateTime() int64 {
  return p.UpdateTime
}

func (p *Item) GetURL() string {
  return p.URL
}
func (p *Item) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Item)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.MediaID = v
}
  return nil
}

func (p *Item)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Name = v
}
  return nil
}

func (p *Item)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.UpdateTime = v
}
  return nil
}

func (p *Item)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.URL = v
}
  return nil
}

func (p *Item) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Item"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Item) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("media_id", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:media_id: ", p), err) }
  if err := oprot.WriteString(string(p.MediaID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.media_id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:media_id: ", p), err) }
  return err
}

func (p *Item) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:name: ", p), err) }
  if err := oprot.WriteString(string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:name: ", p), err) }
  return err
}

func (p *Item) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("update_time", thrift.I64, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:update_time: ", p), err) }
  if err := oprot.WriteI64(int64(p.UpdateTime)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.update_time (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:update_time: ", p), err) }
  return err
}

func (p *Item) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("url", thrift.STRING, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:url: ", p), err) }
  if err := oprot.WriteString(string(p.URL)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.url (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:url: ", p), err) }
  return err
}

func (p *Item) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Item(%+v)", *p)
}

// Attributes:
//  - TotalCount
//  - ItemCount
//  - Item
type Res struct {
  TotalCount string `thrift:"total_count,1" db:"total_count" json:"total_count"`
  ItemCount string `thrift:"item_count,2" db:"item_count" json:"item_count"`
  Item []*Item `thrift:"item,3" db:"item" json:"item"`
}

func NewRes() *Res {
  return &Res{}
}


func (p *Res) GetTotalCount() string {
  return p.TotalCount
}

func (p *Res) GetItemCount() string {
  return p.ItemCount
}

func (p *Res) GetItem() []*Item {
  return p.Item
}
func (p *Res) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Res)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.TotalCount = v
}
  return nil
}

func (p *Res)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.ItemCount = v
}
  return nil
}

func (p *Res)  ReadField3(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*Item, 0, size)
  p.Item =  tSlice
  for i := 0; i < size; i ++ {
    _elem0 := &Item{}
    if err := _elem0.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
    }
    p.Item = append(p.Item, _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *Res) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Res"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Res) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("total_count", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:total_count: ", p), err) }
  if err := oprot.WriteString(string(p.TotalCount)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.total_count (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:total_count: ", p), err) }
  return err
}

func (p *Res) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("item_count", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:item_count: ", p), err) }
  if err := oprot.WriteString(string(p.ItemCount)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.item_count (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:item_count: ", p), err) }
  return err
}

func (p *Res) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("item", thrift.LIST, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:item: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.Item {
    if err := v.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
    }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:item: ", p), err) }
  return err
}

func (p *Res) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Res(%+v)", *p)
}

// Attributes:
//  - Type
//  - Offset
//  - Count
type WxParm struct {
  Type string `thrift:"type,1" db:"type" json:"type"`
  Offset int32 `thrift:"offset,2" db:"offset" json:"offset"`
  Count int32 `thrift:"count,3" db:"count" json:"count"`
}

func NewWxParm() *WxParm {
  return &WxParm{}
}


func (p *WxParm) GetType() string {
  return p.Type
}

func (p *WxParm) GetOffset() int32 {
  return p.Offset
}

func (p *WxParm) GetCount() int32 {
  return p.Count
}
func (p *WxParm) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *WxParm)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Type = v
}
  return nil
}

func (p *WxParm)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Offset = v
}
  return nil
}

func (p *WxParm)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Count = v
}
  return nil
}

func (p *WxParm) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("WxParm"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *WxParm) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("type", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:type: ", p), err) }
  if err := oprot.WriteString(string(p.Type)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.type (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:type: ", p), err) }
  return err
}

func (p *WxParm) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("offset", thrift.I32, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:offset: ", p), err) }
  if err := oprot.WriteI32(int32(p.Offset)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.offset (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:offset: ", p), err) }
  return err
}

func (p *WxParm) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("count", thrift.I32, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:count: ", p), err) }
  if err := oprot.WriteI32(int32(p.Count)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.count (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:count: ", p), err) }
  return err
}

func (p *WxParm) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("WxParm(%+v)", *p)
}

// Attributes:
//  - VoiceCount
//  - VideoCount
//  - ImageCount
//  - NewsCount_
type MaCount struct {
  VoiceCount int64 `thrift:"voice_count,1" db:"voice_count" json:"voice_count"`
  VideoCount int64 `thrift:"video_count,2" db:"video_count" json:"video_count"`
  ImageCount int64 `thrift:"image_count,3" db:"image_count" json:"image_count"`
  NewsCount_ int64 `thrift:"news_count,4" db:"news_count" json:"news_count"`
}

func NewMaCount() *MaCount {
  return &MaCount{}
}


func (p *MaCount) GetVoiceCount() int64 {
  return p.VoiceCount
}

func (p *MaCount) GetVideoCount() int64 {
  return p.VideoCount
}

func (p *MaCount) GetImageCount() int64 {
  return p.ImageCount
}

func (p *MaCount) GetNewsCount_() int64 {
  return p.NewsCount_
}
func (p *MaCount) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MaCount)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.VoiceCount = v
}
  return nil
}

func (p *MaCount)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.VideoCount = v
}
  return nil
}

func (p *MaCount)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.ImageCount = v
}
  return nil
}

func (p *MaCount)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.NewsCount_ = v
}
  return nil
}

func (p *MaCount) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("MaCount"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MaCount) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("voice_count", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:voice_count: ", p), err) }
  if err := oprot.WriteI64(int64(p.VoiceCount)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.voice_count (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:voice_count: ", p), err) }
  return err
}

func (p *MaCount) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("video_count", thrift.I64, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:video_count: ", p), err) }
  if err := oprot.WriteI64(int64(p.VideoCount)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.video_count (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:video_count: ", p), err) }
  return err
}

func (p *MaCount) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("image_count", thrift.I64, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:image_count: ", p), err) }
  if err := oprot.WriteI64(int64(p.ImageCount)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.image_count (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:image_count: ", p), err) }
  return err
}

func (p *MaCount) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("news_count", thrift.I64, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:news_count: ", p), err) }
  if err := oprot.WriteI64(int64(p.NewsCount_)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.news_count (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:news_count: ", p), err) }
  return err
}

func (p *MaCount) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MaCount(%+v)", *p)
}

// Attributes:
//  - Type
//  - MediaID
//  - CreatedAt
//  - Errcode
//  - Errmsg
type WxImage struct {
  Type string `thrift:"type,1" db:"type" json:"type"`
  MediaID string `thrift:"media_id,2" db:"media_id" json:"media_id"`
  CreatedAt int64 `thrift:"created_at,3" db:"created_at" json:"created_at"`
  Errcode int32 `thrift:"Errcode,4" db:"Errcode" json:"Errcode"`
  Errmsg string `thrift:"Errmsg,5" db:"Errmsg" json:"Errmsg"`
}

func NewWxImage() *WxImage {
  return &WxImage{}
}


func (p *WxImage) GetType() string {
  return p.Type
}

func (p *WxImage) GetMediaID() string {
  return p.MediaID
}

func (p *WxImage) GetCreatedAt() int64 {
  return p.CreatedAt
}

func (p *WxImage) GetErrcode() int32 {
  return p.Errcode
}

func (p *WxImage) GetErrmsg() string {
  return p.Errmsg
}
func (p *WxImage) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
        return err
      }
    case 5:
      if err := p.ReadField5(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *WxImage)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Type = v
}
  return nil
}

func (p *WxImage)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.MediaID = v
}
  return nil
}

func (p *WxImage)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.CreatedAt = v
}
  return nil
}

func (p *WxImage)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.Errcode = v
}
  return nil
}

func (p *WxImage)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.Errmsg = v
}
  return nil
}

func (p *WxImage) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("WxImage"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *WxImage) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("type", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:type: ", p), err) }
  if err := oprot.WriteString(string(p.Type)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.type (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:type: ", p), err) }
  return err
}

func (p *WxImage) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("media_id", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:media_id: ", p), err) }
  if err := oprot.WriteString(string(p.MediaID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.media_id (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:media_id: ", p), err) }
  return err
}

func (p *WxImage) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("created_at", thrift.I64, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:created_at: ", p), err) }
  if err := oprot.WriteI64(int64(p.CreatedAt)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.created_at (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:created_at: ", p), err) }
  return err
}

func (p *WxImage) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("Errcode", thrift.I32, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:Errcode: ", p), err) }
  if err := oprot.WriteI32(int32(p.Errcode)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.Errcode (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:Errcode: ", p), err) }
  return err
}

func (p *WxImage) writeField5(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("Errmsg", thrift.STRING, 5); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:Errmsg: ", p), err) }
  if err := oprot.WriteString(string(p.Errmsg)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.Errmsg (5) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 5:Errmsg: ", p), err) }
  return err
}

func (p *WxImage) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("WxImage(%+v)", *p)
}

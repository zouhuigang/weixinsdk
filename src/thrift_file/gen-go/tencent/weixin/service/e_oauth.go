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
//  - URL
//  - State
type AuthCodeURLData struct {
  URL string `thrift:"url,1" db:"url" json:"url"`
  State string `thrift:"state,2" db:"state" json:"state"`
}

func NewAuthCodeURLData() *AuthCodeURLData {
  return &AuthCodeURLData{}
}


func (p *AuthCodeURLData) GetURL() string {
  return p.URL
}

func (p *AuthCodeURLData) GetState() string {
  return p.State
}
func (p *AuthCodeURLData) Read(iprot thrift.TProtocol) error {
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

func (p *AuthCodeURLData)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.URL = v
}
  return nil
}

func (p *AuthCodeURLData)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.State = v
}
  return nil
}

func (p *AuthCodeURLData) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("AuthCodeURLData"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *AuthCodeURLData) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("url", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:url: ", p), err) }
  if err := oprot.WriteString(string(p.URL)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.url (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:url: ", p), err) }
  return err
}

func (p *AuthCodeURLData) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("state", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:state: ", p), err) }
  if err := oprot.WriteString(string(p.State)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.state (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:state: ", p), err) }
  return err
}

func (p *AuthCodeURLData) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("AuthCodeURLData(%+v)", *p)
}


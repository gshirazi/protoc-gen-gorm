// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/contacts/contacts.proto

package contacts

import context "context"
import errors "errors"
import gorm "github.com/jinzhu/gorm"
import ops "github.com/Infoblox-CTO/ngp.api.toolkit/op/gorm"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ContactORM no comment was provided for message type
type ContactORM struct {
	TenantID     string
	ID           uint64
	FirstName    string
	MiddleName   string
	LastName     string
	EmailAddress string
}

// TableName overrides the default tablename generated by GORM
func (ContactORM) TableName() string {
	return "contacts"
}

// ConvertContactToORM takes a pb object and returns an orm object
func ConvertContactToORM(from Contact) (ContactORM, error) {
	to := ContactORM{}
	var err error
	to.ID = from.Id
	to.FirstName = from.FirstName
	to.MiddleName = from.MiddleName
	to.LastName = from.LastName
	to.EmailAddress = from.EmailAddress
	return to, err
}

// ConvertContactFromORM takes an orm object and returns a pb object
func ConvertContactFromORM(from ContactORM) (Contact, error) {
	to := Contact{}
	var err error
	to.Id = from.ID
	to.FirstName = from.FirstName
	to.MiddleName = from.MiddleName
	to.LastName = from.LastName
	to.EmailAddress = from.EmailAddress
	return to, err
}

////////////////////////// CURDL for objects
// DefaultCreateContact executes a basic gorm create call
func DefaultCreateContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateContact")
	}
	ormObj, err := ConvertContactToORM(*in)
	if err != nil {
		return nil, err
	}
	tenantID, tIDErr := auth.GetTenantID(ctx)
	if tIDErr != nil {
		return nil, tIDErr
	}
	ormObj.TenantID = tenantID
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ConvertContactFromORM(ormObj)
	return &pbResponse, err
}

// DefaultReadContact executes a basic gorm read call
func DefaultReadContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadContact")
	}
	ormParams, err := ConvertContactToORM(*in)
	if err != nil {
		return nil, err
	}
	tenantID, tIDErr := auth.GetTenantID(ctx)
	if tIDErr != nil {
		return nil, tIDErr
	}
	ormParams.TenantID = tenantID
	ormResponse := ContactORM{}
	if err = db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ConvertContactFromORM(ormResponse)
	return &pbResponse, err
}

// DefaultUpdateContact executes a basic gorm update call
func DefaultUpdateContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateContact")
	}
	if exists, err := DefaultReadContact(ctx, &Contact{Id: in.GetId()}, db); err != nil {
		return nil, err
	} else if exists == nil {
		return nil, errors.New("Contact not found")
	}
	ormObj, err := ConvertContactToORM(*in)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ConvertContactFromORM(ormObj)
	return &pbResponse, err
}

func DefaultDeleteContact(ctx context.Context, in *Contact, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteContact")
	}
	ormObj, err := ConvertContactToORM(*in)
	if err != nil {
		return err
	}
	tenantID, tIDErr := auth.GetTenantID(ctx)
	if tIDErr != nil {
		return tIDErr
	}
	ormObj.TenantID = tenantID
	err = db.Where(&ormObj).Delete(&ContactORM{}).Error
	return err
}

// DefaultListContact executes a gorm list call
func DefaultListContact(ctx context.Context, db *gorm.DB) ([]*Contact, error) {
	ormResponse := []ContactORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	tenantID, tIDErr := auth.GetTenantID(ctx)
	if tIDErr != nil {
		return nil, tIDErr
	}
	db = db.Where(&ContactORM{TenantID: tenantID})
	if err := db.Set("gorm:auto_preload", true).Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Contact{}
	for _, responseEntry := range ormResponse {
		temp, err := ConvertContactFromORM(responseEntry)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultStrictUpdateContact clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateContact")
	}
	ormObj, err := ConvertContactToORM(*in)
	if err != nil {
		return nil, err
	}
	tenantID, tIDErr := auth.GetTenantID(ctx)
	if tIDErr != nil {
		return nil, tIDErr
	}
	db = db.Where(&ContactORM{TenantID: tenantID})
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ConvertContactFromORM(ormObj)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

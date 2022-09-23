package models

import (
	internalmodel "github.com/SKF/go-asset-component-client/internal/models"
	"github.com/SKF/go-utility/v2/uuid"
)

type (
	Components struct {
		Components []Component
		Count      int64
	}

	Component struct {
		Origin              Origin
		Manufacturer        string
		AttachedTo          uuid.UUID
		Type                string
		Asset               uuid.UUID
		PositionDescription string
		Designation         string
		ID                  uuid.UUID
		SerialNumber        string
		RotatingRing        string
		ShaftSide           string
		Position            int64
		FixedSpeed          int64
	}

	Origin struct {
		ID       string
		Type     string
		Provider uuid.UUID
	}
)

func (c *Components) FromInternal(components internalmodel.GetAssetComponentsResponse) error {
	for _, comp := range components.Components {
		var component Component
		if err := component.FromInternal(comp); err != nil {
			return err
		}

		c.Components = append(c.Components, component)
	}

	c.Count = int64(components.GetCount())

	return nil
}

func (c *Component) FromInternal(component internalmodel.Component) error {
	origin := component.GetOrigin()
	c.Origin = Origin{
		ID:       origin.Id,
		Type:     origin.Type,
		Provider: uuid.UUID(origin.Provider),
	}
	c.Manufacturer = component.GetManufacturer()
	c.AttachedTo = uuid.UUID(component.GetAttachedTo())
	c.Type = component.GetType()
	c.Asset = uuid.UUID(component.GetAsset())
	c.PositionDescription = component.GetPositionDescription()
	c.Designation = component.GetDesignation()
	c.ID = uuid.UUID(component.GetId())
	c.SerialNumber = component.GetSerialNumber()
	c.RotatingRing = component.GetRotatingRing()
	c.ShaftSide = component.GetShaftSide()
	c.Position = int64(component.GetPosition())
	c.FixedSpeed = int64(component.GetFixedSpeed())

	return nil
}

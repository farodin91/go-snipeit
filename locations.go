// Copyright 2020 Eurac Research. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package snipeit

import (
	"net/http"
)

// LocationOptions specifies a subset of optional query parameters for listing
// locations.
type LocationOptions struct {
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
	Search string `url:"search,omitempty"`
	Sort   string `url:"sort,omitempty"`
	Order  string `url:"order,omitempty"`
}

// Location represents a Snipe-IT location.
type Location struct {
	ID             int64     `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Image          string    `json:"image,omitempty"`
	Address        string    `json:"address,omitempty"`
	Address2       string    `json:"address2,omitempty"`
	City           string    `json:"city,omitempty"`
	State          string    `json:"state,omitempty"`
	Country        string    `json:"country,omitempty"`
	Zip            string    `json:"zip,omitempty"`
	AssetsAssigned int64     `json:"assigned_assets_count,omitempty"`
	Assets         int64     `json:"assets_count,omitempty"`
	Users          int64     `json:"users_count,omitempty"`
	Currency       string    `json:"currency,omitempty"`
	CreatedAt      Timestamp `json:"created_at,omitempty"`
	UpdatedAt      Timestamp `json:"updated_at,omitempty"`
	Parent         struct {
		ID   int64  `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"parent,omitempty"`
	Manager  string     `json:"manager,omitempty"`
	Children []Location `json:"children,omitempty"`
	Actions  struct {
		Update bool
		Delete bool
	} `json:"available_actions,omitempty"`
}

// Locations lists all locations.
//
// Snipe-IT API doc: https://snipe-it.readme.io/reference#locations
func (c *Client) Locations(opt *LocationOptions) ([]*Location, *http.Response, error) {
	return listItems[LocationOptions, Location](c, "locations", opt)
}

// Location by ID.
//
// Snipe-IT API doc: https://snipe-it.readme.io/reference#locations-1
func (c *Client) Location(id int64) (*Location, *http.Response, error) {
	return findItem[Location](c, "locations", id)
}

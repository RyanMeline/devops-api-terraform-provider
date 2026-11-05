// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Engineer is the API representation of an engineer as returned by the
// devops API's GET /engineers endpoint. The json tags match the keys the
// API emits (name, id, email).
type Engineer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CREATE engineer.
// (POST /engineers)
// (ID is automatically created by the API and returned in the response)
func (c *Client) CreateEngineer(engineer Engineer) (*Engineer, error) {
	engineerJSON, err := json.Marshal(engineer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.BaseURL+"/engineers", bytes.NewReader(engineerJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode, body)
	}

	var created Engineer
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, fmt.Errorf("decoding create response: %w", err)
	}
	if created.ID == "" {
		return nil, fmt.Errorf("API response did not include an engineer id")
	}

	return &created, nil
}

// READ engineers.
// (GET /engineers)
func (c *Client) GetEngineers() ([]Engineer, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+"/engineers", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode, body)
	}

	var engineers []Engineer
	if err := json.NewDecoder(resp.Body).Decode(&engineers); err != nil {
		return nil, err
	}

	return engineers, nil
}

// UPDATE engineer.
// (PUT /engineers/{id})
func (c *Client) UpdateEngineer(engineer Engineer) (*Engineer, error) {
	engineerJSON, err := json.Marshal(engineer)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, c.BaseURL+"/engineers/"+engineer.ID, bytes.NewReader(engineerJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode, body)
	}

	var updated Engineer
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, fmt.Errorf("decoding update response: %w", err)
	}

	return &updated, nil
}

// DELETE engineer.
// (DELETE /engineers/{id})
func (c *Client) DeleteEngineer(id string) error {
	req, err := http.NewRequest(http.MethodDelete, c.BaseURL+"/engineers/"+id, nil)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("status: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Environment struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type CreateEnvironmentRequest struct {
	Name string `json:"name"`
}

func (c *CoolifyClient) CreateEnvironment(projectUUID, name string) (*Environment, error) {
	reqBody := CreateEnvironmentRequest{
		Name: name,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/v1/projects/%s/environments", c.Endpoint, projectUUID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create environment, status: %d", resp.StatusCode)
	}

	var resultTemp Environment
	if err := json.NewDecoder(resp.Body).Decode(&resultTemp); err != nil {
		return nil, fmt.Errorf("failed to decode create response: %w", err)
	}

	return &resultTemp, nil
}

func (c *CoolifyClient) GetEnvironment(projectUUID, environmentUUID string) (*Environment, error) {
	url := fmt.Sprintf("%s/api/v1/projects/%s/%s", c.Endpoint, projectUUID, environmentUUID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil // Environment not found
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get environment, status: %d", resp.StatusCode)
	}

	var environment Environment
	if err := json.NewDecoder(resp.Body).Decode(&environment); err != nil {
		return nil, fmt.Errorf("failed to decode get response: %w", err)
	}

	return &environment, nil
}

func (c *CoolifyClient) DeleteEnvironment(projectUUID, environmentUUID string) error {
	url := fmt.Sprintf("%s/api/v1/projects/%s/environments/%s", c.Endpoint, projectUUID, environmentUUID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil // Environment already deleted
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete environment, status: %d", resp.StatusCode)
	}

	return nil
}

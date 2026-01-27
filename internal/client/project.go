package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type Project struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type createResponse struct {
	UUID string `json:"uuid"`
}

func (c *CoolifyClient) CreateProject(name, description string) (*Project, error) {
	reqBody := ProjectRequest{
		Name:        name,
		Description: description,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/v1/projects", c.Endpoint)
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
		return nil, fmt.Errorf("failed to create project, status: %d", resp.StatusCode)
	}

	var resultTemp createResponse
	if err := json.NewDecoder(resp.Body).Decode(&resultTemp); err != nil {
		return nil, fmt.Errorf("failed to decode create response: %w", err)
	}

	return &Project{
		UUID:        resultTemp.UUID,
		Name:        name,
		Description: description,
	}, nil
}

func (c *CoolifyClient) GetProject(uuid string) (*Project, error) {
	url := fmt.Sprintf("%s/api/v1/projects/%s", c.Endpoint, uuid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get project, status: %d", resp.StatusCode)
	}

	var result Project
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *CoolifyClient) UpdateProject(uuid, name, description string) (*Project, error) {
	reqBody := ProjectRequest{
		Name:        name,
		Description: description,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/v1/projects/%s", c.Endpoint, uuid)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to update project, status: %d", resp.StatusCode)
	}

	var result Project
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *CoolifyClient) DeleteProject(uuid string) error {
	url := fmt.Sprintf("%s/api/v1/projects/%s", c.Endpoint, uuid)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete project, status: %d", resp.StatusCode)
	}

	return nil
}

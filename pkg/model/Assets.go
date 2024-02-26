package model

import (
	"errors"
	"fmt"
	"strings"
)

type Asset struct {
	ComputerID    string
	UserID        string
	ApplicationID string
	ComputerType  string
	Comment       string
}

const LAPTOP string = "laptop"
const DESKTOP string = "desktop"

type Config struct {
	BatchSize int `json:"batchsize"`
}
func (c *Config) Validate() error {
	if c.BatchSize <= 0 {
		return errors.New("invalid batch size: must be greater than zero")
	}
	return nil
 }
 
 func (a *Asset) Validate() error {
	 validType := strings.EqualFold(a.ComputerType, LAPTOP) || strings.EqualFold(a.ComputerType, DESKTOP)
	 if !validType {
		 return fmt.Errorf("invalid computer type: %s", a.ComputerType)
	 }
	 return nil
 }
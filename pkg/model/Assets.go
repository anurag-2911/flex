package model

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

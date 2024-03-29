# Asset Processor

## Overview
This Go project processes asset data from CSV files, applying specific business logic to calculate 
the minimum number of application copies required based on asset types and user allocations.
This application is a command line utility and running the application is explained in the 
Running the app section below.


## Features
- Read and process asset data from CSV files.
- Calculate the minimum number of application copies required.
- Supports concurrent processing for efficiency.
- Configurable batch size for processing.

## Getting Started

### Prerequisites
- Go 1.15 or later.

### Installation
Clone the repository to your local machine:

## git clone https://github.com/yourusername/assetprocessor.git
## cd cmd\assetmgmt 
## Build the app:
    go build
## Running the app
    go run main.go -appid=374 -filepath="/path/to/your/csvfile.csv"
    .\assetmgmt.exe -appid 374 -filepath sample-small.csv

## For CI/CD
    .github\workflows\buildnPush.yml
    Above workflow will build,run tests and create and push docker image to the docker hub


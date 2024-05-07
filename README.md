# nasa-neows-cli-tool

The NASA NEOs CLI Tool is a command-line interface designed to interact with 
[NASA's Near Earth Object Web Service (NeoWs)](https://api.nasa.gov/). 
Its primary purpose is to retrieve data about near-Earth objects (NEOs) detected within the last 7 days. 
This tool provides a convenient way for users to access essential information about NEOs directly from their terminal.

[Link to task](https://gist.github.com/ahrytsaienko/6209970ef783ca08227ece4d408ad5d3)

## Features

- Retrieve NEO Data: Fetch information about near-Earth objects detected within the last 7 days.
- Concurrent Data Fetching: Utilize concurrent data fetching to improve performance and efficiency.
- Custom API Key: Allow users to use their own API_KEY for accessing the NeoWs API.
- Formatted Output: Present NEO data in a structured and readable format.


### Returned data format:

```JSON
{
  "total": 1,
  "near_earth_objects": [
    {
      "date": "2024-05-01",
      "id": "2030825",
      "name": "30825 (1990 TG1)",
      "is_potentially_hazardous_asteroid": false
    }
  ]
}
```

## Installation

To install the NASA NEOs CLI Tool, follow these steps:

### Clone the repository

```sh
git clone https://github.com/VladyslavKukharchuk/nasa-neows-cli-tool.git
```


### Navigate to the project directory

```sh
cd nasa-neows-cli-tool
```


### Building the Docker Image

```sh
docker build --tag nasa-neows-cli-tool .
```
This command utilizes your Dockerfile and all resources in the current directory to build the image.
Be sure to wait for this process to complete.


### Run the Docker Container

```sh
docker run --name nasa-neows-cli-tool -e API_KEY=<your API key> nasa-neows-cli-tool
```

Replace `<your API key>` with your actual API key.


## Usage

### Start container to get data about near-Earth objects (NEOs) detected within the last 7 days

```sh
docker start -a nasa-neows-cli-tool
```


## Run Tests

```sh
API_KEY=<your API key> go test ./...
```
We need an API key to run the tests

Replace `<your API key>` with your actual API key.
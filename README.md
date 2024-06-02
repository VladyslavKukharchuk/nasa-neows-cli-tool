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


### Using Makefile Commands

The project includes a Makefile to streamline the build, run, and test processes. 
Below are the available commands.

### Building

Build the Docker Image and Run the Container

```sh
make build API_KEY=<your API key>
```

This command builds the Docker image and runs the Docker container. 
Replace `<your API key>` with your actual API key.


## Usage

This command starts the Docker container to get data about 
near-Earth objects (NEOs) detected within the last 7 days.

```sh
make run
```


## Run Tests

```sh
make test API_KEY=<your API key>
```

We need an API key to run the tests
Replace `<your API key>` with your actual API key.


## Remove Docker Container and Image

```sh
make remove
```


## Dev

### Run in Development Mode

```sh
make run-dev API_KEY=<your API key>
```

We need an API key to run app.
Replace `<your API key>` with your actual API key.


### Build the Binary

```sh
make build-binary
```

This command builds the Go binary.


### Run the Binary

```sh
make run-binary API_KEY=<your API key>
```

We need an API key to run the binary.
Replace `<your API key>` with your actual API key.


### Remove the Binary

```sh
make remove-binary
```

This command removes the built Go binary.
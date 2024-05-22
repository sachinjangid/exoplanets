# ExoPlanets Service <br>
ExoPlanets service provides 
- Support to add, update, view and delete the exoplanet data.
- Provides Calculations for fuel cost units  
- Has flexibility to add new exoplanet type and manage other calculations 

# Developer Guidelines

## Setup

### Prerequisites

- Golang (version 1.19 or later)

To set up your development environment, follow these steps:

1. Clone the repository: <br>
`git clone git@github.com:sachinjangid/exoplanets.git` <br>

2. Install dependencies: <br>
`go mod download` <br>

3. Start the development server <br>
`go run main.go` <br>

# API Endpoints

- `POST /exoplanet`

  Create new Exoplanet

  
**Request Payload:**

```  
{
    "name": "My Exoplanet 1",
    "description": "Long Distance from earth",
    "distance": 34,
    "radius": 3.1,
    "mass": 6.3,
    "type": "Terrestrial"
}
```
  **Response**
```
{
    "message": "Exoplanet Stored Successfully!!",
    "payload": null,
    "status": 201
}
```
  **Validation Error Response Example**
```
{
    "message": "Validation Error",
    "payload": {
        "error": [
            "'Name' is a required field",
            "'Distance' is a required field",
            "'Radius' is a required field",
            "'Mass' is invalid or missing",
            "'Type' is invalid or missing"
        ]
    },
    "status": 400
}
```

- `PUT /exoplanet/:exoplanetId`

  Update existing Exoplanet

  
**Request Payload:**

```  
{
    "name": "Updated Exoplanet Name",
    "description": "Updated Description",
    "distance": 34,
    "radius": 3.1,
    "mass": 6.3,
    "type": "Terrestrial"
}
```
  **Response**
```
{
    "message": "Exoplanet Updated Successfully!!",
    "payload": null,
    "status": 200
}
```

- `GET /exoplanet/:exoplanetId`

  Get individual Exoplanet

  **Response**
```
{
    "message": "Exoplanet Fetched Successfully!!",
    "payload": {
        "id": 1,
        "name": "Updated Exoplanet Name",
        "description": "Updated Description",
        "distance": 34,
        "radius": 3.1,
        "mass": 6.3,
        "type": "Terrestrial"
    },
    "status": 200
}
```


- `GET /exoplanet/list`

  Get Exoplanet List

  **Response**
```
{
    "message": "Exoplanet Fetched Successfully!!",
    "payload": [
        {
            "id": 1,
            "name": "Updated Exoplanet Name",
            "description": "Updated Description",
            "distance": 34,
            "radius": 3.1,
            "mass": 6.3,
            "type": "Terrestrial"
        }
    ],
    "status": 200
}
```


- `DELETE /exoplanet/:exoplanetId`

  Delete existing Exoplanet

```
{
    "message": "Exoplanet Deleted Successfully!!",
    "payload": null,
    "status": 200
}
```


- `POST /exoplanet/:exoplanetId/fuel/estimation`

  Create new Exoplanet

  **Request Payload:**

```  
{
    "crewSize": 4,
}
```

  **Response**

```  
{
    "message": "Calcuation Done!!",
    "payload": {
        "fuelCost": 316.44
    },
    "status": 200
}
```
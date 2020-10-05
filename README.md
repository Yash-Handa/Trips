# Trips

[![Go Report Card](https://goreportcard.com/badge/github.com/Yash-Handa/Trips)](https://goreportcard.com/report/github.com/Yash-Handa/Trips)

A GraphQL-Golang backend to emulate Uber like Car booking services

## Stack

**Language**: Golang v1.15.0
**WebServer**: gin-gonic/gin v1.6.3
**GraphQL Layer**: 99designs/gqlgen v0.13.0
**DataBase**: PostgreSQL v13.0.0
**DB ORM**: go-pg/pg/v10 v10.3.2

## Features

- Registration and login of Users
- User Authentication using JWT
- Booking of trip
- Cancellation of trip
- Start a trip
- Stop a trip
- Get logged user info
- Get logged in user's trips
- Subscription to all currently available cabs in your area

The entire schema has been made public to explore at: https://festive-poitras-059dd4.netlify.app (created by 2fd/graphdoc)

## Some Sample queries

### Register User

**Query**:

```graphql
mutation {
  register(
    input: {
      email: "yash@yahoo.com"
      password: "123456"
      confirmPassword: "123456"
      firstName: "Yash"
      lastName: "Handa"
    }
  ) {
    authToken {
      accessToken
      expiredAt
    }
    user {
      createdAt
      firstName
      email
      lastName
    }
  }
}
```

**Result**:

```json
{
  "data": {
    "register": {
      "authToken": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDI0ODk3OTksImp0aSI6IjhlZGIxYzQxLTAzY2YtNGRjYi05YmVkLTQ2ZDE4NTlhZjI5NiIsImlhdCI6MTYwMTg4NDk5OSwiaXNzIjoiVHJpcHMifQ.DA6_yCbKb_1KkoKoxtat2ftuMSSsrEpkvs8pKzmggEU",
        "expiredAt": "2020-10-12T13:33:19+05:30"
      },
      "user": {
        "createdAt": "2020-10-05T08:03:19Z",
        "firstName": "Yash",
        "email": "yash@yahoo.com",
        "lastName": "Handa"
      }
    }
  }
}
```

### Login

**Query**:

```graphql
mutation{
  login(input: {
    email: "yash@yahoo.com"
    password: "123456"
  }) {
    authToken {
      accessToken
      expiredAt
    }
    user {
      firstName
      lastName
      email
      createdAt
    }
  }
}
```

**Result**:

```json
{
  "data": {
    "login": {
      "authToken": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDI0ODk5NzIsImp0aSI6IjhlZGIxYzQxLTAzY2YtNGRjYi05YmVkLTQ2ZDE4NTlhZjI5NiIsImlhdCI6MTYwMTg4NTE3MiwiaXNzIjoiVHJpcHMifQ._RoXKDw66R6p43TD1lYoR0Bz2YRFFj6U1ZgxEHBTVyo",
        "expiredAt": "2020-10-12T13:36:12+05:30"
      },
      "user": {
        "firstName": "Yash",
        "lastName": "Handa",
        "email": "yash@yahoo.com",
        "createdAt": "2020-10-05T08:03:19Z"
      }
    }
  }
}
```

### Book a trip

**Query**:

```graphql
mutation {
  bookTrip(
    input: {
      cabType: SUV
      pickup: { Lat: 1.11, Lon: 1.11 }
      destination: { Lat: -2.22, Lon: -2.22 }
    }
  ) {
    id
    pickup {
      Lat
      Lon
    }
    destination {
      Lat
      Lon
    }
    amount {
      currency
      amount
    }
    cab {
      id
      type
      workingAC
      namePlate
      number
      model
      pic
      driver {
        firstName
        lastName
        rating
        gender
        pic
        id
      }
    }
    canceled {
      cancel
      reason
    }
    completed
    user {
      email
      firstName
      lastName
    }
  }
}
```

**Headers**:

```json
{
  "Authorization": "BEARER eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDI0ODk5NzIsImp0aSI6IjhlZGIxYzQxLTAzY2YtNGRjYi05YmVkLTQ2ZDE4NTlhZjI5NiIsImlhdCI6MTYwMTg4NTE3MiwiaXNzIjoiVHJpcHMifQ._RoXKDw66R6p43TD1lYoR0Bz2YRFFj6U1ZgxEHBTVyo"
}
```

**Result**:

```json
{
  "data": {
    "bookTrip": {
      "id": "29a1b4d2-8065-4f48-9c07-fcf6c8f783da",
      "pickup": {
        "Lat": 1.11,
        "Lon": 1.11
      },
      "destination": {
        "Lat": -2.22,
        "Lon": -2.22
      },
      "amount": {
        "currency": "INR",
        "amount": 500
      },
      "cab": {
        "id": "5596de5e-dad9-4fb1-a43c-60110d6bee0e",
        "type": "SUV",
        "workingAC": true,
        "namePlate": "MH 01A B 4321",
        "number": 4321,
        "model": "Honda CR-V",
        "pic": "/raw/assets/cabs/test_Honda_CR-V.jpg",
        "driver": {
          "firstName": "Goku",
          "lastName": null,
          "rating": 4.9,
          "gender": "M",
          "pic": "/raw/assets/drivers/test_Goku.jpg",
          "id": "acb336c4-15f2-47f6-8080-d744173a1f3e"
        }
      },
      "canceled": null,
      "completed": false,
      "user": {
        "email": "yash@yahoo.com",
        "firstName": "Yash",
        "lastName": "Handa"
      }
    }
  }
}
```

### Subscription of Nearby Cabs

**Query**:

```graphql
subscription {
  nearbyCabs(
    input: { type: Luxury, radius: 5, curLocation: { Lat: 1.11, Lon: 1.11 } }
  ) {
    event
    location {
      Lat
      Lon
    }
    cabID
  }
}
```

**Result**:

```json
15 secs ago----

{
  "data": {
    "nearbyCabs": [
      {
        "event": "ENTER",
        "location": {
          "Lat": 1.1099687803293943,
          "Lon": 1.1099603389909027
        },
        "cabID": "f2d848a9-b1af-4533-93fd-2f8730034fe5"
      },
      {
        "event": "ENTER",
        "location": {
          "Lat": 1.1100332357143243,
          "Lon": 1.1100122799889032
        },
        "cabID": "230b6f63-d66d-4f99-8861-a666f842f57e"
      }
    ]
  }
}

10 secs ago ----

{
  "data": {
    "nearbyCabs": [
      {
        "event": "ENTER",
        "location": {
          "Lat": 1.1099687803293943,
          "Lon": 1.1099603389909027
        },
        "cabID": "f2d848a9-b1af-4533-93fd-2f8730034fe5"
      },
      {
        "event": "ENTER",
        "location": {
          "Lat": 1.1100332357143243,
          "Lon": 1.1100122799889032
        },
        "cabID": "230b6f63-d66d-4f99-8861-a666f842f57e"
      }
    ]
  }
}

5 secs ago ----

{
  "data": {
    "nearbyCabs": [
      {
        "event": "MOVE",
        "location": {
          "Lat": 1.1099571952365093,
          "Lon": 1.1100279253574976
        },
        "cabID": "f2d848a9-b1af-4533-93fd-2f8730034fe5"
      },
      {
        "event": "MOVE",
        "location": {
          "Lat": 1.1099805083540966,
          "Lon": 1.1099723508462003
        },
        "cabID": "230b6f63-d66d-4f99-8861-a666f842f57e"
      }
    ]
  }
}

0 sec ago ----

{
  "data": {
    "nearbyCabs": [
      {
        "event": "MOVE",
        "location": {
          "Lat": 1.1100474882093294,
          "Lon": 1.1099900327014336
        },
        "cabID": "f2d848a9-b1af-4533-93fd-2f8730034fe5"
      },
      {
        "event": "MOVE",
        "location": {
          "Lat": 1.1099936738395368,
          "Lon": 1.1100131956816162
        },
        "cabID": "230b6f63-d66d-4f99-8861-a666f842f57e"
      }
    ]
  }
}
```

## Further Scope of improvement

- **In-memory caching**: To improve the user authentication time and no need to hit the DB for the same, use in-memory cache like **redis** or **memcache**.

- **Use Better system monitoring service**: The system generated logs and matrics could be channeled to a centralized service like **prometheus**

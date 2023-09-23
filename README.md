# DOCUMENTATION

### PARTIES
**Create Party**
URL : http://localhost:5000/api/v1/party
Method: POST
Request Body: 
```json
{
  "nama": "Redbull"
}
```

Response Body :
```json
{
  "code": 200,
  "data": {
    "id": 1,
    "name": "Redbull",
    "created_at": "2023-09-23T20:19:58.363236+07:00",
    "updated_at": "2023-09-23T20:19:58.363236+07:00"
  }
}
```

**Find Parties**
URL : http://localhost:5000/api/v1/parties
Method: GET
Response Body :
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "name": "Wota",
      "created_at": "2023-09-23T18:10:44.944+07:00",
      "updated_at": "2023-09-23T18:10:44.944+07:00"
    },
    {
      "id": 2,
      "name": "Wiboo",
      "created_at": "2023-09-23T18:10:56.082+07:00",
      "updated_at": "2023-09-23T18:10:56.082+07:00"
    }
  ]
}  
```

**Get Party by ID**
URL : http://localhost:5000/api/v1/party/2
Method: GET
Response Body :
```json
{
  "code": 200,
  "data": {
    "id": 2,
    "name": "Wiboo",
    "created_at": "2023-09-23T18:10:56.082+07:00",
    "updated_at": "2023-09-23T18:10:56.082+07:00"
  }
}  
```

**Update Party by ID**
URL : http://localhost:5000/api/v1/party/2
Method: PATCH
Request Body: 
```json
{
  "nama": "Nusantara United"
}
```

Response Body :
```json
{
  "code": 200,
  "data": {
    "id": 2,
    "name": "Nusantara United",
    "created_at": "2023-09-23T18:10:56.082+07:00",
    "updated_at": "2023-09-23T20:32:27.126+07:00"
  }
}  
```

**Delete Party by ID**
URL : http://localhost:5000/api/v1/party/2
Method: DELETE
Response Body :
```json
{
  "code": 200,
  "data": {
    "id": 5,
    "name": "Weboo",
    "created_at": "2023-09-23T20:19:58.363+07:00",
    "updated_at": "2023-09-23T20:19:58.363+07:00"
  }
} 
```

### CAPRES
**Create Paslon**
URL : http://localhost:5000/api/v1/paslon
Method: POST
Request Body: 
```json
{
  "name": "Jhon doe",
  "visi": "merebut isekai dari pemerintahan dunia",
  "party": [2, 3]
}
```

Response Body :
```json
{
  "code": 200,
  "data": {
    "id": 3,
    "name": "Rebbeca eltra",
    "visi": "merebut isekai dari pemerintahan dunia",
    "party": [
      {
        "id": 2,
        "name": "Nusantara United"
      },
      {
        "id": 3,
        "name": "Muhammadiah City"
      }
    ],
    "created_at": "2023-09-23T21:36:44.446+07:00",
    "updated_at": "2023-09-23T21:36:44.446+07:00"
  }
}
```

**Find Paslon**
URL : http://localhost:5000/api/v1/paslons
Method: GET
Response Body :
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "name": "Jhon doe",
      "visi": "merebut isekai dari pemerintahan dunia",
      "party": [
          {
            "id": 1,
            "name": "Wota"
          },
          {
            "id": 3,
            "name": "Wibooo"
          }
      ],
      "created_at": "2023-09-23T18:11:54.223+07:00",
      "updated_at": "2023-09-23T18:11:54.223+07:00"
    },
    {
      "id": 2,
      "name": "Rebbecca Eltra",
      "visi": "merebut isekai dari pemerintahan dunia",
      "party": [
        {
          "id": 2,
          "name": "Nusantara United"
        },
        {
          "id": 3,
          "name": "Muhammadiah City"
        }
      ],
      "created_at": "2023-09-23T19:18:26.666+07:00",
      "updated_at": "2023-09-23T19:18:26.666+07:00"
    }
  ]
}
```

**Get Paslon by ID**
URL : http://localhost:5000/api/v1/paslon/2
Method: GET
Response Body :
```json
{
  "id": 2,
  "name": "Rebbecca Eltra",
  "visi": "merebut isekai dari pemerintahan dunia",
  "party": [
    {
      "id": 2,
      "name": "Nusantara United"
    },
    {
      "id": 3,
      "name": "Muhammadiah City"
    }
  ],
  "created_at": "2023-09-23T19:18:26.666+07:00",
  "updated_at": "2023-09-23T19:18:26.666+07:00"
} 
```

**Update Party by ID**
URL : http://localhost:5000/api/v1/paslon/2
Method: PATCH
Request Body: 
```json
{
  "nama": "Rebbecca Eltra v2.0"
}
```

Response Body :
```json
{
  "id": 2,
  "name": "Rebbecca Eltra v2.0",
  "visi": "merebut isekai dari pemerintahan dunia",
  "party": [
    {
      "id": 2,
      "name": "Nusantara United"
    },
    {
      "id": 3,
      "name": "Muhammadiah City"
    }
  ],
  "created_at": "2023-09-23T19:18:26.666+07:00",
  "updated_at": "2023-09-23T19:18:26.666+07:00"
} 
```

**Delete Paslon by ID**
URL : http://localhost:5000/api/v1/paslon/2
Method: DELETE
Response Body :
```json
{
  "id": 2,
  "name": "Rebbecca Eltra v2.0",
  "visi": "merebut isekai dari pemerintahan dunia",
  "party": [
    {
      "id": 2,
      "name": "Nusantara United"
    },
    {
      "id": 3,
      "name": "Muhammadiah City"
    }
  ],
  "created_at": "2023-09-23T19:18:26.666+07:00",
  "updated_at": "2023-09-23T19:18:26.666+07:00"
} 
```


### VOTES
**Create Vote**
URL : http://localhost:5000/api/v1/vote
Method: POST
Request Body: 
```json
{
  "paslon_id": 3,
  "voter_name": "Obito"
}
```

Response Body :
```json
{
  "code": 200,
  "data": {
    "id": 3,
    "voter_name": "Obito",
    "paslon": {
      "id": 15,
      "name": "Rebbecca Eltra v2.1",
      "visi": "merebut isekai dari pemerintahan dunia"
    }
  }
}
```

**Find Votes**
URL : http://localhost:5000/api/v1/votes
Method: GET
Response Body :
```json
{
  "code": 200,
  "data": [
    {
      "id": 8,
      "voter_name": "Dandi",
      "paslon": {
        "id": 1,
        "name": "Jhon doe",
        "visi": "merebut isekai dari pemerintahan dunia"
      }
    },
    {
      "id": 9,
      "voter_name": "Dandi",
      "paslon": {
        "id": 1,
        "name": "Jhon doe",
        "visi": "merebut isekai dari pemerintahan dunia"
      }
    },
    {
      "id": 10,
      "voter_name": "Dandi",
      "paslon": {
        "id": 1,
        "name": "Jhon doe",
        "visi": "merebut isekai dari pemerintahan dunia"
      }
    },
    {
      "id": 13,
      "voter_name": "Obito",
      "paslon": {
        "id": 15,
        "name": "Rebbecca Eltra v2.1",
        "visi": "merebut isekai dari pemerintahan dunia"
      }
    },
    {
      "id": 14,
      "voter_name": "Newha",
      "paslon": {
        "id": 15,
        "name": "Rebbecca Eltra v2.1",
        "visi": "merebut isekai dari pemerintahan dunia"
      }
    },
    {
      "id": 15,
      "voter_name": "Byakuya",
      "paslon": {
        "id": 15,
        "name": "Rebbecca Eltra v2.1",
        "visi": "merebut isekai dari pemerintahan dunia"
      }
    }
  ]
}
```





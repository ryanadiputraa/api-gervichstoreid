# GERVICHSTORE.ID - API

Backend API for GERVICHSTORE.ID website.

## API SPEC

---

### Products

### Get All Products

- Method : `GET`
- Endpoint : `/api/products/`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
- response :

```json
{
  "status": "Success",
  "code": 200,
  "error": "",
  "data": [
    {
      "product_id": "uuid",
      "product_name": "String",
      "price": "Number",
      "stock": "Numeber",
      "created_at": "TimeStamp",
      "updated_at": "TimeStamp"
    }
  ]
}
```

### Get Product By ID

- Method : `GET`
- Endpoint : `/api/products/:product_id`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
- response :

```json
{
  "status": "Success",
  "code": 200,
  "error": "",
  "data": {
    "product_id": "uuid",
    "product_name": "String",
    "price": "Number",
    "stock": "Numeber",
    "created_at": "TimeStamp",
    "updated_at": "TimeStamp"
  }
}
```

### Create Product

- Method : `POST`
- Endpoint : `/api/products/`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
- body :

```json
{
  "product_name": "String",
  "price": "Number",
  "stock": "Numeber"
}
```

- response :

```json
{
  "status": "Success",
  "code": 201,
  "error": "",
  "data": {
    "message": "String"
  }
}
```

### Update Product

- Method : `PUT`
- Endpoint : `/api/products/`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
- body :

```json
{
  "product_name": "String",
  "price": "Number",
  "stock": "Numeber"
}
```

- response :

```json
{
  "status": "Success",
  "code": 200,
  "error": "",
  "data": {
    "message": "String"
  }
}
```

### Delete Product

- Method : `DELETE`
- Endpoint : `/api/products/`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
- response :

```json
{
  "status": "Success",
  "code": 200,
  "error": "",
  "data": {
    "message": "String"
  }
}
```

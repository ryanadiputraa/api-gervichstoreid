# GERVICHSTORE.ID - API

Backend API for GERVICHSTORE.ID website.

## API SPEC

---

## Users

### Register

- Method : `POST`
- Endpoint : `/auth/v1/register`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
- body :

```json
{
  "fullname": "String",
  "email": "String",
  "username": "String",
  "password": "String",
  "role": "String"
}
```

- response :

```json
{
  "message": "Success",
  "code": 201,
  "error": "",
  "data": null
}
```

### Get User Data

- Method : `GET`
- Endpoint : `/api/v1/users`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
  - Authorization: `Bearer <access_token>`
- response :

```json
{
  "message": "Success",
  "code": 200,
  "error": "",
  "data": {
    "fullname": "String",
    "email": "String",
    "username": "String",
    "role": "String"
  }
}
```

---

## Products

### Get All Products

- Method : `GET`
- Endpoint : `/api/v1/products`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
- response :

```json
{
  "message": "Success",
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
    },
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
- Endpoint : `/api/v1/products/:product_id`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
- response :

```json
{
  "message": "Success",
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

### Add Product

- Method : `POST`
- Endpoint : `/api/v1/products`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
  - Authorization: `Bearer <access_token>`
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
  "message": "Success",
  "code": 201,
  "error": "",
  "data": "null"
}
```

### Update Product

- Method : `PUT`
- Endpoint : `/api/v1/products/:product_id`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
  - Authorization: `Bearer <access_token>`
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
  "message": "Success",
  "code": 200,
  "error": "",
  "data": "null"
}
```

### Delete Product

- Method : `DELETE`
- Endpoint : `/api/v1/products/:product_id`
- Header :
  - Content-Type : `application/json`
  - Accept : `application/json`
  - Authorization: `Bearer <access_token>`
- response :

```json
{
  "message": "Success",
  "code": 200,
  "error": "",
  "data": "null"
}
```

# Coding Challenge

## Requirement

Create a simple CRUD app built with Go and Gin. This let you update track your inventory, add new items, remove items, change items info and export to CSV. Your final application should be containerized.

### Endpoints:


| Endpoint | Description |
| --- | ----------- |
| POST localhost:8080/item  | Post with a json body containing an id, name, quantity and unit_price. Returns the item added with a 201 code. |
| GET localhost:8080/item   | Returns a json with the list of the items and a 200 code.|
| GET localhost:8080/item/1 | Returns a json with the item wanted, and a 200 code. |
| GET localhost:8080/item/1 | Returns a json with the item wanted, and a 200 code. |
| DELETE localhost:8080/item/1 | Delete an item by its ID. Returns 200 if item was removed.|
| PATCH localhost:8080/item/2 | Body needs to contain a json with name, quantity and unit_price. The info of this item will be changed. Returns 200.|
| GET localhost:8080/item/csv | Returns a CSV file.|
| | |

Your application should be packaged as docker container (and docker compose if you have multiple containers).

## Working Demo (Hosted in Linode)
http://cc.sreeramganesan.com:8080/docs/index.html

## Run this project locally
1. Clone this repository
2. Run ```cd coding-challenge```
3. Run ```docker-compose up -d --build```

## Tools/Technologies Used
|Language           | Go      |
|Web Server         | Gin     |
|Database           | MySQL   |
|ORM                | gorm    |
|API Documentation  | Swagger |

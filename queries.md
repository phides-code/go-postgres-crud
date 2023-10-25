curl -X POST -H "Content-Type: application/json" -d '
{
"Make": "Toyota",
"Model": "Corolla",
"Year": 1999
}' http://localhost:8080/api/addcar | jq .

curl -X POST -H "Content-Type: application/json" -d '
{

}' http://localhost:8080/api/getcars | jq .

curl -X PUT -H "Content-Type: application/json" -d '
{
"id": 1,
"year": 2017,
"model": "Sportage FWD"

}' http://localhost:8080/api/updatecar | jq .

curl -X DELETE http://localhost:8080/api/deletecar/1 | jq .

curl -X POST -H "Content-Type: application/json" -d '
{

}' http://192.168.86.140:8080/api/getcars | jq .

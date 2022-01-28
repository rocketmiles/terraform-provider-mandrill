curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"key":"some-api-key","domain":"hybridexampledomain.com"}' \
  localhost:8080/api/1.0/senders/check-domain | jq
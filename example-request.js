const response = await fetch(
  "https://brapi.dev/api/v2/stocks/quote?symbols=PETR4",
  {
    headers: {
      "Authorization": "Bearer <BRAPI_TOKEN>"
    }
  }
);

const data = await response.json();
console.log(data);
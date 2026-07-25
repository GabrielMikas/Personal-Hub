const response = await fetch(
  "https://brapi.dev/api/v2/stocks/quote?symbols=PETR4",
  {
    headers: {
      "Authorization": "Bearer eCbcYdNriKNnUa8q66955V"
    }
  }
);

const data = await response.json();
console.log(data);
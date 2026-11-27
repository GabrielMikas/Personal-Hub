// Run with: node --env-file=apps/personal-hub-back/.env example-request.js
const token = process.env.BRAPI_KEY;
if (!token) {
  throw new Error("BRAPI_KEY is not set (see apps/personal-hub-back/.env.example)");
}

const response = await fetch(
  "https://brapi.dev/api/v2/stocks/quote?symbols=PETR4",
  {
    headers: {
      "Authorization": `Bearer ${token}`
    }
  }
);

const data = await response.json();
console.log(data);

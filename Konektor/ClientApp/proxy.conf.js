const { env } = require('process');

const target = "https://localhost:7141"

console.log(target);
const PROXY_CONFIG = [
  {
    context: [
      "/weatherforecast",
      "/_configuration",
      "/.well-known",
      "/Identity",
      "/connect",
      "/ApplyDatabaseMigrations",
      "/_framework",
      "/Projects",
      "/Project",
      "/projectdetails"
   ],
    target: target,
    secure: false,
    headers: {
      Connection: 'Keep-Alive'
    }
  }
]

module.exports = PROXY_CONFIG;

console.log(module.exports);

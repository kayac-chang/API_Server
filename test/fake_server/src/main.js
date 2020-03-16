const restify = require("restify");

const account = require("./account");
const bet = require("./bet");
const endround = require("./endround");

function main() {
  const server = restify.createServer();

  server.use(restify.plugins.bodyParser());

  // === routers ===
  const prefix = "/api/v1/tgc";

  server.get(`${prefix}/player/check/:account`, account);
  server.post(`${prefix}/transaction/game/bet`, bet);
  server.post(`${prefix}/transaction/game/endround`, endround);

  // === start ===
  server.listen(3000, function() {
    console.log("%s listening at %s", server.name, server.url);
  });
}

main();

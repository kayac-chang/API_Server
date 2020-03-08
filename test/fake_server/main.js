const restify = require("restify");

const prefix = "/api/v1/tgc";

function respond(req, res, next) {
  const respond = {
    data: true,
    balance: {
      balance: 600270,
      currency: "CNY"
    },
    status: {
      code: "0",
      message: "Success",
      datetime: new Date().toISOString()
    }
  };

  res.send(respond);
  next();
}

var server = restify.createServer();
server.get(`${prefix}/player/check/:account`, respond);

server.listen(3000, function() {
  console.log("%s listening at %s", server.name, server.url);
});

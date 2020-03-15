const user = require("./user");

function bet(req, res, next) {
  const session = req.header("session");
  const org_token = req.header("organization_token");

  if (!org_token || !session) {
    const respond = {
      data: {},
      status: {
        code: "0",
        message: "Failed",
        datetime: new Date().toISOString()
      }
    };

    res.send(401, respond);
    return next();
  }

  console.log(req.body);

  const {
    //
    account,
    created_at,
    gamename,
    roundid,
    amount
  } = req.body;

  user.balance -= amount;

  const respond = {
    data: {
      balance: user.balance
    },
    status: {
      code: "0",
      message: "Success",
      datetime: new Date().toISOString()
    }
  };

  res.send(200, respond);
  return next();
}

module.exports = bet;

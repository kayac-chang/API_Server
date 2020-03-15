const user = require("./user");

function account(req, res, next) {
  const respond = {
    data: {
      status: false,
      balance: {
        balance: user.balance
      }
    },
    status: {
      code: "0",
      message: "Success",
      datetime: new Date().toISOString()
    }
  };

  res.send(200, respond);
  next();
}

module.exports = account;

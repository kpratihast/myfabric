var AWS = require("aws-sdk");

AWS.config.loadFromPath("./config.json");

ddb = new AWS.DynamoDB({ apiVersion: "2012-10-08" });

ddb.listTables({ Limit: 10 }, function(err, data) {
  if (err) {
    console.log("Error:", err.code);
  } else {
    console.log("Tables: ", data.TableNames);
  }
});

var params = {
  Key: { Aadhar: { S: "345" } },
  TableName: "Aadhar_Details"
};

ddb.getItem(params, function(err, data) {
  if (err) {
    console.log(err);
  } else {
    console.log(JSON.stringify(data, null, 2));
  }
});

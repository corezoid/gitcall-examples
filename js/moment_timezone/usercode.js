
// build command: npm install moment moment-timezone

const moment = require("moment");
require("moment-timezone");

module.exports = (data) => {
    var t = moment()
    
    data.utc_time = t.tz("UTC").format("HH:mm:ss");
    data.usa_time = t.tz("America/New_York").format("HH:mm:ss");
    data.ukraine_time = t.tz("Europe/Kiev").format("HH:mm:ss");

    return data;

};
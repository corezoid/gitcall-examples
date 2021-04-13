
// build command: npm install uuid

const { v4 : uuidv4 } = require('uuid');

module.exports = (data) => {
    data.uuid1 = uuidv4();

    return data;
};

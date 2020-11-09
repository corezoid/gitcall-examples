const { v4 : uuidv4 } = require('./uuid.js');

/**
 * You should add uuid library, set:
 * Repo URL: https://github.com/uuidjs/uuid.git
 * Commit: 3f44acd0e722e965c14af816e2f658361a6b15f9
 */

module.exports = (data) => {
    data.uuid1 = uuidv4();

    return data;
};

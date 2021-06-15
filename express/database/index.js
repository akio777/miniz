const bcrypt = require('bcrypt')
const config = require('../config.json')
const Pool = require('pg').Pool
const pool = new Pool(config.db)


module.exports = pool

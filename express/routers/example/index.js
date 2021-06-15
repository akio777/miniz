const express = require('express')
const router = express.Router()

router.get("/", async(req, res, next)=>{
    res.status(200).send()
    next()
})

module.exports = router;
// [
//     '_readableState', 'readable',         '_events',
//     '_eventsCount',   '_maxListeners',    'socket',
//     'connection',     'httpVersionMajor', 'httpVersionMinor',
//     'httpVersion',    'complete',         'headers',
//     'rawHeaders',     'trailers',         'rawTrailers',
//     'aborted',        'upgrade',          'url',
//     'method',         'statusCode',       'statusMessage',
//     'client',         '_consuming',       '_dumped',
//     'next',           'baseUrl',          'originalUrl',
//     '_parsedUrl',     'params',           'query',
//     'res',            'body',             'secret',
    // 'cookies',        'signedCookies',    'startTime',
//     "authorization", 'x-api-key'
// ]
// const apiKey = require('./config.json')['api-key']
const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');
// const pool = require('./database');


// const PASS_PATH = {
//     "/" : 'GET',
//     "/login" : 'GET',
//     "/setting_import" : 'GET',
//     '/connection' : 'POST'
// }

module.exports = (req, res, next) => {
    const socket = req.headers["socket-id"] || req.body["socket-id"]
    console.log(req.method + ' ' + req.path + ' from ' + (socket?socket:getIp(req)) + ' Initialized')
    req.startTime = Date.now()
    // // TODO FAKE UID 
    // let token_header = req.headers.authorization
    // if(!(PASS_PATH[req.url]===req.method)){
    //     if(!token_header){
    //         res.status(404).send({msg:"not found token"})
    //         // next()
    //     }
    //     else if(token_header.split(' ')[0]!=='Bearer'){
    //         res.status(404).send({msg:"invalid token type"})
    //     }
    //     else{
    //         let check_token = token_header.split(' ')
    //         let token = check_token[1]
    //         jwt.verify(token, config.secret, (err, decoded)=>{
    //             if(err){
    //                 res.status(401).send({msg:"token expired"})
    //             }
    //             else{
    //                 req.uid = decoded.uid
    //                 req.token = token
    //                 // console.log('-------------------')
    //                 // console.log(req.token)
    //                 // console.log('-------------------')
    //                 // console.log(req.uid)
    //             }
                
    //         })
    //     }
    // }
    next()
}
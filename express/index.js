express = require('express')
path = require('path')
cookieParser = require('cookie-parser')
bodyParser  = require('body-parser')
RateLimiter = require('async-ratelimiter')
Redis = require('ioredis')
http = require('http')
helmet = require('helmet')
log4js = require('log4js')
pretty = require('express-prettify')
cors = require('cors')
AdmZip = require('adm-zip')
dateFileName = require('./functions/dateFileName')

const { getClientIp } = require('request-ip')
getIp = getClientIp

mode = process.env.MODE

// const setting_import = require("./router/setting_import")
// const user_info = require("./router/user_info")
// const login = require("./router/login")
// const import_type = require("./router/import_type")
// const flow_record = require("./router/flow_record")
// const analytic_record = require("./router/analytic_record")
// const connection = require("./router/connection")
// const exercise = require("./router/exercise")
// const pool = require('./database')

const Example = require('./routers/example')

const log4js_config = {
    appenders: {
        logs: { type: "file", filename: 'logs/' + dateFileName() + '.log' },
        console: { type: 'console'}
    },
    categories: { 
        default: {
            appenders: [
                "logs",
                "console"
            ], 
            level: "info"
        } 
    }
}

if(mode !== 'production'){
    delete log4js_config.appenders['logs']
    delete log4js_config.categories.default.appenders[0]
}

log4js.configure(log4js_config)

logger = log4js.getLogger('logs');
log = require('./functions/log')

if(mode === 'production'){
    log('Starting Production Mode')
    rateLimiter = new RateLimiter({
        db: new Redis()
    })
} else {
    log('Starting Development Mode')
}

app = express()
server = http.createServer(app)

client = {
    array : [],
    number : 0
}


config = require('./config.json')

apiQuota = require('./apiQuota')
middleware = require('./middleware')
lastware = require('./lastware.js')

app.use(express.json())
app.use(express.urlencoded({extended: false}))
app.use(cookieParser())
app.use(bodyParser.json())
app.use(bodyParser.urlencoded({extended: true}))
app.use(config.helmet?helmet(config.helmet):helmet())
app.use(cors({
    origin: '*',
    exposedHeaders: ['Content-Disposition']
}))
app.use(pretty({query: 'pretty'}))

if(mode === 'production') app.use(apiQuota)
app.use(middleware)

server.listen(config.port, () => log(`Listening on port ${config.port}`))

END = (req) => {
    const socket = req.headers["socket-id"] || req.body["socket-id"]
    log(req.method + ' ' + req.path + ' from ' + (socket?socket:getIp(req)) + ' Completed. Execution Time: ' + (Date.now() - req.startTime) + ' ms')
}

//################################################################################### PRE ROUTING

app.use('/example', Example)

app.get('/test', (req, res, next)=>{
    res.status(200).send()
    next()
})

//################################################################################### POST ROUTING

app.use(lastware)

//################################################################################### SOCKET IO

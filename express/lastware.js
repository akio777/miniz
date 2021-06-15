module.exports = (req, res) => {
    const socket = req.headers["socket-id"] || req.body["socket-id"]
    log(req.method + ' ' + req.path + ' from ' + (socket?socket:getIp(req)) + ' Completed. Execution Time: ' + (Date.now() - req.startTime) + ' ms')

    if(!res.headersSent) res.status(404).send('NOT FOUND - ' + req.method + ' ' + req.path) 
}
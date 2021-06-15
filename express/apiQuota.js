module.exports = async (req, res, next) => {
    const ip = getIp(req)
    const limit = await rateLimiter.get({ id: ip })

    if(!res.finished && !res.headersSent) {
        res.setHeader('X-Rate-Limit-Limit', limit.total)
        res.setHeader('X-Rate-Limit-Remaining', Math.max(0, limit.remaining - 1))
        res.setHeader('X-Rate-Limit-Reset', limit.reset)
    }

    return !limit.remaining
        ? sendFail({
            req,
            res,
            code: HTTPStatus.TOO_MANY_REQUESTS,
            message: MESSAGES.RATE_LIMIT_EXCEDEED()
        })
        : next()
}
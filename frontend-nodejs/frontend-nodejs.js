const express = require('express')
const axios = require('axios').default

const PORT = 8080

const BACKEND_HOSTS = [
    process.env.BACKEND_HOST1 || 'backend-spring',
    process.env.BACKEND_HOST2 || 'backend-golang'
]

const GREETING_LABEL = process.env.greetingLabel || '#greetingLabel#'

const app = express()

app.get('*', async (req, res) => {
    console.log(req.method, req.url)
    const name = req.url.substr(1)
    try {
        const results = await proxyToBackends(BACKEND_HOSTS, name)
        const joined = results.reduce((acc, result) => `${acc}\n${result}`, `[NODEJS] ${GREETING_LABEL} ${name}`)
        res.send(joined)
    } catch (e) {
        console.error(e)
        res.status(400).send(e)
    }
})

app.listen(PORT, () => {
    console.log(`Server listens on port ${PORT}`)
})

async function proxyToBackends(hosts, name) {
    const urls = hosts.map(host => `http://${host}/${name}`)
    console.log(`Call target URLs ${urls}`)
    const promises = urls.map(url => axios.get(url))
    const responses = await Promise.all(promises)
    const results = responses.map(response => response.data)
    console.log(`Retrieved results ${results}`)
    return results
}

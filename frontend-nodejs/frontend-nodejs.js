const express = require('express')
const axios = require('axios').default

const PORT = 8080
const BACKEND_URL1 = process.env.BACKEND_URL1 || 'backend-golang'
const BACKEND_URL2 = process.env.BACKEND_URL2 || 'backend-spring'

const app = express()

async function queryBackend(name) {
    const target1 = `http://${BACKEND_URL1}/${name}`
    const target2 = `http://${BACKEND_URL2}/${name}`
    console.log(`Call target URLs ${target1} ${target2}`)
    const promises = [axios.get(target1), axios.get(target2)]
    const responses = await Promise.all(promises)
    return responses.reduce((acc, result) => acc + result.data + '\n', '')
}

async function startupProbe() {
    console.log(`Backend URLs: ${BACKEND_URL1} ${BACKEND_URL2}, perform connection test ...`)
    try {
        const result = await queryBackend('Probe')
        console.log(`Backend(s) answered with\n${result}\nProbe ok`)
    } catch (e) {
        console.warn(`Backend(s) returned error ${e.message}, continue anyway`)
    }
}

app.get('*', async (req, res) => {
    console.log(req.method, req.url)
    const name = req.url.substr(1)
    try {
        const result = await queryBackend(name)
        res.send(`${result}[NODEJS] Hello ${name}`)
    } catch (e) {
        console.error(e)
        res.status(400).send(e)
    }
})

app.listen(PORT, () => {
    console.log(`Server listens on port ${PORT}`)
    startupProbe()
})

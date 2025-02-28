import express from 'express'
import axios from 'axios'

const PORT = 8080

const BACKEND_HOSTS = (process.env.BACKEND_HOSTS || "").split(',')
const GREETING_LABEL = process.env.greetingLabel || '#greetingLabel#'

const app = express()

app.get('*', async (req, res) => {
    console.log(req.method, req.url)
    const name = req.url.substr(1)
    try {
        const results = await Promise.all(BACKEND_HOSTS.map(host => proxyToBackend(host, name)))
        const joined = results.reduce((acc, result) => `${acc}\n${result}`, `[NODEJS] ${GREETING_LABEL} ${name}`)
        res.send(joined)
    } catch (e) {
        res.status(500).send(e.message)
    }
})

app.listen(PORT, () => {
    console.log(`Server listens on port ${PORT}`)
})

async function proxyToBackend(host: string, name: string) {
    const url = `http://${host}/${name}`
    console.log(`Call target URL ${url}`)
    try {
        const response = await axios.get(url)
        console.log(`Retrieved result ${response.data} from ${url}`)
        return response.data
    } catch (e) {
        console.warn(e.message)
        throw e
    }
}

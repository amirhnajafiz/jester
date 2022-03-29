import { CreateClient } from "./client";

const express = require('express')
const app = express()
const port = 3000
const cli = CreateClient()

app.get('/sub', (req, res) => {
    cli.sub()

    res.send('Hello World!')
})

app.get('/pub', (req, res) => {
    cli.pub()

    res.send('Temp')
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})

// client creation method
import { CreateClient } from "./client";

// setup express app
const express = require('express')
const app = express()
const port = 5000

// our client
const cli = CreateClient()

// routes
app.get('/check', (req, res) => {
    const topic = req.query.get("topic")

    res.send({"topic": topic})
})

// starting the application
app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})

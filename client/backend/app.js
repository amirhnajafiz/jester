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
    const topic = "test1/stan";
    const send = {
        topic: topic,
        type: "normal"
    };
    let tests = [
        {
            id: "1001",
            topic: topic,
            content: this.id,
            status: false
        },
        {
            id: "1002",
            topic: topic,
            content: this.id,
            status: false
        },
        {
            id: "1003",
            topic: topic,
            content: this.id,
            status: false
        }
    ]

    cli.sub(send, function (err, resp) {
        if (err)
            return null

        tests.forEach((test) => {
            if (test.id === resp.content) {
                test.status = true
            }
        })
    })

    tests.forEach((test) => {
        cli.pub(send, function (err, resp) {
            if (err)
                console.error(err)
            console.log(resp)
        })
    })

    res.send({"tests": tests})
})

// starting the application
app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})

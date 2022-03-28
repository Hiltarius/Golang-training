const express = require('express')

const app = express();

app.use('/', (req,res) => {
    req.pipe(
        request(
            'http://localhost:4001'+ req.url,
            (error) => {
                if ( error && error.code === 'ECONNREFUSED')
                    console.error('connexion refused')
            }
        )
    ).pipe(res)
})

app.listen('7069')
const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
  res.send('Hello World V9!')
})

app.get('/info', (req, res) => {
  res.send('Another info!')
})

app.listen(port, () => {
  console.log(`Listening on port ${port}`)
})

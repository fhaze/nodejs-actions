const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
  res.send('Hello World V1.0.1!')
})

app.listen(port, () => {
  console.log(`Listening on port ${port}`)
})

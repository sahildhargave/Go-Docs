

const express = require('express')
const app = express()
const port = 3000

app.use(express.json());
app.use(express.urlencoded({extended: true}))

app.get('/',(req, res) => {
	res.status(200).send("Welcome to Google store");
})

app.get('/get', (req, res) => {
	res.status(200).json({ message: "Hello from Google support team"})
})

app.post('/post',(req,res) => {
	let myJson = req.body;

	res.status(200).send(myJson);
})


app.post('/postform',(req,res)=>{
	res.status(2000)
})
app.get('/', (req, res) => res.send('Hello World!'))
app.listen(port, () => console.log(`Example app listening on port ${port}!`))
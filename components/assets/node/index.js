require('dotenv').config()

const path = require('path');

const express = require('express');

const bodyParser = require('body-parser');

const cookieParser = require('cookie-parser');

const { body, validationResult } = require('express-validator');

const moment = require('moment');


const app = express();

app.use(bodyParser.urlencoded({ extended: true }));

app.use(cookieParser());




app.get("/", (req, res) => {

    res.send({ message: 'this works' });
});

app.listen(80, () => {

    console.log("Server is running at localhost");

});


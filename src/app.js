const express = require('express');
require('dotenv').config()
const { engine } = require('express-handlebars');
const app = express();
const path = require('path');
const mysql = require('mysql2');
var helpers = require('handlebars-helpers')();

const connection = mysql.createConnection({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PWD,
    database: process.env.DB_NAME,
    port: process.env.DB_PORT,
});

const pool = mysql.createPool({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PWD,
    database: process.env.DB_NAME,
    port: process.env.DB_PORT,
    connectionLimit: 10,
});

connection.connect((err) => {
    if (err) {
        console.error('error connecting db :', err);
        return;
    }
});

app.engine('handlebars', engine({
    partialsDir: `${__dirname}/views/partials`,
}));

app.set('view engine', 'handlebars');
app.set('views', `${__dirname}/views`);

app.use(express.static(path.join(__dirname, '/public/')))

app.get('/', (req, res) => {
    pool.query('SELECT * FROM articles', (err, results) => {
        if (err) {
            console.error('not found :', err);
            return;
        }
        res.render('home', {
            title: 'My App',
            articles: results,
        });
    });
});

app.listen(process.env.PORT, () => {
    console.log(`server launch on port ${process.env.PORT}`);
});
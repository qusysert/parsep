const express = require('express');
const bodyParser = require('body-parser');
const puppeteer = require('puppeteer');
require('dotenv').config();

const app = express();
app.use(bodyParser.json());

const width = 500;
const height = 300;
const host = process.env.HOST || 'localhost';
const port = process.env.PORT || 3000;

app.post('/render', async (req, res) => {
    const html = req.body.query;

    try {
        const browser = await puppeteer.launch();
        const page = await browser.newPage({args: ['--disable-antialiasing']});

        await page.setContent(html);

        await page.setViewport({
            width: width,
            height: height,
            deviceScaleFactor: 5 // set device scale factor to 2x
        });

        const imageBuffer = await page.screenshot({
            clip: { x: 0, y: 0, width, height },
            type: 'jpeg',
            quality: 100 // set JPEG compression quality to 90%
        });

        await browser.close();

        res.writeHead(200, {
            'Content-Type': 'image/jpeg',
            'Content-Length': imageBuffer.length
        });

        res.end(imageBuffer);
    } catch (error) {
        console.error(error);
        res.status(500).send('Internal Server Error');
    }
});

app.listen(port, host,() => {
    console.log('Server started on port 3000');
});

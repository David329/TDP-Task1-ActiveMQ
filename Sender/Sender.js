"use strict";
const stompit = require('stompit');
const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

stompit.connect({ host: 'localhost', port: 61613 }, (err, client) => {
    const frame = client.send({ destination: 'SampleQueue' });

    rl.question('Consola Javascript: ', (answer) => {
        if (answer == "salir") return;
        frame.write(answer);
        frame.end();
        console.log('Mensaje enviado');

        rl.close();
    });

    client.disconnect();
});
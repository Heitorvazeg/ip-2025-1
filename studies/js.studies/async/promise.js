function minhapromise() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            let sucesso = Math.random();

            if (sucesso > 0.5) {
                resolve("sucesso")
            }
            else {
                reject("Nada de sucesso")
            }
        }, 3000);
    })
}
minhapromise().then(resultado => console.log("Sucesso: ", resultado)).catch(erro => console.log("erro: ", erro));

function passearComCachorro() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            if (true) {
                resolve("Você passeou com o cachorro");
            } else {
                reject("Você não passeou com o cachorro");
            }
        }, 3000);
    });
}

function lavarACasa() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            if (true) {
                resolve("Você lavou a casa");
            } else {
                reject("Você não lavou a casa");
            }
        }, 2000);
    });
}

function estudar() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            if (true) {
                resolve("Você estudou");
            } else {
                reject("Você não estudou");
            }
        }, 5000);
    });
}

passearComCachorro()
    .then(resultado => console.log(resultado))
    .then(lavarACasa)
    .then(resultado_2 => console.log(resultado_2))
    .then(estudar)
    .then(resultado_3 => console.log(resultado_3))
    .then(() => console.log("Tudo pronto"))
    .catch(erro => console.error("Erro:", erro));

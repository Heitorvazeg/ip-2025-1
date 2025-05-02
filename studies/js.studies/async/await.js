async function passearComCachorro() {
    return await new Promise((resolve, reject) => setTimeout(() => {
        let real = true;
        if (real) {
            resolve("Voce passeou com o cachorro")
        }
        else {
            reject("Voce nao passou com o cachorro")
        }
    }, 3000));
}

async function lavarACasa() {
    return await new Promise(resolve => setTimeout(() => {
        resolve("Voce lavou a casa");
    }, 1000));
}

async function estudar() {
    await new Promise(resolve => setTimeout(resolve, 3000));
    return "Você estudou";
}

async function fazerTarefas() {
    const passearCmCachorro = await passearComCachorro();
    console.log(passearCmCachorro);

    const lavarCasa = await lavarACasa();
    console.log(lavarCasa);

    const estudos = await estudar();
    console.log(estudos);
}

fazerTarefas().catch(error => console.log(error));

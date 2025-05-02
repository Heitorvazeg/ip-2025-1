async function fetchData() {
    try {
        const input = document.getElementById("myInput").value.toLowerCase();
        const response = await fetch(`https://pokeapi.co/api/v2/pokemon/${input}`);
        if (!response.ok) {
            throw new Error("Erro de fetch");
        }
        const data = await response.json();
        const image = document.getElementById("myImage");
        const pokemon = data.sprites.front_default;
        
        image.src = pokemon;
        image.style.display = "block";
    }
    catch (error) {
        console.error(error);
    }
}
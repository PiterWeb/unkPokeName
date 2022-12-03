<script lang="ts">
  import { onMount } from "svelte";
  import { GetPokemon } from "../wailsjs/go/main/App";

  let pokemonName: string = "";
  let imageUrl: string = "";
  let guessedName: string = "";
  let correct: boolean = false;
  let show: boolean = false;

  function selectPoke() {
    show = false;
    guessedName = "";

    GetPokemon().then((result: string[]) => {
      pokemonName = result[0];
      imageUrl = result[1];
    });
  }

  onMount(() => {
    selectPoke();
  });

  $: correct = pokemonName.toLowerCase() === guessedName.toLowerCase();
</script>

<main>
  <h1>UnkPokeName</h1>

  <input
    disabled={correct || show}
    type="text"
    bind:value={guessedName}
    placeholder="Write the name of the Pokemon"
  />

  {#if correct}
    <h2>Well done, you are right</h2>
  {/if}

  <img src={imageUrl} alt="pokemon sprite" />

  <button
    id="show"
    on:click={() => {
      show = true;
    }}
    disabled={correct || show}
  >
    Show Answer
  </button>

  <button id="next" on:click={selectPoke}>Next Pokemon</button>

  {#if show}
    <p>Current Pokemon: {pokemonName}</p>
  {/if}
</main>

<style>
  @import url("https://fonts.googleapis.com/css2?family=VT323&display=swap");

  main {
    display: flex;
    flex-direction: column;
    align-items: center;
    font-family: "VT323", monospace;
  }

  img {
    width: 200px;
    height: 200px;
    padding: 2rem;
    margin-top: 2rem;
    border-radius: 2rem;
    background-color: black;
    filter: drop-shadow(0 0 0.5rem rgb(122, 19, 24));
  }

  #next {
    margin-top: 2rem;
    padding: 1rem;
    border-radius: 2rem;
    background-color: black;
    color: white;
    font-size: 1.5rem;
  }

  #show {
    margin-top: 1rem;
    padding: 0.5rem;
    border-radius: 2rem;
    background-color: black;
    color: white;
    font-size: 1rem;
  }

  #show:disabled {
    background-color: grey;
  }

  p {
    margin-top: 1rem;
    font-size: 1.5rem;
  }

  input {
    width: 20rem;
    padding: 1rem;
    border-radius: 2rem;
    border: none;
    font-size: 1.25rem;
    text-align: center;
  }

  input::placeholder {
    color: grey;
  }

  input:focus-visible {
    outline: 1rem solid rgb(185, 34, 42);
  }

  input:disabled {
    background-color: grey;
    color: whitesmoke;
  }

</style>

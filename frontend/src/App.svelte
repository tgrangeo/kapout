<script lang="ts">
  import Button from './lib/Button.svelte';
  import QuizCard from './lib/QuizCard.svelte';

  let code = "";

interface Choice {
  id: string;
  name: string;
  correct: boolean;
}

interface QuizQuestion {
  id: string;
  name: string;
  choices: Choice[];
}

interface Quiz {
  id: string;
  Name: string;
  questions: QuizQuestion[];
}

  let quizzes: Quiz[];
  let msg = ""

  async function getQuizzes(){
    let response = await fetch("http://localhost:3000/api/quizzes")
    if (!response.ok){
      alert("failed!")
      return
    }
    let json = await response.json()
    quizzes = json
    console.log(quizzes)
  }

  function connect(){
    let websocket = new WebSocket("ws://localhost:3000/ws")
    websocket.onopen = () =>{
      console.log("opened connection")
      websocket.send(`join:${code}`)
    }
    websocket.onmessage = (event) =>{
      console.log(event.data)
    }

  }

  function hostQuizz(quiz:any){
    let websocket = new WebSocket("ws://localhost:3000/ws")
    websocket.onopen = () =>{
      console.log("opened question")
      websocket.send(`host:${quiz.id}`)
    }
    websocket.onmessage = (event) =>{
      msg = event.data
    }
  }

</script>


  <Button on:click={getQuizzes}>Get quizzes</Button>
  
  {#each quizzes as quiz}
  <QuizCard quiz={quiz.Name} host={hostQuizz(quiz)}/>
  {/each}

  <input bind:value={code} type="text" placeholder="Game Code" class="boder"/>
  <Button on:click={connect}>Join Game</Button>
Message:{msg}



<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .read-the-docs {
    color: #888;
  }
</style>

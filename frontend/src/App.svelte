<script lang="ts">
  import Button from './lib/Button.svelte';
  import QuizCard from './lib/QuizCard.svelte';
  import type { Quiz } from './model/quiz';
  import { NetService } from './service/net';

  let code = "";
  let quizzes: Quiz[];
  let msg = ""
  let netService =  new NetService() 

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

  function hostQuizz(quiz:any){
    netService.sendPacket({
      id:1,
      quiId: quiz.id
    })
  }

  function connect(){
    netService.sendPacket({
      id:1,
      code:"1234",
      name:"bob"
    })
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

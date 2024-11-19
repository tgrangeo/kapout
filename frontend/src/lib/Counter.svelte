<script lang="ts">
  let count: number = $state(69)
  const increment = () => {
    count -= 1
  }
</script>

<button onclick={increment}>
  count is {count} remaining
</button>

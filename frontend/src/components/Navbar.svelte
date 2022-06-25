<script lang="ts">
  let isAuthenticated: boolean = false;

  const auth = () => {
    isAuthenticated = !isAuthenticated;
  };
</script>

<nav class="container flex flex-row justify-between items-center py-6">
  <div class="flex flex-row items space-x-2">
    <div class="w-7 -translate-y-1.5">
      <img src="/kmutt-logo-normal.svg" alt="" />
    </div>
    <div class="font-semibold text-lg">kmutt.in.th</div>
  </div>

  <div class="flex flex-row items-center space-x-4">
    {#if isAuthenticated}
      <span>krid.heprakhone</span>
      <button
        type="button"
        on:click={auth}
        class="btn-primary"
      >
        Logout
      </button>
    {:else}
      <button
        type="button"
        on:click={auth}
        class="btn-primary"
      >
        Login
      </button>
    {/if}
  </div>
</nav>

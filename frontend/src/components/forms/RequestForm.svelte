<script lang="ts">
  let subDomain: string = '';
  let isAvailable: boolean = false;

  const mockData = [
    'comcamp',
  ];

  const changeSubDomain = (event) => {
    const input: string = event.target.value;

    if (/[A-Za-z]/.test(input)) {
      subDomain = input.toLowerCase();
      isAvailable = !mockData.includes(subDomain);
    }
  };
</script>

<form>
  <div class="flex flex-row justify-center">
    <input
      type="text"
      on:input={changeSubDomain}
      bind:value={subDomain}
      spellcheck="false"
      autocomplete="off"
      class="px-4 border border-orange-600 text-orange-600 outline-none"
    />
    <button
      type="text"
      class="px-4 py-2 bg-orange-600 text-white"
    >
      Request
    </button>
  </div>

  <div
    class="mt-4"
    class:text-red-500="{!isAvailable}"
    class:invisible="{!subDomain}"
  >
    {subDomain}.kmutt.in.th is {!isAvailable ? 'not' : ''} available
  </div>

  <div class="flex flex-col mt-4">
    <div class="flex space-x-2 items-center">
      <input type="checkbox" id="error-1" class="accent-red-500" checked />
      <label for="error-1" class="text-red-500">
        Subdomain already in-use
      </label>
    </div>
    <!-- <div class="flex space-x-2 items-center">
      <input type="checkbox" id="error-1" class="accent-red-500" checked />
      <label for="error-1" class="text-red-500">
        Should be only a-z
      </label>
    </div> -->
  </div>
</form>

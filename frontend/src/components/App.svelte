<script lang="ts">
  import RequestForm from '@/components/forms/RequestForm.svelte';
  import Navbar from '@/components/Navbar.svelte';
  import Footer from '@/components/Footer.svelte';
</script>

<main class="flex flex-col h-screen">
  <Navbar />

  <div class="flex flex-col justify-center items-center w-full h-full pb-12">
    <div class="mb-6 text-2xl select-none">
      Search available
      <span class="font-semibold text-orange-600">Subdomain</span>
    </div>
    
    <RequestForm />
  </div>

  <Footer />
</main>

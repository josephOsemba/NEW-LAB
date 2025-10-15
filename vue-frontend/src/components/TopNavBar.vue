<template>
  <nav class="bg-blue-600 text-white p-4 flex justify-between items-center shadow-md">
    <h1 class="text-2xl font-bold tracking-wide">ARLab-PAS</h1>
    <ul class="flex gap-6 text-white">
      <li>
        <a href="/" class="hover:text-yellow-600 not-only:transition-all duration-300">Home</a>
      </li>
      <li>
        <a href="/askAI" class="hover:text-yellow-600 transition-all duration-300">Ask AI</a>
      </li>
      <li>
        <a href="/mainLab" class="hover:text-yellow-600 transition-all duration-300">Main 3D Lab</a>
      </li>
      <li>
        <a href="/about" class="hover:text-yellow-600 transition-all duration-300">About</a>
      </li>
      <li>
        <a href="/contact" class="hover:text-yellow-600 transition-all duration-300">Contact</a>
      </li>
    </ul>
  </nav>
</template>

<script>
export default {
  name: "TopNavBar",
};
</script>

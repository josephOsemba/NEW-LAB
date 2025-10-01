<template>
  <div class="procedure-section py-8 px-4 sm:px-6 md:px-12 bg-gray-50">
    <div class="content-wrapper max-w-5xl mx-auto space-y-6">
      <h2 class="text-2xl md:text-3xl font-bold text-center">Procedure of Pendulum Gravity</h2>

      <div>
        <h3 class="font-semibold text-lg">Procedure</h3>
        <p>The measurement of the acceleration due to gravity:</p>
        <p>
          Measure the period for different lengths <strong>l</strong> and plot a graph of <strong>l</strong> vs <strong>T²</strong>.
        </p>
        <p>The gradient gives <strong>g/4π²</strong>, so <strong>g = 4π²l/T²</strong>.</p>
      </div>

      <div>
        <h4 class="font-semibold">Steps:</h4>
        <ol class="list-decimal ml-6 space-y-2">
          <li>Set length l = 20cm, displace pendulum bob slightly, record time for 20 oscillations.</li>
          <li>Increase length in 20cm steps for 5 different lengths, plot l vs T².</li>
          <li>Calculate slope and determine g.</li>
          <li>Measure 30 oscillations for amplitudes 10° to 70°, plot T1 vs θ.</li>
        </ol>
      </div>

      <div>
        <h4 class="font-semibold">Procedure Diagram</h4>
        <div class="image-preview-container flex justify-center">
          <img src="/icons/GRAVITY-THEORY.png" alt="Procedure Visual" class="w-full max-w-3xl h-auto rounded shadow-md"/>
        </div>
      </div>

      <div>
        <h4 class="font-semibold mt-6">Tables</h4>
        <table class="w-full border border-gray-300 rounded-md table-auto text-sm md:text-base">
          <thead class="bg-gray-200">
            <tr>
              <th class="border px-2 py-1">Length (cm)</th>
              <th class="border px-2 py-1">Time for 20 oscillations (s)</th>
              <th class="border px-2 py-1">Time period (s)</th>
              <th class="border px-2 py-1">T² (s²)</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="i in 4" :key="i">
              <td class="border px-2 py-1">&nbsp;</td>
              <td class="border px-2 py-1">&nbsp;</td>
              <td class="border px-2 py-1">&nbsp;</td>
              <td class="border px-2 py-1">&nbsp;</td>
            </tr>
          </tbody>
        </table>
      </div>

    </div>
  </div>
</template>

<script setup>
</script>

<style scoped>
</style>

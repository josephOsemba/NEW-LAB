<template>
  <div class="theory-section py-10 px-4 sm:px-6 md:px-12 bg-gray-50">
    <div class="content-wrapper max-w-6xl mx-auto">
      <div class="text-content space-y-8">

        <!-- Page Title -->
        <h2 class="text-3xl sm:text-4xl font-bold text-center text-gray-800">
          Theory of Pendulum Gravity
        </h2>

        <!-- Aim -->
        <section class="space-y-2">
          <h3 class="font-semibold text-xl text-gray-700">Aim</h3>
          <p class="text-gray-600">
            This is an experiment to determine the magnitude of the acceleration due to gravity by using a simple pendulum.
          </p>
        </section>

        <!-- Apparatus -->
        <section class="space-y-2">
          <h3 class="font-semibold text-xl text-gray-700">Apparatus</h3>
          <ul class="list-disc list-inside space-y-1 text-gray-600">
            <li>A small mass (bob)</li>
            <li>2m long string</li>
            <li>Stopwatch/clock</li>
            <li>Meter rule</li>
            <li>Stand</li>
            <li>Protractor</li>
          </ul>
        </section>

        <!-- Theory -->
        <section class="space-y-3">
          <h3 class="font-semibold text-xl text-gray-700">Theory</h3>
          <p class="text-gray-600">
            Consider a pendulum of length <strong>l</strong> with a mass <strong>m</strong> at the end displaced through an angle <strong>θ</strong> from the vertical.
            The restoring force <strong>F</strong> is the component of the weight of the bob.
          </p>
          <pre class="bg-gray-100 p-4 rounded-md overflow-x-auto"><code>F = -mg sin(θ) = ma</code></pre>
          <p class="text-gray-600">
            Giving <strong>a = -g sin(θ)</strong>. For small angles, <strong>sin(θ) ≈ θ</strong>:
          </p>
          <pre class="bg-gray-100 p-4 rounded-md overflow-x-auto"><code>a = -gθ = -g(x/l)</code></pre>
          <p class="text-gray-600">
            The acceleration is proportional to the negative of the displacement and the pendulum exhibits simple harmonic motion.
          </p>
          <pre class="bg-gray-100 p-4 rounded-md overflow-x-auto"><code>T = 2π√(l/g)</code></pre>
          <p class="text-gray-600">
            Accurate for small angles of swing. For larger angles, use the corrected formula to account for the amplitude.
          </p>
        </section>

        <!-- Theory Diagram -->
        <section class="space-y-2">
          <h4 class="font-semibold text-lg text-gray-700">Theory Diagram</h4>
          <div class="image-preview-container flex justify-center">
            <img
              src="/icons/GRAVITY-PROCEDURE.png"
              alt="Theory and Procedure"
              class="w-full max-w-4xl h-auto rounded-lg shadow-md"
            />
          </div>
        </section>

      </div>
    </div>
  </div>
</template>

<script setup>
</script>

<style scoped>
/* Optional: smooth image hover effect */
.image-preview-container img:hover {
  transform: scale(1.02);
  transition: transform 0.3s ease-in-out;
}
</style>

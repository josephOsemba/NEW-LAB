<template>
  <div class="flex-1 p-6 md:p-12 bg-gray-100 min-h-screen">
    <h2 class="text-3xl md:text-4xl font-semibold text-center mb-10">
      Virtual Labs at ARLab-PAS
    </h2>

    <div
      class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6 md:gap-8"
    >
      <SubjectCard
        v-for="card in cardData"
        :key="card.title"
        :title="card.title"
        :description="card.description"
        :icon="card.image"
        :path="card.path"
      />
    </div>
  </div>
</template>

<script setup>
import SubjectCard from "@/components/SubjectCard.vue";

const cardData = [
  {
    title: "PHYSICS",
    image: "/icons/health.jpg",
    description:
      "Explore simulations in mechanics, optics, and electricity with interactive AR tools.",
    path: "/physics",
  },
  {
    title: "BIOLOGY",
    image: "/icons/biology.png",
    description:
      "Visualize molecular structures, reactions, and chemical processes in a virtual lab environment.",
    path: "/biology",
  },
  {
    title: "CHEMISTRY",
    image: "/icons/circuit-simulator.jpg",
    description:
      "Dive into cell biology, genetics, and anatomy through interactive 3D simulations.",
    path: "/chemistry",
  },
  {
    title: "ENGINEERING",
    image: "/icons/eng.png",
    description:
      "Access engineering design tools and circuit simulations tailored for practical learning.",
    path: "/engineering",
  },
  {
    title: "EARTH SCIENCE",
    image: "/icons/earth-science.jpg",
    description:
      "Earth Science encompasses a wide range of scientific disciplines focused on understanding our planet.",
    path: "/earth-science",
  },
  {
    title: "MATHEMATICAL SCIENCE",
    image: "/icons/maths.png",
    description:
      "Mathematical sciences apply models and methods to solve real-world problems in various fields.",
    path: "/maths",
  },
  {
    title: "HEALTH EDUCATION",
    image: "/icons/mechanics-lab.jpg",
    description:
      "Engage in virtual health training, anatomy modules, and medical simulations.",
    path: "/health-education",
  },
  {
    title: "COMPUTERS",
    image: "/icons/ictt.jpg",
    description:
      "Learn how technology drives innovation with ICT simulations and coding labs.",
    path: "/computers",
  },
  {
    title: "HOME SCIENCE",
    image: "/icons/hs.jpg",
    description:
      "Learn about nutrition, textiles, and home management through immersive lab setups.",
    path: "/home-science",
  },
];
</script>

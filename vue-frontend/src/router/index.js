import { createRouter, createWebHistory } from 'vue-router'

import HomeView from '@/views/HomeView.vue'
import PhysicsPracticals from '@/views/pages/PhysicsPracticals.vue'
import ProjectileMotion from '@/views/practicals/ProjectileMotion.vue'
import NewtonsSecondLaw from '@/views/practicals/NewtonsSecondLaw.vue'
import MachineEfficiency from '@/views/practicals/MachineEfficiency.vue'

import PendulumLayout from '@/views/practicals/PendulumGravity.vue'
import PendulumTheory from '@/views/practicals/pendulum/PendulumTheory.vue'
import PendulumProcedure from '@/views/practicals/pendulum/PendulumProcedure.vue'
import PendulumEvaluation from '@/views/practicals/pendulum/PendulumEvaluation.vue'
import PendulumSimulator from '@/views/practicals/pendulum/PendulumSimulator.vue'
import PendulumAssignment from '@/views/practicals/pendulum/PendulumAssignment.vue'
import PendulumReference from '@/views/practicals/pendulum/PendulumReference.vue'
import PendulumFeedback from '@/views/practicals/pendulum/PendulumFeedback.vue'
import AskAI from '@/views/pages/AskAI.vue'
import MainLab from '@/views/pages/MainLab.vue'

const routes = [
  { path: '/', name: 'home', component: HomeView },
  {path: '/askAI', name: 'askAI', component: AskAI},
  {path: '/mainLab', name: 'mainLab', component: MainLab},
  { path: '/physics', name: 'physics', component: PhysicsPracticals },
  { path: '/projectile', name: 'projectile', component: ProjectileMotion },
  { path: '/newtons-law', name: 'newtonsLaw', component: NewtonsSecondLaw },
  { path: '/machine-efficiency', name: 'machineEfficiency', component: MachineEfficiency },

  {
    path: '/pendulum',
    component: PendulumLayout,
    children: [
      { path: '', redirect: { name: 'PendulumTheory' } },

      { path: 'theory', name: 'PendulumTheory', component: PendulumTheory },
      { path: 'procedure', name: 'PendulumProcedure', component: PendulumProcedure },
      { path: 'evaluation', name: 'PendulumEvaluation', component: PendulumEvaluation },
      { path: 'simulator', name: 'PendulumSimulator', component: PendulumSimulator },
      { path: 'assignment', name: 'PendulumAssignment', component: PendulumAssignment },
      { path: 'reference', name: 'PendulumReference', component: PendulumReference },
      { path: 'feedback', name: 'PendulumFeedback', component: PendulumFeedback },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router

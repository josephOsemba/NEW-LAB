<template>
  <div class="home-container">
    <!-- Three.js Background - Covering entire background -->
    <div
      ref="mountRef"
      class="canvas-container position-fixed top-0 start-0 w-100 h-100"
      :style="{
        zIndex: 0,
        left: isSidebarCollapsed ? '60px' : '250px',
        width: isSidebarCollapsed ? 'calc(100% - 60px)' : 'calc(100% - 250px)',
        transition: 'all 0.3s ease'
      }"
    >
      <div v-if="isLoading" class="position-absolute top-50 start-50 translate-middle text-white z-1 loading-text">
        Loading Particles...
      </div>
    </div>

    <!-- Sidebar -->
    <SideBar
      :search-history="searchHistory"
      @ask-ai="handleAskAI"
      @history-item-click="handleHistoryItemClick"
      @collapse="handleSidebarCollapse"
    />

    <!-- Main Content -->
    <div
      class="main-content position-relative z-1 d-flex flex-column align-items-center justify-content-center py-5"
      :style="{ marginLeft: isSidebarCollapsed ? '60px' : '250px', transition: 'margin-left 0.3s ease' }"
    >
      <!-- Centered Cards Section -->
      <div class="centered-content text-center">
        <!-- Welcome Text -->
        <div class="welcome-text mb-5">
          <h1 class="display-4 fw-bold text-white mb-3">Welcome to ARLab</h1>
          <p class="lead text-light">Explore the future of augmented reality and AI-powered research</p>
        </div>

        <!-- Cards -->
        <div class="row justify-content-center">
          <div v-for="(title, index) in cards" :key="index" class="col-lg-4 col-md-6 mb-4">
            <div class="card bg-dark text-white shadow-lg hover-card">
              <div class="card-body">
                <h3 class="fw-bold card-title">{{ title }}</h3>
                <p class="text-muted card-description">{{ getCardDescription(title) }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Chat Interface - Improved Positioning and Design -->
      <div v-if="isChatOpen" class="ai-chat-modal">
        <div class="chat-container">
          <!-- Header with Solid Blue Background -->
          <div class="chat-header">
            <div class="header-content">
              <div class="header-icon">
                <i class="fas fa-robot"></i>
              </div>
              <div class="header-text">
                <h5 class="mb-0">ARLab AI Assistant</h5>
                <small class="status-text">Online</small>
              </div>
            </div>
            <button class="close-btn" @click="handleCloseChat">
              <i class="fas fa-times"></i>
            </button>
          </div>

          <!-- Chat Messages -->
          <div ref="chatRef" class="chat-messages">
            <div
              v-for="(msg, index) in messages"
              :key="index"
              :class="`message ${msg.sender}`"
            >
              <div class="message-content">
                <div class="message-text">{{ msg.text }}</div>
                <div class="message-time">{{ formatTime(msg.timestamp) }}</div>
              </div>
            </div>
            <div v-if="isTyping" class="message ai typing">
              <div class="message-content">
                <div class="typing-indicator">
                  <span class="dot"></span>
                  <span class="dot"></span>
                  <span class="dot"></span>
                </div>
              </div>
            </div>
          </div>

          <!-- Chat Input -->
          <div class="chat-input-container">
            <div class="input-group">
              <input
                type="text"
                v-model="input"
                @keypress="handleKeyPress"
                placeholder="Ask me anything about physics..."
                class="chat-input"
              />
              <button @click="handleSend" class="send-btn" :disabled="!input.trim()">
                <i class="fas fa-paper-plane"></i>
              </button>
            </div>
            <div class="input-hint">
              <small>Press Enter to send</small>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive, onMounted, onUnmounted, watch, nextTick } from 'vue'
import * as THREE from 'three'
import SideBar from '@/components/SideBar.vue'
import OpenAI from 'openai'

export default {
  name: 'AskAI',
  components: {
    SideBar
  },
  setup() {
    // Refs
    const mountRef = ref(null)
    const chatRef = ref(null)
    const isLoading = ref(true)
    const error = ref(null)
    const rendererRef = ref(null)
    const animationFrameId = ref(null)
    const isChatOpen = ref(false)
    const isSidebarCollapsed = ref(false)
    const isTyping = ref(false)

    // Reactive data
    const messages = reactive([])
    const input = ref('')
    const searchHistory = ref(JSON.parse(localStorage.getItem('searchHistory')) || [])

    // Card data
    const cards = ref(['ARLab', 'Research', 'Innovation'])

    // Initialize OpenAI client
    const openai = new OpenAI({
      baseURL: "https://api.deepseek.com/v1",
      apiKey: "sk-71ee17ccea4e4f79a6fc7ed3f4ac4bea",
      dangerouslyAllowBrowser: true,
    })

    // Methods
    const getCardDescription = (title) => {
      switch (title) {
        case 'ARLab':
          return 'This is an AI-driven Augmented Reality Lab that allows students to perform virtual scientific experiments.'
        case 'Research':
          return 'We focus on cutting-edge scientific research in physics, AI, and simulations.'
        case 'Innovation':
          return 'Bridging the gap between theory and practice through AR and AI-powered tools.'
        default:
          return ''
      }
    }

    const summarizeQuery = async (query) => {
      try {
        const completion = await openai.chat.completions.create({
          messages: [
            { role: "system", content: "You are a helpful assistant. Summarize the following text into exactly 2-3 words." },
            { role: "user", content: query },
          ],
          model: "deepseek-chat",
          max_tokens: 5,
        })

        const summary = completion.choices[0].message.content
        return summary
      } catch (err) {
        console.error("Error summarizing query:", err)
        if (err.status === 1000) {
          throw new Error("Insufficient balance. Please add funds to your DeepSeek API account.")
        } else {
          throw err
        }
      }
    }

    const answerQuery = async (query) => {
      try {
        const analysis = await openai.chat.completions.create({
          messages: [
            {
              role: "system",
              content: `Analyze the following question and determine if it contains any physics-related components. If it does, extract the physics-related part. If it doesn't, respond with "not physics".`,
            },
            { role: "user", content: query },
          ],
          model: "deepseek-chat",
          max_tokens: 50,
        })

        const analysisResult = analysis.choices[0].message.content.trim()

        if (analysisResult === "not physics") {
          return "I'm primarily trained to answer questions related to physics. Could you ask me something about physics?"
        }

        const completion = await openai.chat.completions.create({
          messages: [
            {
              role: "system",
              content: `You are a physics expert assistant. Focus on the physics-related aspects of the following question and provide a detailed answer. If the question is not entirely about physics, address only the physics part.`,
            },
            { role: "user", content: analysisResult },
          ],
          model: "deepseek-chat",
          max_tokens: 1000,
        })

        return completion.choices[0].message.content
      } catch (err) {
        console.error("Error answering query:", err)
        throw err
      }
    }

    const simulateTyping = (text, callback) => {
      let index = 0
      isTyping.value = true
      const typingInterval = setInterval(() => {
        if (index < text.length) {
          callback(text.slice(0, index + 1))
          index++
        } else {
          clearInterval(typingInterval)
          isTyping.value = false
        }
      }, 30)
    }

    const handleSend = async () => {
      if (input.value.trim() === "") return

      const userMessage = { text: input.value, sender: "user" }
      messages.push(userMessage)
      const currentInput = input.value
      input.value = ""

      try {
        const summary = await summarizeQuery(currentInput)

        const newHistoryItem = {
          query: currentInput,
          summary: summary,
          timestamp: new Date().toISOString(),
        }
        searchHistory.value.unshift(newHistoryItem)

        const answer = await answerQuery(currentInput)

        simulateTyping(answer, (partialText) => {
          const lastMessage = messages[messages.length - 1]
          if (lastMessage && lastMessage.sender === "ai") {
            lastMessage.text = partialText
          } else {
            messages.push({ text: partialText, sender: "ai" })
          }
        })
      } catch (err) {
        const errorMessage = { text: err.message, sender: "ai" }
        messages.push(errorMessage)
      }
    }

    const handleAskAI = () => {
      isChatOpen.value = true
    }

    const handleCloseChat = () => {
      isChatOpen.value = false
      messages.length = 0
      input.value = ""
    }

    const handleHistoryItemClick = (item) => {
      input.value = item.query
      isChatOpen.value = true
    }

    const handleSidebarCollapse = (isCollapsed) => {
      isSidebarCollapsed.value = isCollapsed
    }

    const handleKeyPress = (e) => {
      if (e.key === 'Enter') {
        handleSend()
      }
    }

    // Three.js initialization - Infinite moving particles as background
    const initializeScene = () => {
      let scene, camera, particleSystem, particlesMaterial
      let hue = 0

      try {
        // Scene setup with transparent background
        scene = new THREE.Scene()
        scene.background = null // Remove the dark background
        camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000)
        camera.position.z = 50

        // Renderer setup with transparency
        const renderer = new THREE.WebGLRenderer({
          antialias: true,
          alpha: true // Enable transparency
        })
        renderer.setSize(window.innerWidth, window.innerHeight)
        renderer.setClearColor(0x000000, 0) // Fully transparent background
        rendererRef.value = renderer

        if (mountRef.value) {
          mountRef.value.appendChild(renderer.domElement)
        }

        // Create particles for infinite movement
        const particleCount = 2000 // Increased particle count
        const particles = new THREE.BufferGeometry()
        const positions = new Float32Array(particleCount * 3)
        const velocities = new Float32Array(particleCount * 3)

        // Initialize particles with random positions and velocities
        for (let i = 0; i < particleCount * 3; i += 3) {
          positions[i] = (Math.random() - 0.5) * 300 // Wider spread
          positions[i + 1] = (Math.random() - 0.5) * 300
          positions[i + 2] = (Math.random() - 0.5) * 300

          // More varied velocities for interesting movement
          velocities[i] = (Math.random() - 0.5) * 0.15
          velocities[i + 1] = (Math.random() - 0.5) * 0.15
          velocities[i + 2] = (Math.random() - 0.5) * 0.15
        }

        particles.setAttribute("position", new THREE.BufferAttribute(positions, 3))

        // Enhanced particle material
        particlesMaterial = new THREE.PointsMaterial({
          size: 1.2,
          transparent: true,
          opacity: 0.7,
          blending: THREE.AdditiveBlending,
          sizeAttenuation: true,
          color: 0x4488ff // Initial blue color
        })

        particleSystem = new THREE.Points(particles, particlesMaterial)
        scene.add(particleSystem)

        // Animation loop for infinite movement
        const animate = () => {
          animationFrameId.value = requestAnimationFrame(animate)

          const positions = particleSystem.geometry.attributes.position.array

          for (let i = 0; i < positions.length; i += 3) {
            // Update positions
            positions[i] += velocities[i]
            positions[i + 1] += velocities[i + 1]
            positions[i + 2] += velocities[i + 2]

            // Infinite space - wrap around when particles go too far
            const boundary = 200
            if (Math.abs(positions[i]) > boundary) positions[i] = -positions[i] * 0.9
            if (Math.abs(positions[i + 1]) > boundary) positions[i + 1] = -positions[i + 1] * 0.9
            if (Math.abs(positions[i + 2]) > boundary) positions[i + 2] = -positions[i + 2] * 0.9
          }

          particleSystem.geometry.attributes.position.needsUpdate = true

          // Gentle rotation for dynamic effect
          particleSystem.rotation.x += 0.0003
          particleSystem.rotation.y += 0.0007
          particleSystem.rotation.z += 0.0002

          // Smooth color cycling
          hue += 0.2
          if (hue > 360) hue = 0
          const color = new THREE.Color(`hsl(${hue}, 80%, 65%)`)
          particlesMaterial.color.set(color)

          renderer.render(scene, camera)
        }

        // Handle window resize
        const handleResize = () => {
          const footer = document.querySelector("footer")
          const footerHeight = footer ? footer.offsetHeight : 100
          const newHeight = window.innerHeight - footerHeight
          camera.aspect = window.innerWidth / newHeight
          camera.updateProjectionMatrix()
          renderer.setSize(window.innerWidth, newHeight)
        }

        window.addEventListener("resize", handleResize)
        handleResize()
        animate()
        isLoading.value = false

        // Cleanup function
        return () => {
          window.removeEventListener("resize", handleResize)
          if (animationFrameId.value) {
            cancelAnimationFrame(animationFrameId.value)
          }
          if (mountRef.value && renderer.domElement) {
            mountRef.value.removeChild(renderer.domElement)
          }
          renderer.dispose()
          particles.dispose()
          particlesMaterial.dispose()
          scene.clear()
        }
      } catch (err) {
        error.value = err
        isLoading.value = false
        console.error("Three.js initialization error:", err)
      }
    }

    // Watchers
    watch(searchHistory, (newHistory) => {
      localStorage.setItem('searchHistory', JSON.stringify(newHistory))
    }, { deep: true })

    watch(messages, async () => {
      await nextTick()
      if (chatRef.value) {
        chatRef.value.scrollTop = chatRef.value.scrollHeight
      }
    }, { deep: true })

    // Lifecycle
    onMounted(() => {
      initializeScene()
    })

    onUnmounted(() => {
      if (animationFrameId.value) {
        cancelAnimationFrame(animationFrameId.value)
      }
    })

    return {
      // Refs
      mountRef,
      chatRef,
      isLoading,
      error,
      isChatOpen,
      isSidebarCollapsed,
      isTyping,

      // Reactive data
      messages,
      input,
      searchHistory,
      cards,

      // Methods
      getCardDescription,
      handleSend,
      handleAskAI,
      handleCloseChat,
      handleHistoryItemClick,
      handleSidebarCollapse,
      handleKeyPress,
    }
  }
}
</script>

<style scoped>
.home-container {
  position: relative;
  min-height: 100vh;
  overflow: hidden;
  background: #0a0a0a;
  display: flex;
  flex-direction: column;
}

/* Main content area */
.main-content {
  min-height: 100vh;
  padding: 2rem;
  transition: margin-left 0.3s ease;
  flex: 1;
}

/* Centered content wrapper */
.centered-content {
  max-width: 1200px;
  width: 100%;
  margin: 0 auto;
}

/* Canvas / Particles background */
.canvas-container {
  pointer-events: none;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

/* Sidebar adjustments for responsiveness */
.sidebar-wrapper {
  position: fixed;
  top: 0;
  left: 0;
  height: 100%;
  background: rgba(30, 30, 30, 0.95);
  color: white;
  width: 250px;
  transition: width 0.3s ease, left 0.3s ease;
  z-index: 2;
  overflow-y: auto;
  padding-top: 1rem;
}

/* ===============================
   RESPONSIVE BREAKPOINTS
   =============================== */

/* Tablets (<= 992px) */
@media (max-width: 992px) {
  .main-content {
    padding: 1.5rem;
    margin-left: 200px; /* narrower sidebar */
  }

  .sidebar-wrapper {
    width: 200px;
  }
}

/* Small tablets and large phones (<= 768px) */
@media (max-width: 768px) {
  .main-content {
    margin-left: 0;
    padding: 1rem;
  }

  .sidebar-wrapper {
    position: fixed;
    left: -250px;
    width: 250px;
    transition: left 0.3s ease;
  }

  /* Sidebar toggled open (add a class dynamically with JS or Vue state) */
  .sidebar-wrapper.open {
    left: 0;
  }
}

/* Mobile phones (<= 576px) */
@media (max-width: 576px) {
  .main-content {
    padding: 0.75rem;
  }

  .centered-content {
    padding: 0 0.5rem;
  }

  h3, h4, h2 {
    font-size: 1rem;
  }
}


.loading-text {
  font-size: 1.5rem;
  font-weight: bold;
  text-shadow: 0 0 10px rgba(0, 123, 255, 0.5);
}

/* Enhanced Card Styles with better visibility over particles */
.hover-card {
  transition: all 0.3s ease;
  border: 1px solid rgba(255, 255, 255, 0.3);
  background-color: rgba(15, 15, 15, 0.85) !important;
  backdrop-filter: blur(20px);
  border-radius: 15px;
  position: relative;
  z-index: 2;
}

.hover-card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow:
    0 25px 50px rgba(0, 123, 255, 0.5),
    0 0 0 1px rgba(255, 255, 255, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.2);
  border-color: rgba(0, 123, 255, 0.7);
  background-color: rgba(20, 20, 20, 0.95) !important;
}

.card-title {
  font-size: 1.5rem;
  margin-bottom: 1rem;
  background: linear-gradient(135deg, #ffffff, #a0d8ff);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.7);
}

.card-description {
  font-size: 0.9rem;
  line-height: 1.6;
  color: #e0e0e0 !important;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.6);
}

.welcome-text {
  text-shadow:
    2px 2px 6px rgba(0, 0, 0, 0.9),
    0 0 40px rgba(0, 123, 255, 0.5);
  position: relative;
  z-index: 2;
}

.welcome-text h1 {
  background: linear-gradient(135deg, #ffffff, #66b3ff);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.8));
}

.welcome-text .lead {
  color: #f0f0f0 !important;
  font-size: 1.2rem;
  text-shadow: 0 2px 6px rgba(0, 0, 0, 0.8);
}

/* ================================================
   🌌 ARLab AI Chat Modal — Enhanced Design System
   ================================================= */

/* Base Container */
.ai-chat-modal {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%) scale(1);
  width: 90%;
  max-width: 500px;
  height: 70vh;
  max-height: 650px;
  background: rgba(10, 20, 40, 0.92);
  backdrop-filter: blur(24px) saturate(180%);
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.12);
  box-shadow:
    0 25px 50px rgba(0, 0, 0, 0.55),
    0 0 0 1px rgba(59, 130, 246, 0.1);
  display: flex;
  flex-direction: column;
  overflow: hidden;
  z-index: 1200;
  animation: fadeInUp 0.35s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translate(-50%, -40%) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translate(-50%, -50%) scale(1);
  }
}

/* Inner Structure */
.chat-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

/* Header */
.chat-header {
  background: linear-gradient(135deg, #1e40af 0%, #2563eb 100%);
  padding: 1.25rem 1.5rem;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

.header-content {
  display: flex;
  align-items: center;
  gap: 12px;
}

.header-icon {
  width: 42px;
  height: 42px;
  background: rgba(255, 255, 255, 0.25);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 1.3rem;
  transition: transform 0.3s ease;
}

.header-icon:hover {
  transform: rotate(10deg) scale(1.1);
}

.header-text h5 {
  color: #fff;
  font-weight: 600;
  margin: 0;
  font-size: 1.05rem;
}

.status-text {
  color: rgba(255, 255, 255, 0.85);
  font-size: 0.75rem;
}

.close-btn {
  background: rgba(255, 255, 255, 0.15);
  border: none;
  color: #fff;
  width: 34px;
  height: 34px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.25s ease;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: rotate(90deg);
}

/* Chat Body */
.chat-messages {
  flex: 1;
  padding: 1.25rem 1.5rem;
  background: rgba(30, 41, 59, 0.4);
  overflow-y: auto;
  scroll-behavior: smooth;
}

/* Message Styles */
.message {
  margin-bottom: 1.2rem;
  display: flex;
  align-items: flex-end;
}

.message.user {
  justify-content: flex-end;
}

.message.ai {
  justify-content: flex-start;
}

.message-content {
  max-width: 80%;
  padding: 12px 16px;
  border-radius: 18px;
  word-wrap: break-word;
  font-size: 0.9rem;
  line-height: 1.4;
  transition: background 0.3s ease, transform 0.2s ease;
}

.message.user .message-content {
  background: linear-gradient(135deg, #2563eb 0%, #1e40af 100%);
  color: #fff;
  border-bottom-right-radius: 6px;
}

.message.ai .message-content {
  background: rgba(255, 255, 255, 0.12);
  color: #e2e8f0;
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-bottom-left-radius: 6px;
}

/* Timestamps */
.message-time {
  font-size: 0.7rem;
  opacity: 0.6;
  margin-top: 4px;
}

.message.user .message-time {
  text-align: right;
  color: rgba(255, 255, 255, 0.75);
}

.message.ai .message-time {
  text-align: left;
  color: rgba(226, 232, 240, 0.7);
}

/* Typing Indicator */
.typing-indicator {
  display: flex;
  gap: 5px;
  padding: 8px 0;
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.7);
  animation: typing 1.4s infinite ease-in-out;
}

@keyframes typing {
  0%, 80%, 100% { transform: scale(0.7); opacity: 0.5; }
  40% { transform: scale(1); opacity: 1; }
}

.dot:nth-child(1) { animation-delay: -0.32s; }
.dot:nth-child(2) { animation-delay: -0.16s; }
.dot:nth-child(3) { animation-delay: 0; }

/* Input Section */
.chat-input-container {
  padding: 1.25rem;
  background: rgba(10, 20, 40, 0.9);
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.input-group {
  display: flex;
  gap: 12px;
  align-items: center;
}

.chat-input {
  flex: 1;
  padding: 12px 16px;
  border-radius: 12px;
  border: 1px solid rgba(255, 255, 255, 0.15);
  background: rgba(255, 255, 255, 0.1);
  color: #fff;
  font-size: 0.9rem;
  backdrop-filter: blur(12px);
  transition: border-color 0.3s, background 0.3s;
}

.chat-input:focus {
  outline: none;
  border-color: #3b82f6;
  background: rgba(255, 255, 255, 0.15);
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.25);
}

.chat-input::placeholder {
  color: rgba(255, 255, 255, 0.6);
}

.send-btn {
  background: linear-gradient(135deg, #2563eb 0%, #1d4ed8 100%);
  color: #fff;
  border: none;
  border-radius: 12px;
  padding: 12px 18px;
  cursor: pointer;
  transition: all 0.25s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.send-btn:hover:not(:disabled) {
  background: linear-gradient(135deg, #1d4ed8 0%, #1e3a8a 100%);
  transform: translateY(-1px);
}

.send-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* Input Hint */
.input-hint {
  text-align: center;
  margin-top: 6px;
}

.input-hint small {
  color: rgba(255, 255, 255, 0.5);
  font-size: 0.75rem;
}

/* Scrollbar Customization */
.chat-messages::-webkit-scrollbar {
  width: 6px;
}

.chat-messages::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.05);
  border-radius: 4px;
}

.chat-messages::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.25);
  border-radius: 4px;
}

.chat-messages::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.4);
}

/* ================================
   📱 Responsive Breakpoints
   ================================ */

@media (max-width: 768px) {
  .ai-chat-modal {
    width: 95%;
    height: 80vh;
    max-height: none;
  }

  .chat-header {
    padding: 1rem;
  }

  .chat-messages {
    padding: 1rem;
  }

  .chat-input-container {
    padding: 1rem;
  }

  .message-content {
    max-width: 90%;
  }
}

@media (max-width: 480px) {
  .ai-chat-modal {
    width: 96%;
    border-radius: 16px;
  }

  .header-icon {
    width: 36px;
    height: 36px;
    font-size: 1.1rem;
  }

  .chat-input {
    font-size: 0.85rem;
  }

  .send-btn {
    padding: 10px 14px;
  }
}
</style>

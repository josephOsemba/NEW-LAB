<template>
  <div :class="`sidebar ${isCollapsed ? 'collapsed' : ''}`">
    <!-- Sidebar Header -->
    <div class="sidebar-header">
      <h4 v-if="!isCollapsed">ARLab</h4>
      <button @click="handleCollapse">
        {{ isCollapsed ? '▶' : '◀' }}
      </button>
    </div>

    <!-- Search & Ask AI Section -->
    <div class="search-ask-section">
      <input
        v-if="isSearchActive && !isCollapsed"
        type="text"
        placeholder="Search history..."
        v-model="searchQuery"
        @blur="isSearchActive = false"
        ref="searchInput"
      />
      <button
        v-else-if="!isCollapsed"
        @click="activateSearch"
        class="search-btn"
      >
        <span>🔍</span> <span>Search</span>
      </button>
      <button
        v-else
        @click="activateSearch"
        class="search-btn"
        title="Search"
      >
        <span>🔍</span>
      </button>

      <button
        @click="$emit('ask-ai')"
        class="ask-ai-btn"
        :title="isCollapsed ? 'Ask AI' : ''"
      >
        <span>💬</span> <span v-if="!isCollapsed">Ask AI</span>
      </button>
    </div>

    <!-- History Section -->
    <div class="history-section" v-if="!isCollapsed">
      <h5>History</h5>
      <div v-if="Object.keys(groupedHistory).length > 0">
        <div v-for="[date, items] in Object.entries(groupedHistory)" :key="date">
          <h6 class="history-date">{{ date }}</h6>
          <ul class="history-list">
            <li
              v-for="(item, index) in items"
              :key="index"
              class="history-item"
              @click="$emit('history-item-click', item)"
            >
              {{ item.summary || item.query }}
            </li>
          </ul>
        </div>
      </div>
      <p v-else class="text-muted">No history found</p>
    </div>
    <div v-else class="history-section">
      <!-- Show only icons when collapsed -->
      <div class="history-list">
        <div
          v-for="(item, index) in recentHistory"
          :key="index"
          class="history-item"
          @click="$emit('history-item-click', item)"
          :title="item.summary || item.query"
        >
          💬
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, watch, nextTick } from 'vue'

export default {
  name: 'SideBar',
  props: {
    searchHistory: {
      type: Array,
      default: () => []
    }
  },
  emits: ['ask-ai', 'history-item-click', 'collapse'],
  setup(props, { emit }) {
    const isCollapsed = ref(false)
    const searchQuery = ref('')
    const isSearchActive = ref(false)
    const searchInput = ref(null)

    // Computed
    const filteredHistory = computed(() => {
      return props.searchHistory.filter(item =>
        item.query.toLowerCase().includes(searchQuery.value.toLowerCase())
      )
    })

    const groupedHistory = computed(() => {
      return filteredHistory.value.reduce((acc, item) => {
        const date = new Date(item.timestamp).toLocaleDateString()
        if (!acc[date]) {
          acc[date] = []
        }
        acc[date].push(item)
        return acc
      }, {})
    })

    const recentHistory = computed(() => {
      return props.searchHistory.slice(0, 5) // Show only 5 recent items when collapsed
    })

    // Methods
    const handleCollapse = () => {
      isCollapsed.value = !isCollapsed.value
      emit('collapse', isCollapsed.value)
    }

    const activateSearch = async () => {
      if (isCollapsed.value) {
        // If collapsed, expand first then activate search
        isCollapsed.value = false
        await nextTick()
      }
      isSearchActive.value = true
      await nextTick()
      if (searchInput.value) {
        searchInput.value.focus()
      }
    }

    return {
      isCollapsed,
      searchQuery,
      isSearchActive,
      searchInput,
      filteredHistory,
      groupedHistory,
      recentHistory,
      handleCollapse,
      activateSearch
    }
  }
}
</script>

<style scoped>
.sidebar {
  position: fixed;
  top: 60px;
  left: 0;
  width: 250px;
  height: calc(100vh - 60px);
  background-color: #222;
  color: white;
  padding: 20px;
  overflow-y: auto;
  z-index: 10;
  transition: width 0.3s ease-in-out;
}

.sidebar.collapsed {
  width: 60px;
  padding: 10px;
}

/* Sidebar Header */
.sidebar-header {
  background-color: #23272b;
  text-align: center;
  padding: 15px;
  margin-bottom: 20px;
  border-radius: 8px;
}

.sidebar-header h4 {
  margin: 0;
  font-size: 1.2rem;
  font-weight: bold;
}

.sidebar-header button {
  border: none;
  background: none;
  color: white;
  font-size: 1.2rem;
  cursor: pointer;
}

.sidebar-header button:hover {
  color: #007bff;
}

/* Search Bar & Ask AI Section */
.search-ask-section {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin-bottom: 20px;
}

.search-ask-section input {
  width: 100%;
  padding: 8px;
  border: 1px solid #444;
  border-radius: 4px;
  background-color: #333;
  color: white;
}

.search-ask-section input:focus {
  outline: none;
  border-color: #007bff;
}

.search-ask-section button {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  padding: 8px;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  cursor: pointer;
}

.search-ask-section button:hover {
  background-color: #0056b3;
}

/* History Section */
.history-section {
  margin-top: 20px;
}

.history-section h5 {
  margin-bottom: 10px;
  font-size: 1rem;
  font-weight: bold;
}

.history-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.history-item {
  padding: 10px;
  background-color: #333;
  border-radius: 4px;
  margin-bottom: 8px;
  cursor: pointer;
  transition: background-color 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 8px;
}

.history-item:hover {
  background-color: #444;
}

.history-date {
  color: #ffa500;
  font-weight: bold;
  margin-bottom: 8px;
}

/* Responsive Design */
@media (max-width: 768px) {
  .sidebar {
    width: 60px;
  }

  .sidebar.collapsed {
    width: 60px;
  }

  .sidebar-header h4 {
    display: none;
  }

  .search-ask-section input,
  .search-ask-section button span:not(:first-child),
  .history-section h5,
  .history-item span {
    display: none;
  }

  .search-ask-section button,
  .history-item {
    justify-content: center;
  }
}
</style>

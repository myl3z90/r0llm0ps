<script setup>
import { ref } from 'vue'

const machines = ref([
  {
    id: 1, 
    name: 'Machine One',
    status: 'Running'
  },
  {
    id: 2, 
    name: 'Machine Two',
    status: 'Sleeping'
  },
  {
    id: 3, 
    name: 'Machine Three',
    status: 'Sleeping'
  },
  {
    id: 4, 
    name: 'Machine Four',
    status: 'Sleeping'
  }
])
</script>

<template>
  <h1>Machines</h1>
  <ul>
    <li v-for="machine in machines" :key="machine.id">
      Status: {{ machine.status }} | 
      <router-link :to="{ name: 'detail', params: { id: machine.id }}">
        {{ machine.name }} | 
      </router-link>
      <a href="#">Start</a> | 
      <a href="#">Stop</a>
    </li>
  </ul>
</template>